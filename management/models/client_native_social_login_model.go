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

// ClientNativeSocialLogin Configure native social settings
type ClientNativeSocialLogin struct {
	Apple    NullableClientNativeSocialLoginApple    `json:"apple"`
	Facebook NullableClientNativeSocialLoginFacebook `json:"facebook"`
}

type _ClientNativeSocialLogin ClientNativeSocialLogin

// GetApple returns the Apple field value
// If the value is explicit nil, the zero value for ClientNativeSocialLoginApple will be returned
func (o *ClientNativeSocialLogin) GetApple() ClientNativeSocialLoginApple {
	if o == nil || o.Apple.Get() == nil {
		var ret ClientNativeSocialLoginApple
		return ret
	}

	return *o.Apple.Get()
}

// GetAppleOk returns a tuple with the Apple field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientNativeSocialLogin) GetAppleOk() (*ClientNativeSocialLoginApple, bool) {
	if o == nil {
		return nil, false
	}
	return o.Apple.Get(), o.Apple.IsSet()
}

// SetApple sets field value
func (o *ClientNativeSocialLogin) SetApple(v ClientNativeSocialLoginApple) {
	o.Apple.Set(&v)
}

// GetFacebook returns the Facebook field value
// If the value is explicit nil, the zero value for ClientNativeSocialLoginFacebook will be returned
func (o *ClientNativeSocialLogin) GetFacebook() ClientNativeSocialLoginFacebook {
	if o == nil || o.Facebook.Get() == nil {
		var ret ClientNativeSocialLoginFacebook
		return ret
	}

	return *o.Facebook.Get()
}

// GetFacebookOk returns a tuple with the Facebook field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientNativeSocialLogin) GetFacebookOk() (*ClientNativeSocialLoginFacebook, bool) {
	if o == nil {
		return nil, false
	}
	return o.Facebook.Get(), o.Facebook.IsSet()
}

// SetFacebook sets field value
func (o *ClientNativeSocialLogin) SetFacebook(v ClientNativeSocialLoginFacebook) {
	o.Facebook.Set(&v)
}

func (o ClientNativeSocialLogin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientNativeSocialLogin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apple"] = o.Apple.Get()
	toSerialize["facebook"] = o.Facebook.Get()
	return toSerialize, nil
}

func (o *ClientNativeSocialLogin) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apple",
		"facebook",
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

	varClientNativeSocialLogin := _ClientNativeSocialLogin{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClientNativeSocialLogin)

	if err != nil {
		return err
	}

	*o = ClientNativeSocialLogin(varClientNativeSocialLogin)

	return err
}

type NullableClientNativeSocialLogin struct {
	value *ClientNativeSocialLogin
	isSet bool
}

func (v NullableClientNativeSocialLogin) Get() *ClientNativeSocialLogin {
	return v.value
}

func (v *NullableClientNativeSocialLogin) Set(val *ClientNativeSocialLogin) {
	v.value = val
	v.isSet = true
}

func (v NullableClientNativeSocialLogin) IsSet() bool {
	return v.isSet
}

func (v *NullableClientNativeSocialLogin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientNativeSocialLogin(val *ClientNativeSocialLogin) *NullableClientNativeSocialLogin {
	return &NullableClientNativeSocialLogin{value: val, isSet: true}
}

func (v NullableClientNativeSocialLogin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientNativeSocialLogin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
