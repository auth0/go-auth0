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

// ClientEncryptionKey Encryption used for WsFed responses with this client.
type ClientEncryptionKey struct {
	// Encryption Public RSA Key.
	Pub *string `json:"pub,omitempty"`
	// Encryption certificate for public key in X.590 (.CER) format.
	Cert *string `json:"cert,omitempty"`
	// Encryption certificate name for this certificate in the format `/CN={domain}`.
	Subject              *string `json:"subject,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientEncryptionKey ClientEncryptionKey

// GetPub returns the Pub field value if set, zero value otherwise.
func (o *ClientEncryptionKey) GetPub() string {
	if o == nil || IsNil(o.Pub) {
		var ret string
		return ret
	}
	return *o.Pub
}

// GetPubOk returns a tuple with the Pub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientEncryptionKey) GetPubOk() (*string, bool) {
	if o == nil || IsNil(o.Pub) {
		return nil, false
	}
	return o.Pub, true
}

// HasPub returns a boolean if a field has been set.
func (o *ClientEncryptionKey) HasPub() bool {
	if o != nil && !IsNil(o.Pub) {
		return true
	}

	return false
}

// SetPub gets a reference to the given string and assigns it to the Pub field.
func (o *ClientEncryptionKey) SetPub(v string) {
	o.Pub = &v
}

// GetCert returns the Cert field value if set, zero value otherwise.
func (o *ClientEncryptionKey) GetCert() string {
	if o == nil || IsNil(o.Cert) {
		var ret string
		return ret
	}
	return *o.Cert
}

// GetCertOk returns a tuple with the Cert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientEncryptionKey) GetCertOk() (*string, bool) {
	if o == nil || IsNil(o.Cert) {
		return nil, false
	}
	return o.Cert, true
}

// HasCert returns a boolean if a field has been set.
func (o *ClientEncryptionKey) HasCert() bool {
	if o != nil && !IsNil(o.Cert) {
		return true
	}

	return false
}

// SetCert gets a reference to the given string and assigns it to the Cert field.
func (o *ClientEncryptionKey) SetCert(v string) {
	o.Cert = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *ClientEncryptionKey) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientEncryptionKey) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *ClientEncryptionKey) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *ClientEncryptionKey) SetSubject(v string) {
	o.Subject = &v
}

func (o ClientEncryptionKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientEncryptionKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pub) {
		toSerialize["pub"] = o.Pub
	}
	if !IsNil(o.Cert) {
		toSerialize["cert"] = o.Cert
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientEncryptionKey) UnmarshalJSON(data []byte) (err error) {
	varClientEncryptionKey := _ClientEncryptionKey{}

	err = json.Unmarshal(data, &varClientEncryptionKey)

	if err != nil {
		return err
	}

	*o = ClientEncryptionKey(varClientEncryptionKey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pub")
		delete(additionalProperties, "cert")
		delete(additionalProperties, "subject")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientEncryptionKey struct {
	value *ClientEncryptionKey
	isSet bool
}

func (v NullableClientEncryptionKey) Get() *ClientEncryptionKey {
	return v.value
}

func (v *NullableClientEncryptionKey) Set(val *ClientEncryptionKey) {
	v.value = val
	v.isSet = true
}

func (v NullableClientEncryptionKey) IsSet() bool {
	return v.isSet
}

func (v *NullableClientEncryptionKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientEncryptionKey(val *ClientEncryptionKey) *NullableClientEncryptionKey {
	return &NullableClientEncryptionKey{value: val, isSet: true}
}

func (v NullableClientEncryptionKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientEncryptionKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
