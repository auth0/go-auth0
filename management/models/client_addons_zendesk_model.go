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

// ClientAddonsZendesk Zendesk SSO configuration.
type ClientAddonsZendesk struct {
	// Zendesk account name usually first segment in your Zendesk URL. e.g. `https://acme-org.zendesk.com` would be `acme-org`.
	AccountName          *string `json:"accountName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientAddonsZendesk ClientAddonsZendesk

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *ClientAddonsZendesk) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAddonsZendesk) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *ClientAddonsZendesk) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *ClientAddonsZendesk) SetAccountName(v string) {
	o.AccountName = &v
}

func (o ClientAddonsZendesk) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientAddonsZendesk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountName) {
		toSerialize["accountName"] = o.AccountName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientAddonsZendesk) UnmarshalJSON(data []byte) (err error) {
	varClientAddonsZendesk := _ClientAddonsZendesk{}

	err = json.Unmarshal(data, &varClientAddonsZendesk)

	if err != nil {
		return err
	}

	*o = ClientAddonsZendesk(varClientAddonsZendesk)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accountName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientAddonsZendesk struct {
	value *ClientAddonsZendesk
	isSet bool
}

func (v NullableClientAddonsZendesk) Get() *ClientAddonsZendesk {
	return v.value
}

func (v *NullableClientAddonsZendesk) Set(val *ClientAddonsZendesk) {
	v.value = val
	v.isSet = true
}

func (v NullableClientAddonsZendesk) IsSet() bool {
	return v.isSet
}

func (v *NullableClientAddonsZendesk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientAddonsZendesk(val *ClientAddonsZendesk) *NullableClientAddonsZendesk {
	return &NullableClientAddonsZendesk{value: val, isSet: true}
}

func (v NullableClientAddonsZendesk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientAddonsZendesk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
