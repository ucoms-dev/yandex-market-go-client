# CreateChatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** | Идентификатор заказа на Маркете. | [optional] 
**Context** | Pointer to [**ChatContextDTO**](ChatContextDTO.md) |  | [optional] 

## Methods

### NewCreateChatRequest

`func NewCreateChatRequest() *CreateChatRequest`

NewCreateChatRequest instantiates a new CreateChatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateChatRequestWithDefaults

`func NewCreateChatRequestWithDefaults() *CreateChatRequest`

NewCreateChatRequestWithDefaults instantiates a new CreateChatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *CreateChatRequest) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CreateChatRequest) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CreateChatRequest) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CreateChatRequest) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetContext

`func (o *CreateChatRequest) GetContext() ChatContextDTO`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CreateChatRequest) GetContextOk() (*ChatContextDTO, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CreateChatRequest) SetContext(v ChatContextDTO)`

SetContext sets Context field to given value.

### HasContext

`func (o *CreateChatRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


