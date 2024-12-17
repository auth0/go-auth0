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

// ClientAddonsSalesforce Salesforce SSO configuration.
type ClientAddonsSalesforce struct {
	// Arbitrary logical URL that identifies the Saleforce resource. e.g. `https://acme-org.com`.
	EntityId             *string `json:"entity_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientAddonsSalesforce ClientAddonsSalesforce

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *ClientAddonsSalesforce) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAddonsSalesforce) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *ClientAddonsSalesforce) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *ClientAddonsSalesforce) SetEntityId(v string) {
	o.EntityId = &v
}

func (o ClientAddonsSalesforce) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientAddonsSalesforce) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityId) {
		toSerialize["entity_id"] = o.EntityId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientAddonsSalesforce) UnmarshalJSON(data []byte) (err error) {
	varClientAddonsSalesforce := _ClientAddonsSalesforce{}

	err = json.Unmarshal(data, &varClientAddonsSalesforce)

	if err != nil {
		return err
	}

	*o = ClientAddonsSalesforce(varClientAddonsSalesforce)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entity_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientAddonsSalesforce struct {
	value *ClientAddonsSalesforce
	isSet bool
}

func (v NullableClientAddonsSalesforce) Get() *ClientAddonsSalesforce {
	return v.value
}

func (v *NullableClientAddonsSalesforce) Set(val *ClientAddonsSalesforce) {
	v.value = val
	v.isSet = true
}

func (v NullableClientAddonsSalesforce) IsSet() bool {
	return v.isSet
}

func (v *NullableClientAddonsSalesforce) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientAddonsSalesforce(val *ClientAddonsSalesforce) *NullableClientAddonsSalesforce {
	return &NullableClientAddonsSalesforce{value: val, isSet: true}
}

func (v NullableClientAddonsSalesforce) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientAddonsSalesforce) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
