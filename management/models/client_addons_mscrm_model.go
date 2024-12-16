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

// ClientAddonsMscrm Microsoft Dynamics CRM SSO configuration.
type ClientAddonsMscrm struct {
	// Microsoft Dynamics CRM application URL.
	Url                  string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _ClientAddonsMscrm ClientAddonsMscrm

// GetUrl returns the Url field value
func (o *ClientAddonsMscrm) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ClientAddonsMscrm) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ClientAddonsMscrm) SetUrl(v string) {
	o.Url = v
}

func (o ClientAddonsMscrm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientAddonsMscrm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientAddonsMscrm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varClientAddonsMscrm := _ClientAddonsMscrm{}

	err = json.Unmarshal(data, &varClientAddonsMscrm)

	if err != nil {
		return err
	}

	*o = ClientAddonsMscrm(varClientAddonsMscrm)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientAddonsMscrm struct {
	value *ClientAddonsMscrm
	isSet bool
}

func (v NullableClientAddonsMscrm) Get() *ClientAddonsMscrm {
	return v.value
}

func (v *NullableClientAddonsMscrm) Set(val *ClientAddonsMscrm) {
	v.value = val
	v.isSet = true
}

func (v NullableClientAddonsMscrm) IsSet() bool {
	return v.isSet
}

func (v *NullableClientAddonsMscrm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientAddonsMscrm(val *ClientAddonsMscrm) *NullableClientAddonsMscrm {
	return &NullableClientAddonsMscrm{value: val, isSet: true}
}

func (v NullableClientAddonsMscrm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientAddonsMscrm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
