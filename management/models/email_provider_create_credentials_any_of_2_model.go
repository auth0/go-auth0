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

// EmailProviderCreateCredentialsAnyOf2 struct for EmailProviderCreateCredentialsAnyOf2
type EmailProviderCreateCredentialsAnyOf2 struct {
	// API Key
	ApiKey string                                      `json:"api_key"`
	Region *EmailProviderUpdateCredentialsAnyOf3Region `json:"region,omitempty"`
}

type _EmailProviderCreateCredentialsAnyOf2 EmailProviderCreateCredentialsAnyOf2

// GetApiKey returns the ApiKey field value
func (o *EmailProviderCreateCredentialsAnyOf2) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *EmailProviderCreateCredentialsAnyOf2) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *EmailProviderCreateCredentialsAnyOf2) SetApiKey(v string) {
	o.ApiKey = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *EmailProviderCreateCredentialsAnyOf2) GetRegion() EmailProviderUpdateCredentialsAnyOf3Region {
	if o == nil || IsNil(o.Region) {
		var ret EmailProviderUpdateCredentialsAnyOf3Region
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailProviderCreateCredentialsAnyOf2) GetRegionOk() (*EmailProviderUpdateCredentialsAnyOf3Region, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *EmailProviderCreateCredentialsAnyOf2) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given EmailProviderUpdateCredentialsAnyOf3Region and assigns it to the Region field.
func (o *EmailProviderCreateCredentialsAnyOf2) SetRegion(v EmailProviderUpdateCredentialsAnyOf3Region) {
	o.Region = &v
}

func (o EmailProviderCreateCredentialsAnyOf2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailProviderCreateCredentialsAnyOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["api_key"] = o.ApiKey
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	return toSerialize, nil
}

func (o *EmailProviderCreateCredentialsAnyOf2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"api_key",
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

	varEmailProviderCreateCredentialsAnyOf2 := _EmailProviderCreateCredentialsAnyOf2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmailProviderCreateCredentialsAnyOf2)

	if err != nil {
		return err
	}

	*o = EmailProviderCreateCredentialsAnyOf2(varEmailProviderCreateCredentialsAnyOf2)

	return err
}

type NullableEmailProviderCreateCredentialsAnyOf2 struct {
	value *EmailProviderCreateCredentialsAnyOf2
	isSet bool
}

func (v NullableEmailProviderCreateCredentialsAnyOf2) Get() *EmailProviderCreateCredentialsAnyOf2 {
	return v.value
}

func (v *NullableEmailProviderCreateCredentialsAnyOf2) Set(val *EmailProviderCreateCredentialsAnyOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailProviderCreateCredentialsAnyOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailProviderCreateCredentialsAnyOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailProviderCreateCredentialsAnyOf2(val *EmailProviderCreateCredentialsAnyOf2) *NullableEmailProviderCreateCredentialsAnyOf2 {
	return &NullableEmailProviderCreateCredentialsAnyOf2{value: val, isSet: true}
}

func (v NullableEmailProviderCreateCredentialsAnyOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailProviderCreateCredentialsAnyOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
