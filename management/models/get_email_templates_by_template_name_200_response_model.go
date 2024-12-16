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
)

// GetEmailTemplatesByTemplateName200Response struct for GetEmailTemplatesByTemplateName200Response
type GetEmailTemplatesByTemplateName200Response struct {
	Template PostEmailTemplatesRequestTemplate `json:"template"`
	// Body of the email template.
	Body NullableString `json:"body"`
	// Senders `from` email address.
	From NullableString `json:"from"`
	// URL to redirect the user to after a successful action.
	ResultUrl NullableString `json:"resultUrl"`
	// Subject line of the email.
	Subject NullableString `json:"subject"`
	// Syntax of the template body.
	Syntax NullableString `json:"syntax"`
	// Lifetime in seconds that the link within the email will be valid for.
	UrlLifetimeInSeconds NullableFloat32 `json:"urlLifetimeInSeconds"`
	// Whether the `reset_email` and `verify_email` templates should include the user's email address as the `email` parameter in the returnUrl (true) or whether no email address should be included in the redirect (false). Defaults to true.
	IncludeEmailInRedirect bool `json:"includeEmailInRedirect"`
	// Whether the template is enabled (true) or disabled (false).
	Enabled NullableBool `json:"enabled"`
}

type _GetEmailTemplatesByTemplateName200Response GetEmailTemplatesByTemplateName200Response

// GetTemplate returns the Template field value
func (o *GetEmailTemplatesByTemplateName200Response) GetTemplate() PostEmailTemplatesRequestTemplate {
	if o == nil {
		var ret PostEmailTemplatesRequestTemplate
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *GetEmailTemplatesByTemplateName200Response) GetTemplateOk() (*PostEmailTemplatesRequestTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *GetEmailTemplatesByTemplateName200Response) SetTemplate(v PostEmailTemplatesRequestTemplate) {
	o.Template = v
}

// GetBody returns the Body field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetEmailTemplatesByTemplateName200Response) GetBody() string {
	if o == nil || o.Body.Get() == nil {
		var ret string
		return ret
	}

	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetEmailTemplatesByTemplateName200Response) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// SetBody sets field value
func (o *GetEmailTemplatesByTemplateName200Response) SetBody(v string) {
	o.Body.Set(&v)
}

// GetFrom returns the From field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetEmailTemplatesByTemplateName200Response) GetFrom() string {
	if o == nil || o.From.Get() == nil {
		var ret string
		return ret
	}

	return *o.From.Get()
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetEmailTemplatesByTemplateName200Response) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.From.Get(), o.From.IsSet()
}

// SetFrom sets field value
func (o *GetEmailTemplatesByTemplateName200Response) SetFrom(v string) {
	o.From.Set(&v)
}

// GetResultUrl returns the ResultUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetEmailTemplatesByTemplateName200Response) GetResultUrl() string {
	if o == nil || o.ResultUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ResultUrl.Get()
}

// GetResultUrlOk returns a tuple with the ResultUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetEmailTemplatesByTemplateName200Response) GetResultUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResultUrl.Get(), o.ResultUrl.IsSet()
}

// SetResultUrl sets field value
func (o *GetEmailTemplatesByTemplateName200Response) SetResultUrl(v string) {
	o.ResultUrl.Set(&v)
}

// GetSubject returns the Subject field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetEmailTemplatesByTemplateName200Response) GetSubject() string {
	if o == nil || o.Subject.Get() == nil {
		var ret string
		return ret
	}

	return *o.Subject.Get()
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetEmailTemplatesByTemplateName200Response) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subject.Get(), o.Subject.IsSet()
}

// SetSubject sets field value
func (o *GetEmailTemplatesByTemplateName200Response) SetSubject(v string) {
	o.Subject.Set(&v)
}

// GetSyntax returns the Syntax field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetEmailTemplatesByTemplateName200Response) GetSyntax() string {
	if o == nil || o.Syntax.Get() == nil {
		var ret string
		return ret
	}

	return *o.Syntax.Get()
}

// GetSyntaxOk returns a tuple with the Syntax field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetEmailTemplatesByTemplateName200Response) GetSyntaxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Syntax.Get(), o.Syntax.IsSet()
}

// SetSyntax sets field value
func (o *GetEmailTemplatesByTemplateName200Response) SetSyntax(v string) {
	o.Syntax.Set(&v)
}

// GetUrlLifetimeInSeconds returns the UrlLifetimeInSeconds field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *GetEmailTemplatesByTemplateName200Response) GetUrlLifetimeInSeconds() float32 {
	if o == nil || o.UrlLifetimeInSeconds.Get() == nil {
		var ret float32
		return ret
	}

	return *o.UrlLifetimeInSeconds.Get()
}

// GetUrlLifetimeInSecondsOk returns a tuple with the UrlLifetimeInSeconds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetEmailTemplatesByTemplateName200Response) GetUrlLifetimeInSecondsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UrlLifetimeInSeconds.Get(), o.UrlLifetimeInSeconds.IsSet()
}

// SetUrlLifetimeInSeconds sets field value
func (o *GetEmailTemplatesByTemplateName200Response) SetUrlLifetimeInSeconds(v float32) {
	o.UrlLifetimeInSeconds.Set(&v)
}

// GetIncludeEmailInRedirect returns the IncludeEmailInRedirect field value
func (o *GetEmailTemplatesByTemplateName200Response) GetIncludeEmailInRedirect() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeEmailInRedirect
}

// GetIncludeEmailInRedirectOk returns a tuple with the IncludeEmailInRedirect field value
// and a boolean to check if the value has been set.
func (o *GetEmailTemplatesByTemplateName200Response) GetIncludeEmailInRedirectOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeEmailInRedirect, true
}

// SetIncludeEmailInRedirect sets field value
func (o *GetEmailTemplatesByTemplateName200Response) SetIncludeEmailInRedirect(v bool) {
	o.IncludeEmailInRedirect = v
}

// GetEnabled returns the Enabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *GetEmailTemplatesByTemplateName200Response) GetEnabled() bool {
	if o == nil || o.Enabled.Get() == nil {
		var ret bool
		return ret
	}

	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetEmailTemplatesByTemplateName200Response) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// SetEnabled sets field value
func (o *GetEmailTemplatesByTemplateName200Response) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}

func (o GetEmailTemplatesByTemplateName200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEmailTemplatesByTemplateName200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["template"] = o.Template
	toSerialize["body"] = o.Body.Get()
	toSerialize["from"] = o.From.Get()
	toSerialize["resultUrl"] = o.ResultUrl.Get()
	toSerialize["subject"] = o.Subject.Get()
	toSerialize["syntax"] = o.Syntax.Get()
	toSerialize["urlLifetimeInSeconds"] = o.UrlLifetimeInSeconds.Get()
	toSerialize["includeEmailInRedirect"] = o.IncludeEmailInRedirect
	toSerialize["enabled"] = o.Enabled.Get()
	return toSerialize, nil
}

func (o *GetEmailTemplatesByTemplateName200Response) UnmarshalJSON(data []byte) (err error) {
	varGetEmailTemplatesByTemplateName200Response := _GetEmailTemplatesByTemplateName200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetEmailTemplatesByTemplateName200Response)

	if err != nil {
		return err
	}

	*o = GetEmailTemplatesByTemplateName200Response(varGetEmailTemplatesByTemplateName200Response)

	return err
}

type NullableGetEmailTemplatesByTemplateName200Response struct {
	value *GetEmailTemplatesByTemplateName200Response
	isSet bool
}

func (v NullableGetEmailTemplatesByTemplateName200Response) Get() *GetEmailTemplatesByTemplateName200Response {
	return v.value
}

func (v *NullableGetEmailTemplatesByTemplateName200Response) Set(val *GetEmailTemplatesByTemplateName200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEmailTemplatesByTemplateName200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEmailTemplatesByTemplateName200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEmailTemplatesByTemplateName200Response(val *GetEmailTemplatesByTemplateName200Response) *NullableGetEmailTemplatesByTemplateName200Response {
	return &NullableGetEmailTemplatesByTemplateName200Response{value: val, isSet: true}
}

func (v NullableGetEmailTemplatesByTemplateName200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEmailTemplatesByTemplateName200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
