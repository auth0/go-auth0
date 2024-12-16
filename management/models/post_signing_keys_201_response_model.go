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

// PostSigningKeys201Response struct for PostSigningKeys201Response
type PostSigningKeys201Response struct {
	// Next key certificate
	Cert string `json:"cert"`
	// Next key id
	Kid                  string `json:"kid"`
	AdditionalProperties map[string]interface{}
}

type _PostSigningKeys201Response PostSigningKeys201Response

// GetCert returns the Cert field value
func (o *PostSigningKeys201Response) GetCert() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cert
}

// GetCertOk returns a tuple with the Cert field value
// and a boolean to check if the value has been set.
func (o *PostSigningKeys201Response) GetCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cert, true
}

// SetCert sets field value
func (o *PostSigningKeys201Response) SetCert(v string) {
	o.Cert = v
}

// GetKid returns the Kid field value
func (o *PostSigningKeys201Response) GetKid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kid
}

// GetKidOk returns a tuple with the Kid field value
// and a boolean to check if the value has been set.
func (o *PostSigningKeys201Response) GetKidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kid, true
}

// SetKid sets field value
func (o *PostSigningKeys201Response) SetKid(v string) {
	o.Kid = v
}

func (o PostSigningKeys201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostSigningKeys201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cert"] = o.Cert
	toSerialize["kid"] = o.Kid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostSigningKeys201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cert",
		"kid",
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

	varPostSigningKeys201Response := _PostSigningKeys201Response{}

	err = json.Unmarshal(data, &varPostSigningKeys201Response)

	if err != nil {
		return err
	}

	*o = PostSigningKeys201Response(varPostSigningKeys201Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cert")
		delete(additionalProperties, "kid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostSigningKeys201Response struct {
	value *PostSigningKeys201Response
	isSet bool
}

func (v NullablePostSigningKeys201Response) Get() *PostSigningKeys201Response {
	return v.value
}

func (v *NullablePostSigningKeys201Response) Set(val *PostSigningKeys201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSigningKeys201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSigningKeys201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSigningKeys201Response(val *PostSigningKeys201Response) *NullablePostSigningKeys201Response {
	return &NullablePostSigningKeys201Response{value: val, isSet: true}
}

func (v NullablePostSigningKeys201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSigningKeys201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
