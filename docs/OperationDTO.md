# OperationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор операции. | 
**Type** | [**OperationType**](OperationType.md) |  | 

## Methods

### NewOperationDTO

`func NewOperationDTO(id string, type_ OperationType, ) *OperationDTO`

NewOperationDTO instantiates a new OperationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationDTOWithDefaults

`func NewOperationDTOWithDefaults() *OperationDTO`

NewOperationDTOWithDefaults instantiates a new OperationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OperationDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OperationDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OperationDTO) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *OperationDTO) GetType() OperationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OperationDTO) GetTypeOk() (*OperationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OperationDTO) SetType(v OperationType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


