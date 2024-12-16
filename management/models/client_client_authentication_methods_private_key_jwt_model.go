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

// ClientClientAuthenticationMethodsPrivateKeyJwt Defines `private_key_jwt` client authentication method. If this property is defined, the client is enabled to use the Private Key JWT authentication method.
type ClientClientAuthenticationMethodsPrivateKeyJwt struct {
	// A list of unique and previously created credential IDs enabled on the client for Private Key JWT authentication.
	Credentials []ClientClientAuthenticationMethodsPrivateKeyJwtCredentialsInner `json:"credentials"`
}

type _ClientClientAuthenticationMethodsPrivateKeyJwt ClientClientAuthenticationMethodsPrivateKeyJwt

// GetCredentials returns the Credentials field value
func (o *ClientClientAuthenticationMethodsPrivateKeyJwt) GetCredentials() []ClientClientAuthenticationMethodsPrivateKeyJwtCredentialsInner {
	if o == nil {
		var ret []ClientClientAuthenticationMethodsPrivateKeyJwtCredentialsInner
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *ClientClientAuthenticationMethodsPrivateKeyJwt) GetCredentialsOk() ([]ClientClientAuthenticationMethodsPrivateKeyJwtCredentialsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Credentials, true
}

// SetCredentials sets field value
func (o *ClientClientAuthenticationMethodsPrivateKeyJwt) SetCredentials(v []ClientClientAuthenticationMethodsPrivateKeyJwtCredentialsInner) {
	o.Credentials = v
}

func (o ClientClientAuthenticationMethodsPrivateKeyJwt) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientClientAuthenticationMethodsPrivateKeyJwt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["credentials"] = o.Credentials
	return toSerialize, nil
}

func (o *ClientClientAuthenticationMethodsPrivateKeyJwt) UnmarshalJSON(data []byte) (err error) {
	varClientClientAuthenticationMethodsPrivateKeyJwt := _ClientClientAuthenticationMethodsPrivateKeyJwt{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClientClientAuthenticationMethodsPrivateKeyJwt)

	if err != nil {
		return err
	}

	*o = ClientClientAuthenticationMethodsPrivateKeyJwt(varClientClientAuthenticationMethodsPrivateKeyJwt)

	return err
}

type NullableClientClientAuthenticationMethodsPrivateKeyJwt struct {
	value *ClientClientAuthenticationMethodsPrivateKeyJwt
	isSet bool
}

func (v NullableClientClientAuthenticationMethodsPrivateKeyJwt) Get() *ClientClientAuthenticationMethodsPrivateKeyJwt {
	return v.value
}

func (v *NullableClientClientAuthenticationMethodsPrivateKeyJwt) Set(val *ClientClientAuthenticationMethodsPrivateKeyJwt) {
	v.value = val
	v.isSet = true
}

func (v NullableClientClientAuthenticationMethodsPrivateKeyJwt) IsSet() bool {
	return v.isSet
}

func (v *NullableClientClientAuthenticationMethodsPrivateKeyJwt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientClientAuthenticationMethodsPrivateKeyJwt(val *ClientClientAuthenticationMethodsPrivateKeyJwt) *NullableClientClientAuthenticationMethodsPrivateKeyJwt {
	return &NullableClientClientAuthenticationMethodsPrivateKeyJwt{value: val, isSet: true}
}

func (v NullableClientClientAuthenticationMethodsPrivateKeyJwt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientClientAuthenticationMethodsPrivateKeyJwt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
