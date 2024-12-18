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

// PatchClientGrantsByIdRequest struct for PatchClientGrantsByIdRequest
type PatchClientGrantsByIdRequest struct {
	// Scopes allowed for this client grant.
	Scope             []string                                              `json:"scope,omitempty"`
	OrganizationUsage NullablePatchClientGrantsByIdRequestOrganizationUsage `json:"organization_usage,omitempty"`
	// Controls allowing any organization to be used with this grant
	AllowAnyOrganization NullableBool `json:"allow_any_organization,omitempty"`
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *PatchClientGrantsByIdRequest) GetScope() []string {
	if o == nil || IsNil(o.Scope) {
		var ret []string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchClientGrantsByIdRequest) GetScopeOk() ([]string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *PatchClientGrantsByIdRequest) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given []string and assigns it to the Scope field.
func (o *PatchClientGrantsByIdRequest) SetScope(v []string) {
	o.Scope = v
}

// GetOrganizationUsage returns the OrganizationUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchClientGrantsByIdRequest) GetOrganizationUsage() PatchClientGrantsByIdRequestOrganizationUsage {
	if o == nil || IsNil(o.OrganizationUsage.Get()) {
		var ret PatchClientGrantsByIdRequestOrganizationUsage
		return ret
	}
	return *o.OrganizationUsage.Get()
}

// GetOrganizationUsageOk returns a tuple with the OrganizationUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchClientGrantsByIdRequest) GetOrganizationUsageOk() (*PatchClientGrantsByIdRequestOrganizationUsage, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationUsage.Get(), o.OrganizationUsage.IsSet()
}

// HasOrganizationUsage returns a boolean if a field has been set.
func (o *PatchClientGrantsByIdRequest) HasOrganizationUsage() bool {
	if o != nil && o.OrganizationUsage.IsSet() {
		return true
	}

	return false
}

// SetOrganizationUsage gets a reference to the given NullablePatchClientGrantsByIdRequestOrganizationUsage and assigns it to the OrganizationUsage field.
func (o *PatchClientGrantsByIdRequest) SetOrganizationUsage(v PatchClientGrantsByIdRequestOrganizationUsage) {
	o.OrganizationUsage.Set(&v)
}

// SetOrganizationUsageNil sets the value for OrganizationUsage to be an explicit nil
func (o *PatchClientGrantsByIdRequest) SetOrganizationUsageNil() {
	o.OrganizationUsage.Set(nil)
}

// UnsetOrganizationUsage ensures that no value is present for OrganizationUsage, not even an explicit nil
func (o *PatchClientGrantsByIdRequest) UnsetOrganizationUsage() {
	o.OrganizationUsage.Unset()
}

// GetAllowAnyOrganization returns the AllowAnyOrganization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchClientGrantsByIdRequest) GetAllowAnyOrganization() bool {
	if o == nil || IsNil(o.AllowAnyOrganization.Get()) {
		var ret bool
		return ret
	}
	return *o.AllowAnyOrganization.Get()
}

// GetAllowAnyOrganizationOk returns a tuple with the AllowAnyOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchClientGrantsByIdRequest) GetAllowAnyOrganizationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowAnyOrganization.Get(), o.AllowAnyOrganization.IsSet()
}

// HasAllowAnyOrganization returns a boolean if a field has been set.
func (o *PatchClientGrantsByIdRequest) HasAllowAnyOrganization() bool {
	if o != nil && o.AllowAnyOrganization.IsSet() {
		return true
	}

	return false
}

// SetAllowAnyOrganization gets a reference to the given NullableBool and assigns it to the AllowAnyOrganization field.
func (o *PatchClientGrantsByIdRequest) SetAllowAnyOrganization(v bool) {
	o.AllowAnyOrganization.Set(&v)
}

// SetAllowAnyOrganizationNil sets the value for AllowAnyOrganization to be an explicit nil
func (o *PatchClientGrantsByIdRequest) SetAllowAnyOrganizationNil() {
	o.AllowAnyOrganization.Set(nil)
}

// UnsetAllowAnyOrganization ensures that no value is present for AllowAnyOrganization, not even an explicit nil
func (o *PatchClientGrantsByIdRequest) UnsetAllowAnyOrganization() {
	o.AllowAnyOrganization.Unset()
}

func (o PatchClientGrantsByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchClientGrantsByIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if o.OrganizationUsage.IsSet() {
		toSerialize["organization_usage"] = o.OrganizationUsage.Get()
	}
	if o.AllowAnyOrganization.IsSet() {
		toSerialize["allow_any_organization"] = o.AllowAnyOrganization.Get()
	}
	return toSerialize, nil
}

type NullablePatchClientGrantsByIdRequest struct {
	value *PatchClientGrantsByIdRequest
	isSet bool
}

func (v NullablePatchClientGrantsByIdRequest) Get() *PatchClientGrantsByIdRequest {
	return v.value
}

func (v *NullablePatchClientGrantsByIdRequest) Set(val *PatchClientGrantsByIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchClientGrantsByIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchClientGrantsByIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchClientGrantsByIdRequest(val *PatchClientGrantsByIdRequest) *NullablePatchClientGrantsByIdRequest {
	return &NullablePatchClientGrantsByIdRequest{value: val, isSet: true}
}

func (v NullablePatchClientGrantsByIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchClientGrantsByIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
