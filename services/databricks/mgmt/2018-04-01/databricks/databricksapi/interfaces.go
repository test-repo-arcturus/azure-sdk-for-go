package databricksapi

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
	"github.com/test-repo-arcturus/azure-sdk-for-go/services/databricks/mgmt/2018-04-01/databricks"
)

// WorkspacesClientAPI contains the set of methods on the WorkspacesClient type.
type WorkspacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, parameters databricks.Workspace, resourceGroupName string, workspaceName string) (result databricks.WorkspacesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string) (result databricks.WorkspacesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (result databricks.Workspace, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result databricks.WorkspaceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result databricks.WorkspaceListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result databricks.WorkspaceListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result databricks.WorkspaceListResultIterator, err error)
	Update(ctx context.Context, parameters databricks.WorkspaceUpdate, resourceGroupName string, workspaceName string) (result databricks.WorkspacesUpdateFuture, err error)
}

var _ WorkspacesClientAPI = (*databricks.WorkspacesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result databricks.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result databricks.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*databricks.OperationsClient)(nil)
