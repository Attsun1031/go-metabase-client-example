/*
Metabase

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gometabase

import (
	"encoding/json"
)

// checks if the GroupUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupUpdateRequest{}

// GroupUpdateRequest struct for GroupUpdateRequest
type GroupUpdateRequest struct {
	Name string `json:"name"`
}

// NewGroupUpdateRequest instantiates a new GroupUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupUpdateRequest(name string) *GroupUpdateRequest {
	this := GroupUpdateRequest{}
	this.Name = name
	return &this
}

// NewGroupUpdateRequestWithDefaults instantiates a new GroupUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupUpdateRequestWithDefaults() *GroupUpdateRequest {
	this := GroupUpdateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *GroupUpdateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupUpdateRequest) SetName(v string) {
	o.Name = v
}

func (o GroupUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableGroupUpdateRequest struct {
	value *GroupUpdateRequest
	isSet bool
}

func (v NullableGroupUpdateRequest) Get() *GroupUpdateRequest {
	return v.value
}

func (v *NullableGroupUpdateRequest) Set(val *GroupUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUpdateRequest(val *GroupUpdateRequest) *NullableGroupUpdateRequest {
	return &NullableGroupUpdateRequest{value: val, isSet: true}
}

func (v NullableGroupUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
