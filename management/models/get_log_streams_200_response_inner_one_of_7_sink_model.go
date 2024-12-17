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

// GetLogStreams200ResponseInnerOneOf7Sink struct for GetLogStreams200ResponseInnerOneOf7Sink
type GetLogStreams200ResponseInnerOneOf7Sink struct {
	MixpanelRegion GetLogStreams200ResponseInnerOneOf7SinkMixpanelRegion `json:"mixpanelRegion"`
	// Mixpanel Project Id
	MixpanelProjectId string `json:"mixpanelProjectId"`
	// Mixpanel Service Account Username
	MixpanelServiceAccountUsername string `json:"mixpanelServiceAccountUsername"`
	// Mixpanel Service Account Password
	MixpanelServiceAccountPassword string `json:"mixpanelServiceAccountPassword"`
}

type _GetLogStreams200ResponseInnerOneOf7Sink GetLogStreams200ResponseInnerOneOf7Sink

// GetMixpanelRegion returns the MixpanelRegion field value
func (o *GetLogStreams200ResponseInnerOneOf7Sink) GetMixpanelRegion() GetLogStreams200ResponseInnerOneOf7SinkMixpanelRegion {
	if o == nil {
		var ret GetLogStreams200ResponseInnerOneOf7SinkMixpanelRegion
		return ret
	}

	return o.MixpanelRegion
}

// GetMixpanelRegionOk returns a tuple with the MixpanelRegion field value
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf7Sink) GetMixpanelRegionOk() (*GetLogStreams200ResponseInnerOneOf7SinkMixpanelRegion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MixpanelRegion, true
}

// SetMixpanelRegion sets field value
func (o *GetLogStreams200ResponseInnerOneOf7Sink) SetMixpanelRegion(v GetLogStreams200ResponseInnerOneOf7SinkMixpanelRegion) {
	o.MixpanelRegion = v
}

// GetMixpanelProjectId returns the MixpanelProjectId field value
func (o *GetLogStreams200ResponseInnerOneOf7Sink) GetMixpanelProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MixpanelProjectId
}

// GetMixpanelProjectIdOk returns a tuple with the MixpanelProjectId field value
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf7Sink) GetMixpanelProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MixpanelProjectId, true
}

// SetMixpanelProjectId sets field value
func (o *GetLogStreams200ResponseInnerOneOf7Sink) SetMixpanelProjectId(v string) {
	o.MixpanelProjectId = v
}

// GetMixpanelServiceAccountUsername returns the MixpanelServiceAccountUsername field value
func (o *GetLogStreams200ResponseInnerOneOf7Sink) GetMixpanelServiceAccountUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MixpanelServiceAccountUsername
}

// GetMixpanelServiceAccountUsernameOk returns a tuple with the MixpanelServiceAccountUsername field value
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf7Sink) GetMixpanelServiceAccountUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MixpanelServiceAccountUsername, true
}

// SetMixpanelServiceAccountUsername sets field value
func (o *GetLogStreams200ResponseInnerOneOf7Sink) SetMixpanelServiceAccountUsername(v string) {
	o.MixpanelServiceAccountUsername = v
}

// GetMixpanelServiceAccountPassword returns the MixpanelServiceAccountPassword field value
func (o *GetLogStreams200ResponseInnerOneOf7Sink) GetMixpanelServiceAccountPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MixpanelServiceAccountPassword
}

// GetMixpanelServiceAccountPasswordOk returns a tuple with the MixpanelServiceAccountPassword field value
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf7Sink) GetMixpanelServiceAccountPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MixpanelServiceAccountPassword, true
}

// SetMixpanelServiceAccountPassword sets field value
func (o *GetLogStreams200ResponseInnerOneOf7Sink) SetMixpanelServiceAccountPassword(v string) {
	o.MixpanelServiceAccountPassword = v
}

func (o GetLogStreams200ResponseInnerOneOf7Sink) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLogStreams200ResponseInnerOneOf7Sink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mixpanelRegion"] = o.MixpanelRegion
	toSerialize["mixpanelProjectId"] = o.MixpanelProjectId
	toSerialize["mixpanelServiceAccountUsername"] = o.MixpanelServiceAccountUsername
	toSerialize["mixpanelServiceAccountPassword"] = o.MixpanelServiceAccountPassword
	return toSerialize, nil
}

func (o *GetLogStreams200ResponseInnerOneOf7Sink) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mixpanelRegion",
		"mixpanelProjectId",
		"mixpanelServiceAccountUsername",
		"mixpanelServiceAccountPassword",
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

	varGetLogStreams200ResponseInnerOneOf7Sink := _GetLogStreams200ResponseInnerOneOf7Sink{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetLogStreams200ResponseInnerOneOf7Sink)

	if err != nil {
		return err
	}

	*o = GetLogStreams200ResponseInnerOneOf7Sink(varGetLogStreams200ResponseInnerOneOf7Sink)

	return err
}

type NullableGetLogStreams200ResponseInnerOneOf7Sink struct {
	value *GetLogStreams200ResponseInnerOneOf7Sink
	isSet bool
}

func (v NullableGetLogStreams200ResponseInnerOneOf7Sink) Get() *GetLogStreams200ResponseInnerOneOf7Sink {
	return v.value
}

func (v *NullableGetLogStreams200ResponseInnerOneOf7Sink) Set(val *GetLogStreams200ResponseInnerOneOf7Sink) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLogStreams200ResponseInnerOneOf7Sink) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLogStreams200ResponseInnerOneOf7Sink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLogStreams200ResponseInnerOneOf7Sink(val *GetLogStreams200ResponseInnerOneOf7Sink) *NullableGetLogStreams200ResponseInnerOneOf7Sink {
	return &NullableGetLogStreams200ResponseInnerOneOf7Sink{value: val, isSet: true}
}

func (v NullableGetLogStreams200ResponseInnerOneOf7Sink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLogStreams200ResponseInnerOneOf7Sink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
