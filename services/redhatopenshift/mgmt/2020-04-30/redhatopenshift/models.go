package redhatopenshift

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
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/test-repo-arcturus/azure-sdk-for-go/services/redhatopenshift/mgmt/2020-04-30/redhatopenshift"

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// AdminUpdating ...
	AdminUpdating ProvisioningState = "AdminUpdating"
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{AdminUpdating, Creating, Deleting, Failed, Succeeded, Updating}
}

// Visibility enumerates the values for visibility.
type Visibility string

const (
	// Private ...
	Private Visibility = "Private"
	// Public ...
	Public Visibility = "Public"
)

// PossibleVisibilityValues returns an array of possible values for the Visibility const type.
func PossibleVisibilityValues() []Visibility {
	return []Visibility{Private, Public}
}

// Visibility1 enumerates the values for visibility 1.
type Visibility1 string

const (
	// Visibility1Private ...
	Visibility1Private Visibility1 = "Private"
	// Visibility1Public ...
	Visibility1Public Visibility1 = "Public"
)

// PossibleVisibility1Values returns an array of possible values for the Visibility1 const type.
func PossibleVisibility1Values() []Visibility1 {
	return []Visibility1{Visibility1Private, Visibility1Public}
}

// VMSize enumerates the values for vm size.
type VMSize string

const (
	// StandardD2sV3 ...
	StandardD2sV3 VMSize = "Standard_D2s_v3"
	// StandardD4sV3 ...
	StandardD4sV3 VMSize = "Standard_D4s_v3"
	// StandardD8sV3 ...
	StandardD8sV3 VMSize = "Standard_D8s_v3"
)

// PossibleVMSizeValues returns an array of possible values for the VMSize const type.
func PossibleVMSizeValues() []VMSize {
	return []VMSize{StandardD2sV3, StandardD4sV3, StandardD8sV3}
}

// VMSize1 enumerates the values for vm size 1.
type VMSize1 string

const (
	// VMSize1StandardD2sV3 ...
	VMSize1StandardD2sV3 VMSize1 = "Standard_D2s_v3"
	// VMSize1StandardD4sV3 ...
	VMSize1StandardD4sV3 VMSize1 = "Standard_D4s_v3"
	// VMSize1StandardD8sV3 ...
	VMSize1StandardD8sV3 VMSize1 = "Standard_D8s_v3"
)

// PossibleVMSize1Values returns an array of possible values for the VMSize1 const type.
func PossibleVMSize1Values() []VMSize1 {
	return []VMSize1{VMSize1StandardD2sV3, VMSize1StandardD4sV3, VMSize1StandardD8sV3}
}

// APIServerProfile aPIServerProfile represents an API server profile.
type APIServerProfile struct {
	// Visibility - API server visibility (immutable). Possible values include: 'Private', 'Public'
	Visibility Visibility `json:"visibility,omitempty"`
	// URL - The URL to access the cluster API server (immutable).
	URL *string `json:"url,omitempty"`
	// IP - The IP of the cluster API server (immutable).
	IP *string `json:"ip,omitempty"`
}

// AzureEntityResource the resource model definition for a Azure Resource Manager resource with an etag.
type AzureEntityResource struct {
	// Etag - READ-ONLY; Resource Etag.
	Etag *string `json:"etag,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// CloudError cloudError represents a cloud error.
type CloudError struct {
	// Error - An error response from the service.
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody cloudErrorBody represents the body of a cloud error.
type CloudErrorBody struct {
	// Code - An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string `json:"code,omitempty"`
	// Message - A message describing the error, intended to be suitable for display in a user interface.
	Message *string `json:"message,omitempty"`
	// Target - The target of the particular error. For example, the name of the property in error.
	Target *string `json:"target,omitempty"`
	// Details - A list of additional details about the error.
	Details *[]CloudErrorBody `json:"details,omitempty"`
}

// ClusterProfile clusterProfile represents a cluster profile.
type ClusterProfile struct {
	// PullSecret - The pull secret for the cluster (immutable).
	PullSecret *string `json:"pullSecret,omitempty"`
	// Domain - The domain for the cluster (immutable).
	Domain *string `json:"domain,omitempty"`
	// Version - The version of the cluster (immutable).
	Version *string `json:"version,omitempty"`
	// ResourceGroupID - The ID of the cluster resource group (immutable).
	ResourceGroupID *string `json:"resourceGroupId,omitempty"`
}

// ConsoleProfile consoleProfile represents a console profile.
type ConsoleProfile struct {
	// URL - The URL to access the cluster console (immutable).
	URL *string `json:"url,omitempty"`
}

// Display display represents the display details of an operation.
type Display struct {
	// Provider - Friendly name of the resource provider.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource type on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: read, write, delete, listKeys/action, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Friendly name of the operation.
	Description *string `json:"description,omitempty"`
}

// IngressProfile ingressProfile represents an ingress profile.
type IngressProfile struct {
	// Name - The ingress profile name.  Must be "default" (immutable).
	Name *string `json:"name,omitempty"`
	// Visibility - Ingress visibility (immutable). Possible values include: 'Visibility1Private', 'Visibility1Public'
	Visibility Visibility1 `json:"visibility,omitempty"`
	// IP - The IP of the ingress (immutable).
	IP *string `json:"ip,omitempty"`
}

// MasterProfile masterProfile represents a master profile.
type MasterProfile struct {
	// VMSize - The size of the master VMs (immutable). Possible values include: 'StandardD2sV3', 'StandardD4sV3', 'StandardD8sV3'
	VMSize VMSize `json:"vmSize,omitempty"`
	// SubnetID - The Azure resource ID of the master subnet (immutable).
	SubnetID *string `json:"subnetId,omitempty"`
}

// NetworkProfile networkProfile represents a network profile.
type NetworkProfile struct {
	// PodCidr - The CIDR used for OpenShift/Kubernetes Pods (immutable).
	PodCidr *string `json:"podCidr,omitempty"`
	// ServiceCidr - The CIDR used for OpenShift/Kubernetes Services (immutable).
	ServiceCidr *string `json:"serviceCidr,omitempty"`
}

// OpenShiftCluster openShiftCluster represents an Azure Red Hat OpenShift cluster.
type OpenShiftCluster struct {
	autorest.Response `json:"-"`
	// OpenShiftClusterProperties - The cluster properties.
	*OpenShiftClusterProperties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for OpenShiftCluster.
func (osc OpenShiftCluster) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if osc.OpenShiftClusterProperties != nil {
		objectMap["properties"] = osc.OpenShiftClusterProperties
	}
	if osc.Tags != nil {
		objectMap["tags"] = osc.Tags
	}
	if osc.Location != nil {
		objectMap["location"] = osc.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for OpenShiftCluster struct.
func (osc *OpenShiftCluster) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var openShiftClusterProperties OpenShiftClusterProperties
				err = json.Unmarshal(*v, &openShiftClusterProperties)
				if err != nil {
					return err
				}
				osc.OpenShiftClusterProperties = &openShiftClusterProperties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				osc.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				osc.Location = &location
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				osc.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				osc.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				osc.Type = &typeVar
			}
		}
	}

	return nil
}

// OpenShiftClusterCredentials openShiftClusterCredentials represents an OpenShift cluster's credentials
type OpenShiftClusterCredentials struct {
	autorest.Response `json:"-"`
	// KubeadminUsername - The username for the kubeadmin user
	KubeadminUsername *string `json:"kubeadminUsername,omitempty"`
	// KubeadminPassword - The password for the kubeadmin user
	KubeadminPassword *string `json:"kubeadminPassword,omitempty"`
}

// OpenShiftClusterList openShiftClusterList represents a list of OpenShift clusters.
type OpenShiftClusterList struct {
	autorest.Response `json:"-"`
	// Value - The list of OpenShift clusters.
	Value *[]OpenShiftCluster `json:"value,omitempty"`
	// NextLink - The link used to get the next page of operations.
	NextLink *string `json:"nextLink,omitempty"`
}

// OpenShiftClusterListIterator provides access to a complete listing of OpenShiftCluster values.
type OpenShiftClusterListIterator struct {
	i    int
	page OpenShiftClusterListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OpenShiftClusterListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OpenShiftClusterListIterator.NextWithContext")
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
func (iter *OpenShiftClusterListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OpenShiftClusterListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OpenShiftClusterListIterator) Response() OpenShiftClusterList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OpenShiftClusterListIterator) Value() OpenShiftCluster {
	if !iter.page.NotDone() {
		return OpenShiftCluster{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OpenShiftClusterListIterator type.
func NewOpenShiftClusterListIterator(page OpenShiftClusterListPage) OpenShiftClusterListIterator {
	return OpenShiftClusterListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (oscl OpenShiftClusterList) IsEmpty() bool {
	return oscl.Value == nil || len(*oscl.Value) == 0
}

// openShiftClusterListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (oscl OpenShiftClusterList) openShiftClusterListPreparer(ctx context.Context) (*http.Request, error) {
	if oscl.NextLink == nil || len(to.String(oscl.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(oscl.NextLink)))
}

// OpenShiftClusterListPage contains a page of OpenShiftCluster values.
type OpenShiftClusterListPage struct {
	fn   func(context.Context, OpenShiftClusterList) (OpenShiftClusterList, error)
	oscl OpenShiftClusterList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OpenShiftClusterListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OpenShiftClusterListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.oscl)
	if err != nil {
		return err
	}
	page.oscl = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OpenShiftClusterListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OpenShiftClusterListPage) NotDone() bool {
	return !page.oscl.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OpenShiftClusterListPage) Response() OpenShiftClusterList {
	return page.oscl
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OpenShiftClusterListPage) Values() []OpenShiftCluster {
	if page.oscl.IsEmpty() {
		return nil
	}
	return *page.oscl.Value
}

// Creates a new instance of the OpenShiftClusterListPage type.
func NewOpenShiftClusterListPage(getNextPage func(context.Context, OpenShiftClusterList) (OpenShiftClusterList, error)) OpenShiftClusterListPage {
	return OpenShiftClusterListPage{fn: getNextPage}
}

// OpenShiftClusterProperties openShiftClusterProperties represents an OpenShift cluster's properties.
type OpenShiftClusterProperties struct {
	// ProvisioningState - The cluster provisioning state (immutable). Possible values include: 'AdminUpdating', 'Creating', 'Deleting', 'Failed', 'Succeeded', 'Updating'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// ClusterProfile - The cluster profile.
	ClusterProfile *ClusterProfile `json:"clusterProfile,omitempty"`
	// ConsoleProfile - The console profile.
	ConsoleProfile *ConsoleProfile `json:"consoleProfile,omitempty"`
	// ServicePrincipalProfile - The cluster service principal profile.
	ServicePrincipalProfile *ServicePrincipalProfile `json:"servicePrincipalProfile,omitempty"`
	// NetworkProfile - The cluster network profile.
	NetworkProfile *NetworkProfile `json:"networkProfile,omitempty"`
	// MasterProfile - The cluster master profile.
	MasterProfile *MasterProfile `json:"masterProfile,omitempty"`
	// WorkerProfiles - The cluster worker profiles.
	WorkerProfiles *[]WorkerProfile `json:"workerProfiles,omitempty"`
	// ApiserverProfile - The cluster API server profile.
	ApiserverProfile *APIServerProfile `json:"apiserverProfile,omitempty"`
	// IngressProfiles - The cluster ingress profiles.
	IngressProfiles *[]IngressProfile `json:"ingressProfiles,omitempty"`
}

// OpenShiftClustersCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type OpenShiftClustersCreateOrUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *OpenShiftClustersCreateOrUpdateFuture) Result(client OpenShiftClustersClient) (osc OpenShiftCluster, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redhatopenshift.OpenShiftClustersCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("redhatopenshift.OpenShiftClustersCreateOrUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if osc.Response.Response, err = future.GetResult(sender); err == nil && osc.Response.Response.StatusCode != http.StatusNoContent {
		osc, err = client.CreateOrUpdateResponder(osc.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "redhatopenshift.OpenShiftClustersCreateOrUpdateFuture", "Result", osc.Response.Response, "Failure responding to request")
		}
	}
	return
}

// OpenShiftClustersDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type OpenShiftClustersDeleteFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *OpenShiftClustersDeleteFuture) Result(client OpenShiftClustersClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redhatopenshift.OpenShiftClustersDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("redhatopenshift.OpenShiftClustersDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// OpenShiftClustersUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type OpenShiftClustersUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *OpenShiftClustersUpdateFuture) Result(client OpenShiftClustersClient) (osc OpenShiftCluster, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redhatopenshift.OpenShiftClustersUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("redhatopenshift.OpenShiftClustersUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if osc.Response.Response, err = future.GetResult(sender); err == nil && osc.Response.Response.StatusCode != http.StatusNoContent {
		osc, err = client.UpdateResponder(osc.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "redhatopenshift.OpenShiftClustersUpdateFuture", "Result", osc.Response.Response, "Failure responding to request")
		}
	}
	return
}

// OpenShiftClusterUpdate openShiftCluster represents an Azure Red Hat OpenShift cluster.
type OpenShiftClusterUpdate struct {
	// Tags - The resource tags.
	Tags map[string]*string `json:"tags"`
	// OpenShiftClusterProperties - The cluster properties.
	*OpenShiftClusterProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for OpenShiftClusterUpdate.
func (oscu OpenShiftClusterUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if oscu.Tags != nil {
		objectMap["tags"] = oscu.Tags
	}
	if oscu.OpenShiftClusterProperties != nil {
		objectMap["properties"] = oscu.OpenShiftClusterProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for OpenShiftClusterUpdate struct.
func (oscu *OpenShiftClusterUpdate) UnmarshalJSON(body []byte) error {
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
				oscu.Tags = tags
			}
		case "properties":
			if v != nil {
				var openShiftClusterProperties OpenShiftClusterProperties
				err = json.Unmarshal(*v, &openShiftClusterProperties)
				if err != nil {
					return err
				}
				oscu.OpenShiftClusterProperties = &openShiftClusterProperties
			}
		}
	}

	return nil
}

// Operation operation represents an RP operation.
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The object that describes the operation.
	Display *Display `json:"display,omitempty"`
	// Origin - Sources of requests to this operation.  Comma separated list with valid values user or system, e.g. "user,system".
	Origin *string `json:"origin,omitempty"`
}

// OperationList operationList represents an RP operation list.
type OperationList struct {
	autorest.Response `json:"-"`
	// Value - List of operations supported by the resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - The link used to get the next page of operations.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListIterator provides access to a complete listing of Operation values.
type OperationListIterator struct {
	i    int
	page OperationListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListIterator.NextWithContext")
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
func (iter *OperationListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListIterator) Response() OperationList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListIterator type.
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return OperationListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ol OperationList) IsEmpty() bool {
	return ol.Value == nil || len(*ol.Value) == 0
}

// operationListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ol OperationList) operationListPreparer(ctx context.Context) (*http.Request, error) {
	if ol.NextLink == nil || len(to.String(ol.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ol.NextLink)))
}

// OperationListPage contains a page of Operation values.
type OperationListPage struct {
	fn func(context.Context, OperationList) (OperationList, error)
	ol OperationList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.ol)
	if err != nil {
		return err
	}
	page.ol = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListPage) NotDone() bool {
	return !page.ol.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListPage) Response() OperationList {
	return page.ol
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListPage) Values() []Operation {
	if page.ol.IsEmpty() {
		return nil
	}
	return *page.ol.Value
}

// Creates a new instance of the OperationListPage type.
func NewOperationListPage(getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return OperationListPage{fn: getNextPage}
}

// ProxyResource the resource model definition for a ARM proxy resource. It will have everything other than
// required location and tags
type ProxyResource struct {
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// Resource ...
type Resource struct {
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// ServicePrincipalProfile servicePrincipalProfile represents a service principal profile.
type ServicePrincipalProfile struct {
	// ClientID - The client ID used for the cluster (immutable).
	ClientID *string `json:"clientId,omitempty"`
	// ClientSecret - The client secret used for the cluster (immutable).
	ClientSecret *string `json:"clientSecret,omitempty"`
}

// TrackedResource the resource model definition for a ARM tracked top level resource
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	return json.Marshal(objectMap)
}

// WorkerProfile workerProfile represents a worker profile.
type WorkerProfile struct {
	// Name - The worker profile name.  Must be "worker" (immutable).
	Name *string `json:"name,omitempty"`
	// VMSize - The size of the worker VMs (immutable). Possible values include: 'VMSize1StandardD2sV3', 'VMSize1StandardD4sV3', 'VMSize1StandardD8sV3'
	VMSize VMSize1 `json:"vmSize,omitempty"`
	// DiskSizeGB - The disk size of the worker VMs.  Must be 128 or greater (immutable).
	DiskSizeGB *int32 `json:"diskSizeGB,omitempty"`
	// SubnetID - The Azure resource ID of the worker subnet (immutable).
	SubnetID *string `json:"subnetId,omitempty"`
	// Count - The number of worker VMs.  Must be between 3 and 20 (immutable).
	Count *int32 `json:"count,omitempty"`
}
