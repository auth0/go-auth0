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

// PostEmailTemplatesRequestTemplate Template name. Can be `verify_email`, `verify_email_by_code`, `reset_email`, `reset_email_by_code`, `welcome_email`, `blocked_account`, `stolen_credentials`, `enrollment_email`, `mfa_oob_code`, `user_invitation`, `change_password` (legacy), or `password_reset` (legacy).
type PostEmailTemplatesRequestTemplate string

// List of post_email_templates_request_template
const (
	POSTEMAILTEMPLATESREQUESTTEMPLATE_VERIFY_EMAIL         PostEmailTemplatesRequestTemplate = "verify_email"
	POSTEMAILTEMPLATESREQUESTTEMPLATE_VERIFY_EMAIL_BY_CODE PostEmailTemplatesRequestTemplate = "verify_email_by_code"
	POSTEMAILTEMPLATESREQUESTTEMPLATE_RESET_EMAIL          PostEmailTemplatesRequestTemplate = "reset_email"
	POSTEMAILTEMPLATESREQUESTTEMPLATE_RESET_EMAIL_BY_CODE  PostEmailTemplatesRequestTemplate = "reset_email_by_code"
	POSTEMAILTEMPLATESREQUESTTEMPLATE_WELCOME_EMAIL        PostEmailTemplatesRequestTemplate = "welcome_email"
	POSTEMAILTEMPLATESREQUESTTEMPLATE_BLOCKED_ACCOUNT      PostEmailTemplatesRequestTemplate = "blocked_account"
	POSTEMAILTEMPLATESREQUESTTEMPLATE_STOLEN_CREDENTIALS   PostEmailTemplatesRequestTemplate = "stolen_credentials"
	POSTEMAILTEMPLATESREQUESTTEMPLATE_ENROLLMENT_EMAIL     PostEmailTemplatesRequestTemplate = "enrollment_email"
	POSTEMAILTEMPLATESREQUESTTEMPLATE_MFA_OOB_CODE         PostEmailTemplatesRequestTemplate = "mfa_oob_code"
	POSTEMAILTEMPLATESREQUESTTEMPLATE_USER_INVITATION      PostEmailTemplatesRequestTemplate = "user_invitation"
	POSTEMAILTEMPLATESREQUESTTEMPLATE_CHANGE_PASSWORD      PostEmailTemplatesRequestTemplate = "change_password"
	POSTEMAILTEMPLATESREQUESTTEMPLATE_PASSWORD_RESET       PostEmailTemplatesRequestTemplate = "password_reset"
)

// All allowed values of PostEmailTemplatesRequestTemplate enum
var AllowedPostEmailTemplatesRequestTemplateEnumValues = []PostEmailTemplatesRequestTemplate{
	"verify_email",
	"verify_email_by_code",
	"reset_email",
	"reset_email_by_code",
	"welcome_email",
	"blocked_account",
	"stolen_credentials",
	"enrollment_email",
	"mfa_oob_code",
	"user_invitation",
	"change_password",
	"password_reset",
}

func (v *PostEmailTemplatesRequestTemplate) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PostEmailTemplatesRequestTemplate(value)
	for _, existing := range AllowedPostEmailTemplatesRequestTemplateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PostEmailTemplatesRequestTemplate", value)
}

// NewPostEmailTemplatesRequestTemplateFromValue returns a pointer to a valid PostEmailTemplatesRequestTemplate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPostEmailTemplatesRequestTemplateFromValue(v string) (*PostEmailTemplatesRequestTemplate, error) {
	ev := PostEmailTemplatesRequestTemplate(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PostEmailTemplatesRequestTemplate: valid values are %v", v, AllowedPostEmailTemplatesRequestTemplateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PostEmailTemplatesRequestTemplate) IsValid() bool {
	for _, existing := range AllowedPostEmailTemplatesRequestTemplateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to post_email_templates_request_template value
func (v PostEmailTemplatesRequestTemplate) Ptr() *PostEmailTemplatesRequestTemplate {
	return &v
}

type NullablePostEmailTemplatesRequestTemplate struct {
	value *PostEmailTemplatesRequestTemplate
	isSet bool
}

func (v NullablePostEmailTemplatesRequestTemplate) Get() *PostEmailTemplatesRequestTemplate {
	return v.value
}

func (v *NullablePostEmailTemplatesRequestTemplate) Set(val *PostEmailTemplatesRequestTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullablePostEmailTemplatesRequestTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullablePostEmailTemplatesRequestTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostEmailTemplatesRequestTemplate(val *PostEmailTemplatesRequestTemplate) *NullablePostEmailTemplatesRequestTemplate {
	return &NullablePostEmailTemplatesRequestTemplate{value: val, isSet: true}
}

func (v NullablePostEmailTemplatesRequestTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostEmailTemplatesRequestTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
