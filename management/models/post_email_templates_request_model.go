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

// PostEmailTemplatesRequest struct for PostEmailTemplatesRequest
type PostEmailTemplatesRequest struct {
	Template PostEmailTemplatesRequestTemplate `json:"template"`
	// Body of the email template.
	Body NullableString `json:"body"`
	// Senders `from` email address.
	From NullableString `json:"from"`
	// URL to redirect the user to after a successful action.
	ResultUrl NullableString `json:"resultUrl,omitempty"`
	// Subject line of the email.
	Subject NullableString `json:"subject"`
	// Syntax of the template body.
	Syntax NullableString `json:"syntax"`
	// Lifetime in seconds that the link within the email will be valid for.
	UrlLifetimeInSeconds NullableFloat32 `json:"urlLifetimeInSeconds,omitempty"`
	// Whether the `reset_email` and `verify_email` templates should include the user's email address as the `email` parameter in the returnUrl (true) or whether no email address should be included in the redirect (false). Defaults to true.
	IncludeEmailInRedirect *bool `json:"includeEmailInRedirect,omitempty"`
	// Whether the template is enabled (true) or disabled (false).
	Enabled NullableBool `json:"enabled"`
}

type _PostEmailTemplatesRequest PostEmailTemplatesRequest

// GetTemplate returns the Template field value
func (o *PostEmailTemplatesRequest) GetTemplate() PostEmailTemplatesRequestTemplate {
	if o == nil {
		var ret PostEmailTemplatesRequestTemplate
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *PostEmailTemplatesRequest) GetTemplateOk() (*PostEmailTemplatesRequestTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *PostEmailTemplatesRequest) SetTemplate(v PostEmailTemplatesRequestTemplate) {
	o.Template = v
}

// GetBody returns the Body field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PostEmailTemplatesRequest) GetBody() string {
	if o == nil || o.Body.Get() == nil {
		var ret string
		return ret
	}

	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostEmailTemplatesRequest) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// SetBody sets field value
func (o *PostEmailTemplatesRequest) SetBody(v string) {
	o.Body.Set(&v)
}

// GetFrom returns the From field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PostEmailTemplatesRequest) GetFrom() string {
	if o == nil || o.From.Get() == nil {
		var ret string
		return ret
	}

	return *o.From.Get()
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostEmailTemplatesRequest) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.From.Get(), o.From.IsSet()
}

// SetFrom sets field value
func (o *PostEmailTemplatesRequest) SetFrom(v string) {
	o.From.Set(&v)
}

// GetResultUrl returns the ResultUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostEmailTemplatesRequest) GetResultUrl() string {
	if o == nil || IsNil(o.ResultUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ResultUrl.Get()
}

// GetResultUrlOk returns a tuple with the ResultUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostEmailTemplatesRequest) GetResultUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResultUrl.Get(), o.ResultUrl.IsSet()
}

// HasResultUrl returns a boolean if a field has been set.
func (o *PostEmailTemplatesRequest) HasResultUrl() bool {
	if o != nil && o.ResultUrl.IsSet() {
		return true
	}

	return false
}

// SetResultUrl gets a reference to the given NullableString and assigns it to the ResultUrl field.
func (o *PostEmailTemplatesRequest) SetResultUrl(v string) {
	o.ResultUrl.Set(&v)
}

// SetResultUrlNil sets the value for ResultUrl to be an explicit nil
func (o *PostEmailTemplatesRequest) SetResultUrlNil() {
	o.ResultUrl.Set(nil)
}

// UnsetResultUrl ensures that no value is present for ResultUrl, not even an explicit nil
func (o *PostEmailTemplatesRequest) UnsetResultUrl() {
	o.ResultUrl.Unset()
}

// GetSubject returns the Subject field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PostEmailTemplatesRequest) GetSubject() string {
	if o == nil || o.Subject.Get() == nil {
		var ret string
		return ret
	}

	return *o.Subject.Get()
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostEmailTemplatesRequest) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subject.Get(), o.Subject.IsSet()
}

// SetSubject sets field value
func (o *PostEmailTemplatesRequest) SetSubject(v string) {
	o.Subject.Set(&v)
}

// GetSyntax returns the Syntax field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PostEmailTemplatesRequest) GetSyntax() string {
	if o == nil || o.Syntax.Get() == nil {
		var ret string
		return ret
	}

	return *o.Syntax.Get()
}

// GetSyntaxOk returns a tuple with the Syntax field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostEmailTemplatesRequest) GetSyntaxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Syntax.Get(), o.Syntax.IsSet()
}

// SetSyntax sets field value
func (o *PostEmailTemplatesRequest) SetSyntax(v string) {
	o.Syntax.Set(&v)
}

// GetUrlLifetimeInSeconds returns the UrlLifetimeInSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostEmailTemplatesRequest) GetUrlLifetimeInSeconds() float32 {
	if o == nil || IsNil(o.UrlLifetimeInSeconds.Get()) {
		var ret float32
		return ret
	}
	return *o.UrlLifetimeInSeconds.Get()
}

// GetUrlLifetimeInSecondsOk returns a tuple with the UrlLifetimeInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostEmailTemplatesRequest) GetUrlLifetimeInSecondsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UrlLifetimeInSeconds.Get(), o.UrlLifetimeInSeconds.IsSet()
}

// HasUrlLifetimeInSeconds returns a boolean if a field has been set.
func (o *PostEmailTemplatesRequest) HasUrlLifetimeInSeconds() bool {
	if o != nil && o.UrlLifetimeInSeconds.IsSet() {
		return true
	}

	return false
}

// SetUrlLifetimeInSeconds gets a reference to the given NullableFloat32 and assigns it to the UrlLifetimeInSeconds field.
func (o *PostEmailTemplatesRequest) SetUrlLifetimeInSeconds(v float32) {
	o.UrlLifetimeInSeconds.Set(&v)
}

// SetUrlLifetimeInSecondsNil sets the value for UrlLifetimeInSeconds to be an explicit nil
func (o *PostEmailTemplatesRequest) SetUrlLifetimeInSecondsNil() {
	o.UrlLifetimeInSeconds.Set(nil)
}

// UnsetUrlLifetimeInSeconds ensures that no value is present for UrlLifetimeInSeconds, not even an explicit nil
func (o *PostEmailTemplatesRequest) UnsetUrlLifetimeInSeconds() {
	o.UrlLifetimeInSeconds.Unset()
}

// GetIncludeEmailInRedirect returns the IncludeEmailInRedirect field value if set, zero value otherwise.
func (o *PostEmailTemplatesRequest) GetIncludeEmailInRedirect() bool {
	if o == nil || IsNil(o.IncludeEmailInRedirect) {
		var ret bool
		return ret
	}
	return *o.IncludeEmailInRedirect
}

// GetIncludeEmailInRedirectOk returns a tuple with the IncludeEmailInRedirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostEmailTemplatesRequest) GetIncludeEmailInRedirectOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeEmailInRedirect) {
		return nil, false
	}
	return o.IncludeEmailInRedirect, true
}

// HasIncludeEmailInRedirect returns a boolean if a field has been set.
func (o *PostEmailTemplatesRequest) HasIncludeEmailInRedirect() bool {
	if o != nil && !IsNil(o.IncludeEmailInRedirect) {
		return true
	}

	return false
}

// SetIncludeEmailInRedirect gets a reference to the given bool and assigns it to the IncludeEmailInRedirect field.
func (o *PostEmailTemplatesRequest) SetIncludeEmailInRedirect(v bool) {
	o.IncludeEmailInRedirect = &v
}

// GetEnabled returns the Enabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *PostEmailTemplatesRequest) GetEnabled() bool {
	if o == nil || o.Enabled.Get() == nil {
		var ret bool
		return ret
	}

	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostEmailTemplatesRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// SetEnabled sets field value
func (o *PostEmailTemplatesRequest) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}

func (o PostEmailTemplatesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostEmailTemplatesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["template"] = o.Template
	toSerialize["body"] = o.Body.Get()
	toSerialize["from"] = o.From.Get()
	if o.ResultUrl.IsSet() {
		toSerialize["resultUrl"] = o.ResultUrl.Get()
	}
	toSerialize["subject"] = o.Subject.Get()
	toSerialize["syntax"] = o.Syntax.Get()
	if o.UrlLifetimeInSeconds.IsSet() {
		toSerialize["urlLifetimeInSeconds"] = o.UrlLifetimeInSeconds.Get()
	}
	if !IsNil(o.IncludeEmailInRedirect) {
		toSerialize["includeEmailInRedirect"] = o.IncludeEmailInRedirect
	}
	toSerialize["enabled"] = o.Enabled.Get()
	return toSerialize, nil
}

func (o *PostEmailTemplatesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"template",
		"body",
		"from",
		"subject",
		"syntax",
		"enabled",
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

	varPostEmailTemplatesRequest := _PostEmailTemplatesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostEmailTemplatesRequest)

	if err != nil {
		return err
	}

	*o = PostEmailTemplatesRequest(varPostEmailTemplatesRequest)

	return err
}

type NullablePostEmailTemplatesRequest struct {
	value *PostEmailTemplatesRequest
	isSet bool
}

func (v NullablePostEmailTemplatesRequest) Get() *PostEmailTemplatesRequest {
	return v.value
}

func (v *NullablePostEmailTemplatesRequest) Set(val *PostEmailTemplatesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostEmailTemplatesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostEmailTemplatesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostEmailTemplatesRequest(val *PostEmailTemplatesRequest) *NullablePostEmailTemplatesRequest {
	return &NullablePostEmailTemplatesRequest{value: val, isSet: true}
}

func (v NullablePostEmailTemplatesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostEmailTemplatesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
