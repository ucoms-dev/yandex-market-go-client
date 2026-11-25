# GetSizeTablesResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeTables** | [**[]SizeTableDTO**](SizeTableDTO.md) | Таблицы размеров: сначала созданные бизнесом, затем брендовые, затем стандартные; внутри каждой группы — по убыванию идентификатора.  | 
**NotFoundSizeTableIds** | **[]int64** | Идентификаторы из запроса, по которым таблицу вернуть не удалось: такой таблицы нет или она принадлежит другому бизнесу.  Заполняется только на первой странице. На следующих страницах и в запросе без &#x60;sizeTableIds&#x60; список пустой.  | 
**Paging** | [**PackagingForwardScrollingPagerDTO**](PackagingForwardScrollingPagerDTO.md) |  | 

## Methods

### NewGetSizeTablesResultDTO

`func NewGetSizeTablesResultDTO(sizeTables []SizeTableDTO, notFoundSizeTableIds []int64, paging PackagingForwardScrollingPagerDTO, ) *GetSizeTablesResultDTO`

NewGetSizeTablesResultDTO instantiates a new GetSizeTablesResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSizeTablesResultDTOWithDefaults

`func NewGetSizeTablesResultDTOWithDefaults() *GetSizeTablesResultDTO`

NewGetSizeTablesResultDTOWithDefaults instantiates a new GetSizeTablesResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeTables

`func (o *GetSizeTablesResultDTO) GetSizeTables() []SizeTableDTO`

GetSizeTables returns the SizeTables field if non-nil, zero value otherwise.

### GetSizeTablesOk

`func (o *GetSizeTablesResultDTO) GetSizeTablesOk() (*[]SizeTableDTO, bool)`

GetSizeTablesOk returns a tuple with the SizeTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeTables

`func (o *GetSizeTablesResultDTO) SetSizeTables(v []SizeTableDTO)`

SetSizeTables sets SizeTables field to given value.


### GetNotFoundSizeTableIds

`func (o *GetSizeTablesResultDTO) GetNotFoundSizeTableIds() []int64`

GetNotFoundSizeTableIds returns the NotFoundSizeTableIds field if non-nil, zero value otherwise.

### GetNotFoundSizeTableIdsOk

`func (o *GetSizeTablesResultDTO) GetNotFoundSizeTableIdsOk() (*[]int64, bool)`

GetNotFoundSizeTableIdsOk returns a tuple with the NotFoundSizeTableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotFoundSizeTableIds

`func (o *GetSizeTablesResultDTO) SetNotFoundSizeTableIds(v []int64)`

SetNotFoundSizeTableIds sets NotFoundSizeTableIds field to given value.


### GetPaging

`func (o *GetSizeTablesResultDTO) GetPaging() PackagingForwardScrollingPagerDTO`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *GetSizeTablesResultDTO) GetPagingOk() (*PackagingForwardScrollingPagerDTO, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *GetSizeTablesResultDTO) SetPaging(v PackagingForwardScrollingPagerDTO)`

SetPaging sets Paging field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


