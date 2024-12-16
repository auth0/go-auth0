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

// PostRoleUsersRequest struct for PostRoleUsersRequest
type PostRoleUsersRequest struct {
	// user_id's of the users to assign the role to.
	Users []string `json:"users"`
}

type _PostRoleUsersRequest PostRoleUsersRequest

// GetUsers returns the Users field value
func (o *PostRoleUsersRequest) GetUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *PostRoleUsersRequest) GetUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *PostRoleUsersRequest) SetUsers(v []string) {
	o.Users = v
}

func (o PostRoleUsersRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostRoleUsersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

func (o *PostRoleUsersRequest) UnmarshalJSON(data []byte) (err error) {
	varPostRoleUsersRequest := _PostRoleUsersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostRoleUsersRequest)

	if err != nil {
		return err
	}

	*o = PostRoleUsersRequest(varPostRoleUsersRequest)

	return err
}

type NullablePostRoleUsersRequest struct {
	value *PostRoleUsersRequest
	isSet bool
}

func (v NullablePostRoleUsersRequest) Get() *PostRoleUsersRequest {
	return v.value
}

func (v *NullablePostRoleUsersRequest) Set(val *PostRoleUsersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostRoleUsersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostRoleUsersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostRoleUsersRequest(val *PostRoleUsersRequest) *NullablePostRoleUsersRequest {
	return &NullablePostRoleUsersRequest{value: val, isSet: true}
}

func (v NullablePostRoleUsersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostRoleUsersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
