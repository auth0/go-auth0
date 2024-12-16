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

// GetLogStreams200ResponseInnerOneOf3 struct for GetLogStreams200ResponseInnerOneOf3
type GetLogStreams200ResponseInnerOneOf3 struct {
	// The id of the log stream
	Id string `json:"id"`
	// log stream name
	Name   string                                   `json:"name"`
	Status GetLogStreams200ResponseInnerOneOfStatus `json:"status"`
	Type   GetLogStreams200ResponseInnerOneOf3Type  `json:"type"`
	// Only logs events matching these filters will be delivered by the stream. If omitted or empty, all events will be delivered.
	Filters              []GetLogStreams200ResponseInnerOneOfFiltersInner `json:"filters"`
	Sink                 GetLogStreams200ResponseInnerOneOf3Sink          `json:"sink"`
	AdditionalProperties map[string]interface{}
}

type _GetLogStreams200ResponseInnerOneOf3 GetLogStreams200ResponseInnerOneOf3

// GetId returns the Id field value
func (o *GetLogStreams200ResponseInnerOneOf3) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf3) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetLogStreams200ResponseInnerOneOf3) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GetLogStreams200ResponseInnerOneOf3) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf3) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetLogStreams200ResponseInnerOneOf3) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *GetLogStreams200ResponseInnerOneOf3) GetStatus() GetLogStreams200ResponseInnerOneOfStatus {
	if o == nil {
		var ret GetLogStreams200ResponseInnerOneOfStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf3) GetStatusOk() (*GetLogStreams200ResponseInnerOneOfStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetLogStreams200ResponseInnerOneOf3) SetStatus(v GetLogStreams200ResponseInnerOneOfStatus) {
	o.Status = v
}

// GetType returns the Type field value
func (o *GetLogStreams200ResponseInnerOneOf3) GetType() GetLogStreams200ResponseInnerOneOf3Type {
	if o == nil {
		var ret GetLogStreams200ResponseInnerOneOf3Type
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf3) GetTypeOk() (*GetLogStreams200ResponseInnerOneOf3Type, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetLogStreams200ResponseInnerOneOf3) SetType(v GetLogStreams200ResponseInnerOneOf3Type) {
	o.Type = v
}

// GetFilters returns the Filters field value
func (o *GetLogStreams200ResponseInnerOneOf3) GetFilters() []GetLogStreams200ResponseInnerOneOfFiltersInner {
	if o == nil {
		var ret []GetLogStreams200ResponseInnerOneOfFiltersInner
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf3) GetFiltersOk() ([]GetLogStreams200ResponseInnerOneOfFiltersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filters, true
}

// SetFilters sets field value
func (o *GetLogStreams200ResponseInnerOneOf3) SetFilters(v []GetLogStreams200ResponseInnerOneOfFiltersInner) {
	o.Filters = v
}

// GetSink returns the Sink field value
func (o *GetLogStreams200ResponseInnerOneOf3) GetSink() GetLogStreams200ResponseInnerOneOf3Sink {
	if o == nil {
		var ret GetLogStreams200ResponseInnerOneOf3Sink
		return ret
	}

	return o.Sink
}

// GetSinkOk returns a tuple with the Sink field value
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf3) GetSinkOk() (*GetLogStreams200ResponseInnerOneOf3Sink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sink, true
}

// SetSink sets field value
func (o *GetLogStreams200ResponseInnerOneOf3) SetSink(v GetLogStreams200ResponseInnerOneOf3Sink) {
	o.Sink = v
}

func (o GetLogStreams200ResponseInnerOneOf3) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLogStreams200ResponseInnerOneOf3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	toSerialize["filters"] = o.Filters
	toSerialize["sink"] = o.Sink

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetLogStreams200ResponseInnerOneOf3) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"status",
		"type",
		"filters",
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

	varGetLogStreams200ResponseInnerOneOf3 := _GetLogStreams200ResponseInnerOneOf3{}

	err = json.Unmarshal(data, &varGetLogStreams200ResponseInnerOneOf3)

	if err != nil {
		return err
	}

	*o = GetLogStreams200ResponseInnerOneOf3(varGetLogStreams200ResponseInnerOneOf3)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "filters")
		delete(additionalProperties, "sink")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetLogStreams200ResponseInnerOneOf3 struct {
	value *GetLogStreams200ResponseInnerOneOf3
	isSet bool
}

func (v NullableGetLogStreams200ResponseInnerOneOf3) Get() *GetLogStreams200ResponseInnerOneOf3 {
	return v.value
}

func (v *NullableGetLogStreams200ResponseInnerOneOf3) Set(val *GetLogStreams200ResponseInnerOneOf3) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLogStreams200ResponseInnerOneOf3) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLogStreams200ResponseInnerOneOf3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLogStreams200ResponseInnerOneOf3(val *GetLogStreams200ResponseInnerOneOf3) *NullableGetLogStreams200ResponseInnerOneOf3 {
	return &NullableGetLogStreams200ResponseInnerOneOf3{value: val, isSet: true}
}

func (v NullableGetLogStreams200ResponseInnerOneOf3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLogStreams200ResponseInnerOneOf3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
