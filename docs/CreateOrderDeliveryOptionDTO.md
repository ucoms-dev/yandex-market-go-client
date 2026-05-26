# CreateOrderDeliveryOptionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PickupDelivery** | Pointer to [**OrderPickupDeliveryDTO**](OrderPickupDeliveryDTO.md) |  | [optional] 
**CourierDelivery** | Pointer to [**OrderCourierDeliveryDTO**](OrderCourierDeliveryDTO.md) |  | [optional] 

## Methods

### NewCreateOrderDeliveryOptionDTO

`func NewCreateOrderDeliveryOptionDTO() *CreateOrderDeliveryOptionDTO`

NewCreateOrderDeliveryOptionDTO instantiates a new CreateOrderDeliveryOptionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderDeliveryOptionDTOWithDefaults

`func NewCreateOrderDeliveryOptionDTOWithDefaults() *CreateOrderDeliveryOptionDTO`

NewCreateOrderDeliveryOptionDTOWithDefaults instantiates a new CreateOrderDeliveryOptionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPickupDelivery

`func (o *CreateOrderDeliveryOptionDTO) GetPickupDelivery() OrderPickupDeliveryDTO`

GetPickupDelivery returns the PickupDelivery field if non-nil, zero value otherwise.

### GetPickupDeliveryOk

`func (o *CreateOrderDeliveryOptionDTO) GetPickupDeliveryOk() (*OrderPickupDeliveryDTO, bool)`

GetPickupDeliveryOk returns a tuple with the PickupDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupDelivery

`func (o *CreateOrderDeliveryOptionDTO) SetPickupDelivery(v OrderPickupDeliveryDTO)`

SetPickupDelivery sets PickupDelivery field to given value.

### HasPickupDelivery

`func (o *CreateOrderDeliveryOptionDTO) HasPickupDelivery() bool`

HasPickupDelivery returns a boolean if a field has been set.

### GetCourierDelivery

`func (o *CreateOrderDeliveryOptionDTO) GetCourierDelivery() OrderCourierDeliveryDTO`

GetCourierDelivery returns the CourierDelivery field if non-nil, zero value otherwise.

### GetCourierDeliveryOk

`func (o *CreateOrderDeliveryOptionDTO) GetCourierDeliveryOk() (*OrderCourierDeliveryDTO, bool)`

GetCourierDeliveryOk returns a tuple with the CourierDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourierDelivery

`func (o *CreateOrderDeliveryOptionDTO) SetCourierDelivery(v OrderCourierDeliveryDTO)`

SetCourierDelivery sets CourierDelivery field to given value.

### HasCourierDelivery

`func (o *CreateOrderDeliveryOptionDTO) HasCourierDelivery() bool`

HasCourierDelivery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


