package policyinsightsapi

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
	"github.com/test-repo-arcturus/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-10-17-preview/policyinsights"
	"github.com/Azure/go-autorest/autorest/date"
)

// PolicyEventsClientAPI contains the set of methods on the PolicyEventsClient type.
type PolicyEventsClientAPI interface {
	GetMetadata(ctx context.Context, scope string) (result policyinsights.String, err error)
	ListQueryResultsForManagementGroup(ctx context.Context, managementGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyEventsQueryResults, err error)
	ListQueryResultsForResource(ctx context.Context, resourceID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyEventsQueryResults, err error)
	ListQueryResultsForResourceGroup(ctx context.Context, subscriptionID string, resourceGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyEventsQueryResults, err error)
	ListQueryResultsForSubscription(ctx context.Context, subscriptionID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyEventsQueryResults, err error)
}

var _ PolicyEventsClientAPI = (*policyinsights.PolicyEventsClient)(nil)

// PolicyStatesClientAPI contains the set of methods on the PolicyStatesClient type.
type PolicyStatesClientAPI interface {
	GetMetadata(ctx context.Context, scope string) (result policyinsights.String, err error)
	ListQueryResultsForManagementGroup(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, managementGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyStatesQueryResults, err error)
	ListQueryResultsForResource(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, resourceID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyStatesQueryResults, err error)
	ListQueryResultsForResourceGroup(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, resourceGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyStatesQueryResults, err error)
	ListQueryResultsForSubscription(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyStatesQueryResults, err error)
}

var _ PolicyStatesClientAPI = (*policyinsights.PolicyStatesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result policyinsights.OperationsListResults, err error)
}

var _ OperationsClientAPI = (*policyinsights.OperationsClient)(nil)
