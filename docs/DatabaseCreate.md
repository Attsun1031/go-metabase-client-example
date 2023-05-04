# DatabaseCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsOnDemand** | Pointer to **bool** |  | [optional] 
**IsFullSync** | Pointer to **bool** |  | [optional] 
**IsSample** | Pointer to **bool** |  | [optional] 
**CacheTtl** | Pointer to **string** |  | [optional] 
**Refingerprint** | Pointer to **bool** |  | [optional] 
**AutoRunQueries** | Pointer to **bool** |  | [optional] 
**Schedules** | Pointer to [**DatabaseCreateSchedules**](DatabaseCreateSchedules.md) |  | [optional] 
**Details** | Pointer to [**DatabaseCreateDetails**](DatabaseCreateDetails.md) |  | [optional] 

## Methods

### NewDatabaseCreate

`func NewDatabaseCreate() *DatabaseCreate`

NewDatabaseCreate instantiates a new DatabaseCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseCreateWithDefaults

`func NewDatabaseCreateWithDefaults() *DatabaseCreate`

NewDatabaseCreateWithDefaults instantiates a new DatabaseCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsOnDemand

`func (o *DatabaseCreate) GetIsOnDemand() bool`

GetIsOnDemand returns the IsOnDemand field if non-nil, zero value otherwise.

### GetIsOnDemandOk

`func (o *DatabaseCreate) GetIsOnDemandOk() (*bool, bool)`

GetIsOnDemandOk returns a tuple with the IsOnDemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnDemand

`func (o *DatabaseCreate) SetIsOnDemand(v bool)`

SetIsOnDemand sets IsOnDemand field to given value.

### HasIsOnDemand

`func (o *DatabaseCreate) HasIsOnDemand() bool`

HasIsOnDemand returns a boolean if a field has been set.

### GetIsFullSync

`func (o *DatabaseCreate) GetIsFullSync() bool`

GetIsFullSync returns the IsFullSync field if non-nil, zero value otherwise.

### GetIsFullSyncOk

`func (o *DatabaseCreate) GetIsFullSyncOk() (*bool, bool)`

GetIsFullSyncOk returns a tuple with the IsFullSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullSync

`func (o *DatabaseCreate) SetIsFullSync(v bool)`

SetIsFullSync sets IsFullSync field to given value.

### HasIsFullSync

`func (o *DatabaseCreate) HasIsFullSync() bool`

HasIsFullSync returns a boolean if a field has been set.

### GetIsSample

`func (o *DatabaseCreate) GetIsSample() bool`

GetIsSample returns the IsSample field if non-nil, zero value otherwise.

### GetIsSampleOk

`func (o *DatabaseCreate) GetIsSampleOk() (*bool, bool)`

GetIsSampleOk returns a tuple with the IsSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSample

`func (o *DatabaseCreate) SetIsSample(v bool)`

SetIsSample sets IsSample field to given value.

### HasIsSample

`func (o *DatabaseCreate) HasIsSample() bool`

HasIsSample returns a boolean if a field has been set.

### GetCacheTtl

`func (o *DatabaseCreate) GetCacheTtl() string`

GetCacheTtl returns the CacheTtl field if non-nil, zero value otherwise.

### GetCacheTtlOk

`func (o *DatabaseCreate) GetCacheTtlOk() (*string, bool)`

GetCacheTtlOk returns a tuple with the CacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTtl

`func (o *DatabaseCreate) SetCacheTtl(v string)`

SetCacheTtl sets CacheTtl field to given value.

### HasCacheTtl

`func (o *DatabaseCreate) HasCacheTtl() bool`

HasCacheTtl returns a boolean if a field has been set.

### GetRefingerprint

`func (o *DatabaseCreate) GetRefingerprint() bool`

GetRefingerprint returns the Refingerprint field if non-nil, zero value otherwise.

### GetRefingerprintOk

`func (o *DatabaseCreate) GetRefingerprintOk() (*bool, bool)`

GetRefingerprintOk returns a tuple with the Refingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefingerprint

`func (o *DatabaseCreate) SetRefingerprint(v bool)`

SetRefingerprint sets Refingerprint field to given value.

### HasRefingerprint

`func (o *DatabaseCreate) HasRefingerprint() bool`

HasRefingerprint returns a boolean if a field has been set.

### GetAutoRunQueries

`func (o *DatabaseCreate) GetAutoRunQueries() bool`

GetAutoRunQueries returns the AutoRunQueries field if non-nil, zero value otherwise.

### GetAutoRunQueriesOk

`func (o *DatabaseCreate) GetAutoRunQueriesOk() (*bool, bool)`

GetAutoRunQueriesOk returns a tuple with the AutoRunQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRunQueries

`func (o *DatabaseCreate) SetAutoRunQueries(v bool)`

SetAutoRunQueries sets AutoRunQueries field to given value.

### HasAutoRunQueries

`func (o *DatabaseCreate) HasAutoRunQueries() bool`

HasAutoRunQueries returns a boolean if a field has been set.

### GetSchedules

`func (o *DatabaseCreate) GetSchedules() DatabaseCreateSchedules`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *DatabaseCreate) GetSchedulesOk() (*DatabaseCreateSchedules, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *DatabaseCreate) SetSchedules(v DatabaseCreateSchedules)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *DatabaseCreate) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### GetDetails

`func (o *DatabaseCreate) GetDetails() DatabaseCreateDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DatabaseCreate) GetDetailsOk() (*DatabaseCreateDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DatabaseCreate) SetDetails(v DatabaseCreateDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *DatabaseCreate) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


