package costmanagementapi

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
	"github.com/test-repo-arcturus/azure-sdk-for-go/services/preview/costmanagement/mgmt/2018-08-01-preview/costmanagement"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	QueryBillingAccount(ctx context.Context, billingAccountID string, parameters costmanagement.ReportDefinition) (result costmanagement.QueryResult, err error)
	QueryResourceGroup(ctx context.Context, resourceGroupName string, parameters costmanagement.ReportDefinition) (result costmanagement.QueryResult, err error)
	QuerySubscription(ctx context.Context, parameters costmanagement.ReportDefinition) (result costmanagement.QueryResult, err error)
}

var _ BaseClientAPI = (*costmanagement.BaseClient)(nil)

// ReportsClientAPI contains the set of methods on the ReportsClient type.
type ReportsClientAPI interface {
	CreateOrUpdate(ctx context.Context, reportName string, parameters costmanagement.Report) (result costmanagement.Report, err error)
	CreateOrUpdateByBillingAccount(ctx context.Context, billingAccountID string, reportName string, parameters costmanagement.Report) (result costmanagement.Report, err error)
	CreateOrUpdateByDepartment(ctx context.Context, departmentID string, reportName string, parameters costmanagement.Report) (result costmanagement.Report, err error)
	CreateOrUpdateByResourceGroupName(ctx context.Context, resourceGroupName string, reportName string, parameters costmanagement.Report) (result costmanagement.Report, err error)
	Delete(ctx context.Context, reportName string) (result autorest.Response, err error)
	DeleteByBillingAccount(ctx context.Context, billingAccountID string, reportName string) (result autorest.Response, err error)
	DeleteByDepartment(ctx context.Context, departmentID string, reportName string) (result autorest.Response, err error)
	DeleteByResourceGroupName(ctx context.Context, resourceGroupName string, reportName string) (result autorest.Response, err error)
	Execute(ctx context.Context, reportName string) (result autorest.Response, err error)
	ExecuteByBillingAccount(ctx context.Context, billingAccountID string, reportName string) (result autorest.Response, err error)
	ExecuteByDepartment(ctx context.Context, departmentID string, reportName string) (result autorest.Response, err error)
	ExecuteByResourceGroupName(ctx context.Context, resourceGroupName string, reportName string) (result autorest.Response, err error)
	Get(ctx context.Context, reportName string) (result costmanagement.Report, err error)
	GetByBillingAccount(ctx context.Context, billingAccountID string, reportName string) (result costmanagement.Report, err error)
	GetByDepartment(ctx context.Context, departmentID string, reportName string) (result costmanagement.Report, err error)
	GetByResourceGroupName(ctx context.Context, resourceGroupName string, reportName string) (result costmanagement.Report, err error)
	GetExecutionHistory(ctx context.Context, reportName string) (result costmanagement.ReportExecutionListResult, err error)
	GetExecutionHistoryByBillingAccount(ctx context.Context, billingAccountID string, reportName string) (result costmanagement.ReportExecutionListResult, err error)
	GetExecutionHistoryByDepartment(ctx context.Context, departmentID string, reportName string) (result costmanagement.ReportExecutionListResult, err error)
	GetExecutionHistoryByResourceGroupName(ctx context.Context, resourceGroupName string, reportName string) (result costmanagement.ReportExecutionListResult, err error)
	List(ctx context.Context) (result costmanagement.ReportListResult, err error)
	ListByBillingAccount(ctx context.Context, billingAccountID string) (result costmanagement.ReportListResult, err error)
	ListByDepartment(ctx context.Context, departmentID string) (result costmanagement.ReportListResult, err error)
	ListByResourceGroupName(ctx context.Context, resourceGroupName string) (result costmanagement.ReportListResult, err error)
}

var _ ReportsClientAPI = (*costmanagement.ReportsClient)(nil)

// BillingAccountDimensionsClientAPI contains the set of methods on the BillingAccountDimensionsClient type.
type BillingAccountDimensionsClientAPI interface {
	List(ctx context.Context, billingAccountID string, filter string, expand string, skiptoken string, top *int32) (result costmanagement.DimensionsListResult, err error)
}

var _ BillingAccountDimensionsClientAPI = (*costmanagement.BillingAccountDimensionsClient)(nil)

// SubscriptionDimensionsClientAPI contains the set of methods on the SubscriptionDimensionsClient type.
type SubscriptionDimensionsClientAPI interface {
	List(ctx context.Context, filter string, expand string, skiptoken string, top *int32) (result costmanagement.DimensionsListResult, err error)
}

var _ SubscriptionDimensionsClientAPI = (*costmanagement.SubscriptionDimensionsClient)(nil)

// ResourceGroupDimensionsClientAPI contains the set of methods on the ResourceGroupDimensionsClient type.
type ResourceGroupDimensionsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, filter string, expand string, skiptoken string, top *int32) (result costmanagement.DimensionsListResult, err error)
}

var _ ResourceGroupDimensionsClientAPI = (*costmanagement.ResourceGroupDimensionsClient)(nil)

// ConnectorClientAPI contains the set of methods on the ConnectorClient type.
type ConnectorClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, connectorName string, connector costmanagement.ConnectorDefinition) (result costmanagement.ConnectorDefinition, err error)
	Delete(ctx context.Context, resourceGroupName string, connectorName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, connectorName string) (result costmanagement.ConnectorDefinition, err error)
	List(ctx context.Context) (result costmanagement.ConnectorDefinitionListResult, err error)
	ListByResourceGroupName(ctx context.Context, resourceGroupName string) (result costmanagement.ConnectorDefinitionListResult, err error)
	Update(ctx context.Context, resourceGroupName string, connectorName string, connector costmanagement.ConnectorDefinition) (result costmanagement.ConnectorDefinition, err error)
}

var _ ConnectorClientAPI = (*costmanagement.ConnectorClient)(nil)

// AlertsClientAPI contains the set of methods on the AlertsClient type.
type AlertsClientAPI interface {
	GetAlertByManagementGroups(ctx context.Context, managementGroupID string, alertID string) (result costmanagement.Alert, err error)
	GetByAccount(ctx context.Context, billingAccountID string, enrollmentAccountID string, alertID string) (result costmanagement.Alert, err error)
	GetByDepartment(ctx context.Context, billingAccountID string, departmentID string, alertID string) (result costmanagement.Alert, err error)
	GetByEnrollment(ctx context.Context, billingAccountID string, alertID string) (result costmanagement.Alert, err error)
	GetByResourceGroupName(ctx context.Context, resourceGroupName string, alertID string) (result costmanagement.Alert, err error)
	GetBySubscription(ctx context.Context, alertID string) (result costmanagement.Alert, err error)
	List(ctx context.Context, filter string, skiptoken string, top *int32) (result costmanagement.AlertListResultPage, err error)
	ListComplete(ctx context.Context, filter string, skiptoken string, top *int32) (result costmanagement.AlertListResultIterator, err error)
	ListByAccount(ctx context.Context, billingAccountID string, enrollmentAccountID string, filter string, skiptoken string, top *int32) (result costmanagement.AlertListResultPage, err error)
	ListByAccountComplete(ctx context.Context, billingAccountID string, enrollmentAccountID string, filter string, skiptoken string, top *int32) (result costmanagement.AlertListResultIterator, err error)
	ListByDepartment(ctx context.Context, billingAccountID string, departmentID string, filter string, skiptoken string, top *int32) (result costmanagement.AlertListResultPage, err error)
	ListByDepartmentComplete(ctx context.Context, billingAccountID string, departmentID string, filter string, skiptoken string, top *int32) (result costmanagement.AlertListResultIterator, err error)
	ListByEnrollment(ctx context.Context, billingAccountID string, filter string, skiptoken string, top *int32) (result costmanagement.AlertListResultPage, err error)
	ListByEnrollmentComplete(ctx context.Context, billingAccountID string, filter string, skiptoken string, top *int32) (result costmanagement.AlertListResultIterator, err error)
	ListByManagementGroups(ctx context.Context, managementGroupID string, filter string, skiptoken string, top *int32) (result costmanagement.AlertListResultPage, err error)
	ListByManagementGroupsComplete(ctx context.Context, managementGroupID string, filter string, skiptoken string, top *int32) (result costmanagement.AlertListResultIterator, err error)
	ListByResourceGroupName(ctx context.Context, resourceGroupName string, filter string, skiptoken string, top *int32) (result costmanagement.AlertListResultPage, err error)
	ListByResourceGroupNameComplete(ctx context.Context, resourceGroupName string, filter string, skiptoken string, top *int32) (result costmanagement.AlertListResultIterator, err error)
	UpdateBillingAccountAlertStatus(ctx context.Context, billingAccountID string, alertID string, parameters costmanagement.Alert) (result costmanagement.Alert, err error)
	UpdateDepartmentsAlertStatus(ctx context.Context, billingAccountID string, departmentID string, alertID string, parameters costmanagement.Alert) (result costmanagement.Alert, err error)
	UpdateEnrollmentAccountAlertStatus(ctx context.Context, billingAccountID string, enrollmentAccountID string, alertID string, parameters costmanagement.Alert) (result costmanagement.Alert, err error)
	UpdateManagementGroupAlertStatus(ctx context.Context, managementGroupID string, alertID string, parameters costmanagement.Alert) (result costmanagement.Alert, err error)
	UpdateResourceGroupNameAlertStatus(ctx context.Context, resourceGroupName string, alertID string, parameters costmanagement.Alert) (result costmanagement.Alert, err error)
	UpdateSubscriptionAlertStatus(ctx context.Context, alertID string, parameters costmanagement.Alert) (result costmanagement.Alert, err error)
}

var _ AlertsClientAPI = (*costmanagement.AlertsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result costmanagement.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result costmanagement.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*costmanagement.OperationsClient)(nil)
