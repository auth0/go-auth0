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

// RuleCreate struct for RuleCreate
type RuleCreate struct {
	// Name of this rule.
	Name string `json:"name"`
	// Code to be executed when this rule runs.
	Script string `json:"script"`
	// Order that this rule should execute in relative to other rules. Lower-valued rules execute first.
	Order *float32 `json:"order,omitempty"`
	// Whether the rule is enabled (true), or disabled (false).
	Enabled *bool `json:"enabled,omitempty"`
}

type _RuleCreate RuleCreate

// GetName returns the Name field value
func (o *RuleCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RuleCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RuleCreate) SetName(v string) {
	o.Name = v
}

// GetScript returns the Script field value
func (o *RuleCreate) GetScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *RuleCreate) GetScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *RuleCreate) SetScript(v string) {
	o.Script = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *RuleCreate) GetOrder() float32 {
	if o == nil || IsNil(o.Order) {
		var ret float32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleCreate) GetOrderOk() (*float32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *RuleCreate) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given float32 and assigns it to the Order field.
func (o *RuleCreate) SetOrder(v float32) {
	o.Order = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *RuleCreate) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleCreate) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *RuleCreate) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *RuleCreate) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o RuleCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["script"] = o.Script
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

func (o *RuleCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"script",
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

	varRuleCreate := _RuleCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varRuleCreate)

	if err != nil {
		return err
	}

	*o = RuleCreate(varRuleCreate)

	return err
}

type NullableRuleCreate struct {
	value *RuleCreate
	isSet bool
}

func (v NullableRuleCreate) Get() *RuleCreate {
	return v.value
}

func (v *NullableRuleCreate) Set(val *RuleCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleCreate(val *RuleCreate) *NullableRuleCreate {
	return &NullableRuleCreate{value: val, isSet: true}
}

func (v NullableRuleCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
