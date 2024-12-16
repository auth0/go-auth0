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
)

// EmailProviderCreateCredentialsAnyOf1 struct for EmailProviderCreateCredentialsAnyOf1
type EmailProviderCreateCredentialsAnyOf1 struct {
	SmtpHost EmailProviderUpdateCredentialsAnyOf2SmtpHost `json:"smtp_host"`
	// SMTP port.
	SmtpPort int32 `json:"smtp_port"`
	// SMTP username.
	SmtpUser string `json:"smtp_user"`
	// SMTP password.
	SmtpPass string `json:"smtp_pass"`
}

type _EmailProviderCreateCredentialsAnyOf1 EmailProviderCreateCredentialsAnyOf1

// GetSmtpHost returns the SmtpHost field value
func (o *EmailProviderCreateCredentialsAnyOf1) GetSmtpHost() EmailProviderUpdateCredentialsAnyOf2SmtpHost {
	if o == nil {
		var ret EmailProviderUpdateCredentialsAnyOf2SmtpHost
		return ret
	}

	return o.SmtpHost
}

// GetSmtpHostOk returns a tuple with the SmtpHost field value
// and a boolean to check if the value has been set.
func (o *EmailProviderCreateCredentialsAnyOf1) GetSmtpHostOk() (*EmailProviderUpdateCredentialsAnyOf2SmtpHost, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmtpHost, true
}

// SetSmtpHost sets field value
func (o *EmailProviderCreateCredentialsAnyOf1) SetSmtpHost(v EmailProviderUpdateCredentialsAnyOf2SmtpHost) {
	o.SmtpHost = v
}

// GetSmtpPort returns the SmtpPort field value
func (o *EmailProviderCreateCredentialsAnyOf1) GetSmtpPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SmtpPort
}

// GetSmtpPortOk returns a tuple with the SmtpPort field value
// and a boolean to check if the value has been set.
func (o *EmailProviderCreateCredentialsAnyOf1) GetSmtpPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmtpPort, true
}

// SetSmtpPort sets field value
func (o *EmailProviderCreateCredentialsAnyOf1) SetSmtpPort(v int32) {
	o.SmtpPort = v
}

// GetSmtpUser returns the SmtpUser field value
func (o *EmailProviderCreateCredentialsAnyOf1) GetSmtpUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SmtpUser
}

// GetSmtpUserOk returns a tuple with the SmtpUser field value
// and a boolean to check if the value has been set.
func (o *EmailProviderCreateCredentialsAnyOf1) GetSmtpUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmtpUser, true
}

// SetSmtpUser sets field value
func (o *EmailProviderCreateCredentialsAnyOf1) SetSmtpUser(v string) {
	o.SmtpUser = v
}

// GetSmtpPass returns the SmtpPass field value
func (o *EmailProviderCreateCredentialsAnyOf1) GetSmtpPass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SmtpPass
}

// GetSmtpPassOk returns a tuple with the SmtpPass field value
// and a boolean to check if the value has been set.
func (o *EmailProviderCreateCredentialsAnyOf1) GetSmtpPassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmtpPass, true
}

// SetSmtpPass sets field value
func (o *EmailProviderCreateCredentialsAnyOf1) SetSmtpPass(v string) {
	o.SmtpPass = v
}

func (o EmailProviderCreateCredentialsAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailProviderCreateCredentialsAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["smtp_host"] = o.SmtpHost
	toSerialize["smtp_port"] = o.SmtpPort
	toSerialize["smtp_user"] = o.SmtpUser
	toSerialize["smtp_pass"] = o.SmtpPass
	return toSerialize, nil
}

func (o *EmailProviderCreateCredentialsAnyOf1) UnmarshalJSON(data []byte) (err error) {
	varEmailProviderCreateCredentialsAnyOf1 := _EmailProviderCreateCredentialsAnyOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmailProviderCreateCredentialsAnyOf1)

	if err != nil {
		return err
	}

	*o = EmailProviderCreateCredentialsAnyOf1(varEmailProviderCreateCredentialsAnyOf1)

	return err
}

type NullableEmailProviderCreateCredentialsAnyOf1 struct {
	value *EmailProviderCreateCredentialsAnyOf1
	isSet bool
}

func (v NullableEmailProviderCreateCredentialsAnyOf1) Get() *EmailProviderCreateCredentialsAnyOf1 {
	return v.value
}

func (v *NullableEmailProviderCreateCredentialsAnyOf1) Set(val *EmailProviderCreateCredentialsAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailProviderCreateCredentialsAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailProviderCreateCredentialsAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailProviderCreateCredentialsAnyOf1(val *EmailProviderCreateCredentialsAnyOf1) *NullableEmailProviderCreateCredentialsAnyOf1 {
	return &NullableEmailProviderCreateCredentialsAnyOf1{value: val, isSet: true}
}

func (v NullableEmailProviderCreateCredentialsAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailProviderCreateCredentialsAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
