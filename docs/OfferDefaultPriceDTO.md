# OfferDefaultPriceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinimumForBestseller** | Pointer to **float32** | Минимальная цена товара для попадания в акцию «Бестселлеры Маркета». Подробнее об этом способе участия читайте [в Справке Маркета для продавцов](https://yandex.ru/support/marketplace/ru/marketing/promos/market/bestsellers#minimum).  Передается в методе [POST v2/businesses/{businessId}/offer-prices/updates](../../reference/prices/updateBusinessPrices.md).  | [optional] 
**ExcludedFromBestsellers** | Pointer to **bool** | Признак того, что товар не попадает в акцию «Бестселлеры Маркета». Подробнее об акции читайте [в Справке Маркета для продавцов](https://yandex.ru/support2/marketplace/ru/marketing/promos/market/bestsellers).  Если значение &#x60;true&#x60;, в методе [POST v2/businesses/{businessId}/offer-prices/updates](../../reference/prices/updateBusinessPrices.md) параметр &#x60;minimumForBestseller&#x60; игнорируется.  | [optional] 
**Value** | Pointer to **float32** | Цена товара. | [optional] 
**CurrencyId** | Pointer to [**CurrencyType**](CurrencyType.md) |  | [optional] 
**DiscountBase** | Pointer to **float32** | Зачеркнутая цена.  Число должно быть целым. Вы можете указать цену со скидкой от 5 до 99%.  Передавайте этот параметр при каждом обновлении цены, если предоставляете скидку на товар.  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Время последнего обновления. | [optional] 

## Methods

### NewOfferDefaultPriceDTO

`func NewOfferDefaultPriceDTO() *OfferDefaultPriceDTO`

NewOfferDefaultPriceDTO instantiates a new OfferDefaultPriceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferDefaultPriceDTOWithDefaults

`func NewOfferDefaultPriceDTOWithDefaults() *OfferDefaultPriceDTO`

NewOfferDefaultPriceDTOWithDefaults instantiates a new OfferDefaultPriceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinimumForBestseller

`func (o *OfferDefaultPriceDTO) GetMinimumForBestseller() float32`

GetMinimumForBestseller returns the MinimumForBestseller field if non-nil, zero value otherwise.

### GetMinimumForBestsellerOk

`func (o *OfferDefaultPriceDTO) GetMinimumForBestsellerOk() (*float32, bool)`

GetMinimumForBestsellerOk returns a tuple with the MinimumForBestseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumForBestseller

`func (o *OfferDefaultPriceDTO) SetMinimumForBestseller(v float32)`

SetMinimumForBestseller sets MinimumForBestseller field to given value.

### HasMinimumForBestseller

`func (o *OfferDefaultPriceDTO) HasMinimumForBestseller() bool`

HasMinimumForBestseller returns a boolean if a field has been set.

### GetExcludedFromBestsellers

`func (o *OfferDefaultPriceDTO) GetExcludedFromBestsellers() bool`

GetExcludedFromBestsellers returns the ExcludedFromBestsellers field if non-nil, zero value otherwise.

### GetExcludedFromBestsellersOk

`func (o *OfferDefaultPriceDTO) GetExcludedFromBestsellersOk() (*bool, bool)`

GetExcludedFromBestsellersOk returns a tuple with the ExcludedFromBestsellers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedFromBestsellers

`func (o *OfferDefaultPriceDTO) SetExcludedFromBestsellers(v bool)`

SetExcludedFromBestsellers sets ExcludedFromBestsellers field to given value.

### HasExcludedFromBestsellers

`func (o *OfferDefaultPriceDTO) HasExcludedFromBestsellers() bool`

HasExcludedFromBestsellers returns a boolean if a field has been set.

### GetValue

`func (o *OfferDefaultPriceDTO) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OfferDefaultPriceDTO) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OfferDefaultPriceDTO) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *OfferDefaultPriceDTO) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCurrencyId

`func (o *OfferDefaultPriceDTO) GetCurrencyId() CurrencyType`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *OfferDefaultPriceDTO) GetCurrencyIdOk() (*CurrencyType, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *OfferDefaultPriceDTO) SetCurrencyId(v CurrencyType)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *OfferDefaultPriceDTO) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetDiscountBase

`func (o *OfferDefaultPriceDTO) GetDiscountBase() float32`

GetDiscountBase returns the DiscountBase field if non-nil, zero value otherwise.

### GetDiscountBaseOk

`func (o *OfferDefaultPriceDTO) GetDiscountBaseOk() (*float32, bool)`

GetDiscountBaseOk returns a tuple with the DiscountBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountBase

`func (o *OfferDefaultPriceDTO) SetDiscountBase(v float32)`

SetDiscountBase sets DiscountBase field to given value.

### HasDiscountBase

`func (o *OfferDefaultPriceDTO) HasDiscountBase() bool`

HasDiscountBase returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OfferDefaultPriceDTO) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OfferDefaultPriceDTO) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OfferDefaultPriceDTO) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OfferDefaultPriceDTO) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


