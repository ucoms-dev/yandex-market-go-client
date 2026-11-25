# UpdateSizeTableDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeTableId** | **int64** | Идентификатор таблицы размеров, созданной бизнесом. | 
**SizeTableName** | **string** | Новое название таблицы размеров. Пробелы в начале и в конце удаляются. Название из одних пробелов — ошибка &#x60;INVALID_REQUEST&#x60;.  | 
**Rows** | [**[]SizeTableRowDTO**](SizeTableRowDTO.md) | Все строки таблицы: каждая описывает один размер. Строки и колонки, которых нет в запросе, удаляются. При сохранении порядок строк может измениться.  | 

## Methods

### NewUpdateSizeTableDTO

`func NewUpdateSizeTableDTO(sizeTableId int64, sizeTableName string, rows []SizeTableRowDTO, ) *UpdateSizeTableDTO`

NewUpdateSizeTableDTO instantiates a new UpdateSizeTableDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSizeTableDTOWithDefaults

`func NewUpdateSizeTableDTOWithDefaults() *UpdateSizeTableDTO`

NewUpdateSizeTableDTOWithDefaults instantiates a new UpdateSizeTableDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeTableId

`func (o *UpdateSizeTableDTO) GetSizeTableId() int64`

GetSizeTableId returns the SizeTableId field if non-nil, zero value otherwise.

### GetSizeTableIdOk

`func (o *UpdateSizeTableDTO) GetSizeTableIdOk() (*int64, bool)`

GetSizeTableIdOk returns a tuple with the SizeTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeTableId

`func (o *UpdateSizeTableDTO) SetSizeTableId(v int64)`

SetSizeTableId sets SizeTableId field to given value.


### GetSizeTableName

`func (o *UpdateSizeTableDTO) GetSizeTableName() string`

GetSizeTableName returns the SizeTableName field if non-nil, zero value otherwise.

### GetSizeTableNameOk

`func (o *UpdateSizeTableDTO) GetSizeTableNameOk() (*string, bool)`

GetSizeTableNameOk returns a tuple with the SizeTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeTableName

`func (o *UpdateSizeTableDTO) SetSizeTableName(v string)`

SetSizeTableName sets SizeTableName field to given value.


### GetRows

`func (o *UpdateSizeTableDTO) GetRows() []SizeTableRowDTO`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *UpdateSizeTableDTO) GetRowsOk() (*[]SizeTableRowDTO, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *UpdateSizeTableDTO) SetRows(v []SizeTableRowDTO)`

SetRows sets Rows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


