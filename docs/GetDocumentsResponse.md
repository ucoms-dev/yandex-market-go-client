# GetDocumentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ApiResponseStatusType**](ApiResponseStatusType.md) |  | 
**Result** | [**GetDocumentsResultDTO**](GetDocumentsResultDTO.md) |  | 

## Methods

### NewGetDocumentsResponse

`func NewGetDocumentsResponse(status ApiResponseStatusType, result GetDocumentsResultDTO, ) *GetDocumentsResponse`

NewGetDocumentsResponse instantiates a new GetDocumentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDocumentsResponseWithDefaults

`func NewGetDocumentsResponseWithDefaults() *GetDocumentsResponse`

NewGetDocumentsResponseWithDefaults instantiates a new GetDocumentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetDocumentsResponse) GetStatus() ApiResponseStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDocumentsResponse) GetStatusOk() (*ApiResponseStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDocumentsResponse) SetStatus(v ApiResponseStatusType)`

SetStatus sets Status field to given value.


### GetResult

`func (o *GetDocumentsResponse) GetResult() GetDocumentsResultDTO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDocumentsResponse) GetResultOk() (*GetDocumentsResultDTO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDocumentsResponse) SetResult(v GetDocumentsResultDTO)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


