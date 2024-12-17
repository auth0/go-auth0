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

// PostInvitationsRequest struct for PostInvitationsRequest
type PostInvitationsRequest struct {
	Inviter GetInvitations200ResponseOneOfInnerInviter `json:"inviter"`
	Invitee GetInvitations200ResponseOneOfInnerInvitee `json:"invitee"`
	// Auth0 client ID. Used to resolve the application's login initiation endpoint.
	ClientId string `json:"client_id"`
	// The id of the connection to force invitee to authenticate with.
	ConnectionId *string                                         `json:"connection_id,omitempty"`
	AppMetadata  *GetInvitations200ResponseOneOfInnerAppMetadata `json:"app_metadata,omitempty"`
	// Data related to the user that does not affect the application's core functionality.
	UserMetadata map[string]interface{} `json:"user_metadata,omitempty"`
	// Number of seconds for which the invitation is valid before expiration. If unspecified or set to 0, this value defaults to 604800 seconds (7 days). Max value: 2592000 seconds (30 days).
	TtlSec *int32 `json:"ttl_sec,omitempty"`
	// List of roles IDs to associated with the user.
	Roles []string `json:"roles,omitempty"`
	// Whether the user will receive an invitation email (true) or no email (false), true by default
	SendInvitationEmail *bool `json:"send_invitation_email,omitempty"`
}

type _PostInvitationsRequest PostInvitationsRequest

// GetInviter returns the Inviter field value
func (o *PostInvitationsRequest) GetInviter() GetInvitations200ResponseOneOfInnerInviter {
	if o == nil {
		var ret GetInvitations200ResponseOneOfInnerInviter
		return ret
	}

	return o.Inviter
}

// GetInviterOk returns a tuple with the Inviter field value
// and a boolean to check if the value has been set.
func (o *PostInvitationsRequest) GetInviterOk() (*GetInvitations200ResponseOneOfInnerInviter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inviter, true
}

// SetInviter sets field value
func (o *PostInvitationsRequest) SetInviter(v GetInvitations200ResponseOneOfInnerInviter) {
	o.Inviter = v
}

// GetInvitee returns the Invitee field value
func (o *PostInvitationsRequest) GetInvitee() GetInvitations200ResponseOneOfInnerInvitee {
	if o == nil {
		var ret GetInvitations200ResponseOneOfInnerInvitee
		return ret
	}

	return o.Invitee
}

// GetInviteeOk returns a tuple with the Invitee field value
// and a boolean to check if the value has been set.
func (o *PostInvitationsRequest) GetInviteeOk() (*GetInvitations200ResponseOneOfInnerInvitee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Invitee, true
}

// SetInvitee sets field value
func (o *PostInvitationsRequest) SetInvitee(v GetInvitations200ResponseOneOfInnerInvitee) {
	o.Invitee = v
}

// GetClientId returns the ClientId field value
func (o *PostInvitationsRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *PostInvitationsRequest) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *PostInvitationsRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *PostInvitationsRequest) GetConnectionId() string {
	if o == nil || IsNil(o.ConnectionId) {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInvitationsRequest) GetConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionId) {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *PostInvitationsRequest) HasConnectionId() bool {
	if o != nil && !IsNil(o.ConnectionId) {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *PostInvitationsRequest) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetAppMetadata returns the AppMetadata field value if set, zero value otherwise.
func (o *PostInvitationsRequest) GetAppMetadata() GetInvitations200ResponseOneOfInnerAppMetadata {
	if o == nil || IsNil(o.AppMetadata) {
		var ret GetInvitations200ResponseOneOfInnerAppMetadata
		return ret
	}
	return *o.AppMetadata
}

// GetAppMetadataOk returns a tuple with the AppMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInvitationsRequest) GetAppMetadataOk() (*GetInvitations200ResponseOneOfInnerAppMetadata, bool) {
	if o == nil || IsNil(o.AppMetadata) {
		return nil, false
	}
	return o.AppMetadata, true
}

// HasAppMetadata returns a boolean if a field has been set.
func (o *PostInvitationsRequest) HasAppMetadata() bool {
	if o != nil && !IsNil(o.AppMetadata) {
		return true
	}

	return false
}

// SetAppMetadata gets a reference to the given GetInvitations200ResponseOneOfInnerAppMetadata and assigns it to the AppMetadata field.
func (o *PostInvitationsRequest) SetAppMetadata(v GetInvitations200ResponseOneOfInnerAppMetadata) {
	o.AppMetadata = &v
}

// GetUserMetadata returns the UserMetadata field value if set, zero value otherwise.
func (o *PostInvitationsRequest) GetUserMetadata() map[string]interface{} {
	if o == nil || IsNil(o.UserMetadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.UserMetadata
}

// GetUserMetadataOk returns a tuple with the UserMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInvitationsRequest) GetUserMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UserMetadata) {
		return map[string]interface{}{}, false
	}
	return o.UserMetadata, true
}

// HasUserMetadata returns a boolean if a field has been set.
func (o *PostInvitationsRequest) HasUserMetadata() bool {
	if o != nil && !IsNil(o.UserMetadata) {
		return true
	}

	return false
}

// SetUserMetadata gets a reference to the given map[string]interface{} and assigns it to the UserMetadata field.
func (o *PostInvitationsRequest) SetUserMetadata(v map[string]interface{}) {
	o.UserMetadata = v
}

// GetTtlSec returns the TtlSec field value if set, zero value otherwise.
func (o *PostInvitationsRequest) GetTtlSec() int32 {
	if o == nil || IsNil(o.TtlSec) {
		var ret int32
		return ret
	}
	return *o.TtlSec
}

// GetTtlSecOk returns a tuple with the TtlSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInvitationsRequest) GetTtlSecOk() (*int32, bool) {
	if o == nil || IsNil(o.TtlSec) {
		return nil, false
	}
	return o.TtlSec, true
}

// HasTtlSec returns a boolean if a field has been set.
func (o *PostInvitationsRequest) HasTtlSec() bool {
	if o != nil && !IsNil(o.TtlSec) {
		return true
	}

	return false
}

// SetTtlSec gets a reference to the given int32 and assigns it to the TtlSec field.
func (o *PostInvitationsRequest) SetTtlSec(v int32) {
	o.TtlSec = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *PostInvitationsRequest) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInvitationsRequest) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *PostInvitationsRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *PostInvitationsRequest) SetRoles(v []string) {
	o.Roles = v
}

// GetSendInvitationEmail returns the SendInvitationEmail field value if set, zero value otherwise.
func (o *PostInvitationsRequest) GetSendInvitationEmail() bool {
	if o == nil || IsNil(o.SendInvitationEmail) {
		var ret bool
		return ret
	}
	return *o.SendInvitationEmail
}

// GetSendInvitationEmailOk returns a tuple with the SendInvitationEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInvitationsRequest) GetSendInvitationEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.SendInvitationEmail) {
		return nil, false
	}
	return o.SendInvitationEmail, true
}

// HasSendInvitationEmail returns a boolean if a field has been set.
func (o *PostInvitationsRequest) HasSendInvitationEmail() bool {
	if o != nil && !IsNil(o.SendInvitationEmail) {
		return true
	}

	return false
}

// SetSendInvitationEmail gets a reference to the given bool and assigns it to the SendInvitationEmail field.
func (o *PostInvitationsRequest) SetSendInvitationEmail(v bool) {
	o.SendInvitationEmail = &v
}

func (o PostInvitationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostInvitationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inviter"] = o.Inviter
	toSerialize["invitee"] = o.Invitee
	toSerialize["client_id"] = o.ClientId
	if !IsNil(o.ConnectionId) {
		toSerialize["connection_id"] = o.ConnectionId
	}
	if !IsNil(o.AppMetadata) {
		toSerialize["app_metadata"] = o.AppMetadata
	}
	if !IsNil(o.UserMetadata) {
		toSerialize["user_metadata"] = o.UserMetadata
	}
	if !IsNil(o.TtlSec) {
		toSerialize["ttl_sec"] = o.TtlSec
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.SendInvitationEmail) {
		toSerialize["send_invitation_email"] = o.SendInvitationEmail
	}
	return toSerialize, nil
}

func (o *PostInvitationsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"inviter",
		"invitee",
		"client_id",
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

	varPostInvitationsRequest := _PostInvitationsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostInvitationsRequest)

	if err != nil {
		return err
	}

	*o = PostInvitationsRequest(varPostInvitationsRequest)

	return err
}

type NullablePostInvitationsRequest struct {
	value *PostInvitationsRequest
	isSet bool
}

func (v NullablePostInvitationsRequest) Get() *PostInvitationsRequest {
	return v.value
}

func (v *NullablePostInvitationsRequest) Set(val *PostInvitationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostInvitationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostInvitationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostInvitationsRequest(val *PostInvitationsRequest) *NullablePostInvitationsRequest {
	return &NullablePostInvitationsRequest{value: val, isSet: true}
}

func (v NullablePostInvitationsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostInvitationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
