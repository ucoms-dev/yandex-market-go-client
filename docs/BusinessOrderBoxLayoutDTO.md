# BusinessOrderBoxLayoutDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]BusinessOrderBoxLayoutItemDTO**](BusinessOrderBoxLayoutItemDTO.md) | Список товаров в коробке.  Если в коробке едет часть большого товара, в списке может быть только один пункт.  | 
**BoxId** | **int64** | Идентификатор коробки. | 
**Barcode** | **string** | Идентификатор грузового места в системе магазина. | 

## Methods

### NewBusinessOrderBoxLayoutDTO

`func NewBusinessOrderBoxLayoutDTO(items []BusinessOrderBoxLayoutItemDTO, boxId int64, barcode string, ) *BusinessOrderBoxLayoutDTO`

NewBusinessOrderBoxLayoutDTO instantiates a new BusinessOrderBoxLayoutDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessOrderBoxLayoutDTOWithDefaults

`func NewBusinessOrderBoxLayoutDTOWithDefaults() *BusinessOrderBoxLayoutDTO`

NewBusinessOrderBoxLayoutDTOWithDefaults instantiates a new BusinessOrderBoxLayoutDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BusinessOrderBoxLayoutDTO) GetItems() []BusinessOrderBoxLayoutItemDTO`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BusinessOrderBoxLayoutDTO) GetItemsOk() (*[]BusinessOrderBoxLayoutItemDTO, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BusinessOrderBoxLayoutDTO) SetItems(v []BusinessOrderBoxLayoutItemDTO)`

SetItems sets Items field to given value.


### GetBoxId

`func (o *BusinessOrderBoxLayoutDTO) GetBoxId() int64`

GetBoxId returns the BoxId field if non-nil, zero value otherwise.

### GetBoxIdOk

`func (o *BusinessOrderBoxLayoutDTO) GetBoxIdOk() (*int64, bool)`

GetBoxIdOk returns a tuple with the BoxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxId

`func (o *BusinessOrderBoxLayoutDTO) SetBoxId(v int64)`

SetBoxId sets BoxId field to given value.


### GetBarcode

`func (o *BusinessOrderBoxLayoutDTO) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *BusinessOrderBoxLayoutDTO) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *BusinessOrderBoxLayoutDTO) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


