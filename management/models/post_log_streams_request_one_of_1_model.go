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

// PostLogStreamsRequestOneOf1 struct for PostLogStreamsRequestOneOf1
type PostLogStreamsRequestOneOf1 struct {
	// log stream name
	Name *string                                 `json:"name,omitempty"`
	Type GetLogStreams200ResponseInnerOneOf1Type `json:"type"`
	// Only logs events matching these filters will be delivered by the stream. If omitted or empty, all events will be delivered.
	Filters []GetLogStreams200ResponseInnerOneOfFiltersInner `json:"filters,omitempty"`
	Sink    PostLogStreamsRequestOneOf1Sink                  `json:"sink"`
	// The optional datetime (ISO 8601) to start streaming logs from
	StartFrom *string `json:"startFrom,omitempty"`
}

type _PostLogStreamsRequestOneOf1 PostLogStreamsRequestOneOf1

// GetName returns the Name field value if set, zero value otherwise.
func (o *PostLogStreamsRequestOneOf1) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLogStreamsRequestOneOf1) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PostLogStreamsRequestOneOf1) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PostLogStreamsRequestOneOf1) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value
func (o *PostLogStreamsRequestOneOf1) GetType() GetLogStreams200ResponseInnerOneOf1Type {
	if o == nil {
		var ret GetLogStreams200ResponseInnerOneOf1Type
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostLogStreamsRequestOneOf1) GetTypeOk() (*GetLogStreams200ResponseInnerOneOf1Type, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostLogStreamsRequestOneOf1) SetType(v GetLogStreams200ResponseInnerOneOf1Type) {
	o.Type = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *PostLogStreamsRequestOneOf1) GetFilters() []GetLogStreams200ResponseInnerOneOfFiltersInner {
	if o == nil || IsNil(o.Filters) {
		var ret []GetLogStreams200ResponseInnerOneOfFiltersInner
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLogStreamsRequestOneOf1) GetFiltersOk() ([]GetLogStreams200ResponseInnerOneOfFiltersInner, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *PostLogStreamsRequestOneOf1) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []GetLogStreams200ResponseInnerOneOfFiltersInner and assigns it to the Filters field.
func (o *PostLogStreamsRequestOneOf1) SetFilters(v []GetLogStreams200ResponseInnerOneOfFiltersInner) {
	o.Filters = v
}

// GetSink returns the Sink field value
func (o *PostLogStreamsRequestOneOf1) GetSink() PostLogStreamsRequestOneOf1Sink {
	if o == nil {
		var ret PostLogStreamsRequestOneOf1Sink
		return ret
	}

	return o.Sink
}

// GetSinkOk returns a tuple with the Sink field value
// and a boolean to check if the value has been set.
func (o *PostLogStreamsRequestOneOf1) GetSinkOk() (*PostLogStreamsRequestOneOf1Sink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sink, true
}

// SetSink sets field value
func (o *PostLogStreamsRequestOneOf1) SetSink(v PostLogStreamsRequestOneOf1Sink) {
	o.Sink = v
}

// GetStartFrom returns the StartFrom field value if set, zero value otherwise.
func (o *PostLogStreamsRequestOneOf1) GetStartFrom() string {
	if o == nil || IsNil(o.StartFrom) {
		var ret string
		return ret
	}
	return *o.StartFrom
}

// GetStartFromOk returns a tuple with the StartFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLogStreamsRequestOneOf1) GetStartFromOk() (*string, bool) {
	if o == nil || IsNil(o.StartFrom) {
		return nil, false
	}
	return o.StartFrom, true
}

// HasStartFrom returns a boolean if a field has been set.
func (o *PostLogStreamsRequestOneOf1) HasStartFrom() bool {
	if o != nil && !IsNil(o.StartFrom) {
		return true
	}

	return false
}

// SetStartFrom gets a reference to the given string and assigns it to the StartFrom field.
func (o *PostLogStreamsRequestOneOf1) SetStartFrom(v string) {
	o.StartFrom = &v
}

func (o PostLogStreamsRequestOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostLogStreamsRequestOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	toSerialize["sink"] = o.Sink
	if !IsNil(o.StartFrom) {
		toSerialize["startFrom"] = o.StartFrom
	}
	return toSerialize, nil
}

func (o *PostLogStreamsRequestOneOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"sink",
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

	varPostLogStreamsRequestOneOf1 := _PostLogStreamsRequestOneOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varPostLogStreamsRequestOneOf1)

	if err != nil {
		return err
	}

	*o = PostLogStreamsRequestOneOf1(varPostLogStreamsRequestOneOf1)

	return err
}

type NullablePostLogStreamsRequestOneOf1 struct {
	value *PostLogStreamsRequestOneOf1
	isSet bool
}

func (v NullablePostLogStreamsRequestOneOf1) Get() *PostLogStreamsRequestOneOf1 {
	return v.value
}

func (v *NullablePostLogStreamsRequestOneOf1) Set(val *PostLogStreamsRequestOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullablePostLogStreamsRequestOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullablePostLogStreamsRequestOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostLogStreamsRequestOneOf1(val *PostLogStreamsRequestOneOf1) *NullablePostLogStreamsRequestOneOf1 {
	return &NullablePostLogStreamsRequestOneOf1{value: val, isSet: true}
}

func (v NullablePostLogStreamsRequestOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostLogStreamsRequestOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
