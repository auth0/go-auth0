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

// GetActions200ResponseActionsInnerSupportedTriggersInner struct for GetActions200ResponseActionsInnerSupportedTriggersInner
type GetActions200ResponseActionsInnerSupportedTriggersInner struct {
	Id GetActions200ResponseActionsInnerSupportedTriggersInnerId `json:"id"`
	// The version of a trigger. v1, v2, etc.
	Version *string `json:"version,omitempty"`
	// status points to the trigger status.
	Status *string `json:"status,omitempty"`
	// runtimes supported by this trigger.
	Runtimes []string `json:"runtimes,omitempty"`
	// Runtime that will be used when none is specified when creating an action.
	DefaultRuntime *string `json:"default_runtime,omitempty"`
	// compatible_triggers informs which other trigger supports the same event and api.
	CompatibleTriggers []GetActions200ResponseActionsInnerSupportedTriggersInnerCompatibleTriggersInner `json:"compatible_triggers,omitempty"`
	BindingPolicy      *GetActions200ResponseActionsInnerSupportedTriggersInnerBindingPolicy            `json:"binding_policy,omitempty"`
}

type _GetActions200ResponseActionsInnerSupportedTriggersInner GetActions200ResponseActionsInnerSupportedTriggersInner

// GetId returns the Id field value
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) GetId() GetActions200ResponseActionsInnerSupportedTriggersInnerId {
	if o == nil {
		var ret GetActions200ResponseActionsInnerSupportedTriggersInnerId
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) GetIdOk() (*GetActions200ResponseActionsInnerSupportedTriggersInnerId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) SetId(v GetActions200ResponseActionsInnerSupportedTriggersInnerId) {
	o.Id = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) SetVersion(v string) {
	o.Version = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) SetStatus(v string) {
	o.Status = &v
}

// GetRuntimes returns the Runtimes field value if set, zero value otherwise.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) GetRuntimes() []string {
	if o == nil || IsNil(o.Runtimes) {
		var ret []string
		return ret
	}
	return o.Runtimes
}

// GetRuntimesOk returns a tuple with the Runtimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) GetRuntimesOk() ([]string, bool) {
	if o == nil || IsNil(o.Runtimes) {
		return nil, false
	}
	return o.Runtimes, true
}

// HasRuntimes returns a boolean if a field has been set.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) HasRuntimes() bool {
	if o != nil && !IsNil(o.Runtimes) {
		return true
	}

	return false
}

// SetRuntimes gets a reference to the given []string and assigns it to the Runtimes field.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) SetRuntimes(v []string) {
	o.Runtimes = v
}

// GetDefaultRuntime returns the DefaultRuntime field value if set, zero value otherwise.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) GetDefaultRuntime() string {
	if o == nil || IsNil(o.DefaultRuntime) {
		var ret string
		return ret
	}
	return *o.DefaultRuntime
}

// GetDefaultRuntimeOk returns a tuple with the DefaultRuntime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) GetDefaultRuntimeOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultRuntime) {
		return nil, false
	}
	return o.DefaultRuntime, true
}

// HasDefaultRuntime returns a boolean if a field has been set.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) HasDefaultRuntime() bool {
	if o != nil && !IsNil(o.DefaultRuntime) {
		return true
	}

	return false
}

// SetDefaultRuntime gets a reference to the given string and assigns it to the DefaultRuntime field.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) SetDefaultRuntime(v string) {
	o.DefaultRuntime = &v
}

// GetCompatibleTriggers returns the CompatibleTriggers field value if set, zero value otherwise.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) GetCompatibleTriggers() []GetActions200ResponseActionsInnerSupportedTriggersInnerCompatibleTriggersInner {
	if o == nil || IsNil(o.CompatibleTriggers) {
		var ret []GetActions200ResponseActionsInnerSupportedTriggersInnerCompatibleTriggersInner
		return ret
	}
	return o.CompatibleTriggers
}

// GetCompatibleTriggersOk returns a tuple with the CompatibleTriggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) GetCompatibleTriggersOk() ([]GetActions200ResponseActionsInnerSupportedTriggersInnerCompatibleTriggersInner, bool) {
	if o == nil || IsNil(o.CompatibleTriggers) {
		return nil, false
	}
	return o.CompatibleTriggers, true
}

// HasCompatibleTriggers returns a boolean if a field has been set.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) HasCompatibleTriggers() bool {
	if o != nil && !IsNil(o.CompatibleTriggers) {
		return true
	}

	return false
}

// SetCompatibleTriggers gets a reference to the given []GetActions200ResponseActionsInnerSupportedTriggersInnerCompatibleTriggersInner and assigns it to the CompatibleTriggers field.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) SetCompatibleTriggers(v []GetActions200ResponseActionsInnerSupportedTriggersInnerCompatibleTriggersInner) {
	o.CompatibleTriggers = v
}

// GetBindingPolicy returns the BindingPolicy field value if set, zero value otherwise.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) GetBindingPolicy() GetActions200ResponseActionsInnerSupportedTriggersInnerBindingPolicy {
	if o == nil || IsNil(o.BindingPolicy) {
		var ret GetActions200ResponseActionsInnerSupportedTriggersInnerBindingPolicy
		return ret
	}
	return *o.BindingPolicy
}

// GetBindingPolicyOk returns a tuple with the BindingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) GetBindingPolicyOk() (*GetActions200ResponseActionsInnerSupportedTriggersInnerBindingPolicy, bool) {
	if o == nil || IsNil(o.BindingPolicy) {
		return nil, false
	}
	return o.BindingPolicy, true
}

// HasBindingPolicy returns a boolean if a field has been set.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) HasBindingPolicy() bool {
	if o != nil && !IsNil(o.BindingPolicy) {
		return true
	}

	return false
}

// SetBindingPolicy gets a reference to the given GetActions200ResponseActionsInnerSupportedTriggersInnerBindingPolicy and assigns it to the BindingPolicy field.
func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) SetBindingPolicy(v GetActions200ResponseActionsInnerSupportedTriggersInnerBindingPolicy) {
	o.BindingPolicy = &v
}

func (o GetActions200ResponseActionsInnerSupportedTriggersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetActions200ResponseActionsInnerSupportedTriggersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Runtimes) {
		toSerialize["runtimes"] = o.Runtimes
	}
	if !IsNil(o.DefaultRuntime) {
		toSerialize["default_runtime"] = o.DefaultRuntime
	}
	if !IsNil(o.CompatibleTriggers) {
		toSerialize["compatible_triggers"] = o.CompatibleTriggers
	}
	if !IsNil(o.BindingPolicy) {
		toSerialize["binding_policy"] = o.BindingPolicy
	}
	return toSerialize, nil
}

func (o *GetActions200ResponseActionsInnerSupportedTriggersInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varGetActions200ResponseActionsInnerSupportedTriggersInner := _GetActions200ResponseActionsInnerSupportedTriggersInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetActions200ResponseActionsInnerSupportedTriggersInner)

	if err != nil {
		return err
	}

	*o = GetActions200ResponseActionsInnerSupportedTriggersInner(varGetActions200ResponseActionsInnerSupportedTriggersInner)

	return err
}

type NullableGetActions200ResponseActionsInnerSupportedTriggersInner struct {
	value *GetActions200ResponseActionsInnerSupportedTriggersInner
	isSet bool
}

func (v NullableGetActions200ResponseActionsInnerSupportedTriggersInner) Get() *GetActions200ResponseActionsInnerSupportedTriggersInner {
	return v.value
}

func (v *NullableGetActions200ResponseActionsInnerSupportedTriggersInner) Set(val *GetActions200ResponseActionsInnerSupportedTriggersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetActions200ResponseActionsInnerSupportedTriggersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetActions200ResponseActionsInnerSupportedTriggersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetActions200ResponseActionsInnerSupportedTriggersInner(val *GetActions200ResponseActionsInnerSupportedTriggersInner) *NullableGetActions200ResponseActionsInnerSupportedTriggersInner {
	return &NullableGetActions200ResponseActionsInnerSupportedTriggersInner{value: val, isSet: true}
}

func (v NullableGetActions200ResponseActionsInnerSupportedTriggersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetActions200ResponseActionsInnerSupportedTriggersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
