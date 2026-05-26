# BusinessOrderDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int64** | Идентификатор заказа. | 
**CampaignId** | **int64** | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями.  | 
**ProgramType** | Pointer to [**SellingProgramType**](SellingProgramType.md) |  | [optional] 
**ExternalOrderId** | Pointer to **string** | Внешний идентификатор заказа, который вы передали в [POST v2/campaigns/{campaignId}/orders/{orderId}/external-id](../../reference/orders/updateExternalOrderId.md). | [optional] 
**Status** | [**OrderStatusType**](OrderStatusType.md) |  | 
**Substatus** | [**OrderSubstatusType**](OrderSubstatusType.md) |  | 
**CreationDate** | **time.Time** | Дата и время оформления заказа.  Формат даты: ISO 8601 со смещением относительно UTC.  | 
**UpdateDate** | Pointer to **time.Time** | Дата и время последнего обновления заказа.  Формат даты: ISO 8601 со смещением относительно UTC.  | [optional] 
**PaymentType** | [**OrderPaymentType**](OrderPaymentType.md) |  | 
**PaymentMethod** | [**OrderPaymentMethodType**](OrderPaymentMethodType.md) |  | 
**Fake** | **bool** | Тип заказа:  * &#x60;false&#x60; — настоящий заказ покупателя.  * &#x60;true&#x60; — [тестовый заказ](../../concepts/sandbox.md) Маркета.  | 
**Items** | [**[]BusinessOrderItemDTO**](BusinessOrderItemDTO.md) | Список товаров в заказе. | 
**Prices** | Pointer to [**OrderPriceDTO**](OrderPriceDTO.md) |  | [optional] 
**Delivery** | [**BusinessOrderDeliveryDTO**](BusinessOrderDeliveryDTO.md) |  | 
**Services** | Pointer to [**BusinessOrderServicesDTO**](BusinessOrderServicesDTO.md) |  | [optional] 
**BuyerType** | Pointer to [**OrderBuyerType**](OrderBuyerType.md) |  | [optional] 
**Notes** | Pointer to **string** | Комментарий к заказу. | [optional] 
**CancelRequested** | Pointer to **bool** | **Только для модели DBS**  Запрошена ли отмена.  | [optional] 
**SourcePlatform** | Pointer to [**OrderSourcePlatformType**](OrderSourcePlatformType.md) |  | [optional] 

## Methods

### NewBusinessOrderDTO

`func NewBusinessOrderDTO(orderId int64, campaignId int64, status OrderStatusType, substatus OrderSubstatusType, creationDate time.Time, paymentType OrderPaymentType, paymentMethod OrderPaymentMethodType, fake bool, items []BusinessOrderItemDTO, delivery BusinessOrderDeliveryDTO, ) *BusinessOrderDTO`

NewBusinessOrderDTO instantiates a new BusinessOrderDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessOrderDTOWithDefaults

`func NewBusinessOrderDTOWithDefaults() *BusinessOrderDTO`

NewBusinessOrderDTOWithDefaults instantiates a new BusinessOrderDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *BusinessOrderDTO) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *BusinessOrderDTO) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *BusinessOrderDTO) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetCampaignId

`func (o *BusinessOrderDTO) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *BusinessOrderDTO) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *BusinessOrderDTO) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetProgramType

`func (o *BusinessOrderDTO) GetProgramType() SellingProgramType`

GetProgramType returns the ProgramType field if non-nil, zero value otherwise.

### GetProgramTypeOk

`func (o *BusinessOrderDTO) GetProgramTypeOk() (*SellingProgramType, bool)`

GetProgramTypeOk returns a tuple with the ProgramType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramType

`func (o *BusinessOrderDTO) SetProgramType(v SellingProgramType)`

SetProgramType sets ProgramType field to given value.

### HasProgramType

`func (o *BusinessOrderDTO) HasProgramType() bool`

HasProgramType returns a boolean if a field has been set.

### GetExternalOrderId

`func (o *BusinessOrderDTO) GetExternalOrderId() string`

GetExternalOrderId returns the ExternalOrderId field if non-nil, zero value otherwise.

### GetExternalOrderIdOk

`func (o *BusinessOrderDTO) GetExternalOrderIdOk() (*string, bool)`

GetExternalOrderIdOk returns a tuple with the ExternalOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOrderId

`func (o *BusinessOrderDTO) SetExternalOrderId(v string)`

SetExternalOrderId sets ExternalOrderId field to given value.

### HasExternalOrderId

`func (o *BusinessOrderDTO) HasExternalOrderId() bool`

HasExternalOrderId returns a boolean if a field has been set.

### GetStatus

`func (o *BusinessOrderDTO) GetStatus() OrderStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BusinessOrderDTO) GetStatusOk() (*OrderStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BusinessOrderDTO) SetStatus(v OrderStatusType)`

SetStatus sets Status field to given value.


### GetSubstatus

`func (o *BusinessOrderDTO) GetSubstatus() OrderSubstatusType`

GetSubstatus returns the Substatus field if non-nil, zero value otherwise.

### GetSubstatusOk

`func (o *BusinessOrderDTO) GetSubstatusOk() (*OrderSubstatusType, bool)`

GetSubstatusOk returns a tuple with the Substatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstatus

`func (o *BusinessOrderDTO) SetSubstatus(v OrderSubstatusType)`

SetSubstatus sets Substatus field to given value.


### GetCreationDate

`func (o *BusinessOrderDTO) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *BusinessOrderDTO) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *BusinessOrderDTO) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.


### GetUpdateDate

`func (o *BusinessOrderDTO) GetUpdateDate() time.Time`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *BusinessOrderDTO) GetUpdateDateOk() (*time.Time, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *BusinessOrderDTO) SetUpdateDate(v time.Time)`

SetUpdateDate sets UpdateDate field to given value.

### HasUpdateDate

`func (o *BusinessOrderDTO) HasUpdateDate() bool`

HasUpdateDate returns a boolean if a field has been set.

### GetPaymentType

`func (o *BusinessOrderDTO) GetPaymentType() OrderPaymentType`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *BusinessOrderDTO) GetPaymentTypeOk() (*OrderPaymentType, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *BusinessOrderDTO) SetPaymentType(v OrderPaymentType)`

SetPaymentType sets PaymentType field to given value.


### GetPaymentMethod

`func (o *BusinessOrderDTO) GetPaymentMethod() OrderPaymentMethodType`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *BusinessOrderDTO) GetPaymentMethodOk() (*OrderPaymentMethodType, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *BusinessOrderDTO) SetPaymentMethod(v OrderPaymentMethodType)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetFake

`func (o *BusinessOrderDTO) GetFake() bool`

GetFake returns the Fake field if non-nil, zero value otherwise.

### GetFakeOk

`func (o *BusinessOrderDTO) GetFakeOk() (*bool, bool)`

GetFakeOk returns a tuple with the Fake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFake

`func (o *BusinessOrderDTO) SetFake(v bool)`

SetFake sets Fake field to given value.


### GetItems

`func (o *BusinessOrderDTO) GetItems() []BusinessOrderItemDTO`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BusinessOrderDTO) GetItemsOk() (*[]BusinessOrderItemDTO, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BusinessOrderDTO) SetItems(v []BusinessOrderItemDTO)`

SetItems sets Items field to given value.


### GetPrices

`func (o *BusinessOrderDTO) GetPrices() OrderPriceDTO`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *BusinessOrderDTO) GetPricesOk() (*OrderPriceDTO, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *BusinessOrderDTO) SetPrices(v OrderPriceDTO)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *BusinessOrderDTO) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetDelivery

`func (o *BusinessOrderDTO) GetDelivery() BusinessOrderDeliveryDTO`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *BusinessOrderDTO) GetDeliveryOk() (*BusinessOrderDeliveryDTO, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *BusinessOrderDTO) SetDelivery(v BusinessOrderDeliveryDTO)`

SetDelivery sets Delivery field to given value.


### GetServices

`func (o *BusinessOrderDTO) GetServices() BusinessOrderServicesDTO`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *BusinessOrderDTO) GetServicesOk() (*BusinessOrderServicesDTO, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *BusinessOrderDTO) SetServices(v BusinessOrderServicesDTO)`

SetServices sets Services field to given value.

### HasServices

`func (o *BusinessOrderDTO) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetBuyerType

`func (o *BusinessOrderDTO) GetBuyerType() OrderBuyerType`

GetBuyerType returns the BuyerType field if non-nil, zero value otherwise.

### GetBuyerTypeOk

`func (o *BusinessOrderDTO) GetBuyerTypeOk() (*OrderBuyerType, bool)`

GetBuyerTypeOk returns a tuple with the BuyerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerType

`func (o *BusinessOrderDTO) SetBuyerType(v OrderBuyerType)`

SetBuyerType sets BuyerType field to given value.

### HasBuyerType

`func (o *BusinessOrderDTO) HasBuyerType() bool`

HasBuyerType returns a boolean if a field has been set.

### GetNotes

`func (o *BusinessOrderDTO) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BusinessOrderDTO) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BusinessOrderDTO) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BusinessOrderDTO) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetCancelRequested

`func (o *BusinessOrderDTO) GetCancelRequested() bool`

GetCancelRequested returns the CancelRequested field if non-nil, zero value otherwise.

### GetCancelRequestedOk

`func (o *BusinessOrderDTO) GetCancelRequestedOk() (*bool, bool)`

GetCancelRequestedOk returns a tuple with the CancelRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelRequested

`func (o *BusinessOrderDTO) SetCancelRequested(v bool)`

SetCancelRequested sets CancelRequested field to given value.

### HasCancelRequested

`func (o *BusinessOrderDTO) HasCancelRequested() bool`

HasCancelRequested returns a boolean if a field has been set.

### GetSourcePlatform

`func (o *BusinessOrderDTO) GetSourcePlatform() OrderSourcePlatformType`

GetSourcePlatform returns the SourcePlatform field if non-nil, zero value otherwise.

### GetSourcePlatformOk

`func (o *BusinessOrderDTO) GetSourcePlatformOk() (*OrderSourcePlatformType, bool)`

GetSourcePlatformOk returns a tuple with the SourcePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePlatform

`func (o *BusinessOrderDTO) SetSourcePlatform(v OrderSourcePlatformType)`

SetSourcePlatform sets SourcePlatform field to given value.

### HasSourcePlatform

`func (o *BusinessOrderDTO) HasSourcePlatform() bool`

HasSourcePlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


