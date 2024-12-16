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

// GetActionVersions200ResponseVersionsInnerErrorsInner Error is a generic error with a human readable id which should be easily referenced in support tickets.
type GetActionVersions200ResponseVersionsInnerErrorsInner struct {
	Id  string `json:"id"`
	Msg string `json:"msg"`
	Url string `json:"url"`
}

type _GetActionVersions200ResponseVersionsInnerErrorsInner GetActionVersions200ResponseVersionsInnerErrorsInner

// GetId returns the Id field value
func (o *GetActionVersions200ResponseVersionsInnerErrorsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetActionVersions200ResponseVersionsInnerErrorsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetActionVersions200ResponseVersionsInnerErrorsInner) SetId(v string) {
	o.Id = v
}

// GetMsg returns the Msg field value
func (o *GetActionVersions200ResponseVersionsInnerErrorsInner) GetMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msg
}

// GetMsgOk returns a tuple with the Msg field value
// and a boolean to check if the value has been set.
func (o *GetActionVersions200ResponseVersionsInnerErrorsInner) GetMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msg, true
}

// SetMsg sets field value
func (o *GetActionVersions200ResponseVersionsInnerErrorsInner) SetMsg(v string) {
	o.Msg = v
}

// GetUrl returns the Url field value
func (o *GetActionVersions200ResponseVersionsInnerErrorsInner) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *GetActionVersions200ResponseVersionsInnerErrorsInner) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *GetActionVersions200ResponseVersionsInnerErrorsInner) SetUrl(v string) {
	o.Url = v
}

func (o GetActionVersions200ResponseVersionsInnerErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetActionVersions200ResponseVersionsInnerErrorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["msg"] = o.Msg
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *GetActionVersions200ResponseVersionsInnerErrorsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"msg",
		"url",
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

	varGetActionVersions200ResponseVersionsInnerErrorsInner := _GetActionVersions200ResponseVersionsInnerErrorsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetActionVersions200ResponseVersionsInnerErrorsInner)

	if err != nil {
		return err
	}

	*o = GetActionVersions200ResponseVersionsInnerErrorsInner(varGetActionVersions200ResponseVersionsInnerErrorsInner)

	return err
}

type NullableGetActionVersions200ResponseVersionsInnerErrorsInner struct {
	value *GetActionVersions200ResponseVersionsInnerErrorsInner
	isSet bool
}

func (v NullableGetActionVersions200ResponseVersionsInnerErrorsInner) Get() *GetActionVersions200ResponseVersionsInnerErrorsInner {
	return v.value
}

func (v *NullableGetActionVersions200ResponseVersionsInnerErrorsInner) Set(val *GetActionVersions200ResponseVersionsInnerErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetActionVersions200ResponseVersionsInnerErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetActionVersions200ResponseVersionsInnerErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetActionVersions200ResponseVersionsInnerErrorsInner(val *GetActionVersions200ResponseVersionsInnerErrorsInner) *NullableGetActionVersions200ResponseVersionsInnerErrorsInner {
	return &NullableGetActionVersions200ResponseVersionsInnerErrorsInner{value: val, isSet: true}
}

func (v NullableGetActionVersions200ResponseVersionsInnerErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetActionVersions200ResponseVersionsInnerErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
