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
	"fmt"
)

// HookCreate struct for HookCreate
type HookCreate struct {
	// Name of this hook.
	Name string `json:"name"`
	// Code to be executed when this hook runs.
	Script string `json:"script"`
	// Whether this hook will be executed (true) or ignored (false).
	Enabled *bool `json:"enabled,omitempty"`
	// Dependencies of this hook used by webtask server.
	Dependencies map[string]interface{} `json:"dependencies,omitempty"`
	TriggerId    HookCreateTriggerId    `json:"triggerId"`
}

type _HookCreate HookCreate

// GetName returns the Name field value
func (o *HookCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HookCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HookCreate) SetName(v string) {
	o.Name = v
}

// GetScript returns the Script field value
func (o *HookCreate) GetScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *HookCreate) GetScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *HookCreate) SetScript(v string) {
	o.Script = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *HookCreate) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookCreate) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *HookCreate) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *HookCreate) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *HookCreate) GetDependencies() map[string]interface{} {
	if o == nil || IsNil(o.Dependencies) {
		var ret map[string]interface{}
		return ret
	}
	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookCreate) GetDependenciesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Dependencies) {
		return map[string]interface{}{}, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *HookCreate) HasDependencies() bool {
	if o != nil && !IsNil(o.Dependencies) {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given map[string]interface{} and assigns it to the Dependencies field.
func (o *HookCreate) SetDependencies(v map[string]interface{}) {
	o.Dependencies = v
}

// GetTriggerId returns the TriggerId field value
func (o *HookCreate) GetTriggerId() HookCreateTriggerId {
	if o == nil {
		var ret HookCreateTriggerId
		return ret
	}

	return o.TriggerId
}

// GetTriggerIdOk returns a tuple with the TriggerId field value
// and a boolean to check if the value has been set.
func (o *HookCreate) GetTriggerIdOk() (*HookCreateTriggerId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerId, true
}

// SetTriggerId sets field value
func (o *HookCreate) SetTriggerId(v HookCreateTriggerId) {
	o.TriggerId = v
}

func (o HookCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HookCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["script"] = o.Script
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Dependencies) {
		toSerialize["dependencies"] = o.Dependencies
	}
	toSerialize["triggerId"] = o.TriggerId
	return toSerialize, nil
}

func (o *HookCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"script",
		"triggerId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHookCreate := _HookCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHookCreate)

	if err != nil {
		return err
	}

	*o = HookCreate(varHookCreate)

	return err
}

type NullableHookCreate struct {
	value *HookCreate
	isSet bool
}

func (v NullableHookCreate) Get() *HookCreate {
	return v.value
}

func (v *NullableHookCreate) Set(val *HookCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableHookCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableHookCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHookCreate(val *HookCreate) *NullableHookCreate {
	return &NullableHookCreate{value: val, isSet: true}
}

func (v NullableHookCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHookCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
