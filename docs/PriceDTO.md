# PriceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **float32** | Цена на товар. | [optional] 
**DiscountBase** | Pointer to **float32** | Цена на товар без скидки.  Число должно быть целым. Вы можете указать цену со скидкой от 5 до 99%.  Передавайте этот параметр при каждом обновлении цены, если предоставляете скидку на товар.  | [optional] 
**CurrencyId** | Pointer to [**CurrencyType**](CurrencyType.md) |  | [optional] 
**Vat** | Pointer to **int32** | Идентификатор НДС, применяемый для товара:  * &#x60;2&#x60; — НДС 10%. Например, используется при реализации отдельных продовольственных и медицинских товаров. * &#x60;5&#x60; — НДС 0%. Например, используется при продаже товаров, вывезенных в таможенной процедуре экспорта, или при оказании услуг по международной перевозке товаров. * &#x60;6&#x60; — НДС не облагается, используется только для отдельных видов услуг. * &#x60;7&#x60; — НДС 20%. Основной НДС с 2019 года. * &#x60;10&#x60; — НДС 5%. НДС для упрощенной системы налогообложения (УСН). * &#x60;11&#x60; — НДС 7%. НДС для упрощенной системы налогообложения (УСН).  Если параметр не указан, используется НДС, установленный в кабинете.  **Для продавцов Market Yandex Go** недоступна передача и получение НДС.  | [optional] 

## Methods

### NewPriceDTO

`func NewPriceDTO() *PriceDTO`

NewPriceDTO instantiates a new PriceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceDTOWithDefaults

`func NewPriceDTOWithDefaults() *PriceDTO`

NewPriceDTOWithDefaults instantiates a new PriceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *PriceDTO) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PriceDTO) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PriceDTO) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *PriceDTO) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDiscountBase

`func (o *PriceDTO) GetDiscountBase() float32`

GetDiscountBase returns the DiscountBase field if non-nil, zero value otherwise.

### GetDiscountBaseOk

`func (o *PriceDTO) GetDiscountBaseOk() (*float32, bool)`

GetDiscountBaseOk returns a tuple with the DiscountBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountBase

`func (o *PriceDTO) SetDiscountBase(v float32)`

SetDiscountBase sets DiscountBase field to given value.

### HasDiscountBase

`func (o *PriceDTO) HasDiscountBase() bool`

HasDiscountBase returns a boolean if a field has been set.

### GetCurrencyId

`func (o *PriceDTO) GetCurrencyId() CurrencyType`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *PriceDTO) GetCurrencyIdOk() (*CurrencyType, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *PriceDTO) SetCurrencyId(v CurrencyType)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *PriceDTO) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetVat

`func (o *PriceDTO) GetVat() int32`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *PriceDTO) GetVatOk() (*int32, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *PriceDTO) SetVat(v int32)`

SetVat sets Vat field to given value.

### HasVat

`func (o *PriceDTO) HasVat() bool`

HasVat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


