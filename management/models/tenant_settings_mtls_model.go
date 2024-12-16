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

// TenantSettingsMtls mTLS configuration.
type TenantSettingsMtls struct {
	// If true, enables mTLS endpoint aliases
	EnableEndpointAliases bool `json:"enable_endpoint_aliases"`
}

type _TenantSettingsMtls TenantSettingsMtls

// GetEnableEndpointAliases returns the EnableEndpointAliases field value
func (o *TenantSettingsMtls) GetEnableEndpointAliases() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableEndpointAliases
}

// GetEnableEndpointAliasesOk returns a tuple with the EnableEndpointAliases field value
// and a boolean to check if the value has been set.
func (o *TenantSettingsMtls) GetEnableEndpointAliasesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableEndpointAliases, true
}

// SetEnableEndpointAliases sets field value
func (o *TenantSettingsMtls) SetEnableEndpointAliases(v bool) {
	o.EnableEndpointAliases = v
}

func (o TenantSettingsMtls) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantSettingsMtls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enable_endpoint_aliases"] = o.EnableEndpointAliases
	return toSerialize, nil
}

func (o *TenantSettingsMtls) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enable_endpoint_aliases",
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

	varTenantSettingsMtls := _TenantSettingsMtls{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTenantSettingsMtls)

	if err != nil {
		return err
	}

	*o = TenantSettingsMtls(varTenantSettingsMtls)

	return err
}

type NullableTenantSettingsMtls struct {
	value *TenantSettingsMtls
	isSet bool
}

func (v NullableTenantSettingsMtls) Get() *TenantSettingsMtls {
	return v.value
}

func (v *NullableTenantSettingsMtls) Set(val *TenantSettingsMtls) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantSettingsMtls) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantSettingsMtls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantSettingsMtls(val *TenantSettingsMtls) *NullableTenantSettingsMtls {
	return &NullableTenantSettingsMtls{value: val, isSet: true}
}

func (v NullableTenantSettingsMtls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantSettingsMtls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
