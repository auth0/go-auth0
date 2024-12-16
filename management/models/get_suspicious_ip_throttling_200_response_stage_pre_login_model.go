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

// GetSuspiciousIpThrottling200ResponseStagePreLogin Configuration options that apply before every login attempt.
type GetSuspiciousIpThrottling200ResponseStagePreLogin struct {
	// Total number of attempts allowed per day.
	MaxAttempts int32 `json:"max_attempts"`
	// Interval of time, given in milliseconds, at which new attempts are granted.
	Rate int32 `json:"rate"`
}

type _GetSuspiciousIpThrottling200ResponseStagePreLogin GetSuspiciousIpThrottling200ResponseStagePreLogin

// GetMaxAttempts returns the MaxAttempts field value
func (o *GetSuspiciousIpThrottling200ResponseStagePreLogin) GetMaxAttempts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxAttempts
}

// GetMaxAttemptsOk returns a tuple with the MaxAttempts field value
// and a boolean to check if the value has been set.
func (o *GetSuspiciousIpThrottling200ResponseStagePreLogin) GetMaxAttemptsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxAttempts, true
}

// SetMaxAttempts sets field value
func (o *GetSuspiciousIpThrottling200ResponseStagePreLogin) SetMaxAttempts(v int32) {
	o.MaxAttempts = v
}

// GetRate returns the Rate field value
func (o *GetSuspiciousIpThrottling200ResponseStagePreLogin) GetRate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *GetSuspiciousIpThrottling200ResponseStagePreLogin) GetRateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value
func (o *GetSuspiciousIpThrottling200ResponseStagePreLogin) SetRate(v int32) {
	o.Rate = v
}

func (o GetSuspiciousIpThrottling200ResponseStagePreLogin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSuspiciousIpThrottling200ResponseStagePreLogin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["max_attempts"] = o.MaxAttempts
	toSerialize["rate"] = o.Rate
	return toSerialize, nil
}

func (o *GetSuspiciousIpThrottling200ResponseStagePreLogin) UnmarshalJSON(data []byte) (err error) {
	varGetSuspiciousIpThrottling200ResponseStagePreLogin := _GetSuspiciousIpThrottling200ResponseStagePreLogin{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSuspiciousIpThrottling200ResponseStagePreLogin)

	if err != nil {
		return err
	}

	*o = GetSuspiciousIpThrottling200ResponseStagePreLogin(varGetSuspiciousIpThrottling200ResponseStagePreLogin)

	return err
}

type NullableGetSuspiciousIpThrottling200ResponseStagePreLogin struct {
	value *GetSuspiciousIpThrottling200ResponseStagePreLogin
	isSet bool
}

func (v NullableGetSuspiciousIpThrottling200ResponseStagePreLogin) Get() *GetSuspiciousIpThrottling200ResponseStagePreLogin {
	return v.value
}

func (v *NullableGetSuspiciousIpThrottling200ResponseStagePreLogin) Set(val *GetSuspiciousIpThrottling200ResponseStagePreLogin) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSuspiciousIpThrottling200ResponseStagePreLogin) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSuspiciousIpThrottling200ResponseStagePreLogin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSuspiciousIpThrottling200ResponseStagePreLogin(val *GetSuspiciousIpThrottling200ResponseStagePreLogin) *NullableGetSuspiciousIpThrottling200ResponseStagePreLogin {
	return &NullableGetSuspiciousIpThrottling200ResponseStagePreLogin{value: val, isSet: true}
}

func (v NullableGetSuspiciousIpThrottling200ResponseStagePreLogin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSuspiciousIpThrottling200ResponseStagePreLogin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
