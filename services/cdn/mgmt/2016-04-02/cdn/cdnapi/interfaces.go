package cdnapi

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
	"github.com/test-repo-arcturus/azure-sdk-for-go/services/cdn/mgmt/2016-04-02/cdn"
)

// ProfilesClientAPI contains the set of methods on the ProfilesClient type.
type ProfilesClientAPI interface {
	Create(ctx context.Context, profileName string, profileProperties cdn.ProfileCreateParameters, resourceGroupName string) (result cdn.ProfilesCreateFuture, err error)
	DeleteIfExists(ctx context.Context, profileName string, resourceGroupName string) (result cdn.ProfilesDeleteIfExistsFuture, err error)
	GenerateSsoURI(ctx context.Context, profileName string, resourceGroupName string) (result cdn.SsoURI, err error)
	Get(ctx context.Context, profileName string, resourceGroupName string) (result cdn.Profile, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result cdn.ProfileListResult, err error)
	ListBySubscriptionID(ctx context.Context) (result cdn.ProfileListResult, err error)
	Update(ctx context.Context, profileName string, profileProperties cdn.ProfileUpdateParameters, resourceGroupName string) (result cdn.ProfilesUpdateFuture, err error)
}

var _ ProfilesClientAPI = (*cdn.ProfilesClient)(nil)

// EndpointsClientAPI contains the set of methods on the EndpointsClient type.
type EndpointsClientAPI interface {
	Create(ctx context.Context, endpointName string, endpointProperties cdn.EndpointCreateParameters, profileName string, resourceGroupName string) (result cdn.EndpointsCreateFuture, err error)
	DeleteIfExists(ctx context.Context, endpointName string, profileName string, resourceGroupName string) (result cdn.EndpointsDeleteIfExistsFuture, err error)
	Get(ctx context.Context, endpointName string, profileName string, resourceGroupName string) (result cdn.Endpoint, err error)
	ListByProfile(ctx context.Context, profileName string, resourceGroupName string) (result cdn.EndpointListResult, err error)
	LoadContent(ctx context.Context, endpointName string, contentFilePaths cdn.LoadParameters, profileName string, resourceGroupName string) (result cdn.EndpointsLoadContentFuture, err error)
	PurgeContent(ctx context.Context, endpointName string, contentFilePaths cdn.PurgeParameters, profileName string, resourceGroupName string) (result cdn.EndpointsPurgeContentFuture, err error)
	Start(ctx context.Context, endpointName string, profileName string, resourceGroupName string) (result cdn.EndpointsStartFuture, err error)
	Stop(ctx context.Context, endpointName string, profileName string, resourceGroupName string) (result cdn.EndpointsStopFuture, err error)
	Update(ctx context.Context, endpointName string, endpointProperties cdn.EndpointUpdateParameters, profileName string, resourceGroupName string) (result cdn.EndpointsUpdateFuture, err error)
	ValidateCustomDomain(ctx context.Context, endpointName string, customDomainProperties cdn.ValidateCustomDomainInput, profileName string, resourceGroupName string) (result cdn.ValidateCustomDomainOutput, err error)
}

var _ EndpointsClientAPI = (*cdn.EndpointsClient)(nil)

// OriginsClientAPI contains the set of methods on the OriginsClient type.
type OriginsClientAPI interface {
	Create(ctx context.Context, originName string, originProperties cdn.OriginParameters, endpointName string, profileName string, resourceGroupName string) (result cdn.OriginsCreateFuture, err error)
	DeleteIfExists(ctx context.Context, originName string, endpointName string, profileName string, resourceGroupName string) (result cdn.OriginsDeleteIfExistsFuture, err error)
	Get(ctx context.Context, originName string, endpointName string, profileName string, resourceGroupName string) (result cdn.Origin, err error)
	ListByEndpoint(ctx context.Context, endpointName string, profileName string, resourceGroupName string) (result cdn.OriginListResult, err error)
	Update(ctx context.Context, originName string, originProperties cdn.OriginParameters, endpointName string, profileName string, resourceGroupName string) (result cdn.OriginsUpdateFuture, err error)
}

var _ OriginsClientAPI = (*cdn.OriginsClient)(nil)

// CustomDomainsClientAPI contains the set of methods on the CustomDomainsClient type.
type CustomDomainsClientAPI interface {
	Create(ctx context.Context, customDomainName string, customDomainProperties cdn.CustomDomainParameters, endpointName string, profileName string, resourceGroupName string) (result cdn.CustomDomainsCreateFuture, err error)
	DeleteIfExists(ctx context.Context, customDomainName string, endpointName string, profileName string, resourceGroupName string) (result cdn.CustomDomainsDeleteIfExistsFuture, err error)
	Get(ctx context.Context, customDomainName string, endpointName string, profileName string, resourceGroupName string) (result cdn.CustomDomain, err error)
	ListByEndpoint(ctx context.Context, endpointName string, profileName string, resourceGroupName string) (result cdn.CustomDomainListResult, err error)
	Update(ctx context.Context, customDomainName string, customDomainProperties cdn.CustomDomainParameters, endpointName string, profileName string, resourceGroupName string) (result cdn.ErrorResponse, err error)
}

var _ CustomDomainsClientAPI = (*cdn.CustomDomainsClient)(nil)

// NameAvailabilityClientAPI contains the set of methods on the NameAvailabilityClient type.
type NameAvailabilityClientAPI interface {
	CheckNameAvailability(ctx context.Context, checkNameAvailabilityInput cdn.CheckNameAvailabilityInput) (result cdn.CheckNameAvailabilityOutput, err error)
}

var _ NameAvailabilityClientAPI = (*cdn.NameAvailabilityClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result cdn.OperationListResult, err error)
}

var _ OperationsClientAPI = (*cdn.OperationsClient)(nil)
