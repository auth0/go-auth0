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

// SmsTwilioFactorProvider struct for SmsTwilioFactorProvider
type SmsTwilioFactorProvider struct {
	// From number
	From NullableString `json:"from,omitempty"`
	// Copilot SID
	MessagingServiceSid NullableString `json:"messaging_service_sid,omitempty"`
	// Twilio Authentication token
	AuthToken NullableString `json:"auth_token,omitempty"`
	// Twilio SID
	Sid NullableString `json:"sid,omitempty"`
}

// GetFrom returns the From field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmsTwilioFactorProvider) GetFrom() string {
	if o == nil || IsNil(o.From.Get()) {
		var ret string
		return ret
	}
	return *o.From.Get()
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmsTwilioFactorProvider) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.From.Get(), o.From.IsSet()
}

// HasFrom returns a boolean if a field has been set.
func (o *SmsTwilioFactorProvider) HasFrom() bool {
	if o != nil && o.From.IsSet() {
		return true
	}

	return false
}

// SetFrom gets a reference to the given NullableString and assigns it to the From field.
func (o *SmsTwilioFactorProvider) SetFrom(v string) {
	o.From.Set(&v)
}

// SetFromNil sets the value for From to be an explicit nil
func (o *SmsTwilioFactorProvider) SetFromNil() {
	o.From.Set(nil)
}

// UnsetFrom ensures that no value is present for From, not even an explicit nil
func (o *SmsTwilioFactorProvider) UnsetFrom() {
	o.From.Unset()
}

// GetMessagingServiceSid returns the MessagingServiceSid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmsTwilioFactorProvider) GetMessagingServiceSid() string {
	if o == nil || IsNil(o.MessagingServiceSid.Get()) {
		var ret string
		return ret
	}
	return *o.MessagingServiceSid.Get()
}

// GetMessagingServiceSidOk returns a tuple with the MessagingServiceSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmsTwilioFactorProvider) GetMessagingServiceSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessagingServiceSid.Get(), o.MessagingServiceSid.IsSet()
}

// HasMessagingServiceSid returns a boolean if a field has been set.
func (o *SmsTwilioFactorProvider) HasMessagingServiceSid() bool {
	if o != nil && o.MessagingServiceSid.IsSet() {
		return true
	}

	return false
}

// SetMessagingServiceSid gets a reference to the given NullableString and assigns it to the MessagingServiceSid field.
func (o *SmsTwilioFactorProvider) SetMessagingServiceSid(v string) {
	o.MessagingServiceSid.Set(&v)
}

// SetMessagingServiceSidNil sets the value for MessagingServiceSid to be an explicit nil
func (o *SmsTwilioFactorProvider) SetMessagingServiceSidNil() {
	o.MessagingServiceSid.Set(nil)
}

// UnsetMessagingServiceSid ensures that no value is present for MessagingServiceSid, not even an explicit nil
func (o *SmsTwilioFactorProvider) UnsetMessagingServiceSid() {
	o.MessagingServiceSid.Unset()
}

// GetAuthToken returns the AuthToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmsTwilioFactorProvider) GetAuthToken() string {
	if o == nil || IsNil(o.AuthToken.Get()) {
		var ret string
		return ret
	}
	return *o.AuthToken.Get()
}

// GetAuthTokenOk returns a tuple with the AuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmsTwilioFactorProvider) GetAuthTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthToken.Get(), o.AuthToken.IsSet()
}

// HasAuthToken returns a boolean if a field has been set.
func (o *SmsTwilioFactorProvider) HasAuthToken() bool {
	if o != nil && o.AuthToken.IsSet() {
		return true
	}

	return false
}

// SetAuthToken gets a reference to the given NullableString and assigns it to the AuthToken field.
func (o *SmsTwilioFactorProvider) SetAuthToken(v string) {
	o.AuthToken.Set(&v)
}

// SetAuthTokenNil sets the value for AuthToken to be an explicit nil
func (o *SmsTwilioFactorProvider) SetAuthTokenNil() {
	o.AuthToken.Set(nil)
}

// UnsetAuthToken ensures that no value is present for AuthToken, not even an explicit nil
func (o *SmsTwilioFactorProvider) UnsetAuthToken() {
	o.AuthToken.Unset()
}

// GetSid returns the Sid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmsTwilioFactorProvider) GetSid() string {
	if o == nil || IsNil(o.Sid.Get()) {
		var ret string
		return ret
	}
	return *o.Sid.Get()
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmsTwilioFactorProvider) GetSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sid.Get(), o.Sid.IsSet()
}

// HasSid returns a boolean if a field has been set.
func (o *SmsTwilioFactorProvider) HasSid() bool {
	if o != nil && o.Sid.IsSet() {
		return true
	}

	return false
}

// SetSid gets a reference to the given NullableString and assigns it to the Sid field.
func (o *SmsTwilioFactorProvider) SetSid(v string) {
	o.Sid.Set(&v)
}

// SetSidNil sets the value for Sid to be an explicit nil
func (o *SmsTwilioFactorProvider) SetSidNil() {
	o.Sid.Set(nil)
}

// UnsetSid ensures that no value is present for Sid, not even an explicit nil
func (o *SmsTwilioFactorProvider) UnsetSid() {
	o.Sid.Unset()
}

func (o SmsTwilioFactorProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmsTwilioFactorProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.From.IsSet() {
		toSerialize["from"] = o.From.Get()
	}
	if o.MessagingServiceSid.IsSet() {
		toSerialize["messaging_service_sid"] = o.MessagingServiceSid.Get()
	}
	if o.AuthToken.IsSet() {
		toSerialize["auth_token"] = o.AuthToken.Get()
	}
	if o.Sid.IsSet() {
		toSerialize["sid"] = o.Sid.Get()
	}
	return toSerialize, nil
}

type NullableSmsTwilioFactorProvider struct {
	value *SmsTwilioFactorProvider
	isSet bool
}

func (v NullableSmsTwilioFactorProvider) Get() *SmsTwilioFactorProvider {
	return v.value
}

func (v *NullableSmsTwilioFactorProvider) Set(val *SmsTwilioFactorProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsTwilioFactorProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsTwilioFactorProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsTwilioFactorProvider(val *SmsTwilioFactorProvider) *NullableSmsTwilioFactorProvider {
	return &NullableSmsTwilioFactorProvider{value: val, isSet: true}
}

func (v NullableSmsTwilioFactorProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsTwilioFactorProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
