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

// GetInvitations200ResponseOneOfInnerInvitee struct for GetInvitations200ResponseOneOfInnerInvitee
type GetInvitations200ResponseOneOfInnerInvitee struct {
	// The invitee's email.
	Email string `json:"email"`
}

type _GetInvitations200ResponseOneOfInnerInvitee GetInvitations200ResponseOneOfInnerInvitee

// GetEmail returns the Email field value
func (o *GetInvitations200ResponseOneOfInnerInvitee) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *GetInvitations200ResponseOneOfInnerInvitee) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *GetInvitations200ResponseOneOfInnerInvitee) SetEmail(v string) {
	o.Email = v
}

func (o GetInvitations200ResponseOneOfInnerInvitee) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInvitations200ResponseOneOfInnerInvitee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

func (o *GetInvitations200ResponseOneOfInnerInvitee) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varGetInvitations200ResponseOneOfInnerInvitee := _GetInvitations200ResponseOneOfInnerInvitee{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetInvitations200ResponseOneOfInnerInvitee)

	if err != nil {
		return err
	}

	*o = GetInvitations200ResponseOneOfInnerInvitee(varGetInvitations200ResponseOneOfInnerInvitee)

	return err
}

type NullableGetInvitations200ResponseOneOfInnerInvitee struct {
	value *GetInvitations200ResponseOneOfInnerInvitee
	isSet bool
}

func (v NullableGetInvitations200ResponseOneOfInnerInvitee) Get() *GetInvitations200ResponseOneOfInnerInvitee {
	return v.value
}

func (v *NullableGetInvitations200ResponseOneOfInnerInvitee) Set(val *GetInvitations200ResponseOneOfInnerInvitee) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInvitations200ResponseOneOfInnerInvitee) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInvitations200ResponseOneOfInnerInvitee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInvitations200ResponseOneOfInnerInvitee(val *GetInvitations200ResponseOneOfInnerInvitee) *NullableGetInvitations200ResponseOneOfInnerInvitee {
	return &NullableGetInvitations200ResponseOneOfInnerInvitee{value: val, isSet: true}
}

func (v NullableGetInvitations200ResponseOneOfInnerInvitee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInvitations200ResponseOneOfInnerInvitee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
