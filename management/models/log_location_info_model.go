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

// LogLocationInfo Information about the location that triggered this event based on the `ip`.
type LogLocationInfo struct {
	// Two-letter <a href=\"https://www.iso.org/iso-3166-country-codes.html\">Alpha-2 ISO 3166-1</a> country code.
	CountryCode string `json:"country_code"`
	// Three-letter <a href=\"https://www.iso.org/iso-3166-country-codes.html\">Alpha-3 ISO 3166-1</a> country code.
	CountryCode3 string `json:"country_code3"`
	// Full country name in English.
	CountryName string `json:"country_name"`
	// Full city name in English.
	CityName string `json:"city_name"`
	// Global latitude (horizontal) position.
	Latitude string `json:"latitude"`
	// Global longitude (vertical) position.
	Longitude string `json:"longitude"`
	// Time zone name as found in the <a href=\"https://www.iana.org/time-zones\">tz database</a>.
	TimeZone string `json:"time_zone"`
	// Continent the country is located within. Can be `AF` (Africa), `AN` (Antarctica), `AS` (Asia), `EU` (Europe), `NA` (North America), `OC` (Oceania) or `SA` (South America).
	ContinentCode        string `json:"continent_code"`
	AdditionalProperties map[string]interface{}
}

type _LogLocationInfo LogLocationInfo

// GetCountryCode returns the CountryCode field value
func (o *LogLocationInfo) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *LogLocationInfo) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *LogLocationInfo) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetCountryCode3 returns the CountryCode3 field value
func (o *LogLocationInfo) GetCountryCode3() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode3
}

// GetCountryCode3Ok returns a tuple with the CountryCode3 field value
// and a boolean to check if the value has been set.
func (o *LogLocationInfo) GetCountryCode3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode3, true
}

// SetCountryCode3 sets field value
func (o *LogLocationInfo) SetCountryCode3(v string) {
	o.CountryCode3 = v
}

// GetCountryName returns the CountryName field value
func (o *LogLocationInfo) GetCountryName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryName
}

// GetCountryNameOk returns a tuple with the CountryName field value
// and a boolean to check if the value has been set.
func (o *LogLocationInfo) GetCountryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryName, true
}

// SetCountryName sets field value
func (o *LogLocationInfo) SetCountryName(v string) {
	o.CountryName = v
}

// GetCityName returns the CityName field value
func (o *LogLocationInfo) GetCityName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CityName
}

// GetCityNameOk returns a tuple with the CityName field value
// and a boolean to check if the value has been set.
func (o *LogLocationInfo) GetCityNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CityName, true
}

// SetCityName sets field value
func (o *LogLocationInfo) SetCityName(v string) {
	o.CityName = v
}

// GetLatitude returns the Latitude field value
func (o *LogLocationInfo) GetLatitude() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *LogLocationInfo) GetLatitudeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *LogLocationInfo) SetLatitude(v string) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *LogLocationInfo) GetLongitude() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *LogLocationInfo) GetLongitudeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *LogLocationInfo) SetLongitude(v string) {
	o.Longitude = v
}

// GetTimeZone returns the TimeZone field value
func (o *LogLocationInfo) GetTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value
// and a boolean to check if the value has been set.
func (o *LogLocationInfo) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeZone, true
}

// SetTimeZone sets field value
func (o *LogLocationInfo) SetTimeZone(v string) {
	o.TimeZone = v
}

// GetContinentCode returns the ContinentCode field value
func (o *LogLocationInfo) GetContinentCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContinentCode
}

// GetContinentCodeOk returns a tuple with the ContinentCode field value
// and a boolean to check if the value has been set.
func (o *LogLocationInfo) GetContinentCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContinentCode, true
}

// SetContinentCode sets field value
func (o *LogLocationInfo) SetContinentCode(v string) {
	o.ContinentCode = v
}

func (o LogLocationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogLocationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["country_code"] = o.CountryCode
	toSerialize["country_code3"] = o.CountryCode3
	toSerialize["country_name"] = o.CountryName
	toSerialize["city_name"] = o.CityName
	toSerialize["latitude"] = o.Latitude
	toSerialize["longitude"] = o.Longitude
	toSerialize["time_zone"] = o.TimeZone
	toSerialize["continent_code"] = o.ContinentCode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogLocationInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"country_code",
		"country_code3",
		"country_name",
		"city_name",
		"latitude",
		"longitude",
		"time_zone",
		"continent_code",
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

	varLogLocationInfo := _LogLocationInfo{}

	err = json.Unmarshal(data, &varLogLocationInfo)

	if err != nil {
		return err
	}

	*o = LogLocationInfo(varLogLocationInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "country_code")
		delete(additionalProperties, "country_code3")
		delete(additionalProperties, "country_name")
		delete(additionalProperties, "city_name")
		delete(additionalProperties, "latitude")
		delete(additionalProperties, "longitude")
		delete(additionalProperties, "time_zone")
		delete(additionalProperties, "continent_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogLocationInfo struct {
	value *LogLocationInfo
	isSet bool
}

func (v NullableLogLocationInfo) Get() *LogLocationInfo {
	return v.value
}

func (v *NullableLogLocationInfo) Set(val *LogLocationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLogLocationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLogLocationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogLocationInfo(val *LogLocationInfo) *NullableLogLocationInfo {
	return &NullableLogLocationInfo{value: val, isSet: true}
}

func (v NullableLogLocationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogLocationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
