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

// ResourceServerCreateTokenDialect Dialect of issued access token. `access_token` is a JWT containing standard Auth0 claims; `rfc9068_profile` is a JWT conforming to the IETF JWT Access Token Profile. `access_token_authz` and `rfc9068_profile_authz` additionally include RBAC permissions claims.
type ResourceServerCreateTokenDialect string

// List of resourceServerCreate_token_dialect
const (
	RESOURCESERVERCREATETOKENDIALECT_ACCESS_TOKEN          ResourceServerCreateTokenDialect = "access_token"
	RESOURCESERVERCREATETOKENDIALECT_ACCESS_TOKEN_AUTHZ    ResourceServerCreateTokenDialect = "access_token_authz"
	RESOURCESERVERCREATETOKENDIALECT_RFC9068_PROFILE       ResourceServerCreateTokenDialect = "rfc9068_profile"
	RESOURCESERVERCREATETOKENDIALECT_RFC9068_PROFILE_AUTHZ ResourceServerCreateTokenDialect = "rfc9068_profile_authz"
)

// All allowed values of ResourceServerCreateTokenDialect enum
var AllowedResourceServerCreateTokenDialectEnumValues = []ResourceServerCreateTokenDialect{
	"access_token",
	"access_token_authz",
	"rfc9068_profile",
	"rfc9068_profile_authz",
}

func (v *ResourceServerCreateTokenDialect) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResourceServerCreateTokenDialect(value)
	for _, existing := range AllowedResourceServerCreateTokenDialectEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResourceServerCreateTokenDialect", value)
}

// NewResourceServerCreateTokenDialectFromValue returns a pointer to a valid ResourceServerCreateTokenDialect
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceServerCreateTokenDialectFromValue(v string) (*ResourceServerCreateTokenDialect, error) {
	ev := ResourceServerCreateTokenDialect(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResourceServerCreateTokenDialect: valid values are %v", v, AllowedResourceServerCreateTokenDialectEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResourceServerCreateTokenDialect) IsValid() bool {
	for _, existing := range AllowedResourceServerCreateTokenDialectEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceServerCreate_token_dialect value
func (v ResourceServerCreateTokenDialect) Ptr() *ResourceServerCreateTokenDialect {
	return &v
}

type NullableResourceServerCreateTokenDialect struct {
	value *ResourceServerCreateTokenDialect
	isSet bool
}

func (v NullableResourceServerCreateTokenDialect) Get() *ResourceServerCreateTokenDialect {
	return v.value
}

func (v *NullableResourceServerCreateTokenDialect) Set(val *ResourceServerCreateTokenDialect) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceServerCreateTokenDialect) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceServerCreateTokenDialect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceServerCreateTokenDialect(val *ResourceServerCreateTokenDialect) *NullableResourceServerCreateTokenDialect {
	return &NullableResourceServerCreateTokenDialect{value: val, isSet: true}
}

func (v NullableResourceServerCreateTokenDialect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceServerCreateTokenDialect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
