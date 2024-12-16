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

// GetRoleUser200Response - struct for GetRoleUser200Response
type GetRoleUser200Response struct {
	GetRoleUser200ResponseOneOf             *GetRoleUser200ResponseOneOf
	GetRoleUser200ResponseOneOf1            *GetRoleUser200ResponseOneOf1
	ArrayOfGetRoleUser200ResponseOneOfInner *[]GetRoleUser200ResponseOneOfInner
}

// GetRoleUser200ResponseOneOfAsGetRoleUser200Response is a convenience function that returns GetRoleUser200ResponseOneOf wrapped in GetRoleUser200Response
func GetRoleUser200ResponseOneOfAsGetRoleUser200Response(v *GetRoleUser200ResponseOneOf) GetRoleUser200Response {
	return GetRoleUser200Response{
		GetRoleUser200ResponseOneOf: v,
	}
}

// GetRoleUser200ResponseOneOf1AsGetRoleUser200Response is a convenience function that returns GetRoleUser200ResponseOneOf1 wrapped in GetRoleUser200Response
func GetRoleUser200ResponseOneOf1AsGetRoleUser200Response(v *GetRoleUser200ResponseOneOf1) GetRoleUser200Response {
	return GetRoleUser200Response{
		GetRoleUser200ResponseOneOf1: v,
	}
}

// []GetRoleUser200ResponseOneOfInnerAsGetRoleUser200Response is a convenience function that returns []GetRoleUser200ResponseOneOfInner wrapped in GetRoleUser200Response
func ArrayOfGetRoleUser200ResponseOneOfInnerAsGetRoleUser200Response(v *[]GetRoleUser200ResponseOneOfInner) GetRoleUser200Response {
	return GetRoleUser200Response{
		ArrayOfGetRoleUser200ResponseOneOfInner: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetRoleUser200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetRoleUser200ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.GetRoleUser200ResponseOneOf)
	if err == nil {
		jsonGetRoleUser200ResponseOneOf, _ := json.Marshal(dst.GetRoleUser200ResponseOneOf)
		if string(jsonGetRoleUser200ResponseOneOf) == "{}" { // empty struct
			dst.GetRoleUser200ResponseOneOf = nil
		} else {
			if err = validator.Validate(dst.GetRoleUser200ResponseOneOf); err != nil {
				dst.GetRoleUser200ResponseOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.GetRoleUser200ResponseOneOf = nil
	}

	// try to unmarshal data into GetRoleUser200ResponseOneOf1
	err = newStrictDecoder(data).Decode(&dst.GetRoleUser200ResponseOneOf1)
	if err == nil {
		jsonGetRoleUser200ResponseOneOf1, _ := json.Marshal(dst.GetRoleUser200ResponseOneOf1)
		if string(jsonGetRoleUser200ResponseOneOf1) == "{}" { // empty struct
			dst.GetRoleUser200ResponseOneOf1 = nil
		} else {
			if err = validator.Validate(dst.GetRoleUser200ResponseOneOf1); err != nil {
				dst.GetRoleUser200ResponseOneOf1 = nil
			} else {
				match++
			}
		}
	} else {
		dst.GetRoleUser200ResponseOneOf1 = nil
	}

	// try to unmarshal data into ArrayOfGetRoleUser200ResponseOneOfInner
	err = newStrictDecoder(data).Decode(&dst.ArrayOfGetRoleUser200ResponseOneOfInner)
	if err == nil {
		jsonArrayOfGetRoleUser200ResponseOneOfInner, _ := json.Marshal(dst.ArrayOfGetRoleUser200ResponseOneOfInner)
		if string(jsonArrayOfGetRoleUser200ResponseOneOfInner) == "{}" { // empty struct
			dst.ArrayOfGetRoleUser200ResponseOneOfInner = nil
		} else {
			if err = validator.Validate(dst.ArrayOfGetRoleUser200ResponseOneOfInner); err != nil {
				dst.ArrayOfGetRoleUser200ResponseOneOfInner = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfGetRoleUser200ResponseOneOfInner = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetRoleUser200ResponseOneOf = nil
		dst.GetRoleUser200ResponseOneOf1 = nil
		dst.ArrayOfGetRoleUser200ResponseOneOfInner = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetRoleUser200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetRoleUser200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetRoleUser200Response) MarshalJSON() ([]byte, error) {
	if src.GetRoleUser200ResponseOneOf != nil {
		return json.Marshal(&src.GetRoleUser200ResponseOneOf)
	}

	if src.GetRoleUser200ResponseOneOf1 != nil {
		return json.Marshal(&src.GetRoleUser200ResponseOneOf1)
	}

	if src.ArrayOfGetRoleUser200ResponseOneOfInner != nil {
		return json.Marshal(&src.ArrayOfGetRoleUser200ResponseOneOfInner)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetRoleUser200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetRoleUser200ResponseOneOf != nil {
		return obj.GetRoleUser200ResponseOneOf
	}

	if obj.GetRoleUser200ResponseOneOf1 != nil {
		return obj.GetRoleUser200ResponseOneOf1
	}

	if obj.ArrayOfGetRoleUser200ResponseOneOfInner != nil {
		return obj.ArrayOfGetRoleUser200ResponseOneOfInner
	}

	// all schemas are nil
	return nil
}

type NullableGetRoleUser200Response struct {
	value *GetRoleUser200Response
	isSet bool
}

func (v NullableGetRoleUser200Response) Get() *GetRoleUser200Response {
	return v.value
}

func (v *NullableGetRoleUser200Response) Set(val *GetRoleUser200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRoleUser200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRoleUser200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRoleUser200Response(val *GetRoleUser200Response) *NullableGetRoleUser200Response {
	return &NullableGetRoleUser200Response{value: val, isSet: true}
}

func (v NullableGetRoleUser200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRoleUser200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
