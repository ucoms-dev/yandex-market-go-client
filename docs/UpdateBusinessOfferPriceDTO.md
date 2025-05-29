# UpdateBusinessOfferPriceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferId** | **string** | Ваш SKU — идентификатор товара в вашей системе.  Правила использования SKU:  * У каждого товара SKU должен быть свой.  * Уже заданный SKU нельзя освободить и использовать заново для другого товара. Каждый товар должен получать новый идентификатор, до того никогда не использовавшийся в вашем каталоге.  SKU товара можно изменить в кабинете продавца на Маркете. О том, как это сделать, читайте [в Справке Маркета для продавцов](https://yandex.ru/support2/marketplace/ru/assortment/operations/edit-sku).  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields)  | 
**Price** | [**UpdateBusinessPricesDTO**](UpdateBusinessPricesDTO.md) |  | 

## Methods

### NewUpdateBusinessOfferPriceDTO

`func NewUpdateBusinessOfferPriceDTO(offerId string, price UpdateBusinessPricesDTO, ) *UpdateBusinessOfferPriceDTO`

NewUpdateBusinessOfferPriceDTO instantiates a new UpdateBusinessOfferPriceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBusinessOfferPriceDTOWithDefaults

`func NewUpdateBusinessOfferPriceDTOWithDefaults() *UpdateBusinessOfferPriceDTO`

NewUpdateBusinessOfferPriceDTOWithDefaults instantiates a new UpdateBusinessOfferPriceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferId

`func (o *UpdateBusinessOfferPriceDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *UpdateBusinessOfferPriceDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *UpdateBusinessOfferPriceDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.


### GetPrice

`func (o *UpdateBusinessOfferPriceDTO) GetPrice() UpdateBusinessPricesDTO`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UpdateBusinessOfferPriceDTO) GetPriceOk() (*UpdateBusinessPricesDTO, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UpdateBusinessOfferPriceDTO) SetPrice(v UpdateBusinessPricesDTO)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


