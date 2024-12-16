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

// GetClientGrants200Response - struct for GetClientGrants200Response
type GetClientGrants200Response struct {
	GetClientGrants200ResponseOneOf *GetClientGrants200ResponseOneOf
	ArrayOfClientGrant              *[]ClientGrant
}

// GetClientGrants200ResponseOneOfAsGetClientGrants200Response is a convenience function that returns GetClientGrants200ResponseOneOf wrapped in GetClientGrants200Response
func GetClientGrants200ResponseOneOfAsGetClientGrants200Response(v *GetClientGrants200ResponseOneOf) GetClientGrants200Response {
	return GetClientGrants200Response{
		GetClientGrants200ResponseOneOf: v,
	}
}

// []ClientGrantAsGetClientGrants200Response is a convenience function that returns []ClientGrant wrapped in GetClientGrants200Response
func ArrayOfClientGrantAsGetClientGrants200Response(v *[]ClientGrant) GetClientGrants200Response {
	return GetClientGrants200Response{
		ArrayOfClientGrant: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetClientGrants200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetClientGrants200ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.GetClientGrants200ResponseOneOf)
	if err == nil {
		jsonGetClientGrants200ResponseOneOf, _ := json.Marshal(dst.GetClientGrants200ResponseOneOf)
		if string(jsonGetClientGrants200ResponseOneOf) == "{}" { // empty struct
			dst.GetClientGrants200ResponseOneOf = nil
		} else {
			if err = validator.Validate(dst.GetClientGrants200ResponseOneOf); err != nil {
				dst.GetClientGrants200ResponseOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.GetClientGrants200ResponseOneOf = nil
	}

	// try to unmarshal data into ArrayOfClientGrant
	err = newStrictDecoder(data).Decode(&dst.ArrayOfClientGrant)
	if err == nil {
		jsonArrayOfClientGrant, _ := json.Marshal(dst.ArrayOfClientGrant)
		if string(jsonArrayOfClientGrant) == "{}" { // empty struct
			dst.ArrayOfClientGrant = nil
		} else {
			if err = validator.Validate(dst.ArrayOfClientGrant); err != nil {
				dst.ArrayOfClientGrant = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfClientGrant = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetClientGrants200ResponseOneOf = nil
		dst.ArrayOfClientGrant = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetClientGrants200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetClientGrants200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetClientGrants200Response) MarshalJSON() ([]byte, error) {
	if src.GetClientGrants200ResponseOneOf != nil {
		return json.Marshal(&src.GetClientGrants200ResponseOneOf)
	}

	if src.ArrayOfClientGrant != nil {
		return json.Marshal(&src.ArrayOfClientGrant)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetClientGrants200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetClientGrants200ResponseOneOf != nil {
		return obj.GetClientGrants200ResponseOneOf
	}

	if obj.ArrayOfClientGrant != nil {
		return obj.ArrayOfClientGrant
	}

	// all schemas are nil
	return nil
}

type NullableGetClientGrants200Response struct {
	value *GetClientGrants200Response
	isSet bool
}

func (v NullableGetClientGrants200Response) Get() *GetClientGrants200Response {
	return v.value
}

func (v *NullableGetClientGrants200Response) Set(val *GetClientGrants200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetClientGrants200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetClientGrants200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetClientGrants200Response(val *GetClientGrants200Response) *NullableGetClientGrants200Response {
	return &NullableGetClientGrants200Response{value: val, isSet: true}
}

func (v NullableGetClientGrants200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetClientGrants200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
