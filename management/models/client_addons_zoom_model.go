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

// ClientAddonsZoom Zoom SSO configuration.
type ClientAddonsZoom struct {
	// Zoom account name usually first segment of your Zoom URL, e.g. `https://acme-org.zoom.us` would be `acme-org`.
	Account              string `json:"account"`
	AdditionalProperties map[string]interface{}
}

type _ClientAddonsZoom ClientAddonsZoom

// GetAccount returns the Account field value
func (o *ClientAddonsZoom) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *ClientAddonsZoom) GetAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *ClientAddonsZoom) SetAccount(v string) {
	o.Account = v
}

func (o ClientAddonsZoom) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientAddonsZoom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account"] = o.Account

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientAddonsZoom) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account",
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

	varClientAddonsZoom := _ClientAddonsZoom{}

	err = json.Unmarshal(data, &varClientAddonsZoom)

	if err != nil {
		return err
	}

	*o = ClientAddonsZoom(varClientAddonsZoom)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientAddonsZoom struct {
	value *ClientAddonsZoom
	isSet bool
}

func (v NullableClientAddonsZoom) Get() *ClientAddonsZoom {
	return v.value
}

func (v *NullableClientAddonsZoom) Set(val *ClientAddonsZoom) {
	v.value = val
	v.isSet = true
}

func (v NullableClientAddonsZoom) IsSet() bool {
	return v.isSet
}

func (v *NullableClientAddonsZoom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientAddonsZoom(val *ClientAddonsZoom) *NullableClientAddonsZoom {
	return &NullableClientAddonsZoom{value: val, isSet: true}
}

func (v NullableClientAddonsZoom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientAddonsZoom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
