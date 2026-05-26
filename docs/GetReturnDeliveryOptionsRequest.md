# GetReturnDeliveryOptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]BasicOrderItemDTO**](BasicOrderItemDTO.md) | Товары в возврате. | 
**PickupDelivery** | [**PickupDeliveryParametersDTO**](PickupDeliveryParametersDTO.md) |  | 

## Methods

### NewGetReturnDeliveryOptionsRequest

`func NewGetReturnDeliveryOptionsRequest(items []BasicOrderItemDTO, pickupDelivery PickupDeliveryParametersDTO, ) *GetReturnDeliveryOptionsRequest`

NewGetReturnDeliveryOptionsRequest instantiates a new GetReturnDeliveryOptionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReturnDeliveryOptionsRequestWithDefaults

`func NewGetReturnDeliveryOptionsRequestWithDefaults() *GetReturnDeliveryOptionsRequest`

NewGetReturnDeliveryOptionsRequestWithDefaults instantiates a new GetReturnDeliveryOptionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GetReturnDeliveryOptionsRequest) GetItems() []BasicOrderItemDTO`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetReturnDeliveryOptionsRequest) GetItemsOk() (*[]BasicOrderItemDTO, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetReturnDeliveryOptionsRequest) SetItems(v []BasicOrderItemDTO)`

SetItems sets Items field to given value.


### GetPickupDelivery

`func (o *GetReturnDeliveryOptionsRequest) GetPickupDelivery() PickupDeliveryParametersDTO`

GetPickupDelivery returns the PickupDelivery field if non-nil, zero value otherwise.

### GetPickupDeliveryOk

`func (o *GetReturnDeliveryOptionsRequest) GetPickupDeliveryOk() (*PickupDeliveryParametersDTO, bool)`

GetPickupDeliveryOk returns a tuple with the PickupDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupDelivery

`func (o *GetReturnDeliveryOptionsRequest) SetPickupDelivery(v PickupDeliveryParametersDTO)`

SetPickupDelivery sets PickupDelivery field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


