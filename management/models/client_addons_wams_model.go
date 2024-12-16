/*
Auth0 Management API

Auth0 Management API v2.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
)

// ClientAddonsWams Windows Azure Mobile Services addon configuration.
type ClientAddonsWams struct {
	// Your master key for Windows Azure Mobile Services.
	Masterkey            string `json:"masterkey"`
	AdditionalProperties map[string]interface{}
}

type _ClientAddonsWams ClientAddonsWams

// GetMasterkey returns the Masterkey field value
func (o *ClientAddonsWams) GetMasterkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Masterkey
}

// GetMasterkeyOk returns a tuple with the Masterkey field value
// and a boolean to check if the value has been set.
func (o *ClientAddonsWams) GetMasterkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Masterkey, true
}

// SetMasterkey sets field value
func (o *ClientAddonsWams) SetMasterkey(v string) {
	o.Masterkey = v
}

func (o ClientAddonsWams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientAddonsWams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["masterkey"] = o.Masterkey

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientAddonsWams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"masterkey",
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

	varClientAddonsWams := _ClientAddonsWams{}

	err = json.Unmarshal(data, &varClientAddonsWams)

	if err != nil {
		return err
	}

	*o = ClientAddonsWams(varClientAddonsWams)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "masterkey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientAddonsWams struct {
	value *ClientAddonsWams
	isSet bool
}

func (v NullableClientAddonsWams) Get() *ClientAddonsWams {
	return v.value
}

func (v *NullableClientAddonsWams) Set(val *ClientAddonsWams) {
	v.value = val
	v.isSet = true
}

func (v NullableClientAddonsWams) IsSet() bool {
	return v.isSet
}

func (v *NullableClientAddonsWams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientAddonsWams(val *ClientAddonsWams) *NullableClientAddonsWams {
	return &NullableClientAddonsWams{value: val, isSet: true}
}

func (v NullableClientAddonsWams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientAddonsWams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
