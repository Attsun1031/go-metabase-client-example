# GroupWithMembers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**MemberCount** | Pointer to **float32** |  | [optional] 
**Members** | Pointer to [**[]GroupMember**](GroupMember.md) |  | [optional] 

## Methods

### NewGroupWithMembers

`func NewGroupWithMembers() *GroupWithMembers`

NewGroupWithMembers instantiates a new GroupWithMembers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithMembersWithDefaults

`func NewGroupWithMembersWithDefaults() *GroupWithMembers`

NewGroupWithMembersWithDefaults instantiates a new GroupWithMembers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupWithMembers) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupWithMembers) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupWithMembers) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GroupWithMembers) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GroupWithMembers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupWithMembers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupWithMembers) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupWithMembers) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMemberCount

`func (o *GroupWithMembers) GetMemberCount() float32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *GroupWithMembers) GetMemberCountOk() (*float32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *GroupWithMembers) SetMemberCount(v float32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *GroupWithMembers) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetMembers

`func (o *GroupWithMembers) GetMembers() []GroupMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GroupWithMembers) GetMembersOk() (*[]GroupMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GroupWithMembers) SetMembers(v []GroupMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GroupWithMembers) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


