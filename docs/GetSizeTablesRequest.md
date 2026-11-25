# GetSizeTablesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeTableIds** | Pointer to **[]int64** | Идентификаторы таблиц размеров.  Если параметр не передан, возвращаются все таблицы, созданные бизнесом.  | [optional] 

## Methods

### NewGetSizeTablesRequest

`func NewGetSizeTablesRequest() *GetSizeTablesRequest`

NewGetSizeTablesRequest instantiates a new GetSizeTablesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSizeTablesRequestWithDefaults

`func NewGetSizeTablesRequestWithDefaults() *GetSizeTablesRequest`

NewGetSizeTablesRequestWithDefaults instantiates a new GetSizeTablesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeTableIds

`func (o *GetSizeTablesRequest) GetSizeTableIds() []int64`

GetSizeTableIds returns the SizeTableIds field if non-nil, zero value otherwise.

### GetSizeTableIdsOk

`func (o *GetSizeTablesRequest) GetSizeTableIdsOk() (*[]int64, bool)`

GetSizeTableIdsOk returns a tuple with the SizeTableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeTableIds

`func (o *GetSizeTablesRequest) SetSizeTableIds(v []int64)`

SetSizeTableIds sets SizeTableIds field to given value.

### HasSizeTableIds

`func (o *GetSizeTablesRequest) HasSizeTableIds() bool`

HasSizeTableIds returns a boolean if a field has been set.

### SetSizeTableIdsNil

`func (o *GetSizeTablesRequest) SetSizeTableIdsNil(b bool)`

 SetSizeTableIdsNil sets the value for SizeTableIds to be an explicit nil

### UnsetSizeTableIds
`func (o *GetSizeTablesRequest) UnsetSizeTableIds()`

UnsetSizeTableIds ensures that no value is present for SizeTableIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


