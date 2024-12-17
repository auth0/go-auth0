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

// PatchBindings200Response struct for PatchBindings200Response
type PatchBindings200Response struct {
	Bindings []GetBindings200ResponseBindingsInner `json:"bindings,omitempty"`
}

// GetBindings returns the Bindings field value if set, zero value otherwise.
func (o *PatchBindings200Response) GetBindings() []GetBindings200ResponseBindingsInner {
	if o == nil || IsNil(o.Bindings) {
		var ret []GetBindings200ResponseBindingsInner
		return ret
	}
	return o.Bindings
}

// GetBindingsOk returns a tuple with the Bindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBindings200Response) GetBindingsOk() ([]GetBindings200ResponseBindingsInner, bool) {
	if o == nil || IsNil(o.Bindings) {
		return nil, false
	}
	return o.Bindings, true
}

// HasBindings returns a boolean if a field has been set.
func (o *PatchBindings200Response) HasBindings() bool {
	if o != nil && !IsNil(o.Bindings) {
		return true
	}

	return false
}

// SetBindings gets a reference to the given []GetBindings200ResponseBindingsInner and assigns it to the Bindings field.
func (o *PatchBindings200Response) SetBindings(v []GetBindings200ResponseBindingsInner) {
	o.Bindings = v
}

func (o PatchBindings200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchBindings200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bindings) {
		toSerialize["bindings"] = o.Bindings
	}
	return toSerialize, nil
}

type NullablePatchBindings200Response struct {
	value *PatchBindings200Response
	isSet bool
}

func (v NullablePatchBindings200Response) Get() *PatchBindings200Response {
	return v.value
}

func (v *NullablePatchBindings200Response) Set(val *PatchBindings200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchBindings200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchBindings200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchBindings200Response(val *PatchBindings200Response) *NullablePatchBindings200Response {
	return &NullablePatchBindings200Response{value: val, isSet: true}
}

func (v NullablePatchBindings200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchBindings200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
