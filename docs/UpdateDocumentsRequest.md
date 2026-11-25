# UpdateDocumentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documents** | [**[]UpdateDocumentDTO**](UpdateDocumentDTO.md) | Документы для обновления. Идентификаторы документов не должны повторяться в одном запросе. Если идентификаторы повторяются, запрос не обрабатывается.  | 

## Methods

### NewUpdateDocumentsRequest

`func NewUpdateDocumentsRequest(documents []UpdateDocumentDTO, ) *UpdateDocumentsRequest`

NewUpdateDocumentsRequest instantiates a new UpdateDocumentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDocumentsRequestWithDefaults

`func NewUpdateDocumentsRequestWithDefaults() *UpdateDocumentsRequest`

NewUpdateDocumentsRequestWithDefaults instantiates a new UpdateDocumentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocuments

`func (o *UpdateDocumentsRequest) GetDocuments() []UpdateDocumentDTO`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *UpdateDocumentsRequest) GetDocumentsOk() (*[]UpdateDocumentDTO, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *UpdateDocumentsRequest) SetDocuments(v []UpdateDocumentDTO)`

SetDocuments sets Documents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


