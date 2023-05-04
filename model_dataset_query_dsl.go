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

// checks if the DatasetQueryDsl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasetQueryDsl{}

// DatasetQueryDsl Dataset query request and response object
type DatasetQueryDsl struct {
	SourceTable *int64 `json:"source_table,omitempty"`
	Limit *int64 `json:"limit,omitempty"`
	Page *DatasetQueryDslPage `json:"page,omitempty"`
}

// NewDatasetQueryDsl instantiates a new DatasetQueryDsl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetQueryDsl() *DatasetQueryDsl {
	this := DatasetQueryDsl{}
	return &this
}

// NewDatasetQueryDslWithDefaults instantiates a new DatasetQueryDsl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetQueryDslWithDefaults() *DatasetQueryDsl {
	this := DatasetQueryDsl{}
	return &this
}

// GetSourceTable returns the SourceTable field value if set, zero value otherwise.
func (o *DatasetQueryDsl) GetSourceTable() int64 {
	if o == nil || IsNil(o.SourceTable) {
		var ret int64
		return ret
	}
	return *o.SourceTable
}

// GetSourceTableOk returns a tuple with the SourceTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetQueryDsl) GetSourceTableOk() (*int64, bool) {
	if o == nil || IsNil(o.SourceTable) {
		return nil, false
	}
	return o.SourceTable, true
}

// HasSourceTable returns a boolean if a field has been set.
func (o *DatasetQueryDsl) HasSourceTable() bool {
	if o != nil && !IsNil(o.SourceTable) {
		return true
	}

	return false
}

// SetSourceTable gets a reference to the given int64 and assigns it to the SourceTable field.
func (o *DatasetQueryDsl) SetSourceTable(v int64) {
	o.SourceTable = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *DatasetQueryDsl) GetLimit() int64 {
	if o == nil || IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetQueryDsl) GetLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *DatasetQueryDsl) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *DatasetQueryDsl) SetLimit(v int64) {
	o.Limit = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *DatasetQueryDsl) GetPage() DatasetQueryDslPage {
	if o == nil || IsNil(o.Page) {
		var ret DatasetQueryDslPage
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetQueryDsl) GetPageOk() (*DatasetQueryDslPage, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *DatasetQueryDsl) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given DatasetQueryDslPage and assigns it to the Page field.
func (o *DatasetQueryDsl) SetPage(v DatasetQueryDslPage) {
	o.Page = &v
}

func (o DatasetQueryDsl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasetQueryDsl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceTable) {
		toSerialize["source_table"] = o.SourceTable
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	return toSerialize, nil
}

type NullableDatasetQueryDsl struct {
	value *DatasetQueryDsl
	isSet bool
}

func (v NullableDatasetQueryDsl) Get() *DatasetQueryDsl {
	return v.value
}

func (v *NullableDatasetQueryDsl) Set(val *DatasetQueryDsl) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetQueryDsl) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetQueryDsl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetQueryDsl(val *DatasetQueryDsl) *NullableDatasetQueryDsl {
	return &NullableDatasetQueryDsl{value: val, isSet: true}
}

func (v NullableDatasetQueryDsl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetQueryDsl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

