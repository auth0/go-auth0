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

// GetLogStreams200ResponseInnerOneOf5Type the model 'GetLogStreams200ResponseInnerOneOf5Type'
type GetLogStreams200ResponseInnerOneOf5Type string

// List of get_log_streams_200_response_inner_oneOf_5_type
const (
	GETLOGSTREAMS200RESPONSEINNERONEOF5TYPE_SUMO GetLogStreams200ResponseInnerOneOf5Type = "sumo"
)

// All allowed values of GetLogStreams200ResponseInnerOneOf5Type enum
var AllowedGetLogStreams200ResponseInnerOneOf5TypeEnumValues = []GetLogStreams200ResponseInnerOneOf5Type{
	"sumo",
}

func (v *GetLogStreams200ResponseInnerOneOf5Type) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetLogStreams200ResponseInnerOneOf5Type(value)
	for _, existing := range AllowedGetLogStreams200ResponseInnerOneOf5TypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetLogStreams200ResponseInnerOneOf5Type", value)
}

// NewGetLogStreams200ResponseInnerOneOf5TypeFromValue returns a pointer to a valid GetLogStreams200ResponseInnerOneOf5Type
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetLogStreams200ResponseInnerOneOf5TypeFromValue(v string) (*GetLogStreams200ResponseInnerOneOf5Type, error) {
	ev := GetLogStreams200ResponseInnerOneOf5Type(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetLogStreams200ResponseInnerOneOf5Type: valid values are %v", v, AllowedGetLogStreams200ResponseInnerOneOf5TypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetLogStreams200ResponseInnerOneOf5Type) IsValid() bool {
	for _, existing := range AllowedGetLogStreams200ResponseInnerOneOf5TypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to get_log_streams_200_response_inner_oneOf_5_type value
func (v GetLogStreams200ResponseInnerOneOf5Type) Ptr() *GetLogStreams200ResponseInnerOneOf5Type {
	return &v
}

type NullableGetLogStreams200ResponseInnerOneOf5Type struct {
	value *GetLogStreams200ResponseInnerOneOf5Type
	isSet bool
}

func (v NullableGetLogStreams200ResponseInnerOneOf5Type) Get() *GetLogStreams200ResponseInnerOneOf5Type {
	return v.value
}

func (v *NullableGetLogStreams200ResponseInnerOneOf5Type) Set(val *GetLogStreams200ResponseInnerOneOf5Type) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLogStreams200ResponseInnerOneOf5Type) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLogStreams200ResponseInnerOneOf5Type) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLogStreams200ResponseInnerOneOf5Type(val *GetLogStreams200ResponseInnerOneOf5Type) *NullableGetLogStreams200ResponseInnerOneOf5Type {
	return &NullableGetLogStreams200ResponseInnerOneOf5Type{value: val, isSet: true}
}

func (v NullableGetLogStreams200ResponseInnerOneOf5Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLogStreams200ResponseInnerOneOf5Type) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
