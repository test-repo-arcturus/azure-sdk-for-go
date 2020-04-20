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

package hybriddata

import (
	"context"

	original "github.com/test-repo-arcturus/azure-sdk-for-go/services/hybriddatamanager/mgmt/2016-06-01/hybriddata"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type IsJobCancellable = original.IsJobCancellable

const (
	Cancellable    IsJobCancellable = original.Cancellable
	NotCancellable IsJobCancellable = original.NotCancellable
)

type JobStatus = original.JobStatus

const (
	Cancelled        JobStatus = original.Cancelled
	Cancelling       JobStatus = original.Cancelling
	Failed           JobStatus = original.Failed
	InProgress       JobStatus = original.InProgress
	None             JobStatus = original.None
	Succeeded        JobStatus = original.Succeeded
	WaitingForAction JobStatus = original.WaitingForAction
)

type RunLocation = original.RunLocation

const (
	RunLocationAustraliaeast      RunLocation = original.RunLocationAustraliaeast
	RunLocationAustraliasoutheast RunLocation = original.RunLocationAustraliasoutheast
	RunLocationBrazilsouth        RunLocation = original.RunLocationBrazilsouth
	RunLocationCanadacentral      RunLocation = original.RunLocationCanadacentral
	RunLocationCanadaeast         RunLocation = original.RunLocationCanadaeast
	RunLocationCentralindia       RunLocation = original.RunLocationCentralindia
	RunLocationCentralus          RunLocation = original.RunLocationCentralus
	RunLocationEastasia           RunLocation = original.RunLocationEastasia
	RunLocationEastus             RunLocation = original.RunLocationEastus
	RunLocationEastus2            RunLocation = original.RunLocationEastus2
	RunLocationJapaneast          RunLocation = original.RunLocationJapaneast
	RunLocationJapanwest          RunLocation = original.RunLocationJapanwest
	RunLocationKoreacentral       RunLocation = original.RunLocationKoreacentral
	RunLocationKoreasouth         RunLocation = original.RunLocationKoreasouth
	RunLocationNone               RunLocation = original.RunLocationNone
	RunLocationNorthcentralus     RunLocation = original.RunLocationNorthcentralus
	RunLocationNortheurope        RunLocation = original.RunLocationNortheurope
	RunLocationSouthcentralus     RunLocation = original.RunLocationSouthcentralus
	RunLocationSoutheastasia      RunLocation = original.RunLocationSoutheastasia
	RunLocationSouthindia         RunLocation = original.RunLocationSouthindia
	RunLocationUksouth            RunLocation = original.RunLocationUksouth
	RunLocationUkwest             RunLocation = original.RunLocationUkwest
	RunLocationWestcentralus      RunLocation = original.RunLocationWestcentralus
	RunLocationWesteurope         RunLocation = original.RunLocationWesteurope
	RunLocationWestindia          RunLocation = original.RunLocationWestindia
	RunLocationWestus             RunLocation = original.RunLocationWestus
	RunLocationWestus2            RunLocation = original.RunLocationWestus2
)

type State = original.State

const (
	Disabled  State = original.Disabled
	Enabled   State = original.Enabled
	Supported State = original.Supported
)

type SupportedAlgorithm = original.SupportedAlgorithm

const (
	SupportedAlgorithmNone      SupportedAlgorithm = original.SupportedAlgorithmNone
	SupportedAlgorithmPlainText SupportedAlgorithm = original.SupportedAlgorithmPlainText
	SupportedAlgorithmRSA15     SupportedAlgorithm = original.SupportedAlgorithmRSA15
	SupportedAlgorithmRSAOAEP   SupportedAlgorithm = original.SupportedAlgorithmRSAOAEP
)

type UserConfirmation = original.UserConfirmation

const (
	NotRequired UserConfirmation = original.NotRequired
	Required    UserConfirmation = original.Required
)

type AvailableProviderOperation = original.AvailableProviderOperation
type AvailableProviderOperationDisplay = original.AvailableProviderOperationDisplay
type AvailableProviderOperations = original.AvailableProviderOperations
type AvailableProviderOperationsIterator = original.AvailableProviderOperationsIterator
type AvailableProviderOperationsPage = original.AvailableProviderOperationsPage
type BaseClient = original.BaseClient
type CustomerSecret = original.CustomerSecret
type DataManager = original.DataManager
type DataManagerList = original.DataManagerList
type DataManagerUpdateParameter = original.DataManagerUpdateParameter
type DataManagersClient = original.DataManagersClient
type DataManagersCreateFuture = original.DataManagersCreateFuture
type DataManagersDeleteFuture = original.DataManagersDeleteFuture
type DataManagersUpdateFuture = original.DataManagersUpdateFuture
type DataService = original.DataService
type DataServiceList = original.DataServiceList
type DataServiceListIterator = original.DataServiceListIterator
type DataServiceListPage = original.DataServiceListPage
type DataServiceProperties = original.DataServiceProperties
type DataServicesClient = original.DataServicesClient
type DataStore = original.DataStore
type DataStoreFilter = original.DataStoreFilter
type DataStoreList = original.DataStoreList
type DataStoreListIterator = original.DataStoreListIterator
type DataStoreListPage = original.DataStoreListPage
type DataStoreProperties = original.DataStoreProperties
type DataStoreType = original.DataStoreType
type DataStoreTypeList = original.DataStoreTypeList
type DataStoreTypeListIterator = original.DataStoreTypeListIterator
type DataStoreTypeListPage = original.DataStoreTypeListPage
type DataStoreTypeProperties = original.DataStoreTypeProperties
type DataStoreTypesClient = original.DataStoreTypesClient
type DataStoresClient = original.DataStoresClient
type DataStoresCreateOrUpdateFuture = original.DataStoresCreateOrUpdateFuture
type DataStoresDeleteFuture = original.DataStoresDeleteFuture
type DmsBaseObject = original.DmsBaseObject
type Error = original.Error
type ErrorDetails = original.ErrorDetails
type Job = original.Job
type JobDefinition = original.JobDefinition
type JobDefinitionFilter = original.JobDefinitionFilter
type JobDefinitionList = original.JobDefinitionList
type JobDefinitionListIterator = original.JobDefinitionListIterator
type JobDefinitionListPage = original.JobDefinitionListPage
type JobDefinitionProperties = original.JobDefinitionProperties
type JobDefinitionsClient = original.JobDefinitionsClient
type JobDefinitionsCreateOrUpdateFuture = original.JobDefinitionsCreateOrUpdateFuture
type JobDefinitionsDeleteFuture = original.JobDefinitionsDeleteFuture
type JobDefinitionsRunFuture = original.JobDefinitionsRunFuture
type JobDetails = original.JobDetails
type JobFilter = original.JobFilter
type JobList = original.JobList
type JobListIterator = original.JobListIterator
type JobListPage = original.JobListPage
type JobProperties = original.JobProperties
type JobStages = original.JobStages
type JobsCancelFuture = original.JobsCancelFuture
type JobsClient = original.JobsClient
type JobsResumeFuture = original.JobsResumeFuture
type Key = original.Key
type OperationsClient = original.OperationsClient
type PublicKey = original.PublicKey
type PublicKeyList = original.PublicKeyList
type PublicKeyListIterator = original.PublicKeyListIterator
type PublicKeyListPage = original.PublicKeyListPage
type PublicKeyProperties = original.PublicKeyProperties
type PublicKeysClient = original.PublicKeysClient
type Resource = original.Resource
type RunParameters = original.RunParameters
type Schedule = original.Schedule
type Sku = original.Sku

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAvailableProviderOperationsIterator(page AvailableProviderOperationsPage) AvailableProviderOperationsIterator {
	return original.NewAvailableProviderOperationsIterator(page)
}
func NewAvailableProviderOperationsPage(getNextPage func(context.Context, AvailableProviderOperations) (AvailableProviderOperations, error)) AvailableProviderOperationsPage {
	return original.NewAvailableProviderOperationsPage(getNextPage)
}
func NewDataManagersClient(subscriptionID string) DataManagersClient {
	return original.NewDataManagersClient(subscriptionID)
}
func NewDataManagersClientWithBaseURI(baseURI string, subscriptionID string) DataManagersClient {
	return original.NewDataManagersClientWithBaseURI(baseURI, subscriptionID)
}
func NewDataServiceListIterator(page DataServiceListPage) DataServiceListIterator {
	return original.NewDataServiceListIterator(page)
}
func NewDataServiceListPage(getNextPage func(context.Context, DataServiceList) (DataServiceList, error)) DataServiceListPage {
	return original.NewDataServiceListPage(getNextPage)
}
func NewDataServicesClient(subscriptionID string) DataServicesClient {
	return original.NewDataServicesClient(subscriptionID)
}
func NewDataServicesClientWithBaseURI(baseURI string, subscriptionID string) DataServicesClient {
	return original.NewDataServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewDataStoreListIterator(page DataStoreListPage) DataStoreListIterator {
	return original.NewDataStoreListIterator(page)
}
func NewDataStoreListPage(getNextPage func(context.Context, DataStoreList) (DataStoreList, error)) DataStoreListPage {
	return original.NewDataStoreListPage(getNextPage)
}
func NewDataStoreTypeListIterator(page DataStoreTypeListPage) DataStoreTypeListIterator {
	return original.NewDataStoreTypeListIterator(page)
}
func NewDataStoreTypeListPage(getNextPage func(context.Context, DataStoreTypeList) (DataStoreTypeList, error)) DataStoreTypeListPage {
	return original.NewDataStoreTypeListPage(getNextPage)
}
func NewDataStoreTypesClient(subscriptionID string) DataStoreTypesClient {
	return original.NewDataStoreTypesClient(subscriptionID)
}
func NewDataStoreTypesClientWithBaseURI(baseURI string, subscriptionID string) DataStoreTypesClient {
	return original.NewDataStoreTypesClientWithBaseURI(baseURI, subscriptionID)
}
func NewDataStoresClient(subscriptionID string) DataStoresClient {
	return original.NewDataStoresClient(subscriptionID)
}
func NewDataStoresClientWithBaseURI(baseURI string, subscriptionID string) DataStoresClient {
	return original.NewDataStoresClientWithBaseURI(baseURI, subscriptionID)
}
func NewJobDefinitionListIterator(page JobDefinitionListPage) JobDefinitionListIterator {
	return original.NewJobDefinitionListIterator(page)
}
func NewJobDefinitionListPage(getNextPage func(context.Context, JobDefinitionList) (JobDefinitionList, error)) JobDefinitionListPage {
	return original.NewJobDefinitionListPage(getNextPage)
}
func NewJobDefinitionsClient(subscriptionID string) JobDefinitionsClient {
	return original.NewJobDefinitionsClient(subscriptionID)
}
func NewJobDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) JobDefinitionsClient {
	return original.NewJobDefinitionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewJobListIterator(page JobListPage) JobListIterator {
	return original.NewJobListIterator(page)
}
func NewJobListPage(getNextPage func(context.Context, JobList) (JobList, error)) JobListPage {
	return original.NewJobListPage(getNextPage)
}
func NewJobsClient(subscriptionID string) JobsClient {
	return original.NewJobsClient(subscriptionID)
}
func NewJobsClientWithBaseURI(baseURI string, subscriptionID string) JobsClient {
	return original.NewJobsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPublicKeyListIterator(page PublicKeyListPage) PublicKeyListIterator {
	return original.NewPublicKeyListIterator(page)
}
func NewPublicKeyListPage(getNextPage func(context.Context, PublicKeyList) (PublicKeyList, error)) PublicKeyListPage {
	return original.NewPublicKeyListPage(getNextPage)
}
func NewPublicKeysClient(subscriptionID string) PublicKeysClient {
	return original.NewPublicKeysClient(subscriptionID)
}
func NewPublicKeysClientWithBaseURI(baseURI string, subscriptionID string) PublicKeysClient {
	return original.NewPublicKeysClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleIsJobCancellableValues() []IsJobCancellable {
	return original.PossibleIsJobCancellableValues()
}
func PossibleJobStatusValues() []JobStatus {
	return original.PossibleJobStatusValues()
}
func PossibleRunLocationValues() []RunLocation {
	return original.PossibleRunLocationValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleSupportedAlgorithmValues() []SupportedAlgorithm {
	return original.PossibleSupportedAlgorithmValues()
}
func PossibleUserConfirmationValues() []UserConfirmation {
	return original.PossibleUserConfirmationValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
