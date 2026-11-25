# DocumentWriteDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | Номер, указанный в сертификате, декларации или другом документе. | 
**Type** | [**DocumentType**](DocumentType.md) |  | 
**ActiveFromDate** | Pointer to **string** | Дата начала действия документа. | [optional] 
**ActiveToDate** | Pointer to **string** | Дата окончания действия документа. | [optional] 

## Methods

### NewDocumentWriteDTO

`func NewDocumentWriteDTO(number string, type_ DocumentType, ) *DocumentWriteDTO`

NewDocumentWriteDTO instantiates a new DocumentWriteDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWriteDTOWithDefaults

`func NewDocumentWriteDTOWithDefaults() *DocumentWriteDTO`

NewDocumentWriteDTOWithDefaults instantiates a new DocumentWriteDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *DocumentWriteDTO) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *DocumentWriteDTO) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *DocumentWriteDTO) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetType

`func (o *DocumentWriteDTO) GetType() DocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DocumentWriteDTO) GetTypeOk() (*DocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DocumentWriteDTO) SetType(v DocumentType)`

SetType sets Type field to given value.


### GetActiveFromDate

`func (o *DocumentWriteDTO) GetActiveFromDate() string`

GetActiveFromDate returns the ActiveFromDate field if non-nil, zero value otherwise.

### GetActiveFromDateOk

`func (o *DocumentWriteDTO) GetActiveFromDateOk() (*string, bool)`

GetActiveFromDateOk returns a tuple with the ActiveFromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFromDate

`func (o *DocumentWriteDTO) SetActiveFromDate(v string)`

SetActiveFromDate sets ActiveFromDate field to given value.

### HasActiveFromDate

`func (o *DocumentWriteDTO) HasActiveFromDate() bool`

HasActiveFromDate returns a boolean if a field has been set.

### GetActiveToDate

`func (o *DocumentWriteDTO) GetActiveToDate() string`

GetActiveToDate returns the ActiveToDate field if non-nil, zero value otherwise.

### GetActiveToDateOk

`func (o *DocumentWriteDTO) GetActiveToDateOk() (*string, bool)`

GetActiveToDateOk returns a tuple with the ActiveToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveToDate

`func (o *DocumentWriteDTO) SetActiveToDate(v string)`

SetActiveToDate sets ActiveToDate field to given value.

### HasActiveToDate

`func (o *DocumentWriteDTO) HasActiveToDate() bool`

HasActiveToDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


