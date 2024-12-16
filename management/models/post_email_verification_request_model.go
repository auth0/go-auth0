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

// PostEmailVerificationRequest struct for PostEmailVerificationRequest
type PostEmailVerificationRequest struct {
	// URL the user will be redirected to in the classic Universal Login experience once the ticket is used.
	ResultUrl *string `json:"result_url,omitempty"`
	// user_id of for whom the ticket should be created.
	UserId string `json:"user_id"`
	// ID of the client. If provided for tenants using New Universal Login experience, the user will be prompted to redirect to the default login route of the corresponding application once the ticket is used. See <a target='' href='https://manage.local.dev.auth0.com/docs/universal-login/configure-default-login-routes#completing-the-password-reset-flow'>Configuring Default Login Routes</a> for more details.
	ClientId *string `json:"client_id,omitempty"`
	// (Optional) Organization ID – the ID of the Organization. If provided, organization parameters will be made available to the email template and organization branding will be applied to the prompt. In addition, the redirect link in the prompt will include organization_id and organization_name query string parameters.
	OrganizationId *string `json:"organization_id,omitempty"`
	// Number of seconds for which the ticket is valid before expiration. If unspecified or set to 0, this value defaults to 432000 seconds (5 days).
	TtlSec *int32 `json:"ttl_sec,omitempty"`
	// Whether to include the email address as part of the returnUrl in the reset_email (true), or not (false).
	IncludeEmailInRedirect *bool                                 `json:"includeEmailInRedirect,omitempty"`
	Identity               *PostVerificationEmailRequestIdentity `json:"identity,omitempty"`
}

type _PostEmailVerificationRequest PostEmailVerificationRequest

// GetResultUrl returns the ResultUrl field value if set, zero value otherwise.
func (o *PostEmailVerificationRequest) GetResultUrl() string {
	if o == nil || IsNil(o.ResultUrl) {
		var ret string
		return ret
	}
	return *o.ResultUrl
}

// GetResultUrlOk returns a tuple with the ResultUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostEmailVerificationRequest) GetResultUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ResultUrl) {
		return nil, false
	}
	return o.ResultUrl, true
}

// HasResultUrl returns a boolean if a field has been set.
func (o *PostEmailVerificationRequest) HasResultUrl() bool {
	if o != nil && !IsNil(o.ResultUrl) {
		return true
	}

	return false
}

// SetResultUrl gets a reference to the given string and assigns it to the ResultUrl field.
func (o *PostEmailVerificationRequest) SetResultUrl(v string) {
	o.ResultUrl = &v
}

// GetUserId returns the UserId field value
func (o *PostEmailVerificationRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *PostEmailVerificationRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *PostEmailVerificationRequest) SetUserId(v string) {
	o.UserId = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PostEmailVerificationRequest) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostEmailVerificationRequest) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PostEmailVerificationRequest) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PostEmailVerificationRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *PostEmailVerificationRequest) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostEmailVerificationRequest) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *PostEmailVerificationRequest) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *PostEmailVerificationRequest) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetTtlSec returns the TtlSec field value if set, zero value otherwise.
func (o *PostEmailVerificationRequest) GetTtlSec() int32 {
	if o == nil || IsNil(o.TtlSec) {
		var ret int32
		return ret
	}
	return *o.TtlSec
}

// GetTtlSecOk returns a tuple with the TtlSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostEmailVerificationRequest) GetTtlSecOk() (*int32, bool) {
	if o == nil || IsNil(o.TtlSec) {
		return nil, false
	}
	return o.TtlSec, true
}

// HasTtlSec returns a boolean if a field has been set.
func (o *PostEmailVerificationRequest) HasTtlSec() bool {
	if o != nil && !IsNil(o.TtlSec) {
		return true
	}

	return false
}

// SetTtlSec gets a reference to the given int32 and assigns it to the TtlSec field.
func (o *PostEmailVerificationRequest) SetTtlSec(v int32) {
	o.TtlSec = &v
}

// GetIncludeEmailInRedirect returns the IncludeEmailInRedirect field value if set, zero value otherwise.
func (o *PostEmailVerificationRequest) GetIncludeEmailInRedirect() bool {
	if o == nil || IsNil(o.IncludeEmailInRedirect) {
		var ret bool
		return ret
	}
	return *o.IncludeEmailInRedirect
}

// GetIncludeEmailInRedirectOk returns a tuple with the IncludeEmailInRedirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostEmailVerificationRequest) GetIncludeEmailInRedirectOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeEmailInRedirect) {
		return nil, false
	}
	return o.IncludeEmailInRedirect, true
}

// HasIncludeEmailInRedirect returns a boolean if a field has been set.
func (o *PostEmailVerificationRequest) HasIncludeEmailInRedirect() bool {
	if o != nil && !IsNil(o.IncludeEmailInRedirect) {
		return true
	}

	return false
}

// SetIncludeEmailInRedirect gets a reference to the given bool and assigns it to the IncludeEmailInRedirect field.
func (o *PostEmailVerificationRequest) SetIncludeEmailInRedirect(v bool) {
	o.IncludeEmailInRedirect = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *PostEmailVerificationRequest) GetIdentity() PostVerificationEmailRequestIdentity {
	if o == nil || IsNil(o.Identity) {
		var ret PostVerificationEmailRequestIdentity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostEmailVerificationRequest) GetIdentityOk() (*PostVerificationEmailRequestIdentity, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *PostEmailVerificationRequest) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given PostVerificationEmailRequestIdentity and assigns it to the Identity field.
func (o *PostEmailVerificationRequest) SetIdentity(v PostVerificationEmailRequestIdentity) {
	o.Identity = &v
}

func (o PostEmailVerificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostEmailVerificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResultUrl) {
		toSerialize["result_url"] = o.ResultUrl
	}
	toSerialize["user_id"] = o.UserId
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !IsNil(o.TtlSec) {
		toSerialize["ttl_sec"] = o.TtlSec
	}
	if !IsNil(o.IncludeEmailInRedirect) {
		toSerialize["includeEmailInRedirect"] = o.IncludeEmailInRedirect
	}
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	return toSerialize, nil
}

func (o *PostEmailVerificationRequest) UnmarshalJSON(data []byte) (err error) {
	varPostEmailVerificationRequest := _PostEmailVerificationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostEmailVerificationRequest)

	if err != nil {
		return err
	}

	*o = PostEmailVerificationRequest(varPostEmailVerificationRequest)

	return err
}

type NullablePostEmailVerificationRequest struct {
	value *PostEmailVerificationRequest
	isSet bool
}

func (v NullablePostEmailVerificationRequest) Get() *PostEmailVerificationRequest {
	return v.value
}

func (v *NullablePostEmailVerificationRequest) Set(val *PostEmailVerificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostEmailVerificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostEmailVerificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostEmailVerificationRequest(val *PostEmailVerificationRequest) *NullablePostEmailVerificationRequest {
	return &NullablePostEmailVerificationRequest{value: val, isSet: true}
}

func (v NullablePostEmailVerificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostEmailVerificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
