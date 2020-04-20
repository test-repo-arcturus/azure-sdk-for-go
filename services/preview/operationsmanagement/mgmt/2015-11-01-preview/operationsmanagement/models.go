package operationsmanagement

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
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/test-repo-arcturus/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement"

// ArmTemplateParameter parameter to pass to ARM template
type ArmTemplateParameter struct {
	// Name - name of the parameter.
	Name *string `json:"name,omitempty"`
	// Value - value for the parameter. In Jtoken
	Value *string `json:"value,omitempty"`
}

// CodeMessageError the error body contract.
type CodeMessageError struct {
	// Error - The error details for a failed request.
	Error *CodeMessageErrorError `json:"error,omitempty"`
}

// CodeMessageErrorError the error details for a failed request.
type CodeMessageErrorError struct {
	// Code - The error type.
	Code *string `json:"code,omitempty"`
	// Message - The error message.
	Message *string `json:"message,omitempty"`
}

// ManagementAssociation the container for solution.
type ManagementAssociation struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Properties - Properties for ManagementAssociation object supported by the OperationsManagement resource provider.
	Properties *ManagementAssociationProperties `json:"properties,omitempty"`
}

// ManagementAssociationProperties managementAssociation properties supported by the OperationsManagement
// resource provider.
type ManagementAssociationProperties struct {
	// ApplicationID - The applicationId of the appliance for this association.
	ApplicationID *string `json:"applicationId,omitempty"`
}

// ManagementAssociationPropertiesList the list of ManagementAssociation response
type ManagementAssociationPropertiesList struct {
	autorest.Response `json:"-"`
	// Value - List of Management Association properties within the subscription.
	Value *[]ManagementAssociation `json:"value,omitempty"`
}

// ManagementConfiguration the container for solution.
type ManagementConfiguration struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Properties - Properties for ManagementConfiguration object supported by the OperationsManagement resource provider.
	Properties *ManagementConfigurationProperties `json:"properties,omitempty"`
}

// ManagementConfigurationProperties managementConfiguration properties supported by the
// OperationsManagement resource provider.
type ManagementConfigurationProperties struct {
	// ApplicationID - The applicationId of the appliance for this Management.
	ApplicationID *string `json:"applicationId,omitempty"`
	// ParentResourceType - The type of the parent resource.
	ParentResourceType *string `json:"parentResourceType,omitempty"`
	// Parameters - Parameters to run the ARM template
	Parameters *[]ArmTemplateParameter `json:"parameters,omitempty"`
	// ProvisioningState - READ-ONLY; The provisioning state for the ManagementConfiguration.
	ProvisioningState *string `json:"provisioningState,omitempty"`
	// Template - The Json object containing the ARM template to deploy
	Template interface{} `json:"template,omitempty"`
}

// ManagementConfigurationPropertiesList the list of ManagementConfiguration response
type ManagementConfigurationPropertiesList struct {
	autorest.Response `json:"-"`
	// Value - List of Management Configuration properties within the subscription.
	Value *[]ManagementConfiguration `json:"value,omitempty"`
}

// Operation supported operation of OperationsManagement resource provider.
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - Display metadata associated with the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay display metadata associated with the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft OperationsManagement.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Type of operation: get, read, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result of the request to list solution operations.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of solution operations supported by the OperationsManagement resource provider.
	Value *[]Operation `json:"value,omitempty"`
}

// Solution the container for solution.
type Solution struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
	// Plan - Plan for solution object supported by the OperationsManagement resource provider.
	Plan *SolutionPlan `json:"plan,omitempty"`
	// Properties - Properties for solution object supported by the OperationsManagement resource provider.
	Properties *SolutionProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for Solution.
func (s Solution) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if s.Location != nil {
		objectMap["location"] = s.Location
	}
	if s.Tags != nil {
		objectMap["tags"] = s.Tags
	}
	if s.Plan != nil {
		objectMap["plan"] = s.Plan
	}
	if s.Properties != nil {
		objectMap["properties"] = s.Properties
	}
	return json.Marshal(objectMap)
}

// SolutionPatch the properties of a Solution that can be patched.
type SolutionPatch struct {
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for SolutionPatch.
func (sp SolutionPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sp.Tags != nil {
		objectMap["tags"] = sp.Tags
	}
	return json.Marshal(objectMap)
}

// SolutionPlan plan for solution object supported by the OperationsManagement resource provider.
type SolutionPlan struct {
	// Name - name of the solution to be created. For Microsoft published solution it should be in the format of solutionType(workspaceName). SolutionType part is case sensitive. For third party solution, it can be anything.
	Name *string `json:"name,omitempty"`
	// Publisher - Publisher name. For gallery solution, it is Microsoft.
	Publisher *string `json:"publisher,omitempty"`
	// PromotionCode - promotionCode, Not really used now, can you left as empty
	PromotionCode *string `json:"promotionCode,omitempty"`
	// Product - name of the solution to enabled/add. For Microsoft published gallery solution it should be in the format of OMSGallery/<solutionType>. This is case sensitive
	Product *string `json:"product,omitempty"`
}

// SolutionProperties solution properties supported by the OperationsManagement resource provider.
type SolutionProperties struct {
	// WorkspaceResourceID - The azure resourceId for the workspace where the solution will be deployed/enabled.
	WorkspaceResourceID *string `json:"workspaceResourceId,omitempty"`
	// ProvisioningState - READ-ONLY; The provisioning state for the solution.
	ProvisioningState *string `json:"provisioningState,omitempty"`
	// ContainedResources - The azure resources that will be contained within the solutions. They will be locked and gets deleted automatically when the solution is deleted.
	ContainedResources *[]string `json:"containedResources,omitempty"`
	// ReferencedResources - The resources that will be referenced from this solution. Deleting any of those solution out of band will break the solution.
	ReferencedResources *[]string `json:"referencedResources,omitempty"`
}

// SolutionPropertiesList the list of solution response
type SolutionPropertiesList struct {
	autorest.Response `json:"-"`
	// Value - List of solution properties within the subscription.
	Value *[]Solution `json:"value,omitempty"`
}

// SolutionsCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type SolutionsCreateOrUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *SolutionsCreateOrUpdateFuture) Result(client SolutionsClient) (s Solution, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationsmanagement.SolutionsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("operationsmanagement.SolutionsCreateOrUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if s.Response.Response, err = future.GetResult(sender); err == nil && s.Response.Response.StatusCode != http.StatusNoContent {
		s, err = client.CreateOrUpdateResponder(s.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "operationsmanagement.SolutionsCreateOrUpdateFuture", "Result", s.Response.Response, "Failure responding to request")
		}
	}
	return
}

// SolutionsDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type SolutionsDeleteFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *SolutionsDeleteFuture) Result(client SolutionsClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationsmanagement.SolutionsDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("operationsmanagement.SolutionsDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// SolutionsUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type SolutionsUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *SolutionsUpdateFuture) Result(client SolutionsClient) (s Solution, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationsmanagement.SolutionsUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("operationsmanagement.SolutionsUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if s.Response.Response, err = future.GetResult(sender); err == nil && s.Response.Response.StatusCode != http.StatusNoContent {
		s, err = client.UpdateResponder(s.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "operationsmanagement.SolutionsUpdateFuture", "Result", s.Response.Response, "Failure responding to request")
		}
	}
	return
}
