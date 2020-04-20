package mysqlapi

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
	"github.com/test-repo-arcturus/azure-sdk-for-go/services/preview/mysql/mgmt/2017-12-01-preview/mysql"
)

// ServersClientAPI contains the set of methods on the ServersClient type.
type ServersClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, serverName string, parameters mysql.ServerForCreate) (result mysql.ServersCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string) (result mysql.ServersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string) (result mysql.Server, err error)
	List(ctx context.Context) (result mysql.ServerListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result mysql.ServerListResult, err error)
	Restart(ctx context.Context, resourceGroupName string, serverName string) (result mysql.ServersRestartFuture, err error)
	Update(ctx context.Context, resourceGroupName string, serverName string, parameters mysql.ServerUpdateParameters) (result mysql.ServersUpdateFuture, err error)
}

var _ ServersClientAPI = (*mysql.ServersClient)(nil)

// ReplicasClientAPI contains the set of methods on the ReplicasClient type.
type ReplicasClientAPI interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mysql.ServerListResult, err error)
}

var _ ReplicasClientAPI = (*mysql.ReplicasClient)(nil)

// FirewallRulesClientAPI contains the set of methods on the FirewallRulesClient type.
type FirewallRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters mysql.FirewallRule) (result mysql.FirewallRulesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (result mysql.FirewallRulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (result mysql.FirewallRule, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mysql.FirewallRuleListResult, err error)
}

var _ FirewallRulesClientAPI = (*mysql.FirewallRulesClient)(nil)

// VirtualNetworkRulesClientAPI contains the set of methods on the VirtualNetworkRulesClient type.
type VirtualNetworkRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, parameters mysql.VirtualNetworkRule) (result mysql.VirtualNetworkRulesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string) (result mysql.VirtualNetworkRulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string) (result mysql.VirtualNetworkRule, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mysql.VirtualNetworkRuleListResultPage, err error)
	ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result mysql.VirtualNetworkRuleListResultIterator, err error)
}

var _ VirtualNetworkRulesClientAPI = (*mysql.VirtualNetworkRulesClient)(nil)

// DatabasesClientAPI contains the set of methods on the DatabasesClient type.
type DatabasesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters mysql.Database) (result mysql.DatabasesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result mysql.DatabasesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result mysql.Database, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mysql.DatabaseListResult, err error)
}

var _ DatabasesClientAPI = (*mysql.DatabasesClient)(nil)

// ConfigurationsClientAPI contains the set of methods on the ConfigurationsClient type.
type ConfigurationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, configurationName string, parameters mysql.Configuration) (result mysql.ConfigurationsCreateOrUpdateFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, configurationName string) (result mysql.Configuration, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mysql.ConfigurationListResult, err error)
}

var _ ConfigurationsClientAPI = (*mysql.ConfigurationsClient)(nil)

// LogFilesClientAPI contains the set of methods on the LogFilesClient type.
type LogFilesClientAPI interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mysql.LogFileListResult, err error)
}

var _ LogFilesClientAPI = (*mysql.LogFilesClient)(nil)

// ServerAdministratorsClientAPI contains the set of methods on the ServerAdministratorsClient type.
type ServerAdministratorsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, properties mysql.ServerAdministratorResource) (result mysql.ServerAdministratorsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string) (result mysql.ServerAdministratorsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string) (result mysql.ServerAdministratorResource, err error)
	List(ctx context.Context, resourceGroupName string, serverName string) (result mysql.ServerAdministratorResourceListResult, err error)
}

var _ ServerAdministratorsClientAPI = (*mysql.ServerAdministratorsClient)(nil)

// LocationBasedPerformanceTierClientAPI contains the set of methods on the LocationBasedPerformanceTierClient type.
type LocationBasedPerformanceTierClientAPI interface {
	List(ctx context.Context, locationName string) (result mysql.PerformanceTierListResult, err error)
}

var _ LocationBasedPerformanceTierClientAPI = (*mysql.LocationBasedPerformanceTierClient)(nil)

// CheckNameAvailabilityClientAPI contains the set of methods on the CheckNameAvailabilityClient type.
type CheckNameAvailabilityClientAPI interface {
	Execute(ctx context.Context, nameAvailabilityRequest mysql.NameAvailabilityRequest) (result mysql.NameAvailability, err error)
}

var _ CheckNameAvailabilityClientAPI = (*mysql.CheckNameAvailabilityClient)(nil)

// ServerSecurityAlertPoliciesClientAPI contains the set of methods on the ServerSecurityAlertPoliciesClient type.
type ServerSecurityAlertPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters mysql.ServerSecurityAlertPolicy) (result mysql.ServerSecurityAlertPoliciesCreateOrUpdateFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string) (result mysql.ServerSecurityAlertPolicy, err error)
}

var _ ServerSecurityAlertPoliciesClientAPI = (*mysql.ServerSecurityAlertPoliciesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result mysql.OperationListResult, err error)
}

var _ OperationsClientAPI = (*mysql.OperationsClient)(nil)
