# GetDeliveryOptionsItemDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferId** | **string** | Ваш SKU — идентификатор товара в вашей системе.  Правила использования SKU:  * У каждого товара SKU должен быть свой.  * Уже заданный SKU нельзя освободить и использовать заново для другого товара. Каждый товар должен получать новый идентификатор, до того никогда не использовавшийся в вашем каталоге.  SKU товара можно изменить в кабинете продавца на Маркете. О том, как это сделать, читайте [в Справке Маркета для продавцов](https://yandex.ru/support2/marketplace/ru/assortment/operations/edit-sku).  {% note warning %}  Пробельные символы в начале и конце значения автоматически удаляются. Например, &#x60;\&quot;  SKU123  \&quot;&#x60; и &#x60;\&quot;SKU123\&quot;&#x60; будут обработаны как одинаковые значения.  {% endnote %}  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields)  | 
**Count** | **int32** | Количество единиц товара. | 
**WarehouseId** | Pointer to **int64** | Идентификатор фулфилмент-склада Маркета.  Передайте этот параметр, чтобы вернулись варианты доставки для указанного склада. Иначе Маркет сам выберет склад.  [Как узнать остатки товаров на складах](../../reference/stocks/getStocks.md)  | [optional] 

## Methods

### NewGetDeliveryOptionsItemDTO

`func NewGetDeliveryOptionsItemDTO(offerId string, count int32, ) *GetDeliveryOptionsItemDTO`

NewGetDeliveryOptionsItemDTO instantiates a new GetDeliveryOptionsItemDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeliveryOptionsItemDTOWithDefaults

`func NewGetDeliveryOptionsItemDTOWithDefaults() *GetDeliveryOptionsItemDTO`

NewGetDeliveryOptionsItemDTOWithDefaults instantiates a new GetDeliveryOptionsItemDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferId

`func (o *GetDeliveryOptionsItemDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *GetDeliveryOptionsItemDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *GetDeliveryOptionsItemDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.


### GetCount

`func (o *GetDeliveryOptionsItemDTO) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetDeliveryOptionsItemDTO) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetDeliveryOptionsItemDTO) SetCount(v int32)`

SetCount sets Count field to given value.


### GetWarehouseId

`func (o *GetDeliveryOptionsItemDTO) GetWarehouseId() int64`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *GetDeliveryOptionsItemDTO) GetWarehouseIdOk() (*int64, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *GetDeliveryOptionsItemDTO) SetWarehouseId(v int64)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *GetDeliveryOptionsItemDTO) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


