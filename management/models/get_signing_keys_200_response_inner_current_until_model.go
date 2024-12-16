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

	"gopkg.in/validator.v2"
)

// GetSigningKeys200ResponseInnerCurrentUntil - struct for GetSigningKeys200ResponseInnerCurrentUntil
type GetSigningKeys200ResponseInnerCurrentUntil struct {
	MapmapOfStringAny *map[string]interface{}
	String            *string
}

// map[string]interface{}AsGetSigningKeys200ResponseInnerCurrentUntil is a convenience function that returns map[string]interface{} wrapped in GetSigningKeys200ResponseInnerCurrentUntil
func MapmapOfStringAnyAsGetSigningKeys200ResponseInnerCurrentUntil(v *map[string]interface{}) GetSigningKeys200ResponseInnerCurrentUntil {
	return GetSigningKeys200ResponseInnerCurrentUntil{
		MapmapOfStringAny: v,
	}
}

// stringAsGetSigningKeys200ResponseInnerCurrentUntil is a convenience function that returns string wrapped in GetSigningKeys200ResponseInnerCurrentUntil
func StringAsGetSigningKeys200ResponseInnerCurrentUntil(v *string) GetSigningKeys200ResponseInnerCurrentUntil {
	return GetSigningKeys200ResponseInnerCurrentUntil{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetSigningKeys200ResponseInnerCurrentUntil) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MapmapOfStringAny = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetSigningKeys200ResponseInnerCurrentUntil)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetSigningKeys200ResponseInnerCurrentUntil)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetSigningKeys200ResponseInnerCurrentUntil) MarshalJSON() ([]byte, error) {
	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetSigningKeys200ResponseInnerCurrentUntil) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableGetSigningKeys200ResponseInnerCurrentUntil struct {
	value *GetSigningKeys200ResponseInnerCurrentUntil
	isSet bool
}

func (v NullableGetSigningKeys200ResponseInnerCurrentUntil) Get() *GetSigningKeys200ResponseInnerCurrentUntil {
	return v.value
}

func (v *NullableGetSigningKeys200ResponseInnerCurrentUntil) Set(val *GetSigningKeys200ResponseInnerCurrentUntil) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSigningKeys200ResponseInnerCurrentUntil) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSigningKeys200ResponseInnerCurrentUntil) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSigningKeys200ResponseInnerCurrentUntil(val *GetSigningKeys200ResponseInnerCurrentUntil) *NullableGetSigningKeys200ResponseInnerCurrentUntil {
	return &NullableGetSigningKeys200ResponseInnerCurrentUntil{value: val, isSet: true}
}

func (v NullableGetSigningKeys200ResponseInnerCurrentUntil) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSigningKeys200ResponseInnerCurrentUntil) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
