package hybridcompute

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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/test-repo-arcturus/azure-sdk-for-go/services/preview/hybridcompute/mgmt/2019-03-18-preview/hybridcompute"

// InstanceViewTypes enumerates the values for instance view types.
type InstanceViewTypes string

const (
	// InstanceView ...
	InstanceView InstanceViewTypes = "instanceView"
)

// PossibleInstanceViewTypesValues returns an array of possible values for the InstanceViewTypes const type.
func PossibleInstanceViewTypesValues() []InstanceViewTypes {
	return []InstanceViewTypes{InstanceView}
}

// StatusTypes enumerates the values for status types.
type StatusTypes string

const (
	// Connected ...
	Connected StatusTypes = "Connected"
	// Disconnected ...
	Disconnected StatusTypes = "Disconnected"
	// Error ...
	Error StatusTypes = "Error"
)

// PossibleStatusTypesValues returns an array of possible values for the StatusTypes const type.
func PossibleStatusTypesValues() []StatusTypes {
	return []StatusTypes{Connected, Disconnected, Error}
}

// ErrorDetail ...
type ErrorDetail struct {
	// Code - The error's code.
	Code *string `json:"code,omitempty"`
	// Message - A human readable error message.
	Message *string `json:"message,omitempty"`
	// Target - Indicates which property in the request is responsible for the error.
	Target *string `json:"target,omitempty"`
	// Details - Additional error details.
	Details *[]ErrorDetail `json:"details,omitempty"`
}

// ErrorResponse contains details when the response code indicates an error.
type ErrorResponse struct {
	// Error - The error details.
	Error *ErrorDetail `json:"error,omitempty"`
}

// Identity ...
type Identity struct {
	// Type - The identity type.
	Type *string `json:"type,omitempty"`
	// PrincipalID - READ-ONLY; The identity's principal id.
	PrincipalID *string `json:"principalId,omitempty"`
	// TenantID - READ-ONLY; The identity's tenant id.
	TenantID *string `json:"tenantId,omitempty"`
}

// Machine describes a hybrid machine.
type Machine struct {
	autorest.Response `json:"-"`
	// MachineProperties - Hybrid Compute Machine properties
	*MachineProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
	// Identity - Hybrid Compute Machine Managed Identity
	*Identity `json:"identity,omitempty"`
}

// MarshalJSON is the custom marshaler for Machine.
func (mVar Machine) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if mVar.MachineProperties != nil {
		objectMap["properties"] = mVar.MachineProperties
	}
	if mVar.Location != nil {
		objectMap["location"] = mVar.Location
	}
	if mVar.Tags != nil {
		objectMap["tags"] = mVar.Tags
	}
	if mVar.Identity != nil {
		objectMap["identity"] = mVar.Identity
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Machine struct.
func (mVar *Machine) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var machineProperties MachineProperties
				err = json.Unmarshal(*v, &machineProperties)
				if err != nil {
					return err
				}
				mVar.MachineProperties = &machineProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				mVar.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				mVar.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				mVar.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				mVar.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				mVar.Tags = tags
			}
		case "identity":
			if v != nil {
				var identity Identity
				err = json.Unmarshal(*v, &identity)
				if err != nil {
					return err
				}
				mVar.Identity = &identity
			}
		}
	}

	return nil
}

// MachineListResult the List hybrid machine operation response.
type MachineListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of hybrid machines.
	Value *[]Machine `json:"value,omitempty"`
	// NextLink - The URI to fetch the next page of VMs. Call ListNext() with this URI to fetch the next page of hybrid machines.
	NextLink *string `json:"nextLink,omitempty"`
}

// MachineListResultIterator provides access to a complete listing of Machine values.
type MachineListResultIterator struct {
	i    int
	page MachineListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *MachineListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachineListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *MachineListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter MachineListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter MachineListResultIterator) Response() MachineListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter MachineListResultIterator) Value() Machine {
	if !iter.page.NotDone() {
		return Machine{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the MachineListResultIterator type.
func NewMachineListResultIterator(page MachineListResultPage) MachineListResultIterator {
	return MachineListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (mlr MachineListResult) IsEmpty() bool {
	return mlr.Value == nil || len(*mlr.Value) == 0
}

// machineListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (mlr MachineListResult) machineListResultPreparer(ctx context.Context) (*http.Request, error) {
	if mlr.NextLink == nil || len(to.String(mlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(mlr.NextLink)))
}

// MachineListResultPage contains a page of Machine values.
type MachineListResultPage struct {
	fn  func(context.Context, MachineListResult) (MachineListResult, error)
	mlr MachineListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *MachineListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachineListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.mlr)
	if err != nil {
		return err
	}
	page.mlr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *MachineListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page MachineListResultPage) NotDone() bool {
	return !page.mlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page MachineListResultPage) Response() MachineListResult {
	return page.mlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page MachineListResultPage) Values() []Machine {
	if page.mlr.IsEmpty() {
		return nil
	}
	return *page.mlr.Value
}

// Creates a new instance of the MachineListResultPage type.
func NewMachineListResultPage(getNextPage func(context.Context, MachineListResult) (MachineListResult, error)) MachineListResultPage {
	return MachineListResultPage{fn: getNextPage}
}

// MachineProperties describes the properties of a hybrid machine.
type MachineProperties struct {
	// OsProfile - Specifies the operating system settings for the hybrid machine.
	OsProfile *OSProfile `json:"osProfile,omitempty"`
	// ProvisioningState - READ-ONLY; The provisioning state, which only appears in the response.
	ProvisioningState *string `json:"provisioningState,omitempty"`
	// Status - READ-ONLY; The status of the hybrid machine agent. Possible values include: 'Connected', 'Disconnected', 'Error'
	Status StatusTypes `json:"status,omitempty"`
	// LastStatusChange - READ-ONLY; The time of the last status change.
	LastStatusChange *date.Time `json:"lastStatusChange,omitempty"`
	// ErrorDetails - READ-ONLY; Details about the error state.
	ErrorDetails *[]ErrorDetail `json:"errorDetails,omitempty"`
	// AgentVersion - READ-ONLY; The hybrid machine agent full version.
	AgentVersion *string `json:"agentVersion,omitempty"`
	// VMID - READ-ONLY; Specifies the hybrid machine unique ID.
	VMID *string `json:"vmId,omitempty"`
	// DisplayName - READ-ONLY; Specifies the hybrid machine display name.
	DisplayName *string `json:"displayName,omitempty"`
	// MachineFqdn - READ-ONLY; Specifies the hybrid machine FQDN.
	MachineFqdn *string `json:"machineFqdn,omitempty"`
	// PhysicalLocation - Resource's Physical Location
	PhysicalLocation *string `json:"physicalLocation,omitempty"`
	// ClientPublicKey - Public Key that the client provides to be used during initial resource onboarding
	ClientPublicKey *string `json:"clientPublicKey,omitempty"`
	// OsName - READ-ONLY; The Operating System running on the hybrid machine.
	OsName *string `json:"osName,omitempty"`
	// OsVersion - READ-ONLY; The version of Operating System running on the hybrid machine.
	OsVersion *string `json:"osVersion,omitempty"`
}

// MachineReconnect describes a hybrid machine reconnect.
type MachineReconnect struct {
	// MachineReconnectProperties - Hybrid Compute Machine properties
	*MachineReconnectProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for MachineReconnect.
func (mr MachineReconnect) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if mr.MachineReconnectProperties != nil {
		objectMap["properties"] = mr.MachineReconnectProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for MachineReconnect struct.
func (mr *MachineReconnect) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var machineReconnectProperties MachineReconnectProperties
				err = json.Unmarshal(*v, &machineReconnectProperties)
				if err != nil {
					return err
				}
				mr.MachineReconnectProperties = &machineReconnectProperties
			}
		}
	}

	return nil
}

// MachineReconnectProperties describes the properties required to reconnect a hybrid machine.
type MachineReconnectProperties struct {
	// VMID - Specifies the hybrid machine unique ID.
	VMID *string `json:"vmId,omitempty"`
	// ClientPublicKey - Public Key that the client provides to be used during initial resource onboarding.
	ClientPublicKey *string `json:"clientPublicKey,omitempty"`
}

// MachineUpdate describes a hybrid machine Update.
type MachineUpdate struct {
	// Identity - Hybrid Compute Machine Managed Identity
	*Identity `json:"identity,omitempty"`
	// MachineUpdateProperties - Hybrid Compute Machine properties
	*MachineUpdateProperties `json:"properties,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for MachineUpdate.
func (mu MachineUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if mu.Identity != nil {
		objectMap["identity"] = mu.Identity
	}
	if mu.MachineUpdateProperties != nil {
		objectMap["properties"] = mu.MachineUpdateProperties
	}
	if mu.Tags != nil {
		objectMap["tags"] = mu.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for MachineUpdate struct.
func (mu *MachineUpdate) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "identity":
			if v != nil {
				var identity Identity
				err = json.Unmarshal(*v, &identity)
				if err != nil {
					return err
				}
				mu.Identity = &identity
			}
		case "properties":
			if v != nil {
				var machineUpdateProperties MachineUpdateProperties
				err = json.Unmarshal(*v, &machineUpdateProperties)
				if err != nil {
					return err
				}
				mu.MachineUpdateProperties = &machineUpdateProperties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				mu.Tags = tags
			}
		}
	}

	return nil
}

// MachineUpdateProperties describes the ARM updatable properties of a hybrid machine.
type MachineUpdateProperties struct {
	// PhysicalLocation - Resource's Physical Location
	PhysicalLocation *string `json:"physicalLocation,omitempty"`
}

// OperationListResult the List Compute Operation operation response.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; The list of compute operations
	Value *[]OperationValue `json:"value,omitempty"`
}

// OperationValue describes the properties of a Compute Operation value.
type OperationValue struct {
	// Origin - READ-ONLY; The origin of the compute operation.
	Origin *string `json:"origin,omitempty"`
	// Name - READ-ONLY; The name of the compute operation.
	Name *string `json:"name,omitempty"`
	// OperationValueDisplay - Display properties
	*OperationValueDisplay `json:"display,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationValue.
func (ov OperationValue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ov.OperationValueDisplay != nil {
		objectMap["display"] = ov.OperationValueDisplay
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for OperationValue struct.
func (ov *OperationValue) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "origin":
			if v != nil {
				var origin string
				err = json.Unmarshal(*v, &origin)
				if err != nil {
					return err
				}
				ov.Origin = &origin
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ov.Name = &name
			}
		case "display":
			if v != nil {
				var operationValueDisplay OperationValueDisplay
				err = json.Unmarshal(*v, &operationValueDisplay)
				if err != nil {
					return err
				}
				ov.OperationValueDisplay = &operationValueDisplay
			}
		}
	}

	return nil
}

// OperationValueDisplay describes the properties of a Hybrid Compute Operation Value Display.
type OperationValueDisplay struct {
	// Operation - READ-ONLY; The display name of the compute operation.
	Operation *string `json:"operation,omitempty"`
	// Resource - READ-ONLY; The display name of the resource the operation applies to.
	Resource *string `json:"resource,omitempty"`
	// Description - READ-ONLY; The description of the operation.
	Description *string `json:"description,omitempty"`
	// Provider - READ-ONLY; The resource provider for the operation.
	Provider *string `json:"provider,omitempty"`
}

// OSProfile specifies the operating system settings for the hybrid machine.
type OSProfile struct {
	// ComputerName - READ-ONLY; Specifies the host OS name of the hybrid machine.
	ComputerName *string `json:"computerName,omitempty"`
}

// Resource the Resource model definition.
type Resource struct {
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
	// Identity - Hybrid Compute Machine Managed Identity
	*Identity `json:"identity,omitempty"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	if r.Identity != nil {
		objectMap["identity"] = r.Identity
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Resource struct.
func (r *Resource) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				r.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				r.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				r.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				r.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				r.Tags = tags
			}
		case "identity":
			if v != nil {
				var identity Identity
				err = json.Unmarshal(*v, &identity)
				if err != nil {
					return err
				}
				r.Identity = &identity
			}
		}
	}

	return nil
}

// UpdateResource the Update Resource model definition.
type UpdateResource struct {
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for UpdateResource.
func (ur UpdateResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ur.Tags != nil {
		objectMap["tags"] = ur.Tags
	}
	return json.Marshal(objectMap)
}
