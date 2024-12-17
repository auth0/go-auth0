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

// PostOrganizationMemberRolesRequest struct for PostOrganizationMemberRolesRequest
type PostOrganizationMemberRolesRequest struct {
	// List of roles IDs to associated with the user.
	Roles []string `json:"roles"`
}

type _PostOrganizationMemberRolesRequest PostOrganizationMemberRolesRequest

// GetRoles returns the Roles field value
func (o *PostOrganizationMemberRolesRequest) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *PostOrganizationMemberRolesRequest) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *PostOrganizationMemberRolesRequest) SetRoles(v []string) {
	o.Roles = v
}

func (o PostOrganizationMemberRolesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostOrganizationMemberRolesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roles"] = o.Roles
	return toSerialize, nil
}

func (o *PostOrganizationMemberRolesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"roles",
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

	varPostOrganizationMemberRolesRequest := _PostOrganizationMemberRolesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varPostOrganizationMemberRolesRequest)

	if err != nil {
		return err
	}

	*o = PostOrganizationMemberRolesRequest(varPostOrganizationMemberRolesRequest)

	return err
}

type NullablePostOrganizationMemberRolesRequest struct {
	value *PostOrganizationMemberRolesRequest
	isSet bool
}

func (v NullablePostOrganizationMemberRolesRequest) Get() *PostOrganizationMemberRolesRequest {
	return v.value
}

func (v *NullablePostOrganizationMemberRolesRequest) Set(val *PostOrganizationMemberRolesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostOrganizationMemberRolesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostOrganizationMemberRolesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostOrganizationMemberRolesRequest(val *PostOrganizationMemberRolesRequest) *NullablePostOrganizationMemberRolesRequest {
	return &NullablePostOrganizationMemberRolesRequest{value: val, isSet: true}
}

func (v NullablePostOrganizationMemberRolesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostOrganizationMemberRolesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
