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

// EnrollmentCreate struct for EnrollmentCreate
type EnrollmentCreate struct {
	// user_id for the enrollment ticket
	UserId string `json:"user_id"`
	// alternate email to which the enrollment email will be sent. Optional - by default, the email will be sent to the user's default address
	Email *string `json:"email,omitempty"`
	// Send an email to the user to start the enrollment
	SendMail *bool `json:"send_mail,omitempty"`
	// Optional. Specify the locale of the enrollment email. Used with send_email.
	EmailLocale *string `json:"email_locale,omitempty"`
}

type _EnrollmentCreate EnrollmentCreate

// GetUserId returns the UserId field value
func (o *EnrollmentCreate) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCreate) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *EnrollmentCreate) SetUserId(v string) {
	o.UserId = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *EnrollmentCreate) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentCreate) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *EnrollmentCreate) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *EnrollmentCreate) SetEmail(v string) {
	o.Email = &v
}

// GetSendMail returns the SendMail field value if set, zero value otherwise.
func (o *EnrollmentCreate) GetSendMail() bool {
	if o == nil || IsNil(o.SendMail) {
		var ret bool
		return ret
	}
	return *o.SendMail
}

// GetSendMailOk returns a tuple with the SendMail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentCreate) GetSendMailOk() (*bool, bool) {
	if o == nil || IsNil(o.SendMail) {
		return nil, false
	}
	return o.SendMail, true
}

// HasSendMail returns a boolean if a field has been set.
func (o *EnrollmentCreate) HasSendMail() bool {
	if o != nil && !IsNil(o.SendMail) {
		return true
	}

	return false
}

// SetSendMail gets a reference to the given bool and assigns it to the SendMail field.
func (o *EnrollmentCreate) SetSendMail(v bool) {
	o.SendMail = &v
}

// GetEmailLocale returns the EmailLocale field value if set, zero value otherwise.
func (o *EnrollmentCreate) GetEmailLocale() string {
	if o == nil || IsNil(o.EmailLocale) {
		var ret string
		return ret
	}
	return *o.EmailLocale
}

// GetEmailLocaleOk returns a tuple with the EmailLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentCreate) GetEmailLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.EmailLocale) {
		return nil, false
	}
	return o.EmailLocale, true
}

// HasEmailLocale returns a boolean if a field has been set.
func (o *EnrollmentCreate) HasEmailLocale() bool {
	if o != nil && !IsNil(o.EmailLocale) {
		return true
	}

	return false
}

// SetEmailLocale gets a reference to the given string and assigns it to the EmailLocale field.
func (o *EnrollmentCreate) SetEmailLocale(v string) {
	o.EmailLocale = &v
}

func (o EnrollmentCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrollmentCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_id"] = o.UserId
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.SendMail) {
		toSerialize["send_mail"] = o.SendMail
	}
	if !IsNil(o.EmailLocale) {
		toSerialize["email_locale"] = o.EmailLocale
	}
	return toSerialize, nil
}

func (o *EnrollmentCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_id",
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

	varEnrollmentCreate := _EnrollmentCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varEnrollmentCreate)

	if err != nil {
		return err
	}

	*o = EnrollmentCreate(varEnrollmentCreate)

	return err
}

type NullableEnrollmentCreate struct {
	value *EnrollmentCreate
	isSet bool
}

func (v NullableEnrollmentCreate) Get() *EnrollmentCreate {
	return v.value
}

func (v *NullableEnrollmentCreate) Set(val *EnrollmentCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentCreate(val *EnrollmentCreate) *NullableEnrollmentCreate {
	return &NullableEnrollmentCreate{value: val, isSet: true}
}

func (v NullableEnrollmentCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollmentCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
