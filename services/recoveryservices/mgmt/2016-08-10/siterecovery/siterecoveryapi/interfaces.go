package siterecoveryapi

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
	"github.com/test-repo-arcturus/azure-sdk-for-go/services/recoveryservices/mgmt/2016-08-10/siterecovery"
)

// ReplicationVaultHealthClientAPI contains the set of methods on the ReplicationVaultHealthClient type.
type ReplicationVaultHealthClientAPI interface {
	Get(ctx context.Context) (result siterecovery.VaultHealthDetails, err error)
}

var _ ReplicationVaultHealthClientAPI = (*siterecovery.ReplicationVaultHealthClient)(nil)

// ReplicationProtectedItemsClientAPI contains the set of methods on the ReplicationProtectedItemsClient type.
type ReplicationProtectedItemsClientAPI interface {
	ApplyRecoveryPoint(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string, applyRecoveryPointInput siterecovery.ApplyRecoveryPointInput) (result siterecovery.ReplicationProtectedItemsApplyRecoveryPointFuture, err error)
	Create(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string, input siterecovery.EnableProtectionInput) (result siterecovery.ReplicationProtectedItemsCreateFuture, err error)
	Delete(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string, disableProtectionInput siterecovery.DisableProtectionInput) (result siterecovery.ReplicationProtectedItemsDeleteFuture, err error)
	FailoverCommit(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string) (result siterecovery.ReplicationProtectedItemsFailoverCommitFuture, err error)
	Get(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string) (result siterecovery.ReplicationProtectedItem, err error)
	List(ctx context.Context, skipToken string, filter string) (result siterecovery.ReplicationProtectedItemCollectionPage, err error)
	ListComplete(ctx context.Context, skipToken string, filter string) (result siterecovery.ReplicationProtectedItemCollectionIterator, err error)
	ListByReplicationProtectionContainers(ctx context.Context, fabricName string, protectionContainerName string) (result siterecovery.ReplicationProtectedItemCollectionPage, err error)
	ListByReplicationProtectionContainersComplete(ctx context.Context, fabricName string, protectionContainerName string) (result siterecovery.ReplicationProtectedItemCollectionIterator, err error)
	PlannedFailover(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string, failoverInput siterecovery.PlannedFailoverInput) (result siterecovery.ReplicationProtectedItemsPlannedFailoverFuture, err error)
	Purge(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string) (result siterecovery.ReplicationProtectedItemsPurgeFuture, err error)
	RepairReplication(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string) (result siterecovery.ReplicationProtectedItemsRepairReplicationFuture, err error)
	Reprotect(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string, rrInput siterecovery.ReverseReplicationInput) (result siterecovery.ReplicationProtectedItemsReprotectFuture, err error)
	TestFailover(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string, failoverInput siterecovery.TestFailoverInput) (result siterecovery.ReplicationProtectedItemsTestFailoverFuture, err error)
	TestFailoverCleanup(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string, cleanupInput siterecovery.TestFailoverCleanupInput) (result siterecovery.ReplicationProtectedItemsTestFailoverCleanupFuture, err error)
	UnplannedFailover(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string, failoverInput siterecovery.UnplannedFailoverInput) (result siterecovery.ReplicationProtectedItemsUnplannedFailoverFuture, err error)
	Update(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string, updateProtectionInput siterecovery.UpdateReplicationProtectedItemInput) (result siterecovery.ReplicationProtectedItemsUpdateFuture, err error)
	UpdateMobilityService(ctx context.Context, fabricName string, protectionContainerName string, replicationProtectedItemName string, updateMobilityServiceRequest siterecovery.UpdateMobilityServiceRequest) (result siterecovery.ReplicationProtectedItemsUpdateMobilityServiceFuture, err error)
}

var _ ReplicationProtectedItemsClientAPI = (*siterecovery.ReplicationProtectedItemsClient)(nil)

// ReplicationNetworkMappingsClientAPI contains the set of methods on the ReplicationNetworkMappingsClient type.
type ReplicationNetworkMappingsClientAPI interface {
	Create(ctx context.Context, fabricName string, networkName string, networkMappingName string, input siterecovery.CreateNetworkMappingInput) (result siterecovery.ReplicationNetworkMappingsCreateFuture, err error)
	Delete(ctx context.Context, fabricName string, networkName string, networkMappingName string) (result siterecovery.ReplicationNetworkMappingsDeleteFuture, err error)
	Get(ctx context.Context, fabricName string, networkName string, networkMappingName string) (result siterecovery.NetworkMapping, err error)
	List(ctx context.Context) (result siterecovery.NetworkMappingCollectionPage, err error)
	ListComplete(ctx context.Context) (result siterecovery.NetworkMappingCollectionIterator, err error)
	ListByReplicationNetworks(ctx context.Context, fabricName string, networkName string) (result siterecovery.NetworkMappingCollectionPage, err error)
	ListByReplicationNetworksComplete(ctx context.Context, fabricName string, networkName string) (result siterecovery.NetworkMappingCollectionIterator, err error)
	Update(ctx context.Context, fabricName string, networkName string, networkMappingName string, input siterecovery.UpdateNetworkMappingInput) (result siterecovery.ReplicationNetworkMappingsUpdateFuture, err error)
}

var _ ReplicationNetworkMappingsClientAPI = (*siterecovery.ReplicationNetworkMappingsClient)(nil)

// ReplicationFabricsClientAPI contains the set of methods on the ReplicationFabricsClient type.
type ReplicationFabricsClientAPI interface {
	CheckConsistency(ctx context.Context, fabricName string) (result siterecovery.ReplicationFabricsCheckConsistencyFuture, err error)
	Create(ctx context.Context, fabricName string, input siterecovery.FabricCreationInput) (result siterecovery.ReplicationFabricsCreateFuture, err error)
	Delete(ctx context.Context, fabricName string) (result siterecovery.ReplicationFabricsDeleteFuture, err error)
	Get(ctx context.Context, fabricName string) (result siterecovery.Fabric, err error)
	List(ctx context.Context) (result siterecovery.FabricCollectionPage, err error)
	ListComplete(ctx context.Context) (result siterecovery.FabricCollectionIterator, err error)
	MigrateToAad(ctx context.Context, fabricName string) (result siterecovery.ReplicationFabricsMigrateToAadFuture, err error)
	Purge(ctx context.Context, fabricName string) (result siterecovery.ReplicationFabricsPurgeFuture, err error)
	ReassociateGateway(ctx context.Context, fabricName string, failoverProcessServerRequest siterecovery.FailoverProcessServerRequest) (result siterecovery.ReplicationFabricsReassociateGatewayFuture, err error)
	RenewCertificate(ctx context.Context, fabricName string, renewCertificate siterecovery.RenewCertificateInput) (result siterecovery.ReplicationFabricsRenewCertificateFuture, err error)
}

var _ ReplicationFabricsClientAPI = (*siterecovery.ReplicationFabricsClient)(nil)

// ReplicationvCentersClientAPI contains the set of methods on the ReplicationvCentersClient type.
type ReplicationvCentersClientAPI interface {
	Create(ctx context.Context, fabricName string, vCenterName string, addVCenterRequest siterecovery.AddVCenterRequest) (result siterecovery.ReplicationvCentersCreateFuture, err error)
	Delete(ctx context.Context, fabricName string, vCenterName string) (result siterecovery.ReplicationvCentersDeleteFuture, err error)
	Get(ctx context.Context, fabricName string, vCenterName string) (result siterecovery.VCenter, err error)
	List(ctx context.Context) (result siterecovery.VCenterCollectionPage, err error)
	ListComplete(ctx context.Context) (result siterecovery.VCenterCollectionIterator, err error)
	ListByReplicationFabrics(ctx context.Context, fabricName string) (result siterecovery.VCenterCollectionPage, err error)
	ListByReplicationFabricsComplete(ctx context.Context, fabricName string) (result siterecovery.VCenterCollectionIterator, err error)
	Update(ctx context.Context, fabricName string, vCenterName string, updateVCenterRequest siterecovery.UpdateVCenterRequest) (result siterecovery.ReplicationvCentersUpdateFuture, err error)
}

var _ ReplicationvCentersClientAPI = (*siterecovery.ReplicationvCentersClient)(nil)

// ReplicationStorageClassificationMappingsClientAPI contains the set of methods on the ReplicationStorageClassificationMappingsClient type.
type ReplicationStorageClassificationMappingsClientAPI interface {
	Create(ctx context.Context, fabricName string, storageClassificationName string, storageClassificationMappingName string, pairingInput siterecovery.StorageClassificationMappingInput) (result siterecovery.ReplicationStorageClassificationMappingsCreateFuture, err error)
	Delete(ctx context.Context, fabricName string, storageClassificationName string, storageClassificationMappingName string) (result siterecovery.ReplicationStorageClassificationMappingsDeleteFuture, err error)
	Get(ctx context.Context, fabricName string, storageClassificationName string, storageClassificationMappingName string) (result siterecovery.StorageClassificationMapping, err error)
	List(ctx context.Context) (result siterecovery.StorageClassificationMappingCollectionPage, err error)
	ListComplete(ctx context.Context) (result siterecovery.StorageClassificationMappingCollectionIterator, err error)
	ListByReplicationStorageClassifications(ctx context.Context, fabricName string, storageClassificationName string) (result siterecovery.StorageClassificationMappingCollectionPage, err error)
	ListByReplicationStorageClassificationsComplete(ctx context.Context, fabricName string, storageClassificationName string) (result siterecovery.StorageClassificationMappingCollectionIterator, err error)
}

var _ ReplicationStorageClassificationMappingsClientAPI = (*siterecovery.ReplicationStorageClassificationMappingsClient)(nil)

// ReplicationStorageClassificationsClientAPI contains the set of methods on the ReplicationStorageClassificationsClient type.
type ReplicationStorageClassificationsClientAPI interface {
	Get(ctx context.Context, fabricName string, storageClassificationName string) (result siterecovery.StorageClassification, err error)
	List(ctx context.Context) (result siterecovery.StorageClassificationCollectionPage, err error)
	ListComplete(ctx context.Context) (result siterecovery.StorageClassificationCollectionIterator, err error)
	ListByReplicationFabrics(ctx context.Context, fabricName string) (result siterecovery.StorageClassificationCollectionPage, err error)
	ListByReplicationFabricsComplete(ctx context.Context, fabricName string) (result siterecovery.StorageClassificationCollectionIterator, err error)
}

var _ ReplicationStorageClassificationsClientAPI = (*siterecovery.ReplicationStorageClassificationsClient)(nil)

// ReplicationRecoveryServicesProvidersClientAPI contains the set of methods on the ReplicationRecoveryServicesProvidersClient type.
type ReplicationRecoveryServicesProvidersClientAPI interface {
	Delete(ctx context.Context, fabricName string, providerName string) (result siterecovery.ReplicationRecoveryServicesProvidersDeleteFuture, err error)
	Get(ctx context.Context, fabricName string, providerName string) (result siterecovery.RecoveryServicesProvider, err error)
	List(ctx context.Context) (result siterecovery.RecoveryServicesProviderCollectionPage, err error)
	ListComplete(ctx context.Context) (result siterecovery.RecoveryServicesProviderCollectionIterator, err error)
	ListByReplicationFabrics(ctx context.Context, fabricName string) (result siterecovery.RecoveryServicesProviderCollectionPage, err error)
	ListByReplicationFabricsComplete(ctx context.Context, fabricName string) (result siterecovery.RecoveryServicesProviderCollectionIterator, err error)
	Purge(ctx context.Context, fabricName string, providerName string) (result siterecovery.ReplicationRecoveryServicesProvidersPurgeFuture, err error)
	RefreshProvider(ctx context.Context, fabricName string, providerName string) (result siterecovery.ReplicationRecoveryServicesProvidersRefreshProviderFuture, err error)
}

var _ ReplicationRecoveryServicesProvidersClientAPI = (*siterecovery.ReplicationRecoveryServicesProvidersClient)(nil)

// RecoveryPointsClientAPI contains the set of methods on the RecoveryPointsClient type.
type RecoveryPointsClientAPI interface {
	Get(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string, recoveryPointName string) (result siterecovery.RecoveryPoint, err error)
	ListByReplicationProtectedItems(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string) (result siterecovery.RecoveryPointCollectionPage, err error)
	ListByReplicationProtectedItemsComplete(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string) (result siterecovery.RecoveryPointCollectionIterator, err error)
}

var _ RecoveryPointsClientAPI = (*siterecovery.RecoveryPointsClient)(nil)

// ReplicationRecoveryPlansClientAPI contains the set of methods on the ReplicationRecoveryPlansClient type.
type ReplicationRecoveryPlansClientAPI interface {
	Create(ctx context.Context, recoveryPlanName string, input siterecovery.CreateRecoveryPlanInput) (result siterecovery.ReplicationRecoveryPlansCreateFuture, err error)
	Delete(ctx context.Context, recoveryPlanName string) (result siterecovery.ReplicationRecoveryPlansDeleteFuture, err error)
	FailoverCommit(ctx context.Context, recoveryPlanName string) (result siterecovery.ReplicationRecoveryPlansFailoverCommitFuture, err error)
	Get(ctx context.Context, recoveryPlanName string) (result siterecovery.RecoveryPlan, err error)
	List(ctx context.Context) (result siterecovery.RecoveryPlanCollectionPage, err error)
	ListComplete(ctx context.Context) (result siterecovery.RecoveryPlanCollectionIterator, err error)
	PlannedFailover(ctx context.Context, recoveryPlanName string, input siterecovery.RecoveryPlanPlannedFailoverInput) (result siterecovery.ReplicationRecoveryPlansPlannedFailoverFuture, err error)
	Reprotect(ctx context.Context, recoveryPlanName string) (result siterecovery.ReplicationRecoveryPlansReprotectFuture, err error)
	TestFailover(ctx context.Context, recoveryPlanName string, input siterecovery.RecoveryPlanTestFailoverInput) (result siterecovery.ReplicationRecoveryPlansTestFailoverFuture, err error)
	TestFailoverCleanup(ctx context.Context, recoveryPlanName string, input siterecovery.RecoveryPlanTestFailoverCleanupInput) (result siterecovery.ReplicationRecoveryPlansTestFailoverCleanupFuture, err error)
	UnplannedFailover(ctx context.Context, recoveryPlanName string, input siterecovery.RecoveryPlanUnplannedFailoverInput) (result siterecovery.ReplicationRecoveryPlansUnplannedFailoverFuture, err error)
	Update(ctx context.Context, recoveryPlanName string, input siterecovery.UpdateRecoveryPlanInput) (result siterecovery.ReplicationRecoveryPlansUpdateFuture, err error)
}

var _ ReplicationRecoveryPlansClientAPI = (*siterecovery.ReplicationRecoveryPlansClient)(nil)

// ReplicationProtectionContainersClientAPI contains the set of methods on the ReplicationProtectionContainersClient type.
type ReplicationProtectionContainersClientAPI interface {
	Create(ctx context.Context, fabricName string, protectionContainerName string, creationInput siterecovery.CreateProtectionContainerInput) (result siterecovery.ReplicationProtectionContainersCreateFuture, err error)
	Delete(ctx context.Context, fabricName string, protectionContainerName string) (result siterecovery.ReplicationProtectionContainersDeleteFuture, err error)
	DiscoverProtectableItem(ctx context.Context, fabricName string, protectionContainerName string, discoverProtectableItemRequest siterecovery.DiscoverProtectableItemRequest) (result siterecovery.ReplicationProtectionContainersDiscoverProtectableItemFuture, err error)
	Get(ctx context.Context, fabricName string, protectionContainerName string) (result siterecovery.ProtectionContainer, err error)
	List(ctx context.Context) (result siterecovery.ProtectionContainerCollectionPage, err error)
	ListComplete(ctx context.Context) (result siterecovery.ProtectionContainerCollectionIterator, err error)
	ListByReplicationFabrics(ctx context.Context, fabricName string) (result siterecovery.ProtectionContainerCollectionPage, err error)
	ListByReplicationFabricsComplete(ctx context.Context, fabricName string) (result siterecovery.ProtectionContainerCollectionIterator, err error)
	SwitchProtection(ctx context.Context, fabricName string, protectionContainerName string, switchInput siterecovery.SwitchProtectionInput) (result siterecovery.ReplicationProtectionContainersSwitchProtectionFuture, err error)
}

var _ ReplicationProtectionContainersClientAPI = (*siterecovery.ReplicationProtectionContainersClient)(nil)

// ReplicationProtectionContainerMappingsClientAPI contains the set of methods on the ReplicationProtectionContainerMappingsClient type.
type ReplicationProtectionContainerMappingsClientAPI interface {
	Create(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, creationInput siterecovery.CreateProtectionContainerMappingInput) (result siterecovery.ReplicationProtectionContainerMappingsCreateFuture, err error)
	Delete(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, removalInput siterecovery.RemoveProtectionContainerMappingInput) (result siterecovery.ReplicationProtectionContainerMappingsDeleteFuture, err error)
	Get(ctx context.Context, fabricName string, protectionContainerName string, mappingName string) (result siterecovery.ProtectionContainerMapping, err error)
	List(ctx context.Context) (result siterecovery.ProtectionContainerMappingCollectionPage, err error)
	ListComplete(ctx context.Context) (result siterecovery.ProtectionContainerMappingCollectionIterator, err error)
	ListByReplicationProtectionContainers(ctx context.Context, fabricName string, protectionContainerName string) (result siterecovery.ProtectionContainerMappingCollectionPage, err error)
	ListByReplicationProtectionContainersComplete(ctx context.Context, fabricName string, protectionContainerName string) (result siterecovery.ProtectionContainerMappingCollectionIterator, err error)
	Purge(ctx context.Context, fabricName string, protectionContainerName string, mappingName string) (result siterecovery.ReplicationProtectionContainerMappingsPurgeFuture, err error)
}

var _ ReplicationProtectionContainerMappingsClientAPI = (*siterecovery.ReplicationProtectionContainerMappingsClient)(nil)

// ReplicationProtectableItemsClientAPI contains the set of methods on the ReplicationProtectableItemsClient type.
type ReplicationProtectableItemsClientAPI interface {
	Get(ctx context.Context, fabricName string, protectionContainerName string, protectableItemName string) (result siterecovery.ProtectableItem, err error)
	ListByReplicationProtectionContainers(ctx context.Context, fabricName string, protectionContainerName string) (result siterecovery.ProtectableItemCollectionPage, err error)
	ListByReplicationProtectionContainersComplete(ctx context.Context, fabricName string, protectionContainerName string) (result siterecovery.ProtectableItemCollectionIterator, err error)
}

var _ ReplicationProtectableItemsClientAPI = (*siterecovery.ReplicationProtectableItemsClient)(nil)

// ReplicationPoliciesClientAPI contains the set of methods on the ReplicationPoliciesClient type.
type ReplicationPoliciesClientAPI interface {
	Create(ctx context.Context, policyName string, input siterecovery.CreatePolicyInput) (result siterecovery.ReplicationPoliciesCreateFuture, err error)
	Delete(ctx context.Context, policyName string) (result siterecovery.ReplicationPoliciesDeleteFuture, err error)
	Get(ctx context.Context, policyName string) (result siterecovery.Policy, err error)
	List(ctx context.Context) (result siterecovery.PolicyCollectionPage, err error)
	ListComplete(ctx context.Context) (result siterecovery.PolicyCollectionIterator, err error)
	Update(ctx context.Context, policyName string, input siterecovery.UpdatePolicyInput) (result siterecovery.ReplicationPoliciesUpdateFuture, err error)
}

var _ ReplicationPoliciesClientAPI = (*siterecovery.ReplicationPoliciesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result siterecovery.OperationsDiscoveryCollectionPage, err error)
	ListComplete(ctx context.Context) (result siterecovery.OperationsDiscoveryCollectionIterator, err error)
}

var _ OperationsClientAPI = (*siterecovery.OperationsClient)(nil)

// ReplicationNetworksClientAPI contains the set of methods on the ReplicationNetworksClient type.
type ReplicationNetworksClientAPI interface {
	Get(ctx context.Context, fabricName string, networkName string) (result siterecovery.Network, err error)
	List(ctx context.Context) (result siterecovery.NetworkCollectionPage, err error)
	ListComplete(ctx context.Context) (result siterecovery.NetworkCollectionIterator, err error)
	ListByReplicationFabrics(ctx context.Context, fabricName string) (result siterecovery.NetworkCollectionPage, err error)
	ListByReplicationFabricsComplete(ctx context.Context, fabricName string) (result siterecovery.NetworkCollectionIterator, err error)
}

var _ ReplicationNetworksClientAPI = (*siterecovery.ReplicationNetworksClient)(nil)

// ReplicationLogicalNetworksClientAPI contains the set of methods on the ReplicationLogicalNetworksClient type.
type ReplicationLogicalNetworksClientAPI interface {
	Get(ctx context.Context, fabricName string, logicalNetworkName string) (result siterecovery.LogicalNetwork, err error)
	ListByReplicationFabrics(ctx context.Context, fabricName string) (result siterecovery.LogicalNetworkCollectionPage, err error)
	ListByReplicationFabricsComplete(ctx context.Context, fabricName string) (result siterecovery.LogicalNetworkCollectionIterator, err error)
}

var _ ReplicationLogicalNetworksClientAPI = (*siterecovery.ReplicationLogicalNetworksClient)(nil)

// ReplicationJobsClientAPI contains the set of methods on the ReplicationJobsClient type.
type ReplicationJobsClientAPI interface {
	Cancel(ctx context.Context, jobName string) (result siterecovery.ReplicationJobsCancelFuture, err error)
	Export(ctx context.Context, jobQueryParameter siterecovery.JobQueryParameter) (result siterecovery.ReplicationJobsExportFuture, err error)
	Get(ctx context.Context, jobName string) (result siterecovery.Job, err error)
	List(ctx context.Context, filter string) (result siterecovery.JobCollectionPage, err error)
	ListComplete(ctx context.Context, filter string) (result siterecovery.JobCollectionIterator, err error)
	Restart(ctx context.Context, jobName string) (result siterecovery.ReplicationJobsRestartFuture, err error)
	Resume(ctx context.Context, jobName string, resumeJobParams siterecovery.ResumeJobParams) (result siterecovery.ReplicationJobsResumeFuture, err error)
}

var _ ReplicationJobsClientAPI = (*siterecovery.ReplicationJobsClient)(nil)

// ReplicationEventsClientAPI contains the set of methods on the ReplicationEventsClient type.
type ReplicationEventsClientAPI interface {
	Get(ctx context.Context, eventName string) (result siterecovery.Event, err error)
	List(ctx context.Context, filter string) (result siterecovery.EventCollectionPage, err error)
	ListComplete(ctx context.Context, filter string) (result siterecovery.EventCollectionIterator, err error)
}

var _ ReplicationEventsClientAPI = (*siterecovery.ReplicationEventsClient)(nil)

// ReplicationAlertSettingsClientAPI contains the set of methods on the ReplicationAlertSettingsClient type.
type ReplicationAlertSettingsClientAPI interface {
	Create(ctx context.Context, alertSettingName string, request siterecovery.ConfigureAlertRequest) (result siterecovery.Alert, err error)
	Get(ctx context.Context, alertSettingName string) (result siterecovery.Alert, err error)
	List(ctx context.Context) (result siterecovery.AlertCollectionPage, err error)
	ListComplete(ctx context.Context) (result siterecovery.AlertCollectionIterator, err error)
}

var _ ReplicationAlertSettingsClientAPI = (*siterecovery.ReplicationAlertSettingsClient)(nil)
