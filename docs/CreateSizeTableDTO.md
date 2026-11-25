# CreateSizeTableDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeParameterId** | **int64** | Идентификатор размерной характеристики, к которой относятся размеры из таблицы.  | 
**SizeTableName** | **string** | Название таблицы размеров. Пробелы в начале и в конце удаляются. Название из одних пробелов — ошибка &#x60;INVALID_REQUEST&#x60;.  | 
**Rows** | [**[]SizeTableRowDTO**](SizeTableRowDTO.md) | Строки таблицы: каждая описывает один размер. При сохранении порядок строк может измениться.  | 

## Methods

### NewCreateSizeTableDTO

`func NewCreateSizeTableDTO(sizeParameterId int64, sizeTableName string, rows []SizeTableRowDTO, ) *CreateSizeTableDTO`

NewCreateSizeTableDTO instantiates a new CreateSizeTableDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSizeTableDTOWithDefaults

`func NewCreateSizeTableDTOWithDefaults() *CreateSizeTableDTO`

NewCreateSizeTableDTOWithDefaults instantiates a new CreateSizeTableDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeParameterId

`func (o *CreateSizeTableDTO) GetSizeParameterId() int64`

GetSizeParameterId returns the SizeParameterId field if non-nil, zero value otherwise.

### GetSizeParameterIdOk

`func (o *CreateSizeTableDTO) GetSizeParameterIdOk() (*int64, bool)`

GetSizeParameterIdOk returns a tuple with the SizeParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeParameterId

`func (o *CreateSizeTableDTO) SetSizeParameterId(v int64)`

SetSizeParameterId sets SizeParameterId field to given value.


### GetSizeTableName

`func (o *CreateSizeTableDTO) GetSizeTableName() string`

GetSizeTableName returns the SizeTableName field if non-nil, zero value otherwise.

### GetSizeTableNameOk

`func (o *CreateSizeTableDTO) GetSizeTableNameOk() (*string, bool)`

GetSizeTableNameOk returns a tuple with the SizeTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeTableName

`func (o *CreateSizeTableDTO) SetSizeTableName(v string)`

SetSizeTableName sets SizeTableName field to given value.


### GetRows

`func (o *CreateSizeTableDTO) GetRows() []SizeTableRowDTO`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *CreateSizeTableDTO) GetRowsOk() (*[]SizeTableRowDTO, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *CreateSizeTableDTO) SetRows(v []SizeTableRowDTO)`

SetRows sets Rows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


