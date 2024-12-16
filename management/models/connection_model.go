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

// Connection struct for Connection
type Connection struct {
	// The name of the connection
	Name string `json:"name"`
	// Connection name used in login screen
	DisplayName string                 `json:"display_name"`
	Options     map[string]interface{} `json:"options"`
	// The connection's identifier
	Id string `json:"id"`
	// The type of the connection, related to the identity provider
	Strategy string `json:"strategy"`
	// Defines the realms for which the connection will be used (ie: email domains). If the array is empty or the property is not specified, the connection name will be added as realm.
	Realms []string `json:"realms"`
	// The ids of the clients for which the connection is enabled
	EnabledClients []string `json:"enabled_clients"`
	// True if the connection is domain level
	IsDomainConnection bool `json:"is_domain_connection"`
	// Metadata associated with the connection in the form of an object with string values (max 255 chars).  Maximum of 10 metadata properties allowed.
	Metadata map[string]interface{} `json:"metadata"`
}

type _Connection Connection

// GetName returns the Name field value
func (o *Connection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Connection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Connection) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *Connection) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *Connection) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *Connection) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetOptions returns the Options field value
func (o *Connection) GetOptions() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *Connection) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *Connection) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetId returns the Id field value
func (o *Connection) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Connection) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Connection) SetId(v string) {
	o.Id = v
}

// GetStrategy returns the Strategy field value
func (o *Connection) GetStrategy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *Connection) GetStrategyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value
func (o *Connection) SetStrategy(v string) {
	o.Strategy = v
}

// GetRealms returns the Realms field value
func (o *Connection) GetRealms() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Realms
}

// GetRealmsOk returns a tuple with the Realms field value
// and a boolean to check if the value has been set.
func (o *Connection) GetRealmsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Realms, true
}

// SetRealms sets field value
func (o *Connection) SetRealms(v []string) {
	o.Realms = v
}

// GetEnabledClients returns the EnabledClients field value
func (o *Connection) GetEnabledClients() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EnabledClients
}

// GetEnabledClientsOk returns a tuple with the EnabledClients field value
// and a boolean to check if the value has been set.
func (o *Connection) GetEnabledClientsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnabledClients, true
}

// SetEnabledClients sets field value
func (o *Connection) SetEnabledClients(v []string) {
	o.EnabledClients = v
}

// GetIsDomainConnection returns the IsDomainConnection field value
func (o *Connection) GetIsDomainConnection() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDomainConnection
}

// GetIsDomainConnectionOk returns a tuple with the IsDomainConnection field value
// and a boolean to check if the value has been set.
func (o *Connection) GetIsDomainConnectionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDomainConnection, true
}

// SetIsDomainConnection sets field value
func (o *Connection) SetIsDomainConnection(v bool) {
	o.IsDomainConnection = v
}

// GetMetadata returns the Metadata field value
func (o *Connection) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Connection) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *Connection) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o Connection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Connection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["display_name"] = o.DisplayName
	toSerialize["options"] = o.Options
	toSerialize["id"] = o.Id
	toSerialize["strategy"] = o.Strategy
	toSerialize["realms"] = o.Realms
	toSerialize["enabled_clients"] = o.EnabledClients
	toSerialize["is_domain_connection"] = o.IsDomainConnection
	toSerialize["metadata"] = o.Metadata
	return toSerialize, nil
}

func (o *Connection) UnmarshalJSON(data []byte) (err error) {
	varConnection := _Connection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConnection)

	if err != nil {
		return err
	}

	*o = Connection(varConnection)

	return err
}

type NullableConnection struct {
	value *Connection
	isSet bool
}

func (v NullableConnection) Get() *Connection {
	return v.value
}

func (v *NullableConnection) Set(val *Connection) {
	v.value = val
	v.isSet = true
}

func (v NullableConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnection(val *Connection) *NullableConnection {
	return &NullableConnection{value: val, isSet: true}
}

func (v NullableConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
