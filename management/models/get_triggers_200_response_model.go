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

// GetTriggers200Response struct for GetTriggers200Response
type GetTriggers200Response struct {
	Triggers []GetActions200ResponseActionsInnerSupportedTriggersInner `json:"triggers"`
}

type _GetTriggers200Response GetTriggers200Response

// GetTriggers returns the Triggers field value
func (o *GetTriggers200Response) GetTriggers() []GetActions200ResponseActionsInnerSupportedTriggersInner {
	if o == nil {
		var ret []GetActions200ResponseActionsInnerSupportedTriggersInner
		return ret
	}

	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value
// and a boolean to check if the value has been set.
func (o *GetTriggers200Response) GetTriggersOk() ([]GetActions200ResponseActionsInnerSupportedTriggersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Triggers, true
}

// SetTriggers sets field value
func (o *GetTriggers200Response) SetTriggers(v []GetActions200ResponseActionsInnerSupportedTriggersInner) {
	o.Triggers = v
}

func (o GetTriggers200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTriggers200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["triggers"] = o.Triggers
	return toSerialize, nil
}

func (o *GetTriggers200Response) UnmarshalJSON(data []byte) (err error) {
	varGetTriggers200Response := _GetTriggers200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTriggers200Response)

	if err != nil {
		return err
	}

	*o = GetTriggers200Response(varGetTriggers200Response)

	return err
}

type NullableGetTriggers200Response struct {
	value *GetTriggers200Response
	isSet bool
}

func (v NullableGetTriggers200Response) Get() *GetTriggers200Response {
	return v.value
}

func (v *NullableGetTriggers200Response) Set(val *GetTriggers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTriggers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTriggers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTriggers200Response(val *GetTriggers200Response) *NullableGetTriggers200Response {
	return &NullableGetTriggers200Response{value: val, isSet: true}
}

func (v NullableGetTriggers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTriggers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
