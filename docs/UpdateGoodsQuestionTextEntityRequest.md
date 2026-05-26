# UpdateGoodsQuestionTextEntityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to [**TypedQuestionsTextEntityIdDTO**](TypedQuestionsTextEntityIdDTO.md) |  | [optional] 
**ParentEntityId** | Pointer to [**TypedQuestionsTextEntityIdDTO**](TypedQuestionsTextEntityIdDTO.md) |  | [optional] 
**Text** | Pointer to **string** | Текстовое содержимое.  | [optional] 
**OperationType** | [**QuestionsTextEntityOperationType**](QuestionsTextEntityOperationType.md) |  | 

## Methods

### NewUpdateGoodsQuestionTextEntityRequest

`func NewUpdateGoodsQuestionTextEntityRequest(operationType QuestionsTextEntityOperationType, ) *UpdateGoodsQuestionTextEntityRequest`

NewUpdateGoodsQuestionTextEntityRequest instantiates a new UpdateGoodsQuestionTextEntityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGoodsQuestionTextEntityRequestWithDefaults

`func NewUpdateGoodsQuestionTextEntityRequestWithDefaults() *UpdateGoodsQuestionTextEntityRequest`

NewUpdateGoodsQuestionTextEntityRequestWithDefaults instantiates a new UpdateGoodsQuestionTextEntityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *UpdateGoodsQuestionTextEntityRequest) GetEntityId() TypedQuestionsTextEntityIdDTO`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *UpdateGoodsQuestionTextEntityRequest) GetEntityIdOk() (*TypedQuestionsTextEntityIdDTO, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *UpdateGoodsQuestionTextEntityRequest) SetEntityId(v TypedQuestionsTextEntityIdDTO)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *UpdateGoodsQuestionTextEntityRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetParentEntityId

`func (o *UpdateGoodsQuestionTextEntityRequest) GetParentEntityId() TypedQuestionsTextEntityIdDTO`

GetParentEntityId returns the ParentEntityId field if non-nil, zero value otherwise.

### GetParentEntityIdOk

`func (o *UpdateGoodsQuestionTextEntityRequest) GetParentEntityIdOk() (*TypedQuestionsTextEntityIdDTO, bool)`

GetParentEntityIdOk returns a tuple with the ParentEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEntityId

`func (o *UpdateGoodsQuestionTextEntityRequest) SetParentEntityId(v TypedQuestionsTextEntityIdDTO)`

SetParentEntityId sets ParentEntityId field to given value.

### HasParentEntityId

`func (o *UpdateGoodsQuestionTextEntityRequest) HasParentEntityId() bool`

HasParentEntityId returns a boolean if a field has been set.

### GetText

`func (o *UpdateGoodsQuestionTextEntityRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *UpdateGoodsQuestionTextEntityRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *UpdateGoodsQuestionTextEntityRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *UpdateGoodsQuestionTextEntityRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetOperationType

`func (o *UpdateGoodsQuestionTextEntityRequest) GetOperationType() QuestionsTextEntityOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *UpdateGoodsQuestionTextEntityRequest) GetOperationTypeOk() (*QuestionsTextEntityOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *UpdateGoodsQuestionTextEntityRequest) SetOperationType(v QuestionsTextEntityOperationType)`

SetOperationType sets OperationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


