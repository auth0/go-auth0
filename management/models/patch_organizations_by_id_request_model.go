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

// PatchOrganizationsByIdRequest struct for PatchOrganizationsByIdRequest
type PatchOrganizationsByIdRequest struct {
	// Friendly name of this organization.
	DisplayName *string `json:"display_name,omitempty"`
	// The name of this organization.
	Name     *string                                       `json:"name,omitempty"`
	Branding NullablePatchOrganizationsByIdRequestBranding `json:"branding,omitempty"`
	// Metadata associated with the organization, in the form of an object with string values (max 255 chars).  Maximum of 10 metadata properties allowed.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PatchOrganizationsByIdRequest) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchOrganizationsByIdRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PatchOrganizationsByIdRequest) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PatchOrganizationsByIdRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchOrganizationsByIdRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchOrganizationsByIdRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchOrganizationsByIdRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchOrganizationsByIdRequest) SetName(v string) {
	o.Name = &v
}

// GetBranding returns the Branding field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchOrganizationsByIdRequest) GetBranding() PatchOrganizationsByIdRequestBranding {
	if o == nil || IsNil(o.Branding.Get()) {
		var ret PatchOrganizationsByIdRequestBranding
		return ret
	}
	return *o.Branding.Get()
}

// GetBrandingOk returns a tuple with the Branding field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchOrganizationsByIdRequest) GetBrandingOk() (*PatchOrganizationsByIdRequestBranding, bool) {
	if o == nil {
		return nil, false
	}
	return o.Branding.Get(), o.Branding.IsSet()
}

// HasBranding returns a boolean if a field has been set.
func (o *PatchOrganizationsByIdRequest) HasBranding() bool {
	if o != nil && o.Branding.IsSet() {
		return true
	}

	return false
}

// SetBranding gets a reference to the given NullablePatchOrganizationsByIdRequestBranding and assigns it to the Branding field.
func (o *PatchOrganizationsByIdRequest) SetBranding(v PatchOrganizationsByIdRequestBranding) {
	o.Branding.Set(&v)
}

// SetBrandingNil sets the value for Branding to be an explicit nil
func (o *PatchOrganizationsByIdRequest) SetBrandingNil() {
	o.Branding.Set(nil)
}

// UnsetBranding ensures that no value is present for Branding, not even an explicit nil
func (o *PatchOrganizationsByIdRequest) UnsetBranding() {
	o.Branding.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchOrganizationsByIdRequest) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchOrganizationsByIdRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PatchOrganizationsByIdRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PatchOrganizationsByIdRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PatchOrganizationsByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchOrganizationsByIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Branding.IsSet() {
		toSerialize["branding"] = o.Branding.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullablePatchOrganizationsByIdRequest struct {
	value *PatchOrganizationsByIdRequest
	isSet bool
}

func (v NullablePatchOrganizationsByIdRequest) Get() *PatchOrganizationsByIdRequest {
	return v.value
}

func (v *NullablePatchOrganizationsByIdRequest) Set(val *PatchOrganizationsByIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchOrganizationsByIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchOrganizationsByIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchOrganizationsByIdRequest(val *PatchOrganizationsByIdRequest) *NullablePatchOrganizationsByIdRequest {
	return &NullablePatchOrganizationsByIdRequest{value: val, isSet: true}
}

func (v NullablePatchOrganizationsByIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchOrganizationsByIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
