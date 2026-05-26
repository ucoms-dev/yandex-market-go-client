# QuestionListDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Questions** | [**[]QuestionDTO**](QuestionDTO.md) | Список вопросов. | 
**Paging** | Pointer to [**PackagingForwardScrollingPagerDTO**](PackagingForwardScrollingPagerDTO.md) |  | [optional] 
**TotalCount** | **int64** | Общее количество вопросов, которые попадают под фильтр.  | 

## Methods

### NewQuestionListDTO

`func NewQuestionListDTO(questions []QuestionDTO, totalCount int64, ) *QuestionListDTO`

NewQuestionListDTO instantiates a new QuestionListDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionListDTOWithDefaults

`func NewQuestionListDTOWithDefaults() *QuestionListDTO`

NewQuestionListDTOWithDefaults instantiates a new QuestionListDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestions

`func (o *QuestionListDTO) GetQuestions() []QuestionDTO`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *QuestionListDTO) GetQuestionsOk() (*[]QuestionDTO, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *QuestionListDTO) SetQuestions(v []QuestionDTO)`

SetQuestions sets Questions field to given value.


### GetPaging

`func (o *QuestionListDTO) GetPaging() PackagingForwardScrollingPagerDTO`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *QuestionListDTO) GetPagingOk() (*PackagingForwardScrollingPagerDTO, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *QuestionListDTO) SetPaging(v PackagingForwardScrollingPagerDTO)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *QuestionListDTO) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetTotalCount

`func (o *QuestionListDTO) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *QuestionListDTO) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *QuestionListDTO) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


