# UpdateCampaignOfferDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferId** | **string** | Ваш SKU — идентификатор товара в вашей системе.  Правила использования SKU:  * У каждого товара SKU должен быть свой.  * Уже заданный SKU нельзя освободить и использовать заново для другого товара. Каждый товар должен получать новый идентификатор, до того никогда не использовавшийся в вашем каталоге.  SKU товара можно изменить в кабинете продавца на Маркете. О том, как это сделать, читайте [в Справке Маркета для продавцов](https://yandex.ru/support2/marketplace/ru/assortment/operations/edit-sku).  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields)  | 
**Quantum** | Pointer to [**QuantumDTO**](QuantumDTO.md) |  | [optional] 
**Available** | Pointer to **bool** | {% note warning \&quot;Вместо него используйте методы скрытия товаров с витрины\&quot; %}  * [GET campaigns/{campaignId}/hidden-offers](../../reference/assortment/getHiddenOffers.md) — просмотр скрытых товаров; * [POST campaigns/{campaignId}/hidden-offers](../../reference/assortment/addHiddenOffers.md) — скрытие товаров; * [POST campaigns/{campaignId}/hidden-offers/delete](../../reference/assortment/deleteHiddenOffers.md) — возобновление показа.  {% endnote %}  Есть ли товар в продаже.  | [optional] 
**Vat** | Pointer to **int32** | Идентификатор НДС, применяемый для товара:  * &#x60;2&#x60; — НДС 10%. Например, используется при реализации отдельных продовольственных и медицинских товаров. * &#x60;5&#x60; — НДС 0%. Например, используется при продаже товаров, вывезенных в таможенной процедуре экспорта, или при оказании услуг по международной перевозке товаров. * &#x60;6&#x60; — НДС не облагается, используется только для отдельных видов услуг. * &#x60;7&#x60; — НДС 20%. Основной НДС с 2019 года. * &#x60;10&#x60; — НДС 5%. НДС для упрощенной системы налогообложения (УСН). * &#x60;11&#x60; — НДС 7%. НДС для упрощенной системы налогообложения (УСН).  Если параметр не указан, используется НДС, установленный в кабинете.  **Для продавцов Market Yandex Go** недоступна передача и получение НДС.  | [optional] 

## Methods

### NewUpdateCampaignOfferDTO

`func NewUpdateCampaignOfferDTO(offerId string, ) *UpdateCampaignOfferDTO`

NewUpdateCampaignOfferDTO instantiates a new UpdateCampaignOfferDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCampaignOfferDTOWithDefaults

`func NewUpdateCampaignOfferDTOWithDefaults() *UpdateCampaignOfferDTO`

NewUpdateCampaignOfferDTOWithDefaults instantiates a new UpdateCampaignOfferDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferId

`func (o *UpdateCampaignOfferDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *UpdateCampaignOfferDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *UpdateCampaignOfferDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.


### GetQuantum

`func (o *UpdateCampaignOfferDTO) GetQuantum() QuantumDTO`

GetQuantum returns the Quantum field if non-nil, zero value otherwise.

### GetQuantumOk

`func (o *UpdateCampaignOfferDTO) GetQuantumOk() (*QuantumDTO, bool)`

GetQuantumOk returns a tuple with the Quantum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantum

`func (o *UpdateCampaignOfferDTO) SetQuantum(v QuantumDTO)`

SetQuantum sets Quantum field to given value.

### HasQuantum

`func (o *UpdateCampaignOfferDTO) HasQuantum() bool`

HasQuantum returns a boolean if a field has been set.

### GetAvailable

`func (o *UpdateCampaignOfferDTO) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *UpdateCampaignOfferDTO) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *UpdateCampaignOfferDTO) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *UpdateCampaignOfferDTO) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetVat

`func (o *UpdateCampaignOfferDTO) GetVat() int32`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *UpdateCampaignOfferDTO) GetVatOk() (*int32, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *UpdateCampaignOfferDTO) SetVat(v int32)`

SetVat sets Vat field to given value.

### HasVat

`func (o *UpdateCampaignOfferDTO) HasVat() bool`

HasVat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


