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

// GetUsers200ResponseOneOfInnerAppMetadata User metadata to which this user has read-only access.
type GetUsers200ResponseOneOfInnerAppMetadata struct {
	ClientID             interface{} `json:"clientID,omitempty"`
	GlobalClientID       interface{} `json:"globalClientID,omitempty"`
	GlobalClientId       interface{} `json:"global_client_id,omitempty"`
	EmailVerified        interface{} `json:"email_verified,omitempty"`
	UserId               interface{} `json:"user_id,omitempty"`
	Identities           interface{} `json:"identities,omitempty"`
	LastIP               interface{} `json:"lastIP,omitempty"`
	LastLogin            interface{} `json:"lastLogin,omitempty"`
	Metadata             interface{} `json:"metadata,omitempty"`
	CreatedAt            interface{} `json:"created_at,omitempty"`
	LoginsCount          interface{} `json:"loginsCount,omitempty"`
	Id                   interface{} `json:"_id,omitempty"`
	Email                interface{} `json:"email,omitempty"`
	Blocked              interface{} `json:"blocked,omitempty"`
	Tenant               interface{} `json:"__tenant,omitempty"`
	UpdatedAt            interface{} `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUsers200ResponseOneOfInnerAppMetadata GetUsers200ResponseOneOfInnerAppMetadata

// GetClientID returns the ClientID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetClientID() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ClientID
}

// GetClientIDOk returns a tuple with the ClientID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetClientIDOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ClientID) {
		return nil, false
	}
	return &o.ClientID, true
}

// HasClientID returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) HasClientID() bool {
	if o != nil && !IsNil(o.ClientID) {
		return true
	}

	return false
}

// SetClientID gets a reference to the given interface{} and assigns it to the ClientID field.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) SetClientID(v interface{}) {
	o.ClientID = v
}

// GetGlobalClientID returns the GlobalClientID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetGlobalClientID() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.GlobalClientID
}

// GetGlobalClientIDOk returns a tuple with the GlobalClientID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetGlobalClientIDOk() (*interface{}, bool) {
	if o == nil || IsNil(o.GlobalClientID) {
		return nil, false
	}
	return &o.GlobalClientID, true
}

// HasGlobalClientID returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) HasGlobalClientID() bool {
	if o != nil && !IsNil(o.GlobalClientID) {
		return true
	}

	return false
}

// SetGlobalClientID gets a reference to the given interface{} and assigns it to the GlobalClientID field.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) SetGlobalClientID(v interface{}) {
	o.GlobalClientID = v
}

// GetGlobalClientId returns the GlobalClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetGlobalClientId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.GlobalClientId
}

// GetGlobalClientIdOk returns a tuple with the GlobalClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetGlobalClientIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.GlobalClientId) {
		return nil, false
	}
	return &o.GlobalClientId, true
}

// HasGlobalClientId returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) HasGlobalClientId() bool {
	if o != nil && !IsNil(o.GlobalClientId) {
		return true
	}

	return false
}

// SetGlobalClientId gets a reference to the given interface{} and assigns it to the GlobalClientId field.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) SetGlobalClientId(v interface{}) {
	o.GlobalClientId = v
}

// GetEmailVerified returns the EmailVerified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetEmailVerified() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EmailVerified
}

// GetEmailVerifiedOk returns a tuple with the EmailVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetEmailVerifiedOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EmailVerified) {
		return nil, false
	}
	return &o.EmailVerified, true
}

// HasEmailVerified returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) HasEmailVerified() bool {
	if o != nil && !IsNil(o.EmailVerified) {
		return true
	}

	return false
}

// SetEmailVerified gets a reference to the given interface{} and assigns it to the EmailVerified field.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) SetEmailVerified(v interface{}) {
	o.EmailVerified = v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetUserId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetUserIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return &o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given interface{} and assigns it to the UserId field.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) SetUserId(v interface{}) {
	o.UserId = v
}

// GetIdentities returns the Identities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetIdentities() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Identities
}

// GetIdentitiesOk returns a tuple with the Identities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetIdentitiesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Identities) {
		return nil, false
	}
	return &o.Identities, true
}

// HasIdentities returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) HasIdentities() bool {
	if o != nil && !IsNil(o.Identities) {
		return true
	}

	return false
}

// SetIdentities gets a reference to the given interface{} and assigns it to the Identities field.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) SetIdentities(v interface{}) {
	o.Identities = v
}

// GetLastIP returns the LastIP field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetLastIP() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.LastIP
}

// GetLastIPOk returns a tuple with the LastIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetLastIPOk() (*interface{}, bool) {
	if o == nil || IsNil(o.LastIP) {
		return nil, false
	}
	return &o.LastIP, true
}

// HasLastIP returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) HasLastIP() bool {
	if o != nil && !IsNil(o.LastIP) {
		return true
	}

	return false
}

// SetLastIP gets a reference to the given interface{} and assigns it to the LastIP field.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) SetLastIP(v interface{}) {
	o.LastIP = v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetLastLogin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetLastLoginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.LastLogin) {
		return nil, false
	}
	return &o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) HasLastLogin() bool {
	if o != nil && !IsNil(o.LastLogin) {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given interface{} and assigns it to the LastLogin field.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) SetLastLogin(v interface{}) {
	o.LastLogin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetLoginsCount returns the LoginsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetLoginsCount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.LoginsCount
}

// GetLoginsCountOk returns a tuple with the LoginsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetLoginsCountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.LoginsCount) {
		return nil, false
	}
	return &o.LoginsCount, true
}

// HasLoginsCount returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) HasLoginsCount() bool {
	if o != nil && !IsNil(o.LoginsCount) {
		return true
	}

	return false
}

// SetLoginsCount gets a reference to the given interface{} and assigns it to the LoginsCount field.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) SetLoginsCount(v interface{}) {
	o.LoginsCount = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) SetId(v interface{}) {
	o.Id = v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetEmail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetEmailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return &o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given interface{} and assigns it to the Email field.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) SetEmail(v interface{}) {
	o.Email = v
}

// GetBlocked returns the Blocked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetBlocked() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Blocked
}

// GetBlockedOk returns a tuple with the Blocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetBlockedOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Blocked) {
		return nil, false
	}
	return &o.Blocked, true
}

// HasBlocked returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) HasBlocked() bool {
	if o != nil && !IsNil(o.Blocked) {
		return true
	}

	return false
}

// SetBlocked gets a reference to the given interface{} and assigns it to the Blocked field.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) SetBlocked(v interface{}) {
	o.Blocked = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetTenant() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetTenantOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return &o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given interface{} and assigns it to the Tenant field.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) SetTenant(v interface{}) {
	o.Tenant = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetUsers200ResponseOneOfInnerAppMetadata) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GetUsers200ResponseOneOfInnerAppMetadata) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

func (o GetUsers200ResponseOneOfInnerAppMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUsers200ResponseOneOfInnerAppMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientID != nil {
		toSerialize["clientID"] = o.ClientID
	}
	if o.GlobalClientID != nil {
		toSerialize["globalClientID"] = o.GlobalClientID
	}
	if o.GlobalClientId != nil {
		toSerialize["global_client_id"] = o.GlobalClientId
	}
	if o.EmailVerified != nil {
		toSerialize["email_verified"] = o.EmailVerified
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	if o.Identities != nil {
		toSerialize["identities"] = o.Identities
	}
	if o.LastIP != nil {
		toSerialize["lastIP"] = o.LastIP
	}
	if o.LastLogin != nil {
		toSerialize["lastLogin"] = o.LastLogin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LoginsCount != nil {
		toSerialize["loginsCount"] = o.LoginsCount
	}
	if o.Id != nil {
		toSerialize["_id"] = o.Id
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Blocked != nil {
		toSerialize["blocked"] = o.Blocked
	}
	if o.Tenant != nil {
		toSerialize["__tenant"] = o.Tenant
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetUsers200ResponseOneOfInnerAppMetadata) UnmarshalJSON(data []byte) (err error) {
	varGetUsers200ResponseOneOfInnerAppMetadata := _GetUsers200ResponseOneOfInnerAppMetadata{}

	err = json.Unmarshal(data, &varGetUsers200ResponseOneOfInnerAppMetadata)

	if err != nil {
		return err
	}

	*o = GetUsers200ResponseOneOfInnerAppMetadata(varGetUsers200ResponseOneOfInnerAppMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "clientID")
		delete(additionalProperties, "globalClientID")
		delete(additionalProperties, "global_client_id")
		delete(additionalProperties, "email_verified")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "identities")
		delete(additionalProperties, "lastIP")
		delete(additionalProperties, "lastLogin")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "loginsCount")
		delete(additionalProperties, "_id")
		delete(additionalProperties, "email")
		delete(additionalProperties, "blocked")
		delete(additionalProperties, "__tenant")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUsers200ResponseOneOfInnerAppMetadata struct {
	value *GetUsers200ResponseOneOfInnerAppMetadata
	isSet bool
}

func (v NullableGetUsers200ResponseOneOfInnerAppMetadata) Get() *GetUsers200ResponseOneOfInnerAppMetadata {
	return v.value
}

func (v *NullableGetUsers200ResponseOneOfInnerAppMetadata) Set(val *GetUsers200ResponseOneOfInnerAppMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUsers200ResponseOneOfInnerAppMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUsers200ResponseOneOfInnerAppMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUsers200ResponseOneOfInnerAppMetadata(val *GetUsers200ResponseOneOfInnerAppMetadata) *NullableGetUsers200ResponseOneOfInnerAppMetadata {
	return &NullableGetUsers200ResponseOneOfInnerAppMetadata{value: val, isSet: true}
}

func (v NullableGetUsers200ResponseOneOfInnerAppMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUsers200ResponseOneOfInnerAppMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
