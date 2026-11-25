# SuggestSizeTablesResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeParameterId** | Pointer to **int64** | Идентификатор размерной характеристики категории. Отсутствует, если у категории нет размерной характеристики — тогда и таблиц нет.  | [optional] 
**SizeTables** | [**[]SizeTableDTO**](SizeTableDTO.md) | Подходящие таблицы размеров: сначала созданные бизнесом, затем брендовые, затем стандартные; внутри каждой группы — по убыванию идентификатора.  | 
**Paging** | [**PackagingForwardScrollingPagerDTO**](PackagingForwardScrollingPagerDTO.md) |  | 

## Methods

### NewSuggestSizeTablesResultDTO

`func NewSuggestSizeTablesResultDTO(sizeTables []SizeTableDTO, paging PackagingForwardScrollingPagerDTO, ) *SuggestSizeTablesResultDTO`

NewSuggestSizeTablesResultDTO instantiates a new SuggestSizeTablesResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestSizeTablesResultDTOWithDefaults

`func NewSuggestSizeTablesResultDTOWithDefaults() *SuggestSizeTablesResultDTO`

NewSuggestSizeTablesResultDTOWithDefaults instantiates a new SuggestSizeTablesResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeParameterId

`func (o *SuggestSizeTablesResultDTO) GetSizeParameterId() int64`

GetSizeParameterId returns the SizeParameterId field if non-nil, zero value otherwise.

### GetSizeParameterIdOk

`func (o *SuggestSizeTablesResultDTO) GetSizeParameterIdOk() (*int64, bool)`

GetSizeParameterIdOk returns a tuple with the SizeParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeParameterId

`func (o *SuggestSizeTablesResultDTO) SetSizeParameterId(v int64)`

SetSizeParameterId sets SizeParameterId field to given value.

### HasSizeParameterId

`func (o *SuggestSizeTablesResultDTO) HasSizeParameterId() bool`

HasSizeParameterId returns a boolean if a field has been set.

### GetSizeTables

`func (o *SuggestSizeTablesResultDTO) GetSizeTables() []SizeTableDTO`

GetSizeTables returns the SizeTables field if non-nil, zero value otherwise.

### GetSizeTablesOk

`func (o *SuggestSizeTablesResultDTO) GetSizeTablesOk() (*[]SizeTableDTO, bool)`

GetSizeTablesOk returns a tuple with the SizeTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeTables

`func (o *SuggestSizeTablesResultDTO) SetSizeTables(v []SizeTableDTO)`

SetSizeTables sets SizeTables field to given value.


### GetPaging

`func (o *SuggestSizeTablesResultDTO) GetPaging() PackagingForwardScrollingPagerDTO`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *SuggestSizeTablesResultDTO) GetPagingOk() (*PackagingForwardScrollingPagerDTO, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *SuggestSizeTablesResultDTO) SetPaging(v PackagingForwardScrollingPagerDTO)`

SetPaging sets Paging field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


