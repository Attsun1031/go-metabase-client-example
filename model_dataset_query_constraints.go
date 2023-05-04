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

// checks if the DatasetQueryConstraints type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasetQueryConstraints{}

// DatasetQueryConstraints struct for DatasetQueryConstraints
type DatasetQueryConstraints struct {
	MaxResults *int64 `json:"max-results,omitempty"`
	MaxResultsBareRows *int64 `json:"max-results-bare-rows,omitempty"`
}

// NewDatasetQueryConstraints instantiates a new DatasetQueryConstraints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetQueryConstraints() *DatasetQueryConstraints {
	this := DatasetQueryConstraints{}
	return &this
}

// NewDatasetQueryConstraintsWithDefaults instantiates a new DatasetQueryConstraints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetQueryConstraintsWithDefaults() *DatasetQueryConstraints {
	this := DatasetQueryConstraints{}
	return &this
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *DatasetQueryConstraints) GetMaxResults() int64 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int64
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetQueryConstraints) GetMaxResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *DatasetQueryConstraints) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int64 and assigns it to the MaxResults field.
func (o *DatasetQueryConstraints) SetMaxResults(v int64) {
	o.MaxResults = &v
}

// GetMaxResultsBareRows returns the MaxResultsBareRows field value if set, zero value otherwise.
func (o *DatasetQueryConstraints) GetMaxResultsBareRows() int64 {
	if o == nil || IsNil(o.MaxResultsBareRows) {
		var ret int64
		return ret
	}
	return *o.MaxResultsBareRows
}

// GetMaxResultsBareRowsOk returns a tuple with the MaxResultsBareRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetQueryConstraints) GetMaxResultsBareRowsOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxResultsBareRows) {
		return nil, false
	}
	return o.MaxResultsBareRows, true
}

// HasMaxResultsBareRows returns a boolean if a field has been set.
func (o *DatasetQueryConstraints) HasMaxResultsBareRows() bool {
	if o != nil && !IsNil(o.MaxResultsBareRows) {
		return true
	}

	return false
}

// SetMaxResultsBareRows gets a reference to the given int64 and assigns it to the MaxResultsBareRows field.
func (o *DatasetQueryConstraints) SetMaxResultsBareRows(v int64) {
	o.MaxResultsBareRows = &v
}

func (o DatasetQueryConstraints) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasetQueryConstraints) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxResults) {
		toSerialize["max-results"] = o.MaxResults
	}
	if !IsNil(o.MaxResultsBareRows) {
		toSerialize["max-results-bare-rows"] = o.MaxResultsBareRows
	}
	return toSerialize, nil
}

type NullableDatasetQueryConstraints struct {
	value *DatasetQueryConstraints
	isSet bool
}

func (v NullableDatasetQueryConstraints) Get() *DatasetQueryConstraints {
	return v.value
}

func (v *NullableDatasetQueryConstraints) Set(val *DatasetQueryConstraints) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetQueryConstraints) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetQueryConstraints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetQueryConstraints(val *DatasetQueryConstraints) *NullableDatasetQueryConstraints {
	return &NullableDatasetQueryConstraints{value: val, isSet: true}
}

func (v NullableDatasetQueryConstraints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetQueryConstraints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


