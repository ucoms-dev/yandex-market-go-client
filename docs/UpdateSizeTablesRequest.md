# UpdateSizeTablesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeTables** | [**[]UpdateSizeTableDTO**](UpdateSizeTableDTO.md) | Таблицы размеров с новым содержимым. Обрабатываются по порядку, каждая независимо от остальных.  Идентификаторы таблиц не должны повторяться в одном запросе. Если идентификаторы повторяются, запрос не обрабатывается.  Суммарно во всех таблицах — не больше 5000 ячеек. Если ячеек больше, запрос не обрабатывается.  | 

## Methods

### NewUpdateSizeTablesRequest

`func NewUpdateSizeTablesRequest(sizeTables []UpdateSizeTableDTO, ) *UpdateSizeTablesRequest`

NewUpdateSizeTablesRequest instantiates a new UpdateSizeTablesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSizeTablesRequestWithDefaults

`func NewUpdateSizeTablesRequestWithDefaults() *UpdateSizeTablesRequest`

NewUpdateSizeTablesRequestWithDefaults instantiates a new UpdateSizeTablesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeTables

`func (o *UpdateSizeTablesRequest) GetSizeTables() []UpdateSizeTableDTO`

GetSizeTables returns the SizeTables field if non-nil, zero value otherwise.

### GetSizeTablesOk

`func (o *UpdateSizeTablesRequest) GetSizeTablesOk() (*[]UpdateSizeTableDTO, bool)`

GetSizeTablesOk returns a tuple with the SizeTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeTables

`func (o *UpdateSizeTablesRequest) SetSizeTables(v []UpdateSizeTableDTO)`

SetSizeTables sets SizeTables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


