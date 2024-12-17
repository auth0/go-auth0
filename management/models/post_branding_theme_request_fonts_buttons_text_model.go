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

// PostBrandingThemeRequestFontsButtonsText Buttons text
type PostBrandingThemeRequestFontsButtonsText struct {
	// Buttons text bold
	Bold bool `json:"bold"`
	// Buttons text size
	Size float32 `json:"size"`
}

type _PostBrandingThemeRequestFontsButtonsText PostBrandingThemeRequestFontsButtonsText

// GetBold returns the Bold field value
func (o *PostBrandingThemeRequestFontsButtonsText) GetBold() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Bold
}

// GetBoldOk returns a tuple with the Bold field value
// and a boolean to check if the value has been set.
func (o *PostBrandingThemeRequestFontsButtonsText) GetBoldOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bold, true
}

// SetBold sets field value
func (o *PostBrandingThemeRequestFontsButtonsText) SetBold(v bool) {
	o.Bold = v
}

// GetSize returns the Size field value
func (o *PostBrandingThemeRequestFontsButtonsText) GetSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *PostBrandingThemeRequestFontsButtonsText) GetSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *PostBrandingThemeRequestFontsButtonsText) SetSize(v float32) {
	o.Size = v
}

func (o PostBrandingThemeRequestFontsButtonsText) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostBrandingThemeRequestFontsButtonsText) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bold"] = o.Bold
	toSerialize["size"] = o.Size
	return toSerialize, nil
}

func (o *PostBrandingThemeRequestFontsButtonsText) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bold",
		"size",
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

	varPostBrandingThemeRequestFontsButtonsText := _PostBrandingThemeRequestFontsButtonsText{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varPostBrandingThemeRequestFontsButtonsText)

	if err != nil {
		return err
	}

	*o = PostBrandingThemeRequestFontsButtonsText(varPostBrandingThemeRequestFontsButtonsText)

	return err
}

type NullablePostBrandingThemeRequestFontsButtonsText struct {
	value *PostBrandingThemeRequestFontsButtonsText
	isSet bool
}

func (v NullablePostBrandingThemeRequestFontsButtonsText) Get() *PostBrandingThemeRequestFontsButtonsText {
	return v.value
}

func (v *NullablePostBrandingThemeRequestFontsButtonsText) Set(val *PostBrandingThemeRequestFontsButtonsText) {
	v.value = val
	v.isSet = true
}

func (v NullablePostBrandingThemeRequestFontsButtonsText) IsSet() bool {
	return v.isSet
}

func (v *NullablePostBrandingThemeRequestFontsButtonsText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostBrandingThemeRequestFontsButtonsText(val *PostBrandingThemeRequestFontsButtonsText) *NullablePostBrandingThemeRequestFontsButtonsText {
	return &NullablePostBrandingThemeRequestFontsButtonsText{value: val, isSet: true}
}

func (v NullablePostBrandingThemeRequestFontsButtonsText) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostBrandingThemeRequestFontsButtonsText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
