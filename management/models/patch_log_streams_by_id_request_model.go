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

// PatchLogStreamsByIdRequest struct for PatchLogStreamsByIdRequest
type PatchLogStreamsByIdRequest struct {
	// log stream name
	Name   *string                                   `json:"name,omitempty"`
	Status *GetLogStreams200ResponseInnerOneOfStatus `json:"status,omitempty"`
	// True for priority log streams, false for non-priority
	IsPriority *bool `json:"isPriority,omitempty"`
	// Only logs events matching these filters will be delivered by the stream. If omitted or empty, all events will be delivered.
	Filters []GetLogStreams200ResponseInnerOneOfFiltersInner `json:"filters,omitempty"`
	Sink    *PatchLogStreamsByIdRequestSink                  `json:"sink,omitempty"`
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchLogStreamsByIdRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchLogStreamsByIdRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchLogStreamsByIdRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchLogStreamsByIdRequest) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchLogStreamsByIdRequest) GetStatus() GetLogStreams200ResponseInnerOneOfStatus {
	if o == nil || IsNil(o.Status) {
		var ret GetLogStreams200ResponseInnerOneOfStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchLogStreamsByIdRequest) GetStatusOk() (*GetLogStreams200ResponseInnerOneOfStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchLogStreamsByIdRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GetLogStreams200ResponseInnerOneOfStatus and assigns it to the Status field.
func (o *PatchLogStreamsByIdRequest) SetStatus(v GetLogStreams200ResponseInnerOneOfStatus) {
	o.Status = &v
}

// GetIsPriority returns the IsPriority field value if set, zero value otherwise.
func (o *PatchLogStreamsByIdRequest) GetIsPriority() bool {
	if o == nil || IsNil(o.IsPriority) {
		var ret bool
		return ret
	}
	return *o.IsPriority
}

// GetIsPriorityOk returns a tuple with the IsPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchLogStreamsByIdRequest) GetIsPriorityOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPriority) {
		return nil, false
	}
	return o.IsPriority, true
}

// HasIsPriority returns a boolean if a field has been set.
func (o *PatchLogStreamsByIdRequest) HasIsPriority() bool {
	if o != nil && !IsNil(o.IsPriority) {
		return true
	}

	return false
}

// SetIsPriority gets a reference to the given bool and assigns it to the IsPriority field.
func (o *PatchLogStreamsByIdRequest) SetIsPriority(v bool) {
	o.IsPriority = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *PatchLogStreamsByIdRequest) GetFilters() []GetLogStreams200ResponseInnerOneOfFiltersInner {
	if o == nil || IsNil(o.Filters) {
		var ret []GetLogStreams200ResponseInnerOneOfFiltersInner
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchLogStreamsByIdRequest) GetFiltersOk() ([]GetLogStreams200ResponseInnerOneOfFiltersInner, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *PatchLogStreamsByIdRequest) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []GetLogStreams200ResponseInnerOneOfFiltersInner and assigns it to the Filters field.
func (o *PatchLogStreamsByIdRequest) SetFilters(v []GetLogStreams200ResponseInnerOneOfFiltersInner) {
	o.Filters = v
}

// GetSink returns the Sink field value if set, zero value otherwise.
func (o *PatchLogStreamsByIdRequest) GetSink() PatchLogStreamsByIdRequestSink {
	if o == nil || IsNil(o.Sink) {
		var ret PatchLogStreamsByIdRequestSink
		return ret
	}
	return *o.Sink
}

// GetSinkOk returns a tuple with the Sink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchLogStreamsByIdRequest) GetSinkOk() (*PatchLogStreamsByIdRequestSink, bool) {
	if o == nil || IsNil(o.Sink) {
		return nil, false
	}
	return o.Sink, true
}

// HasSink returns a boolean if a field has been set.
func (o *PatchLogStreamsByIdRequest) HasSink() bool {
	if o != nil && !IsNil(o.Sink) {
		return true
	}

	return false
}

// SetSink gets a reference to the given PatchLogStreamsByIdRequestSink and assigns it to the Sink field.
func (o *PatchLogStreamsByIdRequest) SetSink(v PatchLogStreamsByIdRequestSink) {
	o.Sink = &v
}

func (o PatchLogStreamsByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchLogStreamsByIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.IsPriority) {
		toSerialize["isPriority"] = o.IsPriority
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.Sink) {
		toSerialize["sink"] = o.Sink
	}
	return toSerialize, nil
}

type NullablePatchLogStreamsByIdRequest struct {
	value *PatchLogStreamsByIdRequest
	isSet bool
}

func (v NullablePatchLogStreamsByIdRequest) Get() *PatchLogStreamsByIdRequest {
	return v.value
}

func (v *NullablePatchLogStreamsByIdRequest) Set(val *PatchLogStreamsByIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchLogStreamsByIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchLogStreamsByIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchLogStreamsByIdRequest(val *PatchLogStreamsByIdRequest) *NullablePatchLogStreamsByIdRequest {
	return &NullablePatchLogStreamsByIdRequest{value: val, isSet: true}
}

func (v NullablePatchLogStreamsByIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchLogStreamsByIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
