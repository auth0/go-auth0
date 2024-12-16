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

// GetBindings200ResponseBindingsInner Binding is the associative entity joining a trigger, and an action together.
type GetBindings200ResponseBindingsInner struct {
	// The unique ID of this binding.
	Id        string                                                    `json:"id"`
	TriggerId GetActions200ResponseActionsInnerSupportedTriggersInnerId `json:"trigger_id"`
	// The name of the binding.
	DisplayName string                            `json:"display_name"`
	Action      GetActions200ResponseActionsInner `json:"action"`
	// The time when the binding was created.
	CreatedAt string `json:"created_at"`
	// The time when the binding was updated.
	UpdatedAt string `json:"updated_at"`
}

type _GetBindings200ResponseBindingsInner GetBindings200ResponseBindingsInner

// GetId returns the Id field value
func (o *GetBindings200ResponseBindingsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetBindings200ResponseBindingsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetBindings200ResponseBindingsInner) SetId(v string) {
	o.Id = v
}

// GetTriggerId returns the TriggerId field value
func (o *GetBindings200ResponseBindingsInner) GetTriggerId() GetActions200ResponseActionsInnerSupportedTriggersInnerId {
	if o == nil {
		var ret GetActions200ResponseActionsInnerSupportedTriggersInnerId
		return ret
	}

	return o.TriggerId
}

// GetTriggerIdOk returns a tuple with the TriggerId field value
// and a boolean to check if the value has been set.
func (o *GetBindings200ResponseBindingsInner) GetTriggerIdOk() (*GetActions200ResponseActionsInnerSupportedTriggersInnerId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerId, true
}

// SetTriggerId sets field value
func (o *GetBindings200ResponseBindingsInner) SetTriggerId(v GetActions200ResponseActionsInnerSupportedTriggersInnerId) {
	o.TriggerId = v
}

// GetDisplayName returns the DisplayName field value
func (o *GetBindings200ResponseBindingsInner) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *GetBindings200ResponseBindingsInner) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *GetBindings200ResponseBindingsInner) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetAction returns the Action field value
func (o *GetBindings200ResponseBindingsInner) GetAction() GetActions200ResponseActionsInner {
	if o == nil {
		var ret GetActions200ResponseActionsInner
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *GetBindings200ResponseBindingsInner) GetActionOk() (*GetActions200ResponseActionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *GetBindings200ResponseBindingsInner) SetAction(v GetActions200ResponseActionsInner) {
	o.Action = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GetBindings200ResponseBindingsInner) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GetBindings200ResponseBindingsInner) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GetBindings200ResponseBindingsInner) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *GetBindings200ResponseBindingsInner) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *GetBindings200ResponseBindingsInner) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *GetBindings200ResponseBindingsInner) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o GetBindings200ResponseBindingsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBindings200ResponseBindingsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["trigger_id"] = o.TriggerId
	toSerialize["display_name"] = o.DisplayName
	toSerialize["action"] = o.Action
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *GetBindings200ResponseBindingsInner) UnmarshalJSON(data []byte) (err error) {
	varGetBindings200ResponseBindingsInner := _GetBindings200ResponseBindingsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetBindings200ResponseBindingsInner)

	if err != nil {
		return err
	}

	*o = GetBindings200ResponseBindingsInner(varGetBindings200ResponseBindingsInner)

	return err
}

type NullableGetBindings200ResponseBindingsInner struct {
	value *GetBindings200ResponseBindingsInner
	isSet bool
}

func (v NullableGetBindings200ResponseBindingsInner) Get() *GetBindings200ResponseBindingsInner {
	return v.value
}

func (v *NullableGetBindings200ResponseBindingsInner) Set(val *GetBindings200ResponseBindingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBindings200ResponseBindingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBindings200ResponseBindingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBindings200ResponseBindingsInner(val *GetBindings200ResponseBindingsInner) *NullableGetBindings200ResponseBindingsInner {
	return &NullableGetBindings200ResponseBindingsInner{value: val, isSet: true}
}

func (v NullableGetBindings200ResponseBindingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBindings200ResponseBindingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
