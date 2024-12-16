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

// GetSessionsForUser200ResponseSessionsInner struct for GetSessionsForUser200ResponseSessionsInner
type GetSessionsForUser200ResponseSessionsInner struct {
	// The ID of the session
	Id string `json:"id"`
	// ID of the user which can be used when interacting with other APIs.
	UserId          string                                                    `json:"user_id"`
	CreatedAt       GetSessionsForUser200ResponseSessionsInnerCreatedAt       `json:"created_at"`
	UpdatedAt       GetSessionsForUser200ResponseSessionsInnerUpdatedAt       `json:"updated_at"`
	AuthenticatedAt GetSessionsForUser200ResponseSessionsInnerAuthenticatedAt `json:"authenticated_at"`
	IdleExpiresAt   GetSessionsForUser200ResponseSessionsInnerIdleExpiresAt   `json:"idle_expires_at"`
	ExpiresAt       GetSessionsForUser200ResponseSessionsInnerExpiresAt       `json:"expires_at"`
	Device          GetSessionsForUser200ResponseSessionsInnerDevice          `json:"device"`
	// List of client details for the session
	Clients              []GetSession200ResponseClientsInner                      `json:"clients"`
	Authentication       GetSessionsForUser200ResponseSessionsInnerAuthentication `json:"authentication"`
	AdditionalProperties map[string]interface{}
}

type _GetSessionsForUser200ResponseSessionsInner GetSessionsForUser200ResponseSessionsInner

// GetId returns the Id field value
func (o *GetSessionsForUser200ResponseSessionsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetSessionsForUser200ResponseSessionsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetSessionsForUser200ResponseSessionsInner) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *GetSessionsForUser200ResponseSessionsInner) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *GetSessionsForUser200ResponseSessionsInner) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *GetSessionsForUser200ResponseSessionsInner) SetUserId(v string) {
	o.UserId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GetSessionsForUser200ResponseSessionsInner) GetCreatedAt() GetSessionsForUser200ResponseSessionsInnerCreatedAt {
	if o == nil {
		var ret GetSessionsForUser200ResponseSessionsInnerCreatedAt
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GetSessionsForUser200ResponseSessionsInner) GetCreatedAtOk() (*GetSessionsForUser200ResponseSessionsInnerCreatedAt, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GetSessionsForUser200ResponseSessionsInner) SetCreatedAt(v GetSessionsForUser200ResponseSessionsInnerCreatedAt) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *GetSessionsForUser200ResponseSessionsInner) GetUpdatedAt() GetSessionsForUser200ResponseSessionsInnerUpdatedAt {
	if o == nil {
		var ret GetSessionsForUser200ResponseSessionsInnerUpdatedAt
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *GetSessionsForUser200ResponseSessionsInner) GetUpdatedAtOk() (*GetSessionsForUser200ResponseSessionsInnerUpdatedAt, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *GetSessionsForUser200ResponseSessionsInner) SetUpdatedAt(v GetSessionsForUser200ResponseSessionsInnerUpdatedAt) {
	o.UpdatedAt = v
}

// GetAuthenticatedAt returns the AuthenticatedAt field value
func (o *GetSessionsForUser200ResponseSessionsInner) GetAuthenticatedAt() GetSessionsForUser200ResponseSessionsInnerAuthenticatedAt {
	if o == nil {
		var ret GetSessionsForUser200ResponseSessionsInnerAuthenticatedAt
		return ret
	}

	return o.AuthenticatedAt
}

// GetAuthenticatedAtOk returns a tuple with the AuthenticatedAt field value
// and a boolean to check if the value has been set.
func (o *GetSessionsForUser200ResponseSessionsInner) GetAuthenticatedAtOk() (*GetSessionsForUser200ResponseSessionsInnerAuthenticatedAt, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticatedAt, true
}

// SetAuthenticatedAt sets field value
func (o *GetSessionsForUser200ResponseSessionsInner) SetAuthenticatedAt(v GetSessionsForUser200ResponseSessionsInnerAuthenticatedAt) {
	o.AuthenticatedAt = v
}

// GetIdleExpiresAt returns the IdleExpiresAt field value
func (o *GetSessionsForUser200ResponseSessionsInner) GetIdleExpiresAt() GetSessionsForUser200ResponseSessionsInnerIdleExpiresAt {
	if o == nil {
		var ret GetSessionsForUser200ResponseSessionsInnerIdleExpiresAt
		return ret
	}

	return o.IdleExpiresAt
}

// GetIdleExpiresAtOk returns a tuple with the IdleExpiresAt field value
// and a boolean to check if the value has been set.
func (o *GetSessionsForUser200ResponseSessionsInner) GetIdleExpiresAtOk() (*GetSessionsForUser200ResponseSessionsInnerIdleExpiresAt, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdleExpiresAt, true
}

// SetIdleExpiresAt sets field value
func (o *GetSessionsForUser200ResponseSessionsInner) SetIdleExpiresAt(v GetSessionsForUser200ResponseSessionsInnerIdleExpiresAt) {
	o.IdleExpiresAt = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *GetSessionsForUser200ResponseSessionsInner) GetExpiresAt() GetSessionsForUser200ResponseSessionsInnerExpiresAt {
	if o == nil {
		var ret GetSessionsForUser200ResponseSessionsInnerExpiresAt
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *GetSessionsForUser200ResponseSessionsInner) GetExpiresAtOk() (*GetSessionsForUser200ResponseSessionsInnerExpiresAt, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *GetSessionsForUser200ResponseSessionsInner) SetExpiresAt(v GetSessionsForUser200ResponseSessionsInnerExpiresAt) {
	o.ExpiresAt = v
}

// GetDevice returns the Device field value
func (o *GetSessionsForUser200ResponseSessionsInner) GetDevice() GetSessionsForUser200ResponseSessionsInnerDevice {
	if o == nil {
		var ret GetSessionsForUser200ResponseSessionsInnerDevice
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *GetSessionsForUser200ResponseSessionsInner) GetDeviceOk() (*GetSessionsForUser200ResponseSessionsInnerDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *GetSessionsForUser200ResponseSessionsInner) SetDevice(v GetSessionsForUser200ResponseSessionsInnerDevice) {
	o.Device = v
}

// GetClients returns the Clients field value
func (o *GetSessionsForUser200ResponseSessionsInner) GetClients() []GetSession200ResponseClientsInner {
	if o == nil {
		var ret []GetSession200ResponseClientsInner
		return ret
	}

	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value
// and a boolean to check if the value has been set.
func (o *GetSessionsForUser200ResponseSessionsInner) GetClientsOk() ([]GetSession200ResponseClientsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Clients, true
}

// SetClients sets field value
func (o *GetSessionsForUser200ResponseSessionsInner) SetClients(v []GetSession200ResponseClientsInner) {
	o.Clients = v
}

// GetAuthentication returns the Authentication field value
func (o *GetSessionsForUser200ResponseSessionsInner) GetAuthentication() GetSessionsForUser200ResponseSessionsInnerAuthentication {
	if o == nil {
		var ret GetSessionsForUser200ResponseSessionsInnerAuthentication
		return ret
	}

	return o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value
// and a boolean to check if the value has been set.
func (o *GetSessionsForUser200ResponseSessionsInner) GetAuthenticationOk() (*GetSessionsForUser200ResponseSessionsInnerAuthentication, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authentication, true
}

// SetAuthentication sets field value
func (o *GetSessionsForUser200ResponseSessionsInner) SetAuthentication(v GetSessionsForUser200ResponseSessionsInnerAuthentication) {
	o.Authentication = v
}

func (o GetSessionsForUser200ResponseSessionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSessionsForUser200ResponseSessionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["user_id"] = o.UserId
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["authenticated_at"] = o.AuthenticatedAt
	toSerialize["idle_expires_at"] = o.IdleExpiresAt
	toSerialize["expires_at"] = o.ExpiresAt
	toSerialize["device"] = o.Device
	toSerialize["clients"] = o.Clients
	toSerialize["authentication"] = o.Authentication

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSessionsForUser200ResponseSessionsInner) UnmarshalJSON(data []byte) (err error) {
	varGetSessionsForUser200ResponseSessionsInner := _GetSessionsForUser200ResponseSessionsInner{}

	err = json.Unmarshal(data, &varGetSessionsForUser200ResponseSessionsInner)

	if err != nil {
		return err
	}

	*o = GetSessionsForUser200ResponseSessionsInner(varGetSessionsForUser200ResponseSessionsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "authenticated_at")
		delete(additionalProperties, "idle_expires_at")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "device")
		delete(additionalProperties, "clients")
		delete(additionalProperties, "authentication")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSessionsForUser200ResponseSessionsInner struct {
	value *GetSessionsForUser200ResponseSessionsInner
	isSet bool
}

func (v NullableGetSessionsForUser200ResponseSessionsInner) Get() *GetSessionsForUser200ResponseSessionsInner {
	return v.value
}

func (v *NullableGetSessionsForUser200ResponseSessionsInner) Set(val *GetSessionsForUser200ResponseSessionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSessionsForUser200ResponseSessionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSessionsForUser200ResponseSessionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSessionsForUser200ResponseSessionsInner(val *GetSessionsForUser200ResponseSessionsInner) *NullableGetSessionsForUser200ResponseSessionsInner {
	return &NullableGetSessionsForUser200ResponseSessionsInner{value: val, isSet: true}
}

func (v NullableGetSessionsForUser200ResponseSessionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSessionsForUser200ResponseSessionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
