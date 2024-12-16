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

// GetLogStreams200ResponseInnerOneOf6Sink struct for GetLogStreams200ResponseInnerOneOf6Sink
type GetLogStreams200ResponseInnerOneOf6Sink struct {
	// Segment write key
	SegmentWriteKey string `json:"segmentWriteKey"`
}

type _GetLogStreams200ResponseInnerOneOf6Sink GetLogStreams200ResponseInnerOneOf6Sink

// GetSegmentWriteKey returns the SegmentWriteKey field value
func (o *GetLogStreams200ResponseInnerOneOf6Sink) GetSegmentWriteKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SegmentWriteKey
}

// GetSegmentWriteKeyOk returns a tuple with the SegmentWriteKey field value
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf6Sink) GetSegmentWriteKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SegmentWriteKey, true
}

// SetSegmentWriteKey sets field value
func (o *GetLogStreams200ResponseInnerOneOf6Sink) SetSegmentWriteKey(v string) {
	o.SegmentWriteKey = v
}

func (o GetLogStreams200ResponseInnerOneOf6Sink) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLogStreams200ResponseInnerOneOf6Sink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["segmentWriteKey"] = o.SegmentWriteKey
	return toSerialize, nil
}

func (o *GetLogStreams200ResponseInnerOneOf6Sink) UnmarshalJSON(data []byte) (err error) {
	varGetLogStreams200ResponseInnerOneOf6Sink := _GetLogStreams200ResponseInnerOneOf6Sink{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetLogStreams200ResponseInnerOneOf6Sink)

	if err != nil {
		return err
	}

	*o = GetLogStreams200ResponseInnerOneOf6Sink(varGetLogStreams200ResponseInnerOneOf6Sink)

	return err
}

type NullableGetLogStreams200ResponseInnerOneOf6Sink struct {
	value *GetLogStreams200ResponseInnerOneOf6Sink
	isSet bool
}

func (v NullableGetLogStreams200ResponseInnerOneOf6Sink) Get() *GetLogStreams200ResponseInnerOneOf6Sink {
	return v.value
}

func (v *NullableGetLogStreams200ResponseInnerOneOf6Sink) Set(val *GetLogStreams200ResponseInnerOneOf6Sink) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLogStreams200ResponseInnerOneOf6Sink) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLogStreams200ResponseInnerOneOf6Sink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLogStreams200ResponseInnerOneOf6Sink(val *GetLogStreams200ResponseInnerOneOf6Sink) *NullableGetLogStreams200ResponseInnerOneOf6Sink {
	return &NullableGetLogStreams200ResponseInnerOneOf6Sink{value: val, isSet: true}
}

func (v NullableGetLogStreams200ResponseInnerOneOf6Sink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLogStreams200ResponseInnerOneOf6Sink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
