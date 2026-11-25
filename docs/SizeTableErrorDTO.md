# SizeTableErrorDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | [**SizeTableErrorCodeType**](SizeTableErrorCodeType.md) |  | 
**RowIndex** | Pointer to **int32** | Номер строки с ошибкой в &#x60;rows&#x60; запроса, начиная с 0. Может не совпадать с &#x60;rowIndex&#x60; сохраненной таблицы.  | [optional] 
**SizeChartId** | Pointer to **int64** | Идентификатор размерной сетки или обмера, в колонке которого ошибка. | [optional] 
**SizeChartIds** | Pointer to **[]int64** | Идентификаторы размерных сеток и обмеров, в колонках которых ошибка. | [optional] 
**SizeChartNames** | Pointer to **[]string** | Названия колонок из &#x60;sizeChartName&#x60; запроса, которым не нашлось ровно одной размерной сетки или обмера.  | [optional] 
**RowIndices** | Pointer to **[]int32** | Номера совпадающих строк в &#x60;rows&#x60; запроса, начиная с 0. | [optional] 

## Methods

### NewSizeTableErrorDTO

`func NewSizeTableErrorDTO(code SizeTableErrorCodeType, ) *SizeTableErrorDTO`

NewSizeTableErrorDTO instantiates a new SizeTableErrorDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSizeTableErrorDTOWithDefaults

`func NewSizeTableErrorDTOWithDefaults() *SizeTableErrorDTO`

NewSizeTableErrorDTOWithDefaults instantiates a new SizeTableErrorDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SizeTableErrorDTO) GetCode() SizeTableErrorCodeType`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SizeTableErrorDTO) GetCodeOk() (*SizeTableErrorCodeType, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SizeTableErrorDTO) SetCode(v SizeTableErrorCodeType)`

SetCode sets Code field to given value.


### GetRowIndex

`func (o *SizeTableErrorDTO) GetRowIndex() int32`

GetRowIndex returns the RowIndex field if non-nil, zero value otherwise.

### GetRowIndexOk

`func (o *SizeTableErrorDTO) GetRowIndexOk() (*int32, bool)`

GetRowIndexOk returns a tuple with the RowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowIndex

`func (o *SizeTableErrorDTO) SetRowIndex(v int32)`

SetRowIndex sets RowIndex field to given value.

### HasRowIndex

`func (o *SizeTableErrorDTO) HasRowIndex() bool`

HasRowIndex returns a boolean if a field has been set.

### GetSizeChartId

`func (o *SizeTableErrorDTO) GetSizeChartId() int64`

GetSizeChartId returns the SizeChartId field if non-nil, zero value otherwise.

### GetSizeChartIdOk

`func (o *SizeTableErrorDTO) GetSizeChartIdOk() (*int64, bool)`

GetSizeChartIdOk returns a tuple with the SizeChartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeChartId

`func (o *SizeTableErrorDTO) SetSizeChartId(v int64)`

SetSizeChartId sets SizeChartId field to given value.

### HasSizeChartId

`func (o *SizeTableErrorDTO) HasSizeChartId() bool`

HasSizeChartId returns a boolean if a field has been set.

### GetSizeChartIds

`func (o *SizeTableErrorDTO) GetSizeChartIds() []int64`

GetSizeChartIds returns the SizeChartIds field if non-nil, zero value otherwise.

### GetSizeChartIdsOk

`func (o *SizeTableErrorDTO) GetSizeChartIdsOk() (*[]int64, bool)`

GetSizeChartIdsOk returns a tuple with the SizeChartIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeChartIds

`func (o *SizeTableErrorDTO) SetSizeChartIds(v []int64)`

SetSizeChartIds sets SizeChartIds field to given value.

### HasSizeChartIds

`func (o *SizeTableErrorDTO) HasSizeChartIds() bool`

HasSizeChartIds returns a boolean if a field has been set.

### SetSizeChartIdsNil

`func (o *SizeTableErrorDTO) SetSizeChartIdsNil(b bool)`

 SetSizeChartIdsNil sets the value for SizeChartIds to be an explicit nil

### UnsetSizeChartIds
`func (o *SizeTableErrorDTO) UnsetSizeChartIds()`

UnsetSizeChartIds ensures that no value is present for SizeChartIds, not even an explicit nil
### GetSizeChartNames

`func (o *SizeTableErrorDTO) GetSizeChartNames() []string`

GetSizeChartNames returns the SizeChartNames field if non-nil, zero value otherwise.

### GetSizeChartNamesOk

`func (o *SizeTableErrorDTO) GetSizeChartNamesOk() (*[]string, bool)`

GetSizeChartNamesOk returns a tuple with the SizeChartNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeChartNames

`func (o *SizeTableErrorDTO) SetSizeChartNames(v []string)`

SetSizeChartNames sets SizeChartNames field to given value.

### HasSizeChartNames

`func (o *SizeTableErrorDTO) HasSizeChartNames() bool`

HasSizeChartNames returns a boolean if a field has been set.

### SetSizeChartNamesNil

`func (o *SizeTableErrorDTO) SetSizeChartNamesNil(b bool)`

 SetSizeChartNamesNil sets the value for SizeChartNames to be an explicit nil

### UnsetSizeChartNames
`func (o *SizeTableErrorDTO) UnsetSizeChartNames()`

UnsetSizeChartNames ensures that no value is present for SizeChartNames, not even an explicit nil
### GetRowIndices

`func (o *SizeTableErrorDTO) GetRowIndices() []int32`

GetRowIndices returns the RowIndices field if non-nil, zero value otherwise.

### GetRowIndicesOk

`func (o *SizeTableErrorDTO) GetRowIndicesOk() (*[]int32, bool)`

GetRowIndicesOk returns a tuple with the RowIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowIndices

`func (o *SizeTableErrorDTO) SetRowIndices(v []int32)`

SetRowIndices sets RowIndices field to given value.

### HasRowIndices

`func (o *SizeTableErrorDTO) HasRowIndices() bool`

HasRowIndices returns a boolean if a field has been set.

### SetRowIndicesNil

`func (o *SizeTableErrorDTO) SetRowIndicesNil(b bool)`

 SetRowIndicesNil sets the value for RowIndices to be an explicit nil

### UnsetRowIndices
`func (o *SizeTableErrorDTO) UnsetRowIndices()`

UnsetRowIndices ensures that no value is present for RowIndices, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


