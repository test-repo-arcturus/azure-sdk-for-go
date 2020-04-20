package managednetworkapi

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/test-repo-arcturus/azure-sdk-for-go/services/preview/managednetwork/mgmt/2019-06-01-preview/managednetwork"
)

// ManagedNetworksClientAPI contains the set of methods on the ManagedNetworksClient type.
type ManagedNetworksClientAPI interface {
	CreateOrUpdate(ctx context.Context, managedNetwork managednetwork.ManagedNetwork, resourceGroupName string, managedNetworkName string) (result managednetwork.ManagedNetwork, err error)
	Delete(ctx context.Context, resourceGroupName string, managedNetworkName string) (result managednetwork.ManagedNetworksDeleteFutureType, err error)
	Get(ctx context.Context, resourceGroupName string, managedNetworkName string) (result managednetwork.ManagedNetwork, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, top *int32, skiptoken string) (result managednetwork.ListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, top *int32, skiptoken string) (result managednetwork.ListResultIterator, err error)
	ListBySubscription(ctx context.Context, top *int32, skiptoken string) (result managednetwork.ListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context, top *int32, skiptoken string) (result managednetwork.ListResultIterator, err error)
	Update(ctx context.Context, parameters managednetwork.Update, resourceGroupName string, managedNetworkName string) (result managednetwork.ManagedNetworksUpdateFutureType, err error)
}

var _ ManagedNetworksClientAPI = (*managednetwork.ManagedNetworksClient)(nil)

// ScopeAssignmentsClientAPI contains the set of methods on the ScopeAssignmentsClient type.
type ScopeAssignmentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, parameters managednetwork.ScopeAssignment, scope string, scopeAssignmentName string) (result managednetwork.ScopeAssignment, err error)
	Delete(ctx context.Context, scope string, scopeAssignmentName string) (result autorest.Response, err error)
	Get(ctx context.Context, scope string, scopeAssignmentName string) (result managednetwork.ScopeAssignment, err error)
	List(ctx context.Context, scope string) (result managednetwork.ScopeAssignmentListResultPage, err error)
	ListComplete(ctx context.Context, scope string) (result managednetwork.ScopeAssignmentListResultIterator, err error)
}

var _ ScopeAssignmentsClientAPI = (*managednetwork.ScopeAssignmentsClient)(nil)

// GroupsClientAPI contains the set of methods on the GroupsClient type.
type GroupsClientAPI interface {
	CreateOrUpdate(ctx context.Context, managedNetworkGroup managednetwork.Group, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string) (result managednetwork.GroupsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string) (result managednetwork.GroupsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string) (result managednetwork.Group, err error)
	ListByManagedNetwork(ctx context.Context, resourceGroupName string, managedNetworkName string, top *int32, skiptoken string) (result managednetwork.GroupListResultPage, err error)
	ListByManagedNetworkComplete(ctx context.Context, resourceGroupName string, managedNetworkName string, top *int32, skiptoken string) (result managednetwork.GroupListResultIterator, err error)
}

var _ GroupsClientAPI = (*managednetwork.GroupsClient)(nil)

// PeeringPoliciesClientAPI contains the set of methods on the PeeringPoliciesClient type.
type PeeringPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, managedNetworkPolicy managednetwork.PeeringPolicy, resourceGroupName string, managedNetworkName string, managedNetworkPeeringPolicyName string) (result managednetwork.PeeringPoliciesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkPeeringPolicyName string) (result managednetwork.PeeringPoliciesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkPeeringPolicyName string) (result managednetwork.PeeringPolicy, err error)
	ListByManagedNetwork(ctx context.Context, resourceGroupName string, managedNetworkName string, top *int32, skiptoken string) (result managednetwork.PeeringPolicyListResultPage, err error)
	ListByManagedNetworkComplete(ctx context.Context, resourceGroupName string, managedNetworkName string, top *int32, skiptoken string) (result managednetwork.PeeringPolicyListResultIterator, err error)
}

var _ PeeringPoliciesClientAPI = (*managednetwork.PeeringPoliciesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result managednetwork.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result managednetwork.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*managednetwork.OperationsClient)(nil)
