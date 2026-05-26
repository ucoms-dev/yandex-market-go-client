# AnswerDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Идентификатор ответа на вопрос.  | 
**Text** | **string** | Текстовое содержимое.  | 
**CanModify** | **bool** | Может ли продавец изменять комментарий или удалять его. | 
**Author** | Pointer to [**QuestionsTextContentAuthorDTO**](QuestionsTextContentAuthorDTO.md) |  | [optional] 
**Status** | [**QuestionsTextContentModerationStatusType**](QuestionsTextContentModerationStatusType.md) |  | 
**QuestionId** | **int64** | Идентификатор вопроса.  | 
**CreatedAt** | **time.Time** | Дата и время создания ответа. | 
**Votes** | [**VotesDTO**](VotesDTO.md) |  | 
**Comments** | Pointer to [**[]CommentDTO**](CommentDTO.md) | Список комментариев. | [optional] 

## Methods

### NewAnswerDTO

`func NewAnswerDTO(id int64, text string, canModify bool, status QuestionsTextContentModerationStatusType, questionId int64, createdAt time.Time, votes VotesDTO, ) *AnswerDTO`

NewAnswerDTO instantiates a new AnswerDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerDTOWithDefaults

`func NewAnswerDTOWithDefaults() *AnswerDTO`

NewAnswerDTOWithDefaults instantiates a new AnswerDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AnswerDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnswerDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnswerDTO) SetId(v int64)`

SetId sets Id field to given value.


### GetText

`func (o *AnswerDTO) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AnswerDTO) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AnswerDTO) SetText(v string)`

SetText sets Text field to given value.


### GetCanModify

`func (o *AnswerDTO) GetCanModify() bool`

GetCanModify returns the CanModify field if non-nil, zero value otherwise.

### GetCanModifyOk

`func (o *AnswerDTO) GetCanModifyOk() (*bool, bool)`

GetCanModifyOk returns a tuple with the CanModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanModify

`func (o *AnswerDTO) SetCanModify(v bool)`

SetCanModify sets CanModify field to given value.


### GetAuthor

`func (o *AnswerDTO) GetAuthor() QuestionsTextContentAuthorDTO`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AnswerDTO) GetAuthorOk() (*QuestionsTextContentAuthorDTO, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AnswerDTO) SetAuthor(v QuestionsTextContentAuthorDTO)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *AnswerDTO) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetStatus

`func (o *AnswerDTO) GetStatus() QuestionsTextContentModerationStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnswerDTO) GetStatusOk() (*QuestionsTextContentModerationStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnswerDTO) SetStatus(v QuestionsTextContentModerationStatusType)`

SetStatus sets Status field to given value.


### GetQuestionId

`func (o *AnswerDTO) GetQuestionId() int64`

GetQuestionId returns the QuestionId field if non-nil, zero value otherwise.

### GetQuestionIdOk

`func (o *AnswerDTO) GetQuestionIdOk() (*int64, bool)`

GetQuestionIdOk returns a tuple with the QuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionId

`func (o *AnswerDTO) SetQuestionId(v int64)`

SetQuestionId sets QuestionId field to given value.


### GetCreatedAt

`func (o *AnswerDTO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AnswerDTO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AnswerDTO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetVotes

`func (o *AnswerDTO) GetVotes() VotesDTO`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *AnswerDTO) GetVotesOk() (*VotesDTO, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *AnswerDTO) SetVotes(v VotesDTO)`

SetVotes sets Votes field to given value.


### GetComments

`func (o *AnswerDTO) GetComments() []CommentDTO`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *AnswerDTO) GetCommentsOk() (*[]CommentDTO, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *AnswerDTO) SetComments(v []CommentDTO)`

SetComments sets Comments field to given value.

### HasComments

`func (o *AnswerDTO) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *AnswerDTO) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *AnswerDTO) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


