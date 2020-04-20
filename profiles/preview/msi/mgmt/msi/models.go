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

package msi

import (
	"context"

	original "github.com/test-repo-arcturus/azure-sdk-for-go/services/msi/mgmt/2018-11-30/msi"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type Identity = original.Identity
type IdentityUpdate = original.IdentityUpdate
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type SystemAssignedIdentitiesClient = original.SystemAssignedIdentitiesClient
type SystemAssignedIdentity = original.SystemAssignedIdentity
type SystemAssignedIdentityProperties = original.SystemAssignedIdentityProperties
type TrackedResource = original.TrackedResource
type UserAssignedIdentitiesClient = original.UserAssignedIdentitiesClient
type UserAssignedIdentitiesListResult = original.UserAssignedIdentitiesListResult
type UserAssignedIdentitiesListResultIterator = original.UserAssignedIdentitiesListResultIterator
type UserAssignedIdentitiesListResultPage = original.UserAssignedIdentitiesListResultPage
type UserAssignedIdentityProperties = original.UserAssignedIdentityProperties

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSystemAssignedIdentitiesClient(subscriptionID string) SystemAssignedIdentitiesClient {
	return original.NewSystemAssignedIdentitiesClient(subscriptionID)
}
func NewSystemAssignedIdentitiesClientWithBaseURI(baseURI string, subscriptionID string) SystemAssignedIdentitiesClient {
	return original.NewSystemAssignedIdentitiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewUserAssignedIdentitiesClient(subscriptionID string) UserAssignedIdentitiesClient {
	return original.NewUserAssignedIdentitiesClient(subscriptionID)
}
func NewUserAssignedIdentitiesClientWithBaseURI(baseURI string, subscriptionID string) UserAssignedIdentitiesClient {
	return original.NewUserAssignedIdentitiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewUserAssignedIdentitiesListResultIterator(page UserAssignedIdentitiesListResultPage) UserAssignedIdentitiesListResultIterator {
	return original.NewUserAssignedIdentitiesListResultIterator(page)
}
func NewUserAssignedIdentitiesListResultPage(getNextPage func(context.Context, UserAssignedIdentitiesListResult) (UserAssignedIdentitiesListResult, error)) UserAssignedIdentitiesListResultPage {
	return original.NewUserAssignedIdentitiesListResultPage(getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
