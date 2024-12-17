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

// GetDefaultMapping200Response struct for GetDefaultMapping200Response
type GetDefaultMapping200Response struct {
	// The mapping between auth0 and SCIM
	Mapping []GetScimConfiguration200ResponseMappingInner `json:"mapping,omitempty"`
}

// GetMapping returns the Mapping field value if set, zero value otherwise.
func (o *GetDefaultMapping200Response) GetMapping() []GetScimConfiguration200ResponseMappingInner {
	if o == nil || IsNil(o.Mapping) {
		var ret []GetScimConfiguration200ResponseMappingInner
		return ret
	}
	return o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDefaultMapping200Response) GetMappingOk() ([]GetScimConfiguration200ResponseMappingInner, bool) {
	if o == nil || IsNil(o.Mapping) {
		return nil, false
	}
	return o.Mapping, true
}

// HasMapping returns a boolean if a field has been set.
func (o *GetDefaultMapping200Response) HasMapping() bool {
	if o != nil && !IsNil(o.Mapping) {
		return true
	}

	return false
}

// SetMapping gets a reference to the given []GetScimConfiguration200ResponseMappingInner and assigns it to the Mapping field.
func (o *GetDefaultMapping200Response) SetMapping(v []GetScimConfiguration200ResponseMappingInner) {
	o.Mapping = v
}

func (o GetDefaultMapping200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDefaultMapping200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mapping) {
		toSerialize["mapping"] = o.Mapping
	}
	return toSerialize, nil
}

type NullableGetDefaultMapping200Response struct {
	value *GetDefaultMapping200Response
	isSet bool
}

func (v NullableGetDefaultMapping200Response) Get() *GetDefaultMapping200Response {
	return v.value
}

func (v *NullableGetDefaultMapping200Response) Set(val *GetDefaultMapping200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDefaultMapping200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDefaultMapping200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDefaultMapping200Response(val *GetDefaultMapping200Response) *NullableGetDefaultMapping200Response {
	return &NullableGetDefaultMapping200Response{value: val, isSet: true}
}

func (v NullableGetDefaultMapping200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDefaultMapping200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
