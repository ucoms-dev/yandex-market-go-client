# WarehousesDeliveryOptionsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WarehouseId** | **int64** | Идентификатор фулфилмент-склада Маркета. | 
**DeliveryOptions** | [**WarehouseDeliveryOptionsDTO**](WarehouseDeliveryOptionsDTO.md) |  | 
**Items** | [**[]BasicOrderItemDTO**](BasicOrderItemDTO.md) | Товары в заказе. | 

## Methods

### NewWarehousesDeliveryOptionsDTO

`func NewWarehousesDeliveryOptionsDTO(warehouseId int64, deliveryOptions WarehouseDeliveryOptionsDTO, items []BasicOrderItemDTO, ) *WarehousesDeliveryOptionsDTO`

NewWarehousesDeliveryOptionsDTO instantiates a new WarehousesDeliveryOptionsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehousesDeliveryOptionsDTOWithDefaults

`func NewWarehousesDeliveryOptionsDTOWithDefaults() *WarehousesDeliveryOptionsDTO`

NewWarehousesDeliveryOptionsDTOWithDefaults instantiates a new WarehousesDeliveryOptionsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarehouseId

`func (o *WarehousesDeliveryOptionsDTO) GetWarehouseId() int64`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *WarehousesDeliveryOptionsDTO) GetWarehouseIdOk() (*int64, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *WarehousesDeliveryOptionsDTO) SetWarehouseId(v int64)`

SetWarehouseId sets WarehouseId field to given value.


### GetDeliveryOptions

`func (o *WarehousesDeliveryOptionsDTO) GetDeliveryOptions() WarehouseDeliveryOptionsDTO`

GetDeliveryOptions returns the DeliveryOptions field if non-nil, zero value otherwise.

### GetDeliveryOptionsOk

`func (o *WarehousesDeliveryOptionsDTO) GetDeliveryOptionsOk() (*WarehouseDeliveryOptionsDTO, bool)`

GetDeliveryOptionsOk returns a tuple with the DeliveryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryOptions

`func (o *WarehousesDeliveryOptionsDTO) SetDeliveryOptions(v WarehouseDeliveryOptionsDTO)`

SetDeliveryOptions sets DeliveryOptions field to given value.


### GetItems

`func (o *WarehousesDeliveryOptionsDTO) GetItems() []BasicOrderItemDTO`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *WarehousesDeliveryOptionsDTO) GetItemsOk() (*[]BasicOrderItemDTO, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *WarehousesDeliveryOptionsDTO) SetItems(v []BasicOrderItemDTO)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


