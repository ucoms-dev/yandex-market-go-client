# QuestionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuestionIdentifiers** | [**QuestionIdentifiersDTO**](QuestionIdentifiersDTO.md) |  | 
**BusinessId** | **int64** | Идентификатор кабинета. {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) {% endif %}  | 
**Text** | **string** | Текстовое содержимое.  | 
**CreatedAt** | **time.Time** | Дата и время создания вопроса. | 
**Votes** | [**VotesDTO**](VotesDTO.md) |  | 
**Author** | [**QuestionsTextContentAuthorDTO**](QuestionsTextContentAuthorDTO.md) |  | 

## Methods

### NewQuestionDTO

`func NewQuestionDTO(questionIdentifiers QuestionIdentifiersDTO, businessId int64, text string, createdAt time.Time, votes VotesDTO, author QuestionsTextContentAuthorDTO, ) *QuestionDTO`

NewQuestionDTO instantiates a new QuestionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionDTOWithDefaults

`func NewQuestionDTOWithDefaults() *QuestionDTO`

NewQuestionDTOWithDefaults instantiates a new QuestionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestionIdentifiers

`func (o *QuestionDTO) GetQuestionIdentifiers() QuestionIdentifiersDTO`

GetQuestionIdentifiers returns the QuestionIdentifiers field if non-nil, zero value otherwise.

### GetQuestionIdentifiersOk

`func (o *QuestionDTO) GetQuestionIdentifiersOk() (*QuestionIdentifiersDTO, bool)`

GetQuestionIdentifiersOk returns a tuple with the QuestionIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionIdentifiers

`func (o *QuestionDTO) SetQuestionIdentifiers(v QuestionIdentifiersDTO)`

SetQuestionIdentifiers sets QuestionIdentifiers field to given value.


### GetBusinessId

`func (o *QuestionDTO) GetBusinessId() int64`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *QuestionDTO) GetBusinessIdOk() (*int64, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *QuestionDTO) SetBusinessId(v int64)`

SetBusinessId sets BusinessId field to given value.


### GetText

`func (o *QuestionDTO) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *QuestionDTO) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *QuestionDTO) SetText(v string)`

SetText sets Text field to given value.


### GetCreatedAt

`func (o *QuestionDTO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QuestionDTO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QuestionDTO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetVotes

`func (o *QuestionDTO) GetVotes() VotesDTO`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *QuestionDTO) GetVotesOk() (*VotesDTO, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *QuestionDTO) SetVotes(v VotesDTO)`

SetVotes sets Votes field to given value.


### GetAuthor

`func (o *QuestionDTO) GetAuthor() QuestionsTextContentAuthorDTO`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *QuestionDTO) GetAuthorOk() (*QuestionsTextContentAuthorDTO, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *QuestionDTO) SetAuthor(v QuestionsTextContentAuthorDTO)`

SetAuthor sets Author field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


