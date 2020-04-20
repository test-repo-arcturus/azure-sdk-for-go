package windowsiotapi

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
	"github.com/test-repo-arcturus/azure-sdk-for-go/services/windowsiot/mgmt/2019-06-01/windowsiot"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result windowsiot.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result windowsiot.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*windowsiot.OperationsClient)(nil)

// ServicesClientAPI contains the set of methods on the ServicesClient type.
type ServicesClientAPI interface {
	CheckDeviceServiceNameAvailability(ctx context.Context, deviceServiceCheckNameAvailabilityParameters windowsiot.DeviceServiceCheckNameAvailabilityParameters) (result windowsiot.DeviceServiceNameAvailabilityInfo, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, deviceName string, deviceService windowsiot.DeviceServiceProperties, ifMatch string) (result windowsiot.DeviceService, err error)
	Delete(ctx context.Context, resourceGroupName string, deviceName string) (result windowsiot.DeviceService, err error)
	Get(ctx context.Context, resourceGroupName string, deviceName string) (result windowsiot.DeviceService, err error)
	List(ctx context.Context) (result windowsiot.DeviceServiceDescriptionListResultPage, err error)
	ListComplete(ctx context.Context) (result windowsiot.DeviceServiceDescriptionListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result windowsiot.DeviceServiceDescriptionListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result windowsiot.DeviceServiceDescriptionListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, deviceName string, deviceService windowsiot.DeviceServiceProperties, ifMatch string) (result windowsiot.DeviceService, err error)
}

var _ ServicesClientAPI = (*windowsiot.ServicesClient)(nil)
