# CreateDocumentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documents** | [**[]DocumentWriteDTO**](DocumentWriteDTO.md) | Документы для создания. Номера документов не должны повторяться в одном запросе. Если номера повторяются, запрос не обрабатывается.  | 

## Methods

### NewCreateDocumentsRequest

`func NewCreateDocumentsRequest(documents []DocumentWriteDTO, ) *CreateDocumentsRequest`

NewCreateDocumentsRequest instantiates a new CreateDocumentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDocumentsRequestWithDefaults

`func NewCreateDocumentsRequestWithDefaults() *CreateDocumentsRequest`

NewCreateDocumentsRequestWithDefaults instantiates a new CreateDocumentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocuments

`func (o *CreateDocumentsRequest) GetDocuments() []DocumentWriteDTO`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *CreateDocumentsRequest) GetDocumentsOk() (*[]DocumentWriteDTO, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *CreateDocumentsRequest) SetDocuments(v []DocumentWriteDTO)`

SetDocuments sets Documents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


