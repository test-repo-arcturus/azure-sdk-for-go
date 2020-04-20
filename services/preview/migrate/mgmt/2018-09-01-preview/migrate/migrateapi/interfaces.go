package migrateapi

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
	"github.com/test-repo-arcturus/azure-sdk-for-go/services/preview/migrate/mgmt/2018-09-01-preview/migrate"
)

// DatabaseInstancesClientAPI contains the set of methods on the DatabaseInstancesClient type.
type DatabaseInstancesClientAPI interface {
	EnumerateDatabaseInstances(ctx context.Context, resourceGroupName string, migrateProjectName string, continuationToken string, pageSize *int32) (result migrate.DatabaseInstanceCollection, err error)
	GetDatabaseInstance(ctx context.Context, resourceGroupName string, migrateProjectName string, databaseInstanceName string) (result migrate.DatabaseInstance, err error)
}

var _ DatabaseInstancesClientAPI = (*migrate.DatabaseInstancesClient)(nil)

// DatabasesClientAPI contains the set of methods on the DatabasesClient type.
type DatabasesClientAPI interface {
	EnumerateDatabases(ctx context.Context, resourceGroupName string, migrateProjectName string, continuationToken string, pageSize *int32) (result migrate.DatabaseCollection, err error)
	GetDatabase(ctx context.Context, resourceGroupName string, migrateProjectName string, databaseName string) (result migrate.Database, err error)
}

var _ DatabasesClientAPI = (*migrate.DatabasesClient)(nil)

// EventsClientAPI contains the set of methods on the EventsClient type.
type EventsClientAPI interface {
	DeleteEvent(ctx context.Context, resourceGroupName string, migrateProjectName string, eventName string) (result autorest.Response, err error)
	EnumerateEvents(ctx context.Context, resourceGroupName string, migrateProjectName string, continuationToken string, pageSize *int32) (result migrate.EventCollection, err error)
	GetEvent(ctx context.Context, resourceGroupName string, migrateProjectName string, eventName string) (result migrate.Event, err error)
}

var _ EventsClientAPI = (*migrate.EventsClient)(nil)

// MachinesClientAPI contains the set of methods on the MachinesClient type.
type MachinesClientAPI interface {
	EnumerateMachines(ctx context.Context, resourceGroupName string, migrateProjectName string, continuationToken string, pageSize *int32) (result migrate.MachineCollection, err error)
	GetMachine(ctx context.Context, resourceGroupName string, migrateProjectName string, machineName string) (result migrate.Machine, err error)
}

var _ MachinesClientAPI = (*migrate.MachinesClient)(nil)

// ProjectsClientAPI contains the set of methods on the ProjectsClient type.
type ProjectsClientAPI interface {
	DeleteMigrateProject(ctx context.Context, resourceGroupName string, migrateProjectName string) (result autorest.Response, err error)
	GetMigrateProject(ctx context.Context, resourceGroupName string, migrateProjectName string) (result migrate.Project, err error)
	PatchMigrateProject(ctx context.Context, resourceGroupName string, migrateProjectName string, body migrate.Project) (result migrate.Project, err error)
	PutMigrateProject(ctx context.Context, resourceGroupName string, migrateProjectName string, body migrate.Project) (result migrate.Project, err error)
	RefreshMigrateProjectSummary(ctx context.Context, resourceGroupName string, migrateProjectName string, input migrate.RefreshSummaryInput) (result migrate.RefreshSummaryResult, err error)
	RegisterTool(ctx context.Context, resourceGroupName string, migrateProjectName string, input migrate.RegisterToolInput) (result migrate.RegistrationResult, err error)
}

var _ ProjectsClientAPI = (*migrate.ProjectsClient)(nil)

// SolutionsClientAPI contains the set of methods on the SolutionsClient type.
type SolutionsClientAPI interface {
	CleanupSolutionData(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string) (result autorest.Response, err error)
	DeleteSolution(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string) (result autorest.Response, err error)
	EnumerateSolutions(ctx context.Context, resourceGroupName string, migrateProjectName string) (result migrate.SolutionsCollection, err error)
	GetConfig(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string) (result migrate.SolutionConfig, err error)
	GetSolution(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string) (result migrate.Solution, err error)
	PatchSolution(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string, solutionInput migrate.Solution) (result migrate.Solution, err error)
	PutSolution(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string, solutionInput migrate.Solution) (result migrate.Solution, err error)
}

var _ SolutionsClientAPI = (*migrate.SolutionsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result migrate.OperationResultList, err error)
}

var _ OperationsClientAPI = (*migrate.OperationsClient)(nil)
