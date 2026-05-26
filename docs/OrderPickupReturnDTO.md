# OrderPickupReturnDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogisticPointId** | **int64** | Идентификатор пункта выдачи.  Его можно узнать с помощью метода [POST v1/businesses/{businessId}/logistics-points](../../reference/logistic-points/getLogisticPoints.md).  | 

## Methods

### NewOrderPickupReturnDTO

`func NewOrderPickupReturnDTO(logisticPointId int64, ) *OrderPickupReturnDTO`

NewOrderPickupReturnDTO instantiates a new OrderPickupReturnDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderPickupReturnDTOWithDefaults

`func NewOrderPickupReturnDTOWithDefaults() *OrderPickupReturnDTO`

NewOrderPickupReturnDTOWithDefaults instantiates a new OrderPickupReturnDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogisticPointId

`func (o *OrderPickupReturnDTO) GetLogisticPointId() int64`

GetLogisticPointId returns the LogisticPointId field if non-nil, zero value otherwise.

### GetLogisticPointIdOk

`func (o *OrderPickupReturnDTO) GetLogisticPointIdOk() (*int64, bool)`

GetLogisticPointIdOk returns a tuple with the LogisticPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogisticPointId

`func (o *OrderPickupReturnDTO) SetLogisticPointId(v int64)`

SetLogisticPointId sets LogisticPointId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


