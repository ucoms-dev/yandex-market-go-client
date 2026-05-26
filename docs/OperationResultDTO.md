# OperationResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор операции. | 
**Type** | [**OperationType**](OperationType.md) |  | 
**Status** | [**OperationStatusType**](OperationStatusType.md) |  | 

## Methods

### NewOperationResultDTO

`func NewOperationResultDTO(id string, type_ OperationType, status OperationStatusType, ) *OperationResultDTO`

NewOperationResultDTO instantiates a new OperationResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResultDTOWithDefaults

`func NewOperationResultDTOWithDefaults() *OperationResultDTO`

NewOperationResultDTOWithDefaults instantiates a new OperationResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OperationResultDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OperationResultDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OperationResultDTO) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *OperationResultDTO) GetType() OperationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OperationResultDTO) GetTypeOk() (*OperationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OperationResultDTO) SetType(v OperationType)`

SetType sets Type field to given value.


### GetStatus

`func (o *OperationResultDTO) GetStatus() OperationStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OperationResultDTO) GetStatusOk() (*OperationStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OperationResultDTO) SetStatus(v OperationStatusType)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


