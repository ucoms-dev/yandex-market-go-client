# AnswerListDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answers** | [**[]AnswerDTO**](AnswerDTO.md) | Список ответов. | 
**Paging** | Pointer to [**PackagingForwardScrollingPagerDTO**](PackagingForwardScrollingPagerDTO.md) |  | [optional] 

## Methods

### NewAnswerListDTO

`func NewAnswerListDTO(answers []AnswerDTO, ) *AnswerListDTO`

NewAnswerListDTO instantiates a new AnswerListDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerListDTOWithDefaults

`func NewAnswerListDTOWithDefaults() *AnswerListDTO`

NewAnswerListDTOWithDefaults instantiates a new AnswerListDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswers

`func (o *AnswerListDTO) GetAnswers() []AnswerDTO`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *AnswerListDTO) GetAnswersOk() (*[]AnswerDTO, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *AnswerListDTO) SetAnswers(v []AnswerDTO)`

SetAnswers sets Answers field to given value.


### GetPaging

`func (o *AnswerListDTO) GetPaging() PackagingForwardScrollingPagerDTO`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *AnswerListDTO) GetPagingOk() (*PackagingForwardScrollingPagerDTO, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *AnswerListDTO) SetPaging(v PackagingForwardScrollingPagerDTO)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *AnswerListDTO) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


