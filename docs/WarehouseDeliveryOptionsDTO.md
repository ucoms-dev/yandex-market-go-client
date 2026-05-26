# WarehouseDeliveryOptionsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PickupDelivery** | Pointer to [**PickupDeliveryOptionsDTO**](PickupDeliveryOptionsDTO.md) |  | [optional] 
**CourierDelivery** | Pointer to [**CourierDeliveryOptionsDTO**](CourierDeliveryOptionsDTO.md) |  | [optional] 

## Methods

### NewWarehouseDeliveryOptionsDTO

`func NewWarehouseDeliveryOptionsDTO() *WarehouseDeliveryOptionsDTO`

NewWarehouseDeliveryOptionsDTO instantiates a new WarehouseDeliveryOptionsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseDeliveryOptionsDTOWithDefaults

`func NewWarehouseDeliveryOptionsDTOWithDefaults() *WarehouseDeliveryOptionsDTO`

NewWarehouseDeliveryOptionsDTOWithDefaults instantiates a new WarehouseDeliveryOptionsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPickupDelivery

`func (o *WarehouseDeliveryOptionsDTO) GetPickupDelivery() PickupDeliveryOptionsDTO`

GetPickupDelivery returns the PickupDelivery field if non-nil, zero value otherwise.

### GetPickupDeliveryOk

`func (o *WarehouseDeliveryOptionsDTO) GetPickupDeliveryOk() (*PickupDeliveryOptionsDTO, bool)`

GetPickupDeliveryOk returns a tuple with the PickupDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupDelivery

`func (o *WarehouseDeliveryOptionsDTO) SetPickupDelivery(v PickupDeliveryOptionsDTO)`

SetPickupDelivery sets PickupDelivery field to given value.

### HasPickupDelivery

`func (o *WarehouseDeliveryOptionsDTO) HasPickupDelivery() bool`

HasPickupDelivery returns a boolean if a field has been set.

### GetCourierDelivery

`func (o *WarehouseDeliveryOptionsDTO) GetCourierDelivery() CourierDeliveryOptionsDTO`

GetCourierDelivery returns the CourierDelivery field if non-nil, zero value otherwise.

### GetCourierDeliveryOk

`func (o *WarehouseDeliveryOptionsDTO) GetCourierDeliveryOk() (*CourierDeliveryOptionsDTO, bool)`

GetCourierDeliveryOk returns a tuple with the CourierDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourierDelivery

`func (o *WarehouseDeliveryOptionsDTO) SetCourierDelivery(v CourierDeliveryOptionsDTO)`

SetCourierDelivery sets CourierDelivery field to given value.

### HasCourierDelivery

`func (o *WarehouseDeliveryOptionsDTO) HasCourierDelivery() bool`

HasCourierDelivery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


