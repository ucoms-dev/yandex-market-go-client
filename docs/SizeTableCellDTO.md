# SizeTableCellDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeChartId** | Pointer to **int64** | Идентификатор размерной сетки или обмера из ответа метода &#x60;POST v2/businesses/{businessId}/size-tables/charts&#x60;. У сеток и обмеров общее пространство идентификаторов. Обязателен, если не передан &#x60;sizeChartName&#x60;.  | [optional] 
**SizeChartName** | Pointer to **string** | Название размерной сетки или обмера из ответа метода &#x60;POST v2/businesses/{businessId}/size-tables/charts&#x60;. Сравнивается без учета регистра и пробелов в начале и в конце. Обязательно, если не передан &#x60;sizeChartId&#x60;.  | [optional] 
**SizeValueId** | Pointer to **int64** | Идентификатор значения из справочника размерной сетки. Справочник есть только у стандартных размерных сеток Маркета.  | [optional] 
**SizeValue** | Pointer to **string** | Значение размера или обмера. В колонках без справочника — не длиннее 20 символов.  | [optional] 

## Methods

### NewSizeTableCellDTO

`func NewSizeTableCellDTO() *SizeTableCellDTO`

NewSizeTableCellDTO instantiates a new SizeTableCellDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSizeTableCellDTOWithDefaults

`func NewSizeTableCellDTOWithDefaults() *SizeTableCellDTO`

NewSizeTableCellDTOWithDefaults instantiates a new SizeTableCellDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeChartId

`func (o *SizeTableCellDTO) GetSizeChartId() int64`

GetSizeChartId returns the SizeChartId field if non-nil, zero value otherwise.

### GetSizeChartIdOk

`func (o *SizeTableCellDTO) GetSizeChartIdOk() (*int64, bool)`

GetSizeChartIdOk returns a tuple with the SizeChartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeChartId

`func (o *SizeTableCellDTO) SetSizeChartId(v int64)`

SetSizeChartId sets SizeChartId field to given value.

### HasSizeChartId

`func (o *SizeTableCellDTO) HasSizeChartId() bool`

HasSizeChartId returns a boolean if a field has been set.

### GetSizeChartName

`func (o *SizeTableCellDTO) GetSizeChartName() string`

GetSizeChartName returns the SizeChartName field if non-nil, zero value otherwise.

### GetSizeChartNameOk

`func (o *SizeTableCellDTO) GetSizeChartNameOk() (*string, bool)`

GetSizeChartNameOk returns a tuple with the SizeChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeChartName

`func (o *SizeTableCellDTO) SetSizeChartName(v string)`

SetSizeChartName sets SizeChartName field to given value.

### HasSizeChartName

`func (o *SizeTableCellDTO) HasSizeChartName() bool`

HasSizeChartName returns a boolean if a field has been set.

### GetSizeValueId

`func (o *SizeTableCellDTO) GetSizeValueId() int64`

GetSizeValueId returns the SizeValueId field if non-nil, zero value otherwise.

### GetSizeValueIdOk

`func (o *SizeTableCellDTO) GetSizeValueIdOk() (*int64, bool)`

GetSizeValueIdOk returns a tuple with the SizeValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeValueId

`func (o *SizeTableCellDTO) SetSizeValueId(v int64)`

SetSizeValueId sets SizeValueId field to given value.

### HasSizeValueId

`func (o *SizeTableCellDTO) HasSizeValueId() bool`

HasSizeValueId returns a boolean if a field has been set.

### GetSizeValue

`func (o *SizeTableCellDTO) GetSizeValue() string`

GetSizeValue returns the SizeValue field if non-nil, zero value otherwise.

### GetSizeValueOk

`func (o *SizeTableCellDTO) GetSizeValueOk() (*string, bool)`

GetSizeValueOk returns a tuple with the SizeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeValue

`func (o *SizeTableCellDTO) SetSizeValue(v string)`

SetSizeValue sets SizeValue field to given value.

### HasSizeValue

`func (o *SizeTableCellDTO) HasSizeValue() bool`

HasSizeValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


