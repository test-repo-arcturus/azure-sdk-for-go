package automationapi

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
	"github.com/satori/go.uuid"
	"github.com/test-repo-arcturus/azure-sdk-for-go/services/automation/mgmt/2015-10-31/automation"
	"io"
)

// AccountClientAPI contains the set of methods on the AccountClient type.
type AccountClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, parameters automation.AccountCreateOrUpdateParameters) (result automation.Account, err error)
	Delete(ctx context.Context, resourceGroupName string, automationAccountName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.Account, err error)
	List(ctx context.Context) (result automation.AccountListResultPage, err error)
	ListComplete(ctx context.Context) (result automation.AccountListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result automation.AccountListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result automation.AccountListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, automationAccountName string, parameters automation.AccountUpdateParameters) (result automation.Account, err error)
}

var _ AccountClientAPI = (*automation.AccountClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result automation.OperationListResult, err error)
}

var _ OperationsClientAPI = (*automation.OperationsClient)(nil)

// StatisticsClientAPI contains the set of methods on the StatisticsClient type.
type StatisticsClientAPI interface {
	ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (result automation.StatisticsListResult, err error)
}

var _ StatisticsClientAPI = (*automation.StatisticsClient)(nil)

// UsagesClientAPI contains the set of methods on the UsagesClient type.
type UsagesClientAPI interface {
	ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.UsageListResult, err error)
}

var _ UsagesClientAPI = (*automation.UsagesClient)(nil)

// KeysClientAPI contains the set of methods on the KeysClient type.
type KeysClientAPI interface {
	ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.KeyListResult, err error)
}

var _ KeysClientAPI = (*automation.KeysClient)(nil)

// CertificateClientAPI contains the set of methods on the CertificateClient type.
type CertificateClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string, parameters automation.CertificateCreateOrUpdateParameters) (result automation.Certificate, err error)
	Delete(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string) (result automation.Certificate, err error)
	ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.CertificateListResultPage, err error)
	ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.CertificateListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string, parameters automation.CertificateUpdateParameters) (result automation.Certificate, err error)
}

var _ CertificateClientAPI = (*automation.CertificateClient)(nil)

// ConnectionClientAPI contains the set of methods on the ConnectionClient type.
type ConnectionClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, connectionName string, parameters automation.ConnectionCreateOrUpdateParameters) (result automation.Connection, err error)
	Delete(ctx context.Context, resourceGroupName string, automationAccountName string, connectionName string) (result automation.Connection, err error)
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, connectionName string) (result automation.Connection, err error)
	ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.ConnectionListResultPage, err error)
	ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.ConnectionListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, automationAccountName string, connectionName string, parameters automation.ConnectionUpdateParameters) (result automation.Connection, err error)
}

var _ ConnectionClientAPI = (*automation.ConnectionClient)(nil)

// ConnectionTypeClientAPI contains the set of methods on the ConnectionTypeClient type.
type ConnectionTypeClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, connectionTypeName string, parameters automation.ConnectionTypeCreateOrUpdateParameters) (result automation.ConnectionType, err error)
	Delete(ctx context.Context, resourceGroupName string, automationAccountName string, connectionTypeName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, connectionTypeName string) (result automation.ConnectionType, err error)
	ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.ConnectionTypeListResultPage, err error)
	ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.ConnectionTypeListResultIterator, err error)
}

var _ ConnectionTypeClientAPI = (*automation.ConnectionTypeClient)(nil)

// CredentialClientAPI contains the set of methods on the CredentialClient type.
type CredentialClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, credentialName string, parameters automation.CredentialCreateOrUpdateParameters) (result automation.Credential, err error)
	Delete(ctx context.Context, resourceGroupName string, automationAccountName string, credentialName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, credentialName string) (result automation.Credential, err error)
	ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.CredentialListResultPage, err error)
	ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.CredentialListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, automationAccountName string, credentialName string, parameters automation.CredentialUpdateParameters) (result automation.Credential, err error)
}

var _ CredentialClientAPI = (*automation.CredentialClient)(nil)

// DscCompilationJobClientAPI contains the set of methods on the DscCompilationJobClient type.
type DscCompilationJobClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobID uuid.UUID, parameters automation.DscCompilationJobCreateParameters) (result automation.DscCompilationJob, err error)
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobID uuid.UUID) (result automation.DscCompilationJob, err error)
	GetStream(ctx context.Context, resourceGroupName string, automationAccountName string, jobID uuid.UUID, jobStreamID string) (result automation.JobStream, err error)
	ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (result automation.DscCompilationJobListResultPage, err error)
	ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (result automation.DscCompilationJobListResultIterator, err error)
}

var _ DscCompilationJobClientAPI = (*automation.DscCompilationJobClient)(nil)

// DscCompilationJobStreamClientAPI contains the set of methods on the DscCompilationJobStreamClient type.
type DscCompilationJobStreamClientAPI interface {
	ListByJob(ctx context.Context, resourceGroupName string, automationAccountName string, jobID uuid.UUID) (result automation.JobStreamListResult, err error)
}

var _ DscCompilationJobStreamClientAPI = (*automation.DscCompilationJobStreamClient)(nil)

// DscConfigurationClientAPI contains the set of methods on the DscConfigurationClient type.
type DscConfigurationClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, parameters automation.DscConfigurationCreateOrUpdateParameters) (result automation.DscConfiguration, err error)
	Delete(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string) (result automation.DscConfiguration, err error)
	GetContent(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string) (result automation.ReadCloser, err error)
	ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string, filter string, skip *int32, top *int32, inlinecount string) (result automation.DscConfigurationListResultPage, err error)
	ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string, filter string, skip *int32, top *int32, inlinecount string) (result automation.DscConfigurationListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, parameters *automation.DscConfigurationUpdateParameters) (result automation.DscConfiguration, err error)
}

var _ DscConfigurationClientAPI = (*automation.DscConfigurationClient)(nil)

// AgentRegistrationInformationClientAPI contains the set of methods on the AgentRegistrationInformationClient type.
type AgentRegistrationInformationClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.AgentRegistration, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, automationAccountName string, parameters automation.AgentRegistrationRegenerateKeyParameter) (result automation.AgentRegistration, err error)
}

var _ AgentRegistrationInformationClientAPI = (*automation.AgentRegistrationInformationClient)(nil)

// DscNodeClientAPI contains the set of methods on the DscNodeClient type.
type DscNodeClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string) (result automation.DscNode, err error)
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string) (result automation.DscNode, err error)
	ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (result automation.DscNodeListResultPage, err error)
	ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (result automation.DscNodeListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, parameters automation.DscNodeUpdateParameters) (result automation.DscNode, err error)
}

var _ DscNodeClientAPI = (*automation.DscNodeClient)(nil)

// NodeReportsClientAPI contains the set of methods on the NodeReportsClient type.
type NodeReportsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, reportID string) (result automation.DscNodeReport, err error)
	GetContent(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, reportID string) (result automation.SetObject, err error)
	ListByNode(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, filter string) (result automation.DscNodeReportListResultPage, err error)
	ListByNodeComplete(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, filter string) (result automation.DscNodeReportListResultIterator, err error)
}

var _ NodeReportsClientAPI = (*automation.NodeReportsClient)(nil)

// DscNodeConfigurationClientAPI contains the set of methods on the DscNodeConfigurationClient type.
type DscNodeConfigurationClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, nodeConfigurationName string, parameters automation.DscNodeConfigurationCreateOrUpdateParameters) (result automation.DscNodeConfiguration, err error)
	Delete(ctx context.Context, resourceGroupName string, automationAccountName string, nodeConfigurationName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, nodeConfigurationName string) (result automation.DscNodeConfiguration, err error)
	ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (result automation.DscNodeConfigurationListResultPage, err error)
	ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (result automation.DscNodeConfigurationListResultIterator, err error)
}

var _ DscNodeConfigurationClientAPI = (*automation.DscNodeConfigurationClient)(nil)

// HybridRunbookWorkerGroupClientAPI contains the set of methods on the HybridRunbookWorkerGroupClient type.
type HybridRunbookWorkerGroupClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string) (result automation.HybridRunbookWorkerGroup, err error)
	ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (result automation.HybridRunbookWorkerGroupsListResultPage, err error)
	ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (result automation.HybridRunbookWorkerGroupsListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, parameters automation.HybridRunbookWorkerGroupUpdateParameters) (result automation.HybridRunbookWorkerGroup, err error)
}

var _ HybridRunbookWorkerGroupClientAPI = (*automation.HybridRunbookWorkerGroupClient)(nil)

// JobClientAPI contains the set of methods on the JobClient type.
type JobClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, automationAccountName string, jobID uuid.UUID, parameters automation.JobCreateParameters) (result automation.Job, err error)
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, jobID uuid.UUID) (result automation.Job, err error)
	GetOutput(ctx context.Context, resourceGroupName string, automationAccountName string, jobID string) (result automation.ReadCloser, err error)
	GetRunbookContent(ctx context.Context, resourceGroupName string, automationAccountName string, jobID string) (result automation.ReadCloser, err error)
	ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (result automation.JobListResultPage, err error)
	ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (result automation.JobListResultIterator, err error)
	Resume(ctx context.Context, resourceGroupName string, automationAccountName string, jobID uuid.UUID) (result autorest.Response, err error)
	Stop(ctx context.Context, resourceGroupName string, automationAccountName string, jobID uuid.UUID) (result autorest.Response, err error)
	Suspend(ctx context.Context, resourceGroupName string, automationAccountName string, jobID uuid.UUID) (result autorest.Response, err error)
}

var _ JobClientAPI = (*automation.JobClient)(nil)

// JobStreamClientAPI contains the set of methods on the JobStreamClient type.
type JobStreamClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, jobID string, jobStreamID string) (result automation.JobStream, err error)
	ListByJob(ctx context.Context, resourceGroupName string, automationAccountName string, jobID string, filter string) (result automation.JobStreamListResultPage, err error)
	ListByJobComplete(ctx context.Context, resourceGroupName string, automationAccountName string, jobID string, filter string) (result automation.JobStreamListResultIterator, err error)
}

var _ JobStreamClientAPI = (*automation.JobStreamClient)(nil)

// JobScheduleClientAPI contains the set of methods on the JobScheduleClient type.
type JobScheduleClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, automationAccountName string, jobScheduleID uuid.UUID, parameters automation.JobScheduleCreateParameters) (result automation.JobSchedule, err error)
	Delete(ctx context.Context, resourceGroupName string, automationAccountName string, jobScheduleID uuid.UUID) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, jobScheduleID uuid.UUID) (result automation.JobSchedule, err error)
	ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (result automation.JobScheduleListResultPage, err error)
	ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (result automation.JobScheduleListResultIterator, err error)
}

var _ JobScheduleClientAPI = (*automation.JobScheduleClient)(nil)

// LinkedWorkspaceClientAPI contains the set of methods on the LinkedWorkspaceClient type.
type LinkedWorkspaceClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.LinkedWorkspace, err error)
}

var _ LinkedWorkspaceClientAPI = (*automation.LinkedWorkspaceClient)(nil)

// ActivityClientAPI contains the set of methods on the ActivityClient type.
type ActivityClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, activityName string) (result automation.Activity, err error)
	ListByModule(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string) (result automation.ActivityListResultPage, err error)
	ListByModuleComplete(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string) (result automation.ActivityListResultIterator, err error)
}

var _ ActivityClientAPI = (*automation.ActivityClient)(nil)

// ModuleClientAPI contains the set of methods on the ModuleClient type.
type ModuleClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, parameters automation.ModuleCreateOrUpdateParameters) (result automation.Module, err error)
	Delete(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string) (result automation.Module, err error)
	ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.ModuleListResultPage, err error)
	ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.ModuleListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, parameters automation.ModuleUpdateParameters) (result automation.Module, err error)
}

var _ ModuleClientAPI = (*automation.ModuleClient)(nil)

// ObjectDataTypesClientAPI contains the set of methods on the ObjectDataTypesClient type.
type ObjectDataTypesClientAPI interface {
	ListFieldsByModuleAndType(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, typeName string) (result automation.TypeFieldListResult, err error)
	ListFieldsByType(ctx context.Context, resourceGroupName string, automationAccountName string, typeName string) (result automation.TypeFieldListResult, err error)
}

var _ ObjectDataTypesClientAPI = (*automation.ObjectDataTypesClient)(nil)

// FieldsClientAPI contains the set of methods on the FieldsClient type.
type FieldsClientAPI interface {
	ListByType(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, typeName string) (result automation.TypeFieldListResult, err error)
}

var _ FieldsClientAPI = (*automation.FieldsClient)(nil)

// RunbookDraftClientAPI contains the set of methods on the RunbookDraftClient type.
type RunbookDraftClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (result automation.RunbookDraft, err error)
	GetContent(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (result automation.ReadCloser, err error)
	Publish(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (result automation.RunbookDraftPublishFuture, err error)
	ReplaceContent(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, runbookContent io.ReadCloser) (result automation.RunbookDraftReplaceContentFuture, err error)
	UndoEdit(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (result automation.RunbookDraftUndoEditResult, err error)
}

var _ RunbookDraftClientAPI = (*automation.RunbookDraftClient)(nil)

// RunbookClientAPI contains the set of methods on the RunbookClient type.
type RunbookClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, parameters automation.RunbookCreateOrUpdateParameters) (result automation.Runbook, err error)
	Delete(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (result automation.Runbook, err error)
	GetContent(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (result automation.ReadCloser, err error)
	ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.RunbookListResultPage, err error)
	ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.RunbookListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, parameters automation.RunbookUpdateParameters) (result automation.Runbook, err error)
}

var _ RunbookClientAPI = (*automation.RunbookClient)(nil)

// TestJobStreamsClientAPI contains the set of methods on the TestJobStreamsClient type.
type TestJobStreamsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, jobStreamID string) (result automation.JobStream, err error)
	ListByTestJob(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, filter string) (result automation.JobStreamListResultPage, err error)
	ListByTestJobComplete(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, filter string) (result automation.JobStreamListResultIterator, err error)
}

var _ TestJobStreamsClientAPI = (*automation.TestJobStreamsClient)(nil)

// TestJobClientAPI contains the set of methods on the TestJobClient type.
type TestJobClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, parameters automation.TestJobCreateParameters) (result automation.TestJob, err error)
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (result automation.TestJob, err error)
	Resume(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (result autorest.Response, err error)
	Stop(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (result autorest.Response, err error)
	Suspend(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (result autorest.Response, err error)
}

var _ TestJobClientAPI = (*automation.TestJobClient)(nil)

// ScheduleClientAPI contains the set of methods on the ScheduleClient type.
type ScheduleClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, scheduleName string, parameters automation.ScheduleCreateOrUpdateParameters) (result automation.Schedule, err error)
	Delete(ctx context.Context, resourceGroupName string, automationAccountName string, scheduleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, scheduleName string) (result automation.Schedule, err error)
	ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.ScheduleListResultPage, err error)
	ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.ScheduleListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, automationAccountName string, scheduleName string, parameters automation.ScheduleUpdateParameters) (result automation.Schedule, err error)
}

var _ ScheduleClientAPI = (*automation.ScheduleClient)(nil)

// VariableClientAPI contains the set of methods on the VariableClient type.
type VariableClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, variableName string, parameters automation.VariableCreateOrUpdateParameters) (result automation.Variable, err error)
	Delete(ctx context.Context, resourceGroupName string, automationAccountName string, variableName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, variableName string) (result automation.Variable, err error)
	ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.VariableListResultPage, err error)
	ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.VariableListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, automationAccountName string, variableName string, parameters automation.VariableUpdateParameters) (result automation.Variable, err error)
}

var _ VariableClientAPI = (*automation.VariableClient)(nil)

// WebhookClientAPI contains the set of methods on the WebhookClient type.
type WebhookClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, webhookName string, parameters automation.WebhookCreateOrUpdateParameters) (result automation.Webhook, err error)
	Delete(ctx context.Context, resourceGroupName string, automationAccountName string, webhookName string) (result autorest.Response, err error)
	GenerateURI(ctx context.Context, resourceGroupName string, automationAccountName string) (result automation.String, err error)
	Get(ctx context.Context, resourceGroupName string, automationAccountName string, webhookName string) (result automation.Webhook, err error)
	ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (result automation.WebhookListResultPage, err error)
	ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string, filter string) (result automation.WebhookListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, automationAccountName string, webhookName string, parameters automation.WebhookUpdateParameters) (result automation.Webhook, err error)
}

var _ WebhookClientAPI = (*automation.WebhookClient)(nil)
