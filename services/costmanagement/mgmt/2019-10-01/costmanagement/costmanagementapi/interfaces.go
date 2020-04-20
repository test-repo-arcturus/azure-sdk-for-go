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
	"github.com/test-repo-arcturus/azure-sdk-for-go/services/costmanagement/mgmt/2019-10-01/costmanagement"
	"github.com/Azure/go-autorest/autorest"
)

// DimensionsClientAPI contains the set of methods on the DimensionsClient type.
type DimensionsClientAPI interface {
	List(ctx context.Context, scope string, filter string, expand string, skiptoken string, top *int32) (result costmanagement.DimensionsListResult, err error)
}

var _ DimensionsClientAPI = (*costmanagement.DimensionsClient)(nil)

// QueryClientAPI contains the set of methods on the QueryClient type.
type QueryClientAPI interface {
	Usage(ctx context.Context, scope string, parameters costmanagement.QueryDefinition) (result costmanagement.QueryResult, err error)
}

var _ QueryClientAPI = (*costmanagement.QueryClient)(nil)

// ExportsClientAPI contains the set of methods on the ExportsClient type.
type ExportsClientAPI interface {
	CreateOrUpdate(ctx context.Context, scope string, exportName string, parameters costmanagement.Export) (result costmanagement.Export, err error)
	Delete(ctx context.Context, scope string, exportName string) (result autorest.Response, err error)
	Execute(ctx context.Context, scope string, exportName string) (result autorest.Response, err error)
	Get(ctx context.Context, scope string, exportName string) (result costmanagement.Export, err error)
	GetExecutionHistory(ctx context.Context, scope string, exportName string) (result costmanagement.ExportExecutionListResult, err error)
	List(ctx context.Context, scope string) (result costmanagement.ExportListResult, err error)
}

var _ ExportsClientAPI = (*costmanagement.ExportsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result costmanagement.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result costmanagement.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*costmanagement.OperationsClient)(nil)
