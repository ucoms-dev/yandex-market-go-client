# CreateDocumentsResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documents** | Pointer to [**[]BusinessDocumentDTO**](BusinessDocumentDTO.md) | Документы, созданные этим запросом. | [optional] 
**Errors** | Pointer to [**[]CreateDocumentErrorDTO**](CreateDocumentErrorDTO.md) | Ошибки документов, которые не удалось создать. | [optional] 

## Methods

### NewCreateDocumentsResultDTO

`func NewCreateDocumentsResultDTO() *CreateDocumentsResultDTO`

NewCreateDocumentsResultDTO instantiates a new CreateDocumentsResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDocumentsResultDTOWithDefaults

`func NewCreateDocumentsResultDTOWithDefaults() *CreateDocumentsResultDTO`

NewCreateDocumentsResultDTOWithDefaults instantiates a new CreateDocumentsResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocuments

`func (o *CreateDocumentsResultDTO) GetDocuments() []BusinessDocumentDTO`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *CreateDocumentsResultDTO) GetDocumentsOk() (*[]BusinessDocumentDTO, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *CreateDocumentsResultDTO) SetDocuments(v []BusinessDocumentDTO)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *CreateDocumentsResultDTO) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### SetDocumentsNil

`func (o *CreateDocumentsResultDTO) SetDocumentsNil(b bool)`

 SetDocumentsNil sets the value for Documents to be an explicit nil

### UnsetDocuments
`func (o *CreateDocumentsResultDTO) UnsetDocuments()`

UnsetDocuments ensures that no value is present for Documents, not even an explicit nil
### GetErrors

`func (o *CreateDocumentsResultDTO) GetErrors() []CreateDocumentErrorDTO`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateDocumentsResultDTO) GetErrorsOk() (*[]CreateDocumentErrorDTO, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateDocumentsResultDTO) SetErrors(v []CreateDocumentErrorDTO)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateDocumentsResultDTO) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *CreateDocumentsResultDTO) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *CreateDocumentsResultDTO) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


