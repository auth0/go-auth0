/*
Auth0 Management API

Auth0 Management API v2.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// PutAuthenticationMethods200ResponseInner The successfully created authentication method.
type PutAuthenticationMethods200ResponseInner struct {
	// The ID of the newly created authentication method (automatically generated by the application)
	Id   *string                                      `json:"id,omitempty"`
	Type PutAuthenticationMethods200ResponseInnerType `json:"type"`
	// A human-readable label to identify the authentication method.
	Name *string `json:"name,omitempty"`
	// Base32 encoded secret for TOTP generation
	TotpSecret *string `json:"totp_secret,omitempty"`
	// Applies to phone authentication methods only. The destination phone number used to send verification codes via text and voice.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// Applies to email authentication methods only. The email address used to send verification messages.
	Email                         *string                                                                `json:"email,omitempty"`
	AuthenticationMethods         []PutAuthenticationMethods200ResponseInnerAuthenticationMethodsInner   `json:"authentication_methods,omitempty"`
	PreferredAuthenticationMethod *PutAuthenticationMethods200ResponseInnerPreferredAuthenticationMethod `json:"preferred_authentication_method,omitempty"`
	// Applies to webauthn authenticators only. The id of the credential.
	KeyId *string `json:"key_id,omitempty"`
	// Applies to webauthn authenticators only. The public key.
	PublicKey *string `json:"public_key,omitempty"`
	// Applies to webauthn authenticators only. The relying party identifier.
	RelyingPartyIdentifier *string `json:"relying_party_identifier,omitempty"`
	// Authentication method creation date
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PutAuthenticationMethods200ResponseInner PutAuthenticationMethods200ResponseInner

// GetId returns the Id field value if set, zero value otherwise.
func (o *PutAuthenticationMethods200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutAuthenticationMethods200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PutAuthenticationMethods200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PutAuthenticationMethods200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value
func (o *PutAuthenticationMethods200ResponseInner) GetType() PutAuthenticationMethods200ResponseInnerType {
	if o == nil {
		var ret PutAuthenticationMethods200ResponseInnerType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PutAuthenticationMethods200ResponseInner) GetTypeOk() (*PutAuthenticationMethods200ResponseInnerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PutAuthenticationMethods200ResponseInner) SetType(v PutAuthenticationMethods200ResponseInnerType) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PutAuthenticationMethods200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutAuthenticationMethods200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PutAuthenticationMethods200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PutAuthenticationMethods200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetTotpSecret returns the TotpSecret field value if set, zero value otherwise.
func (o *PutAuthenticationMethods200ResponseInner) GetTotpSecret() string {
	if o == nil || IsNil(o.TotpSecret) {
		var ret string
		return ret
	}
	return *o.TotpSecret
}

// GetTotpSecretOk returns a tuple with the TotpSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutAuthenticationMethods200ResponseInner) GetTotpSecretOk() (*string, bool) {
	if o == nil || IsNil(o.TotpSecret) {
		return nil, false
	}
	return o.TotpSecret, true
}

// HasTotpSecret returns a boolean if a field has been set.
func (o *PutAuthenticationMethods200ResponseInner) HasTotpSecret() bool {
	if o != nil && !IsNil(o.TotpSecret) {
		return true
	}

	return false
}

// SetTotpSecret gets a reference to the given string and assigns it to the TotpSecret field.
func (o *PutAuthenticationMethods200ResponseInner) SetTotpSecret(v string) {
	o.TotpSecret = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *PutAuthenticationMethods200ResponseInner) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutAuthenticationMethods200ResponseInner) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *PutAuthenticationMethods200ResponseInner) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *PutAuthenticationMethods200ResponseInner) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PutAuthenticationMethods200ResponseInner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutAuthenticationMethods200ResponseInner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PutAuthenticationMethods200ResponseInner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PutAuthenticationMethods200ResponseInner) SetEmail(v string) {
	o.Email = &v
}

// GetAuthenticationMethods returns the AuthenticationMethods field value if set, zero value otherwise.
func (o *PutAuthenticationMethods200ResponseInner) GetAuthenticationMethods() []PutAuthenticationMethods200ResponseInnerAuthenticationMethodsInner {
	if o == nil || IsNil(o.AuthenticationMethods) {
		var ret []PutAuthenticationMethods200ResponseInnerAuthenticationMethodsInner
		return ret
	}
	return o.AuthenticationMethods
}

// GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutAuthenticationMethods200ResponseInner) GetAuthenticationMethodsOk() ([]PutAuthenticationMethods200ResponseInnerAuthenticationMethodsInner, bool) {
	if o == nil || IsNil(o.AuthenticationMethods) {
		return nil, false
	}
	return o.AuthenticationMethods, true
}

// HasAuthenticationMethods returns a boolean if a field has been set.
func (o *PutAuthenticationMethods200ResponseInner) HasAuthenticationMethods() bool {
	if o != nil && !IsNil(o.AuthenticationMethods) {
		return true
	}

	return false
}

// SetAuthenticationMethods gets a reference to the given []PutAuthenticationMethods200ResponseInnerAuthenticationMethodsInner and assigns it to the AuthenticationMethods field.
func (o *PutAuthenticationMethods200ResponseInner) SetAuthenticationMethods(v []PutAuthenticationMethods200ResponseInnerAuthenticationMethodsInner) {
	o.AuthenticationMethods = v
}

// GetPreferredAuthenticationMethod returns the PreferredAuthenticationMethod field value if set, zero value otherwise.
func (o *PutAuthenticationMethods200ResponseInner) GetPreferredAuthenticationMethod() PutAuthenticationMethods200ResponseInnerPreferredAuthenticationMethod {
	if o == nil || IsNil(o.PreferredAuthenticationMethod) {
		var ret PutAuthenticationMethods200ResponseInnerPreferredAuthenticationMethod
		return ret
	}
	return *o.PreferredAuthenticationMethod
}

// GetPreferredAuthenticationMethodOk returns a tuple with the PreferredAuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutAuthenticationMethods200ResponseInner) GetPreferredAuthenticationMethodOk() (*PutAuthenticationMethods200ResponseInnerPreferredAuthenticationMethod, bool) {
	if o == nil || IsNil(o.PreferredAuthenticationMethod) {
		return nil, false
	}
	return o.PreferredAuthenticationMethod, true
}

// HasPreferredAuthenticationMethod returns a boolean if a field has been set.
func (o *PutAuthenticationMethods200ResponseInner) HasPreferredAuthenticationMethod() bool {
	if o != nil && !IsNil(o.PreferredAuthenticationMethod) {
		return true
	}

	return false
}

// SetPreferredAuthenticationMethod gets a reference to the given PutAuthenticationMethods200ResponseInnerPreferredAuthenticationMethod and assigns it to the PreferredAuthenticationMethod field.
func (o *PutAuthenticationMethods200ResponseInner) SetPreferredAuthenticationMethod(v PutAuthenticationMethods200ResponseInnerPreferredAuthenticationMethod) {
	o.PreferredAuthenticationMethod = &v
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *PutAuthenticationMethods200ResponseInner) GetKeyId() string {
	if o == nil || IsNil(o.KeyId) {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutAuthenticationMethods200ResponseInner) GetKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.KeyId) {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *PutAuthenticationMethods200ResponseInner) HasKeyId() bool {
	if o != nil && !IsNil(o.KeyId) {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *PutAuthenticationMethods200ResponseInner) SetKeyId(v string) {
	o.KeyId = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *PutAuthenticationMethods200ResponseInner) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutAuthenticationMethods200ResponseInner) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *PutAuthenticationMethods200ResponseInner) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *PutAuthenticationMethods200ResponseInner) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetRelyingPartyIdentifier returns the RelyingPartyIdentifier field value if set, zero value otherwise.
func (o *PutAuthenticationMethods200ResponseInner) GetRelyingPartyIdentifier() string {
	if o == nil || IsNil(o.RelyingPartyIdentifier) {
		var ret string
		return ret
	}
	return *o.RelyingPartyIdentifier
}

// GetRelyingPartyIdentifierOk returns a tuple with the RelyingPartyIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutAuthenticationMethods200ResponseInner) GetRelyingPartyIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.RelyingPartyIdentifier) {
		return nil, false
	}
	return o.RelyingPartyIdentifier, true
}

// HasRelyingPartyIdentifier returns a boolean if a field has been set.
func (o *PutAuthenticationMethods200ResponseInner) HasRelyingPartyIdentifier() bool {
	if o != nil && !IsNil(o.RelyingPartyIdentifier) {
		return true
	}

	return false
}

// SetRelyingPartyIdentifier gets a reference to the given string and assigns it to the RelyingPartyIdentifier field.
func (o *PutAuthenticationMethods200ResponseInner) SetRelyingPartyIdentifier(v string) {
	o.RelyingPartyIdentifier = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PutAuthenticationMethods200ResponseInner) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutAuthenticationMethods200ResponseInner) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PutAuthenticationMethods200ResponseInner) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PutAuthenticationMethods200ResponseInner) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o PutAuthenticationMethods200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutAuthenticationMethods200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
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
	if !IsNil(o.AuthenticationMethods) {
		toSerialize["authentication_methods"] = o.AuthenticationMethods
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
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PutAuthenticationMethods200ResponseInner) UnmarshalJSON(data []byte) (err error) {
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

	varPutAuthenticationMethods200ResponseInner := _PutAuthenticationMethods200ResponseInner{}

	err = json.Unmarshal(data, &varPutAuthenticationMethods200ResponseInner)

	if err != nil {
		return err
	}

	*o = PutAuthenticationMethods200ResponseInner(varPutAuthenticationMethods200ResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "totp_secret")
		delete(additionalProperties, "phone_number")
		delete(additionalProperties, "email")
		delete(additionalProperties, "authentication_methods")
		delete(additionalProperties, "preferred_authentication_method")
		delete(additionalProperties, "key_id")
		delete(additionalProperties, "public_key")
		delete(additionalProperties, "relying_party_identifier")
		delete(additionalProperties, "created_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePutAuthenticationMethods200ResponseInner struct {
	value *PutAuthenticationMethods200ResponseInner
	isSet bool
}

func (v NullablePutAuthenticationMethods200ResponseInner) Get() *PutAuthenticationMethods200ResponseInner {
	return v.value
}

func (v *NullablePutAuthenticationMethods200ResponseInner) Set(val *PutAuthenticationMethods200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePutAuthenticationMethods200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePutAuthenticationMethods200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutAuthenticationMethods200ResponseInner(val *PutAuthenticationMethods200ResponseInner) *NullablePutAuthenticationMethods200ResponseInner {
	return &NullablePutAuthenticationMethods200ResponseInner{value: val, isSet: true}
}

func (v NullablePutAuthenticationMethods200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutAuthenticationMethods200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
