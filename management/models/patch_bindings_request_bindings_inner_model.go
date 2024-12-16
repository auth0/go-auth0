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

// PatchBindingsRequestBindingsInner - struct for PatchBindingsRequestBindingsInner
type PatchBindingsRequestBindingsInner struct {
	PatchBindingsRequestBindingsInnerOneOf *PatchBindingsRequestBindingsInnerOneOf
}

// PatchBindingsRequestBindingsInnerOneOfAsPatchBindingsRequestBindingsInner is a convenience function that returns PatchBindingsRequestBindingsInnerOneOf wrapped in PatchBindingsRequestBindingsInner
func PatchBindingsRequestBindingsInnerOneOfAsPatchBindingsRequestBindingsInner(v *PatchBindingsRequestBindingsInnerOneOf) PatchBindingsRequestBindingsInner {
	return PatchBindingsRequestBindingsInner{
		PatchBindingsRequestBindingsInnerOneOf: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PatchBindingsRequestBindingsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PatchBindingsRequestBindingsInnerOneOf
	err = newStrictDecoder(data).Decode(&dst.PatchBindingsRequestBindingsInnerOneOf)
	if err == nil {
		jsonPatchBindingsRequestBindingsInnerOneOf, _ := json.Marshal(dst.PatchBindingsRequestBindingsInnerOneOf)
		if string(jsonPatchBindingsRequestBindingsInnerOneOf) == "{}" { // empty struct
			dst.PatchBindingsRequestBindingsInnerOneOf = nil
		} else {
			if err = validator.Validate(dst.PatchBindingsRequestBindingsInnerOneOf); err != nil {
				dst.PatchBindingsRequestBindingsInnerOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.PatchBindingsRequestBindingsInnerOneOf = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PatchBindingsRequestBindingsInnerOneOf = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PatchBindingsRequestBindingsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PatchBindingsRequestBindingsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PatchBindingsRequestBindingsInner) MarshalJSON() ([]byte, error) {
	if src.PatchBindingsRequestBindingsInnerOneOf != nil {
		return json.Marshal(&src.PatchBindingsRequestBindingsInnerOneOf)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PatchBindingsRequestBindingsInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.PatchBindingsRequestBindingsInnerOneOf != nil {
		return obj.PatchBindingsRequestBindingsInnerOneOf
	}

	// all schemas are nil
	return nil
}

type NullablePatchBindingsRequestBindingsInner struct {
	value *PatchBindingsRequestBindingsInner
	isSet bool
}

func (v NullablePatchBindingsRequestBindingsInner) Get() *PatchBindingsRequestBindingsInner {
	return v.value
}

func (v *NullablePatchBindingsRequestBindingsInner) Set(val *PatchBindingsRequestBindingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchBindingsRequestBindingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchBindingsRequestBindingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchBindingsRequestBindingsInner(val *PatchBindingsRequestBindingsInner) *NullablePatchBindingsRequestBindingsInner {
	return &NullablePatchBindingsRequestBindingsInner{value: val, isSet: true}
}

func (v NullablePatchBindingsRequestBindingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchBindingsRequestBindingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
