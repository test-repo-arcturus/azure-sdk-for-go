package containerregistry

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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/test-repo-arcturus/azure-sdk-for-go/services/containerregistry/mgmt/2017-03-01/containerregistry"

// PasswordName enumerates the values for password name.
type PasswordName string

const (
	// Password ...
	Password PasswordName = "password"
	// Password2 ...
	Password2 PasswordName = "password2"
)

// PossiblePasswordNameValues returns an array of possible values for the PasswordName const type.
func PossiblePasswordNameValues() []PasswordName {
	return []PasswordName{Password, Password2}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Creating, Succeeded}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Basic ...
	Basic SkuTier = "Basic"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{Basic}
}

// OperationDefinition the definition of a container registry operation.
type OperationDefinition struct {
	// Name - Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The display information for the container registry operation.
	Display *OperationDisplayDefinition `json:"display,omitempty"`
}

// OperationDisplayDefinition the display information for a container registry operation.
type OperationDisplayDefinition struct {
	// Provider - The resource provider name: Microsoft.ContainerRegistry.
	Provider *string `json:"provider,omitempty"`
	// Resource - The resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - The operation that users can perform.
	Operation *string `json:"operation,omitempty"`
	// Description - The description for the operation.
	Description *string `json:"description,omitempty"`
}

// OperationListResult the result of a request to list container registry operations.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of container registry operations. Since this list may be incomplete, the nextLink field should be used to request the next list of operations.
	Value *[]OperationDefinition `json:"value,omitempty"`
	// NextLink - The URI that can be used to request the next list of container registry operations.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of OperationDefinition values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultIterator.NextWithContext")
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
func (iter *OperationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() OperationDefinition {
	if !iter.page.NotDone() {
		return OperationDefinition{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListResultIterator type.
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return OperationListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of OperationDefinition values.
type OperationListResultPage struct {
	fn  func(context.Context, OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.olr)
	if err != nil {
		return err
	}
	page.olr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []OperationDefinition {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Creates a new instance of the OperationListResultPage type.
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return OperationListResultPage{fn: getNextPage}
}

// RegenerateCredentialParameters the parameters used to regenerate the login credential.
type RegenerateCredentialParameters struct {
	// Name - Specifies name of the password which should be regenerated -- password or password2. Possible values include: 'Password', 'Password2'
	Name PasswordName `json:"name,omitempty"`
}

// RegistriesCreateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type RegistriesCreateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *RegistriesCreateFuture) Result(client RegistriesClient) (r Registry, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesCreateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("containerregistry.RegistriesCreateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if r.Response.Response, err = future.GetResult(sender); err == nil && r.Response.Response.StatusCode != http.StatusNoContent {
		r, err = client.CreateResponder(r.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "containerregistry.RegistriesCreateFuture", "Result", r.Response.Response, "Failure responding to request")
		}
	}
	return
}

// Registry an object that represents a container registry.
type Registry struct {
	autorest.Response `json:"-"`
	// Sku - The SKU of the container registry.
	Sku *Sku `json:"sku,omitempty"`
	// RegistryProperties - The properties of the container registry.
	*RegistryProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
	// Location - The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`
	// Tags - The tags of the resource.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Registry.
func (r Registry) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Sku != nil {
		objectMap["sku"] = r.Sku
	}
	if r.RegistryProperties != nil {
		objectMap["properties"] = r.RegistryProperties
	}
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Registry struct.
func (r *Registry) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "sku":
			if v != nil {
				var sku Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				r.Sku = &sku
			}
		case "properties":
			if v != nil {
				var registryProperties RegistryProperties
				err = json.Unmarshal(*v, &registryProperties)
				if err != nil {
					return err
				}
				r.RegistryProperties = &registryProperties
			}
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
		}
	}

	return nil
}

// RegistryCreateParameters the parameters for creating a container registry.
type RegistryCreateParameters struct {
	// Tags - The tags for the container registry.
	Tags map[string]*string `json:"tags"`
	// Location - The location of the container registry. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`
	// Sku - The SKU of the container registry.
	Sku *Sku `json:"sku,omitempty"`
	// RegistryPropertiesCreateParameters - The properties that the container registry will be created with.
	*RegistryPropertiesCreateParameters `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for RegistryCreateParameters.
func (rcp RegistryCreateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rcp.Tags != nil {
		objectMap["tags"] = rcp.Tags
	}
	if rcp.Location != nil {
		objectMap["location"] = rcp.Location
	}
	if rcp.Sku != nil {
		objectMap["sku"] = rcp.Sku
	}
	if rcp.RegistryPropertiesCreateParameters != nil {
		objectMap["properties"] = rcp.RegistryPropertiesCreateParameters
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for RegistryCreateParameters struct.
func (rcp *RegistryCreateParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				rcp.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				rcp.Location = &location
			}
		case "sku":
			if v != nil {
				var sku Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				rcp.Sku = &sku
			}
		case "properties":
			if v != nil {
				var registryPropertiesCreateParameters RegistryPropertiesCreateParameters
				err = json.Unmarshal(*v, &registryPropertiesCreateParameters)
				if err != nil {
					return err
				}
				rcp.RegistryPropertiesCreateParameters = &registryPropertiesCreateParameters
			}
		}
	}

	return nil
}

// RegistryListCredentialsResult the response from the ListCredentials operation.
type RegistryListCredentialsResult struct {
	autorest.Response `json:"-"`
	// Username - The username for a container registry.
	Username *string `json:"username,omitempty"`
	// Passwords - The list of passwords for a container registry.
	Passwords *[]RegistryPassword `json:"passwords,omitempty"`
}

// RegistryListResult the result of a request to list container registries.
type RegistryListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of container registries. Since this list may be incomplete, the nextLink field should be used to request the next list of container registries.
	Value *[]Registry `json:"value,omitempty"`
	// NextLink - The URI that can be used to request the next list of container registries.
	NextLink *string `json:"nextLink,omitempty"`
}

// RegistryListResultIterator provides access to a complete listing of Registry values.
type RegistryListResultIterator struct {
	i    int
	page RegistryListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *RegistryListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistryListResultIterator.NextWithContext")
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
func (iter *RegistryListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter RegistryListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter RegistryListResultIterator) Response() RegistryListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter RegistryListResultIterator) Value() Registry {
	if !iter.page.NotDone() {
		return Registry{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the RegistryListResultIterator type.
func NewRegistryListResultIterator(page RegistryListResultPage) RegistryListResultIterator {
	return RegistryListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (rlr RegistryListResult) IsEmpty() bool {
	return rlr.Value == nil || len(*rlr.Value) == 0
}

// registryListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rlr RegistryListResult) registryListResultPreparer(ctx context.Context) (*http.Request, error) {
	if rlr.NextLink == nil || len(to.String(rlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rlr.NextLink)))
}

// RegistryListResultPage contains a page of Registry values.
type RegistryListResultPage struct {
	fn  func(context.Context, RegistryListResult) (RegistryListResult, error)
	rlr RegistryListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *RegistryListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistryListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.rlr)
	if err != nil {
		return err
	}
	page.rlr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *RegistryListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page RegistryListResultPage) NotDone() bool {
	return !page.rlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page RegistryListResultPage) Response() RegistryListResult {
	return page.rlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page RegistryListResultPage) Values() []Registry {
	if page.rlr.IsEmpty() {
		return nil
	}
	return *page.rlr.Value
}

// Creates a new instance of the RegistryListResultPage type.
func NewRegistryListResultPage(getNextPage func(context.Context, RegistryListResult) (RegistryListResult, error)) RegistryListResultPage {
	return RegistryListResultPage{fn: getNextPage}
}

// RegistryNameCheckRequest a request to check whether a container registry name is available.
type RegistryNameCheckRequest struct {
	// Name - The name of the container registry.
	Name *string `json:"name,omitempty"`
	// Type - The resource type of the container registry. This field must be set to 'Microsoft.ContainerRegistry/registries'.
	Type *string `json:"type,omitempty"`
}

// RegistryNameStatus the result of a request to check the availability of a container registry name.
type RegistryNameStatus struct {
	autorest.Response `json:"-"`
	// NameAvailable - The value that indicates whether the name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - If any, the reason that the name is not available.
	Reason *string `json:"reason,omitempty"`
	// Message - If any, the error message that provides more detail for the reason that the name is not available.
	Message *string `json:"message,omitempty"`
}

// RegistryPassword the login password for the container registry.
type RegistryPassword struct {
	// Name - The password name. Possible values include: 'Password', 'Password2'
	Name PasswordName `json:"name,omitempty"`
	// Value - The password value.
	Value *string `json:"value,omitempty"`
}

// RegistryProperties the properties of a container registry.
type RegistryProperties struct {
	// LoginServer - READ-ONLY; The URL that can be used to log into the container registry.
	LoginServer *string `json:"loginServer,omitempty"`
	// CreationDate - READ-ONLY; The creation date of the container registry in ISO8601 format.
	CreationDate *date.Time `json:"creationDate,omitempty"`
	// ProvisioningState - READ-ONLY; The provisioning state of the container registry at the time the operation was called. Possible values include: 'Creating', 'Succeeded'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// AdminUserEnabled - The value that indicates whether the admin user is enabled.
	AdminUserEnabled *bool `json:"adminUserEnabled,omitempty"`
	// StorageAccount - The properties of the storage account for the container registry.
	StorageAccount *StorageAccountProperties `json:"storageAccount,omitempty"`
}

// RegistryPropertiesCreateParameters the parameters for creating the properties of a container registry.
type RegistryPropertiesCreateParameters struct {
	// AdminUserEnabled - The value that indicates whether the admin user is enabled.
	AdminUserEnabled *bool `json:"adminUserEnabled,omitempty"`
	// StorageAccount - The parameters of a storage account for the container registry. If specified, the storage account must be in the same physical location as the container registry.
	StorageAccount *StorageAccountParameters `json:"storageAccount,omitempty"`
}

// RegistryPropertiesUpdateParameters the parameters for updating the properties of a container registry.
type RegistryPropertiesUpdateParameters struct {
	// AdminUserEnabled - The value that indicates whether the admin user is enabled.
	AdminUserEnabled *bool `json:"adminUserEnabled,omitempty"`
	// StorageAccount - The parameters of a storage account for the container registry. If specified, the storage account must be in the same physical location as the container registry.
	StorageAccount *StorageAccountParameters `json:"storageAccount,omitempty"`
}

// RegistryUpdateParameters the parameters for updating a container registry.
type RegistryUpdateParameters struct {
	// Tags - The tags for the container registry.
	Tags map[string]*string `json:"tags"`
	// RegistryPropertiesUpdateParameters - The properties that the container registry will be updated with.
	*RegistryPropertiesUpdateParameters `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for RegistryUpdateParameters.
func (rup RegistryUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rup.Tags != nil {
		objectMap["tags"] = rup.Tags
	}
	if rup.RegistryPropertiesUpdateParameters != nil {
		objectMap["properties"] = rup.RegistryPropertiesUpdateParameters
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for RegistryUpdateParameters struct.
func (rup *RegistryUpdateParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				rup.Tags = tags
			}
		case "properties":
			if v != nil {
				var registryPropertiesUpdateParameters RegistryPropertiesUpdateParameters
				err = json.Unmarshal(*v, &registryPropertiesUpdateParameters)
				if err != nil {
					return err
				}
				rup.RegistryPropertiesUpdateParameters = &registryPropertiesUpdateParameters
			}
		}
	}

	return nil
}

// Resource an Azure resource.
type Resource struct {
	// ID - READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
	// Location - The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`
	// Tags - The tags of the resource.
	Tags map[string]*string `json:"tags"`
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
	return json.Marshal(objectMap)
}

// Sku the SKU of a container registry.
type Sku struct {
	// Name - The SKU name of the container registry. Required for registry creation. Allowed value: Basic.
	Name *string `json:"name,omitempty"`
	// Tier - READ-ONLY; The SKU tier based on the SKU name. Possible values include: 'Basic'
	Tier SkuTier `json:"tier,omitempty"`
}

// StorageAccountParameters the parameters of a storage account for a container registry.
type StorageAccountParameters struct {
	// Name - The name of the storage account.
	Name *string `json:"name,omitempty"`
	// AccessKey - The access key to the storage account.
	AccessKey *string `json:"accessKey,omitempty"`
}

// StorageAccountProperties the properties of a storage account for a container registry.
type StorageAccountProperties struct {
	// Name - The name of the storage account.
	Name *string `json:"name,omitempty"`
}
