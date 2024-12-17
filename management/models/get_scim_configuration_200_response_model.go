/*
Auth0 Management API

Auth0 Management API v2.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// GetScimConfiguration200Response struct for GetScimConfiguration200Response
type GetScimConfiguration200Response struct {
	// The connection's identifier
	ConnectionId *string `json:"connection_id,omitempty"`
	// The connection's identifier
	ConnectionName *string `json:"connection_name,omitempty"`
	// The connection's strategy
	Strategy *string `json:"strategy,omitempty"`
	// The tenant's name
	TenantName *string `json:"tenant_name,omitempty"`
	// User ID attribute for generating unique user ids
	UserIdAttribute *string `json:"user_id_attribute,omitempty"`
	// The mapping between auth0 and SCIM
	Mapping []GetScimConfiguration200ResponseMappingInner `json:"mapping,omitempty"`
	// The Date Time Scim Configuration was created
	CreatedAt *string `json:"created_at,omitempty"`
	// The Date Time Scim Configuration was last updated
	UpdatedOn *string `json:"updated_on,omitempty"`
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *GetScimConfiguration200Response) GetConnectionId() string {
	if o == nil || IsNil(o.ConnectionId) {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScimConfiguration200Response) GetConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionId) {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *GetScimConfiguration200Response) HasConnectionId() bool {
	if o != nil && !IsNil(o.ConnectionId) {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *GetScimConfiguration200Response) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetConnectionName returns the ConnectionName field value if set, zero value otherwise.
func (o *GetScimConfiguration200Response) GetConnectionName() string {
	if o == nil || IsNil(o.ConnectionName) {
		var ret string
		return ret
	}
	return *o.ConnectionName
}

// GetConnectionNameOk returns a tuple with the ConnectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScimConfiguration200Response) GetConnectionNameOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionName) {
		return nil, false
	}
	return o.ConnectionName, true
}

// HasConnectionName returns a boolean if a field has been set.
func (o *GetScimConfiguration200Response) HasConnectionName() bool {
	if o != nil && !IsNil(o.ConnectionName) {
		return true
	}

	return false
}

// SetConnectionName gets a reference to the given string and assigns it to the ConnectionName field.
func (o *GetScimConfiguration200Response) SetConnectionName(v string) {
	o.ConnectionName = &v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *GetScimConfiguration200Response) GetStrategy() string {
	if o == nil || IsNil(o.Strategy) {
		var ret string
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScimConfiguration200Response) GetStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.Strategy) {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *GetScimConfiguration200Response) HasStrategy() bool {
	if o != nil && !IsNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given string and assigns it to the Strategy field.
func (o *GetScimConfiguration200Response) SetStrategy(v string) {
	o.Strategy = &v
}

// GetTenantName returns the TenantName field value if set, zero value otherwise.
func (o *GetScimConfiguration200Response) GetTenantName() string {
	if o == nil || IsNil(o.TenantName) {
		var ret string
		return ret
	}
	return *o.TenantName
}

// GetTenantNameOk returns a tuple with the TenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScimConfiguration200Response) GetTenantNameOk() (*string, bool) {
	if o == nil || IsNil(o.TenantName) {
		return nil, false
	}
	return o.TenantName, true
}

// HasTenantName returns a boolean if a field has been set.
func (o *GetScimConfiguration200Response) HasTenantName() bool {
	if o != nil && !IsNil(o.TenantName) {
		return true
	}

	return false
}

// SetTenantName gets a reference to the given string and assigns it to the TenantName field.
func (o *GetScimConfiguration200Response) SetTenantName(v string) {
	o.TenantName = &v
}

// GetUserIdAttribute returns the UserIdAttribute field value if set, zero value otherwise.
func (o *GetScimConfiguration200Response) GetUserIdAttribute() string {
	if o == nil || IsNil(o.UserIdAttribute) {
		var ret string
		return ret
	}
	return *o.UserIdAttribute
}

// GetUserIdAttributeOk returns a tuple with the UserIdAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScimConfiguration200Response) GetUserIdAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.UserIdAttribute) {
		return nil, false
	}
	return o.UserIdAttribute, true
}

// HasUserIdAttribute returns a boolean if a field has been set.
func (o *GetScimConfiguration200Response) HasUserIdAttribute() bool {
	if o != nil && !IsNil(o.UserIdAttribute) {
		return true
	}

	return false
}

// SetUserIdAttribute gets a reference to the given string and assigns it to the UserIdAttribute field.
func (o *GetScimConfiguration200Response) SetUserIdAttribute(v string) {
	o.UserIdAttribute = &v
}

// GetMapping returns the Mapping field value if set, zero value otherwise.
func (o *GetScimConfiguration200Response) GetMapping() []GetScimConfiguration200ResponseMappingInner {
	if o == nil || IsNil(o.Mapping) {
		var ret []GetScimConfiguration200ResponseMappingInner
		return ret
	}
	return o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScimConfiguration200Response) GetMappingOk() ([]GetScimConfiguration200ResponseMappingInner, bool) {
	if o == nil || IsNil(o.Mapping) {
		return nil, false
	}
	return o.Mapping, true
}

// HasMapping returns a boolean if a field has been set.
func (o *GetScimConfiguration200Response) HasMapping() bool {
	if o != nil && !IsNil(o.Mapping) {
		return true
	}

	return false
}

// SetMapping gets a reference to the given []GetScimConfiguration200ResponseMappingInner and assigns it to the Mapping field.
func (o *GetScimConfiguration200Response) SetMapping(v []GetScimConfiguration200ResponseMappingInner) {
	o.Mapping = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GetScimConfiguration200Response) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScimConfiguration200Response) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GetScimConfiguration200Response) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GetScimConfiguration200Response) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *GetScimConfiguration200Response) GetUpdatedOn() string {
	if o == nil || IsNil(o.UpdatedOn) {
		var ret string
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScimConfiguration200Response) GetUpdatedOnOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedOn) {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *GetScimConfiguration200Response) HasUpdatedOn() bool {
	if o != nil && !IsNil(o.UpdatedOn) {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given string and assigns it to the UpdatedOn field.
func (o *GetScimConfiguration200Response) SetUpdatedOn(v string) {
	o.UpdatedOn = &v
}

func (o GetScimConfiguration200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetScimConfiguration200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConnectionId) {
		toSerialize["connection_id"] = o.ConnectionId
	}
	if !IsNil(o.ConnectionName) {
		toSerialize["connection_name"] = o.ConnectionName
	}
	if !IsNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	if !IsNil(o.TenantName) {
		toSerialize["tenant_name"] = o.TenantName
	}
	if !IsNil(o.UserIdAttribute) {
		toSerialize["user_id_attribute"] = o.UserIdAttribute
	}
	if !IsNil(o.Mapping) {
		toSerialize["mapping"] = o.Mapping
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedOn) {
		toSerialize["updated_on"] = o.UpdatedOn
	}
	return toSerialize, nil
}

type NullableGetScimConfiguration200Response struct {
	value *GetScimConfiguration200Response
	isSet bool
}

func (v NullableGetScimConfiguration200Response) Get() *GetScimConfiguration200Response {
	return v.value
}

func (v *NullableGetScimConfiguration200Response) Set(val *GetScimConfiguration200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetScimConfiguration200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetScimConfiguration200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetScimConfiguration200Response(val *GetScimConfiguration200Response) *NullableGetScimConfiguration200Response {
	return &NullableGetScimConfiguration200Response{value: val, isSet: true}
}

func (v NullableGetScimConfiguration200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetScimConfiguration200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
