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

// GetMembers200ResponseOneOf struct for GetMembers200ResponseOneOf
type GetMembers200ResponseOneOf struct {
	Start   float32                           `json:"start"`
	Limit   float32                           `json:"limit"`
	Total   float32                           `json:"total"`
	Members []GetMembers200ResponseOneOfInner `json:"members"`
}

type _GetMembers200ResponseOneOf GetMembers200ResponseOneOf

// GetStart returns the Start field value
func (o *GetMembers200ResponseOneOf) GetStart() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *GetMembers200ResponseOneOf) GetStartOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *GetMembers200ResponseOneOf) SetStart(v float32) {
	o.Start = v
}

// GetLimit returns the Limit field value
func (o *GetMembers200ResponseOneOf) GetLimit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *GetMembers200ResponseOneOf) GetLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *GetMembers200ResponseOneOf) SetLimit(v float32) {
	o.Limit = v
}

// GetTotal returns the Total field value
func (o *GetMembers200ResponseOneOf) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *GetMembers200ResponseOneOf) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *GetMembers200ResponseOneOf) SetTotal(v float32) {
	o.Total = v
}

// GetMembers returns the Members field value
func (o *GetMembers200ResponseOneOf) GetMembers() []GetMembers200ResponseOneOfInner {
	if o == nil {
		var ret []GetMembers200ResponseOneOfInner
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *GetMembers200ResponseOneOf) GetMembersOk() ([]GetMembers200ResponseOneOfInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *GetMembers200ResponseOneOf) SetMembers(v []GetMembers200ResponseOneOfInner) {
	o.Members = v
}

func (o GetMembers200ResponseOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMembers200ResponseOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start"] = o.Start
	toSerialize["limit"] = o.Limit
	toSerialize["total"] = o.Total
	toSerialize["members"] = o.Members
	return toSerialize, nil
}

func (o *GetMembers200ResponseOneOf) UnmarshalJSON(data []byte) (err error) {
	varGetMembers200ResponseOneOf := _GetMembers200ResponseOneOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetMembers200ResponseOneOf)

	if err != nil {
		return err
	}

	*o = GetMembers200ResponseOneOf(varGetMembers200ResponseOneOf)

	return err
}

type NullableGetMembers200ResponseOneOf struct {
	value *GetMembers200ResponseOneOf
	isSet bool
}

func (v NullableGetMembers200ResponseOneOf) Get() *GetMembers200ResponseOneOf {
	return v.value
}

func (v *NullableGetMembers200ResponseOneOf) Set(val *GetMembers200ResponseOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMembers200ResponseOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMembers200ResponseOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMembers200ResponseOneOf(val *GetMembers200ResponseOneOf) *NullableGetMembers200ResponseOneOf {
	return &NullableGetMembers200ResponseOneOf{value: val, isSet: true}
}

func (v NullableGetMembers200ResponseOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMembers200ResponseOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
