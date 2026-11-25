# BusinessDocumentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | Номер, указанный в сертификате, декларации или другом документе. | 
**Type** | [**DocumentType**](DocumentType.md) |  | 
**ActiveFromDate** | Pointer to **string** | Дата начала действия документа. | [optional] 
**ActiveToDate** | Pointer to **string** | Дата окончания действия документа. | [optional] 
**Id** | **int64** | Идентификатор документа. | 
**Status** | [**DocumentStatusType**](DocumentStatusType.md) |  | 

## Methods

### NewBusinessDocumentDTO

`func NewBusinessDocumentDTO(number string, type_ DocumentType, id int64, status DocumentStatusType, ) *BusinessDocumentDTO`

NewBusinessDocumentDTO instantiates a new BusinessDocumentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessDocumentDTOWithDefaults

`func NewBusinessDocumentDTOWithDefaults() *BusinessDocumentDTO`

NewBusinessDocumentDTOWithDefaults instantiates a new BusinessDocumentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *BusinessDocumentDTO) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *BusinessDocumentDTO) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *BusinessDocumentDTO) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetType

`func (o *BusinessDocumentDTO) GetType() DocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BusinessDocumentDTO) GetTypeOk() (*DocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BusinessDocumentDTO) SetType(v DocumentType)`

SetType sets Type field to given value.


### GetActiveFromDate

`func (o *BusinessDocumentDTO) GetActiveFromDate() string`

GetActiveFromDate returns the ActiveFromDate field if non-nil, zero value otherwise.

### GetActiveFromDateOk

`func (o *BusinessDocumentDTO) GetActiveFromDateOk() (*string, bool)`

GetActiveFromDateOk returns a tuple with the ActiveFromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFromDate

`func (o *BusinessDocumentDTO) SetActiveFromDate(v string)`

SetActiveFromDate sets ActiveFromDate field to given value.

### HasActiveFromDate

`func (o *BusinessDocumentDTO) HasActiveFromDate() bool`

HasActiveFromDate returns a boolean if a field has been set.

### GetActiveToDate

`func (o *BusinessDocumentDTO) GetActiveToDate() string`

GetActiveToDate returns the ActiveToDate field if non-nil, zero value otherwise.

### GetActiveToDateOk

`func (o *BusinessDocumentDTO) GetActiveToDateOk() (*string, bool)`

GetActiveToDateOk returns a tuple with the ActiveToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveToDate

`func (o *BusinessDocumentDTO) SetActiveToDate(v string)`

SetActiveToDate sets ActiveToDate field to given value.

### HasActiveToDate

`func (o *BusinessDocumentDTO) HasActiveToDate() bool`

HasActiveToDate returns a boolean if a field has been set.

### GetId

`func (o *BusinessDocumentDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BusinessDocumentDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BusinessDocumentDTO) SetId(v int64)`

SetId sets Id field to given value.


### GetStatus

`func (o *BusinessDocumentDTO) GetStatus() DocumentStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BusinessDocumentDTO) GetStatusOk() (*DocumentStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BusinessDocumentDTO) SetStatus(v DocumentStatusType)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


