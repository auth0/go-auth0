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

// GetMembers200ResponseOneOf struct for GetMembers200ResponseOneOf
type GetMembers200ResponseOneOf struct {
	Start   *float32                          `json:"start,omitempty"`
	Limit   *float32                          `json:"limit,omitempty"`
	Total   *float32                          `json:"total,omitempty"`
	Members []GetMembers200ResponseOneOfInner `json:"members,omitempty"`
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *GetMembers200ResponseOneOf) GetStart() float32 {
	if o == nil || IsNil(o.Start) {
		var ret float32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMembers200ResponseOneOf) GetStartOk() (*float32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *GetMembers200ResponseOneOf) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given float32 and assigns it to the Start field.
func (o *GetMembers200ResponseOneOf) SetStart(v float32) {
	o.Start = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetMembers200ResponseOneOf) GetLimit() float32 {
	if o == nil || IsNil(o.Limit) {
		var ret float32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMembers200ResponseOneOf) GetLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetMembers200ResponseOneOf) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given float32 and assigns it to the Limit field.
func (o *GetMembers200ResponseOneOf) SetLimit(v float32) {
	o.Limit = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetMembers200ResponseOneOf) GetTotal() float32 {
	if o == nil || IsNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMembers200ResponseOneOf) GetTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetMembers200ResponseOneOf) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *GetMembers200ResponseOneOf) SetTotal(v float32) {
	o.Total = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *GetMembers200ResponseOneOf) GetMembers() []GetMembers200ResponseOneOfInner {
	if o == nil || IsNil(o.Members) {
		var ret []GetMembers200ResponseOneOfInner
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMembers200ResponseOneOf) GetMembersOk() ([]GetMembers200ResponseOneOfInner, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *GetMembers200ResponseOneOf) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []GetMembers200ResponseOneOfInner and assigns it to the Members field.
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
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	return toSerialize, nil
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
