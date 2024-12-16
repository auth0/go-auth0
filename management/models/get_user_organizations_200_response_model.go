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

// GetUserOrganizations200Response - struct for GetUserOrganizations200Response
type GetUserOrganizations200Response struct {
	GetOrganizations200ResponseOneOf             *GetOrganizations200ResponseOneOf
	ArrayOfGetOrganizations200ResponseOneOfInner *[]GetOrganizations200ResponseOneOfInner
}

// GetOrganizations200ResponseOneOfAsGetUserOrganizations200Response is a convenience function that returns GetOrganizations200ResponseOneOf wrapped in GetUserOrganizations200Response
func GetOrganizations200ResponseOneOfAsGetUserOrganizations200Response(v *GetOrganizations200ResponseOneOf) GetUserOrganizations200Response {
	return GetUserOrganizations200Response{
		GetOrganizations200ResponseOneOf: v,
	}
}

// []GetOrganizations200ResponseOneOfInnerAsGetUserOrganizations200Response is a convenience function that returns []GetOrganizations200ResponseOneOfInner wrapped in GetUserOrganizations200Response
func ArrayOfGetOrganizations200ResponseOneOfInnerAsGetUserOrganizations200Response(v *[]GetOrganizations200ResponseOneOfInner) GetUserOrganizations200Response {
	return GetUserOrganizations200Response{
		ArrayOfGetOrganizations200ResponseOneOfInner: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetUserOrganizations200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetOrganizations200ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.GetOrganizations200ResponseOneOf)
	if err == nil {
		jsonGetOrganizations200ResponseOneOf, _ := json.Marshal(dst.GetOrganizations200ResponseOneOf)
		if string(jsonGetOrganizations200ResponseOneOf) == "{}" { // empty struct
			dst.GetOrganizations200ResponseOneOf = nil
		} else {
			if err = validator.Validate(dst.GetOrganizations200ResponseOneOf); err != nil {
				dst.GetOrganizations200ResponseOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.GetOrganizations200ResponseOneOf = nil
	}

	// try to unmarshal data into ArrayOfGetOrganizations200ResponseOneOfInner
	err = newStrictDecoder(data).Decode(&dst.ArrayOfGetOrganizations200ResponseOneOfInner)
	if err == nil {
		jsonArrayOfGetOrganizations200ResponseOneOfInner, _ := json.Marshal(dst.ArrayOfGetOrganizations200ResponseOneOfInner)
		if string(jsonArrayOfGetOrganizations200ResponseOneOfInner) == "{}" { // empty struct
			dst.ArrayOfGetOrganizations200ResponseOneOfInner = nil
		} else {
			if err = validator.Validate(dst.ArrayOfGetOrganizations200ResponseOneOfInner); err != nil {
				dst.ArrayOfGetOrganizations200ResponseOneOfInner = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfGetOrganizations200ResponseOneOfInner = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetOrganizations200ResponseOneOf = nil
		dst.ArrayOfGetOrganizations200ResponseOneOfInner = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetUserOrganizations200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetUserOrganizations200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetUserOrganizations200Response) MarshalJSON() ([]byte, error) {
	if src.GetOrganizations200ResponseOneOf != nil {
		return json.Marshal(&src.GetOrganizations200ResponseOneOf)
	}

	if src.ArrayOfGetOrganizations200ResponseOneOfInner != nil {
		return json.Marshal(&src.ArrayOfGetOrganizations200ResponseOneOfInner)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetUserOrganizations200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetOrganizations200ResponseOneOf != nil {
		return obj.GetOrganizations200ResponseOneOf
	}

	if obj.ArrayOfGetOrganizations200ResponseOneOfInner != nil {
		return obj.ArrayOfGetOrganizations200ResponseOneOfInner
	}

	// all schemas are nil
	return nil
}

type NullableGetUserOrganizations200Response struct {
	value *GetUserOrganizations200Response
	isSet bool
}

func (v NullableGetUserOrganizations200Response) Get() *GetUserOrganizations200Response {
	return v.value
}

func (v *NullableGetUserOrganizations200Response) Set(val *GetUserOrganizations200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserOrganizations200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserOrganizations200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserOrganizations200Response(val *GetUserOrganizations200Response) *NullableGetUserOrganizations200Response {
	return &NullableGetUserOrganizations200Response{value: val, isSet: true}
}

func (v NullableGetUserOrganizations200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserOrganizations200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
