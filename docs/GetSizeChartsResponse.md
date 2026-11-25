# GetSizeChartsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ApiResponseStatusType**](ApiResponseStatusType.md) |  | 
**Result** | [**GetSizeChartsResultDTO**](GetSizeChartsResultDTO.md) |  | 

## Methods

### NewGetSizeChartsResponse

`func NewGetSizeChartsResponse(status ApiResponseStatusType, result GetSizeChartsResultDTO, ) *GetSizeChartsResponse`

NewGetSizeChartsResponse instantiates a new GetSizeChartsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSizeChartsResponseWithDefaults

`func NewGetSizeChartsResponseWithDefaults() *GetSizeChartsResponse`

NewGetSizeChartsResponseWithDefaults instantiates a new GetSizeChartsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetSizeChartsResponse) GetStatus() ApiResponseStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSizeChartsResponse) GetStatusOk() (*ApiResponseStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSizeChartsResponse) SetStatus(v ApiResponseStatusType)`

SetStatus sets Status field to given value.


### GetResult

`func (o *GetSizeChartsResponse) GetResult() GetSizeChartsResultDTO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetSizeChartsResponse) GetResultOk() (*GetSizeChartsResultDTO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetSizeChartsResponse) SetResult(v GetSizeChartsResultDTO)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


