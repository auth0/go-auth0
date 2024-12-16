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

// PostCredentialsRequestCredentialType Credential type. Supported types: public_key, cert_subject_dn or x509_cert
type PostCredentialsRequestCredentialType string

// List of post_credentials_request_credential_type
const (
	POSTCREDENTIALSREQUESTCREDENTIALTYPE_PUBLIC_KEY      PostCredentialsRequestCredentialType = "public_key"
	POSTCREDENTIALSREQUESTCREDENTIALTYPE_CERT_SUBJECT_DN PostCredentialsRequestCredentialType = "cert_subject_dn"
	POSTCREDENTIALSREQUESTCREDENTIALTYPE_X509_CERT       PostCredentialsRequestCredentialType = "x509_cert"
)

// All allowed values of PostCredentialsRequestCredentialType enum
var AllowedPostCredentialsRequestCredentialTypeEnumValues = []PostCredentialsRequestCredentialType{
	"public_key",
	"cert_subject_dn",
	"x509_cert",
}

func (v *PostCredentialsRequestCredentialType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PostCredentialsRequestCredentialType(value)
	for _, existing := range AllowedPostCredentialsRequestCredentialTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PostCredentialsRequestCredentialType", value)
}

// NewPostCredentialsRequestCredentialTypeFromValue returns a pointer to a valid PostCredentialsRequestCredentialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPostCredentialsRequestCredentialTypeFromValue(v string) (*PostCredentialsRequestCredentialType, error) {
	ev := PostCredentialsRequestCredentialType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PostCredentialsRequestCredentialType: valid values are %v", v, AllowedPostCredentialsRequestCredentialTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PostCredentialsRequestCredentialType) IsValid() bool {
	for _, existing := range AllowedPostCredentialsRequestCredentialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to post_credentials_request_credential_type value
func (v PostCredentialsRequestCredentialType) Ptr() *PostCredentialsRequestCredentialType {
	return &v
}

type NullablePostCredentialsRequestCredentialType struct {
	value *PostCredentialsRequestCredentialType
	isSet bool
}

func (v NullablePostCredentialsRequestCredentialType) Get() *PostCredentialsRequestCredentialType {
	return v.value
}

func (v *NullablePostCredentialsRequestCredentialType) Set(val *PostCredentialsRequestCredentialType) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCredentialsRequestCredentialType) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCredentialsRequestCredentialType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCredentialsRequestCredentialType(val *PostCredentialsRequestCredentialType) *NullablePostCredentialsRequestCredentialType {
	return &NullablePostCredentialsRequestCredentialType{value: val, isSet: true}
}

func (v NullablePostCredentialsRequestCredentialType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCredentialsRequestCredentialType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
