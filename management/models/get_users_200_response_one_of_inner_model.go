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

// GetUsers200ResponseOneOfInner struct for GetUsers200ResponseOneOfInner
type GetUsers200ResponseOneOfInner struct {
	// ID of the user which can be used when interacting with other APIs.
	UserId *string `json:"user_id,omitempty"`
	// Email address of this user.
	Email *string `json:"email,omitempty"`
	// Whether this email address is verified (true) or unverified (false).
	EmailVerified *bool `json:"email_verified,omitempty"`
	// Username of this user.
	Username *string `json:"username,omitempty"`
	// Phone number for this user. Follows the <a href=\"https://en.wikipedia.org/wiki/E.164\">E.164 recommendation</a>.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// Whether this phone number has been verified (true) or not (false).
	PhoneVerified *bool                                   `json:"phone_verified,omitempty"`
	CreatedAt     *GetUsers200ResponseOneOfInnerCreatedAt `json:"created_at,omitempty"`
	UpdatedAt     *GetUsers200ResponseOneOfInnerUpdatedAt `json:"updated_at,omitempty"`
	// Array of user identity objects when accounts are linked.
	Identities  []GetUsers200ResponseOneOfInnerIdentitiesInner `json:"identities,omitempty"`
	AppMetadata *GetUsers200ResponseOneOfInnerAppMetadata      `json:"app_metadata,omitempty"`
	// User metadata to which this user has read/write access.
	UserMetadata map[string]interface{} `json:"user_metadata,omitempty"`
	// URL to picture, photo, or avatar of this user.
	Picture *string `json:"picture,omitempty"`
	// Name of this user.
	Name *string `json:"name,omitempty"`
	// Preferred nickname or alias of this user.
	Nickname *string `json:"nickname,omitempty"`
	// List of multi-factor authentication providers with which this user has enrolled.
	Multifactor []string `json:"multifactor,omitempty"`
	// Last IP address from which this user logged in.
	LastIp    *string                                 `json:"last_ip,omitempty"`
	LastLogin *GetUsers200ResponseOneOfInnerLastLogin `json:"last_login,omitempty"`
	// Total number of logins this user has performed.
	LoginsCount *int32 `json:"logins_count,omitempty"`
	// Whether this user was blocked by an administrator (true) or is not (false).
	Blocked *bool `json:"blocked,omitempty"`
	// Given name/first name/forename of this user.
	GivenName *string `json:"given_name,omitempty"`
	// Family name/last name/surname of this user.
	FamilyName           *string `json:"family_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUsers200ResponseOneOfInner GetUsers200ResponseOneOfInner

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *GetUsers200ResponseOneOfInner) SetUserId(v string) {
	o.UserId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetUsers200ResponseOneOfInner) SetEmail(v string) {
	o.Email = &v
}

// GetEmailVerified returns the EmailVerified field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetEmailVerified() bool {
	if o == nil || IsNil(o.EmailVerified) {
		var ret bool
		return ret
	}
	return *o.EmailVerified
}

// GetEmailVerifiedOk returns a tuple with the EmailVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetEmailVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailVerified) {
		return nil, false
	}
	return o.EmailVerified, true
}

// HasEmailVerified returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasEmailVerified() bool {
	if o != nil && !IsNil(o.EmailVerified) {
		return true
	}

	return false
}

// SetEmailVerified gets a reference to the given bool and assigns it to the EmailVerified field.
func (o *GetUsers200ResponseOneOfInner) SetEmailVerified(v bool) {
	o.EmailVerified = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GetUsers200ResponseOneOfInner) SetUsername(v string) {
	o.Username = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *GetUsers200ResponseOneOfInner) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetPhoneVerified returns the PhoneVerified field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetPhoneVerified() bool {
	if o == nil || IsNil(o.PhoneVerified) {
		var ret bool
		return ret
	}
	return *o.PhoneVerified
}

// GetPhoneVerifiedOk returns a tuple with the PhoneVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetPhoneVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.PhoneVerified) {
		return nil, false
	}
	return o.PhoneVerified, true
}

// HasPhoneVerified returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasPhoneVerified() bool {
	if o != nil && !IsNil(o.PhoneVerified) {
		return true
	}

	return false
}

// SetPhoneVerified gets a reference to the given bool and assigns it to the PhoneVerified field.
func (o *GetUsers200ResponseOneOfInner) SetPhoneVerified(v bool) {
	o.PhoneVerified = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetCreatedAt() GetUsers200ResponseOneOfInnerCreatedAt {
	if o == nil || IsNil(o.CreatedAt) {
		var ret GetUsers200ResponseOneOfInnerCreatedAt
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetCreatedAtOk() (*GetUsers200ResponseOneOfInnerCreatedAt, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given GetUsers200ResponseOneOfInnerCreatedAt and assigns it to the CreatedAt field.
func (o *GetUsers200ResponseOneOfInner) SetCreatedAt(v GetUsers200ResponseOneOfInnerCreatedAt) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetUpdatedAt() GetUsers200ResponseOneOfInnerUpdatedAt {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret GetUsers200ResponseOneOfInnerUpdatedAt
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetUpdatedAtOk() (*GetUsers200ResponseOneOfInnerUpdatedAt, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given GetUsers200ResponseOneOfInnerUpdatedAt and assigns it to the UpdatedAt field.
func (o *GetUsers200ResponseOneOfInner) SetUpdatedAt(v GetUsers200ResponseOneOfInnerUpdatedAt) {
	o.UpdatedAt = &v
}

// GetIdentities returns the Identities field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetIdentities() []GetUsers200ResponseOneOfInnerIdentitiesInner {
	if o == nil || IsNil(o.Identities) {
		var ret []GetUsers200ResponseOneOfInnerIdentitiesInner
		return ret
	}
	return o.Identities
}

// GetIdentitiesOk returns a tuple with the Identities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetIdentitiesOk() ([]GetUsers200ResponseOneOfInnerIdentitiesInner, bool) {
	if o == nil || IsNil(o.Identities) {
		return nil, false
	}
	return o.Identities, true
}

// HasIdentities returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasIdentities() bool {
	if o != nil && !IsNil(o.Identities) {
		return true
	}

	return false
}

// SetIdentities gets a reference to the given []GetUsers200ResponseOneOfInnerIdentitiesInner and assigns it to the Identities field.
func (o *GetUsers200ResponseOneOfInner) SetIdentities(v []GetUsers200ResponseOneOfInnerIdentitiesInner) {
	o.Identities = v
}

// GetAppMetadata returns the AppMetadata field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetAppMetadata() GetUsers200ResponseOneOfInnerAppMetadata {
	if o == nil || IsNil(o.AppMetadata) {
		var ret GetUsers200ResponseOneOfInnerAppMetadata
		return ret
	}
	return *o.AppMetadata
}

// GetAppMetadataOk returns a tuple with the AppMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetAppMetadataOk() (*GetUsers200ResponseOneOfInnerAppMetadata, bool) {
	if o == nil || IsNil(o.AppMetadata) {
		return nil, false
	}
	return o.AppMetadata, true
}

// HasAppMetadata returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasAppMetadata() bool {
	if o != nil && !IsNil(o.AppMetadata) {
		return true
	}

	return false
}

// SetAppMetadata gets a reference to the given GetUsers200ResponseOneOfInnerAppMetadata and assigns it to the AppMetadata field.
func (o *GetUsers200ResponseOneOfInner) SetAppMetadata(v GetUsers200ResponseOneOfInnerAppMetadata) {
	o.AppMetadata = &v
}

// GetUserMetadata returns the UserMetadata field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetUserMetadata() map[string]interface{} {
	if o == nil || IsNil(o.UserMetadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.UserMetadata
}

// GetUserMetadataOk returns a tuple with the UserMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetUserMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UserMetadata) {
		return map[string]interface{}{}, false
	}
	return o.UserMetadata, true
}

// HasUserMetadata returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasUserMetadata() bool {
	if o != nil && !IsNil(o.UserMetadata) {
		return true
	}

	return false
}

// SetUserMetadata gets a reference to the given map[string]interface{} and assigns it to the UserMetadata field.
func (o *GetUsers200ResponseOneOfInner) SetUserMetadata(v map[string]interface{}) {
	o.UserMetadata = v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetPicture() string {
	if o == nil || IsNil(o.Picture) {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetPictureOk() (*string, bool) {
	if o == nil || IsNil(o.Picture) {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasPicture() bool {
	if o != nil && !IsNil(o.Picture) {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *GetUsers200ResponseOneOfInner) SetPicture(v string) {
	o.Picture = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetUsers200ResponseOneOfInner) SetName(v string) {
	o.Name = &v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetNickname() string {
	if o == nil || IsNil(o.Nickname) {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetNicknameOk() (*string, bool) {
	if o == nil || IsNil(o.Nickname) {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasNickname() bool {
	if o != nil && !IsNil(o.Nickname) {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *GetUsers200ResponseOneOfInner) SetNickname(v string) {
	o.Nickname = &v
}

// GetMultifactor returns the Multifactor field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetMultifactor() []string {
	if o == nil || IsNil(o.Multifactor) {
		var ret []string
		return ret
	}
	return o.Multifactor
}

// GetMultifactorOk returns a tuple with the Multifactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetMultifactorOk() ([]string, bool) {
	if o == nil || IsNil(o.Multifactor) {
		return nil, false
	}
	return o.Multifactor, true
}

// HasMultifactor returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasMultifactor() bool {
	if o != nil && !IsNil(o.Multifactor) {
		return true
	}

	return false
}

// SetMultifactor gets a reference to the given []string and assigns it to the Multifactor field.
func (o *GetUsers200ResponseOneOfInner) SetMultifactor(v []string) {
	o.Multifactor = v
}

// GetLastIp returns the LastIp field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetLastIp() string {
	if o == nil || IsNil(o.LastIp) {
		var ret string
		return ret
	}
	return *o.LastIp
}

// GetLastIpOk returns a tuple with the LastIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetLastIpOk() (*string, bool) {
	if o == nil || IsNil(o.LastIp) {
		return nil, false
	}
	return o.LastIp, true
}

// HasLastIp returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasLastIp() bool {
	if o != nil && !IsNil(o.LastIp) {
		return true
	}

	return false
}

// SetLastIp gets a reference to the given string and assigns it to the LastIp field.
func (o *GetUsers200ResponseOneOfInner) SetLastIp(v string) {
	o.LastIp = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetLastLogin() GetUsers200ResponseOneOfInnerLastLogin {
	if o == nil || IsNil(o.LastLogin) {
		var ret GetUsers200ResponseOneOfInnerLastLogin
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetLastLoginOk() (*GetUsers200ResponseOneOfInnerLastLogin, bool) {
	if o == nil || IsNil(o.LastLogin) {
		return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasLastLogin() bool {
	if o != nil && !IsNil(o.LastLogin) {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given GetUsers200ResponseOneOfInnerLastLogin and assigns it to the LastLogin field.
func (o *GetUsers200ResponseOneOfInner) SetLastLogin(v GetUsers200ResponseOneOfInnerLastLogin) {
	o.LastLogin = &v
}

// GetLoginsCount returns the LoginsCount field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetLoginsCount() int32 {
	if o == nil || IsNil(o.LoginsCount) {
		var ret int32
		return ret
	}
	return *o.LoginsCount
}

// GetLoginsCountOk returns a tuple with the LoginsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetLoginsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.LoginsCount) {
		return nil, false
	}
	return o.LoginsCount, true
}

// HasLoginsCount returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasLoginsCount() bool {
	if o != nil && !IsNil(o.LoginsCount) {
		return true
	}

	return false
}

// SetLoginsCount gets a reference to the given int32 and assigns it to the LoginsCount field.
func (o *GetUsers200ResponseOneOfInner) SetLoginsCount(v int32) {
	o.LoginsCount = &v
}

// GetBlocked returns the Blocked field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetBlocked() bool {
	if o == nil || IsNil(o.Blocked) {
		var ret bool
		return ret
	}
	return *o.Blocked
}

// GetBlockedOk returns a tuple with the Blocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetBlockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Blocked) {
		return nil, false
	}
	return o.Blocked, true
}

// HasBlocked returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasBlocked() bool {
	if o != nil && !IsNil(o.Blocked) {
		return true
	}

	return false
}

// SetBlocked gets a reference to the given bool and assigns it to the Blocked field.
func (o *GetUsers200ResponseOneOfInner) SetBlocked(v bool) {
	o.Blocked = &v
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetGivenName() string {
	if o == nil || IsNil(o.GivenName) {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetGivenNameOk() (*string, bool) {
	if o == nil || IsNil(o.GivenName) {
		return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasGivenName() bool {
	if o != nil && !IsNil(o.GivenName) {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *GetUsers200ResponseOneOfInner) SetGivenName(v string) {
	o.GivenName = &v
}

// GetFamilyName returns the FamilyName field value if set, zero value otherwise.
func (o *GetUsers200ResponseOneOfInner) GetFamilyName() string {
	if o == nil || IsNil(o.FamilyName) {
		var ret string
		return ret
	}
	return *o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetFamilyNameOk() (*string, bool) {
	if o == nil || IsNil(o.FamilyName) {
		return nil, false
	}
	return o.FamilyName, true
}

// HasFamilyName returns a boolean if a field has been set.
func (o *GetUsers200ResponseOneOfInner) HasFamilyName() bool {
	if o != nil && !IsNil(o.FamilyName) {
		return true
	}

	return false
}

// SetFamilyName gets a reference to the given string and assigns it to the FamilyName field.
func (o *GetUsers200ResponseOneOfInner) SetFamilyName(v string) {
	o.FamilyName = &v
}

func (o GetUsers200ResponseOneOfInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUsers200ResponseOneOfInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.EmailVerified) {
		toSerialize["email_verified"] = o.EmailVerified
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !IsNil(o.PhoneVerified) {
		toSerialize["phone_verified"] = o.PhoneVerified
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Identities) {
		toSerialize["identities"] = o.Identities
	}
	if !IsNil(o.AppMetadata) {
		toSerialize["app_metadata"] = o.AppMetadata
	}
	if !IsNil(o.UserMetadata) {
		toSerialize["user_metadata"] = o.UserMetadata
	}
	if !IsNil(o.Picture) {
		toSerialize["picture"] = o.Picture
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Nickname) {
		toSerialize["nickname"] = o.Nickname
	}
	if !IsNil(o.Multifactor) {
		toSerialize["multifactor"] = o.Multifactor
	}
	if !IsNil(o.LastIp) {
		toSerialize["last_ip"] = o.LastIp
	}
	if !IsNil(o.LastLogin) {
		toSerialize["last_login"] = o.LastLogin
	}
	if !IsNil(o.LoginsCount) {
		toSerialize["logins_count"] = o.LoginsCount
	}
	if !IsNil(o.Blocked) {
		toSerialize["blocked"] = o.Blocked
	}
	if !IsNil(o.GivenName) {
		toSerialize["given_name"] = o.GivenName
	}
	if !IsNil(o.FamilyName) {
		toSerialize["family_name"] = o.FamilyName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetUsers200ResponseOneOfInner) UnmarshalJSON(data []byte) (err error) {
	varGetUsers200ResponseOneOfInner := _GetUsers200ResponseOneOfInner{}

	err = json.Unmarshal(data, &varGetUsers200ResponseOneOfInner)

	if err != nil {
		return err
	}

	*o = GetUsers200ResponseOneOfInner(varGetUsers200ResponseOneOfInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "email")
		delete(additionalProperties, "email_verified")
		delete(additionalProperties, "username")
		delete(additionalProperties, "phone_number")
		delete(additionalProperties, "phone_verified")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "identities")
		delete(additionalProperties, "app_metadata")
		delete(additionalProperties, "user_metadata")
		delete(additionalProperties, "picture")
		delete(additionalProperties, "name")
		delete(additionalProperties, "nickname")
		delete(additionalProperties, "multifactor")
		delete(additionalProperties, "last_ip")
		delete(additionalProperties, "last_login")
		delete(additionalProperties, "logins_count")
		delete(additionalProperties, "blocked")
		delete(additionalProperties, "given_name")
		delete(additionalProperties, "family_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUsers200ResponseOneOfInner struct {
	value *GetUsers200ResponseOneOfInner
	isSet bool
}

func (v NullableGetUsers200ResponseOneOfInner) Get() *GetUsers200ResponseOneOfInner {
	return v.value
}

func (v *NullableGetUsers200ResponseOneOfInner) Set(val *GetUsers200ResponseOneOfInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUsers200ResponseOneOfInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUsers200ResponseOneOfInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUsers200ResponseOneOfInner(val *GetUsers200ResponseOneOfInner) *NullableGetUsers200ResponseOneOfInner {
	return &NullableGetUsers200ResponseOneOfInner{value: val, isSet: true}
}

func (v NullableGetUsers200ResponseOneOfInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUsers200ResponseOneOfInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
