# OrderItemUnitStatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**OrderItemUnitStatusType**](OrderItemUnitStatusType.md) |  | 
**Count** | **int64** | Количество единиц товара. | 

## Methods

### NewOrderItemUnitStatusDTO

`func NewOrderItemUnitStatusDTO(status OrderItemUnitStatusType, count int64, ) *OrderItemUnitStatusDTO`

NewOrderItemUnitStatusDTO instantiates a new OrderItemUnitStatusDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderItemUnitStatusDTOWithDefaults

`func NewOrderItemUnitStatusDTOWithDefaults() *OrderItemUnitStatusDTO`

NewOrderItemUnitStatusDTOWithDefaults instantiates a new OrderItemUnitStatusDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *OrderItemUnitStatusDTO) GetStatus() OrderItemUnitStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderItemUnitStatusDTO) GetStatusOk() (*OrderItemUnitStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderItemUnitStatusDTO) SetStatus(v OrderItemUnitStatusType)`

SetStatus sets Status field to given value.


### GetCount

`func (o *OrderItemUnitStatusDTO) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OrderItemUnitStatusDTO) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OrderItemUnitStatusDTO) SetCount(v int64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


