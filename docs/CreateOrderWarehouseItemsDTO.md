# CreateOrderWarehouseItemsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WarehouseId** | **int64** | Идентификатор фулфилмент-склада Маркета.  Получите его с помощью метода [POST v2/campaigns/{campaignId}/delivery-options](../../reference/delivery-options/getDeliveryOptions.md).  | 
**Items** | [**[]CreateOrderItemDTO**](CreateOrderItemDTO.md) | Список товаров в заказе.  В рамках одного запроса все значения &#x60;offerId&#x60; должны быть уникальными. Не допускается передача двух объектов с одинаковым &#x60;offerId&#x60;.  | 
**DeliveryDateInterval** | [**DeliveryDateIntervalDTO**](DeliveryDateIntervalDTO.md) |  | 
**DeliveryTimeInterval** | Pointer to [**TimeIntervalDTO**](TimeIntervalDTO.md) |  | [optional] 

## Methods

### NewCreateOrderWarehouseItemsDTO

`func NewCreateOrderWarehouseItemsDTO(warehouseId int64, items []CreateOrderItemDTO, deliveryDateInterval DeliveryDateIntervalDTO, ) *CreateOrderWarehouseItemsDTO`

NewCreateOrderWarehouseItemsDTO instantiates a new CreateOrderWarehouseItemsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderWarehouseItemsDTOWithDefaults

`func NewCreateOrderWarehouseItemsDTOWithDefaults() *CreateOrderWarehouseItemsDTO`

NewCreateOrderWarehouseItemsDTOWithDefaults instantiates a new CreateOrderWarehouseItemsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarehouseId

`func (o *CreateOrderWarehouseItemsDTO) GetWarehouseId() int64`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *CreateOrderWarehouseItemsDTO) GetWarehouseIdOk() (*int64, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *CreateOrderWarehouseItemsDTO) SetWarehouseId(v int64)`

SetWarehouseId sets WarehouseId field to given value.


### GetItems

`func (o *CreateOrderWarehouseItemsDTO) GetItems() []CreateOrderItemDTO`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CreateOrderWarehouseItemsDTO) GetItemsOk() (*[]CreateOrderItemDTO, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CreateOrderWarehouseItemsDTO) SetItems(v []CreateOrderItemDTO)`

SetItems sets Items field to given value.


### GetDeliveryDateInterval

`func (o *CreateOrderWarehouseItemsDTO) GetDeliveryDateInterval() DeliveryDateIntervalDTO`

GetDeliveryDateInterval returns the DeliveryDateInterval field if non-nil, zero value otherwise.

### GetDeliveryDateIntervalOk

`func (o *CreateOrderWarehouseItemsDTO) GetDeliveryDateIntervalOk() (*DeliveryDateIntervalDTO, bool)`

GetDeliveryDateIntervalOk returns a tuple with the DeliveryDateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDateInterval

`func (o *CreateOrderWarehouseItemsDTO) SetDeliveryDateInterval(v DeliveryDateIntervalDTO)`

SetDeliveryDateInterval sets DeliveryDateInterval field to given value.


### GetDeliveryTimeInterval

`func (o *CreateOrderWarehouseItemsDTO) GetDeliveryTimeInterval() TimeIntervalDTO`

GetDeliveryTimeInterval returns the DeliveryTimeInterval field if non-nil, zero value otherwise.

### GetDeliveryTimeIntervalOk

`func (o *CreateOrderWarehouseItemsDTO) GetDeliveryTimeIntervalOk() (*TimeIntervalDTO, bool)`

GetDeliveryTimeIntervalOk returns a tuple with the DeliveryTimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTimeInterval

`func (o *CreateOrderWarehouseItemsDTO) SetDeliveryTimeInterval(v TimeIntervalDTO)`

SetDeliveryTimeInterval sets DeliveryTimeInterval field to given value.

### HasDeliveryTimeInterval

`func (o *CreateOrderWarehouseItemsDTO) HasDeliveryTimeInterval() bool`

HasDeliveryTimeInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


