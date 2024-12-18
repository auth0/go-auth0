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

// PatchClientGrantsByIdRequestOrganizationUsage Controls how organizations may be used with this grant
type PatchClientGrantsByIdRequestOrganizationUsage string

// List of patch_client_grants_by_id_request_organization_usage
const (
	PATCHCLIENTGRANTSBYIDREQUESTORGANIZATIONUSAGE_DENY    PatchClientGrantsByIdRequestOrganizationUsage = "deny"
	PATCHCLIENTGRANTSBYIDREQUESTORGANIZATIONUSAGE_ALLOW   PatchClientGrantsByIdRequestOrganizationUsage = "allow"
	PATCHCLIENTGRANTSBYIDREQUESTORGANIZATIONUSAGE_REQUIRE PatchClientGrantsByIdRequestOrganizationUsage = "require"
)

// All allowed values of PatchClientGrantsByIdRequestOrganizationUsage enum
var AllowedPatchClientGrantsByIdRequestOrganizationUsageEnumValues = []PatchClientGrantsByIdRequestOrganizationUsage{
	"deny",
	"allow",
	"require",
}

func (v *PatchClientGrantsByIdRequestOrganizationUsage) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchClientGrantsByIdRequestOrganizationUsage(value)
	for _, existing := range AllowedPatchClientGrantsByIdRequestOrganizationUsageEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchClientGrantsByIdRequestOrganizationUsage", value)
}

// NewPatchClientGrantsByIdRequestOrganizationUsageFromValue returns a pointer to a valid PatchClientGrantsByIdRequestOrganizationUsage
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchClientGrantsByIdRequestOrganizationUsageFromValue(v string) (*PatchClientGrantsByIdRequestOrganizationUsage, error) {
	ev := PatchClientGrantsByIdRequestOrganizationUsage(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchClientGrantsByIdRequestOrganizationUsage: valid values are %v", v, AllowedPatchClientGrantsByIdRequestOrganizationUsageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchClientGrantsByIdRequestOrganizationUsage) IsValid() bool {
	for _, existing := range AllowedPatchClientGrantsByIdRequestOrganizationUsageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to patch_client_grants_by_id_request_organization_usage value
func (v PatchClientGrantsByIdRequestOrganizationUsage) Ptr() *PatchClientGrantsByIdRequestOrganizationUsage {
	return &v
}

type NullablePatchClientGrantsByIdRequestOrganizationUsage struct {
	value *PatchClientGrantsByIdRequestOrganizationUsage
	isSet bool
}

func (v NullablePatchClientGrantsByIdRequestOrganizationUsage) Get() *PatchClientGrantsByIdRequestOrganizationUsage {
	return v.value
}

func (v *NullablePatchClientGrantsByIdRequestOrganizationUsage) Set(val *PatchClientGrantsByIdRequestOrganizationUsage) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchClientGrantsByIdRequestOrganizationUsage) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchClientGrantsByIdRequestOrganizationUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchClientGrantsByIdRequestOrganizationUsage(val *PatchClientGrantsByIdRequestOrganizationUsage) *NullablePatchClientGrantsByIdRequestOrganizationUsage {
	return &NullablePatchClientGrantsByIdRequestOrganizationUsage{value: val, isSet: true}
}

func (v NullablePatchClientGrantsByIdRequestOrganizationUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchClientGrantsByIdRequestOrganizationUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
