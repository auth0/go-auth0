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

// GetBranding200ResponseFont Custom font settings.
type GetBranding200ResponseFont struct {
	// URL for the custom font. The URL must point to a font file and not a stylesheet. Must use HTTPS.
	Url string `json:"url"`
}

type _GetBranding200ResponseFont GetBranding200ResponseFont

// GetUrl returns the Url field value
func (o *GetBranding200ResponseFont) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *GetBranding200ResponseFont) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *GetBranding200ResponseFont) SetUrl(v string) {
	o.Url = v
}

func (o GetBranding200ResponseFont) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBranding200ResponseFont) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *GetBranding200ResponseFont) UnmarshalJSON(data []byte) (err error) {
	varGetBranding200ResponseFont := _GetBranding200ResponseFont{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetBranding200ResponseFont)

	if err != nil {
		return err
	}

	*o = GetBranding200ResponseFont(varGetBranding200ResponseFont)

	return err
}

type NullableGetBranding200ResponseFont struct {
	value *GetBranding200ResponseFont
	isSet bool
}

func (v NullableGetBranding200ResponseFont) Get() *GetBranding200ResponseFont {
	return v.value
}

func (v *NullableGetBranding200ResponseFont) Set(val *GetBranding200ResponseFont) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBranding200ResponseFont) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBranding200ResponseFont) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBranding200ResponseFont(val *GetBranding200ResponseFont) *NullableGetBranding200ResponseFont {
	return &NullableGetBranding200ResponseFont{value: val, isSet: true}
}

func (v NullableGetBranding200ResponseFont) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBranding200ResponseFont) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
