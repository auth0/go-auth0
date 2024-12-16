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

// GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner struct for GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner
type GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner struct {
	// HTTP header name
	Header *string `json:"header,omitempty"`
	// HTTP header value
	Value                *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner) GetHeader() string {
	if o == nil || IsNil(o.Header) {
		var ret string
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner) GetHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given string and assigns it to the Header field.
func (o *GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner) SetHeader(v string) {
	o.Header = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner) SetValue(v string) {
	o.Value = &v
}

func (o GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner) UnmarshalJSON(data []byte) (err error) {
	varGetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner := _GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner{}

	err = json.Unmarshal(data, &varGetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner)

	if err != nil {
		return err
	}

	*o = GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner(varGetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "header")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner struct {
	value *GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner
	isSet bool
}

func (v NullableGetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner) Get() *GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner {
	return v.value
}

func (v *NullableGetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner) Set(val *GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner(val *GetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner) *NullableGetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner {
	return &NullableGetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner{value: val, isSet: true}
}

func (v NullableGetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLogStreams200ResponseInnerOneOfSinkHttpCustomHeadersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
