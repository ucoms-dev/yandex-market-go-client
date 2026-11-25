# SizeTableDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeTableId** | **int64** | Идентификатор таблицы размеров. | 
**SizeTableName** | **string** | Название таблицы размеров. | 
**SizeTableType** | [**SizeTableType**](SizeTableType.md) |  | 
**SizeParameterId** | **int64** | Идентификатор размерной характеристики, к которой относятся размеры из таблицы. | 
**Editable** | **bool** | Может ли бизнес редактировать таблицу: &#x60;true&#x60; — для таблиц, созданных бизнесом, &#x60;false&#x60; — для таблиц Маркета.  | 
**Key** | Pointer to [**SizeTableKeyDTO**](SizeTableKeyDTO.md) |  | [optional] 
**SizeCharts** | [**[]SizeChartDTO**](SizeChartDTO.md) | Колонки таблицы: размерные сетки и обмеры. Колонки без значений не возвращаются.  | 

## Methods

### NewSizeTableDTO

`func NewSizeTableDTO(sizeTableId int64, sizeTableName string, sizeTableType SizeTableType, sizeParameterId int64, editable bool, sizeCharts []SizeChartDTO, ) *SizeTableDTO`

NewSizeTableDTO instantiates a new SizeTableDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSizeTableDTOWithDefaults

`func NewSizeTableDTOWithDefaults() *SizeTableDTO`

NewSizeTableDTOWithDefaults instantiates a new SizeTableDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeTableId

`func (o *SizeTableDTO) GetSizeTableId() int64`

GetSizeTableId returns the SizeTableId field if non-nil, zero value otherwise.

### GetSizeTableIdOk

`func (o *SizeTableDTO) GetSizeTableIdOk() (*int64, bool)`

GetSizeTableIdOk returns a tuple with the SizeTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeTableId

`func (o *SizeTableDTO) SetSizeTableId(v int64)`

SetSizeTableId sets SizeTableId field to given value.


### GetSizeTableName

`func (o *SizeTableDTO) GetSizeTableName() string`

GetSizeTableName returns the SizeTableName field if non-nil, zero value otherwise.

### GetSizeTableNameOk

`func (o *SizeTableDTO) GetSizeTableNameOk() (*string, bool)`

GetSizeTableNameOk returns a tuple with the SizeTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeTableName

`func (o *SizeTableDTO) SetSizeTableName(v string)`

SetSizeTableName sets SizeTableName field to given value.


### GetSizeTableType

`func (o *SizeTableDTO) GetSizeTableType() SizeTableType`

GetSizeTableType returns the SizeTableType field if non-nil, zero value otherwise.

### GetSizeTableTypeOk

`func (o *SizeTableDTO) GetSizeTableTypeOk() (*SizeTableType, bool)`

GetSizeTableTypeOk returns a tuple with the SizeTableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeTableType

`func (o *SizeTableDTO) SetSizeTableType(v SizeTableType)`

SetSizeTableType sets SizeTableType field to given value.


### GetSizeParameterId

`func (o *SizeTableDTO) GetSizeParameterId() int64`

GetSizeParameterId returns the SizeParameterId field if non-nil, zero value otherwise.

### GetSizeParameterIdOk

`func (o *SizeTableDTO) GetSizeParameterIdOk() (*int64, bool)`

GetSizeParameterIdOk returns a tuple with the SizeParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeParameterId

`func (o *SizeTableDTO) SetSizeParameterId(v int64)`

SetSizeParameterId sets SizeParameterId field to given value.


### GetEditable

`func (o *SizeTableDTO) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *SizeTableDTO) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *SizeTableDTO) SetEditable(v bool)`

SetEditable sets Editable field to given value.


### GetKey

`func (o *SizeTableDTO) GetKey() SizeTableKeyDTO`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SizeTableDTO) GetKeyOk() (*SizeTableKeyDTO, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SizeTableDTO) SetKey(v SizeTableKeyDTO)`

SetKey sets Key field to given value.

### HasKey

`func (o *SizeTableDTO) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSizeCharts

`func (o *SizeTableDTO) GetSizeCharts() []SizeChartDTO`

GetSizeCharts returns the SizeCharts field if non-nil, zero value otherwise.

### GetSizeChartsOk

`func (o *SizeTableDTO) GetSizeChartsOk() (*[]SizeChartDTO, bool)`

GetSizeChartsOk returns a tuple with the SizeCharts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeCharts

`func (o *SizeTableDTO) SetSizeCharts(v []SizeChartDTO)`

SetSizeCharts sets SizeCharts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


