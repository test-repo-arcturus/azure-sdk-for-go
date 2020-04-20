package hdinsight

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/test-repo-arcturus/azure-sdk-for-go/services/preview/hdinsight/2018-11-01-preview/hdinsight"

// ApplicationState enumerates the values for application state.
type ApplicationState string

const (
	// ACCEPTED ...
	ACCEPTED ApplicationState = "ACCEPTED"
	// FAILED ...
	FAILED ApplicationState = "FAILED"
	// FINISHED ...
	FINISHED ApplicationState = "FINISHED"
	// FINISHING ...
	FINISHING ApplicationState = "FINISHING"
	// KILLED ...
	KILLED ApplicationState = "KILLED"
	// NEW ...
	NEW ApplicationState = "NEW"
	// NEWSAVING ...
	NEWSAVING ApplicationState = "NEW_SAVING"
	// RUNNING ...
	RUNNING ApplicationState = "RUNNING"
	// SUBMITTED ...
	SUBMITTED ApplicationState = "SUBMITTED"
)

// PossibleApplicationStateValues returns an array of possible values for the ApplicationState const type.
func PossibleApplicationStateValues() []ApplicationState {
	return []ApplicationState{ACCEPTED, FAILED, FINISHED, FINISHING, KILLED, NEW, NEWSAVING, RUNNING, SUBMITTED}
}

// AppState the State of the application.
type AppState struct {
	autorest.Response `json:"-"`
	// State - The State of the application. Possible values include: 'NEW', 'NEWSAVING', 'SUBMITTED', 'ACCEPTED', 'RUNNING', 'FINISHED', 'FINISHING', 'FAILED', 'KILLED'
	State ApplicationState `json:"state,omitempty"`
}

// JobDetailRootJSONObject the object containing the job details.
type JobDetailRootJSONObject struct {
	autorest.Response `json:"-"`
	// Callback - The callback URL, if any.
	Callback interface{} `json:"callback,omitempty"`
	// Completed - The string representing completed status, for example 'done'.
	Completed *string `json:"completed,omitempty"`
	// ExitValue - The job's exit value.
	ExitValue *int32 `json:"exitValue,omitempty"`
	// ID - The job ID.
	ID *string `json:"id,omitempty"`
	// Msg - The message returned.
	Msg interface{} `json:"msg,omitempty"`
	// ParentID - The parent job ID.
	ParentID *string `json:"parentId,omitempty"`
	// PercentComplete - The job completion percentage, for example '75% complete'.
	PercentComplete *string `json:"percentComplete,omitempty"`
	// Profile - The object containing the job profile information.
	Profile *Profile `json:"profile,omitempty"`
	// Status - The object containing the job status information.
	Status *Status `json:"status,omitempty"`
	// User - The user name of the job creator.
	User *string `json:"user,omitempty"`
	// Userargs - The arguments passed in by the user.
	Userargs *Userargs `json:"userargs,omitempty"`
}

// JobID the object with the Job ID.
type JobID struct {
	// ID - The job number.
	ID *int64 `json:"id,omitempty"`
	// JtIdentifier - The jobTracker identifier.
	JtIdentifier *string `json:"jtIdentifier,omitempty"`
}

// JobListJSONObject the List Job operation response.
type JobListJSONObject struct {
	// Detail - The detail of the job.
	Detail *JobDetailRootJSONObject `json:"detail,omitempty"`
	// ID - The Id of the job.
	ID *string `json:"id,omitempty"`
}

// JobOperationsErrorResponse describes the format of Error response.
type JobOperationsErrorResponse struct {
	// Error - Error message indicating why the operation failed.
	Error *string `json:"error,omitempty"`
}

// JobSubmissionJSONResponse the job submission json response.
type JobSubmissionJSONResponse struct {
	autorest.Response `json:"-"`
	// ID - The Id of the created job.
	ID *string `json:"id,omitempty"`
}

// ListJobListJSONObject ...
type ListJobListJSONObject struct {
	autorest.Response `json:"-"`
	Value             *[]JobListJSONObject `json:"value,omitempty"`
}

// Profile the object containing the job profile information.
type Profile struct {
	// JobFile - The job configuration file.
	JobFile *string `json:"jobFile,omitempty"`
	// JobID - The full ID of the job.
	JobID *string `json:"jobId,omitempty"`
	// JobName - The user-specified job name.
	JobName *string `json:"jobName,omitempty"`
	// QueueName - The name of the queue to which the job is submitted.
	QueueName *string `json:"queueName,omitempty"`
	// URL - The link to the web-ui for details of the job.
	URL *string `json:"url,omitempty"`
	// User - The userid of the person who submitted the job.
	User *string `json:"user,omitempty"`
}

// Status gets or sets the object containing the job status information.
type Status struct {
	// CleanupProgress - The progress made on the cleanup.
	CleanupProgress *float64 `json:"cleanupProgress,omitempty"`
	// FailureInfo - The information about any failures that have occurred.
	FailureInfo *string `json:"failureInfo,omitempty"`
	// FinishTime - The time at which the job completed. It is an integer in milliseconds, as a Unix timestamp relative to 1/1/1970 00:00:00.
	FinishTime *int64 `json:"finishTime,omitempty"`
	// HistoryFile - The history file of the job.
	HistoryFile *string `json:"historyFile,omitempty"`
	// JobACLs - The ACLs of the job.
	JobACLs interface{} `json:"jobACLs,omitempty"`
	// JobComplete - Whether or not the job has completed.
	JobComplete *bool `json:"jobComplete,omitempty"`
	// JobFile - The job configuration file.
	JobFile *string `json:"jobFile,omitempty"`
	// JobID - The full ID of the job.
	JobID *string `json:"jobId,omitempty"`
	// JobName - The user-specified job name.
	JobName *string `json:"jobName,omitempty"`
	// JobPriority - The priority of the job.
	JobPriority *string `json:"jobPriority,omitempty"`
	// MapProgress - The progress made on the maps.
	MapProgress *float64 `json:"mapProgress,omitempty"`
	// NeededMem - The amount of memory needed for the job.
	NeededMem *int64 `json:"neededMem,omitempty"`
	// NumReservedSlots - The number of slots reserved.
	NumReservedSlots *int32 `json:"numReservedSlots,omitempty"`
	// NumUsedSlots - The number of slots used for the job.
	NumUsedSlots *int32 `json:"numUsedSlots,omitempty"`
	// Priority - The priority of the job.
	Priority *string `json:"priority,omitempty"`
	// Queue - The job queue name.
	Queue *string `json:"queue,omitempty"`
	// ReduceProgress - The progress made on the reduces.
	ReduceProgress *float64 `json:"reduceProgress,omitempty"`
	// ReservedMem - The amount of memory reserved for the job.
	ReservedMem *int64 `json:"reservedMem,omitempty"`
	// Retired - Whether or not the job has been retired.
	Retired *bool `json:"retired,omitempty"`
	// RunState - The current state of the job.
	RunState *int32 `json:"runState,omitempty"`
	// SchedulingInfo - The information about the scheduling of the job.
	SchedulingInfo *string `json:"schedulingInfo,omitempty"`
	// SetupProgress - The progress made on the setup.
	SetupProgress *float64 `json:"setupProgress,omitempty"`
	// StartTime - The time at which the job started. It is an integer in milliseconds, as a Unix timestamp relative to 1/1/1970 00:00:00.
	StartTime *int64 `json:"startTime,omitempty"`
	// State - The state of the job.
	State *string `json:"state,omitempty"`
	// TrackingURL - The link to the web-ui for details of the job.
	TrackingURL *string `json:"trackingUrl,omitempty"`
	// Uber - Whether job running in uber mode.
	Uber *bool `json:"uber,omitempty"`
	// UsedMem - The amount of memory used by the job.
	UsedMem *int64 `json:"usedMem,omitempty"`
	// Username - The userid of the person who submitted the job.
	Username *string `json:"username,omitempty"`
}

// Userargs gets or sets the object containing the user arguments.
type Userargs struct {
	// Arg - READ-ONLY; The list of args defined by the user.
	Arg *[]string `json:"arg,omitempty"`
	// Callback - The callback URL, if any.
	Callback interface{} `json:"callback,omitempty"`
	// Define - READ-ONLY; The define properties defined by the user.
	Define *[]string `json:"define,omitempty"`
	// Enablelog - Whether or not the user enabled logs.
	Enablelog *string `json:"enablelog,omitempty"`
	// Execute - The query defined by the user.
	Execute *string `json:"execute,omitempty"`
	// File - The query file provided by the user.
	File interface{} `json:"file,omitempty"`
	// Files - The files defined by the user.
	Files interface{} `json:"files,omitempty"`
	// Jar - The JAR file provided by the user.
	Jar *string `json:"jar,omitempty"`
	// Statusdir - The status directory defined by the user.
	Statusdir interface{} `json:"statusdir,omitempty"`
}
