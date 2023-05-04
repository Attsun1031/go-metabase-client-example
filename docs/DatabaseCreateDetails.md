# DatabaseCreateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **string** |  | [optional] 
**ServiceAccountJson** | Pointer to **string** |  | [optional] 
**DatasetFiltersType** | Pointer to **string** |  | [optional] 
**AdvancedOptions** | Pointer to **bool** |  | [optional] 
**UseJvmTimezone** | Pointer to **bool** |  | [optional] 
**IncludeUserIdAndHash** | Pointer to **bool** |  | [optional] 
**LetUserControlScheduling** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Engine** | Pointer to **string** |  | [optional] 

## Methods

### NewDatabaseCreateDetails

`func NewDatabaseCreateDetails() *DatabaseCreateDetails`

NewDatabaseCreateDetails instantiates a new DatabaseCreateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseCreateDetailsWithDefaults

`func NewDatabaseCreateDetailsWithDefaults() *DatabaseCreateDetails`

NewDatabaseCreateDetailsWithDefaults instantiates a new DatabaseCreateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *DatabaseCreateDetails) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DatabaseCreateDetails) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DatabaseCreateDetails) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *DatabaseCreateDetails) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetServiceAccountJson

`func (o *DatabaseCreateDetails) GetServiceAccountJson() string`

GetServiceAccountJson returns the ServiceAccountJson field if non-nil, zero value otherwise.

### GetServiceAccountJsonOk

`func (o *DatabaseCreateDetails) GetServiceAccountJsonOk() (*string, bool)`

GetServiceAccountJsonOk returns a tuple with the ServiceAccountJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountJson

`func (o *DatabaseCreateDetails) SetServiceAccountJson(v string)`

SetServiceAccountJson sets ServiceAccountJson field to given value.

### HasServiceAccountJson

`func (o *DatabaseCreateDetails) HasServiceAccountJson() bool`

HasServiceAccountJson returns a boolean if a field has been set.

### GetDatasetFiltersType

`func (o *DatabaseCreateDetails) GetDatasetFiltersType() string`

GetDatasetFiltersType returns the DatasetFiltersType field if non-nil, zero value otherwise.

### GetDatasetFiltersTypeOk

`func (o *DatabaseCreateDetails) GetDatasetFiltersTypeOk() (*string, bool)`

GetDatasetFiltersTypeOk returns a tuple with the DatasetFiltersType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetFiltersType

`func (o *DatabaseCreateDetails) SetDatasetFiltersType(v string)`

SetDatasetFiltersType sets DatasetFiltersType field to given value.

### HasDatasetFiltersType

`func (o *DatabaseCreateDetails) HasDatasetFiltersType() bool`

HasDatasetFiltersType returns a boolean if a field has been set.

### GetAdvancedOptions

`func (o *DatabaseCreateDetails) GetAdvancedOptions() bool`

GetAdvancedOptions returns the AdvancedOptions field if non-nil, zero value otherwise.

### GetAdvancedOptionsOk

`func (o *DatabaseCreateDetails) GetAdvancedOptionsOk() (*bool, bool)`

GetAdvancedOptionsOk returns a tuple with the AdvancedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedOptions

`func (o *DatabaseCreateDetails) SetAdvancedOptions(v bool)`

SetAdvancedOptions sets AdvancedOptions field to given value.

### HasAdvancedOptions

`func (o *DatabaseCreateDetails) HasAdvancedOptions() bool`

HasAdvancedOptions returns a boolean if a field has been set.

### GetUseJvmTimezone

`func (o *DatabaseCreateDetails) GetUseJvmTimezone() bool`

GetUseJvmTimezone returns the UseJvmTimezone field if non-nil, zero value otherwise.

### GetUseJvmTimezoneOk

`func (o *DatabaseCreateDetails) GetUseJvmTimezoneOk() (*bool, bool)`

GetUseJvmTimezoneOk returns a tuple with the UseJvmTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseJvmTimezone

`func (o *DatabaseCreateDetails) SetUseJvmTimezone(v bool)`

SetUseJvmTimezone sets UseJvmTimezone field to given value.

### HasUseJvmTimezone

`func (o *DatabaseCreateDetails) HasUseJvmTimezone() bool`

HasUseJvmTimezone returns a boolean if a field has been set.

### GetIncludeUserIdAndHash

`func (o *DatabaseCreateDetails) GetIncludeUserIdAndHash() bool`

GetIncludeUserIdAndHash returns the IncludeUserIdAndHash field if non-nil, zero value otherwise.

### GetIncludeUserIdAndHashOk

`func (o *DatabaseCreateDetails) GetIncludeUserIdAndHashOk() (*bool, bool)`

GetIncludeUserIdAndHashOk returns a tuple with the IncludeUserIdAndHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUserIdAndHash

`func (o *DatabaseCreateDetails) SetIncludeUserIdAndHash(v bool)`

SetIncludeUserIdAndHash sets IncludeUserIdAndHash field to given value.

### HasIncludeUserIdAndHash

`func (o *DatabaseCreateDetails) HasIncludeUserIdAndHash() bool`

HasIncludeUserIdAndHash returns a boolean if a field has been set.

### GetLetUserControlScheduling

`func (o *DatabaseCreateDetails) GetLetUserControlScheduling() bool`

GetLetUserControlScheduling returns the LetUserControlScheduling field if non-nil, zero value otherwise.

### GetLetUserControlSchedulingOk

`func (o *DatabaseCreateDetails) GetLetUserControlSchedulingOk() (*bool, bool)`

GetLetUserControlSchedulingOk returns a tuple with the LetUserControlScheduling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetUserControlScheduling

`func (o *DatabaseCreateDetails) SetLetUserControlScheduling(v bool)`

SetLetUserControlScheduling sets LetUserControlScheduling field to given value.

### HasLetUserControlScheduling

`func (o *DatabaseCreateDetails) HasLetUserControlScheduling() bool`

HasLetUserControlScheduling returns a boolean if a field has been set.

### GetName

`func (o *DatabaseCreateDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseCreateDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseCreateDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DatabaseCreateDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEngine

`func (o *DatabaseCreateDetails) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *DatabaseCreateDetails) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *DatabaseCreateDetails) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *DatabaseCreateDetails) HasEngine() bool`

HasEngine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


