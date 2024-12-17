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

// GetLogStreams200ResponseInnerOneOf7 struct for GetLogStreams200ResponseInnerOneOf7
type GetLogStreams200ResponseInnerOneOf7 struct {
	// The id of the log stream
	Id *string `json:"id,omitempty"`
	// log stream name
	Name   *string                                   `json:"name,omitempty"`
	Status *GetLogStreams200ResponseInnerOneOfStatus `json:"status,omitempty"`
	Type   *GetLogStreams200ResponseInnerOneOf7Type  `json:"type,omitempty"`
	// Only logs events matching these filters will be delivered by the stream. If omitted or empty, all events will be delivered.
	Filters              []GetLogStreams200ResponseInnerOneOfFiltersInner `json:"filters,omitempty"`
	Sink                 *GetLogStreams200ResponseInnerOneOf7Sink         `json:"sink,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetLogStreams200ResponseInnerOneOf7 GetLogStreams200ResponseInnerOneOf7

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetLogStreams200ResponseInnerOneOf7) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf7) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetLogStreams200ResponseInnerOneOf7) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetLogStreams200ResponseInnerOneOf7) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetLogStreams200ResponseInnerOneOf7) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf7) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetLogStreams200ResponseInnerOneOf7) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetLogStreams200ResponseInnerOneOf7) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetLogStreams200ResponseInnerOneOf7) GetStatus() GetLogStreams200ResponseInnerOneOfStatus {
	if o == nil || IsNil(o.Status) {
		var ret GetLogStreams200ResponseInnerOneOfStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf7) GetStatusOk() (*GetLogStreams200ResponseInnerOneOfStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetLogStreams200ResponseInnerOneOf7) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GetLogStreams200ResponseInnerOneOfStatus and assigns it to the Status field.
func (o *GetLogStreams200ResponseInnerOneOf7) SetStatus(v GetLogStreams200ResponseInnerOneOfStatus) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetLogStreams200ResponseInnerOneOf7) GetType() GetLogStreams200ResponseInnerOneOf7Type {
	if o == nil || IsNil(o.Type) {
		var ret GetLogStreams200ResponseInnerOneOf7Type
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf7) GetTypeOk() (*GetLogStreams200ResponseInnerOneOf7Type, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetLogStreams200ResponseInnerOneOf7) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given GetLogStreams200ResponseInnerOneOf7Type and assigns it to the Type field.
func (o *GetLogStreams200ResponseInnerOneOf7) SetType(v GetLogStreams200ResponseInnerOneOf7Type) {
	o.Type = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *GetLogStreams200ResponseInnerOneOf7) GetFilters() []GetLogStreams200ResponseInnerOneOfFiltersInner {
	if o == nil || IsNil(o.Filters) {
		var ret []GetLogStreams200ResponseInnerOneOfFiltersInner
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf7) GetFiltersOk() ([]GetLogStreams200ResponseInnerOneOfFiltersInner, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *GetLogStreams200ResponseInnerOneOf7) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []GetLogStreams200ResponseInnerOneOfFiltersInner and assigns it to the Filters field.
func (o *GetLogStreams200ResponseInnerOneOf7) SetFilters(v []GetLogStreams200ResponseInnerOneOfFiltersInner) {
	o.Filters = v
}

// GetSink returns the Sink field value if set, zero value otherwise.
func (o *GetLogStreams200ResponseInnerOneOf7) GetSink() GetLogStreams200ResponseInnerOneOf7Sink {
	if o == nil || IsNil(o.Sink) {
		var ret GetLogStreams200ResponseInnerOneOf7Sink
		return ret
	}
	return *o.Sink
}

// GetSinkOk returns a tuple with the Sink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf7) GetSinkOk() (*GetLogStreams200ResponseInnerOneOf7Sink, bool) {
	if o == nil || IsNil(o.Sink) {
		return nil, false
	}
	return o.Sink, true
}

// HasSink returns a boolean if a field has been set.
func (o *GetLogStreams200ResponseInnerOneOf7) HasSink() bool {
	if o != nil && !IsNil(o.Sink) {
		return true
	}

	return false
}

// SetSink gets a reference to the given GetLogStreams200ResponseInnerOneOf7Sink and assigns it to the Sink field.
func (o *GetLogStreams200ResponseInnerOneOf7) SetSink(v GetLogStreams200ResponseInnerOneOf7Sink) {
	o.Sink = &v
}

func (o GetLogStreams200ResponseInnerOneOf7) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLogStreams200ResponseInnerOneOf7) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.Sink) {
		toSerialize["sink"] = o.Sink
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetLogStreams200ResponseInnerOneOf7) UnmarshalJSON(data []byte) (err error) {
	varGetLogStreams200ResponseInnerOneOf7 := _GetLogStreams200ResponseInnerOneOf7{}

	err = json.Unmarshal(data, &varGetLogStreams200ResponseInnerOneOf7)

	if err != nil {
		return err
	}

	*o = GetLogStreams200ResponseInnerOneOf7(varGetLogStreams200ResponseInnerOneOf7)

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

type NullableGetLogStreams200ResponseInnerOneOf7 struct {
	value *GetLogStreams200ResponseInnerOneOf7
	isSet bool
}

func (v NullableGetLogStreams200ResponseInnerOneOf7) Get() *GetLogStreams200ResponseInnerOneOf7 {
	return v.value
}

func (v *NullableGetLogStreams200ResponseInnerOneOf7) Set(val *GetLogStreams200ResponseInnerOneOf7) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLogStreams200ResponseInnerOneOf7) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLogStreams200ResponseInnerOneOf7) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLogStreams200ResponseInnerOneOf7(val *GetLogStreams200ResponseInnerOneOf7) *NullableGetLogStreams200ResponseInnerOneOf7 {
	return &NullableGetLogStreams200ResponseInnerOneOf7{value: val, isSet: true}
}

func (v NullableGetLogStreams200ResponseInnerOneOf7) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLogStreams200ResponseInnerOneOf7) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
