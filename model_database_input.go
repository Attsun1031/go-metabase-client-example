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

// checks if the DatabaseInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseInput{}

// DatabaseInput struct for DatabaseInput
type DatabaseInput struct {
	IsOnDemand *bool `json:"is_on_demand,omitempty"`
	IsFullSync *bool `json:"is_full_sync,omitempty"`
	IsSample *bool `json:"is_sample,omitempty"`
	CacheTtl *string `json:"cache_ttl,omitempty"`
	Refingerprint *bool `json:"refingerprint,omitempty"`
	AutoRunQueries *bool `json:"auto_run_queries,omitempty"`
	Schedules *DatabaseInputSchedules `json:"schedules,omitempty"`
	Details *DatabaseInputDetails `json:"details,omitempty"`
}

// NewDatabaseInput instantiates a new DatabaseInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseInput() *DatabaseInput {
	this := DatabaseInput{}
	return &this
}

// NewDatabaseInputWithDefaults instantiates a new DatabaseInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseInputWithDefaults() *DatabaseInput {
	this := DatabaseInput{}
	return &this
}

// GetIsOnDemand returns the IsOnDemand field value if set, zero value otherwise.
func (o *DatabaseInput) GetIsOnDemand() bool {
	if o == nil || IsNil(o.IsOnDemand) {
		var ret bool
		return ret
	}
	return *o.IsOnDemand
}

// GetIsOnDemandOk returns a tuple with the IsOnDemand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseInput) GetIsOnDemandOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOnDemand) {
		return nil, false
	}
	return o.IsOnDemand, true
}

// HasIsOnDemand returns a boolean if a field has been set.
func (o *DatabaseInput) HasIsOnDemand() bool {
	if o != nil && !IsNil(o.IsOnDemand) {
		return true
	}

	return false
}

// SetIsOnDemand gets a reference to the given bool and assigns it to the IsOnDemand field.
func (o *DatabaseInput) SetIsOnDemand(v bool) {
	o.IsOnDemand = &v
}

// GetIsFullSync returns the IsFullSync field value if set, zero value otherwise.
func (o *DatabaseInput) GetIsFullSync() bool {
	if o == nil || IsNil(o.IsFullSync) {
		var ret bool
		return ret
	}
	return *o.IsFullSync
}

// GetIsFullSyncOk returns a tuple with the IsFullSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseInput) GetIsFullSyncOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFullSync) {
		return nil, false
	}
	return o.IsFullSync, true
}

// HasIsFullSync returns a boolean if a field has been set.
func (o *DatabaseInput) HasIsFullSync() bool {
	if o != nil && !IsNil(o.IsFullSync) {
		return true
	}

	return false
}

// SetIsFullSync gets a reference to the given bool and assigns it to the IsFullSync field.
func (o *DatabaseInput) SetIsFullSync(v bool) {
	o.IsFullSync = &v
}

// GetIsSample returns the IsSample field value if set, zero value otherwise.
func (o *DatabaseInput) GetIsSample() bool {
	if o == nil || IsNil(o.IsSample) {
		var ret bool
		return ret
	}
	return *o.IsSample
}

// GetIsSampleOk returns a tuple with the IsSample field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseInput) GetIsSampleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSample) {
		return nil, false
	}
	return o.IsSample, true
}

// HasIsSample returns a boolean if a field has been set.
func (o *DatabaseInput) HasIsSample() bool {
	if o != nil && !IsNil(o.IsSample) {
		return true
	}

	return false
}

// SetIsSample gets a reference to the given bool and assigns it to the IsSample field.
func (o *DatabaseInput) SetIsSample(v bool) {
	o.IsSample = &v
}

// GetCacheTtl returns the CacheTtl field value if set, zero value otherwise.
func (o *DatabaseInput) GetCacheTtl() string {
	if o == nil || IsNil(o.CacheTtl) {
		var ret string
		return ret
	}
	return *o.CacheTtl
}

// GetCacheTtlOk returns a tuple with the CacheTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseInput) GetCacheTtlOk() (*string, bool) {
	if o == nil || IsNil(o.CacheTtl) {
		return nil, false
	}
	return o.CacheTtl, true
}

// HasCacheTtl returns a boolean if a field has been set.
func (o *DatabaseInput) HasCacheTtl() bool {
	if o != nil && !IsNil(o.CacheTtl) {
		return true
	}

	return false
}

// SetCacheTtl gets a reference to the given string and assigns it to the CacheTtl field.
func (o *DatabaseInput) SetCacheTtl(v string) {
	o.CacheTtl = &v
}

// GetRefingerprint returns the Refingerprint field value if set, zero value otherwise.
func (o *DatabaseInput) GetRefingerprint() bool {
	if o == nil || IsNil(o.Refingerprint) {
		var ret bool
		return ret
	}
	return *o.Refingerprint
}

// GetRefingerprintOk returns a tuple with the Refingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseInput) GetRefingerprintOk() (*bool, bool) {
	if o == nil || IsNil(o.Refingerprint) {
		return nil, false
	}
	return o.Refingerprint, true
}

// HasRefingerprint returns a boolean if a field has been set.
func (o *DatabaseInput) HasRefingerprint() bool {
	if o != nil && !IsNil(o.Refingerprint) {
		return true
	}

	return false
}

// SetRefingerprint gets a reference to the given bool and assigns it to the Refingerprint field.
func (o *DatabaseInput) SetRefingerprint(v bool) {
	o.Refingerprint = &v
}

// GetAutoRunQueries returns the AutoRunQueries field value if set, zero value otherwise.
func (o *DatabaseInput) GetAutoRunQueries() bool {
	if o == nil || IsNil(o.AutoRunQueries) {
		var ret bool
		return ret
	}
	return *o.AutoRunQueries
}

// GetAutoRunQueriesOk returns a tuple with the AutoRunQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseInput) GetAutoRunQueriesOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoRunQueries) {
		return nil, false
	}
	return o.AutoRunQueries, true
}

// HasAutoRunQueries returns a boolean if a field has been set.
func (o *DatabaseInput) HasAutoRunQueries() bool {
	if o != nil && !IsNil(o.AutoRunQueries) {
		return true
	}

	return false
}

// SetAutoRunQueries gets a reference to the given bool and assigns it to the AutoRunQueries field.
func (o *DatabaseInput) SetAutoRunQueries(v bool) {
	o.AutoRunQueries = &v
}

// GetSchedules returns the Schedules field value if set, zero value otherwise.
func (o *DatabaseInput) GetSchedules() DatabaseInputSchedules {
	if o == nil || IsNil(o.Schedules) {
		var ret DatabaseInputSchedules
		return ret
	}
	return *o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseInput) GetSchedulesOk() (*DatabaseInputSchedules, bool) {
	if o == nil || IsNil(o.Schedules) {
		return nil, false
	}
	return o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *DatabaseInput) HasSchedules() bool {
	if o != nil && !IsNil(o.Schedules) {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given DatabaseInputSchedules and assigns it to the Schedules field.
func (o *DatabaseInput) SetSchedules(v DatabaseInputSchedules) {
	o.Schedules = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *DatabaseInput) GetDetails() DatabaseInputDetails {
	if o == nil || IsNil(o.Details) {
		var ret DatabaseInputDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseInput) GetDetailsOk() (*DatabaseInputDetails, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *DatabaseInput) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given DatabaseInputDetails and assigns it to the Details field.
func (o *DatabaseInput) SetDetails(v DatabaseInputDetails) {
	o.Details = &v
}

func (o DatabaseInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsOnDemand) {
		toSerialize["is_on_demand"] = o.IsOnDemand
	}
	if !IsNil(o.IsFullSync) {
		toSerialize["is_full_sync"] = o.IsFullSync
	}
	if !IsNil(o.IsSample) {
		toSerialize["is_sample"] = o.IsSample
	}
	if !IsNil(o.CacheTtl) {
		toSerialize["cache_ttl"] = o.CacheTtl
	}
	if !IsNil(o.Refingerprint) {
		toSerialize["refingerprint"] = o.Refingerprint
	}
	if !IsNil(o.AutoRunQueries) {
		toSerialize["auto_run_queries"] = o.AutoRunQueries
	}
	if !IsNil(o.Schedules) {
		toSerialize["schedules"] = o.Schedules
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableDatabaseInput struct {
	value *DatabaseInput
	isSet bool
}

func (v NullableDatabaseInput) Get() *DatabaseInput {
	return v.value
}

func (v *NullableDatabaseInput) Set(val *DatabaseInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseInput(val *DatabaseInput) *NullableDatabaseInput {
	return &NullableDatabaseInput{value: val, isSet: true}
}

func (v NullableDatabaseInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

