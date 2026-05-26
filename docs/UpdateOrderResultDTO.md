# UpdateOrderResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | [**[]OperationDTO**](OperationDTO.md) | Информация о запущенных операциях по изменению заказа.  | 

## Methods

### NewUpdateOrderResultDTO

`func NewUpdateOrderResultDTO(operations []OperationDTO, ) *UpdateOrderResultDTO`

NewUpdateOrderResultDTO instantiates a new UpdateOrderResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrderResultDTOWithDefaults

`func NewUpdateOrderResultDTOWithDefaults() *UpdateOrderResultDTO`

NewUpdateOrderResultDTOWithDefaults instantiates a new UpdateOrderResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *UpdateOrderResultDTO) GetOperations() []OperationDTO`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *UpdateOrderResultDTO) GetOperationsOk() (*[]OperationDTO, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *UpdateOrderResultDTO) SetOperations(v []OperationDTO)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


