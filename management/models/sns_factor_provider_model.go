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

// SnsFactorProvider struct for SnsFactorProvider
type SnsFactorProvider struct {
	AwsAccessKeyId                NullableString `json:"aws_access_key_id"`
	AwsSecretAccessKey            NullableString `json:"aws_secret_access_key"`
	AwsRegion                     NullableString `json:"aws_region"`
	SnsApnsPlatformApplicationArn NullableString `json:"sns_apns_platform_application_arn"`
	SnsGcmPlatformApplicationArn  NullableString `json:"sns_gcm_platform_application_arn"`
}

type _SnsFactorProvider SnsFactorProvider

// GetAwsAccessKeyId returns the AwsAccessKeyId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SnsFactorProvider) GetAwsAccessKeyId() string {
	if o == nil || o.AwsAccessKeyId.Get() == nil {
		var ret string
		return ret
	}

	return *o.AwsAccessKeyId.Get()
}

// GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnsFactorProvider) GetAwsAccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsAccessKeyId.Get(), o.AwsAccessKeyId.IsSet()
}

// SetAwsAccessKeyId sets field value
func (o *SnsFactorProvider) SetAwsAccessKeyId(v string) {
	o.AwsAccessKeyId.Set(&v)
}

// GetAwsSecretAccessKey returns the AwsSecretAccessKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SnsFactorProvider) GetAwsSecretAccessKey() string {
	if o == nil || o.AwsSecretAccessKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.AwsSecretAccessKey.Get()
}

// GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnsFactorProvider) GetAwsSecretAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsSecretAccessKey.Get(), o.AwsSecretAccessKey.IsSet()
}

// SetAwsSecretAccessKey sets field value
func (o *SnsFactorProvider) SetAwsSecretAccessKey(v string) {
	o.AwsSecretAccessKey.Set(&v)
}

// GetAwsRegion returns the AwsRegion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SnsFactorProvider) GetAwsRegion() string {
	if o == nil || o.AwsRegion.Get() == nil {
		var ret string
		return ret
	}

	return *o.AwsRegion.Get()
}

// GetAwsRegionOk returns a tuple with the AwsRegion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnsFactorProvider) GetAwsRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsRegion.Get(), o.AwsRegion.IsSet()
}

// SetAwsRegion sets field value
func (o *SnsFactorProvider) SetAwsRegion(v string) {
	o.AwsRegion.Set(&v)
}

// GetSnsApnsPlatformApplicationArn returns the SnsApnsPlatformApplicationArn field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SnsFactorProvider) GetSnsApnsPlatformApplicationArn() string {
	if o == nil || o.SnsApnsPlatformApplicationArn.Get() == nil {
		var ret string
		return ret
	}

	return *o.SnsApnsPlatformApplicationArn.Get()
}

// GetSnsApnsPlatformApplicationArnOk returns a tuple with the SnsApnsPlatformApplicationArn field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnsFactorProvider) GetSnsApnsPlatformApplicationArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnsApnsPlatformApplicationArn.Get(), o.SnsApnsPlatformApplicationArn.IsSet()
}

// SetSnsApnsPlatformApplicationArn sets field value
func (o *SnsFactorProvider) SetSnsApnsPlatformApplicationArn(v string) {
	o.SnsApnsPlatformApplicationArn.Set(&v)
}

// GetSnsGcmPlatformApplicationArn returns the SnsGcmPlatformApplicationArn field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SnsFactorProvider) GetSnsGcmPlatformApplicationArn() string {
	if o == nil || o.SnsGcmPlatformApplicationArn.Get() == nil {
		var ret string
		return ret
	}

	return *o.SnsGcmPlatformApplicationArn.Get()
}

// GetSnsGcmPlatformApplicationArnOk returns a tuple with the SnsGcmPlatformApplicationArn field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnsFactorProvider) GetSnsGcmPlatformApplicationArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnsGcmPlatformApplicationArn.Get(), o.SnsGcmPlatformApplicationArn.IsSet()
}

// SetSnsGcmPlatformApplicationArn sets field value
func (o *SnsFactorProvider) SetSnsGcmPlatformApplicationArn(v string) {
	o.SnsGcmPlatformApplicationArn.Set(&v)
}

func (o SnsFactorProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnsFactorProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aws_access_key_id"] = o.AwsAccessKeyId.Get()
	toSerialize["aws_secret_access_key"] = o.AwsSecretAccessKey.Get()
	toSerialize["aws_region"] = o.AwsRegion.Get()
	toSerialize["sns_apns_platform_application_arn"] = o.SnsApnsPlatformApplicationArn.Get()
	toSerialize["sns_gcm_platform_application_arn"] = o.SnsGcmPlatformApplicationArn.Get()
	return toSerialize, nil
}

func (o *SnsFactorProvider) UnmarshalJSON(data []byte) (err error) {
	varSnsFactorProvider := _SnsFactorProvider{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSnsFactorProvider)

	if err != nil {
		return err
	}

	*o = SnsFactorProvider(varSnsFactorProvider)

	return err
}

type NullableSnsFactorProvider struct {
	value *SnsFactorProvider
	isSet bool
}

func (v NullableSnsFactorProvider) Get() *SnsFactorProvider {
	return v.value
}

func (v *NullableSnsFactorProvider) Set(val *SnsFactorProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableSnsFactorProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableSnsFactorProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnsFactorProvider(val *SnsFactorProvider) *NullableSnsFactorProvider {
	return &NullableSnsFactorProvider{value: val, isSet: true}
}

func (v NullableSnsFactorProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnsFactorProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
