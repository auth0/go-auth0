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

// GetLogStreams200ResponseInnerOneOf1Sink struct for GetLogStreams200ResponseInnerOneOf1Sink
type GetLogStreams200ResponseInnerOneOf1Sink struct {
	// AWS account ID
	AwsAccountId string                                           `json:"awsAccountId"`
	AwsRegion    GetLogStreams200ResponseInnerOneOf1SinkAwsRegion `json:"awsRegion"`
	// AWS EventBridge partner event source
	AwsPartnerEventSource *string `json:"awsPartnerEventSource,omitempty"`
}

type _GetLogStreams200ResponseInnerOneOf1Sink GetLogStreams200ResponseInnerOneOf1Sink

// GetAwsAccountId returns the AwsAccountId field value
func (o *GetLogStreams200ResponseInnerOneOf1Sink) GetAwsAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsAccountId
}

// GetAwsAccountIdOk returns a tuple with the AwsAccountId field value
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf1Sink) GetAwsAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsAccountId, true
}

// SetAwsAccountId sets field value
func (o *GetLogStreams200ResponseInnerOneOf1Sink) SetAwsAccountId(v string) {
	o.AwsAccountId = v
}

// GetAwsRegion returns the AwsRegion field value
func (o *GetLogStreams200ResponseInnerOneOf1Sink) GetAwsRegion() GetLogStreams200ResponseInnerOneOf1SinkAwsRegion {
	if o == nil {
		var ret GetLogStreams200ResponseInnerOneOf1SinkAwsRegion
		return ret
	}

	return o.AwsRegion
}

// GetAwsRegionOk returns a tuple with the AwsRegion field value
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf1Sink) GetAwsRegionOk() (*GetLogStreams200ResponseInnerOneOf1SinkAwsRegion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsRegion, true
}

// SetAwsRegion sets field value
func (o *GetLogStreams200ResponseInnerOneOf1Sink) SetAwsRegion(v GetLogStreams200ResponseInnerOneOf1SinkAwsRegion) {
	o.AwsRegion = v
}

// GetAwsPartnerEventSource returns the AwsPartnerEventSource field value if set, zero value otherwise.
func (o *GetLogStreams200ResponseInnerOneOf1Sink) GetAwsPartnerEventSource() string {
	if o == nil || IsNil(o.AwsPartnerEventSource) {
		var ret string
		return ret
	}
	return *o.AwsPartnerEventSource
}

// GetAwsPartnerEventSourceOk returns a tuple with the AwsPartnerEventSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogStreams200ResponseInnerOneOf1Sink) GetAwsPartnerEventSourceOk() (*string, bool) {
	if o == nil || IsNil(o.AwsPartnerEventSource) {
		return nil, false
	}
	return o.AwsPartnerEventSource, true
}

// HasAwsPartnerEventSource returns a boolean if a field has been set.
func (o *GetLogStreams200ResponseInnerOneOf1Sink) HasAwsPartnerEventSource() bool {
	if o != nil && !IsNil(o.AwsPartnerEventSource) {
		return true
	}

	return false
}

// SetAwsPartnerEventSource gets a reference to the given string and assigns it to the AwsPartnerEventSource field.
func (o *GetLogStreams200ResponseInnerOneOf1Sink) SetAwsPartnerEventSource(v string) {
	o.AwsPartnerEventSource = &v
}

func (o GetLogStreams200ResponseInnerOneOf1Sink) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLogStreams200ResponseInnerOneOf1Sink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["awsAccountId"] = o.AwsAccountId
	toSerialize["awsRegion"] = o.AwsRegion
	if !IsNil(o.AwsPartnerEventSource) {
		toSerialize["awsPartnerEventSource"] = o.AwsPartnerEventSource
	}
	return toSerialize, nil
}

func (o *GetLogStreams200ResponseInnerOneOf1Sink) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"awsAccountId",
		"awsRegion",
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

	varGetLogStreams200ResponseInnerOneOf1Sink := _GetLogStreams200ResponseInnerOneOf1Sink{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetLogStreams200ResponseInnerOneOf1Sink)

	if err != nil {
		return err
	}

	*o = GetLogStreams200ResponseInnerOneOf1Sink(varGetLogStreams200ResponseInnerOneOf1Sink)

	return err
}

type NullableGetLogStreams200ResponseInnerOneOf1Sink struct {
	value *GetLogStreams200ResponseInnerOneOf1Sink
	isSet bool
}

func (v NullableGetLogStreams200ResponseInnerOneOf1Sink) Get() *GetLogStreams200ResponseInnerOneOf1Sink {
	return v.value
}

func (v *NullableGetLogStreams200ResponseInnerOneOf1Sink) Set(val *GetLogStreams200ResponseInnerOneOf1Sink) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLogStreams200ResponseInnerOneOf1Sink) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLogStreams200ResponseInnerOneOf1Sink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLogStreams200ResponseInnerOneOf1Sink(val *GetLogStreams200ResponseInnerOneOf1Sink) *NullableGetLogStreams200ResponseInnerOneOf1Sink {
	return &NullableGetLogStreams200ResponseInnerOneOf1Sink{value: val, isSet: true}
}

func (v NullableGetLogStreams200ResponseInnerOneOf1Sink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLogStreams200ResponseInnerOneOf1Sink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
