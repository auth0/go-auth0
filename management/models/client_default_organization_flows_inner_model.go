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

// ClientDefaultOrganizationFlowsInner the model 'ClientDefaultOrganizationFlowsInner'
type ClientDefaultOrganizationFlowsInner string

// List of client_default_organization_flows_inner
const (
	CLIENTDEFAULTORGANIZATIONFLOWSINNER_CLIENT_CREDENTIALS ClientDefaultOrganizationFlowsInner = "client_credentials"
)

// All allowed values of ClientDefaultOrganizationFlowsInner enum
var AllowedClientDefaultOrganizationFlowsInnerEnumValues = []ClientDefaultOrganizationFlowsInner{
	"client_credentials",
}

func (v *ClientDefaultOrganizationFlowsInner) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClientDefaultOrganizationFlowsInner(value)
	for _, existing := range AllowedClientDefaultOrganizationFlowsInnerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClientDefaultOrganizationFlowsInner", value)
}

// NewClientDefaultOrganizationFlowsInnerFromValue returns a pointer to a valid ClientDefaultOrganizationFlowsInner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClientDefaultOrganizationFlowsInnerFromValue(v string) (*ClientDefaultOrganizationFlowsInner, error) {
	ev := ClientDefaultOrganizationFlowsInner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClientDefaultOrganizationFlowsInner: valid values are %v", v, AllowedClientDefaultOrganizationFlowsInnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClientDefaultOrganizationFlowsInner) IsValid() bool {
	for _, existing := range AllowedClientDefaultOrganizationFlowsInnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to client_default_organization_flows_inner value
func (v ClientDefaultOrganizationFlowsInner) Ptr() *ClientDefaultOrganizationFlowsInner {
	return &v
}

type NullableClientDefaultOrganizationFlowsInner struct {
	value *ClientDefaultOrganizationFlowsInner
	isSet bool
}

func (v NullableClientDefaultOrganizationFlowsInner) Get() *ClientDefaultOrganizationFlowsInner {
	return v.value
}

func (v *NullableClientDefaultOrganizationFlowsInner) Set(val *ClientDefaultOrganizationFlowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableClientDefaultOrganizationFlowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableClientDefaultOrganizationFlowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientDefaultOrganizationFlowsInner(val *ClientDefaultOrganizationFlowsInner) *NullableClientDefaultOrganizationFlowsInner {
	return &NullableClientDefaultOrganizationFlowsInner{value: val, isSet: true}
}

func (v NullableClientDefaultOrganizationFlowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientDefaultOrganizationFlowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
