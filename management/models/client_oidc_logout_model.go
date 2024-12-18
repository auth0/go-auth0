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

// ClientOidcLogout Configuration for OIDC backchannel logout
type ClientOidcLogout struct {
	// Comma-separated list of URLs that are valid to call back from Auth0 for OIDC backchannel logout. Currently only one URL is allowed.
	BackchannelLogoutUrls       []string                                     `json:"backchannel_logout_urls,omitempty"`
	BackchannelLogoutInitiators *ClientOidcLogoutBackchannelLogoutInitiators `json:"backchannel_logout_initiators,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _ClientOidcLogout ClientOidcLogout

// GetBackchannelLogoutUrls returns the BackchannelLogoutUrls field value if set, zero value otherwise.
func (o *ClientOidcLogout) GetBackchannelLogoutUrls() []string {
	if o == nil || IsNil(o.BackchannelLogoutUrls) {
		var ret []string
		return ret
	}
	return o.BackchannelLogoutUrls
}

// GetBackchannelLogoutUrlsOk returns a tuple with the BackchannelLogoutUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOidcLogout) GetBackchannelLogoutUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.BackchannelLogoutUrls) {
		return nil, false
	}
	return o.BackchannelLogoutUrls, true
}

// HasBackchannelLogoutUrls returns a boolean if a field has been set.
func (o *ClientOidcLogout) HasBackchannelLogoutUrls() bool {
	if o != nil && !IsNil(o.BackchannelLogoutUrls) {
		return true
	}

	return false
}

// SetBackchannelLogoutUrls gets a reference to the given []string and assigns it to the BackchannelLogoutUrls field.
func (o *ClientOidcLogout) SetBackchannelLogoutUrls(v []string) {
	o.BackchannelLogoutUrls = v
}

// GetBackchannelLogoutInitiators returns the BackchannelLogoutInitiators field value if set, zero value otherwise.
func (o *ClientOidcLogout) GetBackchannelLogoutInitiators() ClientOidcLogoutBackchannelLogoutInitiators {
	if o == nil || IsNil(o.BackchannelLogoutInitiators) {
		var ret ClientOidcLogoutBackchannelLogoutInitiators
		return ret
	}
	return *o.BackchannelLogoutInitiators
}

// GetBackchannelLogoutInitiatorsOk returns a tuple with the BackchannelLogoutInitiators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOidcLogout) GetBackchannelLogoutInitiatorsOk() (*ClientOidcLogoutBackchannelLogoutInitiators, bool) {
	if o == nil || IsNil(o.BackchannelLogoutInitiators) {
		return nil, false
	}
	return o.BackchannelLogoutInitiators, true
}

// HasBackchannelLogoutInitiators returns a boolean if a field has been set.
func (o *ClientOidcLogout) HasBackchannelLogoutInitiators() bool {
	if o != nil && !IsNil(o.BackchannelLogoutInitiators) {
		return true
	}

	return false
}

// SetBackchannelLogoutInitiators gets a reference to the given ClientOidcLogoutBackchannelLogoutInitiators and assigns it to the BackchannelLogoutInitiators field.
func (o *ClientOidcLogout) SetBackchannelLogoutInitiators(v ClientOidcLogoutBackchannelLogoutInitiators) {
	o.BackchannelLogoutInitiators = &v
}

func (o ClientOidcLogout) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientOidcLogout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackchannelLogoutUrls) {
		toSerialize["backchannel_logout_urls"] = o.BackchannelLogoutUrls
	}
	if !IsNil(o.BackchannelLogoutInitiators) {
		toSerialize["backchannel_logout_initiators"] = o.BackchannelLogoutInitiators
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientOidcLogout) UnmarshalJSON(data []byte) (err error) {
	varClientOidcLogout := _ClientOidcLogout{}

	err = json.Unmarshal(data, &varClientOidcLogout)

	if err != nil {
		return err
	}

	*o = ClientOidcLogout(varClientOidcLogout)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "backchannel_logout_urls")
		delete(additionalProperties, "backchannel_logout_initiators")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientOidcLogout struct {
	value *ClientOidcLogout
	isSet bool
}

func (v NullableClientOidcLogout) Get() *ClientOidcLogout {
	return v.value
}

func (v *NullableClientOidcLogout) Set(val *ClientOidcLogout) {
	v.value = val
	v.isSet = true
}

func (v NullableClientOidcLogout) IsSet() bool {
	return v.isSet
}

func (v *NullableClientOidcLogout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientOidcLogout(val *ClientOidcLogout) *NullableClientOidcLogout {
	return &NullableClientOidcLogout{value: val, isSet: true}
}

func (v NullableClientOidcLogout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientOidcLogout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
