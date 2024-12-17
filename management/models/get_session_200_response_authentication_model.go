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

// GetSession200ResponseAuthentication Details about authentication signals obtained during the login flow
type GetSession200ResponseAuthentication struct {
	// Contains the authentication methods a user has completed during their session
	Methods              []GetSession200ResponseAuthenticationMethodsInner `json:"methods,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSession200ResponseAuthentication GetSession200ResponseAuthentication

// GetMethods returns the Methods field value if set, zero value otherwise.
func (o *GetSession200ResponseAuthentication) GetMethods() []GetSession200ResponseAuthenticationMethodsInner {
	if o == nil || IsNil(o.Methods) {
		var ret []GetSession200ResponseAuthenticationMethodsInner
		return ret
	}
	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSession200ResponseAuthentication) GetMethodsOk() ([]GetSession200ResponseAuthenticationMethodsInner, bool) {
	if o == nil || IsNil(o.Methods) {
		return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *GetSession200ResponseAuthentication) HasMethods() bool {
	if o != nil && !IsNil(o.Methods) {
		return true
	}

	return false
}

// SetMethods gets a reference to the given []GetSession200ResponseAuthenticationMethodsInner and assigns it to the Methods field.
func (o *GetSession200ResponseAuthentication) SetMethods(v []GetSession200ResponseAuthenticationMethodsInner) {
	o.Methods = v
}

func (o GetSession200ResponseAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSession200ResponseAuthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Methods) {
		toSerialize["methods"] = o.Methods
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSession200ResponseAuthentication) UnmarshalJSON(data []byte) (err error) {
	varGetSession200ResponseAuthentication := _GetSession200ResponseAuthentication{}

	err = json.Unmarshal(data, &varGetSession200ResponseAuthentication)

	if err != nil {
		return err
	}

	*o = GetSession200ResponseAuthentication(varGetSession200ResponseAuthentication)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "methods")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSession200ResponseAuthentication struct {
	value *GetSession200ResponseAuthentication
	isSet bool
}

func (v NullableGetSession200ResponseAuthentication) Get() *GetSession200ResponseAuthentication {
	return v.value
}

func (v *NullableGetSession200ResponseAuthentication) Set(val *GetSession200ResponseAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSession200ResponseAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSession200ResponseAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSession200ResponseAuthentication(val *GetSession200ResponseAuthentication) *NullableGetSession200ResponseAuthentication {
	return &NullableGetSession200ResponseAuthentication{value: val, isSet: true}
}

func (v NullableGetSession200ResponseAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSession200ResponseAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
