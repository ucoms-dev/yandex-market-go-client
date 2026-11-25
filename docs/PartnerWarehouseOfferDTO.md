# PartnerWarehouseOfferDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferId** | **string** | Ваш SKU — идентификатор товара в вашей системе.  Правила использования SKU:  * У каждого товара SKU должен быть свой.  * Уже заданный SKU нельзя освободить и использовать заново для другого товара. Каждый товар должен получать новый идентификатор, до того никогда не использовавшийся в вашем каталоге.  SKU товара можно изменить в кабинете продавца на Маркете. О том, как это сделать, читайте [в Справке Маркета для продавцов](https://yandex.ru/support2/marketplace/ru/assortment/operations/edit-sku).  {% note warning %}  Пробельные символы в начале и конце значения автоматически удаляются. Например, &#x60;\&quot;  SKU123  \&quot;&#x60; и &#x60;\&quot;SKU123\&quot;&#x60; будут обработаны как одинаковые значения.  {% endnote %}  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields)  | 
**Stocks** | [**[]WarehouseStockDTO**](WarehouseStockDTO.md) | Информация об остатках. | 
**UpdatedAt** | Pointer to **time.Time** | Дата и время последнего обновления информации об остатках.  Формат даты и времени: ISO 8601 со смещением относительно UTC. Например, &#x60;2023-11-21T00:42:42+03:00&#x60;.  | [optional] 

## Methods

### NewPartnerWarehouseOfferDTO

`func NewPartnerWarehouseOfferDTO(offerId string, stocks []WarehouseStockDTO, ) *PartnerWarehouseOfferDTO`

NewPartnerWarehouseOfferDTO instantiates a new PartnerWarehouseOfferDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerWarehouseOfferDTOWithDefaults

`func NewPartnerWarehouseOfferDTOWithDefaults() *PartnerWarehouseOfferDTO`

NewPartnerWarehouseOfferDTOWithDefaults instantiates a new PartnerWarehouseOfferDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferId

`func (o *PartnerWarehouseOfferDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *PartnerWarehouseOfferDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *PartnerWarehouseOfferDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.


### GetStocks

`func (o *PartnerWarehouseOfferDTO) GetStocks() []WarehouseStockDTO`

GetStocks returns the Stocks field if non-nil, zero value otherwise.

### GetStocksOk

`func (o *PartnerWarehouseOfferDTO) GetStocksOk() (*[]WarehouseStockDTO, bool)`

GetStocksOk returns a tuple with the Stocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStocks

`func (o *PartnerWarehouseOfferDTO) SetStocks(v []WarehouseStockDTO)`

SetStocks sets Stocks field to given value.


### GetUpdatedAt

`func (o *PartnerWarehouseOfferDTO) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PartnerWarehouseOfferDTO) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PartnerWarehouseOfferDTO) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PartnerWarehouseOfferDTO) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


