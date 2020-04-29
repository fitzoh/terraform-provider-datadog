/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// CheckCanDeleteSLOResponse A service level objective response containing the requested object.
type CheckCanDeleteSLOResponse struct {
	Data *CheckCanDeleteSLOResponseData `json:"data,omitempty"`
	// A mapping of SLO id to it's current usages.
	Errors *map[string]string `json:"errors,omitempty"`
}

// NewCheckCanDeleteSLOResponse instantiates a new CheckCanDeleteSLOResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckCanDeleteSLOResponse() *CheckCanDeleteSLOResponse {
	this := CheckCanDeleteSLOResponse{}
	return &this
}

// NewCheckCanDeleteSLOResponseWithDefaults instantiates a new CheckCanDeleteSLOResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckCanDeleteSLOResponseWithDefaults() *CheckCanDeleteSLOResponse {
	this := CheckCanDeleteSLOResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CheckCanDeleteSLOResponse) GetData() CheckCanDeleteSLOResponseData {
	if o == nil || o.Data == nil {
		var ret CheckCanDeleteSLOResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckCanDeleteSLOResponse) GetDataOk() (*CheckCanDeleteSLOResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CheckCanDeleteSLOResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given CheckCanDeleteSLOResponseData and assigns it to the Data field.
func (o *CheckCanDeleteSLOResponse) SetData(v CheckCanDeleteSLOResponseData) {
	o.Data = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CheckCanDeleteSLOResponse) GetErrors() map[string]string {
	if o == nil || o.Errors == nil {
		var ret map[string]string
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckCanDeleteSLOResponse) GetErrorsOk() (*map[string]string, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CheckCanDeleteSLOResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given map[string]string and assigns it to the Errors field.
func (o *CheckCanDeleteSLOResponse) SetErrors(v map[string]string) {
	o.Errors = &v
}

func (o CheckCanDeleteSLOResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableCheckCanDeleteSLOResponse struct {
	value *CheckCanDeleteSLOResponse
	isSet bool
}

func (v NullableCheckCanDeleteSLOResponse) Get() *CheckCanDeleteSLOResponse {
	return v.value
}

func (v *NullableCheckCanDeleteSLOResponse) Set(val *CheckCanDeleteSLOResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckCanDeleteSLOResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckCanDeleteSLOResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckCanDeleteSLOResponse(val *CheckCanDeleteSLOResponse) *NullableCheckCanDeleteSLOResponse {
	return &NullableCheckCanDeleteSLOResponse{value: val, isSet: true}
}

func (v NullableCheckCanDeleteSLOResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckCanDeleteSLOResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}