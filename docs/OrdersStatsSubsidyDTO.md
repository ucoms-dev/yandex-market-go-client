# OrdersStatsSubsidyDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | [**OrdersStatsSubsidyOperationType**](OrdersStatsSubsidyOperationType.md) |  | 
**Type** | [**OrdersStatsSubsidyType**](OrdersStatsSubsidyType.md) |  | 
**Amount** | **float32** | Количество баллов, которые используются для уменьшения стоимости размещения, с точностью до двух знаков после запятой.  | 

## Methods

### NewOrdersStatsSubsidyDTO

`func NewOrdersStatsSubsidyDTO(operationType OrdersStatsSubsidyOperationType, type_ OrdersStatsSubsidyType, amount float32, ) *OrdersStatsSubsidyDTO`

NewOrdersStatsSubsidyDTO instantiates a new OrdersStatsSubsidyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersStatsSubsidyDTOWithDefaults

`func NewOrdersStatsSubsidyDTOWithDefaults() *OrdersStatsSubsidyDTO`

NewOrdersStatsSubsidyDTOWithDefaults instantiates a new OrdersStatsSubsidyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *OrdersStatsSubsidyDTO) GetOperationType() OrdersStatsSubsidyOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *OrdersStatsSubsidyDTO) GetOperationTypeOk() (*OrdersStatsSubsidyOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *OrdersStatsSubsidyDTO) SetOperationType(v OrdersStatsSubsidyOperationType)`

SetOperationType sets OperationType field to given value.


### GetType

`func (o *OrdersStatsSubsidyDTO) GetType() OrdersStatsSubsidyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrdersStatsSubsidyDTO) GetTypeOk() (*OrdersStatsSubsidyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrdersStatsSubsidyDTO) SetType(v OrdersStatsSubsidyType)`

SetType sets Type field to given value.


### GetAmount

`func (o *OrdersStatsSubsidyDTO) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrdersStatsSubsidyDTO) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrdersStatsSubsidyDTO) SetAmount(v float32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


