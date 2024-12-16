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

// GetOrganizations200ResponseOneOfInner struct for GetOrganizations200ResponseOneOfInner
type GetOrganizations200ResponseOneOfInner struct {
	// Organization identifier
	Id string `json:"id"`
	// The name of this organization.
	Name string `json:"name"`
	// Friendly name of this organization.
	DisplayName string                                        `json:"display_name"`
	Branding    GetOrganizations200ResponseOneOfInnerBranding `json:"branding"`
	// Metadata associated with the organization, in the form of an object with string values (max 255 chars).  Maximum of 10 metadata properties allowed.
	Metadata             map[string]interface{} `json:"metadata"`
	AdditionalProperties map[string]interface{}
}

type _GetOrganizations200ResponseOneOfInner GetOrganizations200ResponseOneOfInner

// GetId returns the Id field value
func (o *GetOrganizations200ResponseOneOfInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetOrganizations200ResponseOneOfInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetOrganizations200ResponseOneOfInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GetOrganizations200ResponseOneOfInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetOrganizations200ResponseOneOfInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetOrganizations200ResponseOneOfInner) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *GetOrganizations200ResponseOneOfInner) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *GetOrganizations200ResponseOneOfInner) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *GetOrganizations200ResponseOneOfInner) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetBranding returns the Branding field value
func (o *GetOrganizations200ResponseOneOfInner) GetBranding() GetOrganizations200ResponseOneOfInnerBranding {
	if o == nil {
		var ret GetOrganizations200ResponseOneOfInnerBranding
		return ret
	}

	return o.Branding
}

// GetBrandingOk returns a tuple with the Branding field value
// and a boolean to check if the value has been set.
func (o *GetOrganizations200ResponseOneOfInner) GetBrandingOk() (*GetOrganizations200ResponseOneOfInnerBranding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Branding, true
}

// SetBranding sets field value
func (o *GetOrganizations200ResponseOneOfInner) SetBranding(v GetOrganizations200ResponseOneOfInnerBranding) {
	o.Branding = v
}

// GetMetadata returns the Metadata field value
func (o *GetOrganizations200ResponseOneOfInner) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *GetOrganizations200ResponseOneOfInner) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *GetOrganizations200ResponseOneOfInner) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GetOrganizations200ResponseOneOfInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizations200ResponseOneOfInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["display_name"] = o.DisplayName
	toSerialize["branding"] = o.Branding
	toSerialize["metadata"] = o.Metadata

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetOrganizations200ResponseOneOfInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"display_name",
		"branding",
		"metadata",
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

	varGetOrganizations200ResponseOneOfInner := _GetOrganizations200ResponseOneOfInner{}

	err = json.Unmarshal(data, &varGetOrganizations200ResponseOneOfInner)

	if err != nil {
		return err
	}

	*o = GetOrganizations200ResponseOneOfInner(varGetOrganizations200ResponseOneOfInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "branding")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOrganizations200ResponseOneOfInner struct {
	value *GetOrganizations200ResponseOneOfInner
	isSet bool
}

func (v NullableGetOrganizations200ResponseOneOfInner) Get() *GetOrganizations200ResponseOneOfInner {
	return v.value
}

func (v *NullableGetOrganizations200ResponseOneOfInner) Set(val *GetOrganizations200ResponseOneOfInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizations200ResponseOneOfInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizations200ResponseOneOfInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizations200ResponseOneOfInner(val *GetOrganizations200ResponseOneOfInner) *NullableGetOrganizations200ResponseOneOfInner {
	return &NullableGetOrganizations200ResponseOneOfInner{value: val, isSet: true}
}

func (v NullableGetOrganizations200ResponseOneOfInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizations200ResponseOneOfInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
