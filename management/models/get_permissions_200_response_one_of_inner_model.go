/*
Auth0 Management API

Auth0 Management API v2.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// GetPermissions200ResponseOneOfInner struct for GetPermissions200ResponseOneOfInner
type GetPermissions200ResponseOneOfInner struct {
	Sources interface{} `json:"sources"`
	// Resource server (API) identifier that this permission is for.
	ResourceServerIdentifier string `json:"resource_server_identifier"`
	// Name of this permission.
	PermissionName string `json:"permission_name"`
	// Resource server (API) name this permission is for.
	ResourceServerName string `json:"resource_server_name"`
	// Description of this permission.
	Description string `json:"description"`
}

type _GetPermissions200ResponseOneOfInner GetPermissions200ResponseOneOfInner

// GetSources returns the Sources field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GetPermissions200ResponseOneOfInner) GetSources() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetPermissions200ResponseOneOfInner) GetSourcesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value
func (o *GetPermissions200ResponseOneOfInner) SetSources(v interface{}) {
	o.Sources = v
}

// GetResourceServerIdentifier returns the ResourceServerIdentifier field value
func (o *GetPermissions200ResponseOneOfInner) GetResourceServerIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceServerIdentifier
}

// GetResourceServerIdentifierOk returns a tuple with the ResourceServerIdentifier field value
// and a boolean to check if the value has been set.
func (o *GetPermissions200ResponseOneOfInner) GetResourceServerIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceServerIdentifier, true
}

// SetResourceServerIdentifier sets field value
func (o *GetPermissions200ResponseOneOfInner) SetResourceServerIdentifier(v string) {
	o.ResourceServerIdentifier = v
}

// GetPermissionName returns the PermissionName field value
func (o *GetPermissions200ResponseOneOfInner) GetPermissionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PermissionName
}

// GetPermissionNameOk returns a tuple with the PermissionName field value
// and a boolean to check if the value has been set.
func (o *GetPermissions200ResponseOneOfInner) GetPermissionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PermissionName, true
}

// SetPermissionName sets field value
func (o *GetPermissions200ResponseOneOfInner) SetPermissionName(v string) {
	o.PermissionName = v
}

// GetResourceServerName returns the ResourceServerName field value
func (o *GetPermissions200ResponseOneOfInner) GetResourceServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceServerName
}

// GetResourceServerNameOk returns a tuple with the ResourceServerName field value
// and a boolean to check if the value has been set.
func (o *GetPermissions200ResponseOneOfInner) GetResourceServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceServerName, true
}

// SetResourceServerName sets field value
func (o *GetPermissions200ResponseOneOfInner) SetResourceServerName(v string) {
	o.ResourceServerName = v
}

// GetDescription returns the Description field value
func (o *GetPermissions200ResponseOneOfInner) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GetPermissions200ResponseOneOfInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *GetPermissions200ResponseOneOfInner) SetDescription(v string) {
	o.Description = v
}

func (o GetPermissions200ResponseOneOfInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPermissions200ResponseOneOfInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	toSerialize["resource_server_identifier"] = o.ResourceServerIdentifier
	toSerialize["permission_name"] = o.PermissionName
	toSerialize["resource_server_name"] = o.ResourceServerName
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

func (o *GetPermissions200ResponseOneOfInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sources",
		"resource_server_identifier",
		"permission_name",
		"resource_server_name",
		"description",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetPermissions200ResponseOneOfInner := _GetPermissions200ResponseOneOfInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetPermissions200ResponseOneOfInner)

	if err != nil {
		return err
	}

	*o = GetPermissions200ResponseOneOfInner(varGetPermissions200ResponseOneOfInner)

	return err
}

type NullableGetPermissions200ResponseOneOfInner struct {
	value *GetPermissions200ResponseOneOfInner
	isSet bool
}

func (v NullableGetPermissions200ResponseOneOfInner) Get() *GetPermissions200ResponseOneOfInner {
	return v.value
}

func (v *NullableGetPermissions200ResponseOneOfInner) Set(val *GetPermissions200ResponseOneOfInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPermissions200ResponseOneOfInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPermissions200ResponseOneOfInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPermissions200ResponseOneOfInner(val *GetPermissions200ResponseOneOfInner) *NullableGetPermissions200ResponseOneOfInner {
	return &NullableGetPermissions200ResponseOneOfInner{value: val, isSet: true}
}

func (v NullableGetPermissions200ResponseOneOfInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPermissions200ResponseOneOfInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
