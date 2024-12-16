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

// GetBindings200Response struct for GetBindings200Response
type GetBindings200Response struct {
	// The total result count.
	Total float32 `json:"total"`
	// Page index of the results being returned. First page is 0.
	Page float32 `json:"page"`
	// Number of results per page.
	PerPage float32 `json:"per_page"`
	// The list of actions that are bound to this trigger in the order in which they will be executed.
	Bindings []GetBindings200ResponseBindingsInner `json:"bindings"`
}

type _GetBindings200Response GetBindings200Response

// GetTotal returns the Total field value
func (o *GetBindings200Response) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *GetBindings200Response) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *GetBindings200Response) SetTotal(v float32) {
	o.Total = v
}

// GetPage returns the Page field value
func (o *GetBindings200Response) GetPage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *GetBindings200Response) GetPageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *GetBindings200Response) SetPage(v float32) {
	o.Page = v
}

// GetPerPage returns the PerPage field value
func (o *GetBindings200Response) GetPerPage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value
// and a boolean to check if the value has been set.
func (o *GetBindings200Response) GetPerPageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PerPage, true
}

// SetPerPage sets field value
func (o *GetBindings200Response) SetPerPage(v float32) {
	o.PerPage = v
}

// GetBindings returns the Bindings field value
func (o *GetBindings200Response) GetBindings() []GetBindings200ResponseBindingsInner {
	if o == nil {
		var ret []GetBindings200ResponseBindingsInner
		return ret
	}

	return o.Bindings
}

// GetBindingsOk returns a tuple with the Bindings field value
// and a boolean to check if the value has been set.
func (o *GetBindings200Response) GetBindingsOk() ([]GetBindings200ResponseBindingsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bindings, true
}

// SetBindings sets field value
func (o *GetBindings200Response) SetBindings(v []GetBindings200ResponseBindingsInner) {
	o.Bindings = v
}

func (o GetBindings200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBindings200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total"] = o.Total
	toSerialize["page"] = o.Page
	toSerialize["per_page"] = o.PerPage
	toSerialize["bindings"] = o.Bindings
	return toSerialize, nil
}

func (o *GetBindings200Response) UnmarshalJSON(data []byte) (err error) {
	varGetBindings200Response := _GetBindings200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetBindings200Response)

	if err != nil {
		return err
	}

	*o = GetBindings200Response(varGetBindings200Response)

	return err
}

type NullableGetBindings200Response struct {
	value *GetBindings200Response
	isSet bool
}

func (v NullableGetBindings200Response) Get() *GetBindings200Response {
	return v.value
}

func (v *NullableGetBindings200Response) Set(val *GetBindings200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBindings200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBindings200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBindings200Response(val *GetBindings200Response) *NullableGetBindings200Response {
	return &NullableGetBindings200Response{value: val, isSet: true}
}

func (v NullableGetBindings200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBindings200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
