package policy

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
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/test-repo-arcturus/azure-sdk-for-go/services/resources/mgmt/2018-03-01/policy"

// Mode enumerates the values for mode.
type Mode string

const (
	// All ...
	All Mode = "All"
	// Indexed ...
	Indexed Mode = "Indexed"
	// NotSpecified ...
	NotSpecified Mode = "NotSpecified"
)

// PossibleModeValues returns an array of possible values for the Mode const type.
func PossibleModeValues() []Mode {
	return []Mode{All, Indexed, NotSpecified}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeBuiltIn ...
	TypeBuiltIn Type = "BuiltIn"
	// TypeCustom ...
	TypeCustom Type = "Custom"
	// TypeNotSpecified ...
	TypeNotSpecified Type = "NotSpecified"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeBuiltIn, TypeCustom, TypeNotSpecified}
}

// Assignment the policy assignment.
type Assignment struct {
	autorest.Response `json:"-"`
	// AssignmentProperties - Properties for the policy assignment.
	*AssignmentProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; The ID of the policy assignment.
	ID *string `json:"id,omitempty"`
	// Type - READ-ONLY; The type of the policy assignment.
	Type *string `json:"type,omitempty"`
	// Name - READ-ONLY; The name of the policy assignment.
	Name *string `json:"name,omitempty"`
	// Sku - The policy sku. This property is optional, obsolete, and will be ignored.
	Sku *Sku `json:"sku,omitempty"`
}

// MarshalJSON is the custom marshaler for Assignment.
func (a Assignment) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.AssignmentProperties != nil {
		objectMap["properties"] = a.AssignmentProperties
	}
	if a.Sku != nil {
		objectMap["sku"] = a.Sku
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Assignment struct.
func (a *Assignment) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var assignmentProperties AssignmentProperties
				err = json.Unmarshal(*v, &assignmentProperties)
				if err != nil {
					return err
				}
				a.AssignmentProperties = &assignmentProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				a.ID = &ID
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				a.Type = &typeVar
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				a.Name = &name
			}
		case "sku":
			if v != nil {
				var sku Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				a.Sku = &sku
			}
		}
	}

	return nil
}

// AssignmentListResult list of policy assignments.
type AssignmentListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of policy assignments.
	Value *[]Assignment `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// AssignmentListResultIterator provides access to a complete listing of Assignment values.
type AssignmentListResultIterator struct {
	i    int
	page AssignmentListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *AssignmentListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentListResultIterator.NextWithContext")
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
func (iter *AssignmentListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter AssignmentListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter AssignmentListResultIterator) Response() AssignmentListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter AssignmentListResultIterator) Value() Assignment {
	if !iter.page.NotDone() {
		return Assignment{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the AssignmentListResultIterator type.
func NewAssignmentListResultIterator(page AssignmentListResultPage) AssignmentListResultIterator {
	return AssignmentListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (alr AssignmentListResult) IsEmpty() bool {
	return alr.Value == nil || len(*alr.Value) == 0
}

// assignmentListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (alr AssignmentListResult) assignmentListResultPreparer(ctx context.Context) (*http.Request, error) {
	if alr.NextLink == nil || len(to.String(alr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(alr.NextLink)))
}

// AssignmentListResultPage contains a page of Assignment values.
type AssignmentListResultPage struct {
	fn  func(context.Context, AssignmentListResult) (AssignmentListResult, error)
	alr AssignmentListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AssignmentListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.alr)
	if err != nil {
		return err
	}
	page.alr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *AssignmentListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page AssignmentListResultPage) NotDone() bool {
	return !page.alr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page AssignmentListResultPage) Response() AssignmentListResult {
	return page.alr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page AssignmentListResultPage) Values() []Assignment {
	if page.alr.IsEmpty() {
		return nil
	}
	return *page.alr.Value
}

// Creates a new instance of the AssignmentListResultPage type.
func NewAssignmentListResultPage(getNextPage func(context.Context, AssignmentListResult) (AssignmentListResult, error)) AssignmentListResultPage {
	return AssignmentListResultPage{fn: getNextPage}
}

// AssignmentProperties the policy assignment properties.
type AssignmentProperties struct {
	// DisplayName - The display name of the policy assignment.
	DisplayName *string `json:"displayName,omitempty"`
	// PolicyDefinitionID - The ID of the policy definition or policy set definition being assigned.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty"`
	// Scope - The scope for the policy assignment.
	Scope *string `json:"scope,omitempty"`
	// NotScopes - The policy's excluded scopes.
	NotScopes *[]string `json:"notScopes,omitempty"`
	// Parameters - Required if a parameter is used in policy rule.
	Parameters interface{} `json:"parameters,omitempty"`
	// Description - This message will be part of response in case of policy violation.
	Description *string `json:"description,omitempty"`
	// Metadata - The policy assignment metadata.
	Metadata interface{} `json:"metadata,omitempty"`
}

// Definition the policy definition.
type Definition struct {
	autorest.Response `json:"-"`
	// DefinitionProperties - The policy definition properties.
	*DefinitionProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; The ID of the policy definition.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the policy definition.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource (Microsoft.Authorization/policyDefinitions).
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Definition.
func (d Definition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if d.DefinitionProperties != nil {
		objectMap["properties"] = d.DefinitionProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Definition struct.
func (d *Definition) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var definitionProperties DefinitionProperties
				err = json.Unmarshal(*v, &definitionProperties)
				if err != nil {
					return err
				}
				d.DefinitionProperties = &definitionProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				d.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				d.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				d.Type = &typeVar
			}
		}
	}

	return nil
}

// DefinitionListResult list of policy definitions.
type DefinitionListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of policy definitions.
	Value *[]Definition `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// DefinitionListResultIterator provides access to a complete listing of Definition values.
type DefinitionListResultIterator struct {
	i    int
	page DefinitionListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *DefinitionListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionListResultIterator.NextWithContext")
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
func (iter *DefinitionListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter DefinitionListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter DefinitionListResultIterator) Response() DefinitionListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter DefinitionListResultIterator) Value() Definition {
	if !iter.page.NotDone() {
		return Definition{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the DefinitionListResultIterator type.
func NewDefinitionListResultIterator(page DefinitionListResultPage) DefinitionListResultIterator {
	return DefinitionListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (dlr DefinitionListResult) IsEmpty() bool {
	return dlr.Value == nil || len(*dlr.Value) == 0
}

// definitionListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (dlr DefinitionListResult) definitionListResultPreparer(ctx context.Context) (*http.Request, error) {
	if dlr.NextLink == nil || len(to.String(dlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(dlr.NextLink)))
}

// DefinitionListResultPage contains a page of Definition values.
type DefinitionListResultPage struct {
	fn  func(context.Context, DefinitionListResult) (DefinitionListResult, error)
	dlr DefinitionListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *DefinitionListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.dlr)
	if err != nil {
		return err
	}
	page.dlr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *DefinitionListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page DefinitionListResultPage) NotDone() bool {
	return !page.dlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page DefinitionListResultPage) Response() DefinitionListResult {
	return page.dlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page DefinitionListResultPage) Values() []Definition {
	if page.dlr.IsEmpty() {
		return nil
	}
	return *page.dlr.Value
}

// Creates a new instance of the DefinitionListResultPage type.
func NewDefinitionListResultPage(getNextPage func(context.Context, DefinitionListResult) (DefinitionListResult, error)) DefinitionListResultPage {
	return DefinitionListResultPage{fn: getNextPage}
}

// DefinitionProperties the policy definition properties.
type DefinitionProperties struct {
	// PolicyType - The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom. Possible values include: 'TypeNotSpecified', 'TypeBuiltIn', 'TypeCustom'
	PolicyType Type `json:"policyType,omitempty"`
	// Mode - The policy definition mode. Possible values are NotSpecified, Indexed, and All. Possible values include: 'NotSpecified', 'Indexed', 'All'
	Mode Mode `json:"mode,omitempty"`
	// DisplayName - The display name of the policy definition.
	DisplayName *string `json:"displayName,omitempty"`
	// Description - The policy definition description.
	Description *string `json:"description,omitempty"`
	// PolicyRule - The policy rule.
	PolicyRule interface{} `json:"policyRule,omitempty"`
	// Metadata - The policy definition metadata.
	Metadata interface{} `json:"metadata,omitempty"`
	// Parameters - Required if a parameter is used in policy rule.
	Parameters interface{} `json:"parameters,omitempty"`
}

// DefinitionReference the policy definition reference.
type DefinitionReference struct {
	// PolicyDefinitionID - The ID of the policy definition or policy set definition.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty"`
	// Parameters - Required if a parameter is used in policy rule.
	Parameters interface{} `json:"parameters,omitempty"`
}

// ErrorResponse error response indicates Azure Resource Manager is not able to process the incoming
// request. The reason is provided in the error message.
type ErrorResponse struct {
	// HTTPStatus - Http status code.
	HTTPStatus *string `json:"httpStatus,omitempty"`
	// ErrorCode - Error code.
	ErrorCode *string `json:"errorCode,omitempty"`
	// ErrorMessage - Error message indicating why the operation failed.
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// SetDefinition the policy set definition.
type SetDefinition struct {
	autorest.Response `json:"-"`
	// SetDefinitionProperties - The policy definition properties.
	*SetDefinitionProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; The ID of the policy set definition.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the policy set definition.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource (Microsoft.Authorization/policySetDefinitions).
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for SetDefinition.
func (sd SetDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sd.SetDefinitionProperties != nil {
		objectMap["properties"] = sd.SetDefinitionProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for SetDefinition struct.
func (sd *SetDefinition) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var setDefinitionProperties SetDefinitionProperties
				err = json.Unmarshal(*v, &setDefinitionProperties)
				if err != nil {
					return err
				}
				sd.SetDefinitionProperties = &setDefinitionProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				sd.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				sd.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				sd.Type = &typeVar
			}
		}
	}

	return nil
}

// SetDefinitionListResult list of policy set definitions.
type SetDefinitionListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of policy set definitions.
	Value *[]SetDefinition `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// SetDefinitionListResultIterator provides access to a complete listing of SetDefinition values.
type SetDefinitionListResultIterator struct {
	i    int
	page SetDefinitionListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *SetDefinitionListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SetDefinitionListResultIterator.NextWithContext")
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
func (iter *SetDefinitionListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter SetDefinitionListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter SetDefinitionListResultIterator) Response() SetDefinitionListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter SetDefinitionListResultIterator) Value() SetDefinition {
	if !iter.page.NotDone() {
		return SetDefinition{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the SetDefinitionListResultIterator type.
func NewSetDefinitionListResultIterator(page SetDefinitionListResultPage) SetDefinitionListResultIterator {
	return SetDefinitionListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (sdlr SetDefinitionListResult) IsEmpty() bool {
	return sdlr.Value == nil || len(*sdlr.Value) == 0
}

// setDefinitionListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (sdlr SetDefinitionListResult) setDefinitionListResultPreparer(ctx context.Context) (*http.Request, error) {
	if sdlr.NextLink == nil || len(to.String(sdlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(sdlr.NextLink)))
}

// SetDefinitionListResultPage contains a page of SetDefinition values.
type SetDefinitionListResultPage struct {
	fn   func(context.Context, SetDefinitionListResult) (SetDefinitionListResult, error)
	sdlr SetDefinitionListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *SetDefinitionListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SetDefinitionListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.sdlr)
	if err != nil {
		return err
	}
	page.sdlr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *SetDefinitionListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page SetDefinitionListResultPage) NotDone() bool {
	return !page.sdlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page SetDefinitionListResultPage) Response() SetDefinitionListResult {
	return page.sdlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page SetDefinitionListResultPage) Values() []SetDefinition {
	if page.sdlr.IsEmpty() {
		return nil
	}
	return *page.sdlr.Value
}

// Creates a new instance of the SetDefinitionListResultPage type.
func NewSetDefinitionListResultPage(getNextPage func(context.Context, SetDefinitionListResult) (SetDefinitionListResult, error)) SetDefinitionListResultPage {
	return SetDefinitionListResultPage{fn: getNextPage}
}

// SetDefinitionProperties the policy set definition properties.
type SetDefinitionProperties struct {
	// PolicyType - The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom. Possible values include: 'TypeNotSpecified', 'TypeBuiltIn', 'TypeCustom'
	PolicyType Type `json:"policyType,omitempty"`
	// DisplayName - The display name of the policy set definition.
	DisplayName *string `json:"displayName,omitempty"`
	// Description - The policy set definition description.
	Description *string `json:"description,omitempty"`
	// Metadata - The policy set definition metadata.
	Metadata interface{} `json:"metadata,omitempty"`
	// Parameters - The policy set definition parameters that can be used in policy definition references.
	Parameters interface{} `json:"parameters,omitempty"`
	// PolicyDefinitions - An array of policy definition references.
	PolicyDefinitions *[]DefinitionReference `json:"policyDefinitions,omitempty"`
}

// Sku the policy sku. This property is optional, obsolete, and will be ignored.
type Sku struct {
	// Name - The name of the policy sku. Possible values are A0 and A1.
	Name *string `json:"name,omitempty"`
	// Tier - The policy sku tier. Possible values are Free and Standard.
	Tier *string `json:"tier,omitempty"`
}
