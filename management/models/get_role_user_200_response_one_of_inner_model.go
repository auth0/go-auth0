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

// GetRoleUser200ResponseOneOfInner struct for GetRoleUser200ResponseOneOfInner
type GetRoleUser200ResponseOneOfInner struct {
	// ID of this user.
	UserId string `json:"user_id"`
	// URL to a picture for this user.
	Picture string `json:"picture"`
	// Name of this user.
	Name string `json:"name"`
	// Email address of this user.
	Email string `json:"email"`
}

type _GetRoleUser200ResponseOneOfInner GetRoleUser200ResponseOneOfInner

// GetUserId returns the UserId field value
func (o *GetRoleUser200ResponseOneOfInner) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *GetRoleUser200ResponseOneOfInner) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *GetRoleUser200ResponseOneOfInner) SetUserId(v string) {
	o.UserId = v
}

// GetPicture returns the Picture field value
func (o *GetRoleUser200ResponseOneOfInner) GetPicture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Picture
}

// GetPictureOk returns a tuple with the Picture field value
// and a boolean to check if the value has been set.
func (o *GetRoleUser200ResponseOneOfInner) GetPictureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Picture, true
}

// SetPicture sets field value
func (o *GetRoleUser200ResponseOneOfInner) SetPicture(v string) {
	o.Picture = v
}

// GetName returns the Name field value
func (o *GetRoleUser200ResponseOneOfInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetRoleUser200ResponseOneOfInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetRoleUser200ResponseOneOfInner) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *GetRoleUser200ResponseOneOfInner) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *GetRoleUser200ResponseOneOfInner) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *GetRoleUser200ResponseOneOfInner) SetEmail(v string) {
	o.Email = v
}

func (o GetRoleUser200ResponseOneOfInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRoleUser200ResponseOneOfInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_id"] = o.UserId
	toSerialize["picture"] = o.Picture
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

func (o *GetRoleUser200ResponseOneOfInner) UnmarshalJSON(data []byte) (err error) {
	varGetRoleUser200ResponseOneOfInner := _GetRoleUser200ResponseOneOfInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetRoleUser200ResponseOneOfInner)

	if err != nil {
		return err
	}

	*o = GetRoleUser200ResponseOneOfInner(varGetRoleUser200ResponseOneOfInner)

	return err
}

type NullableGetRoleUser200ResponseOneOfInner struct {
	value *GetRoleUser200ResponseOneOfInner
	isSet bool
}

func (v NullableGetRoleUser200ResponseOneOfInner) Get() *GetRoleUser200ResponseOneOfInner {
	return v.value
}

func (v *NullableGetRoleUser200ResponseOneOfInner) Set(val *GetRoleUser200ResponseOneOfInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRoleUser200ResponseOneOfInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRoleUser200ResponseOneOfInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRoleUser200ResponseOneOfInner(val *GetRoleUser200ResponseOneOfInner) *NullableGetRoleUser200ResponseOneOfInner {
	return &NullableGetRoleUser200ResponseOneOfInner{value: val, isSet: true}
}

func (v NullableGetRoleUser200ResponseOneOfInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRoleUser200ResponseOneOfInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
