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

// ConnectionCreateOptionsCustomScripts A map of scripts used to integrate with a custom database.
type ConnectionCreateOptionsCustomScripts struct {
	Login                *string `json:"login,omitempty"`
	GetUser              *string `json:"get_user,omitempty"`
	Delete               *string `json:"delete,omitempty"`
	ChangePassword       *string `json:"change_password,omitempty"`
	Verify               *string `json:"verify,omitempty"`
	Create               *string `json:"create,omitempty"`
	ChangeUsername       *string `json:"change_username,omitempty"`
	ChangeEmail          *string `json:"change_email,omitempty"`
	ChangePhoneNumber    *string `json:"change_phone_number,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectionCreateOptionsCustomScripts ConnectionCreateOptionsCustomScripts

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *ConnectionCreateOptionsCustomScripts) GetLogin() string {
	if o == nil || IsNil(o.Login) {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCreateOptionsCustomScripts) GetLoginOk() (*string, bool) {
	if o == nil || IsNil(o.Login) {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *ConnectionCreateOptionsCustomScripts) HasLogin() bool {
	if o != nil && !IsNil(o.Login) {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *ConnectionCreateOptionsCustomScripts) SetLogin(v string) {
	o.Login = &v
}

// GetGetUser returns the GetUser field value if set, zero value otherwise.
func (o *ConnectionCreateOptionsCustomScripts) GetGetUser() string {
	if o == nil || IsNil(o.GetUser) {
		var ret string
		return ret
	}
	return *o.GetUser
}

// GetGetUserOk returns a tuple with the GetUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCreateOptionsCustomScripts) GetGetUserOk() (*string, bool) {
	if o == nil || IsNil(o.GetUser) {
		return nil, false
	}
	return o.GetUser, true
}

// HasGetUser returns a boolean if a field has been set.
func (o *ConnectionCreateOptionsCustomScripts) HasGetUser() bool {
	if o != nil && !IsNil(o.GetUser) {
		return true
	}

	return false
}

// SetGetUser gets a reference to the given string and assigns it to the GetUser field.
func (o *ConnectionCreateOptionsCustomScripts) SetGetUser(v string) {
	o.GetUser = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *ConnectionCreateOptionsCustomScripts) GetDelete() string {
	if o == nil || IsNil(o.Delete) {
		var ret string
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCreateOptionsCustomScripts) GetDeleteOk() (*string, bool) {
	if o == nil || IsNil(o.Delete) {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *ConnectionCreateOptionsCustomScripts) HasDelete() bool {
	if o != nil && !IsNil(o.Delete) {
		return true
	}

	return false
}

// SetDelete gets a reference to the given string and assigns it to the Delete field.
func (o *ConnectionCreateOptionsCustomScripts) SetDelete(v string) {
	o.Delete = &v
}

// GetChangePassword returns the ChangePassword field value if set, zero value otherwise.
func (o *ConnectionCreateOptionsCustomScripts) GetChangePassword() string {
	if o == nil || IsNil(o.ChangePassword) {
		var ret string
		return ret
	}
	return *o.ChangePassword
}

// GetChangePasswordOk returns a tuple with the ChangePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCreateOptionsCustomScripts) GetChangePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.ChangePassword) {
		return nil, false
	}
	return o.ChangePassword, true
}

// HasChangePassword returns a boolean if a field has been set.
func (o *ConnectionCreateOptionsCustomScripts) HasChangePassword() bool {
	if o != nil && !IsNil(o.ChangePassword) {
		return true
	}

	return false
}

// SetChangePassword gets a reference to the given string and assigns it to the ChangePassword field.
func (o *ConnectionCreateOptionsCustomScripts) SetChangePassword(v string) {
	o.ChangePassword = &v
}

// GetVerify returns the Verify field value if set, zero value otherwise.
func (o *ConnectionCreateOptionsCustomScripts) GetVerify() string {
	if o == nil || IsNil(o.Verify) {
		var ret string
		return ret
	}
	return *o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCreateOptionsCustomScripts) GetVerifyOk() (*string, bool) {
	if o == nil || IsNil(o.Verify) {
		return nil, false
	}
	return o.Verify, true
}

// HasVerify returns a boolean if a field has been set.
func (o *ConnectionCreateOptionsCustomScripts) HasVerify() bool {
	if o != nil && !IsNil(o.Verify) {
		return true
	}

	return false
}

// SetVerify gets a reference to the given string and assigns it to the Verify field.
func (o *ConnectionCreateOptionsCustomScripts) SetVerify(v string) {
	o.Verify = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *ConnectionCreateOptionsCustomScripts) GetCreate() string {
	if o == nil || IsNil(o.Create) {
		var ret string
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCreateOptionsCustomScripts) GetCreateOk() (*string, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *ConnectionCreateOptionsCustomScripts) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given string and assigns it to the Create field.
func (o *ConnectionCreateOptionsCustomScripts) SetCreate(v string) {
	o.Create = &v
}

// GetChangeUsername returns the ChangeUsername field value if set, zero value otherwise.
func (o *ConnectionCreateOptionsCustomScripts) GetChangeUsername() string {
	if o == nil || IsNil(o.ChangeUsername) {
		var ret string
		return ret
	}
	return *o.ChangeUsername
}

// GetChangeUsernameOk returns a tuple with the ChangeUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCreateOptionsCustomScripts) GetChangeUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeUsername) {
		return nil, false
	}
	return o.ChangeUsername, true
}

// HasChangeUsername returns a boolean if a field has been set.
func (o *ConnectionCreateOptionsCustomScripts) HasChangeUsername() bool {
	if o != nil && !IsNil(o.ChangeUsername) {
		return true
	}

	return false
}

// SetChangeUsername gets a reference to the given string and assigns it to the ChangeUsername field.
func (o *ConnectionCreateOptionsCustomScripts) SetChangeUsername(v string) {
	o.ChangeUsername = &v
}

// GetChangeEmail returns the ChangeEmail field value if set, zero value otherwise.
func (o *ConnectionCreateOptionsCustomScripts) GetChangeEmail() string {
	if o == nil || IsNil(o.ChangeEmail) {
		var ret string
		return ret
	}
	return *o.ChangeEmail
}

// GetChangeEmailOk returns a tuple with the ChangeEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCreateOptionsCustomScripts) GetChangeEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeEmail) {
		return nil, false
	}
	return o.ChangeEmail, true
}

// HasChangeEmail returns a boolean if a field has been set.
func (o *ConnectionCreateOptionsCustomScripts) HasChangeEmail() bool {
	if o != nil && !IsNil(o.ChangeEmail) {
		return true
	}

	return false
}

// SetChangeEmail gets a reference to the given string and assigns it to the ChangeEmail field.
func (o *ConnectionCreateOptionsCustomScripts) SetChangeEmail(v string) {
	o.ChangeEmail = &v
}

// GetChangePhoneNumber returns the ChangePhoneNumber field value if set, zero value otherwise.
func (o *ConnectionCreateOptionsCustomScripts) GetChangePhoneNumber() string {
	if o == nil || IsNil(o.ChangePhoneNumber) {
		var ret string
		return ret
	}
	return *o.ChangePhoneNumber
}

// GetChangePhoneNumberOk returns a tuple with the ChangePhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCreateOptionsCustomScripts) GetChangePhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ChangePhoneNumber) {
		return nil, false
	}
	return o.ChangePhoneNumber, true
}

// HasChangePhoneNumber returns a boolean if a field has been set.
func (o *ConnectionCreateOptionsCustomScripts) HasChangePhoneNumber() bool {
	if o != nil && !IsNil(o.ChangePhoneNumber) {
		return true
	}

	return false
}

// SetChangePhoneNumber gets a reference to the given string and assigns it to the ChangePhoneNumber field.
func (o *ConnectionCreateOptionsCustomScripts) SetChangePhoneNumber(v string) {
	o.ChangePhoneNumber = &v
}

func (o ConnectionCreateOptionsCustomScripts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionCreateOptionsCustomScripts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Login) {
		toSerialize["login"] = o.Login
	}
	if !IsNil(o.GetUser) {
		toSerialize["get_user"] = o.GetUser
	}
	if !IsNil(o.Delete) {
		toSerialize["delete"] = o.Delete
	}
	if !IsNil(o.ChangePassword) {
		toSerialize["change_password"] = o.ChangePassword
	}
	if !IsNil(o.Verify) {
		toSerialize["verify"] = o.Verify
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.ChangeUsername) {
		toSerialize["change_username"] = o.ChangeUsername
	}
	if !IsNil(o.ChangeEmail) {
		toSerialize["change_email"] = o.ChangeEmail
	}
	if !IsNil(o.ChangePhoneNumber) {
		toSerialize["change_phone_number"] = o.ChangePhoneNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectionCreateOptionsCustomScripts) UnmarshalJSON(data []byte) (err error) {
	varConnectionCreateOptionsCustomScripts := _ConnectionCreateOptionsCustomScripts{}

	err = json.Unmarshal(data, &varConnectionCreateOptionsCustomScripts)

	if err != nil {
		return err
	}

	*o = ConnectionCreateOptionsCustomScripts(varConnectionCreateOptionsCustomScripts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "login")
		delete(additionalProperties, "get_user")
		delete(additionalProperties, "delete")
		delete(additionalProperties, "change_password")
		delete(additionalProperties, "verify")
		delete(additionalProperties, "create")
		delete(additionalProperties, "change_username")
		delete(additionalProperties, "change_email")
		delete(additionalProperties, "change_phone_number")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectionCreateOptionsCustomScripts struct {
	value *ConnectionCreateOptionsCustomScripts
	isSet bool
}

func (v NullableConnectionCreateOptionsCustomScripts) Get() *ConnectionCreateOptionsCustomScripts {
	return v.value
}

func (v *NullableConnectionCreateOptionsCustomScripts) Set(val *ConnectionCreateOptionsCustomScripts) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionCreateOptionsCustomScripts) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionCreateOptionsCustomScripts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionCreateOptionsCustomScripts(val *ConnectionCreateOptionsCustomScripts) *NullableConnectionCreateOptionsCustomScripts {
	return &NullableConnectionCreateOptionsCustomScripts{value: val, isSet: true}
}

func (v NullableConnectionCreateOptionsCustomScripts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionCreateOptionsCustomScripts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
