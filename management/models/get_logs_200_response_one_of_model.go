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

// GetLogs200ResponseOneOf struct for GetLogs200ResponseOneOf
type GetLogs200ResponseOneOf struct {
	Start  *float32 `json:"start,omitempty"`
	Limit  *float32 `json:"limit,omitempty"`
	Length *float32 `json:"length,omitempty"`
	Total  *float32 `json:"total,omitempty"`
	Logs   []Log    `json:"logs,omitempty"`
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *GetLogs200ResponseOneOf) GetStart() float32 {
	if o == nil || IsNil(o.Start) {
		var ret float32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogs200ResponseOneOf) GetStartOk() (*float32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *GetLogs200ResponseOneOf) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given float32 and assigns it to the Start field.
func (o *GetLogs200ResponseOneOf) SetStart(v float32) {
	o.Start = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetLogs200ResponseOneOf) GetLimit() float32 {
	if o == nil || IsNil(o.Limit) {
		var ret float32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogs200ResponseOneOf) GetLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetLogs200ResponseOneOf) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given float32 and assigns it to the Limit field.
func (o *GetLogs200ResponseOneOf) SetLimit(v float32) {
	o.Limit = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *GetLogs200ResponseOneOf) GetLength() float32 {
	if o == nil || IsNil(o.Length) {
		var ret float32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogs200ResponseOneOf) GetLengthOk() (*float32, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *GetLogs200ResponseOneOf) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given float32 and assigns it to the Length field.
func (o *GetLogs200ResponseOneOf) SetLength(v float32) {
	o.Length = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetLogs200ResponseOneOf) GetTotal() float32 {
	if o == nil || IsNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogs200ResponseOneOf) GetTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetLogs200ResponseOneOf) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *GetLogs200ResponseOneOf) SetTotal(v float32) {
	o.Total = &v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *GetLogs200ResponseOneOf) GetLogs() []Log {
	if o == nil || IsNil(o.Logs) {
		var ret []Log
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogs200ResponseOneOf) GetLogsOk() ([]Log, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *GetLogs200ResponseOneOf) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []Log and assigns it to the Logs field.
func (o *GetLogs200ResponseOneOf) SetLogs(v []Log) {
	o.Logs = v
}

func (o GetLogs200ResponseOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLogs200ResponseOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	return toSerialize, nil
}

type NullableGetLogs200ResponseOneOf struct {
	value *GetLogs200ResponseOneOf
	isSet bool
}

func (v NullableGetLogs200ResponseOneOf) Get() *GetLogs200ResponseOneOf {
	return v.value
}

func (v *NullableGetLogs200ResponseOneOf) Set(val *GetLogs200ResponseOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLogs200ResponseOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLogs200ResponseOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLogs200ResponseOneOf(val *GetLogs200ResponseOneOf) *NullableGetLogs200ResponseOneOf {
	return &NullableGetLogs200ResponseOneOf{value: val, isSet: true}
}

func (v NullableGetLogs200ResponseOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLogs200ResponseOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
