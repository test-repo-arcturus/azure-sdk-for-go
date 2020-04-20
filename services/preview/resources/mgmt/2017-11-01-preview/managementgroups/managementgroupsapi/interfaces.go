package managementgroupsapi

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
	"context"
	"github.com/test-repo-arcturus/azure-sdk-for-go/services/preview/resources/mgmt/2017-11-01-preview/managementgroups"
	"github.com/Azure/go-autorest/autorest"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CreateOrUpdate(ctx context.Context, groupID string, createManagementGroupRequest managementgroups.CreateManagementGroupRequest, cacheControl string) (result managementgroups.ManagementGroup, err error)
	Delete(ctx context.Context, groupID string, cacheControl string) (result autorest.Response, err error)
	Get(ctx context.Context, groupID string, expand string, recurse *bool, cacheControl string) (result managementgroups.ManagementGroup, err error)
	List(ctx context.Context, cacheControl string, skiptoken string) (result managementgroups.ListResultPage, err error)
	ListComplete(ctx context.Context, cacheControl string, skiptoken string) (result managementgroups.ListResultIterator, err error)
	Update(ctx context.Context, groupID string, createManagementGroupRequest managementgroups.CreateManagementGroupRequest, cacheControl string) (result managementgroups.ManagementGroup, err error)
}

var _ ClientAPI = (*managementgroups.Client)(nil)

// SubscriptionsClientAPI contains the set of methods on the SubscriptionsClient type.
type SubscriptionsClientAPI interface {
	Create(ctx context.Context, groupID string, subscriptionID string, cacheControl string) (result autorest.Response, err error)
	Delete(ctx context.Context, groupID string, subscriptionID string, cacheControl string) (result autorest.Response, err error)
}

var _ SubscriptionsClientAPI = (*managementgroups.SubscriptionsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result managementgroups.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result managementgroups.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*managementgroups.OperationsClient)(nil)
