# OrderCourierDeliveryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**CourierDeliveryAddressDTO**](CourierDeliveryAddressDTO.md) |  | 
**Notes** | Pointer to **string** | Комментарий к заказу. | [optional] 

## Methods

### NewOrderCourierDeliveryDTO

`func NewOrderCourierDeliveryDTO(address CourierDeliveryAddressDTO, ) *OrderCourierDeliveryDTO`

NewOrderCourierDeliveryDTO instantiates a new OrderCourierDeliveryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCourierDeliveryDTOWithDefaults

`func NewOrderCourierDeliveryDTOWithDefaults() *OrderCourierDeliveryDTO`

NewOrderCourierDeliveryDTOWithDefaults instantiates a new OrderCourierDeliveryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *OrderCourierDeliveryDTO) GetAddress() CourierDeliveryAddressDTO`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OrderCourierDeliveryDTO) GetAddressOk() (*CourierDeliveryAddressDTO, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OrderCourierDeliveryDTO) SetAddress(v CourierDeliveryAddressDTO)`

SetAddress sets Address field to given value.


### GetNotes

`func (o *OrderCourierDeliveryDTO) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OrderCourierDeliveryDTO) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OrderCourierDeliveryDTO) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OrderCourierDeliveryDTO) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


