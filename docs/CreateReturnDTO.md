# CreateReturnDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalReturnId** | **string** | Внешний идентификатор возврата в системе магазина. | 
**OrderId** | **int64** | Идентификатор заказа, по которому нужно сделать возврат. | 
**Items** | [**[]CreateReturnItemDTO**](CreateReturnItemDTO.md) | Список товаров в возврате. | 
**Customer** | [**CustomerDTO**](CustomerDTO.md) |  | 
**ReturnOption** | [**CreateReturnOptionDTO**](CreateReturnOptionDTO.md) |  | 

## Methods

### NewCreateReturnDTO

`func NewCreateReturnDTO(externalReturnId string, orderId int64, items []CreateReturnItemDTO, customer CustomerDTO, returnOption CreateReturnOptionDTO, ) *CreateReturnDTO`

NewCreateReturnDTO instantiates a new CreateReturnDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReturnDTOWithDefaults

`func NewCreateReturnDTOWithDefaults() *CreateReturnDTO`

NewCreateReturnDTOWithDefaults instantiates a new CreateReturnDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalReturnId

`func (o *CreateReturnDTO) GetExternalReturnId() string`

GetExternalReturnId returns the ExternalReturnId field if non-nil, zero value otherwise.

### GetExternalReturnIdOk

`func (o *CreateReturnDTO) GetExternalReturnIdOk() (*string, bool)`

GetExternalReturnIdOk returns a tuple with the ExternalReturnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReturnId

`func (o *CreateReturnDTO) SetExternalReturnId(v string)`

SetExternalReturnId sets ExternalReturnId field to given value.


### GetOrderId

`func (o *CreateReturnDTO) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CreateReturnDTO) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CreateReturnDTO) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetItems

`func (o *CreateReturnDTO) GetItems() []CreateReturnItemDTO`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CreateReturnDTO) GetItemsOk() (*[]CreateReturnItemDTO, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CreateReturnDTO) SetItems(v []CreateReturnItemDTO)`

SetItems sets Items field to given value.


### GetCustomer

`func (o *CreateReturnDTO) GetCustomer() CustomerDTO`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CreateReturnDTO) GetCustomerOk() (*CustomerDTO, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CreateReturnDTO) SetCustomer(v CustomerDTO)`

SetCustomer sets Customer field to given value.


### GetReturnOption

`func (o *CreateReturnDTO) GetReturnOption() CreateReturnOptionDTO`

GetReturnOption returns the ReturnOption field if non-nil, zero value otherwise.

### GetReturnOptionOk

`func (o *CreateReturnDTO) GetReturnOptionOk() (*CreateReturnOptionDTO, bool)`

GetReturnOptionOk returns a tuple with the ReturnOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnOption

`func (o *CreateReturnDTO) SetReturnOption(v CreateReturnOptionDTO)`

SetReturnOption sets ReturnOption field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


