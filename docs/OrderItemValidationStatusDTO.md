# OrderItemValidationStatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Идентификатор товара в заказе. | 
**Uin** | [**[]UinDTO**](UinDTO.md) | Статусы проверки УИНов. | 

## Methods

### NewOrderItemValidationStatusDTO

`func NewOrderItemValidationStatusDTO(id int64, uin []UinDTO, ) *OrderItemValidationStatusDTO`

NewOrderItemValidationStatusDTO instantiates a new OrderItemValidationStatusDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderItemValidationStatusDTOWithDefaults

`func NewOrderItemValidationStatusDTOWithDefaults() *OrderItemValidationStatusDTO`

NewOrderItemValidationStatusDTOWithDefaults instantiates a new OrderItemValidationStatusDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderItemValidationStatusDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderItemValidationStatusDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderItemValidationStatusDTO) SetId(v int64)`

SetId sets Id field to given value.


### GetUin

`func (o *OrderItemValidationStatusDTO) GetUin() []UinDTO`

GetUin returns the Uin field if non-nil, zero value otherwise.

### GetUinOk

`func (o *OrderItemValidationStatusDTO) GetUinOk() (*[]UinDTO, bool)`

GetUinOk returns a tuple with the Uin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUin

`func (o *OrderItemValidationStatusDTO) SetUin(v []UinDTO)`

SetUin sets Uin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


