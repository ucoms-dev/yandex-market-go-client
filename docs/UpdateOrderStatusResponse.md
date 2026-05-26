# UpdateOrderStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**OrderDTO**](OrderDTO.md) |  | [optional] 
**Operation** | Pointer to [**OperationDTO**](OperationDTO.md) |  | [optional] 

## Methods

### NewUpdateOrderStatusResponse

`func NewUpdateOrderStatusResponse() *UpdateOrderStatusResponse`

NewUpdateOrderStatusResponse instantiates a new UpdateOrderStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrderStatusResponseWithDefaults

`func NewUpdateOrderStatusResponseWithDefaults() *UpdateOrderStatusResponse`

NewUpdateOrderStatusResponseWithDefaults instantiates a new UpdateOrderStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *UpdateOrderStatusResponse) GetOrder() OrderDTO`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *UpdateOrderStatusResponse) GetOrderOk() (*OrderDTO, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *UpdateOrderStatusResponse) SetOrder(v OrderDTO)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *UpdateOrderStatusResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOperation

`func (o *UpdateOrderStatusResponse) GetOperation() OperationDTO`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *UpdateOrderStatusResponse) GetOperationOk() (*OperationDTO, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *UpdateOrderStatusResponse) SetOperation(v OperationDTO)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *UpdateOrderStatusResponse) HasOperation() bool`

HasOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


