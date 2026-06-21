# GoodsFeedbackUrbanadsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedbackId** | **int64** | Идентификатор отзыва.  | 
**CreatedAt** | **time.Time** | Дата и время создания отзыва. | 
**NeedReaction** | **bool** | Прочитан ли отзыв.  Принимает значение &#x60;false&#x60;, если рекламодатель:  * Прочитал отзыв в кабинете UrbanAds. * Пропустил реакцию на отзыв — метод [POST v2/businesses/{businessId}/goods-feedback/skip-reaction](../../reference/goods-feedback/skipGoodsFeedbacksReaction.md). * Оставил комментарий к отзыву — метод [POST v2/businesses/{businessId}/goods-feedback/comments/update](../../reference/goods-feedback/updateGoodsFeedbackComment.md).  | 
**Context** | [**GoodsFeedbackContextUrbanadsDTO**](GoodsFeedbackContextUrbanadsDTO.md) |  | 
**Author** | Pointer to **string** | Имя автора отзыва. | [optional] 
**Description** | Pointer to [**GoodsFeedbackDescriptionDTO**](GoodsFeedbackDescriptionDTO.md) |  | [optional] 
**Media** | Pointer to [**GoodsFeedbackMediaDTO**](GoodsFeedbackMediaDTO.md) |  | [optional] 
**Statistics** | [**GoodsFeedbackStatisticsDTO**](GoodsFeedbackStatisticsDTO.md) |  | 

## Methods

### NewGoodsFeedbackUrbanadsDTO

`func NewGoodsFeedbackUrbanadsDTO(feedbackId int64, createdAt time.Time, needReaction bool, context GoodsFeedbackContextUrbanadsDTO, statistics GoodsFeedbackStatisticsDTO, ) *GoodsFeedbackUrbanadsDTO`

NewGoodsFeedbackUrbanadsDTO instantiates a new GoodsFeedbackUrbanadsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoodsFeedbackUrbanadsDTOWithDefaults

`func NewGoodsFeedbackUrbanadsDTOWithDefaults() *GoodsFeedbackUrbanadsDTO`

NewGoodsFeedbackUrbanadsDTOWithDefaults instantiates a new GoodsFeedbackUrbanadsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedbackId

`func (o *GoodsFeedbackUrbanadsDTO) GetFeedbackId() int64`

GetFeedbackId returns the FeedbackId field if non-nil, zero value otherwise.

### GetFeedbackIdOk

`func (o *GoodsFeedbackUrbanadsDTO) GetFeedbackIdOk() (*int64, bool)`

GetFeedbackIdOk returns a tuple with the FeedbackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackId

`func (o *GoodsFeedbackUrbanadsDTO) SetFeedbackId(v int64)`

SetFeedbackId sets FeedbackId field to given value.


### GetCreatedAt

`func (o *GoodsFeedbackUrbanadsDTO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GoodsFeedbackUrbanadsDTO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GoodsFeedbackUrbanadsDTO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetNeedReaction

`func (o *GoodsFeedbackUrbanadsDTO) GetNeedReaction() bool`

GetNeedReaction returns the NeedReaction field if non-nil, zero value otherwise.

### GetNeedReactionOk

`func (o *GoodsFeedbackUrbanadsDTO) GetNeedReactionOk() (*bool, bool)`

GetNeedReactionOk returns a tuple with the NeedReaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedReaction

`func (o *GoodsFeedbackUrbanadsDTO) SetNeedReaction(v bool)`

SetNeedReaction sets NeedReaction field to given value.


### GetContext

`func (o *GoodsFeedbackUrbanadsDTO) GetContext() GoodsFeedbackContextUrbanadsDTO`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GoodsFeedbackUrbanadsDTO) GetContextOk() (*GoodsFeedbackContextUrbanadsDTO, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GoodsFeedbackUrbanadsDTO) SetContext(v GoodsFeedbackContextUrbanadsDTO)`

SetContext sets Context field to given value.


### GetAuthor

`func (o *GoodsFeedbackUrbanadsDTO) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *GoodsFeedbackUrbanadsDTO) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *GoodsFeedbackUrbanadsDTO) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *GoodsFeedbackUrbanadsDTO) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetDescription

`func (o *GoodsFeedbackUrbanadsDTO) GetDescription() GoodsFeedbackDescriptionDTO`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GoodsFeedbackUrbanadsDTO) GetDescriptionOk() (*GoodsFeedbackDescriptionDTO, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GoodsFeedbackUrbanadsDTO) SetDescription(v GoodsFeedbackDescriptionDTO)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GoodsFeedbackUrbanadsDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMedia

`func (o *GoodsFeedbackUrbanadsDTO) GetMedia() GoodsFeedbackMediaDTO`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *GoodsFeedbackUrbanadsDTO) GetMediaOk() (*GoodsFeedbackMediaDTO, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *GoodsFeedbackUrbanadsDTO) SetMedia(v GoodsFeedbackMediaDTO)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *GoodsFeedbackUrbanadsDTO) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetStatistics

`func (o *GoodsFeedbackUrbanadsDTO) GetStatistics() GoodsFeedbackStatisticsDTO`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *GoodsFeedbackUrbanadsDTO) GetStatisticsOk() (*GoodsFeedbackStatisticsDTO, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *GoodsFeedbackUrbanadsDTO) SetStatistics(v GoodsFeedbackStatisticsDTO)`

SetStatistics sets Statistics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


