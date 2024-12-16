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

// ClientCreateAddonsAzureSb Azure Storage Bus addon configuration.
type ClientCreateAddonsAzureSb struct {
	// Your Azure Service Bus namespace. Usually the first segment of your Service Bus URL (e.g. `https://acme-org.servicebus.windows.net` would be `acme-org`).
	Namespace *string `json:"namespace,omitempty"`
	// Your shared access policy name defined in your Service Bus entity.
	SasKeyName *string `json:"sasKeyName,omitempty"`
	// Primary Key associated with your shared access policy.
	SasKey *string `json:"sasKey,omitempty"`
	// Entity you want to request a token for. e.g. `my-queue`.'
	EntityPath *string `json:"entityPath,omitempty"`
	// Optional expiration in minutes for the generated token. Defaults to 5 minutes.
	Expiration           *int32 `json:"expiration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientCreateAddonsAzureSb ClientCreateAddonsAzureSb

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ClientCreateAddonsAzureSb) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsAzureSb) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ClientCreateAddonsAzureSb) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ClientCreateAddonsAzureSb) SetNamespace(v string) {
	o.Namespace = &v
}

// GetSasKeyName returns the SasKeyName field value if set, zero value otherwise.
func (o *ClientCreateAddonsAzureSb) GetSasKeyName() string {
	if o == nil || IsNil(o.SasKeyName) {
		var ret string
		return ret
	}
	return *o.SasKeyName
}

// GetSasKeyNameOk returns a tuple with the SasKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsAzureSb) GetSasKeyNameOk() (*string, bool) {
	if o == nil || IsNil(o.SasKeyName) {
		return nil, false
	}
	return o.SasKeyName, true
}

// HasSasKeyName returns a boolean if a field has been set.
func (o *ClientCreateAddonsAzureSb) HasSasKeyName() bool {
	if o != nil && !IsNil(o.SasKeyName) {
		return true
	}

	return false
}

// SetSasKeyName gets a reference to the given string and assigns it to the SasKeyName field.
func (o *ClientCreateAddonsAzureSb) SetSasKeyName(v string) {
	o.SasKeyName = &v
}

// GetSasKey returns the SasKey field value if set, zero value otherwise.
func (o *ClientCreateAddonsAzureSb) GetSasKey() string {
	if o == nil || IsNil(o.SasKey) {
		var ret string
		return ret
	}
	return *o.SasKey
}

// GetSasKeyOk returns a tuple with the SasKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsAzureSb) GetSasKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SasKey) {
		return nil, false
	}
	return o.SasKey, true
}

// HasSasKey returns a boolean if a field has been set.
func (o *ClientCreateAddonsAzureSb) HasSasKey() bool {
	if o != nil && !IsNil(o.SasKey) {
		return true
	}

	return false
}

// SetSasKey gets a reference to the given string and assigns it to the SasKey field.
func (o *ClientCreateAddonsAzureSb) SetSasKey(v string) {
	o.SasKey = &v
}

// GetEntityPath returns the EntityPath field value if set, zero value otherwise.
func (o *ClientCreateAddonsAzureSb) GetEntityPath() string {
	if o == nil || IsNil(o.EntityPath) {
		var ret string
		return ret
	}
	return *o.EntityPath
}

// GetEntityPathOk returns a tuple with the EntityPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsAzureSb) GetEntityPathOk() (*string, bool) {
	if o == nil || IsNil(o.EntityPath) {
		return nil, false
	}
	return o.EntityPath, true
}

// HasEntityPath returns a boolean if a field has been set.
func (o *ClientCreateAddonsAzureSb) HasEntityPath() bool {
	if o != nil && !IsNil(o.EntityPath) {
		return true
	}

	return false
}

// SetEntityPath gets a reference to the given string and assigns it to the EntityPath field.
func (o *ClientCreateAddonsAzureSb) SetEntityPath(v string) {
	o.EntityPath = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *ClientCreateAddonsAzureSb) GetExpiration() int32 {
	if o == nil || IsNil(o.Expiration) {
		var ret int32
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCreateAddonsAzureSb) GetExpirationOk() (*int32, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *ClientCreateAddonsAzureSb) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given int32 and assigns it to the Expiration field.
func (o *ClientCreateAddonsAzureSb) SetExpiration(v int32) {
	o.Expiration = &v
}

func (o ClientCreateAddonsAzureSb) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientCreateAddonsAzureSb) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.SasKeyName) {
		toSerialize["sasKeyName"] = o.SasKeyName
	}
	if !IsNil(o.SasKey) {
		toSerialize["sasKey"] = o.SasKey
	}
	if !IsNil(o.EntityPath) {
		toSerialize["entityPath"] = o.EntityPath
	}
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientCreateAddonsAzureSb) UnmarshalJSON(data []byte) (err error) {
	varClientCreateAddonsAzureSb := _ClientCreateAddonsAzureSb{}

	err = json.Unmarshal(data, &varClientCreateAddonsAzureSb)

	if err != nil {
		return err
	}

	*o = ClientCreateAddonsAzureSb(varClientCreateAddonsAzureSb)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "namespace")
		delete(additionalProperties, "sasKeyName")
		delete(additionalProperties, "sasKey")
		delete(additionalProperties, "entityPath")
		delete(additionalProperties, "expiration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientCreateAddonsAzureSb struct {
	value *ClientCreateAddonsAzureSb
	isSet bool
}

func (v NullableClientCreateAddonsAzureSb) Get() *ClientCreateAddonsAzureSb {
	return v.value
}

func (v *NullableClientCreateAddonsAzureSb) Set(val *ClientCreateAddonsAzureSb) {
	v.value = val
	v.isSet = true
}

func (v NullableClientCreateAddonsAzureSb) IsSet() bool {
	return v.isSet
}

func (v *NullableClientCreateAddonsAzureSb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientCreateAddonsAzureSb(val *ClientCreateAddonsAzureSb) *NullableClientCreateAddonsAzureSb {
	return &NullableClientCreateAddonsAzureSb{value: val, isSet: true}
}

func (v NullableClientCreateAddonsAzureSb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientCreateAddonsAzureSb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
