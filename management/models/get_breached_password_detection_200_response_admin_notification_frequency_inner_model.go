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

// GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner the model 'GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner'
type GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner string

// List of get_breached_password_detection_200_response_admin_notification_frequency_inner
const (
	GETBREACHEDPASSWORDDETECTION200RESPONSEADMINNOTIFICATIONFREQUENCYINNER_IMMEDIATELY GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner = "immediately"
	GETBREACHEDPASSWORDDETECTION200RESPONSEADMINNOTIFICATIONFREQUENCYINNER_DAILY       GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner = "daily"
	GETBREACHEDPASSWORDDETECTION200RESPONSEADMINNOTIFICATIONFREQUENCYINNER_WEEKLY      GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner = "weekly"
	GETBREACHEDPASSWORDDETECTION200RESPONSEADMINNOTIFICATIONFREQUENCYINNER_MONTHLY     GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner = "monthly"
)

// All allowed values of GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner enum
var AllowedGetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInnerEnumValues = []GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner{
	"immediately",
	"daily",
	"weekly",
	"monthly",
}

func (v *GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner(value)
	for _, existing := range AllowedGetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInnerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner", value)
}

// NewGetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInnerFromValue returns a pointer to a valid GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInnerFromValue(v string) (*GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner, error) {
	ev := GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner: valid values are %v", v, AllowedGetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner) IsValid() bool {
	for _, existing := range AllowedGetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to get_breached_password_detection_200_response_admin_notification_frequency_inner value
func (v GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner) Ptr() *GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner {
	return &v
}

type NullableGetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner struct {
	value *GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner
	isSet bool
}

func (v NullableGetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner) Get() *GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner {
	return v.value
}

func (v *NullableGetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner) Set(val *GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner(val *GetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner) *NullableGetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner {
	return &NullableGetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner{value: val, isSet: true}
}

func (v NullableGetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBreachedPasswordDetection200ResponseAdminNotificationFrequencyInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
