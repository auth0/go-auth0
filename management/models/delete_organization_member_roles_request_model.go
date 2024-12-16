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

// DeleteOrganizationMemberRolesRequest struct for DeleteOrganizationMemberRolesRequest
type DeleteOrganizationMemberRolesRequest struct {
	// List of roles IDs associated with the organization member to remove.
	Roles []string `json:"roles"`
}

type _DeleteOrganizationMemberRolesRequest DeleteOrganizationMemberRolesRequest

// GetRoles returns the Roles field value
func (o *DeleteOrganizationMemberRolesRequest) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *DeleteOrganizationMemberRolesRequest) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *DeleteOrganizationMemberRolesRequest) SetRoles(v []string) {
	o.Roles = v
}

func (o DeleteOrganizationMemberRolesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteOrganizationMemberRolesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roles"] = o.Roles
	return toSerialize, nil
}

func (o *DeleteOrganizationMemberRolesRequest) UnmarshalJSON(data []byte) (err error) {
	varDeleteOrganizationMemberRolesRequest := _DeleteOrganizationMemberRolesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteOrganizationMemberRolesRequest)

	if err != nil {
		return err
	}

	*o = DeleteOrganizationMemberRolesRequest(varDeleteOrganizationMemberRolesRequest)

	return err
}

type NullableDeleteOrganizationMemberRolesRequest struct {
	value *DeleteOrganizationMemberRolesRequest
	isSet bool
}

func (v NullableDeleteOrganizationMemberRolesRequest) Get() *DeleteOrganizationMemberRolesRequest {
	return v.value
}

func (v *NullableDeleteOrganizationMemberRolesRequest) Set(val *DeleteOrganizationMemberRolesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteOrganizationMemberRolesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteOrganizationMemberRolesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteOrganizationMemberRolesRequest(val *DeleteOrganizationMemberRolesRequest) *NullableDeleteOrganizationMemberRolesRequest {
	return &NullableDeleteOrganizationMemberRolesRequest{value: val, isSet: true}
}

func (v NullableDeleteOrganizationMemberRolesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteOrganizationMemberRolesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
