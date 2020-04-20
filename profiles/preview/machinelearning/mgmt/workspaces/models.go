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

package workspaces

import (
	"context"

	original "github.com/test-repo-arcturus/azure-sdk-for-go/services/machinelearning/mgmt/2019-10-01/workspaces"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type State = original.State

const (
	Deleted      State = original.Deleted
	Disabled     State = original.Disabled
	Enabled      State = original.Enabled
	Migrated     State = original.Migrated
	Registered   State = original.Registered
	Unregistered State = original.Unregistered
	Updated      State = original.Updated
)

type WorkspaceType = original.WorkspaceType

const (
	Anonymous    WorkspaceType = original.Anonymous
	Free         WorkspaceType = original.Free
	PaidPremium  WorkspaceType = original.PaidPremium
	PaidStandard WorkspaceType = original.PaidStandard
	Production   WorkspaceType = original.Production
)

type BaseClient = original.BaseClient
type Client = original.Client
type ErrorResponse = original.ErrorResponse
type KeysResponse = original.KeysResponse
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type Properties = original.Properties
type PropertiesUpdateParameters = original.PropertiesUpdateParameters
type Resource = original.Resource
type Sku = original.Sku
type UpdateParameters = original.UpdateParameters
type Workspace = original.Workspace

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return original.NewListResultIterator(page)
}
func NewListResultPage(getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return original.NewListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleWorkspaceTypeValues() []WorkspaceType {
	return original.PossibleWorkspaceTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
