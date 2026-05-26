# UpdateGoodsQuestionTextEntityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ApiResponseStatusType**](ApiResponseStatusType.md) |  | 
**Result** | Pointer to [**UpdateGoodsQuestionTextEntityDTO**](UpdateGoodsQuestionTextEntityDTO.md) |  | [optional] 

## Methods

### NewUpdateGoodsQuestionTextEntityResponse

`func NewUpdateGoodsQuestionTextEntityResponse(status ApiResponseStatusType, ) *UpdateGoodsQuestionTextEntityResponse`

NewUpdateGoodsQuestionTextEntityResponse instantiates a new UpdateGoodsQuestionTextEntityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGoodsQuestionTextEntityResponseWithDefaults

`func NewUpdateGoodsQuestionTextEntityResponseWithDefaults() *UpdateGoodsQuestionTextEntityResponse`

NewUpdateGoodsQuestionTextEntityResponseWithDefaults instantiates a new UpdateGoodsQuestionTextEntityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateGoodsQuestionTextEntityResponse) GetStatus() ApiResponseStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateGoodsQuestionTextEntityResponse) GetStatusOk() (*ApiResponseStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateGoodsQuestionTextEntityResponse) SetStatus(v ApiResponseStatusType)`

SetStatus sets Status field to given value.


### GetResult

`func (o *UpdateGoodsQuestionTextEntityResponse) GetResult() UpdateGoodsQuestionTextEntityDTO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UpdateGoodsQuestionTextEntityResponse) GetResultOk() (*UpdateGoodsQuestionTextEntityDTO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UpdateGoodsQuestionTextEntityResponse) SetResult(v UpdateGoodsQuestionTextEntityDTO)`

SetResult sets Result field to given value.

### HasResult

`func (o *UpdateGoodsQuestionTextEntityResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


