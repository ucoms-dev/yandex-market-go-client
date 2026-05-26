# GetAnswersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuestionId** | Pointer to **int64** | Идентификатор вопроса.  | [optional] 
**AnswerIds** | Pointer to **[]int64** | Идентификаторы ответов.  | [optional] 

## Methods

### NewGetAnswersRequest

`func NewGetAnswersRequest() *GetAnswersRequest`

NewGetAnswersRequest instantiates a new GetAnswersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnswersRequestWithDefaults

`func NewGetAnswersRequestWithDefaults() *GetAnswersRequest`

NewGetAnswersRequestWithDefaults instantiates a new GetAnswersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestionId

`func (o *GetAnswersRequest) GetQuestionId() int64`

GetQuestionId returns the QuestionId field if non-nil, zero value otherwise.

### GetQuestionIdOk

`func (o *GetAnswersRequest) GetQuestionIdOk() (*int64, bool)`

GetQuestionIdOk returns a tuple with the QuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionId

`func (o *GetAnswersRequest) SetQuestionId(v int64)`

SetQuestionId sets QuestionId field to given value.

### HasQuestionId

`func (o *GetAnswersRequest) HasQuestionId() bool`

HasQuestionId returns a boolean if a field has been set.

### GetAnswerIds

`func (o *GetAnswersRequest) GetAnswerIds() []int64`

GetAnswerIds returns the AnswerIds field if non-nil, zero value otherwise.

### GetAnswerIdsOk

`func (o *GetAnswersRequest) GetAnswerIdsOk() (*[]int64, bool)`

GetAnswerIdsOk returns a tuple with the AnswerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerIds

`func (o *GetAnswersRequest) SetAnswerIds(v []int64)`

SetAnswerIds sets AnswerIds field to given value.

### HasAnswerIds

`func (o *GetAnswersRequest) HasAnswerIds() bool`

HasAnswerIds returns a boolean if a field has been set.

### SetAnswerIdsNil

`func (o *GetAnswersRequest) SetAnswerIdsNil(b bool)`

 SetAnswerIdsNil sets the value for AnswerIds to be an explicit nil

### UnsetAnswerIds
`func (o *GetAnswersRequest) UnsetAnswerIds()`

UnsetAnswerIds ensures that no value is present for AnswerIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


