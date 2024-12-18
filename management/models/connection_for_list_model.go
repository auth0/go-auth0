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

// ConnectionForList struct for ConnectionForList
type ConnectionForList struct {
	// The name of the connection
	Name *string `json:"name,omitempty"`
	// Connection name used in login screen
	DisplayName *string `json:"display_name,omitempty"`
	// In order to return options in the response, the `read:connections_options` scope must be present
	Options map[string]interface{} `json:"options,omitempty"`
	// The connection's identifier
	Id *string `json:"id,omitempty"`
	// The type of the connection, related to the identity provider
	Strategy *string `json:"strategy,omitempty"`
	// Defines the realms for which the connection will be used (ie: email domains). If the array is empty or the property is not specified, the connection name will be added as realm.
	Realms []string `json:"realms,omitempty"`
	// True if the connection is domain level
	IsDomainConnection *bool `json:"is_domain_connection,omitempty"`
	// Metadata associated with the connection in the form of an object with string values (max 255 chars).  Maximum of 10 metadata properties allowed.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectionForList) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionForList) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectionForList) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectionForList) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ConnectionForList) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionForList) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ConnectionForList) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ConnectionForList) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ConnectionForList) GetOptions() map[string]interface{} {
	if o == nil || IsNil(o.Options) {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionForList) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return map[string]interface{}{}, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ConnectionForList) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *ConnectionForList) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConnectionForList) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionForList) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConnectionForList) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConnectionForList) SetId(v string) {
	o.Id = &v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *ConnectionForList) GetStrategy() string {
	if o == nil || IsNil(o.Strategy) {
		var ret string
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionForList) GetStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.Strategy) {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *ConnectionForList) HasStrategy() bool {
	if o != nil && !IsNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given string and assigns it to the Strategy field.
func (o *ConnectionForList) SetStrategy(v string) {
	o.Strategy = &v
}

// GetRealms returns the Realms field value if set, zero value otherwise.
func (o *ConnectionForList) GetRealms() []string {
	if o == nil || IsNil(o.Realms) {
		var ret []string
		return ret
	}
	return o.Realms
}

// GetRealmsOk returns a tuple with the Realms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionForList) GetRealmsOk() ([]string, bool) {
	if o == nil || IsNil(o.Realms) {
		return nil, false
	}
	return o.Realms, true
}

// HasRealms returns a boolean if a field has been set.
func (o *ConnectionForList) HasRealms() bool {
	if o != nil && !IsNil(o.Realms) {
		return true
	}

	return false
}

// SetRealms gets a reference to the given []string and assigns it to the Realms field.
func (o *ConnectionForList) SetRealms(v []string) {
	o.Realms = v
}

// GetIsDomainConnection returns the IsDomainConnection field value if set, zero value otherwise.
func (o *ConnectionForList) GetIsDomainConnection() bool {
	if o == nil || IsNil(o.IsDomainConnection) {
		var ret bool
		return ret
	}
	return *o.IsDomainConnection
}

// GetIsDomainConnectionOk returns a tuple with the IsDomainConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionForList) GetIsDomainConnectionOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDomainConnection) {
		return nil, false
	}
	return o.IsDomainConnection, true
}

// HasIsDomainConnection returns a boolean if a field has been set.
func (o *ConnectionForList) HasIsDomainConnection() bool {
	if o != nil && !IsNil(o.IsDomainConnection) {
		return true
	}

	return false
}

// SetIsDomainConnection gets a reference to the given bool and assigns it to the IsDomainConnection field.
func (o *ConnectionForList) SetIsDomainConnection(v bool) {
	o.IsDomainConnection = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ConnectionForList) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionForList) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ConnectionForList) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ConnectionForList) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o ConnectionForList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionForList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	if !IsNil(o.Realms) {
		toSerialize["realms"] = o.Realms
	}
	if !IsNil(o.IsDomainConnection) {
		toSerialize["is_domain_connection"] = o.IsDomainConnection
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableConnectionForList struct {
	value *ConnectionForList
	isSet bool
}

func (v NullableConnectionForList) Get() *ConnectionForList {
	return v.value
}

func (v *NullableConnectionForList) Set(val *ConnectionForList) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionForList) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionForList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionForList(val *ConnectionForList) *NullableConnectionForList {
	return &NullableConnectionForList{value: val, isSet: true}
}

func (v NullableConnectionForList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionForList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
