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

// ActionsDraftUpdate struct for ActionsDraftUpdate
type ActionsDraftUpdate struct {
	// True if the draft of the action should be updated with the reverted version.
	UpdateDraft *bool `json:"update_draft,omitempty"`
}

// GetUpdateDraft returns the UpdateDraft field value if set, zero value otherwise.
func (o *ActionsDraftUpdate) GetUpdateDraft() bool {
	if o == nil || IsNil(o.UpdateDraft) {
		var ret bool
		return ret
	}
	return *o.UpdateDraft
}

// GetUpdateDraftOk returns a tuple with the UpdateDraft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsDraftUpdate) GetUpdateDraftOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateDraft) {
		return nil, false
	}
	return o.UpdateDraft, true
}

// HasUpdateDraft returns a boolean if a field has been set.
func (o *ActionsDraftUpdate) HasUpdateDraft() bool {
	if o != nil && !IsNil(o.UpdateDraft) {
		return true
	}

	return false
}

// SetUpdateDraft gets a reference to the given bool and assigns it to the UpdateDraft field.
func (o *ActionsDraftUpdate) SetUpdateDraft(v bool) {
	o.UpdateDraft = &v
}

func (o ActionsDraftUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionsDraftUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UpdateDraft) {
		toSerialize["update_draft"] = o.UpdateDraft
	}
	return toSerialize, nil
}

type NullableActionsDraftUpdate struct {
	value *ActionsDraftUpdate
	isSet bool
}

func (v NullableActionsDraftUpdate) Get() *ActionsDraftUpdate {
	return v.value
}

func (v *NullableActionsDraftUpdate) Set(val *ActionsDraftUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableActionsDraftUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableActionsDraftUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionsDraftUpdate(val *ActionsDraftUpdate) *NullableActionsDraftUpdate {
	return &NullableActionsDraftUpdate{value: val, isSet: true}
}

func (v NullableActionsDraftUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionsDraftUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
