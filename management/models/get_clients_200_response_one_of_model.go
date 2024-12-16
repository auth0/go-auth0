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

// GetClients200ResponseOneOf struct for GetClients200ResponseOneOf
type GetClients200ResponseOneOf struct {
	Start   float32  `json:"start"`
	Limit   float32  `json:"limit"`
	Total   float32  `json:"total"`
	Clients []Client `json:"clients"`
}

type _GetClients200ResponseOneOf GetClients200ResponseOneOf

// GetStart returns the Start field value
func (o *GetClients200ResponseOneOf) GetStart() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *GetClients200ResponseOneOf) GetStartOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *GetClients200ResponseOneOf) SetStart(v float32) {
	o.Start = v
}

// GetLimit returns the Limit field value
func (o *GetClients200ResponseOneOf) GetLimit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *GetClients200ResponseOneOf) GetLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *GetClients200ResponseOneOf) SetLimit(v float32) {
	o.Limit = v
}

// GetTotal returns the Total field value
func (o *GetClients200ResponseOneOf) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *GetClients200ResponseOneOf) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *GetClients200ResponseOneOf) SetTotal(v float32) {
	o.Total = v
}

// GetClients returns the Clients field value
func (o *GetClients200ResponseOneOf) GetClients() []Client {
	if o == nil {
		var ret []Client
		return ret
	}

	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value
// and a boolean to check if the value has been set.
func (o *GetClients200ResponseOneOf) GetClientsOk() ([]Client, bool) {
	if o == nil {
		return nil, false
	}
	return o.Clients, true
}

// SetClients sets field value
func (o *GetClients200ResponseOneOf) SetClients(v []Client) {
	o.Clients = v
}

func (o GetClients200ResponseOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetClients200ResponseOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start"] = o.Start
	toSerialize["limit"] = o.Limit
	toSerialize["total"] = o.Total
	toSerialize["clients"] = o.Clients
	return toSerialize, nil
}

func (o *GetClients200ResponseOneOf) UnmarshalJSON(data []byte) (err error) {
	varGetClients200ResponseOneOf := _GetClients200ResponseOneOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetClients200ResponseOneOf)

	if err != nil {
		return err
	}

	*o = GetClients200ResponseOneOf(varGetClients200ResponseOneOf)

	return err
}

type NullableGetClients200ResponseOneOf struct {
	value *GetClients200ResponseOneOf
	isSet bool
}

func (v NullableGetClients200ResponseOneOf) Get() *GetClients200ResponseOneOf {
	return v.value
}

func (v *NullableGetClients200ResponseOneOf) Set(val *GetClients200ResponseOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetClients200ResponseOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetClients200ResponseOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetClients200ResponseOneOf(val *GetClients200ResponseOneOf) *NullableGetClients200ResponseOneOf {
	return &NullableGetClients200ResponseOneOf{value: val, isSet: true}
}

func (v NullableGetClients200ResponseOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetClients200ResponseOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
