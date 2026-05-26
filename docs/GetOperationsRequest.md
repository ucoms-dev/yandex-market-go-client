# GetOperationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | [**OperationType**](OperationType.md) |  | 
**OperationIds** | **[]string** | Список идентификаторов операций.  | 

## Methods

### NewGetOperationsRequest

`func NewGetOperationsRequest(operationType OperationType, operationIds []string, ) *GetOperationsRequest`

NewGetOperationsRequest instantiates a new GetOperationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOperationsRequestWithDefaults

`func NewGetOperationsRequestWithDefaults() *GetOperationsRequest`

NewGetOperationsRequestWithDefaults instantiates a new GetOperationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *GetOperationsRequest) GetOperationType() OperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *GetOperationsRequest) GetOperationTypeOk() (*OperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *GetOperationsRequest) SetOperationType(v OperationType)`

SetOperationType sets OperationType field to given value.


### GetOperationIds

`func (o *GetOperationsRequest) GetOperationIds() []string`

GetOperationIds returns the OperationIds field if non-nil, zero value otherwise.

### GetOperationIdsOk

`func (o *GetOperationsRequest) GetOperationIdsOk() (*[]string, bool)`

GetOperationIdsOk returns a tuple with the OperationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationIds

`func (o *GetOperationsRequest) SetOperationIds(v []string)`

SetOperationIds sets OperationIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


