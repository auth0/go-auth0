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

// ClientCreateClientAuthenticationMethods Defines client authentication methods.
type ClientCreateClientAuthenticationMethods struct {
	PrivateKeyJwt           *ClientCreateClientAuthenticationMethodsPrivateKeyJwt           `json:"private_key_jwt,omitempty"`
	TlsClientAuth           *ClientCreateClientAuthenticationMethodsTlsClientAuth           `json:"tls_client_auth,omitempty"`
	SelfSignedTlsClientAuth *ClientCreateClientAuthenticationMethodsSelfSignedTlsClientAuth `json:"self_signed_tls_client_auth,omitempty"`
}

// GetPrivateKeyJwt returns the PrivateKeyJwt field value if set, zero value otherwise.
func (o *ClientCreateClientAuthenticationMethods) GetPrivateKeyJwt() ClientCreateClientAuthenticationMethodsPrivateKeyJwt {
	if o == nil || IsNil(o.PrivateKeyJwt) {
		var ret ClientCreateClientAuthenticationMethodsPrivateKeyJwt
		return ret
	}
	return *o.PrivateKeyJwt
}

// GetPrivateKeyJwtOk returns a tuple with the PrivateKeyJwt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateClientAuthenticationMethods) GetPrivateKeyJwtOk() (*ClientCreateClientAuthenticationMethodsPrivateKeyJwt, bool) {
	if o == nil || IsNil(o.PrivateKeyJwt) {
		return nil, false
	}
	return o.PrivateKeyJwt, true
}

// HasPrivateKeyJwt returns a boolean if a field has been set.
func (o *ClientCreateClientAuthenticationMethods) HasPrivateKeyJwt() bool {
	if o != nil && !IsNil(o.PrivateKeyJwt) {
		return true
	}

	return false
}

// SetPrivateKeyJwt gets a reference to the given ClientCreateClientAuthenticationMethodsPrivateKeyJwt and assigns it to the PrivateKeyJwt field.
func (o *ClientCreateClientAuthenticationMethods) SetPrivateKeyJwt(v ClientCreateClientAuthenticationMethodsPrivateKeyJwt) {
	o.PrivateKeyJwt = &v
}

// GetTlsClientAuth returns the TlsClientAuth field value if set, zero value otherwise.
func (o *ClientCreateClientAuthenticationMethods) GetTlsClientAuth() ClientCreateClientAuthenticationMethodsTlsClientAuth {
	if o == nil || IsNil(o.TlsClientAuth) {
		var ret ClientCreateClientAuthenticationMethodsTlsClientAuth
		return ret
	}
	return *o.TlsClientAuth
}

// GetTlsClientAuthOk returns a tuple with the TlsClientAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateClientAuthenticationMethods) GetTlsClientAuthOk() (*ClientCreateClientAuthenticationMethodsTlsClientAuth, bool) {
	if o == nil || IsNil(o.TlsClientAuth) {
		return nil, false
	}
	return o.TlsClientAuth, true
}

// HasTlsClientAuth returns a boolean if a field has been set.
func (o *ClientCreateClientAuthenticationMethods) HasTlsClientAuth() bool {
	if o != nil && !IsNil(o.TlsClientAuth) {
		return true
	}

	return false
}

// SetTlsClientAuth gets a reference to the given ClientCreateClientAuthenticationMethodsTlsClientAuth and assigns it to the TlsClientAuth field.
func (o *ClientCreateClientAuthenticationMethods) SetTlsClientAuth(v ClientCreateClientAuthenticationMethodsTlsClientAuth) {
	o.TlsClientAuth = &v
}

// GetSelfSignedTlsClientAuth returns the SelfSignedTlsClientAuth field value if set, zero value otherwise.
func (o *ClientCreateClientAuthenticationMethods) GetSelfSignedTlsClientAuth() ClientCreateClientAuthenticationMethodsSelfSignedTlsClientAuth {
	if o == nil || IsNil(o.SelfSignedTlsClientAuth) {
		var ret ClientCreateClientAuthenticationMethodsSelfSignedTlsClientAuth
		return ret
	}
	return *o.SelfSignedTlsClientAuth
}

// GetSelfSignedTlsClientAuthOk returns a tuple with the SelfSignedTlsClientAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateClientAuthenticationMethods) GetSelfSignedTlsClientAuthOk() (*ClientCreateClientAuthenticationMethodsSelfSignedTlsClientAuth, bool) {
	if o == nil || IsNil(o.SelfSignedTlsClientAuth) {
		return nil, false
	}
	return o.SelfSignedTlsClientAuth, true
}

// HasSelfSignedTlsClientAuth returns a boolean if a field has been set.
func (o *ClientCreateClientAuthenticationMethods) HasSelfSignedTlsClientAuth() bool {
	if o != nil && !IsNil(o.SelfSignedTlsClientAuth) {
		return true
	}

	return false
}

// SetSelfSignedTlsClientAuth gets a reference to the given ClientCreateClientAuthenticationMethodsSelfSignedTlsClientAuth and assigns it to the SelfSignedTlsClientAuth field.
func (o *ClientCreateClientAuthenticationMethods) SetSelfSignedTlsClientAuth(v ClientCreateClientAuthenticationMethodsSelfSignedTlsClientAuth) {
	o.SelfSignedTlsClientAuth = &v
}

func (o ClientCreateClientAuthenticationMethods) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientCreateClientAuthenticationMethods) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrivateKeyJwt) {
		toSerialize["private_key_jwt"] = o.PrivateKeyJwt
	}
	if !IsNil(o.TlsClientAuth) {
		toSerialize["tls_client_auth"] = o.TlsClientAuth
	}
	if !IsNil(o.SelfSignedTlsClientAuth) {
		toSerialize["self_signed_tls_client_auth"] = o.SelfSignedTlsClientAuth
	}
	return toSerialize, nil
}

type NullableClientCreateClientAuthenticationMethods struct {
	value *ClientCreateClientAuthenticationMethods
	isSet bool
}

func (v NullableClientCreateClientAuthenticationMethods) Get() *ClientCreateClientAuthenticationMethods {
	return v.value
}

func (v *NullableClientCreateClientAuthenticationMethods) Set(val *ClientCreateClientAuthenticationMethods) {
	v.value = val
	v.isSet = true
}

func (v NullableClientCreateClientAuthenticationMethods) IsSet() bool {
	return v.isSet
}

func (v *NullableClientCreateClientAuthenticationMethods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientCreateClientAuthenticationMethods(val *ClientCreateClientAuthenticationMethods) *NullableClientCreateClientAuthenticationMethods {
	return &NullableClientCreateClientAuthenticationMethods{value: val, isSet: true}
}

func (v NullableClientCreateClientAuthenticationMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientCreateClientAuthenticationMethods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
