# GetDocumentsResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documents** | [**[]BusinessDocumentDTO**](BusinessDocumentDTO.md) | Документы на странице. | 
**Paging** | [**PackagingForwardScrollingPagerDTO**](PackagingForwardScrollingPagerDTO.md) |  | 

## Methods

### NewGetDocumentsResultDTO

`func NewGetDocumentsResultDTO(documents []BusinessDocumentDTO, paging PackagingForwardScrollingPagerDTO, ) *GetDocumentsResultDTO`

NewGetDocumentsResultDTO instantiates a new GetDocumentsResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDocumentsResultDTOWithDefaults

`func NewGetDocumentsResultDTOWithDefaults() *GetDocumentsResultDTO`

NewGetDocumentsResultDTOWithDefaults instantiates a new GetDocumentsResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocuments

`func (o *GetDocumentsResultDTO) GetDocuments() []BusinessDocumentDTO`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *GetDocumentsResultDTO) GetDocumentsOk() (*[]BusinessDocumentDTO, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *GetDocumentsResultDTO) SetDocuments(v []BusinessDocumentDTO)`

SetDocuments sets Documents field to given value.


### GetPaging

`func (o *GetDocumentsResultDTO) GetPaging() PackagingForwardScrollingPagerDTO`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *GetDocumentsResultDTO) GetPagingOk() (*PackagingForwardScrollingPagerDTO, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *GetDocumentsResultDTO) SetPaging(v PackagingForwardScrollingPagerDTO)`

SetPaging sets Paging field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


