# UpdateOrderDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Идентификатор заказа, в котором нужны изменения. | 
**DeliveryInterval** | Pointer to [**DeliveryIntervalsUpdateOptionDTO**](DeliveryIntervalsUpdateOptionDTO.md) |  | [optional] 
**Customer** | Pointer to [**CustomerDTO**](CustomerDTO.md) |  | [optional] 

## Methods

### NewUpdateOrderDTO

`func NewUpdateOrderDTO(id int64, ) *UpdateOrderDTO`

NewUpdateOrderDTO instantiates a new UpdateOrderDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrderDTOWithDefaults

`func NewUpdateOrderDTOWithDefaults() *UpdateOrderDTO`

NewUpdateOrderDTOWithDefaults instantiates a new UpdateOrderDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateOrderDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateOrderDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateOrderDTO) SetId(v int64)`

SetId sets Id field to given value.


### GetDeliveryInterval

`func (o *UpdateOrderDTO) GetDeliveryInterval() DeliveryIntervalsUpdateOptionDTO`

GetDeliveryInterval returns the DeliveryInterval field if non-nil, zero value otherwise.

### GetDeliveryIntervalOk

`func (o *UpdateOrderDTO) GetDeliveryIntervalOk() (*DeliveryIntervalsUpdateOptionDTO, bool)`

GetDeliveryIntervalOk returns a tuple with the DeliveryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryInterval

`func (o *UpdateOrderDTO) SetDeliveryInterval(v DeliveryIntervalsUpdateOptionDTO)`

SetDeliveryInterval sets DeliveryInterval field to given value.

### HasDeliveryInterval

`func (o *UpdateOrderDTO) HasDeliveryInterval() bool`

HasDeliveryInterval returns a boolean if a field has been set.

### GetCustomer

`func (o *UpdateOrderDTO) GetCustomer() CustomerDTO`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *UpdateOrderDTO) GetCustomerOk() (*CustomerDTO, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *UpdateOrderDTO) SetCustomer(v CustomerDTO)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *UpdateOrderDTO) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


