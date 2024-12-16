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

// PatchLogStreamsByIdRequestSinkOneOf1 struct for PatchLogStreamsByIdRequestSinkOneOf1
type PatchLogStreamsByIdRequestSinkOneOf1 struct {
	// Splunk URL Endpoint
	SplunkDomain string `json:"splunkDomain"`
	// Port
	SplunkPort string `json:"splunkPort"`
	// Splunk token
	SplunkToken *string `json:"splunkToken,omitempty"`
	// Verify TLS certificate
	SplunkSecure bool `json:"splunkSecure"`
}

type _PatchLogStreamsByIdRequestSinkOneOf1 PatchLogStreamsByIdRequestSinkOneOf1

// GetSplunkDomain returns the SplunkDomain field value
func (o *PatchLogStreamsByIdRequestSinkOneOf1) GetSplunkDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SplunkDomain
}

// GetSplunkDomainOk returns a tuple with the SplunkDomain field value
// and a boolean to check if the value has been set.
func (o *PatchLogStreamsByIdRequestSinkOneOf1) GetSplunkDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SplunkDomain, true
}

// SetSplunkDomain sets field value
func (o *PatchLogStreamsByIdRequestSinkOneOf1) SetSplunkDomain(v string) {
	o.SplunkDomain = v
}

// GetSplunkPort returns the SplunkPort field value
func (o *PatchLogStreamsByIdRequestSinkOneOf1) GetSplunkPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SplunkPort
}

// GetSplunkPortOk returns a tuple with the SplunkPort field value
// and a boolean to check if the value has been set.
func (o *PatchLogStreamsByIdRequestSinkOneOf1) GetSplunkPortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SplunkPort, true
}

// SetSplunkPort sets field value
func (o *PatchLogStreamsByIdRequestSinkOneOf1) SetSplunkPort(v string) {
	o.SplunkPort = v
}

// GetSplunkToken returns the SplunkToken field value if set, zero value otherwise.
func (o *PatchLogStreamsByIdRequestSinkOneOf1) GetSplunkToken() string {
	if o == nil || IsNil(o.SplunkToken) {
		var ret string
		return ret
	}
	return *o.SplunkToken
}

// GetSplunkTokenOk returns a tuple with the SplunkToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchLogStreamsByIdRequestSinkOneOf1) GetSplunkTokenOk() (*string, bool) {
	if o == nil || IsNil(o.SplunkToken) {
		return nil, false
	}
	return o.SplunkToken, true
}

// HasSplunkToken returns a boolean if a field has been set.
func (o *PatchLogStreamsByIdRequestSinkOneOf1) HasSplunkToken() bool {
	if o != nil && !IsNil(o.SplunkToken) {
		return true
	}

	return false
}

// SetSplunkToken gets a reference to the given string and assigns it to the SplunkToken field.
func (o *PatchLogStreamsByIdRequestSinkOneOf1) SetSplunkToken(v string) {
	o.SplunkToken = &v
}

// GetSplunkSecure returns the SplunkSecure field value
func (o *PatchLogStreamsByIdRequestSinkOneOf1) GetSplunkSecure() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SplunkSecure
}

// GetSplunkSecureOk returns a tuple with the SplunkSecure field value
// and a boolean to check if the value has been set.
func (o *PatchLogStreamsByIdRequestSinkOneOf1) GetSplunkSecureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SplunkSecure, true
}

// SetSplunkSecure sets field value
func (o *PatchLogStreamsByIdRequestSinkOneOf1) SetSplunkSecure(v bool) {
	o.SplunkSecure = v
}

func (o PatchLogStreamsByIdRequestSinkOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchLogStreamsByIdRequestSinkOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["splunkDomain"] = o.SplunkDomain
	toSerialize["splunkPort"] = o.SplunkPort
	if !IsNil(o.SplunkToken) {
		toSerialize["splunkToken"] = o.SplunkToken
	}
	toSerialize["splunkSecure"] = o.SplunkSecure
	return toSerialize, nil
}

func (o *PatchLogStreamsByIdRequestSinkOneOf1) UnmarshalJSON(data []byte) (err error) {
	varPatchLogStreamsByIdRequestSinkOneOf1 := _PatchLogStreamsByIdRequestSinkOneOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPatchLogStreamsByIdRequestSinkOneOf1)

	if err != nil {
		return err
	}

	*o = PatchLogStreamsByIdRequestSinkOneOf1(varPatchLogStreamsByIdRequestSinkOneOf1)

	return err
}

type NullablePatchLogStreamsByIdRequestSinkOneOf1 struct {
	value *PatchLogStreamsByIdRequestSinkOneOf1
	isSet bool
}

func (v NullablePatchLogStreamsByIdRequestSinkOneOf1) Get() *PatchLogStreamsByIdRequestSinkOneOf1 {
	return v.value
}

func (v *NullablePatchLogStreamsByIdRequestSinkOneOf1) Set(val *PatchLogStreamsByIdRequestSinkOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchLogStreamsByIdRequestSinkOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchLogStreamsByIdRequestSinkOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchLogStreamsByIdRequestSinkOneOf1(val *PatchLogStreamsByIdRequestSinkOneOf1) *NullablePatchLogStreamsByIdRequestSinkOneOf1 {
	return &NullablePatchLogStreamsByIdRequestSinkOneOf1{value: val, isSet: true}
}

func (v NullablePatchLogStreamsByIdRequestSinkOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchLogStreamsByIdRequestSinkOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
