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

// UserEnrollmentStatus Status of this enrollment. Can be `pending` or `confirmed`.
type UserEnrollmentStatus string

// List of userEnrollment_status
const (
	USERENROLLMENTSTATUS_PENDING   UserEnrollmentStatus = "pending"
	USERENROLLMENTSTATUS_CONFIRMED UserEnrollmentStatus = "confirmed"
)

// All allowed values of UserEnrollmentStatus enum
var AllowedUserEnrollmentStatusEnumValues = []UserEnrollmentStatus{
	"pending",
	"confirmed",
}

func (v *UserEnrollmentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserEnrollmentStatus(value)
	for _, existing := range AllowedUserEnrollmentStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserEnrollmentStatus", value)
}

// NewUserEnrollmentStatusFromValue returns a pointer to a valid UserEnrollmentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserEnrollmentStatusFromValue(v string) (*UserEnrollmentStatus, error) {
	ev := UserEnrollmentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserEnrollmentStatus: valid values are %v", v, AllowedUserEnrollmentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserEnrollmentStatus) IsValid() bool {
	for _, existing := range AllowedUserEnrollmentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to userEnrollment_status value
func (v UserEnrollmentStatus) Ptr() *UserEnrollmentStatus {
	return &v
}

type NullableUserEnrollmentStatus struct {
	value *UserEnrollmentStatus
	isSet bool
}

func (v NullableUserEnrollmentStatus) Get() *UserEnrollmentStatus {
	return v.value
}

func (v *NullableUserEnrollmentStatus) Set(val *UserEnrollmentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableUserEnrollmentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableUserEnrollmentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserEnrollmentStatus(val *UserEnrollmentStatus) *NullableUserEnrollmentStatus {
	return &NullableUserEnrollmentStatus{value: val, isSet: true}
}

func (v NullableUserEnrollmentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserEnrollmentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
