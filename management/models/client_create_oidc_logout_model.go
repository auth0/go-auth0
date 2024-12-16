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

// ClientCreateOidcLogout Configuration for OIDC backchannel logout
type ClientCreateOidcLogout struct {
	// Comma-separated list of URLs that are valid to call back from Auth0 for OIDC backchannel logout. Currently only one URL is allowed.
	BackchannelLogoutUrls []string `json:"backchannel_logout_urls"`
}

type _ClientCreateOidcLogout ClientCreateOidcLogout

// GetBackchannelLogoutUrls returns the BackchannelLogoutUrls field value
func (o *ClientCreateOidcLogout) GetBackchannelLogoutUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BackchannelLogoutUrls
}

// GetBackchannelLogoutUrlsOk returns a tuple with the BackchannelLogoutUrls field value
// and a boolean to check if the value has been set.
func (o *ClientCreateOidcLogout) GetBackchannelLogoutUrlsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackchannelLogoutUrls, true
}

// SetBackchannelLogoutUrls sets field value
func (o *ClientCreateOidcLogout) SetBackchannelLogoutUrls(v []string) {
	o.BackchannelLogoutUrls = v
}

func (o ClientCreateOidcLogout) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientCreateOidcLogout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["backchannel_logout_urls"] = o.BackchannelLogoutUrls
	return toSerialize, nil
}

func (o *ClientCreateOidcLogout) UnmarshalJSON(data []byte) (err error) {
	varClientCreateOidcLogout := _ClientCreateOidcLogout{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClientCreateOidcLogout)

	if err != nil {
		return err
	}

	*o = ClientCreateOidcLogout(varClientCreateOidcLogout)

	return err
}

type NullableClientCreateOidcLogout struct {
	value *ClientCreateOidcLogout
	isSet bool
}

func (v NullableClientCreateOidcLogout) Get() *ClientCreateOidcLogout {
	return v.value
}

func (v *NullableClientCreateOidcLogout) Set(val *ClientCreateOidcLogout) {
	v.value = val
	v.isSet = true
}

func (v NullableClientCreateOidcLogout) IsSet() bool {
	return v.isSet
}

func (v *NullableClientCreateOidcLogout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientCreateOidcLogout(val *ClientCreateOidcLogout) *NullableClientCreateOidcLogout {
	return &NullableClientCreateOidcLogout{value: val, isSet: true}
}

func (v NullableClientCreateOidcLogout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientCreateOidcLogout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
