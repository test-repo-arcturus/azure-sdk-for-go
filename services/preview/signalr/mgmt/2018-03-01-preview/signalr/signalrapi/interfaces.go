package signalrapi

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
	"github.com/test-repo-arcturus/azure-sdk-for-go/services/preview/signalr/mgmt/2018-03-01-preview/signalr"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result signalr.OperationListPage, err error)
	ListComplete(ctx context.Context) (result signalr.OperationListIterator, err error)
}

var _ OperationsClientAPI = (*signalr.OperationsClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CheckNameAvailability(ctx context.Context, location string, parameters *signalr.NameAvailabilityParameters) (result signalr.NameAvailability, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters *signalr.CreateParameters) (result signalr.CreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result signalr.DeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result signalr.ResourceType, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result signalr.ResourceListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result signalr.ResourceListIterator, err error)
	ListBySubscription(ctx context.Context) (result signalr.ResourceListPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result signalr.ResourceListIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, resourceName string) (result signalr.Keys, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, resourceName string, parameters *signalr.RegenerateKeyParameters) (result signalr.RegenerateKeyFuture, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, parameters *signalr.UpdateParameters) (result signalr.UpdateFuture, err error)
}

var _ ClientAPI = (*signalr.Client)(nil)

// UsagesClientAPI contains the set of methods on the UsagesClient type.
type UsagesClientAPI interface {
	List(ctx context.Context, location string) (result signalr.UsageListPage, err error)
	ListComplete(ctx context.Context, location string) (result signalr.UsageListIterator, err error)
}

var _ UsagesClientAPI = (*signalr.UsagesClient)(nil)
