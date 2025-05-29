# SupplyRequestItemDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferId** | **string** | Ваш SKU — идентификатор товара в вашей системе.  Правила использования SKU:  * У каждого товара SKU должен быть свой.  * Уже заданный SKU нельзя освободить и использовать заново для другого товара. Каждый товар должен получать новый идентификатор, до того никогда не использовавшийся в вашем каталоге.  SKU товара можно изменить в кабинете продавца на Маркете. О том, как это сделать, читайте [в Справке Маркета для продавцов](https://yandex.ru/support2/marketplace/ru/assortment/operations/edit-sku).  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields)  | 
**Name** | **string** | Название товара. | 
**Price** | Pointer to [**CurrencyValueDTO**](CurrencyValueDTO.md) |  | [optional] 
**Counters** | [**SupplyRequestItemCountersDTO**](SupplyRequestItemCountersDTO.md) |  | 

## Methods

### NewSupplyRequestItemDTO

`func NewSupplyRequestItemDTO(offerId string, name string, counters SupplyRequestItemCountersDTO, ) *SupplyRequestItemDTO`

NewSupplyRequestItemDTO instantiates a new SupplyRequestItemDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplyRequestItemDTOWithDefaults

`func NewSupplyRequestItemDTOWithDefaults() *SupplyRequestItemDTO`

NewSupplyRequestItemDTOWithDefaults instantiates a new SupplyRequestItemDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferId

`func (o *SupplyRequestItemDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *SupplyRequestItemDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *SupplyRequestItemDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.


### GetName

`func (o *SupplyRequestItemDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SupplyRequestItemDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SupplyRequestItemDTO) SetName(v string)`

SetName sets Name field to given value.


### GetPrice

`func (o *SupplyRequestItemDTO) GetPrice() CurrencyValueDTO`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SupplyRequestItemDTO) GetPriceOk() (*CurrencyValueDTO, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SupplyRequestItemDTO) SetPrice(v CurrencyValueDTO)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SupplyRequestItemDTO) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCounters

`func (o *SupplyRequestItemDTO) GetCounters() SupplyRequestItemCountersDTO`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *SupplyRequestItemDTO) GetCountersOk() (*SupplyRequestItemCountersDTO, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *SupplyRequestItemDTO) SetCounters(v SupplyRequestItemCountersDTO)`

SetCounters sets Counters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


