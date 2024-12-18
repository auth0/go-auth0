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

// GetRefreshToken200ResponseDevice Device used while issuing/exchanging the refresh token
type GetRefreshToken200ResponseDevice struct {
	// First IP address associated with the refresh token
	InitialIp *string `json:"initial_ip,omitempty"`
	// First autonomous system number associated with the refresh token
	InitialAsn *string `json:"initial_asn,omitempty"`
	// First user agent associated with the refresh token
	InitialUserAgent *string `json:"initial_user_agent,omitempty"`
	// Last IP address associated with the refresh token
	LastIp *string `json:"last_ip,omitempty"`
	// Last autonomous system number associated with the refresh token
	LastAsn *string `json:"last_asn,omitempty"`
	// Last user agent associated with the refresh token
	LastUserAgent        *string `json:"last_user_agent,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetRefreshToken200ResponseDevice GetRefreshToken200ResponseDevice

// GetInitialIp returns the InitialIp field value if set, zero value otherwise.
func (o *GetRefreshToken200ResponseDevice) GetInitialIp() string {
	if o == nil || IsNil(o.InitialIp) {
		var ret string
		return ret
	}
	return *o.InitialIp
}

// GetInitialIpOk returns a tuple with the InitialIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRefreshToken200ResponseDevice) GetInitialIpOk() (*string, bool) {
	if o == nil || IsNil(o.InitialIp) {
		return nil, false
	}
	return o.InitialIp, true
}

// HasInitialIp returns a boolean if a field has been set.
func (o *GetRefreshToken200ResponseDevice) HasInitialIp() bool {
	if o != nil && !IsNil(o.InitialIp) {
		return true
	}

	return false
}

// SetInitialIp gets a reference to the given string and assigns it to the InitialIp field.
func (o *GetRefreshToken200ResponseDevice) SetInitialIp(v string) {
	o.InitialIp = &v
}

// GetInitialAsn returns the InitialAsn field value if set, zero value otherwise.
func (o *GetRefreshToken200ResponseDevice) GetInitialAsn() string {
	if o == nil || IsNil(o.InitialAsn) {
		var ret string
		return ret
	}
	return *o.InitialAsn
}

// GetInitialAsnOk returns a tuple with the InitialAsn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRefreshToken200ResponseDevice) GetInitialAsnOk() (*string, bool) {
	if o == nil || IsNil(o.InitialAsn) {
		return nil, false
	}
	return o.InitialAsn, true
}

// HasInitialAsn returns a boolean if a field has been set.
func (o *GetRefreshToken200ResponseDevice) HasInitialAsn() bool {
	if o != nil && !IsNil(o.InitialAsn) {
		return true
	}

	return false
}

// SetInitialAsn gets a reference to the given string and assigns it to the InitialAsn field.
func (o *GetRefreshToken200ResponseDevice) SetInitialAsn(v string) {
	o.InitialAsn = &v
}

// GetInitialUserAgent returns the InitialUserAgent field value if set, zero value otherwise.
func (o *GetRefreshToken200ResponseDevice) GetInitialUserAgent() string {
	if o == nil || IsNil(o.InitialUserAgent) {
		var ret string
		return ret
	}
	return *o.InitialUserAgent
}

// GetInitialUserAgentOk returns a tuple with the InitialUserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRefreshToken200ResponseDevice) GetInitialUserAgentOk() (*string, bool) {
	if o == nil || IsNil(o.InitialUserAgent) {
		return nil, false
	}
	return o.InitialUserAgent, true
}

// HasInitialUserAgent returns a boolean if a field has been set.
func (o *GetRefreshToken200ResponseDevice) HasInitialUserAgent() bool {
	if o != nil && !IsNil(o.InitialUserAgent) {
		return true
	}

	return false
}

// SetInitialUserAgent gets a reference to the given string and assigns it to the InitialUserAgent field.
func (o *GetRefreshToken200ResponseDevice) SetInitialUserAgent(v string) {
	o.InitialUserAgent = &v
}

// GetLastIp returns the LastIp field value if set, zero value otherwise.
func (o *GetRefreshToken200ResponseDevice) GetLastIp() string {
	if o == nil || IsNil(o.LastIp) {
		var ret string
		return ret
	}
	return *o.LastIp
}

// GetLastIpOk returns a tuple with the LastIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRefreshToken200ResponseDevice) GetLastIpOk() (*string, bool) {
	if o == nil || IsNil(o.LastIp) {
		return nil, false
	}
	return o.LastIp, true
}

// HasLastIp returns a boolean if a field has been set.
func (o *GetRefreshToken200ResponseDevice) HasLastIp() bool {
	if o != nil && !IsNil(o.LastIp) {
		return true
	}

	return false
}

// SetLastIp gets a reference to the given string and assigns it to the LastIp field.
func (o *GetRefreshToken200ResponseDevice) SetLastIp(v string) {
	o.LastIp = &v
}

// GetLastAsn returns the LastAsn field value if set, zero value otherwise.
func (o *GetRefreshToken200ResponseDevice) GetLastAsn() string {
	if o == nil || IsNil(o.LastAsn) {
		var ret string
		return ret
	}
	return *o.LastAsn
}

// GetLastAsnOk returns a tuple with the LastAsn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRefreshToken200ResponseDevice) GetLastAsnOk() (*string, bool) {
	if o == nil || IsNil(o.LastAsn) {
		return nil, false
	}
	return o.LastAsn, true
}

// HasLastAsn returns a boolean if a field has been set.
func (o *GetRefreshToken200ResponseDevice) HasLastAsn() bool {
	if o != nil && !IsNil(o.LastAsn) {
		return true
	}

	return false
}

// SetLastAsn gets a reference to the given string and assigns it to the LastAsn field.
func (o *GetRefreshToken200ResponseDevice) SetLastAsn(v string) {
	o.LastAsn = &v
}

// GetLastUserAgent returns the LastUserAgent field value if set, zero value otherwise.
func (o *GetRefreshToken200ResponseDevice) GetLastUserAgent() string {
	if o == nil || IsNil(o.LastUserAgent) {
		var ret string
		return ret
	}
	return *o.LastUserAgent
}

// GetLastUserAgentOk returns a tuple with the LastUserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRefreshToken200ResponseDevice) GetLastUserAgentOk() (*string, bool) {
	if o == nil || IsNil(o.LastUserAgent) {
		return nil, false
	}
	return o.LastUserAgent, true
}

// HasLastUserAgent returns a boolean if a field has been set.
func (o *GetRefreshToken200ResponseDevice) HasLastUserAgent() bool {
	if o != nil && !IsNil(o.LastUserAgent) {
		return true
	}

	return false
}

// SetLastUserAgent gets a reference to the given string and assigns it to the LastUserAgent field.
func (o *GetRefreshToken200ResponseDevice) SetLastUserAgent(v string) {
	o.LastUserAgent = &v
}

func (o GetRefreshToken200ResponseDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRefreshToken200ResponseDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InitialIp) {
		toSerialize["initial_ip"] = o.InitialIp
	}
	if !IsNil(o.InitialAsn) {
		toSerialize["initial_asn"] = o.InitialAsn
	}
	if !IsNil(o.InitialUserAgent) {
		toSerialize["initial_user_agent"] = o.InitialUserAgent
	}
	if !IsNil(o.LastIp) {
		toSerialize["last_ip"] = o.LastIp
	}
	if !IsNil(o.LastAsn) {
		toSerialize["last_asn"] = o.LastAsn
	}
	if !IsNil(o.LastUserAgent) {
		toSerialize["last_user_agent"] = o.LastUserAgent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetRefreshToken200ResponseDevice) UnmarshalJSON(data []byte) (err error) {
	varGetRefreshToken200ResponseDevice := _GetRefreshToken200ResponseDevice{}

	err = json.Unmarshal(data, &varGetRefreshToken200ResponseDevice)

	if err != nil {
		return err
	}

	*o = GetRefreshToken200ResponseDevice(varGetRefreshToken200ResponseDevice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "initial_ip")
		delete(additionalProperties, "initial_asn")
		delete(additionalProperties, "initial_user_agent")
		delete(additionalProperties, "last_ip")
		delete(additionalProperties, "last_asn")
		delete(additionalProperties, "last_user_agent")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRefreshToken200ResponseDevice struct {
	value *GetRefreshToken200ResponseDevice
	isSet bool
}

func (v NullableGetRefreshToken200ResponseDevice) Get() *GetRefreshToken200ResponseDevice {
	return v.value
}

func (v *NullableGetRefreshToken200ResponseDevice) Set(val *GetRefreshToken200ResponseDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRefreshToken200ResponseDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRefreshToken200ResponseDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRefreshToken200ResponseDevice(val *GetRefreshToken200ResponseDevice) *NullableGetRefreshToken200ResponseDevice {
	return &NullableGetRefreshToken200ResponseDevice{value: val, isSet: true}
}

func (v NullableGetRefreshToken200ResponseDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRefreshToken200ResponseDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
