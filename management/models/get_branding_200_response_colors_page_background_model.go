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

// GetBranding200ResponseColorsPageBackground - Page Background Color or Gradient. Property contains either <code>null</code> to unset, a solid color as a string value <code>#FFFFFF</code>, or a gradient as an object.  <pre><code> {   type: 'linear-gradient',   start: '#FFFFFF',   end: '#000000',   angle_deg: 35 } </code></pre>
type GetBranding200ResponseColorsPageBackground struct {
	MapmapOfStringAny *map[string]interface{}
	String            *string
}

// map[string]interface{}AsGetBranding200ResponseColorsPageBackground is a convenience function that returns map[string]interface{} wrapped in GetBranding200ResponseColorsPageBackground
func MapmapOfStringAnyAsGetBranding200ResponseColorsPageBackground(v *map[string]interface{}) GetBranding200ResponseColorsPageBackground {
	return GetBranding200ResponseColorsPageBackground{
		MapmapOfStringAny: v,
	}
}

// stringAsGetBranding200ResponseColorsPageBackground is a convenience function that returns string wrapped in GetBranding200ResponseColorsPageBackground
func StringAsGetBranding200ResponseColorsPageBackground(v *string) GetBranding200ResponseColorsPageBackground {
	return GetBranding200ResponseColorsPageBackground{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetBranding200ResponseColorsPageBackground) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(GetBranding200ResponseColorsPageBackground)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetBranding200ResponseColorsPageBackground)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetBranding200ResponseColorsPageBackground) MarshalJSON() ([]byte, error) {
	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetBranding200ResponseColorsPageBackground) GetActualInstance() interface{} {
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

type NullableGetBranding200ResponseColorsPageBackground struct {
	value *GetBranding200ResponseColorsPageBackground
	isSet bool
}

func (v NullableGetBranding200ResponseColorsPageBackground) Get() *GetBranding200ResponseColorsPageBackground {
	return v.value
}

func (v *NullableGetBranding200ResponseColorsPageBackground) Set(val *GetBranding200ResponseColorsPageBackground) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBranding200ResponseColorsPageBackground) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBranding200ResponseColorsPageBackground) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBranding200ResponseColorsPageBackground(val *GetBranding200ResponseColorsPageBackground) *NullableGetBranding200ResponseColorsPageBackground {
	return &NullableGetBranding200ResponseColorsPageBackground{value: val, isSet: true}
}

func (v NullableGetBranding200ResponseColorsPageBackground) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBranding200ResponseColorsPageBackground) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
