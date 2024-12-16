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
)

// ConnectionCreateOptionsPasswordDictionary Options for password dictionary policy
type ConnectionCreateOptionsPasswordDictionary struct {
	Enable bool `json:"enable"`
	// Custom Password Dictionary. An array of up to 200 entries.
	Dictionary []string `json:"dictionary,omitempty"`
}

type _ConnectionCreateOptionsPasswordDictionary ConnectionCreateOptionsPasswordDictionary

// GetEnable returns the Enable field value
func (o *ConnectionCreateOptionsPasswordDictionary) GetEnable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value
// and a boolean to check if the value has been set.
func (o *ConnectionCreateOptionsPasswordDictionary) GetEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enable, true
}

// SetEnable sets field value
func (o *ConnectionCreateOptionsPasswordDictionary) SetEnable(v bool) {
	o.Enable = v
}

// GetDictionary returns the Dictionary field value if set, zero value otherwise.
func (o *ConnectionCreateOptionsPasswordDictionary) GetDictionary() []string {
	if o == nil || IsNil(o.Dictionary) {
		var ret []string
		return ret
	}
	return o.Dictionary
}

// GetDictionaryOk returns a tuple with the Dictionary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCreateOptionsPasswordDictionary) GetDictionaryOk() ([]string, bool) {
	if o == nil || IsNil(o.Dictionary) {
		return nil, false
	}
	return o.Dictionary, true
}

// HasDictionary returns a boolean if a field has been set.
func (o *ConnectionCreateOptionsPasswordDictionary) HasDictionary() bool {
	if o != nil && !IsNil(o.Dictionary) {
		return true
	}

	return false
}

// SetDictionary gets a reference to the given []string and assigns it to the Dictionary field.
func (o *ConnectionCreateOptionsPasswordDictionary) SetDictionary(v []string) {
	o.Dictionary = v
}

func (o ConnectionCreateOptionsPasswordDictionary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionCreateOptionsPasswordDictionary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enable"] = o.Enable
	if !IsNil(o.Dictionary) {
		toSerialize["dictionary"] = o.Dictionary
	}
	return toSerialize, nil
}

func (o *ConnectionCreateOptionsPasswordDictionary) UnmarshalJSON(data []byte) (err error) {
	varConnectionCreateOptionsPasswordDictionary := _ConnectionCreateOptionsPasswordDictionary{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConnectionCreateOptionsPasswordDictionary)

	if err != nil {
		return err
	}

	*o = ConnectionCreateOptionsPasswordDictionary(varConnectionCreateOptionsPasswordDictionary)

	return err
}

type NullableConnectionCreateOptionsPasswordDictionary struct {
	value *ConnectionCreateOptionsPasswordDictionary
	isSet bool
}

func (v NullableConnectionCreateOptionsPasswordDictionary) Get() *ConnectionCreateOptionsPasswordDictionary {
	return v.value
}

func (v *NullableConnectionCreateOptionsPasswordDictionary) Set(val *ConnectionCreateOptionsPasswordDictionary) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionCreateOptionsPasswordDictionary) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionCreateOptionsPasswordDictionary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionCreateOptionsPasswordDictionary(val *ConnectionCreateOptionsPasswordDictionary) *NullableConnectionCreateOptionsPasswordDictionary {
	return &NullableConnectionCreateOptionsPasswordDictionary{value: val, isSet: true}
}

func (v NullableConnectionCreateOptionsPasswordDictionary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionCreateOptionsPasswordDictionary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
