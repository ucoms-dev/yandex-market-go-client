# GetDocumentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentIds** | Pointer to **[]int64** | Идентификаторы документов. | [optional] 
**DocumentNumbers** | Pointer to **[]string** | Номера документов. | [optional] 
**DocumentTypes** | Pointer to [**[]DocumentType**](DocumentType.md) | Типы документов. | [optional] 
**DocumentStatuses** | Pointer to [**[]DocumentStatusType**](DocumentStatusType.md) | Статусы документов. | [optional] 

## Methods

### NewGetDocumentsRequest

`func NewGetDocumentsRequest() *GetDocumentsRequest`

NewGetDocumentsRequest instantiates a new GetDocumentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDocumentsRequestWithDefaults

`func NewGetDocumentsRequestWithDefaults() *GetDocumentsRequest`

NewGetDocumentsRequestWithDefaults instantiates a new GetDocumentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentIds

`func (o *GetDocumentsRequest) GetDocumentIds() []int64`

GetDocumentIds returns the DocumentIds field if non-nil, zero value otherwise.

### GetDocumentIdsOk

`func (o *GetDocumentsRequest) GetDocumentIdsOk() (*[]int64, bool)`

GetDocumentIdsOk returns a tuple with the DocumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIds

`func (o *GetDocumentsRequest) SetDocumentIds(v []int64)`

SetDocumentIds sets DocumentIds field to given value.

### HasDocumentIds

`func (o *GetDocumentsRequest) HasDocumentIds() bool`

HasDocumentIds returns a boolean if a field has been set.

### SetDocumentIdsNil

`func (o *GetDocumentsRequest) SetDocumentIdsNil(b bool)`

 SetDocumentIdsNil sets the value for DocumentIds to be an explicit nil

### UnsetDocumentIds
`func (o *GetDocumentsRequest) UnsetDocumentIds()`

UnsetDocumentIds ensures that no value is present for DocumentIds, not even an explicit nil
### GetDocumentNumbers

`func (o *GetDocumentsRequest) GetDocumentNumbers() []string`

GetDocumentNumbers returns the DocumentNumbers field if non-nil, zero value otherwise.

### GetDocumentNumbersOk

`func (o *GetDocumentsRequest) GetDocumentNumbersOk() (*[]string, bool)`

GetDocumentNumbersOk returns a tuple with the DocumentNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumbers

`func (o *GetDocumentsRequest) SetDocumentNumbers(v []string)`

SetDocumentNumbers sets DocumentNumbers field to given value.

### HasDocumentNumbers

`func (o *GetDocumentsRequest) HasDocumentNumbers() bool`

HasDocumentNumbers returns a boolean if a field has been set.

### SetDocumentNumbersNil

`func (o *GetDocumentsRequest) SetDocumentNumbersNil(b bool)`

 SetDocumentNumbersNil sets the value for DocumentNumbers to be an explicit nil

### UnsetDocumentNumbers
`func (o *GetDocumentsRequest) UnsetDocumentNumbers()`

UnsetDocumentNumbers ensures that no value is present for DocumentNumbers, not even an explicit nil
### GetDocumentTypes

`func (o *GetDocumentsRequest) GetDocumentTypes() []DocumentType`

GetDocumentTypes returns the DocumentTypes field if non-nil, zero value otherwise.

### GetDocumentTypesOk

`func (o *GetDocumentsRequest) GetDocumentTypesOk() (*[]DocumentType, bool)`

GetDocumentTypesOk returns a tuple with the DocumentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypes

`func (o *GetDocumentsRequest) SetDocumentTypes(v []DocumentType)`

SetDocumentTypes sets DocumentTypes field to given value.

### HasDocumentTypes

`func (o *GetDocumentsRequest) HasDocumentTypes() bool`

HasDocumentTypes returns a boolean if a field has been set.

### SetDocumentTypesNil

`func (o *GetDocumentsRequest) SetDocumentTypesNil(b bool)`

 SetDocumentTypesNil sets the value for DocumentTypes to be an explicit nil

### UnsetDocumentTypes
`func (o *GetDocumentsRequest) UnsetDocumentTypes()`

UnsetDocumentTypes ensures that no value is present for DocumentTypes, not even an explicit nil
### GetDocumentStatuses

`func (o *GetDocumentsRequest) GetDocumentStatuses() []DocumentStatusType`

GetDocumentStatuses returns the DocumentStatuses field if non-nil, zero value otherwise.

### GetDocumentStatusesOk

`func (o *GetDocumentsRequest) GetDocumentStatusesOk() (*[]DocumentStatusType, bool)`

GetDocumentStatusesOk returns a tuple with the DocumentStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentStatuses

`func (o *GetDocumentsRequest) SetDocumentStatuses(v []DocumentStatusType)`

SetDocumentStatuses sets DocumentStatuses field to given value.

### HasDocumentStatuses

`func (o *GetDocumentsRequest) HasDocumentStatuses() bool`

HasDocumentStatuses returns a boolean if a field has been set.

### SetDocumentStatusesNil

`func (o *GetDocumentsRequest) SetDocumentStatusesNil(b bool)`

 SetDocumentStatusesNil sets the value for DocumentStatuses to be an explicit nil

### UnsetDocumentStatuses
`func (o *GetDocumentsRequest) UnsetDocumentStatuses()`

UnsetDocumentStatuses ensures that no value is present for DocumentStatuses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


