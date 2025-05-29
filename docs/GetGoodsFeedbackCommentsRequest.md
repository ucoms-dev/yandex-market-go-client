# GetGoodsFeedbackCommentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedbackId** | Pointer to **int64** | Идентификатор отзыва.  | [optional] 
**CommentIds** | Pointer to **[]int64** | Идентификаторы комментариев.  ⚠️ Не используйте это поле одновременно с другими фильтрами. Если вы хотите воспользоваться ими, оставьте поле пустым.  | [optional] 

## Methods

### NewGetGoodsFeedbackCommentsRequest

`func NewGetGoodsFeedbackCommentsRequest() *GetGoodsFeedbackCommentsRequest`

NewGetGoodsFeedbackCommentsRequest instantiates a new GetGoodsFeedbackCommentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGoodsFeedbackCommentsRequestWithDefaults

`func NewGetGoodsFeedbackCommentsRequestWithDefaults() *GetGoodsFeedbackCommentsRequest`

NewGetGoodsFeedbackCommentsRequestWithDefaults instantiates a new GetGoodsFeedbackCommentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedbackId

`func (o *GetGoodsFeedbackCommentsRequest) GetFeedbackId() int64`

GetFeedbackId returns the FeedbackId field if non-nil, zero value otherwise.

### GetFeedbackIdOk

`func (o *GetGoodsFeedbackCommentsRequest) GetFeedbackIdOk() (*int64, bool)`

GetFeedbackIdOk returns a tuple with the FeedbackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackId

`func (o *GetGoodsFeedbackCommentsRequest) SetFeedbackId(v int64)`

SetFeedbackId sets FeedbackId field to given value.

### HasFeedbackId

`func (o *GetGoodsFeedbackCommentsRequest) HasFeedbackId() bool`

HasFeedbackId returns a boolean if a field has been set.

### GetCommentIds

`func (o *GetGoodsFeedbackCommentsRequest) GetCommentIds() []int64`

GetCommentIds returns the CommentIds field if non-nil, zero value otherwise.

### GetCommentIdsOk

`func (o *GetGoodsFeedbackCommentsRequest) GetCommentIdsOk() (*[]int64, bool)`

GetCommentIdsOk returns a tuple with the CommentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentIds

`func (o *GetGoodsFeedbackCommentsRequest) SetCommentIds(v []int64)`

SetCommentIds sets CommentIds field to given value.

### HasCommentIds

`func (o *GetGoodsFeedbackCommentsRequest) HasCommentIds() bool`

HasCommentIds returns a boolean if a field has been set.

### SetCommentIdsNil

`func (o *GetGoodsFeedbackCommentsRequest) SetCommentIdsNil(b bool)`

 SetCommentIdsNil sets the value for CommentIds to be an explicit nil

### UnsetCommentIds
`func (o *GetGoodsFeedbackCommentsRequest) UnsetCommentIds()`

UnsetCommentIds ensures that no value is present for CommentIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


