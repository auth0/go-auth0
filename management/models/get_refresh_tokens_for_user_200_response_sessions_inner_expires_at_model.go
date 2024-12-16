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

// GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt - struct for GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt
type GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt struct {
	MapmapOfStringAny *map[string]interface{}
	TimeTime          *time.Time
}

// map[string]interface{}AsGetRefreshTokensForUser200ResponseSessionsInnerExpiresAt is a convenience function that returns map[string]interface{} wrapped in GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt
func MapmapOfStringAnyAsGetRefreshTokensForUser200ResponseSessionsInnerExpiresAt(v *map[string]interface{}) GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt {
	return GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt{
		MapmapOfStringAny: v,
	}
}

// time.TimeAsGetRefreshTokensForUser200ResponseSessionsInnerExpiresAt is a convenience function that returns time.Time wrapped in GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt
func TimeTimeAsGetRefreshTokensForUser200ResponseSessionsInnerExpiresAt(v *time.Time) GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt {
	return GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt{
		TimeTime: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt) UnmarshalJSON(data []byte) error {
	var err error
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

		return fmt.Errorf("data matches more than one schema in oneOf(GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt) MarshalJSON() ([]byte, error) {
	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	if src.TimeTime != nil {
		return json.Marshal(&src.TimeTime)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt) GetActualInstance() interface{} {
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

type NullableGetRefreshTokensForUser200ResponseSessionsInnerExpiresAt struct {
	value *GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt
	isSet bool
}

func (v NullableGetRefreshTokensForUser200ResponseSessionsInnerExpiresAt) Get() *GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt {
	return v.value
}

func (v *NullableGetRefreshTokensForUser200ResponseSessionsInnerExpiresAt) Set(val *GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRefreshTokensForUser200ResponseSessionsInnerExpiresAt) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRefreshTokensForUser200ResponseSessionsInnerExpiresAt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRefreshTokensForUser200ResponseSessionsInnerExpiresAt(val *GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt) *NullableGetRefreshTokensForUser200ResponseSessionsInnerExpiresAt {
	return &NullableGetRefreshTokensForUser200ResponseSessionsInnerExpiresAt{value: val, isSet: true}
}

func (v NullableGetRefreshTokensForUser200ResponseSessionsInnerExpiresAt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRefreshTokensForUser200ResponseSessionsInnerExpiresAt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
