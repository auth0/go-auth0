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

// GetExecution200Response The result of a specific execution of a trigger.
type GetExecution200Response struct {
	// ID identifies this specific execution simulation. These IDs would resemble real executions in production.
	Id        *string                                                    `json:"id,omitempty"`
	TriggerId *GetActions200ResponseActionsInnerSupportedTriggersInnerId `json:"trigger_id,omitempty"`
	Status    *GetExecution200ResponseStatus                             `json:"status,omitempty"`
	Results   []GetExecution200ResponseResultsInner                      `json:"results,omitempty"`
	// The time that the execution was started.
	CreatedAt *string `json:"created_at,omitempty"`
	// The time that the exeution finished executing.
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetExecution200Response) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExecution200Response) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetExecution200Response) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetExecution200Response) SetId(v string) {
	o.Id = &v
}

// GetTriggerId returns the TriggerId field value if set, zero value otherwise.
func (o *GetExecution200Response) GetTriggerId() GetActions200ResponseActionsInnerSupportedTriggersInnerId {
	if o == nil || IsNil(o.TriggerId) {
		var ret GetActions200ResponseActionsInnerSupportedTriggersInnerId
		return ret
	}
	return *o.TriggerId
}

// GetTriggerIdOk returns a tuple with the TriggerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExecution200Response) GetTriggerIdOk() (*GetActions200ResponseActionsInnerSupportedTriggersInnerId, bool) {
	if o == nil || IsNil(o.TriggerId) {
		return nil, false
	}
	return o.TriggerId, true
}

// HasTriggerId returns a boolean if a field has been set.
func (o *GetExecution200Response) HasTriggerId() bool {
	if o != nil && !IsNil(o.TriggerId) {
		return true
	}

	return false
}

// SetTriggerId gets a reference to the given GetActions200ResponseActionsInnerSupportedTriggersInnerId and assigns it to the TriggerId field.
func (o *GetExecution200Response) SetTriggerId(v GetActions200ResponseActionsInnerSupportedTriggersInnerId) {
	o.TriggerId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetExecution200Response) GetStatus() GetExecution200ResponseStatus {
	if o == nil || IsNil(o.Status) {
		var ret GetExecution200ResponseStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExecution200Response) GetStatusOk() (*GetExecution200ResponseStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetExecution200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GetExecution200ResponseStatus and assigns it to the Status field.
func (o *GetExecution200Response) SetStatus(v GetExecution200ResponseStatus) {
	o.Status = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *GetExecution200Response) GetResults() []GetExecution200ResponseResultsInner {
	if o == nil || IsNil(o.Results) {
		var ret []GetExecution200ResponseResultsInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExecution200Response) GetResultsOk() ([]GetExecution200ResponseResultsInner, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *GetExecution200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []GetExecution200ResponseResultsInner and assigns it to the Results field.
func (o *GetExecution200Response) SetResults(v []GetExecution200ResponseResultsInner) {
	o.Results = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GetExecution200Response) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExecution200Response) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GetExecution200Response) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GetExecution200Response) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GetExecution200Response) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExecution200Response) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GetExecution200Response) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GetExecution200Response) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o GetExecution200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExecution200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TriggerId) {
		toSerialize["trigger_id"] = o.TriggerId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableGetExecution200Response struct {
	value *GetExecution200Response
	isSet bool
}

func (v NullableGetExecution200Response) Get() *GetExecution200Response {
	return v.value
}

func (v *NullableGetExecution200Response) Set(val *GetExecution200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExecution200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExecution200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExecution200Response(val *GetExecution200Response) *NullableGetExecution200Response {
	return &NullableGetExecution200Response{value: val, isSet: true}
}

func (v NullableGetExecution200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExecution200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
