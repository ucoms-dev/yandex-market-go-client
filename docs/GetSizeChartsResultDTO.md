# GetSizeChartsResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeParameterId** | **int64** | Идентификатор размерной характеристики из запроса. | 
**SizeCharts** | [**[]SizeChartInfoDTO**](SizeChartInfoDTO.md) | Размерные сетки и обмеры, которые можно использовать в таблицах размеров: сначала сетки, затем обмеры.  | 

## Methods

### NewGetSizeChartsResultDTO

`func NewGetSizeChartsResultDTO(sizeParameterId int64, sizeCharts []SizeChartInfoDTO, ) *GetSizeChartsResultDTO`

NewGetSizeChartsResultDTO instantiates a new GetSizeChartsResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSizeChartsResultDTOWithDefaults

`func NewGetSizeChartsResultDTOWithDefaults() *GetSizeChartsResultDTO`

NewGetSizeChartsResultDTOWithDefaults instantiates a new GetSizeChartsResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeParameterId

`func (o *GetSizeChartsResultDTO) GetSizeParameterId() int64`

GetSizeParameterId returns the SizeParameterId field if non-nil, zero value otherwise.

### GetSizeParameterIdOk

`func (o *GetSizeChartsResultDTO) GetSizeParameterIdOk() (*int64, bool)`

GetSizeParameterIdOk returns a tuple with the SizeParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeParameterId

`func (o *GetSizeChartsResultDTO) SetSizeParameterId(v int64)`

SetSizeParameterId sets SizeParameterId field to given value.


### GetSizeCharts

`func (o *GetSizeChartsResultDTO) GetSizeCharts() []SizeChartInfoDTO`

GetSizeCharts returns the SizeCharts field if non-nil, zero value otherwise.

### GetSizeChartsOk

`func (o *GetSizeChartsResultDTO) GetSizeChartsOk() (*[]SizeChartInfoDTO, bool)`

GetSizeChartsOk returns a tuple with the SizeCharts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeCharts

`func (o *GetSizeChartsResultDTO) SetSizeCharts(v []SizeChartInfoDTO)`

SetSizeCharts sets SizeCharts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


