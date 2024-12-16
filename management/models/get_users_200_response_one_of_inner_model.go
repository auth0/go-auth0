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

// GetUsers200ResponseOneOfInner struct for GetUsers200ResponseOneOfInner
type GetUsers200ResponseOneOfInner struct {
	// ID of the user which can be used when interacting with other APIs.
	UserId string `json:"user_id"`
	// Email address of this user.
	Email string `json:"email"`
	// Whether this email address is verified (true) or unverified (false).
	EmailVerified bool `json:"email_verified"`
	// Username of this user.
	Username string `json:"username"`
	// Phone number for this user when using SMS connections. Follows the <a href=\"https://en.wikipedia.org/wiki/E.164\">E.164 recommendation</a>.
	PhoneNumber string `json:"phone_number"`
	// Whether this phone number has been verified (true) or not (false).
	PhoneVerified bool                                   `json:"phone_verified"`
	CreatedAt     GetUsers200ResponseOneOfInnerCreatedAt `json:"created_at"`
	UpdatedAt     GetUsers200ResponseOneOfInnerUpdatedAt `json:"updated_at"`
	// Array of user identity objects when accounts are linked.
	Identities  []GetUsers200ResponseOneOfInnerIdentitiesInner `json:"identities"`
	AppMetadata GetUsers200ResponseOneOfInnerAppMetadata       `json:"app_metadata"`
	// User metadata to which this user has read/write access.
	UserMetadata map[string]interface{} `json:"user_metadata"`
	// URL to picture, photo, or avatar of this user.
	Picture string `json:"picture"`
	// Name of this user.
	Name string `json:"name"`
	// Preferred nickname or alias of this user.
	Nickname string `json:"nickname"`
	// List of multi-factor authentication providers with which this user has enrolled.
	Multifactor []string `json:"multifactor"`
	// Last IP address from which this user logged in.
	LastIp    string                                 `json:"last_ip"`
	LastLogin GetUsers200ResponseOneOfInnerLastLogin `json:"last_login"`
	// Total number of logins this user has performed.
	LoginsCount int32 `json:"logins_count"`
	// Whether this user was blocked by an administrator (true) or is not (false).
	Blocked bool `json:"blocked"`
	// Given name/first name/forename of this user.
	GivenName string `json:"given_name"`
	// Family name/last name/surname of this user.
	FamilyName           string `json:"family_name"`
	AdditionalProperties map[string]interface{}
}

type _GetUsers200ResponseOneOfInner GetUsers200ResponseOneOfInner

// GetUserId returns the UserId field value
func (o *GetUsers200ResponseOneOfInner) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *GetUsers200ResponseOneOfInner) SetUserId(v string) {
	o.UserId = v
}

// GetEmail returns the Email field value
func (o *GetUsers200ResponseOneOfInner) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *GetUsers200ResponseOneOfInner) SetEmail(v string) {
	o.Email = v
}

// GetEmailVerified returns the EmailVerified field value
func (o *GetUsers200ResponseOneOfInner) GetEmailVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EmailVerified
}

// GetEmailVerifiedOk returns a tuple with the EmailVerified field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetEmailVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailVerified, true
}

// SetEmailVerified sets field value
func (o *GetUsers200ResponseOneOfInner) SetEmailVerified(v bool) {
	o.EmailVerified = v
}

// GetUsername returns the Username field value
func (o *GetUsers200ResponseOneOfInner) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *GetUsers200ResponseOneOfInner) SetUsername(v string) {
	o.Username = v
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *GetUsers200ResponseOneOfInner) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *GetUsers200ResponseOneOfInner) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

// GetPhoneVerified returns the PhoneVerified field value
func (o *GetUsers200ResponseOneOfInner) GetPhoneVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PhoneVerified
}

// GetPhoneVerifiedOk returns a tuple with the PhoneVerified field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetPhoneVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneVerified, true
}

// SetPhoneVerified sets field value
func (o *GetUsers200ResponseOneOfInner) SetPhoneVerified(v bool) {
	o.PhoneVerified = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GetUsers200ResponseOneOfInner) GetCreatedAt() GetUsers200ResponseOneOfInnerCreatedAt {
	if o == nil {
		var ret GetUsers200ResponseOneOfInnerCreatedAt
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetCreatedAtOk() (*GetUsers200ResponseOneOfInnerCreatedAt, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GetUsers200ResponseOneOfInner) SetCreatedAt(v GetUsers200ResponseOneOfInnerCreatedAt) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *GetUsers200ResponseOneOfInner) GetUpdatedAt() GetUsers200ResponseOneOfInnerUpdatedAt {
	if o == nil {
		var ret GetUsers200ResponseOneOfInnerUpdatedAt
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetUpdatedAtOk() (*GetUsers200ResponseOneOfInnerUpdatedAt, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *GetUsers200ResponseOneOfInner) SetUpdatedAt(v GetUsers200ResponseOneOfInnerUpdatedAt) {
	o.UpdatedAt = v
}

// GetIdentities returns the Identities field value
func (o *GetUsers200ResponseOneOfInner) GetIdentities() []GetUsers200ResponseOneOfInnerIdentitiesInner {
	if o == nil {
		var ret []GetUsers200ResponseOneOfInnerIdentitiesInner
		return ret
	}

	return o.Identities
}

// GetIdentitiesOk returns a tuple with the Identities field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetIdentitiesOk() ([]GetUsers200ResponseOneOfInnerIdentitiesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identities, true
}

// SetIdentities sets field value
func (o *GetUsers200ResponseOneOfInner) SetIdentities(v []GetUsers200ResponseOneOfInnerIdentitiesInner) {
	o.Identities = v
}

// GetAppMetadata returns the AppMetadata field value
func (o *GetUsers200ResponseOneOfInner) GetAppMetadata() GetUsers200ResponseOneOfInnerAppMetadata {
	if o == nil {
		var ret GetUsers200ResponseOneOfInnerAppMetadata
		return ret
	}

	return o.AppMetadata
}

// GetAppMetadataOk returns a tuple with the AppMetadata field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetAppMetadataOk() (*GetUsers200ResponseOneOfInnerAppMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppMetadata, true
}

// SetAppMetadata sets field value
func (o *GetUsers200ResponseOneOfInner) SetAppMetadata(v GetUsers200ResponseOneOfInnerAppMetadata) {
	o.AppMetadata = v
}

// GetUserMetadata returns the UserMetadata field value
func (o *GetUsers200ResponseOneOfInner) GetUserMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.UserMetadata
}

// GetUserMetadataOk returns a tuple with the UserMetadata field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetUserMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.UserMetadata, true
}

// SetUserMetadata sets field value
func (o *GetUsers200ResponseOneOfInner) SetUserMetadata(v map[string]interface{}) {
	o.UserMetadata = v
}

// GetPicture returns the Picture field value
func (o *GetUsers200ResponseOneOfInner) GetPicture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Picture
}

// GetPictureOk returns a tuple with the Picture field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetPictureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Picture, true
}

// SetPicture sets field value
func (o *GetUsers200ResponseOneOfInner) SetPicture(v string) {
	o.Picture = v
}

// GetName returns the Name field value
func (o *GetUsers200ResponseOneOfInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetUsers200ResponseOneOfInner) SetName(v string) {
	o.Name = v
}

// GetNickname returns the Nickname field value
func (o *GetUsers200ResponseOneOfInner) GetNickname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetNicknameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nickname, true
}

// SetNickname sets field value
func (o *GetUsers200ResponseOneOfInner) SetNickname(v string) {
	o.Nickname = v
}

// GetMultifactor returns the Multifactor field value
func (o *GetUsers200ResponseOneOfInner) GetMultifactor() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Multifactor
}

// GetMultifactorOk returns a tuple with the Multifactor field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetMultifactorOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Multifactor, true
}

// SetMultifactor sets field value
func (o *GetUsers200ResponseOneOfInner) SetMultifactor(v []string) {
	o.Multifactor = v
}

// GetLastIp returns the LastIp field value
func (o *GetUsers200ResponseOneOfInner) GetLastIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastIp
}

// GetLastIpOk returns a tuple with the LastIp field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetLastIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastIp, true
}

// SetLastIp sets field value
func (o *GetUsers200ResponseOneOfInner) SetLastIp(v string) {
	o.LastIp = v
}

// GetLastLogin returns the LastLogin field value
func (o *GetUsers200ResponseOneOfInner) GetLastLogin() GetUsers200ResponseOneOfInnerLastLogin {
	if o == nil {
		var ret GetUsers200ResponseOneOfInnerLastLogin
		return ret
	}

	return o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetLastLoginOk() (*GetUsers200ResponseOneOfInnerLastLogin, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastLogin, true
}

// SetLastLogin sets field value
func (o *GetUsers200ResponseOneOfInner) SetLastLogin(v GetUsers200ResponseOneOfInnerLastLogin) {
	o.LastLogin = v
}

// GetLoginsCount returns the LoginsCount field value
func (o *GetUsers200ResponseOneOfInner) GetLoginsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LoginsCount
}

// GetLoginsCountOk returns a tuple with the LoginsCount field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetLoginsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoginsCount, true
}

// SetLoginsCount sets field value
func (o *GetUsers200ResponseOneOfInner) SetLoginsCount(v int32) {
	o.LoginsCount = v
}

// GetBlocked returns the Blocked field value
func (o *GetUsers200ResponseOneOfInner) GetBlocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Blocked
}

// GetBlockedOk returns a tuple with the Blocked field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetBlockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blocked, true
}

// SetBlocked sets field value
func (o *GetUsers200ResponseOneOfInner) SetBlocked(v bool) {
	o.Blocked = v
}

// GetGivenName returns the GivenName field value
func (o *GetUsers200ResponseOneOfInner) GetGivenName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetGivenNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GivenName, true
}

// SetGivenName sets field value
func (o *GetUsers200ResponseOneOfInner) SetGivenName(v string) {
	o.GivenName = v
}

// GetFamilyName returns the FamilyName field value
func (o *GetUsers200ResponseOneOfInner) GetFamilyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value
// and a boolean to check if the value has been set.
func (o *GetUsers200ResponseOneOfInner) GetFamilyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FamilyName, true
}

// SetFamilyName sets field value
func (o *GetUsers200ResponseOneOfInner) SetFamilyName(v string) {
	o.FamilyName = v
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
	toSerialize["user_id"] = o.UserId
	toSerialize["email"] = o.Email
	toSerialize["email_verified"] = o.EmailVerified
	toSerialize["username"] = o.Username
	toSerialize["phone_number"] = o.PhoneNumber
	toSerialize["phone_verified"] = o.PhoneVerified
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["identities"] = o.Identities
	toSerialize["app_metadata"] = o.AppMetadata
	toSerialize["user_metadata"] = o.UserMetadata
	toSerialize["picture"] = o.Picture
	toSerialize["name"] = o.Name
	toSerialize["nickname"] = o.Nickname
	toSerialize["multifactor"] = o.Multifactor
	toSerialize["last_ip"] = o.LastIp
	toSerialize["last_login"] = o.LastLogin
	toSerialize["logins_count"] = o.LoginsCount
	toSerialize["blocked"] = o.Blocked
	toSerialize["given_name"] = o.GivenName
	toSerialize["family_name"] = o.FamilyName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetUsers200ResponseOneOfInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_id",
		"email",
		"email_verified",
		"username",
		"phone_number",
		"phone_verified",
		"created_at",
		"updated_at",
		"identities",
		"app_metadata",
		"user_metadata",
		"picture",
		"name",
		"nickname",
		"multifactor",
		"last_ip",
		"last_login",
		"logins_count",
		"blocked",
		"given_name",
		"family_name",
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
