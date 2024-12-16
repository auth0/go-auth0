/*
Auth0 Management API

Auth0 Management API v2.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// CustomDomain struct for CustomDomain
type CustomDomain struct {
	// ID of the custom domain.
	CustomDomainId string `json:"custom_domain_id"`
	// Domain name.
	Domain string `json:"domain"`
	// Whether this is a primary domain (true) or not (false).
	Primary bool                               `json:"primary"`
	Status  PostCustomDomains201ResponseStatus `json:"status"`
	Type    PostCustomDomains201ResponseType   `json:"type"`
	// Intermediate address.
	OriginDomainName *string                                   `json:"origin_domain_name,omitempty"`
	Verification     *PostCustomDomains201ResponseVerification `json:"verification,omitempty"`
	// The HTTP header to fetch the client's IP address
	CustomClientIpHeader NullableString `json:"custom_client_ip_header,omitempty"`
	// The TLS version policy
	TlsPolicy *string `json:"tls_policy,omitempty"`
}

type _CustomDomain CustomDomain

// GetCustomDomainId returns the CustomDomainId field value
func (o *CustomDomain) GetCustomDomainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomDomainId
}

// GetCustomDomainIdOk returns a tuple with the CustomDomainId field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetCustomDomainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomDomainId, true
}

// SetCustomDomainId sets field value
func (o *CustomDomain) SetCustomDomainId(v string) {
	o.CustomDomainId = v
}

// GetDomain returns the Domain field value
func (o *CustomDomain) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *CustomDomain) SetDomain(v string) {
	o.Domain = v
}

// GetPrimary returns the Primary field value
func (o *CustomDomain) GetPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Primary, true
}

// SetPrimary sets field value
func (o *CustomDomain) SetPrimary(v bool) {
	o.Primary = v
}

// GetStatus returns the Status field value
func (o *CustomDomain) GetStatus() PostCustomDomains201ResponseStatus {
	if o == nil {
		var ret PostCustomDomains201ResponseStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetStatusOk() (*PostCustomDomains201ResponseStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CustomDomain) SetStatus(v PostCustomDomains201ResponseStatus) {
	o.Status = v
}

// GetType returns the Type field value
func (o *CustomDomain) GetType() PostCustomDomains201ResponseType {
	if o == nil {
		var ret PostCustomDomains201ResponseType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetTypeOk() (*PostCustomDomains201ResponseType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomDomain) SetType(v PostCustomDomains201ResponseType) {
	o.Type = v
}

// GetOriginDomainName returns the OriginDomainName field value if set, zero value otherwise.
func (o *CustomDomain) GetOriginDomainName() string {
	if o == nil || IsNil(o.OriginDomainName) {
		var ret string
		return ret
	}
	return *o.OriginDomainName
}

// GetOriginDomainNameOk returns a tuple with the OriginDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetOriginDomainNameOk() (*string, bool) {
	if o == nil || IsNil(o.OriginDomainName) {
		return nil, false
	}
	return o.OriginDomainName, true
}

// HasOriginDomainName returns a boolean if a field has been set.
func (o *CustomDomain) HasOriginDomainName() bool {
	if o != nil && !IsNil(o.OriginDomainName) {
		return true
	}

	return false
}

// SetOriginDomainName gets a reference to the given string and assigns it to the OriginDomainName field.
func (o *CustomDomain) SetOriginDomainName(v string) {
	o.OriginDomainName = &v
}

// GetVerification returns the Verification field value if set, zero value otherwise.
func (o *CustomDomain) GetVerification() PostCustomDomains201ResponseVerification {
	if o == nil || IsNil(o.Verification) {
		var ret PostCustomDomains201ResponseVerification
		return ret
	}
	return *o.Verification
}

// GetVerificationOk returns a tuple with the Verification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetVerificationOk() (*PostCustomDomains201ResponseVerification, bool) {
	if o == nil || IsNil(o.Verification) {
		return nil, false
	}
	return o.Verification, true
}

// HasVerification returns a boolean if a field has been set.
func (o *CustomDomain) HasVerification() bool {
	if o != nil && !IsNil(o.Verification) {
		return true
	}

	return false
}

// SetVerification gets a reference to the given PostCustomDomains201ResponseVerification and assigns it to the Verification field.
func (o *CustomDomain) SetVerification(v PostCustomDomains201ResponseVerification) {
	o.Verification = &v
}

// GetCustomClientIpHeader returns the CustomClientIpHeader field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomDomain) GetCustomClientIpHeader() string {
	if o == nil || IsNil(o.CustomClientIpHeader.Get()) {
		var ret string
		return ret
	}
	return *o.CustomClientIpHeader.Get()
}

// GetCustomClientIpHeaderOk returns a tuple with the CustomClientIpHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomDomain) GetCustomClientIpHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomClientIpHeader.Get(), o.CustomClientIpHeader.IsSet()
}

// HasCustomClientIpHeader returns a boolean if a field has been set.
func (o *CustomDomain) HasCustomClientIpHeader() bool {
	if o != nil && o.CustomClientIpHeader.IsSet() {
		return true
	}

	return false
}

// SetCustomClientIpHeader gets a reference to the given NullableString and assigns it to the CustomClientIpHeader field.
func (o *CustomDomain) SetCustomClientIpHeader(v string) {
	o.CustomClientIpHeader.Set(&v)
}

// SetCustomClientIpHeaderNil sets the value for CustomClientIpHeader to be an explicit nil
func (o *CustomDomain) SetCustomClientIpHeaderNil() {
	o.CustomClientIpHeader.Set(nil)
}

// UnsetCustomClientIpHeader ensures that no value is present for CustomClientIpHeader, not even an explicit nil
func (o *CustomDomain) UnsetCustomClientIpHeader() {
	o.CustomClientIpHeader.Unset()
}

// GetTlsPolicy returns the TlsPolicy field value if set, zero value otherwise.
func (o *CustomDomain) GetTlsPolicy() string {
	if o == nil || IsNil(o.TlsPolicy) {
		var ret string
		return ret
	}
	return *o.TlsPolicy
}

// GetTlsPolicyOk returns a tuple with the TlsPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetTlsPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.TlsPolicy) {
		return nil, false
	}
	return o.TlsPolicy, true
}

// HasTlsPolicy returns a boolean if a field has been set.
func (o *CustomDomain) HasTlsPolicy() bool {
	if o != nil && !IsNil(o.TlsPolicy) {
		return true
	}

	return false
}

// SetTlsPolicy gets a reference to the given string and assigns it to the TlsPolicy field.
func (o *CustomDomain) SetTlsPolicy(v string) {
	o.TlsPolicy = &v
}

func (o CustomDomain) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["custom_domain_id"] = o.CustomDomainId
	toSerialize["domain"] = o.Domain
	toSerialize["primary"] = o.Primary
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	if !IsNil(o.OriginDomainName) {
		toSerialize["origin_domain_name"] = o.OriginDomainName
	}
	if !IsNil(o.Verification) {
		toSerialize["verification"] = o.Verification
	}
	if o.CustomClientIpHeader.IsSet() {
		toSerialize["custom_client_ip_header"] = o.CustomClientIpHeader.Get()
	}
	if !IsNil(o.TlsPolicy) {
		toSerialize["tls_policy"] = o.TlsPolicy
	}
	return toSerialize, nil
}

func (o *CustomDomain) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"custom_domain_id",
		"domain",
		"primary",
		"status",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCustomDomain := _CustomDomain{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomDomain)

	if err != nil {
		return err
	}

	*o = CustomDomain(varCustomDomain)

	return err
}

type NullableCustomDomain struct {
	value *CustomDomain
	isSet bool
}

func (v NullableCustomDomain) Get() *CustomDomain {
	return v.value
}

func (v *NullableCustomDomain) Set(val *CustomDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDomain(val *CustomDomain) *NullableCustomDomain {
	return &NullableCustomDomain{value: val, isSet: true}
}

func (v NullableCustomDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
