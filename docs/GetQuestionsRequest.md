# GetQuestionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryIds** | Pointer to **[]int64** | Идентификаторы категорий товаров. | [optional] 
**QuestionIds** | Pointer to **[]int64** | Идентификаторы вопросов. | [optional] 
**DateFrom** | Pointer to **string** | Дата начала периода создания вопроса.  Если параметр не указан, возвращается информация за 1 месяц до указанной в &#x60;dateTo&#x60; даты.  Максимальный интервал 1 месяц.  | [optional] 
**DateTo** | Pointer to **string** | Дата окончания периода создания вопроса.  Если параметр не указан, используется текущая дата.  Максимальный интервал 1 месяц.  | [optional] 
**NeedAnswer** | Pointer to **bool** | Нужен ли ответ на вопрос.  * &#x60;true&#x60; — только вопросы, которые ждут ответа. * &#x60;false&#x60; — все вопросы.  | [optional] [default to false]
**Sort** | Pointer to [**QuestionSortOrderType**](QuestionSortOrderType.md) |  | [optional] 

## Methods

### NewGetQuestionsRequest

`func NewGetQuestionsRequest() *GetQuestionsRequest`

NewGetQuestionsRequest instantiates a new GetQuestionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetQuestionsRequestWithDefaults

`func NewGetQuestionsRequestWithDefaults() *GetQuestionsRequest`

NewGetQuestionsRequestWithDefaults instantiates a new GetQuestionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryIds

`func (o *GetQuestionsRequest) GetCategoryIds() []int64`

GetCategoryIds returns the CategoryIds field if non-nil, zero value otherwise.

### GetCategoryIdsOk

`func (o *GetQuestionsRequest) GetCategoryIdsOk() (*[]int64, bool)`

GetCategoryIdsOk returns a tuple with the CategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryIds

`func (o *GetQuestionsRequest) SetCategoryIds(v []int64)`

SetCategoryIds sets CategoryIds field to given value.

### HasCategoryIds

`func (o *GetQuestionsRequest) HasCategoryIds() bool`

HasCategoryIds returns a boolean if a field has been set.

### SetCategoryIdsNil

`func (o *GetQuestionsRequest) SetCategoryIdsNil(b bool)`

 SetCategoryIdsNil sets the value for CategoryIds to be an explicit nil

### UnsetCategoryIds
`func (o *GetQuestionsRequest) UnsetCategoryIds()`

UnsetCategoryIds ensures that no value is present for CategoryIds, not even an explicit nil
### GetQuestionIds

`func (o *GetQuestionsRequest) GetQuestionIds() []int64`

GetQuestionIds returns the QuestionIds field if non-nil, zero value otherwise.

### GetQuestionIdsOk

`func (o *GetQuestionsRequest) GetQuestionIdsOk() (*[]int64, bool)`

GetQuestionIdsOk returns a tuple with the QuestionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionIds

`func (o *GetQuestionsRequest) SetQuestionIds(v []int64)`

SetQuestionIds sets QuestionIds field to given value.

### HasQuestionIds

`func (o *GetQuestionsRequest) HasQuestionIds() bool`

HasQuestionIds returns a boolean if a field has been set.

### SetQuestionIdsNil

`func (o *GetQuestionsRequest) SetQuestionIdsNil(b bool)`

 SetQuestionIdsNil sets the value for QuestionIds to be an explicit nil

### UnsetQuestionIds
`func (o *GetQuestionsRequest) UnsetQuestionIds()`

UnsetQuestionIds ensures that no value is present for QuestionIds, not even an explicit nil
### GetDateFrom

`func (o *GetQuestionsRequest) GetDateFrom() string`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *GetQuestionsRequest) GetDateFromOk() (*string, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *GetQuestionsRequest) SetDateFrom(v string)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *GetQuestionsRequest) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### GetDateTo

`func (o *GetQuestionsRequest) GetDateTo() string`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *GetQuestionsRequest) GetDateToOk() (*string, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *GetQuestionsRequest) SetDateTo(v string)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *GetQuestionsRequest) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.

### GetNeedAnswer

`func (o *GetQuestionsRequest) GetNeedAnswer() bool`

GetNeedAnswer returns the NeedAnswer field if non-nil, zero value otherwise.

### GetNeedAnswerOk

`func (o *GetQuestionsRequest) GetNeedAnswerOk() (*bool, bool)`

GetNeedAnswerOk returns a tuple with the NeedAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedAnswer

`func (o *GetQuestionsRequest) SetNeedAnswer(v bool)`

SetNeedAnswer sets NeedAnswer field to given value.

### HasNeedAnswer

`func (o *GetQuestionsRequest) HasNeedAnswer() bool`

HasNeedAnswer returns a boolean if a field has been set.

### GetSort

`func (o *GetQuestionsRequest) GetSort() QuestionSortOrderType`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *GetQuestionsRequest) GetSortOk() (*QuestionSortOrderType, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *GetQuestionsRequest) SetSort(v QuestionSortOrderType)`

SetSort sets Sort field to given value.

### HasSort

`func (o *GetQuestionsRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


