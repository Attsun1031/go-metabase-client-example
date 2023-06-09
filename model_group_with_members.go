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

// checks if the GroupWithMembers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupWithMembers{}

// GroupWithMembers struct for GroupWithMembers
type GroupWithMembers struct {
	Id      *int64        `json:"id,omitempty"`
	Name    *string       `json:"name,omitempty"`
	Members []GroupMember `json:"members,omitempty"`
}

// NewGroupWithMembers instantiates a new GroupWithMembers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupWithMembers() *GroupWithMembers {
	this := GroupWithMembers{}
	return &this
}

// NewGroupWithMembersWithDefaults instantiates a new GroupWithMembers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWithMembersWithDefaults() *GroupWithMembers {
	this := GroupWithMembers{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupWithMembers) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupWithMembers) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupWithMembers) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GroupWithMembers) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupWithMembers) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupWithMembers) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupWithMembers) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupWithMembers) SetName(v string) {
	o.Name = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *GroupWithMembers) GetMembers() []GroupMember {
	if o == nil || IsNil(o.Members) {
		var ret []GroupMember
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupWithMembers) GetMembersOk() ([]GroupMember, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *GroupWithMembers) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []GroupMember and assigns it to the Members field.
func (o *GroupWithMembers) SetMembers(v []GroupMember) {
	o.Members = v
}

func (o GroupWithMembers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupWithMembers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	return toSerialize, nil
}

type NullableGroupWithMembers struct {
	value *GroupWithMembers
	isSet bool
}

func (v NullableGroupWithMembers) Get() *GroupWithMembers {
	return v.value
}

func (v *NullableGroupWithMembers) Set(val *GroupWithMembers) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupWithMembers) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupWithMembers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupWithMembers(val *GroupWithMembers) *NullableGroupWithMembers {
	return &NullableGroupWithMembers{value: val, isSet: true}
}

func (v NullableGroupWithMembers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupWithMembers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
