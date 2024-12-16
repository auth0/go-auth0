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

// GetOrganizations200ResponseOneOfInnerBrandingColors Color scheme used to customize the login pages
type GetOrganizations200ResponseOneOfInnerBrandingColors struct {
	// HEX Color for primary elements
	Primary string `json:"primary"`
	// HEX Color for background
	PageBackground string `json:"page_background"`
}

type _GetOrganizations200ResponseOneOfInnerBrandingColors GetOrganizations200ResponseOneOfInnerBrandingColors

// GetPrimary returns the Primary field value
func (o *GetOrganizations200ResponseOneOfInnerBrandingColors) GetPrimary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value
// and a boolean to check if the value has been set.
func (o *GetOrganizations200ResponseOneOfInnerBrandingColors) GetPrimaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Primary, true
}

// SetPrimary sets field value
func (o *GetOrganizations200ResponseOneOfInnerBrandingColors) SetPrimary(v string) {
	o.Primary = v
}

// GetPageBackground returns the PageBackground field value
func (o *GetOrganizations200ResponseOneOfInnerBrandingColors) GetPageBackground() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PageBackground
}

// GetPageBackgroundOk returns a tuple with the PageBackground field value
// and a boolean to check if the value has been set.
func (o *GetOrganizations200ResponseOneOfInnerBrandingColors) GetPageBackgroundOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageBackground, true
}

// SetPageBackground sets field value
func (o *GetOrganizations200ResponseOneOfInnerBrandingColors) SetPageBackground(v string) {
	o.PageBackground = v
}

func (o GetOrganizations200ResponseOneOfInnerBrandingColors) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizations200ResponseOneOfInnerBrandingColors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["primary"] = o.Primary
	toSerialize["page_background"] = o.PageBackground
	return toSerialize, nil
}

func (o *GetOrganizations200ResponseOneOfInnerBrandingColors) UnmarshalJSON(data []byte) (err error) {
	varGetOrganizations200ResponseOneOfInnerBrandingColors := _GetOrganizations200ResponseOneOfInnerBrandingColors{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetOrganizations200ResponseOneOfInnerBrandingColors)

	if err != nil {
		return err
	}

	*o = GetOrganizations200ResponseOneOfInnerBrandingColors(varGetOrganizations200ResponseOneOfInnerBrandingColors)

	return err
}

type NullableGetOrganizations200ResponseOneOfInnerBrandingColors struct {
	value *GetOrganizations200ResponseOneOfInnerBrandingColors
	isSet bool
}

func (v NullableGetOrganizations200ResponseOneOfInnerBrandingColors) Get() *GetOrganizations200ResponseOneOfInnerBrandingColors {
	return v.value
}

func (v *NullableGetOrganizations200ResponseOneOfInnerBrandingColors) Set(val *GetOrganizations200ResponseOneOfInnerBrandingColors) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizations200ResponseOneOfInnerBrandingColors) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizations200ResponseOneOfInnerBrandingColors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizations200ResponseOneOfInnerBrandingColors(val *GetOrganizations200ResponseOneOfInnerBrandingColors) *NullableGetOrganizations200ResponseOneOfInnerBrandingColors {
	return &NullableGetOrganizations200ResponseOneOfInnerBrandingColors{value: val, isSet: true}
}

func (v NullableGetOrganizations200ResponseOneOfInnerBrandingColors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizations200ResponseOneOfInnerBrandingColors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
