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

// checks if the DatabaseCreateDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseCreateDetails{}

// DatabaseCreateDetails struct for DatabaseCreateDetails
type DatabaseCreateDetails struct {
	ProjectId                *string `json:"project-id,omitempty"`
	ServiceAccountJson       *string `json:"service-account-json,omitempty"`
	DatasetFiltersType       *string `json:"dataset-filters-type,omitempty"`
	AdvancedOptions          *bool   `json:"advanced-options,omitempty"`
	UseJvmTimezone           *bool   `json:"use-jvm-timezone,omitempty"`
	IncludeUserIdAndHash     *bool   `json:"include-user-id-and-hash,omitempty"`
	LetUserControlScheduling *bool   `json:"let-user-control-scheduling,omitempty"`
	Name                     *string `json:"name,omitempty"`
	Engine                   *string `json:"engine,omitempty"`
}

// NewDatabaseCreateDetails instantiates a new DatabaseCreateDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseCreateDetails() *DatabaseCreateDetails {
	this := DatabaseCreateDetails{}
	return &this
}

// NewDatabaseCreateDetailsWithDefaults instantiates a new DatabaseCreateDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseCreateDetailsWithDefaults() *DatabaseCreateDetails {
	this := DatabaseCreateDetails{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DatabaseCreateDetails) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCreateDetails) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DatabaseCreateDetails) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *DatabaseCreateDetails) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetServiceAccountJson returns the ServiceAccountJson field value if set, zero value otherwise.
func (o *DatabaseCreateDetails) GetServiceAccountJson() string {
	if o == nil || IsNil(o.ServiceAccountJson) {
		var ret string
		return ret
	}
	return *o.ServiceAccountJson
}

// GetServiceAccountJsonOk returns a tuple with the ServiceAccountJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCreateDetails) GetServiceAccountJsonOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceAccountJson) {
		return nil, false
	}
	return o.ServiceAccountJson, true
}

// HasServiceAccountJson returns a boolean if a field has been set.
func (o *DatabaseCreateDetails) HasServiceAccountJson() bool {
	if o != nil && !IsNil(o.ServiceAccountJson) {
		return true
	}

	return false
}

// SetServiceAccountJson gets a reference to the given string and assigns it to the ServiceAccountJson field.
func (o *DatabaseCreateDetails) SetServiceAccountJson(v string) {
	o.ServiceAccountJson = &v
}

// GetDatasetFiltersType returns the DatasetFiltersType field value if set, zero value otherwise.
func (o *DatabaseCreateDetails) GetDatasetFiltersType() string {
	if o == nil || IsNil(o.DatasetFiltersType) {
		var ret string
		return ret
	}
	return *o.DatasetFiltersType
}

// GetDatasetFiltersTypeOk returns a tuple with the DatasetFiltersType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCreateDetails) GetDatasetFiltersTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DatasetFiltersType) {
		return nil, false
	}
	return o.DatasetFiltersType, true
}

// HasDatasetFiltersType returns a boolean if a field has been set.
func (o *DatabaseCreateDetails) HasDatasetFiltersType() bool {
	if o != nil && !IsNil(o.DatasetFiltersType) {
		return true
	}

	return false
}

// SetDatasetFiltersType gets a reference to the given string and assigns it to the DatasetFiltersType field.
func (o *DatabaseCreateDetails) SetDatasetFiltersType(v string) {
	o.DatasetFiltersType = &v
}

// GetAdvancedOptions returns the AdvancedOptions field value if set, zero value otherwise.
func (o *DatabaseCreateDetails) GetAdvancedOptions() bool {
	if o == nil || IsNil(o.AdvancedOptions) {
		var ret bool
		return ret
	}
	return *o.AdvancedOptions
}

// GetAdvancedOptionsOk returns a tuple with the AdvancedOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCreateDetails) GetAdvancedOptionsOk() (*bool, bool) {
	if o == nil || IsNil(o.AdvancedOptions) {
		return nil, false
	}
	return o.AdvancedOptions, true
}

// HasAdvancedOptions returns a boolean if a field has been set.
func (o *DatabaseCreateDetails) HasAdvancedOptions() bool {
	if o != nil && !IsNil(o.AdvancedOptions) {
		return true
	}

	return false
}

// SetAdvancedOptions gets a reference to the given bool and assigns it to the AdvancedOptions field.
func (o *DatabaseCreateDetails) SetAdvancedOptions(v bool) {
	o.AdvancedOptions = &v
}

// GetUseJvmTimezone returns the UseJvmTimezone field value if set, zero value otherwise.
func (o *DatabaseCreateDetails) GetUseJvmTimezone() bool {
	if o == nil || IsNil(o.UseJvmTimezone) {
		var ret bool
		return ret
	}
	return *o.UseJvmTimezone
}

// GetUseJvmTimezoneOk returns a tuple with the UseJvmTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCreateDetails) GetUseJvmTimezoneOk() (*bool, bool) {
	if o == nil || IsNil(o.UseJvmTimezone) {
		return nil, false
	}
	return o.UseJvmTimezone, true
}

// HasUseJvmTimezone returns a boolean if a field has been set.
func (o *DatabaseCreateDetails) HasUseJvmTimezone() bool {
	if o != nil && !IsNil(o.UseJvmTimezone) {
		return true
	}

	return false
}

// SetUseJvmTimezone gets a reference to the given bool and assigns it to the UseJvmTimezone field.
func (o *DatabaseCreateDetails) SetUseJvmTimezone(v bool) {
	o.UseJvmTimezone = &v
}

// GetIncludeUserIdAndHash returns the IncludeUserIdAndHash field value if set, zero value otherwise.
func (o *DatabaseCreateDetails) GetIncludeUserIdAndHash() bool {
	if o == nil || IsNil(o.IncludeUserIdAndHash) {
		var ret bool
		return ret
	}
	return *o.IncludeUserIdAndHash
}

// GetIncludeUserIdAndHashOk returns a tuple with the IncludeUserIdAndHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCreateDetails) GetIncludeUserIdAndHashOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeUserIdAndHash) {
		return nil, false
	}
	return o.IncludeUserIdAndHash, true
}

// HasIncludeUserIdAndHash returns a boolean if a field has been set.
func (o *DatabaseCreateDetails) HasIncludeUserIdAndHash() bool {
	if o != nil && !IsNil(o.IncludeUserIdAndHash) {
		return true
	}

	return false
}

// SetIncludeUserIdAndHash gets a reference to the given bool and assigns it to the IncludeUserIdAndHash field.
func (o *DatabaseCreateDetails) SetIncludeUserIdAndHash(v bool) {
	o.IncludeUserIdAndHash = &v
}

// GetLetUserControlScheduling returns the LetUserControlScheduling field value if set, zero value otherwise.
func (o *DatabaseCreateDetails) GetLetUserControlScheduling() bool {
	if o == nil || IsNil(o.LetUserControlScheduling) {
		var ret bool
		return ret
	}
	return *o.LetUserControlScheduling
}

// GetLetUserControlSchedulingOk returns a tuple with the LetUserControlScheduling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCreateDetails) GetLetUserControlSchedulingOk() (*bool, bool) {
	if o == nil || IsNil(o.LetUserControlScheduling) {
		return nil, false
	}
	return o.LetUserControlScheduling, true
}

// HasLetUserControlScheduling returns a boolean if a field has been set.
func (o *DatabaseCreateDetails) HasLetUserControlScheduling() bool {
	if o != nil && !IsNil(o.LetUserControlScheduling) {
		return true
	}

	return false
}

// SetLetUserControlScheduling gets a reference to the given bool and assigns it to the LetUserControlScheduling field.
func (o *DatabaseCreateDetails) SetLetUserControlScheduling(v bool) {
	o.LetUserControlScheduling = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DatabaseCreateDetails) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCreateDetails) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DatabaseCreateDetails) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DatabaseCreateDetails) SetName(v string) {
	o.Name = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *DatabaseCreateDetails) GetEngine() string {
	if o == nil || IsNil(o.Engine) {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseCreateDetails) GetEngineOk() (*string, bool) {
	if o == nil || IsNil(o.Engine) {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *DatabaseCreateDetails) HasEngine() bool {
	if o != nil && !IsNil(o.Engine) {
		return true
	}

	return false
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *DatabaseCreateDetails) SetEngine(v string) {
	o.Engine = &v
}

func (o DatabaseCreateDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseCreateDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["project-id"] = o.ProjectId
	}
	if !IsNil(o.ServiceAccountJson) {
		toSerialize["service-account-json"] = o.ServiceAccountJson
	}
	if !IsNil(o.DatasetFiltersType) {
		toSerialize["dataset-filters-type"] = o.DatasetFiltersType
	}
	if !IsNil(o.AdvancedOptions) {
		toSerialize["advanced-options"] = o.AdvancedOptions
	}
	if !IsNil(o.UseJvmTimezone) {
		toSerialize["use-jvm-timezone"] = o.UseJvmTimezone
	}
	if !IsNil(o.IncludeUserIdAndHash) {
		toSerialize["include-user-id-and-hash"] = o.IncludeUserIdAndHash
	}
	if !IsNil(o.LetUserControlScheduling) {
		toSerialize["let-user-control-scheduling"] = o.LetUserControlScheduling
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Engine) {
		toSerialize["engine"] = o.Engine
	}
	return toSerialize, nil
}

type NullableDatabaseCreateDetails struct {
	value *DatabaseCreateDetails
	isSet bool
}

func (v NullableDatabaseCreateDetails) Get() *DatabaseCreateDetails {
	return v.value
}

func (v *NullableDatabaseCreateDetails) Set(val *DatabaseCreateDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseCreateDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseCreateDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseCreateDetails(val *DatabaseCreateDetails) *NullableDatabaseCreateDetails {
	return &NullableDatabaseCreateDetails{value: val, isSet: true}
}

func (v NullableDatabaseCreateDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseCreateDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
