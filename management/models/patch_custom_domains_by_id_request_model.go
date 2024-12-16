/*
Auth0 Management API

Auth0 Management API v2.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// PatchCustomDomainsByIdRequest struct for PatchCustomDomainsByIdRequest
type PatchCustomDomainsByIdRequest struct {
	TlsPolicy            *PostCustomDomainsRequestTlsPolicy                 `json:"tls_policy,omitempty"`
	CustomClientIpHeader *PatchCustomDomainsByIdRequestCustomClientIpHeader `json:"custom_client_ip_header,omitempty"`
}

// GetTlsPolicy returns the TlsPolicy field value if set, zero value otherwise.
func (o *PatchCustomDomainsByIdRequest) GetTlsPolicy() PostCustomDomainsRequestTlsPolicy {
	if o == nil || IsNil(o.TlsPolicy) {
		var ret PostCustomDomainsRequestTlsPolicy
		return ret
	}
	return *o.TlsPolicy
}

// GetTlsPolicyOk returns a tuple with the TlsPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchCustomDomainsByIdRequest) GetTlsPolicyOk() (*PostCustomDomainsRequestTlsPolicy, bool) {
	if o == nil || IsNil(o.TlsPolicy) {
		return nil, false
	}
	return o.TlsPolicy, true
}

// HasTlsPolicy returns a boolean if a field has been set.
func (o *PatchCustomDomainsByIdRequest) HasTlsPolicy() bool {
	if o != nil && !IsNil(o.TlsPolicy) {
		return true
	}

	return false
}

// SetTlsPolicy gets a reference to the given PostCustomDomainsRequestTlsPolicy and assigns it to the TlsPolicy field.
func (o *PatchCustomDomainsByIdRequest) SetTlsPolicy(v PostCustomDomainsRequestTlsPolicy) {
	o.TlsPolicy = &v
}

// GetCustomClientIpHeader returns the CustomClientIpHeader field value if set, zero value otherwise.
func (o *PatchCustomDomainsByIdRequest) GetCustomClientIpHeader() PatchCustomDomainsByIdRequestCustomClientIpHeader {
	if o == nil || IsNil(o.CustomClientIpHeader) {
		var ret PatchCustomDomainsByIdRequestCustomClientIpHeader
		return ret
	}
	return *o.CustomClientIpHeader
}

// GetCustomClientIpHeaderOk returns a tuple with the CustomClientIpHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchCustomDomainsByIdRequest) GetCustomClientIpHeaderOk() (*PatchCustomDomainsByIdRequestCustomClientIpHeader, bool) {
	if o == nil || IsNil(o.CustomClientIpHeader) {
		return nil, false
	}
	return o.CustomClientIpHeader, true
}

// HasCustomClientIpHeader returns a boolean if a field has been set.
func (o *PatchCustomDomainsByIdRequest) HasCustomClientIpHeader() bool {
	if o != nil && !IsNil(o.CustomClientIpHeader) {
		return true
	}

	return false
}

// SetCustomClientIpHeader gets a reference to the given PatchCustomDomainsByIdRequestCustomClientIpHeader and assigns it to the CustomClientIpHeader field.
func (o *PatchCustomDomainsByIdRequest) SetCustomClientIpHeader(v PatchCustomDomainsByIdRequestCustomClientIpHeader) {
	o.CustomClientIpHeader = &v
}

func (o PatchCustomDomainsByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchCustomDomainsByIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TlsPolicy) {
		toSerialize["tls_policy"] = o.TlsPolicy
	}
	if !IsNil(o.CustomClientIpHeader) {
		toSerialize["custom_client_ip_header"] = o.CustomClientIpHeader
	}
	return toSerialize, nil
}

type NullablePatchCustomDomainsByIdRequest struct {
	value *PatchCustomDomainsByIdRequest
	isSet bool
}

func (v NullablePatchCustomDomainsByIdRequest) Get() *PatchCustomDomainsByIdRequest {
	return v.value
}

func (v *NullablePatchCustomDomainsByIdRequest) Set(val *PatchCustomDomainsByIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchCustomDomainsByIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchCustomDomainsByIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchCustomDomainsByIdRequest(val *PatchCustomDomainsByIdRequest) *NullablePatchCustomDomainsByIdRequest {
	return &NullablePatchCustomDomainsByIdRequest{value: val, isSet: true}
}

func (v NullablePatchCustomDomainsByIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchCustomDomainsByIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
