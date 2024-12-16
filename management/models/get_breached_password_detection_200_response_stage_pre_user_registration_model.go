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

// GetBreachedPasswordDetection200ResponseStagePreUserRegistration struct for GetBreachedPasswordDetection200ResponseStagePreUserRegistration
type GetBreachedPasswordDetection200ResponseStagePreUserRegistration struct {
	// Action to take when a breached password is detected during a signup.               Possible values: <code>block</code>, <code>admin_notification</code>.
	Shields []GetBreachedPasswordDetection200ResponseStagePreUserRegistrationShieldsInner `json:"shields"`
}

type _GetBreachedPasswordDetection200ResponseStagePreUserRegistration GetBreachedPasswordDetection200ResponseStagePreUserRegistration

// GetShields returns the Shields field value
func (o *GetBreachedPasswordDetection200ResponseStagePreUserRegistration) GetShields() []GetBreachedPasswordDetection200ResponseStagePreUserRegistrationShieldsInner {
	if o == nil {
		var ret []GetBreachedPasswordDetection200ResponseStagePreUserRegistrationShieldsInner
		return ret
	}

	return o.Shields
}

// GetShieldsOk returns a tuple with the Shields field value
// and a boolean to check if the value has been set.
func (o *GetBreachedPasswordDetection200ResponseStagePreUserRegistration) GetShieldsOk() ([]GetBreachedPasswordDetection200ResponseStagePreUserRegistrationShieldsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shields, true
}

// SetShields sets field value
func (o *GetBreachedPasswordDetection200ResponseStagePreUserRegistration) SetShields(v []GetBreachedPasswordDetection200ResponseStagePreUserRegistrationShieldsInner) {
	o.Shields = v
}

func (o GetBreachedPasswordDetection200ResponseStagePreUserRegistration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBreachedPasswordDetection200ResponseStagePreUserRegistration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shields"] = o.Shields
	return toSerialize, nil
}

func (o *GetBreachedPasswordDetection200ResponseStagePreUserRegistration) UnmarshalJSON(data []byte) (err error) {
	varGetBreachedPasswordDetection200ResponseStagePreUserRegistration := _GetBreachedPasswordDetection200ResponseStagePreUserRegistration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetBreachedPasswordDetection200ResponseStagePreUserRegistration)

	if err != nil {
		return err
	}

	*o = GetBreachedPasswordDetection200ResponseStagePreUserRegistration(varGetBreachedPasswordDetection200ResponseStagePreUserRegistration)

	return err
}

type NullableGetBreachedPasswordDetection200ResponseStagePreUserRegistration struct {
	value *GetBreachedPasswordDetection200ResponseStagePreUserRegistration
	isSet bool
}

func (v NullableGetBreachedPasswordDetection200ResponseStagePreUserRegistration) Get() *GetBreachedPasswordDetection200ResponseStagePreUserRegistration {
	return v.value
}

func (v *NullableGetBreachedPasswordDetection200ResponseStagePreUserRegistration) Set(val *GetBreachedPasswordDetection200ResponseStagePreUserRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBreachedPasswordDetection200ResponseStagePreUserRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBreachedPasswordDetection200ResponseStagePreUserRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBreachedPasswordDetection200ResponseStagePreUserRegistration(val *GetBreachedPasswordDetection200ResponseStagePreUserRegistration) *NullableGetBreachedPasswordDetection200ResponseStagePreUserRegistration {
	return &NullableGetBreachedPasswordDetection200ResponseStagePreUserRegistration{value: val, isSet: true}
}

func (v NullableGetBreachedPasswordDetection200ResponseStagePreUserRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBreachedPasswordDetection200ResponseStagePreUserRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
