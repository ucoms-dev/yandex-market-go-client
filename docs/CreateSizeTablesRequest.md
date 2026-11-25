# CreateSizeTablesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeTables** | [**[]CreateSizeTableDTO**](CreateSizeTableDTO.md) | Таблицы размеров. Обрабатываются по порядку, каждая независимо от остальных.  Суммарно во всех таблицах — не больше 5000 ячеек. Если ячеек больше, запрос не обрабатывается.  | 

## Methods

### NewCreateSizeTablesRequest

`func NewCreateSizeTablesRequest(sizeTables []CreateSizeTableDTO, ) *CreateSizeTablesRequest`

NewCreateSizeTablesRequest instantiates a new CreateSizeTablesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSizeTablesRequestWithDefaults

`func NewCreateSizeTablesRequestWithDefaults() *CreateSizeTablesRequest`

NewCreateSizeTablesRequestWithDefaults instantiates a new CreateSizeTablesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeTables

`func (o *CreateSizeTablesRequest) GetSizeTables() []CreateSizeTableDTO`

GetSizeTables returns the SizeTables field if non-nil, zero value otherwise.

### GetSizeTablesOk

`func (o *CreateSizeTablesRequest) GetSizeTablesOk() (*[]CreateSizeTableDTO, bool)`

GetSizeTablesOk returns a tuple with the SizeTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeTables

`func (o *CreateSizeTablesRequest) SetSizeTables(v []CreateSizeTableDTO)`

SetSizeTables sets SizeTables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


