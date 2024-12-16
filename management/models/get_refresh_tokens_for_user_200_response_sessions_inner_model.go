/*
Auth0 Management API

Auth0 Management API v2.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
)

// GetRefreshTokensForUser200ResponseSessionsInner struct for GetRefreshTokensForUser200ResponseSessionsInner
type GetRefreshTokensForUser200ResponseSessionsInner struct {
	// The ID of the refresh token
	Id string `json:"id"`
	// ID of the user which can be used when interacting with other APIs.
	UserId        string                                                       `json:"user_id"`
	CreatedAt     GetRefreshTokensForUser200ResponseSessionsInnerCreatedAt     `json:"created_at"`
	IdleExpiresAt GetRefreshTokensForUser200ResponseSessionsInnerIdleExpiresAt `json:"idle_expires_at"`
	ExpiresAt     GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt     `json:"expires_at"`
	// ID of the client application granted with this refresh token
	ClientId string `json:"client_id"`
	// ID of the authenticated session used to obtain this refresh-token
	SessionId string `json:"session_id"`
	// True if the token is a rotating refresh token
	Rotating bool `json:"rotating"`
	// A list of the resource server IDs associated to this refresh-token and their granted scopes
	ResourceServers      []GetRefreshToken200ResponseResourceServersInner `json:"resource_servers"`
	AdditionalProperties map[string]interface{}
}

type _GetRefreshTokensForUser200ResponseSessionsInner GetRefreshTokensForUser200ResponseSessionsInner

// GetId returns the Id field value
func (o *GetRefreshTokensForUser200ResponseSessionsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetRefreshTokensForUser200ResponseSessionsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetRefreshTokensForUser200ResponseSessionsInner) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *GetRefreshTokensForUser200ResponseSessionsInner) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *GetRefreshTokensForUser200ResponseSessionsInner) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *GetRefreshTokensForUser200ResponseSessionsInner) SetUserId(v string) {
	o.UserId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GetRefreshTokensForUser200ResponseSessionsInner) GetCreatedAt() GetRefreshTokensForUser200ResponseSessionsInnerCreatedAt {
	if o == nil {
		var ret GetRefreshTokensForUser200ResponseSessionsInnerCreatedAt
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GetRefreshTokensForUser200ResponseSessionsInner) GetCreatedAtOk() (*GetRefreshTokensForUser200ResponseSessionsInnerCreatedAt, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GetRefreshTokensForUser200ResponseSessionsInner) SetCreatedAt(v GetRefreshTokensForUser200ResponseSessionsInnerCreatedAt) {
	o.CreatedAt = v
}

// GetIdleExpiresAt returns the IdleExpiresAt field value
func (o *GetRefreshTokensForUser200ResponseSessionsInner) GetIdleExpiresAt() GetRefreshTokensForUser200ResponseSessionsInnerIdleExpiresAt {
	if o == nil {
		var ret GetRefreshTokensForUser200ResponseSessionsInnerIdleExpiresAt
		return ret
	}

	return o.IdleExpiresAt
}

// GetIdleExpiresAtOk returns a tuple with the IdleExpiresAt field value
// and a boolean to check if the value has been set.
func (o *GetRefreshTokensForUser200ResponseSessionsInner) GetIdleExpiresAtOk() (*GetRefreshTokensForUser200ResponseSessionsInnerIdleExpiresAt, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdleExpiresAt, true
}

// SetIdleExpiresAt sets field value
func (o *GetRefreshTokensForUser200ResponseSessionsInner) SetIdleExpiresAt(v GetRefreshTokensForUser200ResponseSessionsInnerIdleExpiresAt) {
	o.IdleExpiresAt = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *GetRefreshTokensForUser200ResponseSessionsInner) GetExpiresAt() GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt {
	if o == nil {
		var ret GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *GetRefreshTokensForUser200ResponseSessionsInner) GetExpiresAtOk() (*GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *GetRefreshTokensForUser200ResponseSessionsInner) SetExpiresAt(v GetRefreshTokensForUser200ResponseSessionsInnerExpiresAt) {
	o.ExpiresAt = v
}

// GetClientId returns the ClientId field value
func (o *GetRefreshTokensForUser200ResponseSessionsInner) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *GetRefreshTokensForUser200ResponseSessionsInner) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *GetRefreshTokensForUser200ResponseSessionsInner) SetClientId(v string) {
	o.ClientId = v
}

// GetSessionId returns the SessionId field value
func (o *GetRefreshTokensForUser200ResponseSessionsInner) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *GetRefreshTokensForUser200ResponseSessionsInner) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *GetRefreshTokensForUser200ResponseSessionsInner) SetSessionId(v string) {
	o.SessionId = v
}

// GetRotating returns the Rotating field value
func (o *GetRefreshTokensForUser200ResponseSessionsInner) GetRotating() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Rotating
}

// GetRotatingOk returns a tuple with the Rotating field value
// and a boolean to check if the value has been set.
func (o *GetRefreshTokensForUser200ResponseSessionsInner) GetRotatingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rotating, true
}

// SetRotating sets field value
func (o *GetRefreshTokensForUser200ResponseSessionsInner) SetRotating(v bool) {
	o.Rotating = v
}

// GetResourceServers returns the ResourceServers field value
func (o *GetRefreshTokensForUser200ResponseSessionsInner) GetResourceServers() []GetRefreshToken200ResponseResourceServersInner {
	if o == nil {
		var ret []GetRefreshToken200ResponseResourceServersInner
		return ret
	}

	return o.ResourceServers
}

// GetResourceServersOk returns a tuple with the ResourceServers field value
// and a boolean to check if the value has been set.
func (o *GetRefreshTokensForUser200ResponseSessionsInner) GetResourceServersOk() ([]GetRefreshToken200ResponseResourceServersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceServers, true
}

// SetResourceServers sets field value
func (o *GetRefreshTokensForUser200ResponseSessionsInner) SetResourceServers(v []GetRefreshToken200ResponseResourceServersInner) {
	o.ResourceServers = v
}

func (o GetRefreshTokensForUser200ResponseSessionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRefreshTokensForUser200ResponseSessionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["user_id"] = o.UserId
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["idle_expires_at"] = o.IdleExpiresAt
	toSerialize["expires_at"] = o.ExpiresAt
	toSerialize["client_id"] = o.ClientId
	toSerialize["session_id"] = o.SessionId
	toSerialize["rotating"] = o.Rotating
	toSerialize["resource_servers"] = o.ResourceServers

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetRefreshTokensForUser200ResponseSessionsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"user_id",
		"created_at",
		"idle_expires_at",
		"expires_at",
		"client_id",
		"session_id",
		"rotating",
		"resource_servers",
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

	varGetRefreshTokensForUser200ResponseSessionsInner := _GetRefreshTokensForUser200ResponseSessionsInner{}

	err = json.Unmarshal(data, &varGetRefreshTokensForUser200ResponseSessionsInner)

	if err != nil {
		return err
	}

	*o = GetRefreshTokensForUser200ResponseSessionsInner(varGetRefreshTokensForUser200ResponseSessionsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "idle_expires_at")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "session_id")
		delete(additionalProperties, "rotating")
		delete(additionalProperties, "resource_servers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRefreshTokensForUser200ResponseSessionsInner struct {
	value *GetRefreshTokensForUser200ResponseSessionsInner
	isSet bool
}

func (v NullableGetRefreshTokensForUser200ResponseSessionsInner) Get() *GetRefreshTokensForUser200ResponseSessionsInner {
	return v.value
}

func (v *NullableGetRefreshTokensForUser200ResponseSessionsInner) Set(val *GetRefreshTokensForUser200ResponseSessionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRefreshTokensForUser200ResponseSessionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRefreshTokensForUser200ResponseSessionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRefreshTokensForUser200ResponseSessionsInner(val *GetRefreshTokensForUser200ResponseSessionsInner) *NullableGetRefreshTokensForUser200ResponseSessionsInner {
	return &NullableGetRefreshTokensForUser200ResponseSessionsInner{value: val, isSet: true}
}

func (v NullableGetRefreshTokensForUser200ResponseSessionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRefreshTokensForUser200ResponseSessionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
