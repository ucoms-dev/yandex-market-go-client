# CommentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Идентификатор комментария к ответу.  | 
**Text** | **string** | Текстовое содержимое.  | 
**CanModify** | Pointer to **bool** | Может ли продавец изменять комментарий или удалять его. | [optional] 
**ParentId** | Pointer to **int64** | Идентификатор комментария к ответу.  | [optional] 
**Author** | Pointer to [**QuestionsTextContentAuthorDTO**](QuestionsTextContentAuthorDTO.md) |  | [optional] 
**Status** | [**QuestionsTextContentModerationStatusType**](QuestionsTextContentModerationStatusType.md) |  | 
**AnswerId** | **int64** | Идентификатор ответа на вопрос.  | 
**CreatedAt** | **time.Time** | Дата создания комментария. | 
**Votes** | Pointer to [**VotesDTO**](VotesDTO.md) |  | [optional] 

## Methods

### NewCommentDTO

`func NewCommentDTO(id int64, text string, status QuestionsTextContentModerationStatusType, answerId int64, createdAt time.Time, ) *CommentDTO`

NewCommentDTO instantiates a new CommentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentDTOWithDefaults

`func NewCommentDTOWithDefaults() *CommentDTO`

NewCommentDTOWithDefaults instantiates a new CommentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommentDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommentDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommentDTO) SetId(v int64)`

SetId sets Id field to given value.


### GetText

`func (o *CommentDTO) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CommentDTO) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CommentDTO) SetText(v string)`

SetText sets Text field to given value.


### GetCanModify

`func (o *CommentDTO) GetCanModify() bool`

GetCanModify returns the CanModify field if non-nil, zero value otherwise.

### GetCanModifyOk

`func (o *CommentDTO) GetCanModifyOk() (*bool, bool)`

GetCanModifyOk returns a tuple with the CanModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanModify

`func (o *CommentDTO) SetCanModify(v bool)`

SetCanModify sets CanModify field to given value.

### HasCanModify

`func (o *CommentDTO) HasCanModify() bool`

HasCanModify returns a boolean if a field has been set.

### GetParentId

`func (o *CommentDTO) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CommentDTO) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CommentDTO) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CommentDTO) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetAuthor

`func (o *CommentDTO) GetAuthor() QuestionsTextContentAuthorDTO`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CommentDTO) GetAuthorOk() (*QuestionsTextContentAuthorDTO, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CommentDTO) SetAuthor(v QuestionsTextContentAuthorDTO)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *CommentDTO) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetStatus

`func (o *CommentDTO) GetStatus() QuestionsTextContentModerationStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommentDTO) GetStatusOk() (*QuestionsTextContentModerationStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommentDTO) SetStatus(v QuestionsTextContentModerationStatusType)`

SetStatus sets Status field to given value.


### GetAnswerId

`func (o *CommentDTO) GetAnswerId() int64`

GetAnswerId returns the AnswerId field if non-nil, zero value otherwise.

### GetAnswerIdOk

`func (o *CommentDTO) GetAnswerIdOk() (*int64, bool)`

GetAnswerIdOk returns a tuple with the AnswerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerId

`func (o *CommentDTO) SetAnswerId(v int64)`

SetAnswerId sets AnswerId field to given value.


### GetCreatedAt

`func (o *CommentDTO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommentDTO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommentDTO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetVotes

`func (o *CommentDTO) GetVotes() VotesDTO`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *CommentDTO) GetVotesOk() (*VotesDTO, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *CommentDTO) SetVotes(v VotesDTO)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *CommentDTO) HasVotes() bool`

HasVotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


