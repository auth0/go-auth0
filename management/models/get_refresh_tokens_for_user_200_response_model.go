/*
Auth0 Management API

Auth0 Management API v2.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
)

// GetRefreshTokensForUser200Response struct for GetRefreshTokensForUser200Response
type GetRefreshTokensForUser200Response struct {
	Sessions             []GetRefreshTokensForUser200ResponseSessionsInner `json:"sessions"`
	AdditionalProperties map[string]interface{}
}

type _GetRefreshTokensForUser200Response GetRefreshTokensForUser200Response

// GetSessions returns the Sessions field value
func (o *GetRefreshTokensForUser200Response) GetSessions() []GetRefreshTokensForUser200ResponseSessionsInner {
	if o == nil {
		var ret []GetRefreshTokensForUser200ResponseSessionsInner
		return ret
	}

	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value
// and a boolean to check if the value has been set.
func (o *GetRefreshTokensForUser200Response) GetSessionsOk() ([]GetRefreshTokensForUser200ResponseSessionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sessions, true
}

// SetSessions sets field value
func (o *GetRefreshTokensForUser200Response) SetSessions(v []GetRefreshTokensForUser200ResponseSessionsInner) {
	o.Sessions = v
}

func (o GetRefreshTokensForUser200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRefreshTokensForUser200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessions"] = o.Sessions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetRefreshTokensForUser200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sessions",
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

	varGetRefreshTokensForUser200Response := _GetRefreshTokensForUser200Response{}

	err = json.Unmarshal(data, &varGetRefreshTokensForUser200Response)

	if err != nil {
		return err
	}

	*o = GetRefreshTokensForUser200Response(varGetRefreshTokensForUser200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sessions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRefreshTokensForUser200Response struct {
	value *GetRefreshTokensForUser200Response
	isSet bool
}

func (v NullableGetRefreshTokensForUser200Response) Get() *GetRefreshTokensForUser200Response {
	return v.value
}

func (v *NullableGetRefreshTokensForUser200Response) Set(val *GetRefreshTokensForUser200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRefreshTokensForUser200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRefreshTokensForUser200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRefreshTokensForUser200Response(val *GetRefreshTokensForUser200Response) *NullableGetRefreshTokensForUser200Response {
	return &NullableGetRefreshTokensForUser200Response{value: val, isSet: true}
}

func (v NullableGetRefreshTokensForUser200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRefreshTokensForUser200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
