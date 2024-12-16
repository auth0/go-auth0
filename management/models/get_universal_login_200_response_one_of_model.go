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

// GetUniversalLogin200ResponseOneOf struct for GetUniversalLogin200ResponseOneOf
type GetUniversalLogin200ResponseOneOf struct {
	// The custom page template for the New Universal Login Experience
	Body string `json:"body"`
}

type _GetUniversalLogin200ResponseOneOf GetUniversalLogin200ResponseOneOf

// GetBody returns the Body field value
func (o *GetUniversalLogin200ResponseOneOf) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *GetUniversalLogin200ResponseOneOf) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *GetUniversalLogin200ResponseOneOf) SetBody(v string) {
	o.Body = v
}

func (o GetUniversalLogin200ResponseOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUniversalLogin200ResponseOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["body"] = o.Body
	return toSerialize, nil
}

func (o *GetUniversalLogin200ResponseOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"body",
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

	varGetUniversalLogin200ResponseOneOf := _GetUniversalLogin200ResponseOneOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetUniversalLogin200ResponseOneOf)

	if err != nil {
		return err
	}

	*o = GetUniversalLogin200ResponseOneOf(varGetUniversalLogin200ResponseOneOf)

	return err
}

type NullableGetUniversalLogin200ResponseOneOf struct {
	value *GetUniversalLogin200ResponseOneOf
	isSet bool
}

func (v NullableGetUniversalLogin200ResponseOneOf) Get() *GetUniversalLogin200ResponseOneOf {
	return v.value
}

func (v *NullableGetUniversalLogin200ResponseOneOf) Set(val *GetUniversalLogin200ResponseOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUniversalLogin200ResponseOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUniversalLogin200ResponseOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUniversalLogin200ResponseOneOf(val *GetUniversalLogin200ResponseOneOf) *NullableGetUniversalLogin200ResponseOneOf {
	return &NullableGetUniversalLogin200ResponseOneOf{value: val, isSet: true}
}

func (v NullableGetUniversalLogin200ResponseOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUniversalLogin200ResponseOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
