# SuggestSizeTablesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ApiResponseStatusType**](ApiResponseStatusType.md) |  | 
**Result** | Pointer to [**SuggestSizeTablesResultDTO**](SuggestSizeTablesResultDTO.md) |  | [optional] 

## Methods

### NewSuggestSizeTablesResponse

`func NewSuggestSizeTablesResponse(status ApiResponseStatusType, ) *SuggestSizeTablesResponse`

NewSuggestSizeTablesResponse instantiates a new SuggestSizeTablesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestSizeTablesResponseWithDefaults

`func NewSuggestSizeTablesResponseWithDefaults() *SuggestSizeTablesResponse`

NewSuggestSizeTablesResponseWithDefaults instantiates a new SuggestSizeTablesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SuggestSizeTablesResponse) GetStatus() ApiResponseStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SuggestSizeTablesResponse) GetStatusOk() (*ApiResponseStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SuggestSizeTablesResponse) SetStatus(v ApiResponseStatusType)`

SetStatus sets Status field to given value.


### GetResult

`func (o *SuggestSizeTablesResponse) GetResult() SuggestSizeTablesResultDTO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SuggestSizeTablesResponse) GetResultOk() (*SuggestSizeTablesResultDTO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SuggestSizeTablesResponse) SetResult(v SuggestSizeTablesResultDTO)`

SetResult sets Result field to given value.

### HasResult

`func (o *SuggestSizeTablesResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


