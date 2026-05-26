# GetAnswersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ApiResponseStatusType**](ApiResponseStatusType.md) |  | 
**Result** | Pointer to [**AnswerListDTO**](AnswerListDTO.md) |  | [optional] 

## Methods

### NewGetAnswersResponse

`func NewGetAnswersResponse(status ApiResponseStatusType, ) *GetAnswersResponse`

NewGetAnswersResponse instantiates a new GetAnswersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnswersResponseWithDefaults

`func NewGetAnswersResponseWithDefaults() *GetAnswersResponse`

NewGetAnswersResponseWithDefaults instantiates a new GetAnswersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetAnswersResponse) GetStatus() ApiResponseStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAnswersResponse) GetStatusOk() (*ApiResponseStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAnswersResponse) SetStatus(v ApiResponseStatusType)`

SetStatus sets Status field to given value.


### GetResult

`func (o *GetAnswersResponse) GetResult() AnswerListDTO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetAnswersResponse) GetResultOk() (*AnswerListDTO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetAnswersResponse) SetResult(v AnswerListDTO)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetAnswersResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


