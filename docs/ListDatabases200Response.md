# ListDatabases200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Database**](Database.md) |  | [optional] 

## Methods

### NewListDatabases200Response

`func NewListDatabases200Response() *ListDatabases200Response`

NewListDatabases200Response instantiates a new ListDatabases200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDatabases200ResponseWithDefaults

`func NewListDatabases200ResponseWithDefaults() *ListDatabases200Response`

NewListDatabases200ResponseWithDefaults instantiates a new ListDatabases200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListDatabases200Response) GetData() []Database`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListDatabases200Response) GetDataOk() (*[]Database, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListDatabases200Response) SetData(v []Database)`

SetData sets Data field to given value.

### HasData

`func (o *ListDatabases200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


