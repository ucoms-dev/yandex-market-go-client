# GetDeliveryOptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]GetDeliveryOptionsItemDTO**](GetDeliveryOptionsItemDTO.md) | Товары на складах, для которых нужно вернуть варианты доставки.  В рамках одного запроса все значения &#x60;offerId&#x60; должны быть уникальными. Не допускается передача двух объектов с одинаковым &#x60;offerId&#x60;.  | 
**PickupDelivery** | Pointer to [**PickupDeliveryParametersDTO**](PickupDeliveryParametersDTO.md) |  | [optional] 
**CourierDelivery** | Pointer to [**CourierDeliveryParametersDTO**](CourierDeliveryParametersDTO.md) |  | [optional] 

## Methods

### NewGetDeliveryOptionsRequest

`func NewGetDeliveryOptionsRequest(items []GetDeliveryOptionsItemDTO, ) *GetDeliveryOptionsRequest`

NewGetDeliveryOptionsRequest instantiates a new GetDeliveryOptionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeliveryOptionsRequestWithDefaults

`func NewGetDeliveryOptionsRequestWithDefaults() *GetDeliveryOptionsRequest`

NewGetDeliveryOptionsRequestWithDefaults instantiates a new GetDeliveryOptionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GetDeliveryOptionsRequest) GetItems() []GetDeliveryOptionsItemDTO`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetDeliveryOptionsRequest) GetItemsOk() (*[]GetDeliveryOptionsItemDTO, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetDeliveryOptionsRequest) SetItems(v []GetDeliveryOptionsItemDTO)`

SetItems sets Items field to given value.


### GetPickupDelivery

`func (o *GetDeliveryOptionsRequest) GetPickupDelivery() PickupDeliveryParametersDTO`

GetPickupDelivery returns the PickupDelivery field if non-nil, zero value otherwise.

### GetPickupDeliveryOk

`func (o *GetDeliveryOptionsRequest) GetPickupDeliveryOk() (*PickupDeliveryParametersDTO, bool)`

GetPickupDeliveryOk returns a tuple with the PickupDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupDelivery

`func (o *GetDeliveryOptionsRequest) SetPickupDelivery(v PickupDeliveryParametersDTO)`

SetPickupDelivery sets PickupDelivery field to given value.

### HasPickupDelivery

`func (o *GetDeliveryOptionsRequest) HasPickupDelivery() bool`

HasPickupDelivery returns a boolean if a field has been set.

### GetCourierDelivery

`func (o *GetDeliveryOptionsRequest) GetCourierDelivery() CourierDeliveryParametersDTO`

GetCourierDelivery returns the CourierDelivery field if non-nil, zero value otherwise.

### GetCourierDeliveryOk

`func (o *GetDeliveryOptionsRequest) GetCourierDeliveryOk() (*CourierDeliveryParametersDTO, bool)`

GetCourierDeliveryOk returns a tuple with the CourierDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourierDelivery

`func (o *GetDeliveryOptionsRequest) SetCourierDelivery(v CourierDeliveryParametersDTO)`

SetCourierDelivery sets CourierDelivery field to given value.

### HasCourierDelivery

`func (o *GetDeliveryOptionsRequest) HasCourierDelivery() bool`

HasCourierDelivery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


