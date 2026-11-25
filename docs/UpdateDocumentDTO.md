# UpdateDocumentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | Номер, указанный в сертификате, декларации или другом документе. | 
**Type** | [**DocumentType**](DocumentType.md) |  | 
**ActiveFromDate** | Pointer to **string** | Дата начала действия документа. | [optional] 
**ActiveToDate** | Pointer to **string** | Дата окончания действия документа. | [optional] 
**Id** | **int64** | Идентификатор документа. | 

## Methods

### NewUpdateDocumentDTO

`func NewUpdateDocumentDTO(number string, type_ DocumentType, id int64, ) *UpdateDocumentDTO`

NewUpdateDocumentDTO instantiates a new UpdateDocumentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDocumentDTOWithDefaults

`func NewUpdateDocumentDTOWithDefaults() *UpdateDocumentDTO`

NewUpdateDocumentDTOWithDefaults instantiates a new UpdateDocumentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *UpdateDocumentDTO) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *UpdateDocumentDTO) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *UpdateDocumentDTO) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetType

`func (o *UpdateDocumentDTO) GetType() DocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateDocumentDTO) GetTypeOk() (*DocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateDocumentDTO) SetType(v DocumentType)`

SetType sets Type field to given value.


### GetActiveFromDate

`func (o *UpdateDocumentDTO) GetActiveFromDate() string`

GetActiveFromDate returns the ActiveFromDate field if non-nil, zero value otherwise.

### GetActiveFromDateOk

`func (o *UpdateDocumentDTO) GetActiveFromDateOk() (*string, bool)`

GetActiveFromDateOk returns a tuple with the ActiveFromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFromDate

`func (o *UpdateDocumentDTO) SetActiveFromDate(v string)`

SetActiveFromDate sets ActiveFromDate field to given value.

### HasActiveFromDate

`func (o *UpdateDocumentDTO) HasActiveFromDate() bool`

HasActiveFromDate returns a boolean if a field has been set.

### GetActiveToDate

`func (o *UpdateDocumentDTO) GetActiveToDate() string`

GetActiveToDate returns the ActiveToDate field if non-nil, zero value otherwise.

### GetActiveToDateOk

`func (o *UpdateDocumentDTO) GetActiveToDateOk() (*string, bool)`

GetActiveToDateOk returns a tuple with the ActiveToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveToDate

`func (o *UpdateDocumentDTO) SetActiveToDate(v string)`

SetActiveToDate sets ActiveToDate field to given value.

### HasActiveToDate

`func (o *UpdateDocumentDTO) HasActiveToDate() bool`

HasActiveToDate returns a boolean if a field has been set.

### GetId

`func (o *UpdateDocumentDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateDocumentDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateDocumentDTO) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


