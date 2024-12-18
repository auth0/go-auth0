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

// PostAuthenticationMethodsRequest struct for PostAuthenticationMethodsRequest
type PostAuthenticationMethodsRequest struct {
	Type PostAuthenticationMethodsRequestType `json:"type"`
	// A human-readable label to identify the authentication method.
	Name *string `json:"name,omitempty"`
	// Base32 encoded secret for TOTP generation.
	TotpSecret *string `json:"totp_secret,omitempty"`
	// Applies to phone authentication methods only. The destination phone number used to send verification codes via text and voice.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// Applies to email authentication methods only. The email address used to send verification messages.
	Email                         *string                                                            `json:"email,omitempty"`
	PreferredAuthenticationMethod *PutAuthenticationMethodsRequestInnerPreferredAuthenticationMethod `json:"preferred_authentication_method,omitempty"`
	// Applies to webauthn authentication methods only. The id of the credential.
	KeyId *string `json:"key_id,omitempty"`
	// Applies to webauthn authentication methods only. The public key.
	PublicKey *string `json:"public_key,omitempty"`
	// Applies to webauthn authentication methods only. The relying party identifier.
	RelyingPartyIdentifier *string `json:"relying_party_identifier,omitempty"`
}

type _PostAuthenticationMethodsRequest PostAuthenticationMethodsRequest

// GetType returns the Type field value
func (o *PostAuthenticationMethodsRequest) GetType() PostAuthenticationMethodsRequestType {
	if o == nil {
		var ret PostAuthenticationMethodsRequestType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostAuthenticationMethodsRequest) GetTypeOk() (*PostAuthenticationMethodsRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostAuthenticationMethodsRequest) SetType(v PostAuthenticationMethodsRequestType) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PostAuthenticationMethodsRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthenticationMethodsRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PostAuthenticationMethodsRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PostAuthenticationMethodsRequest) SetName(v string) {
	o.Name = &v
}

// GetTotpSecret returns the TotpSecret field value if set, zero value otherwise.
func (o *PostAuthenticationMethodsRequest) GetTotpSecret() string {
	if o == nil || IsNil(o.TotpSecret) {
		var ret string
		return ret
	}
	return *o.TotpSecret
}

// GetTotpSecretOk returns a tuple with the TotpSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthenticationMethodsRequest) GetTotpSecretOk() (*string, bool) {
	if o == nil || IsNil(o.TotpSecret) {
		return nil, false
	}
	return o.TotpSecret, true
}

// HasTotpSecret returns a boolean if a field has been set.
func (o *PostAuthenticationMethodsRequest) HasTotpSecret() bool {
	if o != nil && !IsNil(o.TotpSecret) {
		return true
	}

	return false
}

// SetTotpSecret gets a reference to the given string and assigns it to the TotpSecret field.
func (o *PostAuthenticationMethodsRequest) SetTotpSecret(v string) {
	o.TotpSecret = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *PostAuthenticationMethodsRequest) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthenticationMethodsRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *PostAuthenticationMethodsRequest) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *PostAuthenticationMethodsRequest) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PostAuthenticationMethodsRequest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthenticationMethodsRequest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PostAuthenticationMethodsRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PostAuthenticationMethodsRequest) SetEmail(v string) {
	o.Email = &v
}

// GetPreferredAuthenticationMethod returns the PreferredAuthenticationMethod field value if set, zero value otherwise.
func (o *PostAuthenticationMethodsRequest) GetPreferredAuthenticationMethod() PutAuthenticationMethodsRequestInnerPreferredAuthenticationMethod {
	if o == nil || IsNil(o.PreferredAuthenticationMethod) {
		var ret PutAuthenticationMethodsRequestInnerPreferredAuthenticationMethod
		return ret
	}
	return *o.PreferredAuthenticationMethod
}

// GetPreferredAuthenticationMethodOk returns a tuple with the PreferredAuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthenticationMethodsRequest) GetPreferredAuthenticationMethodOk() (*PutAuthenticationMethodsRequestInnerPreferredAuthenticationMethod, bool) {
	if o == nil || IsNil(o.PreferredAuthenticationMethod) {
		return nil, false
	}
	return o.PreferredAuthenticationMethod, true
}

// HasPreferredAuthenticationMethod returns a boolean if a field has been set.
func (o *PostAuthenticationMethodsRequest) HasPreferredAuthenticationMethod() bool {
	if o != nil && !IsNil(o.PreferredAuthenticationMethod) {
		return true
	}

	return false
}

// SetPreferredAuthenticationMethod gets a reference to the given PutAuthenticationMethodsRequestInnerPreferredAuthenticationMethod and assigns it to the PreferredAuthenticationMethod field.
func (o *PostAuthenticationMethodsRequest) SetPreferredAuthenticationMethod(v PutAuthenticationMethodsRequestInnerPreferredAuthenticationMethod) {
	o.PreferredAuthenticationMethod = &v
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *PostAuthenticationMethodsRequest) GetKeyId() string {
	if o == nil || IsNil(o.KeyId) {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthenticationMethodsRequest) GetKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.KeyId) {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *PostAuthenticationMethodsRequest) HasKeyId() bool {
	if o != nil && !IsNil(o.KeyId) {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *PostAuthenticationMethodsRequest) SetKeyId(v string) {
	o.KeyId = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *PostAuthenticationMethodsRequest) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthenticationMethodsRequest) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *PostAuthenticationMethodsRequest) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *PostAuthenticationMethodsRequest) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetRelyingPartyIdentifier returns the RelyingPartyIdentifier field value if set, zero value otherwise.
func (o *PostAuthenticationMethodsRequest) GetRelyingPartyIdentifier() string {
	if o == nil || IsNil(o.RelyingPartyIdentifier) {
		var ret string
		return ret
	}
	return *o.RelyingPartyIdentifier
}

// GetRelyingPartyIdentifierOk returns a tuple with the RelyingPartyIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthenticationMethodsRequest) GetRelyingPartyIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.RelyingPartyIdentifier) {
		return nil, false
	}
	return o.RelyingPartyIdentifier, true
}

// HasRelyingPartyIdentifier returns a boolean if a field has been set.
func (o *PostAuthenticationMethodsRequest) HasRelyingPartyIdentifier() bool {
	if o != nil && !IsNil(o.RelyingPartyIdentifier) {
		return true
	}

	return false
}

// SetRelyingPartyIdentifier gets a reference to the given string and assigns it to the RelyingPartyIdentifier field.
func (o *PostAuthenticationMethodsRequest) SetRelyingPartyIdentifier(v string) {
	o.RelyingPartyIdentifier = &v
}

func (o PostAuthenticationMethodsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAuthenticationMethodsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TotpSecret) {
		toSerialize["totp_secret"] = o.TotpSecret
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.PreferredAuthenticationMethod) {
		toSerialize["preferred_authentication_method"] = o.PreferredAuthenticationMethod
	}
	if !IsNil(o.KeyId) {
		toSerialize["key_id"] = o.KeyId
	}
	if !IsNil(o.PublicKey) {
		toSerialize["public_key"] = o.PublicKey
	}
	if !IsNil(o.RelyingPartyIdentifier) {
		toSerialize["relying_party_identifier"] = o.RelyingPartyIdentifier
	}
	return toSerialize, nil
}

func (o *PostAuthenticationMethodsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varPostAuthenticationMethodsRequest := _PostAuthenticationMethodsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varPostAuthenticationMethodsRequest)

	if err != nil {
		return err
	}

	*o = PostAuthenticationMethodsRequest(varPostAuthenticationMethodsRequest)

	return err
}

type NullablePostAuthenticationMethodsRequest struct {
	value *PostAuthenticationMethodsRequest
	isSet bool
}

func (v NullablePostAuthenticationMethodsRequest) Get() *PostAuthenticationMethodsRequest {
	return v.value
}

func (v *NullablePostAuthenticationMethodsRequest) Set(val *PostAuthenticationMethodsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAuthenticationMethodsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAuthenticationMethodsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAuthenticationMethodsRequest(val *PostAuthenticationMethodsRequest) *NullablePostAuthenticationMethodsRequest {
	return &NullablePostAuthenticationMethodsRequest{value: val, isSet: true}
}

func (v NullablePostAuthenticationMethodsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAuthenticationMethodsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
