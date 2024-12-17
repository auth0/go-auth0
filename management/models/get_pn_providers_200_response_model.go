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

// GetPnProviders200Response struct for GetPnProviders200Response
type GetPnProviders200Response struct {
	Provider *GetPnProviders200ResponseProvider `json:"provider,omitempty"`
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *GetPnProviders200Response) GetProvider() GetPnProviders200ResponseProvider {
	if o == nil || IsNil(o.Provider) {
		var ret GetPnProviders200ResponseProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPnProviders200Response) GetProviderOk() (*GetPnProviders200ResponseProvider, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *GetPnProviders200Response) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given GetPnProviders200ResponseProvider and assigns it to the Provider field.
func (o *GetPnProviders200Response) SetProvider(v GetPnProviders200ResponseProvider) {
	o.Provider = &v
}

func (o GetPnProviders200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPnProviders200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	return toSerialize, nil
}

type NullableGetPnProviders200Response struct {
	value *GetPnProviders200Response
	isSet bool
}

func (v NullableGetPnProviders200Response) Get() *GetPnProviders200Response {
	return v.value
}

func (v *NullableGetPnProviders200Response) Set(val *GetPnProviders200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPnProviders200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPnProviders200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPnProviders200Response(val *GetPnProviders200Response) *NullableGetPnProviders200Response {
	return &NullableGetPnProviders200Response{value: val, isSet: true}
}

func (v NullableGetPnProviders200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPnProviders200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
