# SizeChartInfoDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeChartId** | **int64** | Идентификатор размерной сетки или обмера. Передавайте его в &#x60;sizeChartId&#x60; ячейки таблицы. | 
**SizeChartName** | **string** | Название размерной сетки или обмера. | 
**ColumnType** | [**SizeTableColumnType**](SizeTableColumnType.md) |  | 
**Required** | **bool** | Обязательная ли размерная сетка. В таблице должна быть хотя бы одна обязательная сетка, а каждую обязательную сетку в таблице нужно заполнить во всех строках. У обмеров всегда &#x60;false&#x60;.  | 
**SizeValues** | [**[]SizeChartValueDTO**](SizeChartValueDTO.md) | Справочник значений размерной сетки. Заполнен только у стандартных размерных сеток Маркета: в ячейке такой колонки передавайте значение из справочника.  Если список пустой, значение задается текстом.  | 

## Methods

### NewSizeChartInfoDTO

`func NewSizeChartInfoDTO(sizeChartId int64, sizeChartName string, columnType SizeTableColumnType, required bool, sizeValues []SizeChartValueDTO, ) *SizeChartInfoDTO`

NewSizeChartInfoDTO instantiates a new SizeChartInfoDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSizeChartInfoDTOWithDefaults

`func NewSizeChartInfoDTOWithDefaults() *SizeChartInfoDTO`

NewSizeChartInfoDTOWithDefaults instantiates a new SizeChartInfoDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeChartId

`func (o *SizeChartInfoDTO) GetSizeChartId() int64`

GetSizeChartId returns the SizeChartId field if non-nil, zero value otherwise.

### GetSizeChartIdOk

`func (o *SizeChartInfoDTO) GetSizeChartIdOk() (*int64, bool)`

GetSizeChartIdOk returns a tuple with the SizeChartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeChartId

`func (o *SizeChartInfoDTO) SetSizeChartId(v int64)`

SetSizeChartId sets SizeChartId field to given value.


### GetSizeChartName

`func (o *SizeChartInfoDTO) GetSizeChartName() string`

GetSizeChartName returns the SizeChartName field if non-nil, zero value otherwise.

### GetSizeChartNameOk

`func (o *SizeChartInfoDTO) GetSizeChartNameOk() (*string, bool)`

GetSizeChartNameOk returns a tuple with the SizeChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeChartName

`func (o *SizeChartInfoDTO) SetSizeChartName(v string)`

SetSizeChartName sets SizeChartName field to given value.


### GetColumnType

`func (o *SizeChartInfoDTO) GetColumnType() SizeTableColumnType`

GetColumnType returns the ColumnType field if non-nil, zero value otherwise.

### GetColumnTypeOk

`func (o *SizeChartInfoDTO) GetColumnTypeOk() (*SizeTableColumnType, bool)`

GetColumnTypeOk returns a tuple with the ColumnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnType

`func (o *SizeChartInfoDTO) SetColumnType(v SizeTableColumnType)`

SetColumnType sets ColumnType field to given value.


### GetRequired

`func (o *SizeChartInfoDTO) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *SizeChartInfoDTO) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *SizeChartInfoDTO) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetSizeValues

`func (o *SizeChartInfoDTO) GetSizeValues() []SizeChartValueDTO`

GetSizeValues returns the SizeValues field if non-nil, zero value otherwise.

### GetSizeValuesOk

`func (o *SizeChartInfoDTO) GetSizeValuesOk() (*[]SizeChartValueDTO, bool)`

GetSizeValuesOk returns a tuple with the SizeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeValues

`func (o *SizeChartInfoDTO) SetSizeValues(v []SizeChartValueDTO)`

SetSizeValues sets SizeValues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


