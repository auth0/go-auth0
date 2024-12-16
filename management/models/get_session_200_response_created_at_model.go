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
	"time"

	"gopkg.in/validator.v2"
)

// GetSession200ResponseCreatedAt - struct for GetSession200ResponseCreatedAt
type GetSession200ResponseCreatedAt struct {
	MapmapOfStringAny *map[string]interface{}
	TimeTime          *time.Time
}

// map[string]interface{}AsGetSession200ResponseCreatedAt is a convenience function that returns map[string]interface{} wrapped in GetSession200ResponseCreatedAt
func MapmapOfStringAnyAsGetSession200ResponseCreatedAt(v *map[string]interface{}) GetSession200ResponseCreatedAt {
	return GetSession200ResponseCreatedAt{
		MapmapOfStringAny: v,
	}
}

// time.TimeAsGetSession200ResponseCreatedAt is a convenience function that returns time.Time wrapped in GetSession200ResponseCreatedAt
func TimeTimeAsGetSession200ResponseCreatedAt(v *time.Time) GetSession200ResponseCreatedAt {
	return GetSession200ResponseCreatedAt{
		TimeTime: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetSession200ResponseCreatedAt) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into MapmapOfStringAny
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringAny)
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			if err = validator.Validate(dst.MapmapOfStringAny); err != nil {
				dst.MapmapOfStringAny = nil
			} else {
				match++
			}
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	// try to unmarshal data into TimeTime
	err = newStrictDecoder(data).Decode(&dst.TimeTime)
	if err == nil {
		jsonTimeTime, _ := json.Marshal(dst.TimeTime)
		if string(jsonTimeTime) == "{}" { // empty struct
			dst.TimeTime = nil
		} else {
			if err = validator.Validate(dst.TimeTime); err != nil {
				dst.TimeTime = nil
			} else {
				match++
			}
		}
	} else {
		dst.TimeTime = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MapmapOfStringAny = nil
		dst.TimeTime = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetSession200ResponseCreatedAt)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetSession200ResponseCreatedAt)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetSession200ResponseCreatedAt) MarshalJSON() ([]byte, error) {
	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	if src.TimeTime != nil {
		return json.Marshal(&src.TimeTime)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetSession200ResponseCreatedAt) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	if obj.TimeTime != nil {
		return obj.TimeTime
	}

	// all schemas are nil
	return nil
}

type NullableGetSession200ResponseCreatedAt struct {
	value *GetSession200ResponseCreatedAt
	isSet bool
}

func (v NullableGetSession200ResponseCreatedAt) Get() *GetSession200ResponseCreatedAt {
	return v.value
}

func (v *NullableGetSession200ResponseCreatedAt) Set(val *GetSession200ResponseCreatedAt) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSession200ResponseCreatedAt) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSession200ResponseCreatedAt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSession200ResponseCreatedAt(val *GetSession200ResponseCreatedAt) *NullableGetSession200ResponseCreatedAt {
	return &NullableGetSession200ResponseCreatedAt{value: val, isSet: true}
}

func (v NullableGetSession200ResponseCreatedAt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSession200ResponseCreatedAt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
