# BarcodeOfferInfoDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferId** | **string** | Ваш SKU — идентификатор товара в вашей системе.  Правила использования SKU:  * У каждого товара SKU должен быть свой.  * Уже заданный SKU нельзя освободить и использовать заново для другого товара. Каждый товар должен получать новый идентификатор, до того никогда не использовавшийся в вашем каталоге.  SKU товара можно изменить в кабинете продавца на Маркете. О том, как это сделать, читайте [в Справке Маркета для продавцов](https://yandex.ru/support2/marketplace/ru/assortment/operations/edit-sku).  {% note warning %}  Пробельные символы в начале и конце значения автоматически удаляются. Например, &#x60;\&quot;  SKU123  \&quot;&#x60; и &#x60;\&quot;SKU123\&quot;&#x60; будут обработаны как одинаковые значения.  {% endnote %}  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields)  | 
**BarcodeCount** | **int64** | Количество штрихкодов для печати. | 

## Methods

### NewBarcodeOfferInfoDTO

`func NewBarcodeOfferInfoDTO(offerId string, barcodeCount int64, ) *BarcodeOfferInfoDTO`

NewBarcodeOfferInfoDTO instantiates a new BarcodeOfferInfoDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBarcodeOfferInfoDTOWithDefaults

`func NewBarcodeOfferInfoDTOWithDefaults() *BarcodeOfferInfoDTO`

NewBarcodeOfferInfoDTOWithDefaults instantiates a new BarcodeOfferInfoDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferId

`func (o *BarcodeOfferInfoDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *BarcodeOfferInfoDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *BarcodeOfferInfoDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.


### GetBarcodeCount

`func (o *BarcodeOfferInfoDTO) GetBarcodeCount() int64`

GetBarcodeCount returns the BarcodeCount field if non-nil, zero value otherwise.

### GetBarcodeCountOk

`func (o *BarcodeOfferInfoDTO) GetBarcodeCountOk() (*int64, bool)`

GetBarcodeCountOk returns a tuple with the BarcodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeCount

`func (o *BarcodeOfferInfoDTO) SetBarcodeCount(v int64)`

SetBarcodeCount sets BarcodeCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


