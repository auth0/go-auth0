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

// GetEnabledConnections200ResponseOneOfInner struct for GetEnabledConnections200ResponseOneOfInner
type GetEnabledConnections200ResponseOneOfInner struct {
	// ID of the connection.
	ConnectionId string `json:"connection_id"`
	// When true, all users that log in with this connection will be automatically granted membership in the organization. When false, users must be granted membership in the organization before logging in with this connection.
	AssignMembershipOnLogin bool `json:"assign_membership_on_login"`
	// Determines whether a connection should be displayed on this organization’s login prompt. Only applicable for enterprise connections. Default: true.
	ShowAsButton bool `json:"show_as_button"`
	// Determines whether organization signup should be enabled for this organization connection. Only applicable for database connections. Default: false.
	IsSignupEnabled bool                                                          `json:"is_signup_enabled"`
	Connection      PostOrganizations201ResponseEnabledConnectionsInnerConnection `json:"connection"`
}

type _GetEnabledConnections200ResponseOneOfInner GetEnabledConnections200ResponseOneOfInner

// GetConnectionId returns the ConnectionId field value
func (o *GetEnabledConnections200ResponseOneOfInner) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *GetEnabledConnections200ResponseOneOfInner) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *GetEnabledConnections200ResponseOneOfInner) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetAssignMembershipOnLogin returns the AssignMembershipOnLogin field value
func (o *GetEnabledConnections200ResponseOneOfInner) GetAssignMembershipOnLogin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AssignMembershipOnLogin
}

// GetAssignMembershipOnLoginOk returns a tuple with the AssignMembershipOnLogin field value
// and a boolean to check if the value has been set.
func (o *GetEnabledConnections200ResponseOneOfInner) GetAssignMembershipOnLoginOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignMembershipOnLogin, true
}

// SetAssignMembershipOnLogin sets field value
func (o *GetEnabledConnections200ResponseOneOfInner) SetAssignMembershipOnLogin(v bool) {
	o.AssignMembershipOnLogin = v
}

// GetShowAsButton returns the ShowAsButton field value
func (o *GetEnabledConnections200ResponseOneOfInner) GetShowAsButton() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowAsButton
}

// GetShowAsButtonOk returns a tuple with the ShowAsButton field value
// and a boolean to check if the value has been set.
func (o *GetEnabledConnections200ResponseOneOfInner) GetShowAsButtonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowAsButton, true
}

// SetShowAsButton sets field value
func (o *GetEnabledConnections200ResponseOneOfInner) SetShowAsButton(v bool) {
	o.ShowAsButton = v
}

// GetIsSignupEnabled returns the IsSignupEnabled field value
func (o *GetEnabledConnections200ResponseOneOfInner) GetIsSignupEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSignupEnabled
}

// GetIsSignupEnabledOk returns a tuple with the IsSignupEnabled field value
// and a boolean to check if the value has been set.
func (o *GetEnabledConnections200ResponseOneOfInner) GetIsSignupEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSignupEnabled, true
}

// SetIsSignupEnabled sets field value
func (o *GetEnabledConnections200ResponseOneOfInner) SetIsSignupEnabled(v bool) {
	o.IsSignupEnabled = v
}

// GetConnection returns the Connection field value
func (o *GetEnabledConnections200ResponseOneOfInner) GetConnection() PostOrganizations201ResponseEnabledConnectionsInnerConnection {
	if o == nil {
		var ret PostOrganizations201ResponseEnabledConnectionsInnerConnection
		return ret
	}

	return o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value
// and a boolean to check if the value has been set.
func (o *GetEnabledConnections200ResponseOneOfInner) GetConnectionOk() (*PostOrganizations201ResponseEnabledConnectionsInnerConnection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connection, true
}

// SetConnection sets field value
func (o *GetEnabledConnections200ResponseOneOfInner) SetConnection(v PostOrganizations201ResponseEnabledConnectionsInnerConnection) {
	o.Connection = v
}

func (o GetEnabledConnections200ResponseOneOfInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEnabledConnections200ResponseOneOfInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connection_id"] = o.ConnectionId
	toSerialize["assign_membership_on_login"] = o.AssignMembershipOnLogin
	toSerialize["show_as_button"] = o.ShowAsButton
	toSerialize["is_signup_enabled"] = o.IsSignupEnabled
	toSerialize["connection"] = o.Connection
	return toSerialize, nil
}

func (o *GetEnabledConnections200ResponseOneOfInner) UnmarshalJSON(data []byte) (err error) {
	varGetEnabledConnections200ResponseOneOfInner := _GetEnabledConnections200ResponseOneOfInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetEnabledConnections200ResponseOneOfInner)

	if err != nil {
		return err
	}

	*o = GetEnabledConnections200ResponseOneOfInner(varGetEnabledConnections200ResponseOneOfInner)

	return err
}

type NullableGetEnabledConnections200ResponseOneOfInner struct {
	value *GetEnabledConnections200ResponseOneOfInner
	isSet bool
}

func (v NullableGetEnabledConnections200ResponseOneOfInner) Get() *GetEnabledConnections200ResponseOneOfInner {
	return v.value
}

func (v *NullableGetEnabledConnections200ResponseOneOfInner) Set(val *GetEnabledConnections200ResponseOneOfInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEnabledConnections200ResponseOneOfInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEnabledConnections200ResponseOneOfInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEnabledConnections200ResponseOneOfInner(val *GetEnabledConnections200ResponseOneOfInner) *NullableGetEnabledConnections200ResponseOneOfInner {
	return &NullableGetEnabledConnections200ResponseOneOfInner{value: val, isSet: true}
}

func (v NullableGetEnabledConnections200ResponseOneOfInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEnabledConnections200ResponseOneOfInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
