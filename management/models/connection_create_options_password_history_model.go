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

// ConnectionCreateOptionsPasswordHistory Options for password history policy
type ConnectionCreateOptionsPasswordHistory struct {
	Enable bool   `json:"enable"`
	Size   *int32 `json:"size,omitempty"`
}

type _ConnectionCreateOptionsPasswordHistory ConnectionCreateOptionsPasswordHistory

// GetEnable returns the Enable field value
func (o *ConnectionCreateOptionsPasswordHistory) GetEnable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value
// and a boolean to check if the value has been set.
func (o *ConnectionCreateOptionsPasswordHistory) GetEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enable, true
}

// SetEnable sets field value
func (o *ConnectionCreateOptionsPasswordHistory) SetEnable(v bool) {
	o.Enable = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ConnectionCreateOptionsPasswordHistory) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCreateOptionsPasswordHistory) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ConnectionCreateOptionsPasswordHistory) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *ConnectionCreateOptionsPasswordHistory) SetSize(v int32) {
	o.Size = &v
}

func (o ConnectionCreateOptionsPasswordHistory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionCreateOptionsPasswordHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enable"] = o.Enable
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

func (o *ConnectionCreateOptionsPasswordHistory) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enable",
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

	varConnectionCreateOptionsPasswordHistory := _ConnectionCreateOptionsPasswordHistory{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varConnectionCreateOptionsPasswordHistory)

	if err != nil {
		return err
	}

	*o = ConnectionCreateOptionsPasswordHistory(varConnectionCreateOptionsPasswordHistory)

	return err
}

type NullableConnectionCreateOptionsPasswordHistory struct {
	value *ConnectionCreateOptionsPasswordHistory
	isSet bool
}

func (v NullableConnectionCreateOptionsPasswordHistory) Get() *ConnectionCreateOptionsPasswordHistory {
	return v.value
}

func (v *NullableConnectionCreateOptionsPasswordHistory) Set(val *ConnectionCreateOptionsPasswordHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionCreateOptionsPasswordHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionCreateOptionsPasswordHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionCreateOptionsPasswordHistory(val *ConnectionCreateOptionsPasswordHistory) *NullableConnectionCreateOptionsPasswordHistory {
	return &NullableConnectionCreateOptionsPasswordHistory{value: val, isSet: true}
}

func (v NullableConnectionCreateOptionsPasswordHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionCreateOptionsPasswordHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
