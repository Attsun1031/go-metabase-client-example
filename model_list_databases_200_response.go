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

// checks if the ListDatabases200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDatabases200Response{}

// ListDatabases200Response struct for ListDatabases200Response
type ListDatabases200Response struct {
	Data []Database `json:"data,omitempty"`
}

// NewListDatabases200Response instantiates a new ListDatabases200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDatabases200Response() *ListDatabases200Response {
	this := ListDatabases200Response{}
	return &this
}

// NewListDatabases200ResponseWithDefaults instantiates a new ListDatabases200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDatabases200ResponseWithDefaults() *ListDatabases200Response {
	this := ListDatabases200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListDatabases200Response) GetData() []Database {
	if o == nil || IsNil(o.Data) {
		var ret []Database
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDatabases200Response) GetDataOk() ([]Database, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListDatabases200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Database and assigns it to the Data field.
func (o *ListDatabases200Response) SetData(v []Database) {
	o.Data = v
}

func (o ListDatabases200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDatabases200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListDatabases200Response struct {
	value *ListDatabases200Response
	isSet bool
}

func (v NullableListDatabases200Response) Get() *ListDatabases200Response {
	return v.value
}

func (v *NullableListDatabases200Response) Set(val *ListDatabases200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListDatabases200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListDatabases200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDatabases200Response(val *ListDatabases200Response) *NullableListDatabases200Response {
	return &NullableListDatabases200Response{value: val, isSet: true}
}

func (v NullableListDatabases200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDatabases200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

