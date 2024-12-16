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

// PostMembersRequest struct for PostMembersRequest
type PostMembersRequest struct {
	// List of user IDs to add to the organization as members.
	Members              []string `json:"members"`
	AdditionalProperties map[string]interface{}
}

type _PostMembersRequest PostMembersRequest

// GetMembers returns the Members field value
func (o *PostMembersRequest) GetMembers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *PostMembersRequest) GetMembersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *PostMembersRequest) SetMembers(v []string) {
	o.Members = v
}

func (o PostMembersRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostMembersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["members"] = o.Members

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostMembersRequest) UnmarshalJSON(data []byte) (err error) {
	varPostMembersRequest := _PostMembersRequest{}

	err = json.Unmarshal(data, &varPostMembersRequest)

	if err != nil {
		return err
	}

	*o = PostMembersRequest(varPostMembersRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "members")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostMembersRequest struct {
	value *PostMembersRequest
	isSet bool
}

func (v NullablePostMembersRequest) Get() *PostMembersRequest {
	return v.value
}

func (v *NullablePostMembersRequest) Set(val *PostMembersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostMembersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostMembersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostMembersRequest(val *PostMembersRequest) *NullablePostMembersRequest {
	return &NullablePostMembersRequest{value: val, isSet: true}
}

func (v NullablePostMembersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostMembersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
