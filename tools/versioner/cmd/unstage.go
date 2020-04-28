// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/test-repo-arcturus/azure-sdk-for-go/tools/internal/files"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/test-repo-arcturus/azure-sdk-for-go/tools/internal/log"
	"github.com/test-repo-arcturus/azure-sdk-for-go/tools/internal/modinfo"
	"github.com/spf13/cobra"
)

func unstageCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "unstage <staging dir>",
		Short: "Creates or updates the latest major version for a package from staged content.",
		Long: `This tool will compare a staged package against its latest major version to detect
breaking changes.  If there are no breaking changes the latest major version is updated
with the staged content.  If there are breaking changes the staged content becomes the
next latest major version and the go.mod file is updated.
The default version for new modules is v1.0.0, and for preview modules is v0.0.0.
`,
		Args: cobra.RangeArgs(1, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			root := args[0]
			startingVer := viper.GetString("starting-version")
			if !modinfo.IsValidModuleVersion(startingVer) {
				return fmt.Errorf("the string '%s' is not a valid module version", startingVer)
			}
			startingVerPreview := viper.GetString("starting-preview-version")
			if !modinfo.IsValidModuleVersion(startingVerPreview) {
				return fmt.Errorf("the string '%s' is not a valid module version", startingVerPreview)
			}
			versionSetting := &VersionSetting{
				InitialVersion:        startingVer,
				InitialVersionPreview: startingVerPreview,
			}
			repoRoot := viper.GetString("gomod-root")
			_, _, err := ExecuteUnstage(root, repoRoot, versionSetting, getTags)
			return err
		},
	}
}

var (
	semverRegex    = regexp.MustCompile(`v\d+\.\d+\.\d+$`)
	versionGoRegex = regexp.MustCompile(`\d+\.\d+\.\d+`)
	verSuffixRegex = regexp.MustCompile(`v\d+`)
)

const (
	goModFilename = "go.mod"
	changeLogName   = "CHANGELOG.md"
	versionFilename = "version.go"
	interfacesName = "interfaces.go"
)

// TagsHookFunc is a func used for get tags from remote
type TagsHookFunc func(root string, tagPrefix string) ([]string, error)

func ExecuteUnstage(s, repoRoot string, versionSetting *VersionSetting, getTagsHook TagsHookFunc) (string, string, error) {
	stage, err := filepath.Abs(s)
	if err != nil {
		return "", "", fmt.Errorf("failed to get absolute path from '%s': %+v", s, err)
	}
	// find target directory, which should just be the parent of the stage directory
	baseline := filepath.Dir(stage)
	// format stage folder first to avoid unnecessary changes detected by apidiff
	// and gofmt will automatically convert CRLF to LF, therefore we have to do the same to both baseline and stage to avoid any differences caused by the line separators
	if err := formatCode(baseline); err != nil {
		return "", "", fmt.Errorf("failed to format baseline directory: %+v", err)
	}
	if err := formatCode(stage); err != nil {
		return "", "", fmt.Errorf("failed to format stage directory: %+v", err)
	}
	log.Infof("Target directory path of stage '%s': %s", stage, baseline)
	log.Infof("Checking if '%s' and '%s' are identical", baseline, stage)
	// first we need to check if the baseline and stage is identical, if no change, we just remove the stage directory and return empty tag
	identical, err := checkIdentical(baseline, stage)
	if err != nil {
		return "", "", fmt.Errorf("failed to check identical: %+v", err)
	}
	if identical {
		// directly remove the stage directory
		log.Infof("Stage '%s' is identical to '%s', removing stage", stage, baseline)
		tag, err := updateIdenticalPackage(baseline, stage, repoRoot, getTagsHook)
		if err != nil {
			return "", "", fmt.Errorf("failed to update identical package: %+v", err)
		}
		return baseline, tag, nil
	}
	log.Infof("Generating changes between '%s' and '%s'", baseline, stage)
	mod, err := modinfo.GetModuleInfo(baseline, stage)
	if err != nil {
		return "", "", fmt.Errorf("failed to create module info: %v", err)
	}
	// despite that to have a major version directory for those major version greater than 1 would have better compatibility,
	// consider that now go 1.12 is not in the supporting list of golang version, we just simply do all the update inplace to reduce complexity
	tag, err := updatePackage(baseline, stage, repoRoot, versionSetting, mod, getTagsHook)
	return baseline, tag, err
}

func checkIdentical(baseline, stage string) (bool, error) {
	// first list all the files and directories in baseline and stage
	// and since stage should be a subdirectory of baseline, we need to escape stage when list all file and directories
	fileListInBaseline, err := listAllFiles(baseline)
	if err != nil {
		return false, err
	}
	// strip out those are in stage directory
	fileListInBaseline = escapeStage(fileListInBaseline, stage)
	fileListInBaseline = escapeSpecialFiles(fileListInBaseline)
	fileListInStage, err := listAllFiles(stage)
	if err != nil {
		return false, err
	}
	fileListInStage = escapeSpecialFiles(fileListInStage)
	if len(fileListInBaseline) != len(fileListInStage) {
		return false, nil
	}
	// filepath.Walk follows the lexical order, therefore the two lists should all be in lexical order
	for i, file := range fileListInBaseline {
		fileInStage := fileListInStage[i]
		same, err := files.DeepCompare(file, fileInStage)
		if err != nil {
			return false, err
		}
		if !same {
			return false, nil
		}
	}
	return true, nil
}

func listAllFiles(root string) ([]string, error) {
	if _, err := os.Stat(root); os.IsNotExist(err) {
		return nil, fmt.Errorf("the root path '%s' does not exist", root)
	}
	var results []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			results = append(results, path)
		}
		return nil
	})
	return results, err
}

func escapeStage(fileList []string, stage string) []string {
	var result []string
	for _, path := range fileList {
		if !strings.HasPrefix(path, stage) {
			result = append(result, path)
		}
	}
	return result
}

func escapeSpecialFiles(fileList []string) []string {
	var result []string
	for _, path := range fileList {
		base := filepath.Base(path)
		if base != goModFilename && base != changeLogName && base != versionFilename && base != interfacesName {
			result = append(result, path)
		}
	}
	return result
}

// updateIdenticalPackage will update the code in baseline directory
// (mainly the version.go, since the version.go might be changed in a legacy version release)
func updateIdenticalPackage(baseline, stage, repoRoot string, getTagsHook TagsHookFunc) (string, error) {
	tagPrefix, err := getTagPrefix(baseline, repoRoot)
	if err != nil {
		return "", fmt.Errorf("failed to get tag prefix: %+v", err)
	}
	tags, err := getTagsHook(baseline, tagPrefix)
	if err != nil {
		return "", fmt.Errorf("failed to list tags: %+v", err)
	}
	latestVersion, err := getLatestSemver(tags, tagPrefix)
	if err != nil {
		return "", fmt.Errorf("failed to get latest version: %+v", err)
	}

	if latestVersion == nil {
		return "", fmt.Errorf("baseline and stage are identical, but do not find any latest version of this module")
	}

	tag := fmt.Sprintf("%s/v%s", tagPrefix, latestVersion.String())
	log.Info("Update mod.go")
	if err := updateGoModFile(stage, tag); err != nil {
		return "", fmt.Errorf("failed to update go.mod file: %+v", err)
	}
	log.Info("Update version.go")
	if err := updateVersionFile(stage, tag); err != nil {
		return "", fmt.Errorf("failed to update version.go file: %+v", err)
	}
	// override the go.mod and version.go in baseline, since the go.mod and version.go file may get overridden by the legacy version release
	if err := os.Rename(filepath.Join(stage, goModFilename), filepath.Join(baseline, goModFilename)); err != nil {
		return "", fmt.Errorf("failed to override go.mod file: %+v", err)
	}
	if err := os.Rename(filepath.Join(stage, versionFilename), filepath.Join(baseline, versionFilename)); err != nil {
		return "", fmt.Errorf("failed to override version.go file: %+v", err)
	}
	// remove the stage directory
	if err := os.RemoveAll(stage); err != nil {
		return "", fmt.Errorf("failed to remove stage directory: %+v", err)
	}

	return tag, nil
}

// updatePackage updates the code in baseline directory from the stage directory, and returns the tag of this new module
func updatePackage(baseline, stage, repoRoot string, versionSetting *VersionSetting, mod modinfo.Provider, getTagsHook TagsHookFunc) (string, error) {
	log.Infof("Updating code base in '%s' from stage '%s'", baseline, stage)
	// get the tag for this module
	tag, err := calculateModuleTag(baseline, versionSetting, repoRoot, mod, getTagsHook)
	if err != nil {
		return "", fmt.Errorf("failed to calculate module tag: %+v", err)
	}
	log.Infof("Tag for stage '%s': %s", stage, tag)
	// use the major version suffix to update the go.mod file
	log.Info("Update mod.go")
	if err := updateGoModFile(stage, tag); err != nil {
		return "", fmt.Errorf("failed to update go.mod file: %+v", err)
	}
	// use the tag to update version.go file
	log.Info("Update version.go")
	if err := updateVersionFile(stage, tag); err != nil {
		return "", fmt.Errorf("failed to update version.go: %v", err)
	}
	// generate CHANGELOG.md
	log.Info("Write CHANGELOG file")
	if err := writeChangelog(stage, mod); err != nil {
		return "", fmt.Errorf("failed to write changelog: %v", err)
	}
	if err := overrideLMVFromStageDirectory(baseline, stage); err != nil {
		return "", fmt.Errorf("failed to override '%s' from stage '%s': %+v", baseline, stage, err)
	}
	return tag, nil
}

func overrideLMVFromStageDirectory(baseline, stage string) error {
	// in our case, stage should always be a child of baseline
	// first move stage to a temp directory outside of baseline, then remove the whole baseline directory, and finally move temp back to baseline
	temp := filepath.Join(filepath.Dir(baseline), "temp")
	if err := os.Rename(stage, temp); err != nil {
		return fmt.Errorf("failed to rename '%s' to '%s': %+v", stage, temp, err)
	}
	if err := os.RemoveAll(baseline); err != nil {
		return fmt.Errorf("failed to delete '%s': %+v", baseline, err)
	}
	if err := os.Rename(temp, baseline); err != nil {
		return fmt.Errorf("failed to rename '%s' to '%s': %+v", temp, baseline, err)
	}
	return nil
}

func updateGoModFile(directory, tag string) error {
	path := filepath.Join(directory, goModFilename)
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("failed to open for read '%s': %+v", directory, err)
	}
	defer file.Close()
	// find if the new module will have a major version suffix (/v2, /v3 etc)
	ver := findVersionSuffixInTag(tag)
	if err := updateGoMod(file, ver); err != nil {
		return fmt.Errorf("failed to update go.mod file: %+v", err)
	}
	return nil
}

// here we only update the version number in version.go, the api-version in User-Agent method will be taken care of
// when this file is generated by autorest
func updateVersionFile(directory, tag string) error {
	version := semverRegex.FindString(tag)
	log.Infof("Updating version.go file in %s with version %s", directory, version)
	// version.go file must exists
	file := filepath.Join(directory, versionFilename)
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return errors.New("version.go file does not exist")
	}
	verFile, err := os.OpenFile(file, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("failed to version.go file: %+v", err)
	}
	defer verFile.Close()
	scanner := bufio.NewScanner(verFile)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	_, err = verFile.Seek(0, io.SeekStart) // set pointer to start of the file
	if err != nil {
		return fmt.Errorf("failed to update version.go file: %+v", err)
	}
	hasTag := false
	for _, line := range lines {
		if !strings.HasPrefix(line, "// ") && versionGoRegex.MatchString(line) {
			line = versionGoRegex.ReplaceAllString(line, version)
		}
		if strings.HasPrefix(line, "// tag: ") {
			line = fmt.Sprintf("// tag: %s", tag)
			hasTag = true
		}
		fmt.Fprintln(verFile, line)
	}
	if !hasTag {
		fmt.Fprintf(verFile, "\n// tag: %s\n", tag)
	}
	return nil
}

// returns the absolute path to the latest major version based on the provided staging directory.
// it's assumed that the staging directory is a subdirectory of the actual package directory.
func findLatestMajorVersion(stage string) (string, error) {
	// example input:
	// ~/work/src/github.com/test-repo-arcturus/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis/stage
	// finds:
	// ~/work/src/github.com/test-repo-arcturus/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis
	// ~/work/src/github.com/test-repo-arcturus/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis/v2
	// returns:
	// ~/work/src/github.com/test-repo-arcturus/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis/v2
	parent := filepath.Dir(stage)
	dirs, err := modinfo.GetModuleSubdirs(parent)
	if err != nil {
		return "", fmt.Errorf("failed to get module subdirs '%s': %v", parent, err)
	}
	// no dirs means this is a v1 package
	if len(dirs) == 0 {
		return parent, nil
	}
	sort.Strings(dirs)
	// last one in the slice is the largest
	v := filepath.Join(parent, dirs[len(dirs)-1])
	log.Infof("Latest major version: %v", v)
	return v, nil
}

// updates the module version inside the go.mod file
func updateGoMod(goMod io.ReadWriteSeeker, newVer string) error {
	if newVer == "" {
		return nil
	}
	scanner := bufio.NewScanner(goMod)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	_, err := goMod.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("failed to seek to start: %v", err)
	}
	for _, line := range lines {
		if strings.Index(line, "module") > -1 {
			if modinfo.HasVersionSuffix(line) {
				line = strings.Replace(line, "/"+modinfo.FindVersionSuffix(line), "/"+newVer, 1)
			} else {
				line = line + "/" + newVer
			}
		}
		fmt.Fprintln(goMod, line)
	}
	return nil
}

// traversal all go source files in the stage folder, and replace the import statement with new ones
func updateImportStatement(stage, currentPath, ver string) error {
	index := strings.Index(currentPath, "github.com")
	if index < 0 {
		return fmt.Errorf("github.com does not find in path '%s', this should never happen", currentPath)
	}
	newImport := strings.ReplaceAll(currentPath[index:], "\\", "/")
	oldImport := newImport[:strings.LastIndex(newImport, "/"+ver)]
	log.Infof("Attempting to replace import statement from '%s' to '%s'", oldImport, newImport)
	files, err := findAllFilesContainImportStatement(stage, oldImport)
	if err != nil {
		return err
	}
	log.Infof("Found %d files with import statement of '%s'", len(files), oldImport)
	log.Debugf("Files: \n%s", strings.Join(files, "\n"))
	err = replaceImportStatement(files, oldImport, newImport)
	if err != nil {
		return err
	}
	return nil
}

func findAllFilesContainImportStatement(path, importStatement string) ([]string, error) {
	files := make([]string, 0) // files stores filenames for those content contained the given import statements
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			// read every line of this file
			file, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("failed to open file '%s': %v", path, err)
			}
			defer file.Close()
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.Index(line, importStatement) > -1 {
					files = append(files, path)
				}
			}
			return nil
		}
		return nil
	})
	return files, err
}

func replaceImportStatement(files []string, oldImport, newImport string) error {
	for _, file := range files {
		err := replaceImportInFile(file, oldImport, newImport)
		if err != nil {
			return fmt.Errorf("failed to preform replace in file '%s'", file)
		}
	}
	return nil
}

func replaceImportInFile(filepath, oldContent, newContent string) error {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to open file '%s': %v", filepath, err)
	}
	content := string(bytes)
	importStatements, err := findImportStatements(content)
	if err != nil {
		return err
	}
	newImportStatements := strings.ReplaceAll(importStatements, oldContent, newContent)
	newFileContent := strings.ReplaceAll(content, importStatements, newImportStatements)
	if err := ioutil.WriteFile(filepath, []byte(newFileContent), 0755); err != nil {
		return err
	}
	return nil
}

func findImportStatements(content string) (string, error) {
	oneLineImport := regexp.MustCompile(`import ".*"`)
	if oneLineImport.MatchString(content) {
		return oneLineImport.FindString(content), nil
	}
	multiLineRegex := `import \(\r?\n(\s*\".*\"\r?\n)+\s*\)`
	multiLineImport := regexp.MustCompile(multiLineRegex)
	if multiLineImport.MatchString(content) {
		return multiLineImport.FindString(content), nil
	}
	return "", fmt.Errorf("failed to match import statement")
}

func writeChangelog(stage string, mod modinfo.Provider) error {
	// don't write a changelog for a new module
	if mod.NewModule() {
		return nil
	}
	rpt := mod.GenerateReport()
	changelog, err := os.Create(filepath.Join(stage, changeLogName))
	if err != nil {
		return fmt.Errorf("failed to create %s: %v", changeLogName, err)
	}
	defer changelog.Close()
	if rpt.IsEmpty() {
		_, err = changelog.WriteString("No changes to exported content compared to the previous release.\n")
		return err
	}
	_, err = changelog.WriteString(rpt.ToMarkdown())
	return err
}

func formatCode(folder string) error {
	c := exec.Command("gofmt", "-w", folder)
	if output, err := c.CombinedOutput(); err != nil {
		return errors.New(string(output))
	}
	return nil
}
