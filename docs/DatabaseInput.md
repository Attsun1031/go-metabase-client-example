# DatabaseInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsOnDemand** | Pointer to **bool** |  | [optional] 
**IsFullSync** | Pointer to **bool** |  | [optional] 
**IsSample** | Pointer to **bool** |  | [optional] 
**CacheTtl** | Pointer to **string** |  | [optional] 
**Refingerprint** | Pointer to **bool** |  | [optional] 
**AutoRunQueries** | Pointer to **bool** |  | [optional] 
**Schedules** | Pointer to [**DatabaseInputSchedules**](DatabaseInputSchedules.md) |  | [optional] 
**Details** | Pointer to [**DatabaseInputDetails**](DatabaseInputDetails.md) |  | [optional] 

## Methods

### NewDatabaseInput

`func NewDatabaseInput() *DatabaseInput`

NewDatabaseInput instantiates a new DatabaseInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseInputWithDefaults

`func NewDatabaseInputWithDefaults() *DatabaseInput`

NewDatabaseInputWithDefaults instantiates a new DatabaseInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsOnDemand

`func (o *DatabaseInput) GetIsOnDemand() bool`

GetIsOnDemand returns the IsOnDemand field if non-nil, zero value otherwise.

### GetIsOnDemandOk

`func (o *DatabaseInput) GetIsOnDemandOk() (*bool, bool)`

GetIsOnDemandOk returns a tuple with the IsOnDemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnDemand

`func (o *DatabaseInput) SetIsOnDemand(v bool)`

SetIsOnDemand sets IsOnDemand field to given value.

### HasIsOnDemand

`func (o *DatabaseInput) HasIsOnDemand() bool`

HasIsOnDemand returns a boolean if a field has been set.

### GetIsFullSync

`func (o *DatabaseInput) GetIsFullSync() bool`

GetIsFullSync returns the IsFullSync field if non-nil, zero value otherwise.

### GetIsFullSyncOk

`func (o *DatabaseInput) GetIsFullSyncOk() (*bool, bool)`

GetIsFullSyncOk returns a tuple with the IsFullSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullSync

`func (o *DatabaseInput) SetIsFullSync(v bool)`

SetIsFullSync sets IsFullSync field to given value.

### HasIsFullSync

`func (o *DatabaseInput) HasIsFullSync() bool`

HasIsFullSync returns a boolean if a field has been set.

### GetIsSample

`func (o *DatabaseInput) GetIsSample() bool`

GetIsSample returns the IsSample field if non-nil, zero value otherwise.

### GetIsSampleOk

`func (o *DatabaseInput) GetIsSampleOk() (*bool, bool)`

GetIsSampleOk returns a tuple with the IsSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSample

`func (o *DatabaseInput) SetIsSample(v bool)`

SetIsSample sets IsSample field to given value.

### HasIsSample

`func (o *DatabaseInput) HasIsSample() bool`

HasIsSample returns a boolean if a field has been set.

### GetCacheTtl

`func (o *DatabaseInput) GetCacheTtl() string`

GetCacheTtl returns the CacheTtl field if non-nil, zero value otherwise.

### GetCacheTtlOk

`func (o *DatabaseInput) GetCacheTtlOk() (*string, bool)`

GetCacheTtlOk returns a tuple with the CacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTtl

`func (o *DatabaseInput) SetCacheTtl(v string)`

SetCacheTtl sets CacheTtl field to given value.

### HasCacheTtl

`func (o *DatabaseInput) HasCacheTtl() bool`

HasCacheTtl returns a boolean if a field has been set.

### GetRefingerprint

`func (o *DatabaseInput) GetRefingerprint() bool`

GetRefingerprint returns the Refingerprint field if non-nil, zero value otherwise.

### GetRefingerprintOk

`func (o *DatabaseInput) GetRefingerprintOk() (*bool, bool)`

GetRefingerprintOk returns a tuple with the Refingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefingerprint

`func (o *DatabaseInput) SetRefingerprint(v bool)`

SetRefingerprint sets Refingerprint field to given value.

### HasRefingerprint

`func (o *DatabaseInput) HasRefingerprint() bool`

HasRefingerprint returns a boolean if a field has been set.

### GetAutoRunQueries

`func (o *DatabaseInput) GetAutoRunQueries() bool`

GetAutoRunQueries returns the AutoRunQueries field if non-nil, zero value otherwise.

### GetAutoRunQueriesOk

`func (o *DatabaseInput) GetAutoRunQueriesOk() (*bool, bool)`

GetAutoRunQueriesOk returns a tuple with the AutoRunQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRunQueries

`func (o *DatabaseInput) SetAutoRunQueries(v bool)`

SetAutoRunQueries sets AutoRunQueries field to given value.

### HasAutoRunQueries

`func (o *DatabaseInput) HasAutoRunQueries() bool`

HasAutoRunQueries returns a boolean if a field has been set.

### GetSchedules

`func (o *DatabaseInput) GetSchedules() DatabaseInputSchedules`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *DatabaseInput) GetSchedulesOk() (*DatabaseInputSchedules, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *DatabaseInput) SetSchedules(v DatabaseInputSchedules)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *DatabaseInput) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### GetDetails

`func (o *DatabaseInput) GetDetails() DatabaseInputDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DatabaseInput) GetDetailsOk() (*DatabaseInputDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DatabaseInput) SetDetails(v DatabaseInputDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *DatabaseInput) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


