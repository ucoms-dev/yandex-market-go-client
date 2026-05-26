# OrderUpdateOptionsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryIntervals** | [**DeliveryIntervalsUpdateOptionsDTO**](DeliveryIntervalsUpdateOptionsDTO.md) |  | 

## Methods

### NewOrderUpdateOptionsDTO

`func NewOrderUpdateOptionsDTO(deliveryIntervals DeliveryIntervalsUpdateOptionsDTO, ) *OrderUpdateOptionsDTO`

NewOrderUpdateOptionsDTO instantiates a new OrderUpdateOptionsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderUpdateOptionsDTOWithDefaults

`func NewOrderUpdateOptionsDTOWithDefaults() *OrderUpdateOptionsDTO`

NewOrderUpdateOptionsDTOWithDefaults instantiates a new OrderUpdateOptionsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryIntervals

`func (o *OrderUpdateOptionsDTO) GetDeliveryIntervals() DeliveryIntervalsUpdateOptionsDTO`

GetDeliveryIntervals returns the DeliveryIntervals field if non-nil, zero value otherwise.

### GetDeliveryIntervalsOk

`func (o *OrderUpdateOptionsDTO) GetDeliveryIntervalsOk() (*DeliveryIntervalsUpdateOptionsDTO, bool)`

GetDeliveryIntervalsOk returns a tuple with the DeliveryIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryIntervals

`func (o *OrderUpdateOptionsDTO) SetDeliveryIntervals(v DeliveryIntervalsUpdateOptionsDTO)`

SetDeliveryIntervals sets DeliveryIntervals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


