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

// GetSessionsForUser200ResponseSessionsInnerDevice Metadata related to the device used in the session
type GetSessionsForUser200ResponseSessionsInnerDevice struct {
	// First IP address associated with this session
	InitialIp string `json:"initial_ip"`
	// First autonomous system number associated with this session
	InitialAsn string `json:"initial_asn"`
	// Last user agent of the device from which this user logged in
	LastUserAgent string `json:"last_user_agent"`
	// Last IP address from which this user logged in
	LastIp string `json:"last_ip"`
	// Last autonomous system number from which this user logged in
	LastAsn              string `json:"last_asn"`
	AdditionalProperties map[string]interface{}
}

type _GetSessionsForUser200ResponseSessionsInnerDevice GetSessionsForUser200ResponseSessionsInnerDevice

// GetInitialIp returns the InitialIp field value
func (o *GetSessionsForUser200ResponseSessionsInnerDevice) GetInitialIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InitialIp
}

// GetInitialIpOk returns a tuple with the InitialIp field value
// and a boolean to check if the value has been set.
func (o *GetSessionsForUser200ResponseSessionsInnerDevice) GetInitialIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitialIp, true
}

// SetInitialIp sets field value
func (o *GetSessionsForUser200ResponseSessionsInnerDevice) SetInitialIp(v string) {
	o.InitialIp = v
}

// GetInitialAsn returns the InitialAsn field value
func (o *GetSessionsForUser200ResponseSessionsInnerDevice) GetInitialAsn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InitialAsn
}

// GetInitialAsnOk returns a tuple with the InitialAsn field value
// and a boolean to check if the value has been set.
func (o *GetSessionsForUser200ResponseSessionsInnerDevice) GetInitialAsnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitialAsn, true
}

// SetInitialAsn sets field value
func (o *GetSessionsForUser200ResponseSessionsInnerDevice) SetInitialAsn(v string) {
	o.InitialAsn = v
}

// GetLastUserAgent returns the LastUserAgent field value
func (o *GetSessionsForUser200ResponseSessionsInnerDevice) GetLastUserAgent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUserAgent
}

// GetLastUserAgentOk returns a tuple with the LastUserAgent field value
// and a boolean to check if the value has been set.
func (o *GetSessionsForUser200ResponseSessionsInnerDevice) GetLastUserAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUserAgent, true
}

// SetLastUserAgent sets field value
func (o *GetSessionsForUser200ResponseSessionsInnerDevice) SetLastUserAgent(v string) {
	o.LastUserAgent = v
}

// GetLastIp returns the LastIp field value
func (o *GetSessionsForUser200ResponseSessionsInnerDevice) GetLastIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastIp
}

// GetLastIpOk returns a tuple with the LastIp field value
// and a boolean to check if the value has been set.
func (o *GetSessionsForUser200ResponseSessionsInnerDevice) GetLastIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastIp, true
}

// SetLastIp sets field value
func (o *GetSessionsForUser200ResponseSessionsInnerDevice) SetLastIp(v string) {
	o.LastIp = v
}

// GetLastAsn returns the LastAsn field value
func (o *GetSessionsForUser200ResponseSessionsInnerDevice) GetLastAsn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastAsn
}

// GetLastAsnOk returns a tuple with the LastAsn field value
// and a boolean to check if the value has been set.
func (o *GetSessionsForUser200ResponseSessionsInnerDevice) GetLastAsnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastAsn, true
}

// SetLastAsn sets field value
func (o *GetSessionsForUser200ResponseSessionsInnerDevice) SetLastAsn(v string) {
	o.LastAsn = v
}

func (o GetSessionsForUser200ResponseSessionsInnerDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSessionsForUser200ResponseSessionsInnerDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["initial_ip"] = o.InitialIp
	toSerialize["initial_asn"] = o.InitialAsn
	toSerialize["last_user_agent"] = o.LastUserAgent
	toSerialize["last_ip"] = o.LastIp
	toSerialize["last_asn"] = o.LastAsn

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSessionsForUser200ResponseSessionsInnerDevice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"initial_ip",
		"initial_asn",
		"last_user_agent",
		"last_ip",
		"last_asn",
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

	varGetSessionsForUser200ResponseSessionsInnerDevice := _GetSessionsForUser200ResponseSessionsInnerDevice{}

	err = json.Unmarshal(data, &varGetSessionsForUser200ResponseSessionsInnerDevice)

	if err != nil {
		return err
	}

	*o = GetSessionsForUser200ResponseSessionsInnerDevice(varGetSessionsForUser200ResponseSessionsInnerDevice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "initial_ip")
		delete(additionalProperties, "initial_asn")
		delete(additionalProperties, "last_user_agent")
		delete(additionalProperties, "last_ip")
		delete(additionalProperties, "last_asn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSessionsForUser200ResponseSessionsInnerDevice struct {
	value *GetSessionsForUser200ResponseSessionsInnerDevice
	isSet bool
}

func (v NullableGetSessionsForUser200ResponseSessionsInnerDevice) Get() *GetSessionsForUser200ResponseSessionsInnerDevice {
	return v.value
}

func (v *NullableGetSessionsForUser200ResponseSessionsInnerDevice) Set(val *GetSessionsForUser200ResponseSessionsInnerDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSessionsForUser200ResponseSessionsInnerDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSessionsForUser200ResponseSessionsInnerDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSessionsForUser200ResponseSessionsInnerDevice(val *GetSessionsForUser200ResponseSessionsInnerDevice) *NullableGetSessionsForUser200ResponseSessionsInnerDevice {
	return &NullableGetSessionsForUser200ResponseSessionsInnerDevice{value: val, isSet: true}
}

func (v NullableGetSessionsForUser200ResponseSessionsInnerDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSessionsForUser200ResponseSessionsInnerDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
