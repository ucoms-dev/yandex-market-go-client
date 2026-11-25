# SizeValueDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeValueId** | Pointer to **int64** | Идентификатор значения размера. Возвращается, только если у колонки есть справочник значений — это стандартные размерные сетки Маркета.  В сетках, созданных бизнесом, брендовых сетках и обмерах значение задается только текстом.  | [optional] 
**SizeValue** | **string** | Значение размера или обмера. | 
**RowIndex** | **int32** | Номер строки таблицы, начиная с 0. Значения с одинаковым &#x60;rowIndex&#x60; в разных колонках описывают один размер.  Номера могут идти с пропусками, а порядок строк может не совпадать с порядком размеров. Собирайте строки по значению &#x60;rowIndex&#x60;, а не по позиции значения в &#x60;sizes&#x60;.  | 

## Methods

### NewSizeValueDTO

`func NewSizeValueDTO(sizeValue string, rowIndex int32, ) *SizeValueDTO`

NewSizeValueDTO instantiates a new SizeValueDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSizeValueDTOWithDefaults

`func NewSizeValueDTOWithDefaults() *SizeValueDTO`

NewSizeValueDTOWithDefaults instantiates a new SizeValueDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeValueId

`func (o *SizeValueDTO) GetSizeValueId() int64`

GetSizeValueId returns the SizeValueId field if non-nil, zero value otherwise.

### GetSizeValueIdOk

`func (o *SizeValueDTO) GetSizeValueIdOk() (*int64, bool)`

GetSizeValueIdOk returns a tuple with the SizeValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeValueId

`func (o *SizeValueDTO) SetSizeValueId(v int64)`

SetSizeValueId sets SizeValueId field to given value.

### HasSizeValueId

`func (o *SizeValueDTO) HasSizeValueId() bool`

HasSizeValueId returns a boolean if a field has been set.

### GetSizeValue

`func (o *SizeValueDTO) GetSizeValue() string`

GetSizeValue returns the SizeValue field if non-nil, zero value otherwise.

### GetSizeValueOk

`func (o *SizeValueDTO) GetSizeValueOk() (*string, bool)`

GetSizeValueOk returns a tuple with the SizeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeValue

`func (o *SizeValueDTO) SetSizeValue(v string)`

SetSizeValue sets SizeValue field to given value.


### GetRowIndex

`func (o *SizeValueDTO) GetRowIndex() int32`

GetRowIndex returns the RowIndex field if non-nil, zero value otherwise.

### GetRowIndexOk

`func (o *SizeValueDTO) GetRowIndexOk() (*int32, bool)`

GetRowIndexOk returns a tuple with the RowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowIndex

`func (o *SizeValueDTO) SetRowIndex(v int32)`

SetRowIndex sets RowIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


