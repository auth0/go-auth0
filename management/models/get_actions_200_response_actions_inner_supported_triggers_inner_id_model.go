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

// GetActions200ResponseActionsInnerSupportedTriggersInnerId An actions extensibility point.
type GetActions200ResponseActionsInnerSupportedTriggersInnerId string

// List of get_actions_200_response_actions_inner_supported_triggers_inner_id
const (
	GETACTIONS200RESPONSEACTIONSINNERSUPPORTEDTRIGGERSINNERID_POST_LOGIN                    GetActions200ResponseActionsInnerSupportedTriggersInnerId = "post-login"
	GETACTIONS200RESPONSEACTIONSINNERSUPPORTEDTRIGGERSINNERID_CREDENTIALS_EXCHANGE          GetActions200ResponseActionsInnerSupportedTriggersInnerId = "credentials-exchange"
	GETACTIONS200RESPONSEACTIONSINNERSUPPORTEDTRIGGERSINNERID_PRE_USER_REGISTRATION         GetActions200ResponseActionsInnerSupportedTriggersInnerId = "pre-user-registration"
	GETACTIONS200RESPONSEACTIONSINNERSUPPORTEDTRIGGERSINNERID_POST_USER_REGISTRATION        GetActions200ResponseActionsInnerSupportedTriggersInnerId = "post-user-registration"
	GETACTIONS200RESPONSEACTIONSINNERSUPPORTEDTRIGGERSINNERID_POST_CHANGE_PASSWORD          GetActions200ResponseActionsInnerSupportedTriggersInnerId = "post-change-password"
	GETACTIONS200RESPONSEACTIONSINNERSUPPORTEDTRIGGERSINNERID_SEND_PHONE_MESSAGE            GetActions200ResponseActionsInnerSupportedTriggersInnerId = "send-phone-message"
	GETACTIONS200RESPONSEACTIONSINNERSUPPORTEDTRIGGERSINNERID_CUSTOM_PHONE_PROVIDER         GetActions200ResponseActionsInnerSupportedTriggersInnerId = "custom-phone-provider"
	GETACTIONS200RESPONSEACTIONSINNERSUPPORTEDTRIGGERSINNERID_CUSTOM_EMAIL_PROVIDER         GetActions200ResponseActionsInnerSupportedTriggersInnerId = "custom-email-provider"
	GETACTIONS200RESPONSEACTIONSINNERSUPPORTEDTRIGGERSINNERID_PASSWORD_RESET_POST_CHALLENGE GetActions200ResponseActionsInnerSupportedTriggersInnerId = "password-reset-post-challenge"
	GETACTIONS200RESPONSEACTIONSINNERSUPPORTEDTRIGGERSINNERID_CUSTOM_TOKEN_EXCHANGE_BETA    GetActions200ResponseActionsInnerSupportedTriggersInnerId = "custom-token-exchange-beta"
	GETACTIONS200RESPONSEACTIONSINNERSUPPORTEDTRIGGERSINNERID_CUSTOM_TOKEN_EXCHANGE         GetActions200ResponseActionsInnerSupportedTriggersInnerId = "custom-token-exchange"
)

// All allowed values of GetActions200ResponseActionsInnerSupportedTriggersInnerId enum
var AllowedGetActions200ResponseActionsInnerSupportedTriggersInnerIdEnumValues = []GetActions200ResponseActionsInnerSupportedTriggersInnerId{
	"post-login",
	"credentials-exchange",
	"pre-user-registration",
	"post-user-registration",
	"post-change-password",
	"send-phone-message",
	"custom-phone-provider",
	"custom-email-provider",
	"password-reset-post-challenge",
	"custom-token-exchange-beta",
	"custom-token-exchange",
}

func (v *GetActions200ResponseActionsInnerSupportedTriggersInnerId) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetActions200ResponseActionsInnerSupportedTriggersInnerId(value)
	for _, existing := range AllowedGetActions200ResponseActionsInnerSupportedTriggersInnerIdEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetActions200ResponseActionsInnerSupportedTriggersInnerId", value)
}

// NewGetActions200ResponseActionsInnerSupportedTriggersInnerIdFromValue returns a pointer to a valid GetActions200ResponseActionsInnerSupportedTriggersInnerId
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetActions200ResponseActionsInnerSupportedTriggersInnerIdFromValue(v string) (*GetActions200ResponseActionsInnerSupportedTriggersInnerId, error) {
	ev := GetActions200ResponseActionsInnerSupportedTriggersInnerId(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetActions200ResponseActionsInnerSupportedTriggersInnerId: valid values are %v", v, AllowedGetActions200ResponseActionsInnerSupportedTriggersInnerIdEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetActions200ResponseActionsInnerSupportedTriggersInnerId) IsValid() bool {
	for _, existing := range AllowedGetActions200ResponseActionsInnerSupportedTriggersInnerIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to get_actions_200_response_actions_inner_supported_triggers_inner_id value
func (v GetActions200ResponseActionsInnerSupportedTriggersInnerId) Ptr() *GetActions200ResponseActionsInnerSupportedTriggersInnerId {
	return &v
}

type NullableGetActions200ResponseActionsInnerSupportedTriggersInnerId struct {
	value *GetActions200ResponseActionsInnerSupportedTriggersInnerId
	isSet bool
}

func (v NullableGetActions200ResponseActionsInnerSupportedTriggersInnerId) Get() *GetActions200ResponseActionsInnerSupportedTriggersInnerId {
	return v.value
}

func (v *NullableGetActions200ResponseActionsInnerSupportedTriggersInnerId) Set(val *GetActions200ResponseActionsInnerSupportedTriggersInnerId) {
	v.value = val
	v.isSet = true
}

func (v NullableGetActions200ResponseActionsInnerSupportedTriggersInnerId) IsSet() bool {
	return v.isSet
}

func (v *NullableGetActions200ResponseActionsInnerSupportedTriggersInnerId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetActions200ResponseActionsInnerSupportedTriggersInnerId(val *GetActions200ResponseActionsInnerSupportedTriggersInnerId) *NullableGetActions200ResponseActionsInnerSupportedTriggersInnerId {
	return &NullableGetActions200ResponseActionsInnerSupportedTriggersInnerId{value: val, isSet: true}
}

func (v NullableGetActions200ResponseActionsInnerSupportedTriggersInnerId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetActions200ResponseActionsInnerSupportedTriggersInnerId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
