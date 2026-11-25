# DeleteDocumentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentIds** | **[]int64** | Идентификаторы документов. Значения не должны повторяться.  | 

## Methods

### NewDeleteDocumentsRequest

`func NewDeleteDocumentsRequest(documentIds []int64, ) *DeleteDocumentsRequest`

NewDeleteDocumentsRequest instantiates a new DeleteDocumentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDocumentsRequestWithDefaults

`func NewDeleteDocumentsRequestWithDefaults() *DeleteDocumentsRequest`

NewDeleteDocumentsRequestWithDefaults instantiates a new DeleteDocumentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentIds

`func (o *DeleteDocumentsRequest) GetDocumentIds() []int64`

GetDocumentIds returns the DocumentIds field if non-nil, zero value otherwise.

### GetDocumentIdsOk

`func (o *DeleteDocumentsRequest) GetDocumentIdsOk() (*[]int64, bool)`

GetDocumentIdsOk returns a tuple with the DocumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIds

`func (o *DeleteDocumentsRequest) SetDocumentIds(v []int64)`

SetDocumentIds sets DocumentIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


