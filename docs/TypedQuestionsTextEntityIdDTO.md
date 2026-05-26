# TypedQuestionsTextEntityIdDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Идентификатор вопроса, ответа или комментария. | 
**Type** | [**QuestionsTextEntityType**](QuestionsTextEntityType.md) |  | 

## Methods

### NewTypedQuestionsTextEntityIdDTO

`func NewTypedQuestionsTextEntityIdDTO(id int64, type_ QuestionsTextEntityType, ) *TypedQuestionsTextEntityIdDTO`

NewTypedQuestionsTextEntityIdDTO instantiates a new TypedQuestionsTextEntityIdDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypedQuestionsTextEntityIdDTOWithDefaults

`func NewTypedQuestionsTextEntityIdDTOWithDefaults() *TypedQuestionsTextEntityIdDTO`

NewTypedQuestionsTextEntityIdDTOWithDefaults instantiates a new TypedQuestionsTextEntityIdDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TypedQuestionsTextEntityIdDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypedQuestionsTextEntityIdDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypedQuestionsTextEntityIdDTO) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *TypedQuestionsTextEntityIdDTO) GetType() QuestionsTextEntityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TypedQuestionsTextEntityIdDTO) GetTypeOk() (*QuestionsTextEntityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TypedQuestionsTextEntityIdDTO) SetType(v QuestionsTextEntityType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


