# OrderItemValidationStatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Идентификатор товара в заказе. | 
**Uin** | Pointer to [**[]UinDTO**](UinDTO.md) | Информация по проверке :no-translate[УИНов]. | [optional] 
**Cis** | Pointer to [**[]CisDTO**](CisDTO.md) | Информация по проверке кодов маркировки в системе :no-translate[«Честный ЗНАК»]. | [optional] 

## Methods

### NewOrderItemValidationStatusDTO

`func NewOrderItemValidationStatusDTO(id int64, ) *OrderItemValidationStatusDTO`

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

### HasUin

`func (o *OrderItemValidationStatusDTO) HasUin() bool`

HasUin returns a boolean if a field has been set.

### SetUinNil

`func (o *OrderItemValidationStatusDTO) SetUinNil(b bool)`

 SetUinNil sets the value for Uin to be an explicit nil

### UnsetUin
`func (o *OrderItemValidationStatusDTO) UnsetUin()`

UnsetUin ensures that no value is present for Uin, not even an explicit nil
### GetCis

`func (o *OrderItemValidationStatusDTO) GetCis() []CisDTO`

GetCis returns the Cis field if non-nil, zero value otherwise.

### GetCisOk

`func (o *OrderItemValidationStatusDTO) GetCisOk() (*[]CisDTO, bool)`

GetCisOk returns a tuple with the Cis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCis

`func (o *OrderItemValidationStatusDTO) SetCis(v []CisDTO)`

SetCis sets Cis field to given value.

### HasCis

`func (o *OrderItemValidationStatusDTO) HasCis() bool`

HasCis returns a boolean if a field has been set.

### SetCisNil

`func (o *OrderItemValidationStatusDTO) SetCisNil(b bool)`

 SetCisNil sets the value for Cis to be an explicit nil

### UnsetCis
`func (o *OrderItemValidationStatusDTO) UnsetCis()`

UnsetCis ensures that no value is present for Cis, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


