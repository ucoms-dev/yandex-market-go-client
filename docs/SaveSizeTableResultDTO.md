# SaveSizeTableResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeTableId** | Pointer to **int64** | Идентификатор сохраненной таблицы размеров. | [optional] 
**SizeTableName** | Pointer to **string** | Название сохраненной таблицы размеров. | [optional] 
**Errors** | Pointer to [**[]SizeTableErrorDTO**](SizeTableErrorDTO.md) | Ошибки, из-за которых таблица не сохранена.  Сначала проверяются структура таблицы, ее колонки и значения ячеек — ошибки &#x60;INVALID_REQUEST&#x60;, &#x60;UNKNOWN_SIZE_CHART&#x60; и &#x60;INVALID_SIZE_VALUE&#x60;. Если они есть, остальные проверки не выполняются, поэтому после исправления могут вернуться другие ошибки.  | [optional] 

## Methods

### NewSaveSizeTableResultDTO

`func NewSaveSizeTableResultDTO() *SaveSizeTableResultDTO`

NewSaveSizeTableResultDTO instantiates a new SaveSizeTableResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveSizeTableResultDTOWithDefaults

`func NewSaveSizeTableResultDTOWithDefaults() *SaveSizeTableResultDTO`

NewSaveSizeTableResultDTOWithDefaults instantiates a new SaveSizeTableResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeTableId

`func (o *SaveSizeTableResultDTO) GetSizeTableId() int64`

GetSizeTableId returns the SizeTableId field if non-nil, zero value otherwise.

### GetSizeTableIdOk

`func (o *SaveSizeTableResultDTO) GetSizeTableIdOk() (*int64, bool)`

GetSizeTableIdOk returns a tuple with the SizeTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeTableId

`func (o *SaveSizeTableResultDTO) SetSizeTableId(v int64)`

SetSizeTableId sets SizeTableId field to given value.

### HasSizeTableId

`func (o *SaveSizeTableResultDTO) HasSizeTableId() bool`

HasSizeTableId returns a boolean if a field has been set.

### GetSizeTableName

`func (o *SaveSizeTableResultDTO) GetSizeTableName() string`

GetSizeTableName returns the SizeTableName field if non-nil, zero value otherwise.

### GetSizeTableNameOk

`func (o *SaveSizeTableResultDTO) GetSizeTableNameOk() (*string, bool)`

GetSizeTableNameOk returns a tuple with the SizeTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeTableName

`func (o *SaveSizeTableResultDTO) SetSizeTableName(v string)`

SetSizeTableName sets SizeTableName field to given value.

### HasSizeTableName

`func (o *SaveSizeTableResultDTO) HasSizeTableName() bool`

HasSizeTableName returns a boolean if a field has been set.

### GetErrors

`func (o *SaveSizeTableResultDTO) GetErrors() []SizeTableErrorDTO`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SaveSizeTableResultDTO) GetErrorsOk() (*[]SizeTableErrorDTO, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SaveSizeTableResultDTO) SetErrors(v []SizeTableErrorDTO)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SaveSizeTableResultDTO) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *SaveSizeTableResultDTO) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *SaveSizeTableResultDTO) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


