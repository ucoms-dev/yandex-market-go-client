# SizeChartDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeChartId** | **int64** | Идентификатор размерной сетки или обмера. У сеток и обмеров общее пространство идентификаторов. | 
**SizeChartName** | **string** | Название размерной сетки или обмера. | 
**ColumnType** | [**SizeTableColumnType**](SizeTableColumnType.md) |  | 
**Sizes** | [**[]SizeValueDTO**](SizeValueDTO.md) | Значения колонки в порядке строк таблицы. Пустые ячейки не возвращаются. | 

## Methods

### NewSizeChartDTO

`func NewSizeChartDTO(sizeChartId int64, sizeChartName string, columnType SizeTableColumnType, sizes []SizeValueDTO, ) *SizeChartDTO`

NewSizeChartDTO instantiates a new SizeChartDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSizeChartDTOWithDefaults

`func NewSizeChartDTOWithDefaults() *SizeChartDTO`

NewSizeChartDTOWithDefaults instantiates a new SizeChartDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeChartId

`func (o *SizeChartDTO) GetSizeChartId() int64`

GetSizeChartId returns the SizeChartId field if non-nil, zero value otherwise.

### GetSizeChartIdOk

`func (o *SizeChartDTO) GetSizeChartIdOk() (*int64, bool)`

GetSizeChartIdOk returns a tuple with the SizeChartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeChartId

`func (o *SizeChartDTO) SetSizeChartId(v int64)`

SetSizeChartId sets SizeChartId field to given value.


### GetSizeChartName

`func (o *SizeChartDTO) GetSizeChartName() string`

GetSizeChartName returns the SizeChartName field if non-nil, zero value otherwise.

### GetSizeChartNameOk

`func (o *SizeChartDTO) GetSizeChartNameOk() (*string, bool)`

GetSizeChartNameOk returns a tuple with the SizeChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeChartName

`func (o *SizeChartDTO) SetSizeChartName(v string)`

SetSizeChartName sets SizeChartName field to given value.


### GetColumnType

`func (o *SizeChartDTO) GetColumnType() SizeTableColumnType`

GetColumnType returns the ColumnType field if non-nil, zero value otherwise.

### GetColumnTypeOk

`func (o *SizeChartDTO) GetColumnTypeOk() (*SizeTableColumnType, bool)`

GetColumnTypeOk returns a tuple with the ColumnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnType

`func (o *SizeChartDTO) SetColumnType(v SizeTableColumnType)`

SetColumnType sets ColumnType field to given value.


### GetSizes

`func (o *SizeChartDTO) GetSizes() []SizeValueDTO`

GetSizes returns the Sizes field if non-nil, zero value otherwise.

### GetSizesOk

`func (o *SizeChartDTO) GetSizesOk() (*[]SizeValueDTO, bool)`

GetSizesOk returns a tuple with the Sizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizes

`func (o *SizeChartDTO) SetSizes(v []SizeValueDTO)`

SetSizes sets Sizes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


