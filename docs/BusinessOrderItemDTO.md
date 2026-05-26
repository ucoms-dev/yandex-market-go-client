# BusinessOrderItemDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Идентификатор товара в заказе.  Позволяет идентифицировать товар в рамках заказа.  | 
**OfferId** | **string** | Ваш SKU — идентификатор товара в вашей системе.  Правила использования SKU:  * У каждого товара SKU должен быть свой.  * Уже заданный SKU нельзя освободить и использовать заново для другого товара. Каждый товар должен получать новый идентификатор, до того никогда не использовавшийся в вашем каталоге.  SKU товара можно изменить в кабинете продавца на Маркете. О том, как это сделать, читайте [в Справке Маркета для продавцов](https://yandex.ru/support2/marketplace/ru/assortment/operations/edit-sku).  {% note warning %}  Пробельные символы в начале и конце значения автоматически удаляются. Например, &#x60;\&quot;  SKU123  \&quot;&#x60; и &#x60;\&quot;SKU123\&quot;&#x60; будут обработаны как одинаковые значения.  {% endnote %}  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields)  | 
**OfferName** | **string** | Название товара. | 
**Count** | **int32** | Количество единиц товара. | 
**Prices** | Pointer to [**ItemPriceDTO**](ItemPriceDTO.md) |  | [optional] 
**Instances** | Pointer to [**[]OrderItemInstanceDTO**](OrderItemInstanceDTO.md) | Информация о маркировке единиц товара.  Возвращаются данные для маркировки, переданные в запросе:  * Для DBS — [PUT v2/campaigns/{campaignId}/orders/{orderId}/identifiers](../../reference/orders/provideOrderItemIdentifiers.md) или [PUT v2/campaigns/{campaignId}/orders/{orderId}/boxes](../../reference/orders/setOrderBoxLayout.md). * Для FBS и EXPRESS — [PUT v2/campaigns/{campaignId}/orders/{orderId}/boxes](../../reference/orders/setOrderBoxLayout.md).  Для FBY возвращаются коды маркировки, переданные во время поставки.  Если магазин еще не передавал коды для этого заказа, &#x60;instances&#x60; отсутствует.  | [optional] 
**RequiredInstanceTypes** | Pointer to [**[]OrderItemInstanceType**](OrderItemInstanceType.md) | Список необходимых маркировок товара. | [optional] 
**Tags** | Pointer to [**[]OrderItemTagType**](OrderItemTagType.md) | Признаки товара. | [optional] 

## Methods

### NewBusinessOrderItemDTO

`func NewBusinessOrderItemDTO(id int64, offerId string, offerName string, count int32, ) *BusinessOrderItemDTO`

NewBusinessOrderItemDTO instantiates a new BusinessOrderItemDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessOrderItemDTOWithDefaults

`func NewBusinessOrderItemDTOWithDefaults() *BusinessOrderItemDTO`

NewBusinessOrderItemDTOWithDefaults instantiates a new BusinessOrderItemDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BusinessOrderItemDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BusinessOrderItemDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BusinessOrderItemDTO) SetId(v int64)`

SetId sets Id field to given value.


### GetOfferId

`func (o *BusinessOrderItemDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *BusinessOrderItemDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *BusinessOrderItemDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.


### GetOfferName

`func (o *BusinessOrderItemDTO) GetOfferName() string`

GetOfferName returns the OfferName field if non-nil, zero value otherwise.

### GetOfferNameOk

`func (o *BusinessOrderItemDTO) GetOfferNameOk() (*string, bool)`

GetOfferNameOk returns a tuple with the OfferName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferName

`func (o *BusinessOrderItemDTO) SetOfferName(v string)`

SetOfferName sets OfferName field to given value.


### GetCount

`func (o *BusinessOrderItemDTO) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BusinessOrderItemDTO) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BusinessOrderItemDTO) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPrices

`func (o *BusinessOrderItemDTO) GetPrices() ItemPriceDTO`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *BusinessOrderItemDTO) GetPricesOk() (*ItemPriceDTO, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *BusinessOrderItemDTO) SetPrices(v ItemPriceDTO)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *BusinessOrderItemDTO) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetInstances

`func (o *BusinessOrderItemDTO) GetInstances() []OrderItemInstanceDTO`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *BusinessOrderItemDTO) GetInstancesOk() (*[]OrderItemInstanceDTO, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *BusinessOrderItemDTO) SetInstances(v []OrderItemInstanceDTO)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *BusinessOrderItemDTO) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### SetInstancesNil

`func (o *BusinessOrderItemDTO) SetInstancesNil(b bool)`

 SetInstancesNil sets the value for Instances to be an explicit nil

### UnsetInstances
`func (o *BusinessOrderItemDTO) UnsetInstances()`

UnsetInstances ensures that no value is present for Instances, not even an explicit nil
### GetRequiredInstanceTypes

`func (o *BusinessOrderItemDTO) GetRequiredInstanceTypes() []OrderItemInstanceType`

GetRequiredInstanceTypes returns the RequiredInstanceTypes field if non-nil, zero value otherwise.

### GetRequiredInstanceTypesOk

`func (o *BusinessOrderItemDTO) GetRequiredInstanceTypesOk() (*[]OrderItemInstanceType, bool)`

GetRequiredInstanceTypesOk returns a tuple with the RequiredInstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredInstanceTypes

`func (o *BusinessOrderItemDTO) SetRequiredInstanceTypes(v []OrderItemInstanceType)`

SetRequiredInstanceTypes sets RequiredInstanceTypes field to given value.

### HasRequiredInstanceTypes

`func (o *BusinessOrderItemDTO) HasRequiredInstanceTypes() bool`

HasRequiredInstanceTypes returns a boolean if a field has been set.

### SetRequiredInstanceTypesNil

`func (o *BusinessOrderItemDTO) SetRequiredInstanceTypesNil(b bool)`

 SetRequiredInstanceTypesNil sets the value for RequiredInstanceTypes to be an explicit nil

### UnsetRequiredInstanceTypes
`func (o *BusinessOrderItemDTO) UnsetRequiredInstanceTypes()`

UnsetRequiredInstanceTypes ensures that no value is present for RequiredInstanceTypes, not even an explicit nil
### GetTags

`func (o *BusinessOrderItemDTO) GetTags() []OrderItemTagType`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BusinessOrderItemDTO) GetTagsOk() (*[]OrderItemTagType, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BusinessOrderItemDTO) SetTags(v []OrderItemTagType)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BusinessOrderItemDTO) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *BusinessOrderItemDTO) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *BusinessOrderItemDTO) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


