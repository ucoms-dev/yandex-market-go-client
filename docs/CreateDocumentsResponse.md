# CreateDocumentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ApiResponseStatusType**](ApiResponseStatusType.md) |  | 
**Result** | [**CreateDocumentsResultDTO**](CreateDocumentsResultDTO.md) |  | 

## Methods

### NewCreateDocumentsResponse

`func NewCreateDocumentsResponse(status ApiResponseStatusType, result CreateDocumentsResultDTO, ) *CreateDocumentsResponse`

NewCreateDocumentsResponse instantiates a new CreateDocumentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDocumentsResponseWithDefaults

`func NewCreateDocumentsResponseWithDefaults() *CreateDocumentsResponse`

NewCreateDocumentsResponseWithDefaults instantiates a new CreateDocumentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateDocumentsResponse) GetStatus() ApiResponseStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateDocumentsResponse) GetStatusOk() (*ApiResponseStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateDocumentsResponse) SetStatus(v ApiResponseStatusType)`

SetStatus sets Status field to given value.


### GetResult

`func (o *CreateDocumentsResponse) GetResult() CreateDocumentsResultDTO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CreateDocumentsResponse) GetResultOk() (*CreateDocumentsResultDTO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CreateDocumentsResponse) SetResult(v CreateDocumentsResultDTO)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


