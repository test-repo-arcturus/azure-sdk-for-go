package servicefabricapi

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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/satori/go.uuid"
	"github.com/test-repo-arcturus/azure-sdk-for-go/services/servicefabric/6.2/servicefabric"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	BackupPartition(ctx context.Context, partitionID uuid.UUID, backupPartitionDescription *servicefabric.BackupPartitionDescription, backupTimeout *int32, timeout *int64) (result autorest.Response, err error)
	CancelOperation(ctx context.Context, operationID uuid.UUID, force bool, timeout *int64) (result autorest.Response, err error)
	CancelRepairTask(ctx context.Context, repairTaskCancelDescription servicefabric.RepairTaskCancelDescription) (result servicefabric.RepairTaskUpdateInfo, err error)
	CommitImageStoreUploadSession(ctx context.Context, sessionID uuid.UUID, timeout *int64) (result autorest.Response, err error)
	CopyImageStoreContent(ctx context.Context, imageStoreCopyDescription servicefabric.ImageStoreCopyDescription, timeout *int64) (result autorest.Response, err error)
	CreateApplication(ctx context.Context, applicationDescription servicefabric.ApplicationDescription, timeout *int64) (result autorest.Response, err error)
	CreateBackupPolicy(ctx context.Context, backupPolicyDescription servicefabric.BackupPolicyDescription, timeout *int64) (result autorest.Response, err error)
	CreateComposeDeployment(ctx context.Context, createComposeDeploymentDescription servicefabric.CreateComposeDeploymentDescription, timeout *int64) (result autorest.Response, err error)
	CreateName(ctx context.Context, nameDescription servicefabric.NameDescription, timeout *int64) (result autorest.Response, err error)
	CreateRepairTask(ctx context.Context, repairTask servicefabric.RepairTask) (result servicefabric.RepairTaskUpdateInfo, err error)
	CreateService(ctx context.Context, applicationID string, serviceDescription servicefabric.BasicServiceDescription, timeout *int64) (result autorest.Response, err error)
	CreateServiceFromTemplate(ctx context.Context, applicationID string, serviceFromTemplateDescription servicefabric.ServiceFromTemplateDescription, timeout *int64) (result autorest.Response, err error)
	DeleteApplication(ctx context.Context, applicationID string, forceRemove *bool, timeout *int64) (result autorest.Response, err error)
	DeleteBackupPolicy(ctx context.Context, backupPolicyName string, timeout *int64) (result autorest.Response, err error)
	DeleteImageStoreContent(ctx context.Context, contentPath string, timeout *int64) (result autorest.Response, err error)
	DeleteImageStoreUploadSession(ctx context.Context, sessionID uuid.UUID, timeout *int64) (result autorest.Response, err error)
	DeleteName(ctx context.Context, nameID string, timeout *int64) (result autorest.Response, err error)
	DeleteProperty(ctx context.Context, nameID string, propertyName string, timeout *int64) (result autorest.Response, err error)
	DeleteRepairTask(ctx context.Context, repairTaskDeleteDescription servicefabric.RepairTaskDeleteDescription) (result autorest.Response, err error)
	DeleteService(ctx context.Context, serviceID string, forceRemove *bool, timeout *int64) (result autorest.Response, err error)
	DeployServicePackageToNode(ctx context.Context, nodeName string, deployServicePackageToNodeDescription servicefabric.DeployServicePackageToNodeDescription, timeout *int64) (result autorest.Response, err error)
	DisableApplicationBackup(ctx context.Context, applicationID string, timeout *int64) (result autorest.Response, err error)
	DisableNode(ctx context.Context, nodeName string, deactivationIntentDescription servicefabric.DeactivationIntentDescription, timeout *int64) (result autorest.Response, err error)
	DisablePartitionBackup(ctx context.Context, partitionID uuid.UUID, timeout *int64) (result autorest.Response, err error)
	DisableServiceBackup(ctx context.Context, serviceID string, timeout *int64) (result autorest.Response, err error)
	EnableApplicationBackup(ctx context.Context, applicationID string, enableBackupDescription servicefabric.EnableBackupDescription, timeout *int64) (result autorest.Response, err error)
	EnableNode(ctx context.Context, nodeName string, timeout *int64) (result autorest.Response, err error)
	EnablePartitionBackup(ctx context.Context, partitionID uuid.UUID, enableBackupDescription servicefabric.EnableBackupDescription, timeout *int64) (result autorest.Response, err error)
	EnableServiceBackup(ctx context.Context, serviceID string, enableBackupDescription servicefabric.EnableBackupDescription, timeout *int64) (result autorest.Response, err error)
	ForceApproveRepairTask(ctx context.Context, repairTaskApproveDescription servicefabric.RepairTaskApproveDescription) (result servicefabric.RepairTaskUpdateInfo, err error)
	GetAadMetadata(ctx context.Context, timeout *int64) (result servicefabric.AadMetadataObject, err error)
	GetAllEntitiesBackedUpByPolicy(ctx context.Context, backupPolicyName string, continuationToken string, maxResults *int64, timeout *int64) (result servicefabric.PagedBackupEntityList, err error)
	GetApplicationBackupConfigurationInfo(ctx context.Context, applicationID string, continuationToken string, maxResults *int64, timeout *int64) (result servicefabric.PagedBackupConfigurationInfoList, err error)
	GetApplicationBackupList(ctx context.Context, applicationID string, timeout *int64, latest *bool, startDateTimeFilter *date.Time, endDateTimeFilter *date.Time, continuationToken string, maxResults *int64) (result servicefabric.PagedBackupInfoList, err error)
	GetApplicationEventList(ctx context.Context, applicationID string, startTimeUtc string, endTimeUtc string, timeout *int64, eventsTypesFilter string, excludeAnalysisEvents *bool, skipCorrelationLookup *bool) (result servicefabric.ListApplicationEvent, err error)
	GetApplicationHealth(ctx context.Context, applicationID string, eventsHealthStateFilter *int32, deployedApplicationsHealthStateFilter *int32, servicesHealthStateFilter *int32, excludeHealthStatistics *bool, timeout *int64) (result servicefabric.ApplicationHealth, err error)
	GetApplicationHealthUsingPolicy(ctx context.Context, applicationID string, eventsHealthStateFilter *int32, deployedApplicationsHealthStateFilter *int32, servicesHealthStateFilter *int32, excludeHealthStatistics *bool, applicationHealthPolicy *servicefabric.ApplicationHealthPolicy, timeout *int64) (result servicefabric.ApplicationHealth, err error)
	GetApplicationInfo(ctx context.Context, applicationID string, excludeApplicationParameters *bool, timeout *int64) (result servicefabric.ApplicationInfo, err error)
	GetApplicationInfoList(ctx context.Context, applicationDefinitionKindFilter *int32, applicationTypeName string, excludeApplicationParameters *bool, continuationToken string, maxResults *int64, timeout *int64) (result servicefabric.PagedApplicationInfoList, err error)
	GetApplicationLoadInfo(ctx context.Context, applicationID string, timeout *int64) (result servicefabric.ApplicationLoadInfo, err error)
	GetApplicationManifest(ctx context.Context, applicationTypeName string, applicationTypeVersion string, timeout *int64) (result servicefabric.ApplicationTypeManifest, err error)
	GetApplicationNameInfo(ctx context.Context, serviceID string, timeout *int64) (result servicefabric.ApplicationNameInfo, err error)
	GetApplicationsEventList(ctx context.Context, startTimeUtc string, endTimeUtc string, timeout *int64, eventsTypesFilter string, excludeAnalysisEvents *bool, skipCorrelationLookup *bool) (result servicefabric.ListApplicationEvent, err error)
	GetApplicationTypeInfoList(ctx context.Context, applicationTypeDefinitionKindFilter *int32, excludeApplicationParameters *bool, continuationToken string, maxResults *int64, timeout *int64) (result servicefabric.PagedApplicationTypeInfoList, err error)
	GetApplicationTypeInfoListByName(ctx context.Context, applicationTypeName string, applicationTypeVersion string, excludeApplicationParameters *bool, continuationToken string, maxResults *int64, timeout *int64) (result servicefabric.PagedApplicationTypeInfoList, err error)
	GetApplicationUpgrade(ctx context.Context, applicationID string, timeout *int64) (result servicefabric.ApplicationUpgradeProgressInfo, err error)
	GetBackupPolicyByName(ctx context.Context, backupPolicyName string, timeout *int64) (result servicefabric.BackupPolicyDescription, err error)
	GetBackupPolicyList(ctx context.Context, continuationToken string, maxResults *int64, timeout *int64) (result servicefabric.PagedBackupPolicyDescriptionList, err error)
	GetBackupsFromBackupLocation(ctx context.Context, getBackupByStorageQueryDescription servicefabric.GetBackupByStorageQueryDescription, timeout *int64, continuationToken string, maxResults *int64) (result servicefabric.PagedBackupInfoList, err error)
	GetChaos(ctx context.Context, timeout *int64) (result servicefabric.Chaos, err error)
	GetChaosEvents(ctx context.Context, continuationToken string, startTimeUtc string, endTimeUtc string, maxResults *int64, timeout *int64) (result servicefabric.ChaosEventsSegment, err error)
	GetChaosSchedule(ctx context.Context) (result servicefabric.ChaosScheduleDescription, err error)
	GetClusterConfiguration(ctx context.Context, configurationAPIVersion string, timeout *int64) (result servicefabric.ClusterConfiguration, err error)
	GetClusterConfigurationUpgradeStatus(ctx context.Context, timeout *int64) (result servicefabric.ClusterConfigurationUpgradeStatusInfo, err error)
	GetClusterEventList(ctx context.Context, startTimeUtc string, endTimeUtc string, timeout *int64, eventsTypesFilter string, excludeAnalysisEvents *bool, skipCorrelationLookup *bool) (result servicefabric.ListClusterEvent, err error)
	GetClusterHealth(ctx context.Context, nodesHealthStateFilter *int32, applicationsHealthStateFilter *int32, eventsHealthStateFilter *int32, excludeHealthStatistics *bool, includeSystemApplicationHealthStatistics *bool, timeout *int64) (result servicefabric.ClusterHealth, err error)
	GetClusterHealthChunk(ctx context.Context, timeout *int64) (result servicefabric.ClusterHealthChunk, err error)
	GetClusterHealthChunkUsingPolicyAndAdvancedFilters(ctx context.Context, clusterHealthChunkQueryDescription *servicefabric.ClusterHealthChunkQueryDescription, timeout *int64) (result servicefabric.ClusterHealthChunk, err error)
	GetClusterHealthUsingPolicy(ctx context.Context, nodesHealthStateFilter *int32, applicationsHealthStateFilter *int32, eventsHealthStateFilter *int32, excludeHealthStatistics *bool, includeSystemApplicationHealthStatistics *bool, clusterHealthPolicies *servicefabric.ClusterHealthPolicies, timeout *int64) (result servicefabric.ClusterHealth, err error)
	GetClusterManifest(ctx context.Context, timeout *int64) (result servicefabric.ClusterManifest, err error)
	GetClusterUpgradeProgress(ctx context.Context, timeout *int64) (result servicefabric.ClusterUpgradeProgressObject, err error)
	GetComposeDeploymentStatus(ctx context.Context, deploymentName string, timeout *int64) (result servicefabric.ComposeDeploymentStatusInfo, err error)
	GetComposeDeploymentStatusList(ctx context.Context, continuationToken string, maxResults *int64, timeout *int64) (result servicefabric.PagedComposeDeploymentStatusInfoList, err error)
	GetComposeDeploymentUpgradeProgress(ctx context.Context, deploymentName string, timeout *int64) (result servicefabric.ComposeDeploymentUpgradeProgressInfo, err error)
	GetContainerLogsDeployedOnNode(ctx context.Context, nodeName string, applicationID string, serviceManifestName string, codePackageName string, tail string, previous *bool, timeout *int64) (result servicefabric.ContainerLogs, err error)
	GetContainersEventList(ctx context.Context, startTimeUtc string, endTimeUtc string, timeout *int64, eventsTypesFilter string, excludeAnalysisEvents *bool, skipCorrelationLookup *bool) (result servicefabric.ListContainerInstanceEvent, err error)
	GetCorrelatedEventList(ctx context.Context, eventInstanceID string, timeout *int64) (result servicefabric.ListFabricEvent, err error)
	GetDataLossProgress(ctx context.Context, serviceID string, partitionID uuid.UUID, operationID uuid.UUID, timeout *int64) (result servicefabric.PartitionDataLossProgress, err error)
	GetDeployedApplicationHealth(ctx context.Context, nodeName string, applicationID string, eventsHealthStateFilter *int32, deployedServicePackagesHealthStateFilter *int32, excludeHealthStatistics *bool, timeout *int64) (result servicefabric.DeployedApplicationHealth, err error)
	GetDeployedApplicationHealthUsingPolicy(ctx context.Context, nodeName string, applicationID string, eventsHealthStateFilter *int32, deployedServicePackagesHealthStateFilter *int32, applicationHealthPolicy *servicefabric.ApplicationHealthPolicy, excludeHealthStatistics *bool, timeout *int64) (result servicefabric.DeployedApplicationHealth, err error)
	GetDeployedApplicationInfo(ctx context.Context, nodeName string, applicationID string, timeout *int64, includeHealthState *bool) (result servicefabric.DeployedApplicationInfo, err error)
	GetDeployedApplicationInfoList(ctx context.Context, nodeName string, timeout *int64, includeHealthState *bool, continuationToken string, maxResults *int64) (result servicefabric.PagedDeployedApplicationInfoList, err error)
	GetDeployedCodePackageInfoList(ctx context.Context, nodeName string, applicationID string, serviceManifestName string, codePackageName string, timeout *int64) (result servicefabric.ListDeployedCodePackageInfo, err error)
	GetDeployedServicePackageHealth(ctx context.Context, nodeName string, applicationID string, servicePackageName string, eventsHealthStateFilter *int32, timeout *int64) (result servicefabric.DeployedServicePackageHealth, err error)
	GetDeployedServicePackageHealthUsingPolicy(ctx context.Context, nodeName string, applicationID string, servicePackageName string, eventsHealthStateFilter *int32, applicationHealthPolicy *servicefabric.ApplicationHealthPolicy, timeout *int64) (result servicefabric.DeployedServicePackageHealth, err error)
	GetDeployedServicePackageInfoList(ctx context.Context, nodeName string, applicationID string, timeout *int64) (result servicefabric.ListDeployedServicePackageInfo, err error)
	GetDeployedServicePackageInfoListByName(ctx context.Context, nodeName string, applicationID string, servicePackageName string, timeout *int64) (result servicefabric.ListDeployedServicePackageInfo, err error)
	GetDeployedServiceReplicaDetailInfo(ctx context.Context, nodeName string, partitionID uuid.UUID, replicaID string, timeout *int64) (result servicefabric.DeployedServiceReplicaDetailInfoModel, err error)
	GetDeployedServiceReplicaDetailInfoByPartitionID(ctx context.Context, nodeName string, partitionID uuid.UUID, timeout *int64) (result servicefabric.DeployedServiceReplicaDetailInfoModel, err error)
	GetDeployedServiceReplicaInfoList(ctx context.Context, nodeName string, applicationID string, partitionID *uuid.UUID, serviceManifestName string, timeout *int64) (result servicefabric.ListDeployedServiceReplicaInfo, err error)
	GetDeployedServiceTypeInfoByName(ctx context.Context, nodeName string, applicationID string, serviceTypeName string, serviceManifestName string, timeout *int64) (result servicefabric.ListDeployedServiceTypeInfo, err error)
	GetDeployedServiceTypeInfoList(ctx context.Context, nodeName string, applicationID string, serviceManifestName string, timeout *int64) (result servicefabric.ListDeployedServiceTypeInfo, err error)
	GetFaultOperationList(ctx context.Context, typeFilter int32, stateFilter int32, timeout *int64) (result servicefabric.ListOperationStatus, err error)
	GetImageStoreContent(ctx context.Context, contentPath string, timeout *int64) (result servicefabric.ImageStoreContent, err error)
	GetImageStoreRootContent(ctx context.Context, timeout *int64) (result servicefabric.ImageStoreContent, err error)
	GetImageStoreUploadSessionByID(ctx context.Context, sessionID uuid.UUID, timeout *int64) (result servicefabric.UploadSession, err error)
	GetImageStoreUploadSessionByPath(ctx context.Context, contentPath string, timeout *int64) (result servicefabric.UploadSession, err error)
	GetNameExistsInfo(ctx context.Context, nameID string, timeout *int64) (result autorest.Response, err error)
	GetNodeEventList(ctx context.Context, nodeName string, startTimeUtc string, endTimeUtc string, timeout *int64, eventsTypesFilter string, excludeAnalysisEvents *bool, skipCorrelationLookup *bool) (result servicefabric.ListNodeEvent, err error)
	GetNodeHealth(ctx context.Context, nodeName string, eventsHealthStateFilter *int32, timeout *int64) (result servicefabric.NodeHealth, err error)
	GetNodeHealthUsingPolicy(ctx context.Context, nodeName string, eventsHealthStateFilter *int32, clusterHealthPolicy *servicefabric.ClusterHealthPolicy, timeout *int64) (result servicefabric.NodeHealth, err error)
	GetNodeInfo(ctx context.Context, nodeName string, timeout *int64) (result servicefabric.NodeInfo, err error)
	GetNodeInfoList(ctx context.Context, continuationToken string, nodeStatusFilter servicefabric.NodeStatusFilter, timeout *int64) (result servicefabric.PagedNodeInfoList, err error)
	GetNodeLoadInfo(ctx context.Context, nodeName string, timeout *int64) (result servicefabric.NodeLoadInfo, err error)
	GetNodesEventList(ctx context.Context, startTimeUtc string, endTimeUtc string, timeout *int64, eventsTypesFilter string, excludeAnalysisEvents *bool, skipCorrelationLookup *bool) (result servicefabric.ListNodeEvent, err error)
	GetNodeTransitionProgress(ctx context.Context, nodeName string, operationID uuid.UUID, timeout *int64) (result servicefabric.NodeTransitionProgress, err error)
	GetPartitionBackupConfigurationInfo(ctx context.Context, partitionID uuid.UUID, timeout *int64) (result servicefabric.PartitionBackupConfigurationInfo, err error)
	GetPartitionBackupList(ctx context.Context, partitionID uuid.UUID, timeout *int64, latest *bool, startDateTimeFilter *date.Time, endDateTimeFilter *date.Time) (result servicefabric.PagedBackupInfoList, err error)
	GetPartitionBackupProgress(ctx context.Context, partitionID uuid.UUID, timeout *int64) (result servicefabric.BackupProgressInfo, err error)
	GetPartitionEventList(ctx context.Context, partitionID uuid.UUID, startTimeUtc string, endTimeUtc string, timeout *int64, eventsTypesFilter string, excludeAnalysisEvents *bool, skipCorrelationLookup *bool) (result servicefabric.ListPartitionEvent, err error)
	GetPartitionHealth(ctx context.Context, partitionID uuid.UUID, eventsHealthStateFilter *int32, replicasHealthStateFilter *int32, excludeHealthStatistics *bool, timeout *int64) (result servicefabric.PartitionHealth, err error)
	GetPartitionHealthUsingPolicy(ctx context.Context, partitionID uuid.UUID, eventsHealthStateFilter *int32, replicasHealthStateFilter *int32, applicationHealthPolicy *servicefabric.ApplicationHealthPolicy, excludeHealthStatistics *bool, timeout *int64) (result servicefabric.PartitionHealth, err error)
	GetPartitionInfo(ctx context.Context, partitionID uuid.UUID, timeout *int64) (result servicefabric.ServicePartitionInfoModel, err error)
	GetPartitionInfoList(ctx context.Context, serviceID string, continuationToken string, timeout *int64) (result servicefabric.PagedServicePartitionInfoList, err error)
	GetPartitionLoadInformation(ctx context.Context, partitionID uuid.UUID, timeout *int64) (result servicefabric.PartitionLoadInformation, err error)
	GetPartitionReplicaEventList(ctx context.Context, partitionID uuid.UUID, replicaID string, startTimeUtc string, endTimeUtc string, timeout *int64, eventsTypesFilter string, excludeAnalysisEvents *bool, skipCorrelationLookup *bool) (result servicefabric.ListReplicaEvent, err error)
	GetPartitionReplicasEventList(ctx context.Context, partitionID uuid.UUID, startTimeUtc string, endTimeUtc string, timeout *int64, eventsTypesFilter string, excludeAnalysisEvents *bool, skipCorrelationLookup *bool) (result servicefabric.ListReplicaEvent, err error)
	GetPartitionRestartProgress(ctx context.Context, serviceID string, partitionID uuid.UUID, operationID uuid.UUID, timeout *int64) (result servicefabric.PartitionRestartProgress, err error)
	GetPartitionRestoreProgress(ctx context.Context, partitionID uuid.UUID, timeout *int64) (result servicefabric.RestoreProgressInfo, err error)
	GetPartitionsEventList(ctx context.Context, startTimeUtc string, endTimeUtc string, timeout *int64, eventsTypesFilter string, excludeAnalysisEvents *bool, skipCorrelationLookup *bool) (result servicefabric.ListPartitionEvent, err error)
	GetPropertyInfo(ctx context.Context, nameID string, propertyName string, timeout *int64) (result servicefabric.PropertyInfo, err error)
	GetPropertyInfoList(ctx context.Context, nameID string, includeValues *bool, continuationToken string, timeout *int64) (result servicefabric.PagedPropertyInfoList, err error)
	GetProvisionedFabricCodeVersionInfoList(ctx context.Context, codeVersion string, timeout *int64) (result servicefabric.ListFabricCodeVersionInfo, err error)
	GetProvisionedFabricConfigVersionInfoList(ctx context.Context, configVersion string, timeout *int64) (result servicefabric.ListFabricConfigVersionInfo, err error)
	GetQuorumLossProgress(ctx context.Context, serviceID string, partitionID uuid.UUID, operationID uuid.UUID, timeout *int64) (result servicefabric.PartitionQuorumLossProgress, err error)
	GetRepairTaskList(ctx context.Context, taskIDFilter string, stateFilter *int32, executorFilter string) (result servicefabric.ListRepairTask, err error)
	GetReplicaHealth(ctx context.Context, partitionID uuid.UUID, replicaID string, eventsHealthStateFilter *int32, timeout *int64) (result servicefabric.ReplicaHealthModel, err error)
	GetReplicaHealthUsingPolicy(ctx context.Context, partitionID uuid.UUID, replicaID string, eventsHealthStateFilter *int32, applicationHealthPolicy *servicefabric.ApplicationHealthPolicy, timeout *int64) (result servicefabric.ReplicaHealthModel, err error)
	GetReplicaInfo(ctx context.Context, partitionID uuid.UUID, replicaID string, timeout *int64) (result servicefabric.ReplicaInfoModel, err error)
	GetReplicaInfoList(ctx context.Context, partitionID uuid.UUID, continuationToken string, timeout *int64) (result servicefabric.PagedReplicaInfoList, err error)
	GetServiceBackupConfigurationInfo(ctx context.Context, serviceID string, continuationToken string, maxResults *int64, timeout *int64) (result servicefabric.PagedBackupConfigurationInfoList, err error)
	GetServiceBackupList(ctx context.Context, serviceID string, timeout *int64, latest *bool, startDateTimeFilter *date.Time, endDateTimeFilter *date.Time, continuationToken string, maxResults *int64) (result servicefabric.PagedBackupInfoList, err error)
	GetServiceDescription(ctx context.Context, serviceID string, timeout *int64) (result servicefabric.ServiceDescriptionModel, err error)
	GetServiceEventList(ctx context.Context, serviceID string, startTimeUtc string, endTimeUtc string, timeout *int64, eventsTypesFilter string, excludeAnalysisEvents *bool, skipCorrelationLookup *bool) (result servicefabric.ListServiceEvent, err error)
	GetServiceHealth(ctx context.Context, serviceID string, eventsHealthStateFilter *int32, partitionsHealthStateFilter *int32, excludeHealthStatistics *bool, timeout *int64) (result servicefabric.ServiceHealth, err error)
	GetServiceHealthUsingPolicy(ctx context.Context, serviceID string, eventsHealthStateFilter *int32, partitionsHealthStateFilter *int32, applicationHealthPolicy *servicefabric.ApplicationHealthPolicy, excludeHealthStatistics *bool, timeout *int64) (result servicefabric.ServiceHealth, err error)
	GetServiceInfo(ctx context.Context, applicationID string, serviceID string, timeout *int64) (result servicefabric.ServiceInfoModel, err error)
	GetServiceInfoList(ctx context.Context, applicationID string, serviceTypeName string, continuationToken string, timeout *int64) (result servicefabric.PagedServiceInfoList, err error)
	GetServiceManifest(ctx context.Context, applicationTypeName string, applicationTypeVersion string, serviceManifestName string, timeout *int64) (result servicefabric.ServiceTypeManifest, err error)
	GetServiceNameInfo(ctx context.Context, partitionID uuid.UUID, timeout *int64) (result servicefabric.ServiceNameInfo, err error)
	GetServicesEventList(ctx context.Context, startTimeUtc string, endTimeUtc string, timeout *int64, eventsTypesFilter string, excludeAnalysisEvents *bool, skipCorrelationLookup *bool) (result servicefabric.ListServiceEvent, err error)
	GetServiceTypeInfoByName(ctx context.Context, applicationTypeName string, applicationTypeVersion string, serviceTypeName string, timeout *int64) (result servicefabric.ServiceTypeInfo, err error)
	GetServiceTypeInfoList(ctx context.Context, applicationTypeName string, applicationTypeVersion string, timeout *int64) (result servicefabric.ListServiceTypeInfo, err error)
	GetSubNameInfoList(ctx context.Context, nameID string, recursive *bool, continuationToken string, timeout *int64) (result servicefabric.PagedSubNameInfoList, err error)
	GetUpgradeOrchestrationServiceState(ctx context.Context, timeout *int64) (result servicefabric.UpgradeOrchestrationServiceState, err error)
	InvokeContainerAPI(ctx context.Context, nodeName string, applicationID string, serviceManifestName string, codePackageName string, codePackageInstanceID string, containerAPIRequestBody servicefabric.ContainerAPIRequestBody, timeout *int64) (result servicefabric.ContainerAPIResponse, err error)
	InvokeInfrastructureCommand(ctx context.Context, command string, serviceID string, timeout *int64) (result servicefabric.String, err error)
	InvokeInfrastructureQuery(ctx context.Context, command string, serviceID string, timeout *int64) (result servicefabric.String, err error)
	PostChaosSchedule(ctx context.Context, chaosSchedule servicefabric.ChaosScheduleDescription) (result autorest.Response, err error)
	ProvisionApplicationType(ctx context.Context, provisionApplicationTypeDescriptionBaseRequiredBodyParam servicefabric.BasicProvisionApplicationTypeDescriptionBase, timeout *int64) (result autorest.Response, err error)
	ProvisionCluster(ctx context.Context, provisionFabricDescription servicefabric.ProvisionFabricDescription, timeout *int64) (result autorest.Response, err error)
	PutProperty(ctx context.Context, nameID string, propertyDescription servicefabric.PropertyDescription, timeout *int64) (result autorest.Response, err error)
	RecoverAllPartitions(ctx context.Context, timeout *int64) (result autorest.Response, err error)
	RecoverPartition(ctx context.Context, partitionID uuid.UUID, timeout *int64) (result autorest.Response, err error)
	RecoverServicePartitions(ctx context.Context, serviceID string, timeout *int64) (result autorest.Response, err error)
	RecoverSystemPartitions(ctx context.Context, timeout *int64) (result autorest.Response, err error)
	RemoveComposeDeployment(ctx context.Context, deploymentName string, timeout *int64) (result autorest.Response, err error)
	RemoveNodeState(ctx context.Context, nodeName string, timeout *int64) (result autorest.Response, err error)
	RemoveReplica(ctx context.Context, nodeName string, partitionID uuid.UUID, replicaID string, forceRemove *bool, timeout *int64) (result autorest.Response, err error)
	ReportApplicationHealth(ctx context.Context, applicationID string, healthInformation servicefabric.HealthInformation, immediate *bool, timeout *int64) (result autorest.Response, err error)
	ReportClusterHealth(ctx context.Context, healthInformation servicefabric.HealthInformation, immediate *bool, timeout *int64) (result autorest.Response, err error)
	ReportDeployedApplicationHealth(ctx context.Context, nodeName string, applicationID string, healthInformation servicefabric.HealthInformation, immediate *bool, timeout *int64) (result autorest.Response, err error)
	ReportDeployedServicePackageHealth(ctx context.Context, nodeName string, applicationID string, servicePackageName string, healthInformation servicefabric.HealthInformation, immediate *bool, timeout *int64) (result autorest.Response, err error)
	ReportNodeHealth(ctx context.Context, nodeName string, healthInformation servicefabric.HealthInformation, immediate *bool, timeout *int64) (result autorest.Response, err error)
	ReportPartitionHealth(ctx context.Context, partitionID uuid.UUID, healthInformation servicefabric.HealthInformation, immediate *bool, timeout *int64) (result autorest.Response, err error)
	ReportReplicaHealth(ctx context.Context, partitionID uuid.UUID, replicaID string, replicaHealthReportServiceKind servicefabric.ReplicaHealthReportServiceKind, healthInformation servicefabric.HealthInformation, immediate *bool, timeout *int64) (result autorest.Response, err error)
	ReportServiceHealth(ctx context.Context, serviceID string, healthInformation servicefabric.HealthInformation, immediate *bool, timeout *int64) (result autorest.Response, err error)
	ResetPartitionLoad(ctx context.Context, partitionID uuid.UUID, timeout *int64) (result autorest.Response, err error)
	ResolveService(ctx context.Context, serviceID string, partitionKeyType *int32, partitionKeyValue string, previousRspVersion string, timeout *int64) (result servicefabric.ResolvedServicePartition, err error)
	RestartDeployedCodePackage(ctx context.Context, nodeName string, applicationID string, restartDeployedCodePackageDescription servicefabric.RestartDeployedCodePackageDescription, timeout *int64) (result autorest.Response, err error)
	RestartNode(ctx context.Context, nodeName string, restartNodeDescription servicefabric.RestartNodeDescription, timeout *int64) (result autorest.Response, err error)
	RestartReplica(ctx context.Context, nodeName string, partitionID uuid.UUID, replicaID string, timeout *int64) (result autorest.Response, err error)
	RestorePartition(ctx context.Context, partitionID uuid.UUID, restorePartitionDescription servicefabric.RestorePartitionDescription, restoreTimeout *int32, timeout *int64) (result autorest.Response, err error)
	ResumeApplicationBackup(ctx context.Context, applicationID string, timeout *int64) (result autorest.Response, err error)
	ResumeApplicationUpgrade(ctx context.Context, applicationID string, resumeApplicationUpgradeDescription servicefabric.ResumeApplicationUpgradeDescription, timeout *int64) (result autorest.Response, err error)
	ResumeClusterUpgrade(ctx context.Context, resumeClusterUpgradeDescription servicefabric.ResumeClusterUpgradeDescription, timeout *int64) (result autorest.Response, err error)
	ResumePartitionBackup(ctx context.Context, partitionID uuid.UUID, timeout *int64) (result autorest.Response, err error)
	ResumeServiceBackup(ctx context.Context, serviceID string, timeout *int64) (result autorest.Response, err error)
	RollbackApplicationUpgrade(ctx context.Context, applicationID string, timeout *int64) (result autorest.Response, err error)
	RollbackClusterUpgrade(ctx context.Context, timeout *int64) (result autorest.Response, err error)
	SetUpgradeOrchestrationServiceState(ctx context.Context, upgradeOrchestrationServiceState servicefabric.UpgradeOrchestrationServiceState, timeout *int64) (result servicefabric.UpgradeOrchestrationServiceStateSummary, err error)
	StartApplicationUpgrade(ctx context.Context, applicationID string, applicationUpgradeDescription servicefabric.ApplicationUpgradeDescription, timeout *int64) (result autorest.Response, err error)
	StartChaos(ctx context.Context, chaosParameters servicefabric.ChaosParameters, timeout *int64) (result autorest.Response, err error)
	StartClusterConfigurationUpgrade(ctx context.Context, clusterConfigurationUpgradeDescription servicefabric.ClusterConfigurationUpgradeDescription, timeout *int64) (result autorest.Response, err error)
	StartClusterUpgrade(ctx context.Context, startClusterUpgradeDescription servicefabric.StartClusterUpgradeDescription, timeout *int64) (result autorest.Response, err error)
	StartComposeDeploymentUpgrade(ctx context.Context, deploymentName string, composeDeploymentUpgradeDescription servicefabric.ComposeDeploymentUpgradeDescription, timeout *int64) (result autorest.Response, err error)
	StartDataLoss(ctx context.Context, serviceID string, partitionID uuid.UUID, operationID uuid.UUID, dataLossMode servicefabric.DataLossMode, timeout *int64) (result autorest.Response, err error)
	StartNodeTransition(ctx context.Context, nodeName string, operationID uuid.UUID, nodeTransitionType servicefabric.NodeTransitionType, nodeInstanceID string, stopDurationInSeconds int32, timeout *int64) (result autorest.Response, err error)
	StartPartitionRestart(ctx context.Context, serviceID string, partitionID uuid.UUID, operationID uuid.UUID, restartPartitionMode servicefabric.RestartPartitionMode, timeout *int64) (result autorest.Response, err error)
	StartQuorumLoss(ctx context.Context, serviceID string, partitionID uuid.UUID, operationID uuid.UUID, quorumLossMode servicefabric.QuorumLossMode, quorumLossDuration int32, timeout *int64) (result autorest.Response, err error)
	StopChaos(ctx context.Context, timeout *int64) (result autorest.Response, err error)
	SubmitPropertyBatch(ctx context.Context, nameID string, propertyBatchDescriptionList servicefabric.PropertyBatchDescriptionList, timeout *int64) (result servicefabric.PropertyBatchInfoModel, err error)
	SuspendApplicationBackup(ctx context.Context, applicationID string, timeout *int64) (result autorest.Response, err error)
	SuspendPartitionBackup(ctx context.Context, partitionID uuid.UUID, timeout *int64) (result autorest.Response, err error)
	SuspendServiceBackup(ctx context.Context, serviceID string, timeout *int64) (result autorest.Response, err error)
	UnprovisionApplicationType(ctx context.Context, applicationTypeName string, unprovisionApplicationTypeDescriptionInfo servicefabric.UnprovisionApplicationTypeDescriptionInfo, timeout *int64) (result autorest.Response, err error)
	UnprovisionCluster(ctx context.Context, unprovisionFabricDescription servicefabric.UnprovisionFabricDescription, timeout *int64) (result autorest.Response, err error)
	UpdateApplicationUpgrade(ctx context.Context, applicationID string, applicationUpgradeUpdateDescription servicefabric.ApplicationUpgradeUpdateDescription, timeout *int64) (result autorest.Response, err error)
	UpdateBackupPolicy(ctx context.Context, backupPolicyDescription servicefabric.BackupPolicyDescription, backupPolicyName string, timeout *int64) (result autorest.Response, err error)
	UpdateClusterUpgrade(ctx context.Context, updateClusterUpgradeDescription servicefabric.UpdateClusterUpgradeDescription, timeout *int64) (result autorest.Response, err error)
	UpdateRepairExecutionState(ctx context.Context, repairTask servicefabric.RepairTask) (result servicefabric.RepairTaskUpdateInfo, err error)
	UpdateRepairTaskHealthPolicy(ctx context.Context, repairTaskUpdateHealthPolicyDescription servicefabric.RepairTaskUpdateHealthPolicyDescription) (result servicefabric.RepairTaskUpdateInfo, err error)
	UpdateService(ctx context.Context, serviceID string, serviceUpdateDescription servicefabric.BasicServiceUpdateDescription, timeout *int64) (result autorest.Response, err error)
	UploadFile(ctx context.Context, contentPath string, timeout *int64) (result autorest.Response, err error)
	UploadFileChunk(ctx context.Context, contentPath string, sessionID uuid.UUID, contentRange string, timeout *int64) (result autorest.Response, err error)
}

var _ BaseClientAPI = (*servicefabric.BaseClient)(nil)
