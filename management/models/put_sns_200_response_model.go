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

// PutSns200Response struct for PutSns200Response
type PutSns200Response struct {
	AwsAccessKeyId                NullableString `json:"aws_access_key_id"`
	AwsSecretAccessKey            NullableString `json:"aws_secret_access_key"`
	AwsRegion                     NullableString `json:"aws_region"`
	SnsApnsPlatformApplicationArn NullableString `json:"sns_apns_platform_application_arn"`
	SnsGcmPlatformApplicationArn  NullableString `json:"sns_gcm_platform_application_arn"`
}

type _PutSns200Response PutSns200Response

// GetAwsAccessKeyId returns the AwsAccessKeyId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PutSns200Response) GetAwsAccessKeyId() string {
	if o == nil || o.AwsAccessKeyId.Get() == nil {
		var ret string
		return ret
	}

	return *o.AwsAccessKeyId.Get()
}

// GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutSns200Response) GetAwsAccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsAccessKeyId.Get(), o.AwsAccessKeyId.IsSet()
}

// SetAwsAccessKeyId sets field value
func (o *PutSns200Response) SetAwsAccessKeyId(v string) {
	o.AwsAccessKeyId.Set(&v)
}

// GetAwsSecretAccessKey returns the AwsSecretAccessKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PutSns200Response) GetAwsSecretAccessKey() string {
	if o == nil || o.AwsSecretAccessKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.AwsSecretAccessKey.Get()
}

// GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutSns200Response) GetAwsSecretAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsSecretAccessKey.Get(), o.AwsSecretAccessKey.IsSet()
}

// SetAwsSecretAccessKey sets field value
func (o *PutSns200Response) SetAwsSecretAccessKey(v string) {
	o.AwsSecretAccessKey.Set(&v)
}

// GetAwsRegion returns the AwsRegion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PutSns200Response) GetAwsRegion() string {
	if o == nil || o.AwsRegion.Get() == nil {
		var ret string
		return ret
	}

	return *o.AwsRegion.Get()
}

// GetAwsRegionOk returns a tuple with the AwsRegion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutSns200Response) GetAwsRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsRegion.Get(), o.AwsRegion.IsSet()
}

// SetAwsRegion sets field value
func (o *PutSns200Response) SetAwsRegion(v string) {
	o.AwsRegion.Set(&v)
}

// GetSnsApnsPlatformApplicationArn returns the SnsApnsPlatformApplicationArn field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PutSns200Response) GetSnsApnsPlatformApplicationArn() string {
	if o == nil || o.SnsApnsPlatformApplicationArn.Get() == nil {
		var ret string
		return ret
	}

	return *o.SnsApnsPlatformApplicationArn.Get()
}

// GetSnsApnsPlatformApplicationArnOk returns a tuple with the SnsApnsPlatformApplicationArn field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutSns200Response) GetSnsApnsPlatformApplicationArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnsApnsPlatformApplicationArn.Get(), o.SnsApnsPlatformApplicationArn.IsSet()
}

// SetSnsApnsPlatformApplicationArn sets field value
func (o *PutSns200Response) SetSnsApnsPlatformApplicationArn(v string) {
	o.SnsApnsPlatformApplicationArn.Set(&v)
}

// GetSnsGcmPlatformApplicationArn returns the SnsGcmPlatformApplicationArn field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PutSns200Response) GetSnsGcmPlatformApplicationArn() string {
	if o == nil || o.SnsGcmPlatformApplicationArn.Get() == nil {
		var ret string
		return ret
	}

	return *o.SnsGcmPlatformApplicationArn.Get()
}

// GetSnsGcmPlatformApplicationArnOk returns a tuple with the SnsGcmPlatformApplicationArn field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutSns200Response) GetSnsGcmPlatformApplicationArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnsGcmPlatformApplicationArn.Get(), o.SnsGcmPlatformApplicationArn.IsSet()
}

// SetSnsGcmPlatformApplicationArn sets field value
func (o *PutSns200Response) SetSnsGcmPlatformApplicationArn(v string) {
	o.SnsGcmPlatformApplicationArn.Set(&v)
}

func (o PutSns200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutSns200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aws_access_key_id"] = o.AwsAccessKeyId.Get()
	toSerialize["aws_secret_access_key"] = o.AwsSecretAccessKey.Get()
	toSerialize["aws_region"] = o.AwsRegion.Get()
	toSerialize["sns_apns_platform_application_arn"] = o.SnsApnsPlatformApplicationArn.Get()
	toSerialize["sns_gcm_platform_application_arn"] = o.SnsGcmPlatformApplicationArn.Get()
	return toSerialize, nil
}

func (o *PutSns200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"aws_access_key_id",
		"aws_secret_access_key",
		"aws_region",
		"sns_apns_platform_application_arn",
		"sns_gcm_platform_application_arn",
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

	varPutSns200Response := _PutSns200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPutSns200Response)

	if err != nil {
		return err
	}

	*o = PutSns200Response(varPutSns200Response)

	return err
}

type NullablePutSns200Response struct {
	value *PutSns200Response
	isSet bool
}

func (v NullablePutSns200Response) Get() *PutSns200Response {
	return v.value
}

func (v *NullablePutSns200Response) Set(val *PutSns200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePutSns200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePutSns200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutSns200Response(val *PutSns200Response) *NullablePutSns200Response {
	return &NullablePutSns200Response{value: val, isSet: true}
}

func (v NullablePutSns200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutSns200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
