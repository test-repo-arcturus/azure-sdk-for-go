package devicesapi

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
	"github.com/test-repo-arcturus/azure-sdk-for-go/services/preview/iothub/mgmt/2018-12-01-preview/devices"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result devices.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result devices.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*devices.OperationsClient)(nil)

// IotHubResourceClientAPI contains the set of methods on the IotHubResourceClient type.
type IotHubResourceClientAPI interface {
	CheckNameAvailability(ctx context.Context, operationInputs devices.OperationInputs) (result devices.IotHubNameAvailabilityInfo, err error)
	CreateEventHubConsumerGroup(ctx context.Context, resourceGroupName string, resourceName string, eventHubEndpointName string, name string) (result devices.EventHubConsumerGroupInfo, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, iotHubDescription devices.IotHubDescription, ifMatch string) (result devices.IotHubResourceCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result devices.IotHubResourceDeleteFuture, err error)
	DeleteEventHubConsumerGroup(ctx context.Context, resourceGroupName string, resourceName string, eventHubEndpointName string, name string) (result autorest.Response, err error)
	ExportDevices(ctx context.Context, resourceGroupName string, resourceName string, exportDevicesParameters devices.ExportDevicesRequest) (result devices.JobResponse, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result devices.IotHubDescription, err error)
	GetEndpointHealth(ctx context.Context, resourceGroupName string, iotHubName string) (result devices.EndpointHealthDataListResultPage, err error)
	GetEndpointHealthComplete(ctx context.Context, resourceGroupName string, iotHubName string) (result devices.EndpointHealthDataListResultIterator, err error)
	GetEventHubConsumerGroup(ctx context.Context, resourceGroupName string, resourceName string, eventHubEndpointName string, name string) (result devices.EventHubConsumerGroupInfo, err error)
	GetJob(ctx context.Context, resourceGroupName string, resourceName string, jobID string) (result devices.JobResponse, err error)
	GetKeysForKeyName(ctx context.Context, resourceGroupName string, resourceName string, keyName string) (result devices.SharedAccessSignatureAuthorizationRule, err error)
	GetQuotaMetrics(ctx context.Context, resourceGroupName string, resourceName string) (result devices.IotHubQuotaMetricInfoListResultPage, err error)
	GetQuotaMetricsComplete(ctx context.Context, resourceGroupName string, resourceName string) (result devices.IotHubQuotaMetricInfoListResultIterator, err error)
	GetStats(ctx context.Context, resourceGroupName string, resourceName string) (result devices.RegistryStatistics, err error)
	GetValidSkus(ctx context.Context, resourceGroupName string, resourceName string) (result devices.IotHubSkuDescriptionListResultPage, err error)
	GetValidSkusComplete(ctx context.Context, resourceGroupName string, resourceName string) (result devices.IotHubSkuDescriptionListResultIterator, err error)
	ImportDevices(ctx context.Context, resourceGroupName string, resourceName string, importDevicesParameters devices.ImportDevicesRequest) (result devices.JobResponse, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result devices.IotHubDescriptionListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result devices.IotHubDescriptionListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result devices.IotHubDescriptionListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result devices.IotHubDescriptionListResultIterator, err error)
	ListEventHubConsumerGroups(ctx context.Context, resourceGroupName string, resourceName string, eventHubEndpointName string) (result devices.EventHubConsumerGroupsListResultPage, err error)
	ListEventHubConsumerGroupsComplete(ctx context.Context, resourceGroupName string, resourceName string, eventHubEndpointName string) (result devices.EventHubConsumerGroupsListResultIterator, err error)
	ListJobs(ctx context.Context, resourceGroupName string, resourceName string) (result devices.JobResponseListResultPage, err error)
	ListJobsComplete(ctx context.Context, resourceGroupName string, resourceName string) (result devices.JobResponseListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, resourceName string) (result devices.SharedAccessSignatureAuthorizationRuleListResultPage, err error)
	ListKeysComplete(ctx context.Context, resourceGroupName string, resourceName string) (result devices.SharedAccessSignatureAuthorizationRuleListResultIterator, err error)
	TestAllRoutes(ctx context.Context, input devices.TestAllRoutesInput, iotHubName string, resourceGroupName string) (result devices.TestAllRoutesResult, err error)
	TestRoute(ctx context.Context, input devices.TestRouteInput, iotHubName string, resourceGroupName string) (result devices.TestRouteResult, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, iotHubTags devices.TagsResource) (result devices.IotHubResourceUpdateFuture, err error)
}

var _ IotHubResourceClientAPI = (*devices.IotHubResourceClient)(nil)

// ResourceProviderCommonClientAPI contains the set of methods on the ResourceProviderCommonClient type.
type ResourceProviderCommonClientAPI interface {
	GetSubscriptionQuota(ctx context.Context) (result devices.UserSubscriptionQuotaListResult, err error)
}

var _ ResourceProviderCommonClientAPI = (*devices.ResourceProviderCommonClient)(nil)

// CertificatesClientAPI contains the set of methods on the CertificatesClient type.
type CertificatesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, certificateName string, certificateDescription devices.CertificateBodyDescription, ifMatch string) (result devices.CertificateDescription, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string, certificateName string, ifMatch string) (result autorest.Response, err error)
	GenerateVerificationCode(ctx context.Context, resourceGroupName string, resourceName string, certificateName string, ifMatch string) (result devices.CertificateWithNonceDescription, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string, certificateName string) (result devices.CertificateDescription, err error)
	ListByIotHub(ctx context.Context, resourceGroupName string, resourceName string) (result devices.CertificateListDescription, err error)
	Verify(ctx context.Context, resourceGroupName string, resourceName string, certificateName string, certificateVerificationBody devices.CertificateVerificationDescription, ifMatch string) (result devices.CertificateDescription, err error)
}

var _ CertificatesClientAPI = (*devices.CertificatesClient)(nil)
