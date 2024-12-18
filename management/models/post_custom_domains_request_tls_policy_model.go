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

// PostCustomDomainsRequestTlsPolicy recommended includes TLS 1.2
type PostCustomDomainsRequestTlsPolicy string

// List of post_custom_domains_request_tls_policy
const (
	POSTCUSTOMDOMAINSREQUESTTLSPOLICY_RECOMMENDED PostCustomDomainsRequestTlsPolicy = "recommended"
	POSTCUSTOMDOMAINSREQUESTTLSPOLICY_COMPATIBLE  PostCustomDomainsRequestTlsPolicy = "compatible"
)

// All allowed values of PostCustomDomainsRequestTlsPolicy enum
var AllowedPostCustomDomainsRequestTlsPolicyEnumValues = []PostCustomDomainsRequestTlsPolicy{
	"recommended",
	"compatible",
}

func (v *PostCustomDomainsRequestTlsPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PostCustomDomainsRequestTlsPolicy(value)
	for _, existing := range AllowedPostCustomDomainsRequestTlsPolicyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PostCustomDomainsRequestTlsPolicy", value)
}

// NewPostCustomDomainsRequestTlsPolicyFromValue returns a pointer to a valid PostCustomDomainsRequestTlsPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPostCustomDomainsRequestTlsPolicyFromValue(v string) (*PostCustomDomainsRequestTlsPolicy, error) {
	ev := PostCustomDomainsRequestTlsPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PostCustomDomainsRequestTlsPolicy: valid values are %v", v, AllowedPostCustomDomainsRequestTlsPolicyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PostCustomDomainsRequestTlsPolicy) IsValid() bool {
	for _, existing := range AllowedPostCustomDomainsRequestTlsPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to post_custom_domains_request_tls_policy value
func (v PostCustomDomainsRequestTlsPolicy) Ptr() *PostCustomDomainsRequestTlsPolicy {
	return &v
}

type NullablePostCustomDomainsRequestTlsPolicy struct {
	value *PostCustomDomainsRequestTlsPolicy
	isSet bool
}

func (v NullablePostCustomDomainsRequestTlsPolicy) Get() *PostCustomDomainsRequestTlsPolicy {
	return v.value
}

func (v *NullablePostCustomDomainsRequestTlsPolicy) Set(val *PostCustomDomainsRequestTlsPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCustomDomainsRequestTlsPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCustomDomainsRequestTlsPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCustomDomainsRequestTlsPolicy(val *PostCustomDomainsRequestTlsPolicy) *NullablePostCustomDomainsRequestTlsPolicy {
	return &NullablePostCustomDomainsRequestTlsPolicy{value: val, isSet: true}
}

func (v NullablePostCustomDomainsRequestTlsPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCustomDomainsRequestTlsPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
