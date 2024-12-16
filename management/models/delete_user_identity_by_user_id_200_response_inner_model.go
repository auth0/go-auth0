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

// DeleteUserIdentityByUserId200ResponseInner struct for DeleteUserIdentityByUserId200ResponseInner
type DeleteUserIdentityByUserId200ResponseInner struct {
	// The name of the connection for the identity.
	Connection string `json:"connection"`
	// The unique identifier for the user for the identity.
	UserId string `json:"user_id"`
	// The type of identity provider.
	Provider string `json:"provider"`
	// <code>true</code> if the identity provider is a social provider, <code>false</code>s otherwise
	IsSocial *bool `json:"isSocial,omitempty"`
	// IDP access token returned only if scope read:user_idp_tokens is defined
	AccessToken *string `json:"access_token,omitempty"`
	// IDP access token secret returned only if scope read:user_idp_tokens is defined.
	AccessTokenSecret *string `json:"access_token_secret,omitempty"`
	// IDP refresh token returned only if scope read:user_idp_tokens is defined.
	RefreshToken *string      `json:"refresh_token,omitempty"`
	ProfileData  *UserProfile `json:"profileData,omitempty"`
}

type _DeleteUserIdentityByUserId200ResponseInner DeleteUserIdentityByUserId200ResponseInner

// GetConnection returns the Connection field value
func (o *DeleteUserIdentityByUserId200ResponseInner) GetConnection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value
// and a boolean to check if the value has been set.
func (o *DeleteUserIdentityByUserId200ResponseInner) GetConnectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connection, true
}

// SetConnection sets field value
func (o *DeleteUserIdentityByUserId200ResponseInner) SetConnection(v string) {
	o.Connection = v
}

// GetUserId returns the UserId field value
func (o *DeleteUserIdentityByUserId200ResponseInner) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *DeleteUserIdentityByUserId200ResponseInner) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *DeleteUserIdentityByUserId200ResponseInner) SetUserId(v string) {
	o.UserId = v
}

// GetProvider returns the Provider field value
func (o *DeleteUserIdentityByUserId200ResponseInner) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *DeleteUserIdentityByUserId200ResponseInner) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *DeleteUserIdentityByUserId200ResponseInner) SetProvider(v string) {
	o.Provider = v
}

// GetIsSocial returns the IsSocial field value if set, zero value otherwise.
func (o *DeleteUserIdentityByUserId200ResponseInner) GetIsSocial() bool {
	if o == nil || IsNil(o.IsSocial) {
		var ret bool
		return ret
	}
	return *o.IsSocial
}

// GetIsSocialOk returns a tuple with the IsSocial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteUserIdentityByUserId200ResponseInner) GetIsSocialOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSocial) {
		return nil, false
	}
	return o.IsSocial, true
}

// HasIsSocial returns a boolean if a field has been set.
func (o *DeleteUserIdentityByUserId200ResponseInner) HasIsSocial() bool {
	if o != nil && !IsNil(o.IsSocial) {
		return true
	}

	return false
}

// SetIsSocial gets a reference to the given bool and assigns it to the IsSocial field.
func (o *DeleteUserIdentityByUserId200ResponseInner) SetIsSocial(v bool) {
	o.IsSocial = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *DeleteUserIdentityByUserId200ResponseInner) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteUserIdentityByUserId200ResponseInner) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *DeleteUserIdentityByUserId200ResponseInner) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *DeleteUserIdentityByUserId200ResponseInner) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetAccessTokenSecret returns the AccessTokenSecret field value if set, zero value otherwise.
func (o *DeleteUserIdentityByUserId200ResponseInner) GetAccessTokenSecret() string {
	if o == nil || IsNil(o.AccessTokenSecret) {
		var ret string
		return ret
	}
	return *o.AccessTokenSecret
}

// GetAccessTokenSecretOk returns a tuple with the AccessTokenSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteUserIdentityByUserId200ResponseInner) GetAccessTokenSecretOk() (*string, bool) {
	if o == nil || IsNil(o.AccessTokenSecret) {
		return nil, false
	}
	return o.AccessTokenSecret, true
}

// HasAccessTokenSecret returns a boolean if a field has been set.
func (o *DeleteUserIdentityByUserId200ResponseInner) HasAccessTokenSecret() bool {
	if o != nil && !IsNil(o.AccessTokenSecret) {
		return true
	}

	return false
}

// SetAccessTokenSecret gets a reference to the given string and assigns it to the AccessTokenSecret field.
func (o *DeleteUserIdentityByUserId200ResponseInner) SetAccessTokenSecret(v string) {
	o.AccessTokenSecret = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *DeleteUserIdentityByUserId200ResponseInner) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteUserIdentityByUserId200ResponseInner) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *DeleteUserIdentityByUserId200ResponseInner) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *DeleteUserIdentityByUserId200ResponseInner) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetProfileData returns the ProfileData field value if set, zero value otherwise.
func (o *DeleteUserIdentityByUserId200ResponseInner) GetProfileData() UserProfile {
	if o == nil || IsNil(o.ProfileData) {
		var ret UserProfile
		return ret
	}
	return *o.ProfileData
}

// GetProfileDataOk returns a tuple with the ProfileData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteUserIdentityByUserId200ResponseInner) GetProfileDataOk() (*UserProfile, bool) {
	if o == nil || IsNil(o.ProfileData) {
		return nil, false
	}
	return o.ProfileData, true
}

// HasProfileData returns a boolean if a field has been set.
func (o *DeleteUserIdentityByUserId200ResponseInner) HasProfileData() bool {
	if o != nil && !IsNil(o.ProfileData) {
		return true
	}

	return false
}

// SetProfileData gets a reference to the given UserProfile and assigns it to the ProfileData field.
func (o *DeleteUserIdentityByUserId200ResponseInner) SetProfileData(v UserProfile) {
	o.ProfileData = &v
}

func (o DeleteUserIdentityByUserId200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteUserIdentityByUserId200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connection"] = o.Connection
	toSerialize["user_id"] = o.UserId
	toSerialize["provider"] = o.Provider
	if !IsNil(o.IsSocial) {
		toSerialize["isSocial"] = o.IsSocial
	}
	if !IsNil(o.AccessToken) {
		toSerialize["access_token"] = o.AccessToken
	}
	if !IsNil(o.AccessTokenSecret) {
		toSerialize["access_token_secret"] = o.AccessTokenSecret
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	if !IsNil(o.ProfileData) {
		toSerialize["profileData"] = o.ProfileData
	}
	return toSerialize, nil
}

func (o *DeleteUserIdentityByUserId200ResponseInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connection",
		"user_id",
		"provider",
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

	varDeleteUserIdentityByUserId200ResponseInner := _DeleteUserIdentityByUserId200ResponseInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteUserIdentityByUserId200ResponseInner)

	if err != nil {
		return err
	}

	*o = DeleteUserIdentityByUserId200ResponseInner(varDeleteUserIdentityByUserId200ResponseInner)

	return err
}

type NullableDeleteUserIdentityByUserId200ResponseInner struct {
	value *DeleteUserIdentityByUserId200ResponseInner
	isSet bool
}

func (v NullableDeleteUserIdentityByUserId200ResponseInner) Get() *DeleteUserIdentityByUserId200ResponseInner {
	return v.value
}

func (v *NullableDeleteUserIdentityByUserId200ResponseInner) Set(val *DeleteUserIdentityByUserId200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteUserIdentityByUserId200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteUserIdentityByUserId200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteUserIdentityByUserId200ResponseInner(val *DeleteUserIdentityByUserId200ResponseInner) *NullableDeleteUserIdentityByUserId200ResponseInner {
	return &NullableDeleteUserIdentityByUserId200ResponseInner{value: val, isSet: true}
}

func (v NullableDeleteUserIdentityByUserId200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteUserIdentityByUserId200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
