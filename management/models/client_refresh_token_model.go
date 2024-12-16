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

// ClientRefreshToken Refresh token configuration
type ClientRefreshToken struct {
	RotationType   ClientRefreshTokenRotationType   `json:"rotation_type"`
	ExpirationType ClientRefreshTokenExpirationType `json:"expiration_type"`
	// Period in seconds where the previous refresh token can be exchanged without triggering breach detection
	Leeway *int32 `json:"leeway,omitempty"`
	// Period (in seconds) for which refresh tokens will remain valid
	TokenLifetime *int32 `json:"token_lifetime,omitempty"`
	// Prevents tokens from having a set lifetime when `true` (takes precedence over `token_lifetime` values)
	InfiniteTokenLifetime *bool `json:"infinite_token_lifetime,omitempty"`
	// Period (in seconds) for which refresh tokens will remain valid without use
	IdleTokenLifetime *int32 `json:"idle_token_lifetime,omitempty"`
	// Prevents tokens from expiring without use when `true` (takes precedence over `idle_token_lifetime` values)
	InfiniteIdleTokenLifetime *bool `json:"infinite_idle_token_lifetime,omitempty"`
}

type _ClientRefreshToken ClientRefreshToken

// GetRotationType returns the RotationType field value
func (o *ClientRefreshToken) GetRotationType() ClientRefreshTokenRotationType {
	if o == nil {
		var ret ClientRefreshTokenRotationType
		return ret
	}

	return o.RotationType
}

// GetRotationTypeOk returns a tuple with the RotationType field value
// and a boolean to check if the value has been set.
func (o *ClientRefreshToken) GetRotationTypeOk() (*ClientRefreshTokenRotationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RotationType, true
}

// SetRotationType sets field value
func (o *ClientRefreshToken) SetRotationType(v ClientRefreshTokenRotationType) {
	o.RotationType = v
}

// GetExpirationType returns the ExpirationType field value
func (o *ClientRefreshToken) GetExpirationType() ClientRefreshTokenExpirationType {
	if o == nil {
		var ret ClientRefreshTokenExpirationType
		return ret
	}

	return o.ExpirationType
}

// GetExpirationTypeOk returns a tuple with the ExpirationType field value
// and a boolean to check if the value has been set.
func (o *ClientRefreshToken) GetExpirationTypeOk() (*ClientRefreshTokenExpirationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationType, true
}

// SetExpirationType sets field value
func (o *ClientRefreshToken) SetExpirationType(v ClientRefreshTokenExpirationType) {
	o.ExpirationType = v
}

// GetLeeway returns the Leeway field value if set, zero value otherwise.
func (o *ClientRefreshToken) GetLeeway() int32 {
	if o == nil || IsNil(o.Leeway) {
		var ret int32
		return ret
	}
	return *o.Leeway
}

// GetLeewayOk returns a tuple with the Leeway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRefreshToken) GetLeewayOk() (*int32, bool) {
	if o == nil || IsNil(o.Leeway) {
		return nil, false
	}
	return o.Leeway, true
}

// HasLeeway returns a boolean if a field has been set.
func (o *ClientRefreshToken) HasLeeway() bool {
	if o != nil && !IsNil(o.Leeway) {
		return true
	}

	return false
}

// SetLeeway gets a reference to the given int32 and assigns it to the Leeway field.
func (o *ClientRefreshToken) SetLeeway(v int32) {
	o.Leeway = &v
}

// GetTokenLifetime returns the TokenLifetime field value if set, zero value otherwise.
func (o *ClientRefreshToken) GetTokenLifetime() int32 {
	if o == nil || IsNil(o.TokenLifetime) {
		var ret int32
		return ret
	}
	return *o.TokenLifetime
}

// GetTokenLifetimeOk returns a tuple with the TokenLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRefreshToken) GetTokenLifetimeOk() (*int32, bool) {
	if o == nil || IsNil(o.TokenLifetime) {
		return nil, false
	}
	return o.TokenLifetime, true
}

// HasTokenLifetime returns a boolean if a field has been set.
func (o *ClientRefreshToken) HasTokenLifetime() bool {
	if o != nil && !IsNil(o.TokenLifetime) {
		return true
	}

	return false
}

// SetTokenLifetime gets a reference to the given int32 and assigns it to the TokenLifetime field.
func (o *ClientRefreshToken) SetTokenLifetime(v int32) {
	o.TokenLifetime = &v
}

// GetInfiniteTokenLifetime returns the InfiniteTokenLifetime field value if set, zero value otherwise.
func (o *ClientRefreshToken) GetInfiniteTokenLifetime() bool {
	if o == nil || IsNil(o.InfiniteTokenLifetime) {
		var ret bool
		return ret
	}
	return *o.InfiniteTokenLifetime
}

// GetInfiniteTokenLifetimeOk returns a tuple with the InfiniteTokenLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRefreshToken) GetInfiniteTokenLifetimeOk() (*bool, bool) {
	if o == nil || IsNil(o.InfiniteTokenLifetime) {
		return nil, false
	}
	return o.InfiniteTokenLifetime, true
}

// HasInfiniteTokenLifetime returns a boolean if a field has been set.
func (o *ClientRefreshToken) HasInfiniteTokenLifetime() bool {
	if o != nil && !IsNil(o.InfiniteTokenLifetime) {
		return true
	}

	return false
}

// SetInfiniteTokenLifetime gets a reference to the given bool and assigns it to the InfiniteTokenLifetime field.
func (o *ClientRefreshToken) SetInfiniteTokenLifetime(v bool) {
	o.InfiniteTokenLifetime = &v
}

// GetIdleTokenLifetime returns the IdleTokenLifetime field value if set, zero value otherwise.
func (o *ClientRefreshToken) GetIdleTokenLifetime() int32 {
	if o == nil || IsNil(o.IdleTokenLifetime) {
		var ret int32
		return ret
	}
	return *o.IdleTokenLifetime
}

// GetIdleTokenLifetimeOk returns a tuple with the IdleTokenLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRefreshToken) GetIdleTokenLifetimeOk() (*int32, bool) {
	if o == nil || IsNil(o.IdleTokenLifetime) {
		return nil, false
	}
	return o.IdleTokenLifetime, true
}

// HasIdleTokenLifetime returns a boolean if a field has been set.
func (o *ClientRefreshToken) HasIdleTokenLifetime() bool {
	if o != nil && !IsNil(o.IdleTokenLifetime) {
		return true
	}

	return false
}

// SetIdleTokenLifetime gets a reference to the given int32 and assigns it to the IdleTokenLifetime field.
func (o *ClientRefreshToken) SetIdleTokenLifetime(v int32) {
	o.IdleTokenLifetime = &v
}

// GetInfiniteIdleTokenLifetime returns the InfiniteIdleTokenLifetime field value if set, zero value otherwise.
func (o *ClientRefreshToken) GetInfiniteIdleTokenLifetime() bool {
	if o == nil || IsNil(o.InfiniteIdleTokenLifetime) {
		var ret bool
		return ret
	}
	return *o.InfiniteIdleTokenLifetime
}

// GetInfiniteIdleTokenLifetimeOk returns a tuple with the InfiniteIdleTokenLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRefreshToken) GetInfiniteIdleTokenLifetimeOk() (*bool, bool) {
	if o == nil || IsNil(o.InfiniteIdleTokenLifetime) {
		return nil, false
	}
	return o.InfiniteIdleTokenLifetime, true
}

// HasInfiniteIdleTokenLifetime returns a boolean if a field has been set.
func (o *ClientRefreshToken) HasInfiniteIdleTokenLifetime() bool {
	if o != nil && !IsNil(o.InfiniteIdleTokenLifetime) {
		return true
	}

	return false
}

// SetInfiniteIdleTokenLifetime gets a reference to the given bool and assigns it to the InfiniteIdleTokenLifetime field.
func (o *ClientRefreshToken) SetInfiniteIdleTokenLifetime(v bool) {
	o.InfiniteIdleTokenLifetime = &v
}

func (o ClientRefreshToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientRefreshToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rotation_type"] = o.RotationType
	toSerialize["expiration_type"] = o.ExpirationType
	if !IsNil(o.Leeway) {
		toSerialize["leeway"] = o.Leeway
	}
	if !IsNil(o.TokenLifetime) {
		toSerialize["token_lifetime"] = o.TokenLifetime
	}
	if !IsNil(o.InfiniteTokenLifetime) {
		toSerialize["infinite_token_lifetime"] = o.InfiniteTokenLifetime
	}
	if !IsNil(o.IdleTokenLifetime) {
		toSerialize["idle_token_lifetime"] = o.IdleTokenLifetime
	}
	if !IsNil(o.InfiniteIdleTokenLifetime) {
		toSerialize["infinite_idle_token_lifetime"] = o.InfiniteIdleTokenLifetime
	}
	return toSerialize, nil
}

func (o *ClientRefreshToken) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rotation_type",
		"expiration_type",
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

	varClientRefreshToken := _ClientRefreshToken{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClientRefreshToken)

	if err != nil {
		return err
	}

	*o = ClientRefreshToken(varClientRefreshToken)

	return err
}

type NullableClientRefreshToken struct {
	value *ClientRefreshToken
	isSet bool
}

func (v NullableClientRefreshToken) Get() *ClientRefreshToken {
	return v.value
}

func (v *NullableClientRefreshToken) Set(val *ClientRefreshToken) {
	v.value = val
	v.isSet = true
}

func (v NullableClientRefreshToken) IsSet() bool {
	return v.isSet
}

func (v *NullableClientRefreshToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientRefreshToken(val *ClientRefreshToken) *NullableClientRefreshToken {
	return &NullableClientRefreshToken{value: val, isSet: true}
}

func (v NullableClientRefreshToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientRefreshToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
