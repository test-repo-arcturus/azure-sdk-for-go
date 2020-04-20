package redis

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
const fqdn = "github.com/test-repo-arcturus/azure-sdk-for-go/services/redis/mgmt/2015-08-01/redis"

// KeyType enumerates the values for key type.
type KeyType string

const (
	// Primary ...
	Primary KeyType = "Primary"
	// Secondary ...
	Secondary KeyType = "Secondary"
)

// PossibleKeyTypeValues returns an array of possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{Primary, Secondary}
}

// RebootType enumerates the values for reboot type.
type RebootType string

const (
	// AllNodes ...
	AllNodes RebootType = "AllNodes"
	// PrimaryNode ...
	PrimaryNode RebootType = "PrimaryNode"
	// SecondaryNode ...
	SecondaryNode RebootType = "SecondaryNode"
)

// PossibleRebootTypeValues returns an array of possible values for the RebootType const type.
func PossibleRebootTypeValues() []RebootType {
	return []RebootType{AllNodes, PrimaryNode, SecondaryNode}
}

// SkuFamily enumerates the values for sku family.
type SkuFamily string

const (
	// C ...
	C SkuFamily = "C"
	// P ...
	P SkuFamily = "P"
)

// PossibleSkuFamilyValues returns an array of possible values for the SkuFamily const type.
func PossibleSkuFamilyValues() []SkuFamily {
	return []SkuFamily{C, P}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Basic ...
	Basic SkuName = "Basic"
	// Premium ...
	Premium SkuName = "Premium"
	// Standard ...
	Standard SkuName = "Standard"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{Basic, Premium, Standard}
}

// AccessKeys redis cache access keys.
type AccessKeys struct {
	// PrimaryKey - The current primary key that clients can use to authenticate with Redis cache.
	PrimaryKey *string `json:"primaryKey,omitempty"`
	// SecondaryKey - The current secondary key that clients can use to authenticate with Redis cache.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
}

// CreateOrUpdateParameters parameters supplied to the CreateOrUpdate Redis operation.
type CreateOrUpdateParameters struct {
	// Properties - Redis cache properties.
	*Properties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for CreateOrUpdateParameters.
func (coup CreateOrUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if coup.Properties != nil {
		objectMap["properties"] = coup.Properties
	}
	if coup.Location != nil {
		objectMap["location"] = coup.Location
	}
	if coup.Tags != nil {
		objectMap["tags"] = coup.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for CreateOrUpdateParameters struct.
func (coup *CreateOrUpdateParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var properties Properties
				err = json.Unmarshal(*v, &properties)
				if err != nil {
					return err
				}
				coup.Properties = &properties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				coup.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				coup.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				coup.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				coup.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				coup.Tags = tags
			}
		}
	}

	return nil
}

// ListKeysResult the response of Redis list keys operation.
type ListKeysResult struct {
	autorest.Response `json:"-"`
	// PrimaryKey - The current primary key that clients can use to authenticate with Redis cache.
	PrimaryKey *string `json:"primaryKey,omitempty"`
	// SecondaryKey - The current secondary key that clients can use to authenticate with Redis cache.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
}

// ListResult the response of list Redis operation.
type ListResult struct {
	autorest.Response `json:"-"`
	// Value - Results of the list operation.
	Value *[]ResourceType `json:"value,omitempty"`
	// NextLink - Link for next set of locations.
	NextLink *string `json:"nextLink,omitempty"`
}

// ListResultIterator provides access to a complete listing of ResourceType values.
type ListResultIterator struct {
	i    int
	page ListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListResultIterator.NextWithContext")
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
func (iter *ListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ListResultIterator) Response() ListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ListResultIterator) Value() ResourceType {
	if !iter.page.NotDone() {
		return ResourceType{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ListResultIterator type.
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return ListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (lr ListResult) IsEmpty() bool {
	return lr.Value == nil || len(*lr.Value) == 0
}

// listResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (lr ListResult) listResultPreparer(ctx context.Context) (*http.Request, error) {
	if lr.NextLink == nil || len(to.String(lr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(lr.NextLink)))
}

// ListResultPage contains a page of ResourceType values.
type ListResultPage struct {
	fn func(context.Context, ListResult) (ListResult, error)
	lr ListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.lr)
	if err != nil {
		return err
	}
	page.lr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ListResultPage) NotDone() bool {
	return !page.lr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ListResultPage) Response() ListResult {
	return page.lr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ListResultPage) Values() []ResourceType {
	if page.lr.IsEmpty() {
		return nil
	}
	return *page.lr.Value
}

// Creates a new instance of the ListResultPage type.
func NewListResultPage(getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return ListResultPage{fn: getNextPage}
}

// Properties parameters supplied to CreateOrUpdate Redis operation.
type Properties struct {
	// RedisVersion - RedisVersion parameter has been deprecated. As such, it is no longer necessary to provide this parameter and any value specified is ignored.
	RedisVersion *string `json:"redisVersion,omitempty"`
	// Sku - What SKU of Redis cache to deploy.
	Sku *Sku `json:"sku,omitempty"`
	// RedisConfiguration - All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
	RedisConfiguration map[string]*string `json:"redisConfiguration"`
	// EnableNonSslPort - If the value is true, then the non-SLL Redis server port (6379) will be enabled.
	EnableNonSslPort *bool `json:"enableNonSslPort,omitempty"`
	// TenantSettings - tenantSettings
	TenantSettings map[string]*string `json:"tenantSettings"`
	// ShardCount - The number of shards to be created on a Premium Cluster Cache.
	ShardCount *int32 `json:"shardCount,omitempty"`
	// VirtualNetwork - The exact ARM resource ID of the virtual network to deploy the Redis cache in. Example format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.ClassicNetwork/VirtualNetworks/vnet1
	VirtualNetwork *string `json:"virtualNetwork,omitempty"`
	// Subnet - Required when deploying a Redis cache inside an existing Azure Virtual Network.
	Subnet *string `json:"subnet,omitempty"`
	// StaticIP - Required when deploying a Redis cache inside an existing Azure Virtual Network.
	StaticIP *string `json:"staticIP,omitempty"`
}

// MarshalJSON is the custom marshaler for Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if p.RedisVersion != nil {
		objectMap["redisVersion"] = p.RedisVersion
	}
	if p.Sku != nil {
		objectMap["sku"] = p.Sku
	}
	if p.RedisConfiguration != nil {
		objectMap["redisConfiguration"] = p.RedisConfiguration
	}
	if p.EnableNonSslPort != nil {
		objectMap["enableNonSslPort"] = p.EnableNonSslPort
	}
	if p.TenantSettings != nil {
		objectMap["tenantSettings"] = p.TenantSettings
	}
	if p.ShardCount != nil {
		objectMap["shardCount"] = p.ShardCount
	}
	if p.VirtualNetwork != nil {
		objectMap["virtualNetwork"] = p.VirtualNetwork
	}
	if p.Subnet != nil {
		objectMap["subnet"] = p.Subnet
	}
	if p.StaticIP != nil {
		objectMap["staticIP"] = p.StaticIP
	}
	return json.Marshal(objectMap)
}

// ReadableProperties parameters describing a Redis instance
type ReadableProperties struct {
	// ProvisioningState - Redis instance provisioning status.
	ProvisioningState *string `json:"provisioningState,omitempty"`
	// HostName - Redis host name.
	HostName *string `json:"hostName,omitempty"`
	// Port - Redis non-SSL port.
	Port *int32 `json:"port,omitempty"`
	// SslPort - Redis SSL port.
	SslPort *int32 `json:"sslPort,omitempty"`
	// RedisVersion - RedisVersion parameter has been deprecated. As such, it is no longer necessary to provide this parameter and any value specified is ignored.
	RedisVersion *string `json:"redisVersion,omitempty"`
	// Sku - What SKU of Redis cache to deploy.
	Sku *Sku `json:"sku,omitempty"`
	// RedisConfiguration - All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
	RedisConfiguration map[string]*string `json:"redisConfiguration"`
	// EnableNonSslPort - If the value is true, then the non-SLL Redis server port (6379) will be enabled.
	EnableNonSslPort *bool `json:"enableNonSslPort,omitempty"`
	// TenantSettings - tenantSettings
	TenantSettings map[string]*string `json:"tenantSettings"`
	// ShardCount - The number of shards to be created on a Premium Cluster Cache.
	ShardCount *int32 `json:"shardCount,omitempty"`
	// VirtualNetwork - The exact ARM resource ID of the virtual network to deploy the Redis cache in. Example format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.ClassicNetwork/VirtualNetworks/vnet1
	VirtualNetwork *string `json:"virtualNetwork,omitempty"`
	// Subnet - Required when deploying a Redis cache inside an existing Azure Virtual Network.
	Subnet *string `json:"subnet,omitempty"`
	// StaticIP - Required when deploying a Redis cache inside an existing Azure Virtual Network.
	StaticIP *string `json:"staticIP,omitempty"`
}

// MarshalJSON is the custom marshaler for ReadableProperties.
func (rp ReadableProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rp.ProvisioningState != nil {
		objectMap["provisioningState"] = rp.ProvisioningState
	}
	if rp.HostName != nil {
		objectMap["hostName"] = rp.HostName
	}
	if rp.Port != nil {
		objectMap["port"] = rp.Port
	}
	if rp.SslPort != nil {
		objectMap["sslPort"] = rp.SslPort
	}
	if rp.RedisVersion != nil {
		objectMap["redisVersion"] = rp.RedisVersion
	}
	if rp.Sku != nil {
		objectMap["sku"] = rp.Sku
	}
	if rp.RedisConfiguration != nil {
		objectMap["redisConfiguration"] = rp.RedisConfiguration
	}
	if rp.EnableNonSslPort != nil {
		objectMap["enableNonSslPort"] = rp.EnableNonSslPort
	}
	if rp.TenantSettings != nil {
		objectMap["tenantSettings"] = rp.TenantSettings
	}
	if rp.ShardCount != nil {
		objectMap["shardCount"] = rp.ShardCount
	}
	if rp.VirtualNetwork != nil {
		objectMap["virtualNetwork"] = rp.VirtualNetwork
	}
	if rp.Subnet != nil {
		objectMap["subnet"] = rp.Subnet
	}
	if rp.StaticIP != nil {
		objectMap["staticIP"] = rp.StaticIP
	}
	return json.Marshal(objectMap)
}

// ReadablePropertiesWithAccessKey properties generated only in response to CreateOrUpdate Redis operation.
type ReadablePropertiesWithAccessKey struct {
	// AccessKeys - Redis cache access keys.
	AccessKeys *AccessKeys `json:"accessKeys,omitempty"`
	// ProvisioningState - Redis instance provisioning status.
	ProvisioningState *string `json:"provisioningState,omitempty"`
	// HostName - Redis host name.
	HostName *string `json:"hostName,omitempty"`
	// Port - Redis non-SSL port.
	Port *int32 `json:"port,omitempty"`
	// SslPort - Redis SSL port.
	SslPort *int32 `json:"sslPort,omitempty"`
	// RedisVersion - RedisVersion parameter has been deprecated. As such, it is no longer necessary to provide this parameter and any value specified is ignored.
	RedisVersion *string `json:"redisVersion,omitempty"`
	// Sku - What SKU of Redis cache to deploy.
	Sku *Sku `json:"sku,omitempty"`
	// RedisConfiguration - All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
	RedisConfiguration map[string]*string `json:"redisConfiguration"`
	// EnableNonSslPort - If the value is true, then the non-SLL Redis server port (6379) will be enabled.
	EnableNonSslPort *bool `json:"enableNonSslPort,omitempty"`
	// TenantSettings - tenantSettings
	TenantSettings map[string]*string `json:"tenantSettings"`
	// ShardCount - The number of shards to be created on a Premium Cluster Cache.
	ShardCount *int32 `json:"shardCount,omitempty"`
	// VirtualNetwork - The exact ARM resource ID of the virtual network to deploy the Redis cache in. Example format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.ClassicNetwork/VirtualNetworks/vnet1
	VirtualNetwork *string `json:"virtualNetwork,omitempty"`
	// Subnet - Required when deploying a Redis cache inside an existing Azure Virtual Network.
	Subnet *string `json:"subnet,omitempty"`
	// StaticIP - Required when deploying a Redis cache inside an existing Azure Virtual Network.
	StaticIP *string `json:"staticIP,omitempty"`
}

// MarshalJSON is the custom marshaler for ReadablePropertiesWithAccessKey.
func (rpwak ReadablePropertiesWithAccessKey) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rpwak.AccessKeys != nil {
		objectMap["accessKeys"] = rpwak.AccessKeys
	}
	if rpwak.ProvisioningState != nil {
		objectMap["provisioningState"] = rpwak.ProvisioningState
	}
	if rpwak.HostName != nil {
		objectMap["hostName"] = rpwak.HostName
	}
	if rpwak.Port != nil {
		objectMap["port"] = rpwak.Port
	}
	if rpwak.SslPort != nil {
		objectMap["sslPort"] = rpwak.SslPort
	}
	if rpwak.RedisVersion != nil {
		objectMap["redisVersion"] = rpwak.RedisVersion
	}
	if rpwak.Sku != nil {
		objectMap["sku"] = rpwak.Sku
	}
	if rpwak.RedisConfiguration != nil {
		objectMap["redisConfiguration"] = rpwak.RedisConfiguration
	}
	if rpwak.EnableNonSslPort != nil {
		objectMap["enableNonSslPort"] = rpwak.EnableNonSslPort
	}
	if rpwak.TenantSettings != nil {
		objectMap["tenantSettings"] = rpwak.TenantSettings
	}
	if rpwak.ShardCount != nil {
		objectMap["shardCount"] = rpwak.ShardCount
	}
	if rpwak.VirtualNetwork != nil {
		objectMap["virtualNetwork"] = rpwak.VirtualNetwork
	}
	if rpwak.Subnet != nil {
		objectMap["subnet"] = rpwak.Subnet
	}
	if rpwak.StaticIP != nil {
		objectMap["staticIP"] = rpwak.StaticIP
	}
	return json.Marshal(objectMap)
}

// RebootParameters specifies which Redis node(s) to reboot.
type RebootParameters struct {
	// RebootType - Which Redis node(s) to reboot. Depending on this value data loss is possible. Possible values include: 'PrimaryNode', 'SecondaryNode', 'AllNodes'
	RebootType RebootType `json:"rebootType,omitempty"`
	// ShardID - If clustering is enabled, the ID of the shared be rebooted.
	ShardID *int32 `json:"shardId,omitempty"`
}

// RegenerateKeyParameters specifies which Redis access keys to reset.
type RegenerateKeyParameters struct {
	// KeyType - Which Redis access key to reset. Possible values include: 'Primary', 'Secondary'
	KeyType KeyType `json:"keyType,omitempty"`
}

// Resource ...
type Resource struct {
	// ID - READ-ONLY; Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags.
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

// ResourceType a single Redis item in List or Get Operation.
type ResourceType struct {
	autorest.Response `json:"-"`
	// ReadableProperties - Redis cache properties.
	*ReadableProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ResourceType.
func (rt ResourceType) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rt.ReadableProperties != nil {
		objectMap["properties"] = rt.ReadableProperties
	}
	if rt.Location != nil {
		objectMap["location"] = rt.Location
	}
	if rt.Tags != nil {
		objectMap["tags"] = rt.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ResourceType struct.
func (rt *ResourceType) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var readableProperties ReadableProperties
				err = json.Unmarshal(*v, &readableProperties)
				if err != nil {
					return err
				}
				rt.ReadableProperties = &readableProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				rt.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				rt.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				rt.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				rt.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				rt.Tags = tags
			}
		}
	}

	return nil
}

// ResourceWithAccessKey a Redis item in CreateOrUpdate Operation response.
type ResourceWithAccessKey struct {
	autorest.Response `json:"-"`
	// ReadablePropertiesWithAccessKey - Redis cache properties.
	*ReadablePropertiesWithAccessKey `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ResourceWithAccessKey.
func (rwak ResourceWithAccessKey) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rwak.ReadablePropertiesWithAccessKey != nil {
		objectMap["properties"] = rwak.ReadablePropertiesWithAccessKey
	}
	if rwak.Location != nil {
		objectMap["location"] = rwak.Location
	}
	if rwak.Tags != nil {
		objectMap["tags"] = rwak.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ResourceWithAccessKey struct.
func (rwak *ResourceWithAccessKey) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var readablePropertiesWithAccessKey ReadablePropertiesWithAccessKey
				err = json.Unmarshal(*v, &readablePropertiesWithAccessKey)
				if err != nil {
					return err
				}
				rwak.ReadablePropertiesWithAccessKey = &readablePropertiesWithAccessKey
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				rwak.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				rwak.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				rwak.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				rwak.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				rwak.Tags = tags
			}
		}
	}

	return nil
}

// Sku SKU parameters supplied to the create Redis operation.
type Sku struct {
	// Name - What type of Redis cache to deploy. Valid values: (Basic, Standard, Premium). Possible values include: 'Basic', 'Standard', 'Premium'
	Name SkuName `json:"name,omitempty"`
	// Family - Which family to use. Valid values: (C, P). Possible values include: 'C', 'P'
	Family SkuFamily `json:"family,omitempty"`
	// Capacity - What size of Redis cache to deploy. Valid values: for C family (0, 1, 2, 3, 4, 5, 6), for P family (1, 2, 3, 4).
	Capacity *int32 `json:"capacity,omitempty"`
}
