# SizeTableRowDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | [**[]SizeTableCellDTO**](SizeTableCellDTO.md) | Ячейки строки в любом порядке. Каждая размерная сетка или обмер — не больше одного раза в строке, иначе вернется ошибка &#x60;INVALID_REQUEST&#x60;.  Ячейку без значения можно не передавать или передать с пустым &#x60;sizeValue&#x60;: она сохранится как «нет значения». В обязательной размерной сетке это ошибка &#x60;EMPTY_CELL&#x60;.  | 

## Methods

### NewSizeTableRowDTO

`func NewSizeTableRowDTO(values []SizeTableCellDTO, ) *SizeTableRowDTO`

NewSizeTableRowDTO instantiates a new SizeTableRowDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSizeTableRowDTOWithDefaults

`func NewSizeTableRowDTOWithDefaults() *SizeTableRowDTO`

NewSizeTableRowDTOWithDefaults instantiates a new SizeTableRowDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *SizeTableRowDTO) GetValues() []SizeTableCellDTO`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SizeTableRowDTO) GetValuesOk() (*[]SizeTableCellDTO, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SizeTableRowDTO) SetValues(v []SizeTableCellDTO)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


