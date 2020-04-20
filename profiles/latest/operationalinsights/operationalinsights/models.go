// +build go1.9

// Copyright 2020 Microsoft Corporation
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

// This code was auto-generated by:
// github.com/test-repo-arcturus/azure-sdk-for-go/tools/profileBuilder

package operationalinsights

import original "github.com/test-repo-arcturus/azure-sdk-for-go/services/operationalinsights/v1/operationalinsights"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type MetadataColumnDataType = original.MetadataColumnDataType

const (
	Bool     MetadataColumnDataType = original.Bool
	Datetime MetadataColumnDataType = original.Datetime
	Dynamic  MetadataColumnDataType = original.Dynamic
	Int      MetadataColumnDataType = original.Int
	Long     MetadataColumnDataType = original.Long
	Real     MetadataColumnDataType = original.Real
	String   MetadataColumnDataType = original.String
)

type BaseClient = original.BaseClient
type Column = original.Column
type ErrorDetail = original.ErrorDetail
type ErrorInfo = original.ErrorInfo
type ErrorResponse = original.ErrorResponse
type GetClient = original.GetClient
type MetadataApplication = original.MetadataApplication
type MetadataApplicationRelated = original.MetadataApplicationRelated
type MetadataCategory = original.MetadataCategory
type MetadataCategoryRelated = original.MetadataCategoryRelated
type MetadataFunction = original.MetadataFunction
type MetadataFunctionRelated = original.MetadataFunctionRelated
type MetadataPermissions = original.MetadataPermissions
type MetadataPermissionsApplicationsItem = original.MetadataPermissionsApplicationsItem
type MetadataPermissionsResourcesItem = original.MetadataPermissionsResourcesItem
type MetadataPermissionsWorkspacesItem = original.MetadataPermissionsWorkspacesItem
type MetadataQuery = original.MetadataQuery
type MetadataQueryRelated = original.MetadataQueryRelated
type MetadataResourceType = original.MetadataResourceType
type MetadataResourceTypeRelated = original.MetadataResourceTypeRelated
type MetadataResults = original.MetadataResults
type MetadataSolution = original.MetadataSolution
type MetadataSolutionRelated = original.MetadataSolutionRelated
type MetadataTable = original.MetadataTable
type MetadataTableColumnsItem = original.MetadataTableColumnsItem
type MetadataTableRelated = original.MetadataTableRelated
type MetadataWorkspace = original.MetadataWorkspace
type MetadataWorkspaceRelated = original.MetadataWorkspaceRelated
type PostClient = original.PostClient
type QueryBody = original.QueryBody
type QueryClient = original.QueryClient
type QueryResults = original.QueryResults
type Table = original.Table

func New() BaseClient {
	return original.New()
}
func NewGetClient() GetClient {
	return original.NewGetClient()
}
func NewGetClientWithBaseURI(baseURI string) GetClient {
	return original.NewGetClientWithBaseURI(baseURI)
}
func NewPostClient() PostClient {
	return original.NewPostClient()
}
func NewPostClientWithBaseURI(baseURI string) PostClient {
	return original.NewPostClientWithBaseURI(baseURI)
}
func NewQueryClient() QueryClient {
	return original.NewQueryClient()
}
func NewQueryClientWithBaseURI(baseURI string) QueryClient {
	return original.NewQueryClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleMetadataColumnDataTypeValues() []MetadataColumnDataType {
	return original.PossibleMetadataColumnDataTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
