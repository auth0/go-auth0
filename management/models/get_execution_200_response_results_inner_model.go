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

// GetExecution200ResponseResultsInner Captures the results of a single action being executed.
type GetExecution200ResponseResultsInner struct {
	// The name of the action that was executed.
	ActionName *string                                               `json:"action_name,omitempty"`
	Error      *GetActionVersions200ResponseVersionsInnerErrorsInner `json:"error,omitempty"`
	// The time when the action was started.
	StartedAt *string `json:"started_at,omitempty"`
	// The time when the action finished executing.
	EndedAt *string `json:"ended_at,omitempty"`
}

// GetActionName returns the ActionName field value if set, zero value otherwise.
func (o *GetExecution200ResponseResultsInner) GetActionName() string {
	if o == nil || IsNil(o.ActionName) {
		var ret string
		return ret
	}
	return *o.ActionName
}

// GetActionNameOk returns a tuple with the ActionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExecution200ResponseResultsInner) GetActionNameOk() (*string, bool) {
	if o == nil || IsNil(o.ActionName) {
		return nil, false
	}
	return o.ActionName, true
}

// HasActionName returns a boolean if a field has been set.
func (o *GetExecution200ResponseResultsInner) HasActionName() bool {
	if o != nil && !IsNil(o.ActionName) {
		return true
	}

	return false
}

// SetActionName gets a reference to the given string and assigns it to the ActionName field.
func (o *GetExecution200ResponseResultsInner) SetActionName(v string) {
	o.ActionName = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetExecution200ResponseResultsInner) GetError() GetActionVersions200ResponseVersionsInnerErrorsInner {
	if o == nil || IsNil(o.Error) {
		var ret GetActionVersions200ResponseVersionsInnerErrorsInner
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExecution200ResponseResultsInner) GetErrorOk() (*GetActionVersions200ResponseVersionsInnerErrorsInner, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetExecution200ResponseResultsInner) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given GetActionVersions200ResponseVersionsInnerErrorsInner and assigns it to the Error field.
func (o *GetExecution200ResponseResultsInner) SetError(v GetActionVersions200ResponseVersionsInnerErrorsInner) {
	o.Error = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *GetExecution200ResponseResultsInner) GetStartedAt() string {
	if o == nil || IsNil(o.StartedAt) {
		var ret string
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExecution200ResponseResultsInner) GetStartedAtOk() (*string, bool) {
	if o == nil || IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *GetExecution200ResponseResultsInner) HasStartedAt() bool {
	if o != nil && !IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given string and assigns it to the StartedAt field.
func (o *GetExecution200ResponseResultsInner) SetStartedAt(v string) {
	o.StartedAt = &v
}

// GetEndedAt returns the EndedAt field value if set, zero value otherwise.
func (o *GetExecution200ResponseResultsInner) GetEndedAt() string {
	if o == nil || IsNil(o.EndedAt) {
		var ret string
		return ret
	}
	return *o.EndedAt
}

// GetEndedAtOk returns a tuple with the EndedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExecution200ResponseResultsInner) GetEndedAtOk() (*string, bool) {
	if o == nil || IsNil(o.EndedAt) {
		return nil, false
	}
	return o.EndedAt, true
}

// HasEndedAt returns a boolean if a field has been set.
func (o *GetExecution200ResponseResultsInner) HasEndedAt() bool {
	if o != nil && !IsNil(o.EndedAt) {
		return true
	}

	return false
}

// SetEndedAt gets a reference to the given string and assigns it to the EndedAt field.
func (o *GetExecution200ResponseResultsInner) SetEndedAt(v string) {
	o.EndedAt = &v
}

func (o GetExecution200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExecution200ResponseResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionName) {
		toSerialize["action_name"] = o.ActionName
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.StartedAt) {
		toSerialize["started_at"] = o.StartedAt
	}
	if !IsNil(o.EndedAt) {
		toSerialize["ended_at"] = o.EndedAt
	}
	return toSerialize, nil
}

type NullableGetExecution200ResponseResultsInner struct {
	value *GetExecution200ResponseResultsInner
	isSet bool
}

func (v NullableGetExecution200ResponseResultsInner) Get() *GetExecution200ResponseResultsInner {
	return v.value
}

func (v *NullableGetExecution200ResponseResultsInner) Set(val *GetExecution200ResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExecution200ResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExecution200ResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExecution200ResponseResultsInner(val *GetExecution200ResponseResultsInner) *NullableGetExecution200ResponseResultsInner {
	return &NullableGetExecution200ResponseResultsInner{value: val, isSet: true}
}

func (v NullableGetExecution200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExecution200ResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
