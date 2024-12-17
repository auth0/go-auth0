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

// UserCreate struct for UserCreate
type UserCreate struct {
	// The user's email.
	Email *string `json:"email,omitempty"`
	// The user's phone number (following the E.164 recommendation), only valid for users from SMS connections.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// Data related to the user that does not affect the application's core functionality.
	UserMetadata map[string]interface{} `json:"user_metadata,omitempty"`
	// Whether this user was blocked by an administrator (true) or not (false).
	Blocked *bool `json:"blocked,omitempty"`
	// Whether this email address is verified (true) or unverified (false). User will receive a verification email after creation if `email_verified` is false or not specified
	EmailVerified *bool `json:"email_verified,omitempty"`
	// Whether this phone number has been verified (true) or not (false).
	PhoneVerified *bool                                           `json:"phone_verified,omitempty"`
	AppMetadata   *GetInvitations200ResponseOneOfInnerAppMetadata `json:"app_metadata,omitempty"`
	// The user's given name(s).
	GivenName *string `json:"given_name,omitempty"`
	// The user's family name(s).
	FamilyName *string `json:"family_name,omitempty"`
	// The user's full name.
	Name *string `json:"name,omitempty"`
	// The user's nickname.
	Nickname *string `json:"nickname,omitempty"`
	// A URI pointing to the user's picture.
	Picture *string `json:"picture,omitempty"`
	// The external user's id provided by the identity provider.
	UserId *string `json:"user_id,omitempty"`
	// Name of the connection this user should be created in.
	Connection string `json:"connection"`
	// Initial password for this user (mandatory only for auth0 connection strategy).
	Password *string `json:"password,omitempty"`
	// Whether the user will receive a verification email after creation (true) or no email (false). Overrides behavior of `email_verified` parameter.
	VerifyEmail *bool `json:"verify_email,omitempty"`
	// The user's username. Only valid if the connection requires a username.
	Username *string `json:"username,omitempty"`
}

type _UserCreate UserCreate

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserCreate) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserCreate) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserCreate) SetEmail(v string) {
	o.Email = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *UserCreate) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *UserCreate) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *UserCreate) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetUserMetadata returns the UserMetadata field value if set, zero value otherwise.
func (o *UserCreate) GetUserMetadata() map[string]interface{} {
	if o == nil || IsNil(o.UserMetadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.UserMetadata
}

// GetUserMetadataOk returns a tuple with the UserMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetUserMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UserMetadata) {
		return map[string]interface{}{}, false
	}
	return o.UserMetadata, true
}

// HasUserMetadata returns a boolean if a field has been set.
func (o *UserCreate) HasUserMetadata() bool {
	if o != nil && !IsNil(o.UserMetadata) {
		return true
	}

	return false
}

// SetUserMetadata gets a reference to the given map[string]interface{} and assigns it to the UserMetadata field.
func (o *UserCreate) SetUserMetadata(v map[string]interface{}) {
	o.UserMetadata = v
}

// GetBlocked returns the Blocked field value if set, zero value otherwise.
func (o *UserCreate) GetBlocked() bool {
	if o == nil || IsNil(o.Blocked) {
		var ret bool
		return ret
	}
	return *o.Blocked
}

// GetBlockedOk returns a tuple with the Blocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetBlockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Blocked) {
		return nil, false
	}
	return o.Blocked, true
}

// HasBlocked returns a boolean if a field has been set.
func (o *UserCreate) HasBlocked() bool {
	if o != nil && !IsNil(o.Blocked) {
		return true
	}

	return false
}

// SetBlocked gets a reference to the given bool and assigns it to the Blocked field.
func (o *UserCreate) SetBlocked(v bool) {
	o.Blocked = &v
}

// GetEmailVerified returns the EmailVerified field value if set, zero value otherwise.
func (o *UserCreate) GetEmailVerified() bool {
	if o == nil || IsNil(o.EmailVerified) {
		var ret bool
		return ret
	}
	return *o.EmailVerified
}

// GetEmailVerifiedOk returns a tuple with the EmailVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetEmailVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailVerified) {
		return nil, false
	}
	return o.EmailVerified, true
}

// HasEmailVerified returns a boolean if a field has been set.
func (o *UserCreate) HasEmailVerified() bool {
	if o != nil && !IsNil(o.EmailVerified) {
		return true
	}

	return false
}

// SetEmailVerified gets a reference to the given bool and assigns it to the EmailVerified field.
func (o *UserCreate) SetEmailVerified(v bool) {
	o.EmailVerified = &v
}

// GetPhoneVerified returns the PhoneVerified field value if set, zero value otherwise.
func (o *UserCreate) GetPhoneVerified() bool {
	if o == nil || IsNil(o.PhoneVerified) {
		var ret bool
		return ret
	}
	return *o.PhoneVerified
}

// GetPhoneVerifiedOk returns a tuple with the PhoneVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetPhoneVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.PhoneVerified) {
		return nil, false
	}
	return o.PhoneVerified, true
}

// HasPhoneVerified returns a boolean if a field has been set.
func (o *UserCreate) HasPhoneVerified() bool {
	if o != nil && !IsNil(o.PhoneVerified) {
		return true
	}

	return false
}

// SetPhoneVerified gets a reference to the given bool and assigns it to the PhoneVerified field.
func (o *UserCreate) SetPhoneVerified(v bool) {
	o.PhoneVerified = &v
}

// GetAppMetadata returns the AppMetadata field value if set, zero value otherwise.
func (o *UserCreate) GetAppMetadata() GetInvitations200ResponseOneOfInnerAppMetadata {
	if o == nil || IsNil(o.AppMetadata) {
		var ret GetInvitations200ResponseOneOfInnerAppMetadata
		return ret
	}
	return *o.AppMetadata
}

// GetAppMetadataOk returns a tuple with the AppMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetAppMetadataOk() (*GetInvitations200ResponseOneOfInnerAppMetadata, bool) {
	if o == nil || IsNil(o.AppMetadata) {
		return nil, false
	}
	return o.AppMetadata, true
}

// HasAppMetadata returns a boolean if a field has been set.
func (o *UserCreate) HasAppMetadata() bool {
	if o != nil && !IsNil(o.AppMetadata) {
		return true
	}

	return false
}

// SetAppMetadata gets a reference to the given GetInvitations200ResponseOneOfInnerAppMetadata and assigns it to the AppMetadata field.
func (o *UserCreate) SetAppMetadata(v GetInvitations200ResponseOneOfInnerAppMetadata) {
	o.AppMetadata = &v
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *UserCreate) GetGivenName() string {
	if o == nil || IsNil(o.GivenName) {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetGivenNameOk() (*string, bool) {
	if o == nil || IsNil(o.GivenName) {
		return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *UserCreate) HasGivenName() bool {
	if o != nil && !IsNil(o.GivenName) {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *UserCreate) SetGivenName(v string) {
	o.GivenName = &v
}

// GetFamilyName returns the FamilyName field value if set, zero value otherwise.
func (o *UserCreate) GetFamilyName() string {
	if o == nil || IsNil(o.FamilyName) {
		var ret string
		return ret
	}
	return *o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetFamilyNameOk() (*string, bool) {
	if o == nil || IsNil(o.FamilyName) {
		return nil, false
	}
	return o.FamilyName, true
}

// HasFamilyName returns a boolean if a field has been set.
func (o *UserCreate) HasFamilyName() bool {
	if o != nil && !IsNil(o.FamilyName) {
		return true
	}

	return false
}

// SetFamilyName gets a reference to the given string and assigns it to the FamilyName field.
func (o *UserCreate) SetFamilyName(v string) {
	o.FamilyName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserCreate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserCreate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserCreate) SetName(v string) {
	o.Name = &v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *UserCreate) GetNickname() string {
	if o == nil || IsNil(o.Nickname) {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetNicknameOk() (*string, bool) {
	if o == nil || IsNil(o.Nickname) {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *UserCreate) HasNickname() bool {
	if o != nil && !IsNil(o.Nickname) {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *UserCreate) SetNickname(v string) {
	o.Nickname = &v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *UserCreate) GetPicture() string {
	if o == nil || IsNil(o.Picture) {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetPictureOk() (*string, bool) {
	if o == nil || IsNil(o.Picture) {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *UserCreate) HasPicture() bool {
	if o != nil && !IsNil(o.Picture) {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *UserCreate) SetPicture(v string) {
	o.Picture = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UserCreate) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UserCreate) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *UserCreate) SetUserId(v string) {
	o.UserId = &v
}

// GetConnection returns the Connection field value
func (o *UserCreate) GetConnection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value
// and a boolean to check if the value has been set.
func (o *UserCreate) GetConnectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connection, true
}

// SetConnection sets field value
func (o *UserCreate) SetConnection(v string) {
	o.Connection = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UserCreate) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UserCreate) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UserCreate) SetPassword(v string) {
	o.Password = &v
}

// GetVerifyEmail returns the VerifyEmail field value if set, zero value otherwise.
func (o *UserCreate) GetVerifyEmail() bool {
	if o == nil || IsNil(o.VerifyEmail) {
		var ret bool
		return ret
	}
	return *o.VerifyEmail
}

// GetVerifyEmailOk returns a tuple with the VerifyEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetVerifyEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyEmail) {
		return nil, false
	}
	return o.VerifyEmail, true
}

// HasVerifyEmail returns a boolean if a field has been set.
func (o *UserCreate) HasVerifyEmail() bool {
	if o != nil && !IsNil(o.VerifyEmail) {
		return true
	}

	return false
}

// SetVerifyEmail gets a reference to the given bool and assigns it to the VerifyEmail field.
func (o *UserCreate) SetVerifyEmail(v bool) {
	o.VerifyEmail = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UserCreate) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UserCreate) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UserCreate) SetUsername(v string) {
	o.Username = &v
}

func (o UserCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !IsNil(o.UserMetadata) {
		toSerialize["user_metadata"] = o.UserMetadata
	}
	if !IsNil(o.Blocked) {
		toSerialize["blocked"] = o.Blocked
	}
	if !IsNil(o.EmailVerified) {
		toSerialize["email_verified"] = o.EmailVerified
	}
	if !IsNil(o.PhoneVerified) {
		toSerialize["phone_verified"] = o.PhoneVerified
	}
	if !IsNil(o.AppMetadata) {
		toSerialize["app_metadata"] = o.AppMetadata
	}
	if !IsNil(o.GivenName) {
		toSerialize["given_name"] = o.GivenName
	}
	if !IsNil(o.FamilyName) {
		toSerialize["family_name"] = o.FamilyName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Nickname) {
		toSerialize["nickname"] = o.Nickname
	}
	if !IsNil(o.Picture) {
		toSerialize["picture"] = o.Picture
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	toSerialize["connection"] = o.Connection
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.VerifyEmail) {
		toSerialize["verify_email"] = o.VerifyEmail
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

func (o *UserCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connection",
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

	varUserCreate := _UserCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserCreate)

	if err != nil {
		return err
	}

	*o = UserCreate(varUserCreate)

	return err
}

type NullableUserCreate struct {
	value *UserCreate
	isSet bool
}

func (v NullableUserCreate) Get() *UserCreate {
	return v.value
}

func (v *NullableUserCreate) Set(val *UserCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCreate(val *UserCreate) *NullableUserCreate {
	return &NullableUserCreate{value: val, isSet: true}
}

func (v NullableUserCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
