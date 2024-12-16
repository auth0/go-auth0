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

// PostOrganizations201Response struct for PostOrganizations201Response
type PostOrganizations201Response struct {
	// Organization identifier
	Id string `json:"id"`
	// The name of this organization.
	Name string `json:"name"`
	// Friendly name of this organization.
	DisplayName string                                        `json:"display_name"`
	Branding    GetOrganizations200ResponseOneOfInnerBranding `json:"branding"`
	// Metadata associated with the organization, in the form of an object with string values (max 255 chars).  Maximum of 10 metadata properties allowed.
	Metadata             map[string]interface{}                                `json:"metadata"`
	EnabledConnections   []PostOrganizations201ResponseEnabledConnectionsInner `json:"enabled_connections"`
	AdditionalProperties map[string]interface{}
}

type _PostOrganizations201Response PostOrganizations201Response

// GetId returns the Id field value
func (o *PostOrganizations201Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostOrganizations201Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostOrganizations201Response) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PostOrganizations201Response) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PostOrganizations201Response) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PostOrganizations201Response) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *PostOrganizations201Response) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *PostOrganizations201Response) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *PostOrganizations201Response) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetBranding returns the Branding field value
func (o *PostOrganizations201Response) GetBranding() GetOrganizations200ResponseOneOfInnerBranding {
	if o == nil {
		var ret GetOrganizations200ResponseOneOfInnerBranding
		return ret
	}

	return o.Branding
}

// GetBrandingOk returns a tuple with the Branding field value
// and a boolean to check if the value has been set.
func (o *PostOrganizations201Response) GetBrandingOk() (*GetOrganizations200ResponseOneOfInnerBranding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Branding, true
}

// SetBranding sets field value
func (o *PostOrganizations201Response) SetBranding(v GetOrganizations200ResponseOneOfInnerBranding) {
	o.Branding = v
}

// GetMetadata returns the Metadata field value
func (o *PostOrganizations201Response) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *PostOrganizations201Response) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *PostOrganizations201Response) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetEnabledConnections returns the EnabledConnections field value
func (o *PostOrganizations201Response) GetEnabledConnections() []PostOrganizations201ResponseEnabledConnectionsInner {
	if o == nil {
		var ret []PostOrganizations201ResponseEnabledConnectionsInner
		return ret
	}

	return o.EnabledConnections
}

// GetEnabledConnectionsOk returns a tuple with the EnabledConnections field value
// and a boolean to check if the value has been set.
func (o *PostOrganizations201Response) GetEnabledConnectionsOk() ([]PostOrganizations201ResponseEnabledConnectionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnabledConnections, true
}

// SetEnabledConnections sets field value
func (o *PostOrganizations201Response) SetEnabledConnections(v []PostOrganizations201ResponseEnabledConnectionsInner) {
	o.EnabledConnections = v
}

func (o PostOrganizations201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostOrganizations201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["display_name"] = o.DisplayName
	toSerialize["branding"] = o.Branding
	toSerialize["metadata"] = o.Metadata
	toSerialize["enabled_connections"] = o.EnabledConnections

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostOrganizations201Response) UnmarshalJSON(data []byte) (err error) {
	varPostOrganizations201Response := _PostOrganizations201Response{}

	err = json.Unmarshal(data, &varPostOrganizations201Response)

	if err != nil {
		return err
	}

	*o = PostOrganizations201Response(varPostOrganizations201Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "branding")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "enabled_connections")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostOrganizations201Response struct {
	value *PostOrganizations201Response
	isSet bool
}

func (v NullablePostOrganizations201Response) Get() *PostOrganizations201Response {
	return v.value
}

func (v *NullablePostOrganizations201Response) Set(val *PostOrganizations201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostOrganizations201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostOrganizations201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostOrganizations201Response(val *PostOrganizations201Response) *NullablePostOrganizations201Response {
	return &NullablePostOrganizations201Response{value: val, isSet: true}
}

func (v NullablePostOrganizations201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostOrganizations201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
