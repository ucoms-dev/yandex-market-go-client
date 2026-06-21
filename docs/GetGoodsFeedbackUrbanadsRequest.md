# GetGoodsFeedbackUrbanadsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedbackIds** | Pointer to **[]int64** | Идентификаторы отзывов.  ⚠️ Не используйте это поле одновременно с другими фильтрами. Если вы хотите воспользоваться ими, оставьте поле пустым.  | [optional] 
**DateTimeFrom** | Pointer to **time.Time** | Начало периода. Не включительно.  Если параметр не указан, возвращается информация за 6 месяцев до указанной в &#x60;dateTimeTo&#x60; даты.  Максимальный интервал 6 месяцев.  | [optional] 
**DateTimeTo** | Pointer to **time.Time** | Конец периода. Не включительно.  Если параметр не указан, используется текущая дата.  Максимальный интервал 6 месяцев.  | [optional] 
**ReactionStatus** | Pointer to [**FeedbackReactionStatusType**](FeedbackReactionStatusType.md) |  | [optional] 
**RatingValues** | Pointer to **[]int32** | Оценка товара. | [optional] 
**Paid** | Pointer to **bool** | Фильтр отзывов за баллы Плюса. | [optional] 

## Methods

### NewGetGoodsFeedbackUrbanadsRequest

`func NewGetGoodsFeedbackUrbanadsRequest() *GetGoodsFeedbackUrbanadsRequest`

NewGetGoodsFeedbackUrbanadsRequest instantiates a new GetGoodsFeedbackUrbanadsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGoodsFeedbackUrbanadsRequestWithDefaults

`func NewGetGoodsFeedbackUrbanadsRequestWithDefaults() *GetGoodsFeedbackUrbanadsRequest`

NewGetGoodsFeedbackUrbanadsRequestWithDefaults instantiates a new GetGoodsFeedbackUrbanadsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedbackIds

`func (o *GetGoodsFeedbackUrbanadsRequest) GetFeedbackIds() []int64`

GetFeedbackIds returns the FeedbackIds field if non-nil, zero value otherwise.

### GetFeedbackIdsOk

`func (o *GetGoodsFeedbackUrbanadsRequest) GetFeedbackIdsOk() (*[]int64, bool)`

GetFeedbackIdsOk returns a tuple with the FeedbackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackIds

`func (o *GetGoodsFeedbackUrbanadsRequest) SetFeedbackIds(v []int64)`

SetFeedbackIds sets FeedbackIds field to given value.

### HasFeedbackIds

`func (o *GetGoodsFeedbackUrbanadsRequest) HasFeedbackIds() bool`

HasFeedbackIds returns a boolean if a field has been set.

### SetFeedbackIdsNil

`func (o *GetGoodsFeedbackUrbanadsRequest) SetFeedbackIdsNil(b bool)`

 SetFeedbackIdsNil sets the value for FeedbackIds to be an explicit nil

### UnsetFeedbackIds
`func (o *GetGoodsFeedbackUrbanadsRequest) UnsetFeedbackIds()`

UnsetFeedbackIds ensures that no value is present for FeedbackIds, not even an explicit nil
### GetDateTimeFrom

`func (o *GetGoodsFeedbackUrbanadsRequest) GetDateTimeFrom() time.Time`

GetDateTimeFrom returns the DateTimeFrom field if non-nil, zero value otherwise.

### GetDateTimeFromOk

`func (o *GetGoodsFeedbackUrbanadsRequest) GetDateTimeFromOk() (*time.Time, bool)`

GetDateTimeFromOk returns a tuple with the DateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFrom

`func (o *GetGoodsFeedbackUrbanadsRequest) SetDateTimeFrom(v time.Time)`

SetDateTimeFrom sets DateTimeFrom field to given value.

### HasDateTimeFrom

`func (o *GetGoodsFeedbackUrbanadsRequest) HasDateTimeFrom() bool`

HasDateTimeFrom returns a boolean if a field has been set.

### GetDateTimeTo

`func (o *GetGoodsFeedbackUrbanadsRequest) GetDateTimeTo() time.Time`

GetDateTimeTo returns the DateTimeTo field if non-nil, zero value otherwise.

### GetDateTimeToOk

`func (o *GetGoodsFeedbackUrbanadsRequest) GetDateTimeToOk() (*time.Time, bool)`

GetDateTimeToOk returns a tuple with the DateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeTo

`func (o *GetGoodsFeedbackUrbanadsRequest) SetDateTimeTo(v time.Time)`

SetDateTimeTo sets DateTimeTo field to given value.

### HasDateTimeTo

`func (o *GetGoodsFeedbackUrbanadsRequest) HasDateTimeTo() bool`

HasDateTimeTo returns a boolean if a field has been set.

### GetReactionStatus

`func (o *GetGoodsFeedbackUrbanadsRequest) GetReactionStatus() FeedbackReactionStatusType`

GetReactionStatus returns the ReactionStatus field if non-nil, zero value otherwise.

### GetReactionStatusOk

`func (o *GetGoodsFeedbackUrbanadsRequest) GetReactionStatusOk() (*FeedbackReactionStatusType, bool)`

GetReactionStatusOk returns a tuple with the ReactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionStatus

`func (o *GetGoodsFeedbackUrbanadsRequest) SetReactionStatus(v FeedbackReactionStatusType)`

SetReactionStatus sets ReactionStatus field to given value.

### HasReactionStatus

`func (o *GetGoodsFeedbackUrbanadsRequest) HasReactionStatus() bool`

HasReactionStatus returns a boolean if a field has been set.

### GetRatingValues

`func (o *GetGoodsFeedbackUrbanadsRequest) GetRatingValues() []int32`

GetRatingValues returns the RatingValues field if non-nil, zero value otherwise.

### GetRatingValuesOk

`func (o *GetGoodsFeedbackUrbanadsRequest) GetRatingValuesOk() (*[]int32, bool)`

GetRatingValuesOk returns a tuple with the RatingValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingValues

`func (o *GetGoodsFeedbackUrbanadsRequest) SetRatingValues(v []int32)`

SetRatingValues sets RatingValues field to given value.

### HasRatingValues

`func (o *GetGoodsFeedbackUrbanadsRequest) HasRatingValues() bool`

HasRatingValues returns a boolean if a field has been set.

### SetRatingValuesNil

`func (o *GetGoodsFeedbackUrbanadsRequest) SetRatingValuesNil(b bool)`

 SetRatingValuesNil sets the value for RatingValues to be an explicit nil

### UnsetRatingValues
`func (o *GetGoodsFeedbackUrbanadsRequest) UnsetRatingValues()`

UnsetRatingValues ensures that no value is present for RatingValues, not even an explicit nil
### GetPaid

`func (o *GetGoodsFeedbackUrbanadsRequest) GetPaid() bool`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *GetGoodsFeedbackUrbanadsRequest) GetPaidOk() (*bool, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *GetGoodsFeedbackUrbanadsRequest) SetPaid(v bool)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *GetGoodsFeedbackUrbanadsRequest) HasPaid() bool`

HasPaid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


