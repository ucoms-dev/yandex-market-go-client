# UpdateBusinessPricesDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **float32** | Цена товара. | 
**CurrencyId** | [**CurrencyType**](CurrencyType.md) |  | 
**DiscountBase** | Pointer to **float32** | Зачеркнутая цена.  Число должно быть целым. Вы можете указать цену со скидкой от 5 до 99%.  Передавайте этот параметр при каждом обновлении цены, если предоставляете скидку на товар.  | [optional] 
**MinimumForBestseller** | Pointer to **float32** | Минимальная цена товара для попадания в акцию «Бестселлеры Маркета». Подробнее об этом способе участия читайте [в Справке Маркета для продавцов](https://yandex.ru/support/marketplace/ru/marketing/promos/market/bestsellers#minimum).  При передаче цены ориентируйтесь на значение параметра &#x60;maxPromoPrice&#x60; (максимально возможная цена для участия в акции) в методе [POST v2/businesses/{businessId}/promos/offers](../../reference/promos/getPromoOffers.md).  Товар не попадет в акцию с помощью этого способа, если:  * Не передать этот параметр. Удалится значение, которое вы указали ранее. * В методе [POST v2/businesses/{businessId}/offer-prices](../../reference/prices/getDefaultPrices.md) для этого товара возвращается параметр &#x60;excludedFromBestsellers&#x60; со значением &#x60;true&#x60;.    Но товар по-прежнему сможет попасть в акцию через [автоматическое участие](*auto) или [ручное добавление](*updatePromoOffers).  | [optional] 

## Methods

### NewUpdateBusinessPricesDTO

`func NewUpdateBusinessPricesDTO(value float32, currencyId CurrencyType, ) *UpdateBusinessPricesDTO`

NewUpdateBusinessPricesDTO instantiates a new UpdateBusinessPricesDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBusinessPricesDTOWithDefaults

`func NewUpdateBusinessPricesDTOWithDefaults() *UpdateBusinessPricesDTO`

NewUpdateBusinessPricesDTOWithDefaults instantiates a new UpdateBusinessPricesDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *UpdateBusinessPricesDTO) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateBusinessPricesDTO) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateBusinessPricesDTO) SetValue(v float32)`

SetValue sets Value field to given value.


### GetCurrencyId

`func (o *UpdateBusinessPricesDTO) GetCurrencyId() CurrencyType`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *UpdateBusinessPricesDTO) GetCurrencyIdOk() (*CurrencyType, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *UpdateBusinessPricesDTO) SetCurrencyId(v CurrencyType)`

SetCurrencyId sets CurrencyId field to given value.


### GetDiscountBase

`func (o *UpdateBusinessPricesDTO) GetDiscountBase() float32`

GetDiscountBase returns the DiscountBase field if non-nil, zero value otherwise.

### GetDiscountBaseOk

`func (o *UpdateBusinessPricesDTO) GetDiscountBaseOk() (*float32, bool)`

GetDiscountBaseOk returns a tuple with the DiscountBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountBase

`func (o *UpdateBusinessPricesDTO) SetDiscountBase(v float32)`

SetDiscountBase sets DiscountBase field to given value.

### HasDiscountBase

`func (o *UpdateBusinessPricesDTO) HasDiscountBase() bool`

HasDiscountBase returns a boolean if a field has been set.

### GetMinimumForBestseller

`func (o *UpdateBusinessPricesDTO) GetMinimumForBestseller() float32`

GetMinimumForBestseller returns the MinimumForBestseller field if non-nil, zero value otherwise.

### GetMinimumForBestsellerOk

`func (o *UpdateBusinessPricesDTO) GetMinimumForBestsellerOk() (*float32, bool)`

GetMinimumForBestsellerOk returns a tuple with the MinimumForBestseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumForBestseller

`func (o *UpdateBusinessPricesDTO) SetMinimumForBestseller(v float32)`

SetMinimumForBestseller sets MinimumForBestseller field to given value.

### HasMinimumForBestseller

`func (o *UpdateBusinessPricesDTO) HasMinimumForBestseller() bool`

HasMinimumForBestseller returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


