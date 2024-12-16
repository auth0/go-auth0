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

// PostScimToken201Response struct for PostScimToken201Response
type PostScimToken201Response struct {
	// The token's identifier
	TokenId string `json:"token_id"`
	// The scim client's token
	Token string `json:"token"`
	// The scopes of the scim token
	Scopes []string `json:"scopes"`
	// The token's created at timestamp
	CreatedAt string `json:"created_at"`
	// The token's valid until at timestamp
	ValidUntil string `json:"valid_until"`
}

type _PostScimToken201Response PostScimToken201Response

// GetTokenId returns the TokenId field value
func (o *PostScimToken201Response) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *PostScimToken201Response) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *PostScimToken201Response) SetTokenId(v string) {
	o.TokenId = v
}

// GetToken returns the Token field value
func (o *PostScimToken201Response) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *PostScimToken201Response) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *PostScimToken201Response) SetToken(v string) {
	o.Token = v
}

// GetScopes returns the Scopes field value
func (o *PostScimToken201Response) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *PostScimToken201Response) GetScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *PostScimToken201Response) SetScopes(v []string) {
	o.Scopes = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PostScimToken201Response) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PostScimToken201Response) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PostScimToken201Response) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetValidUntil returns the ValidUntil field value
func (o *PostScimToken201Response) GetValidUntil() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value
// and a boolean to check if the value has been set.
func (o *PostScimToken201Response) GetValidUntilOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidUntil, true
}

// SetValidUntil sets field value
func (o *PostScimToken201Response) SetValidUntil(v string) {
	o.ValidUntil = v
}

func (o PostScimToken201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostScimToken201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token_id"] = o.TokenId
	toSerialize["token"] = o.Token
	toSerialize["scopes"] = o.Scopes
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["valid_until"] = o.ValidUntil
	return toSerialize, nil
}

func (o *PostScimToken201Response) UnmarshalJSON(data []byte) (err error) {
	varPostScimToken201Response := _PostScimToken201Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostScimToken201Response)

	if err != nil {
		return err
	}

	*o = PostScimToken201Response(varPostScimToken201Response)

	return err
}

type NullablePostScimToken201Response struct {
	value *PostScimToken201Response
	isSet bool
}

func (v NullablePostScimToken201Response) Get() *PostScimToken201Response {
	return v.value
}

func (v *NullablePostScimToken201Response) Set(val *PostScimToken201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostScimToken201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostScimToken201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostScimToken201Response(val *PostScimToken201Response) *NullablePostScimToken201Response {
	return &NullablePostScimToken201Response{value: val, isSet: true}
}

func (v NullablePostScimToken201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostScimToken201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
