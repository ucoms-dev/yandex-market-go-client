# GoodsFeedbackUrbanadsListDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feedbacks** | [**[]GoodsFeedbackUrbanadsDTO**](GoodsFeedbackUrbanadsDTO.md) | Список отзывов. | 
**Paging** | Pointer to [**PackagingForwardScrollingPagerDTO**](PackagingForwardScrollingPagerDTO.md) |  | [optional] 

## Methods

### NewGoodsFeedbackUrbanadsListDTO

`func NewGoodsFeedbackUrbanadsListDTO(feedbacks []GoodsFeedbackUrbanadsDTO, ) *GoodsFeedbackUrbanadsListDTO`

NewGoodsFeedbackUrbanadsListDTO instantiates a new GoodsFeedbackUrbanadsListDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoodsFeedbackUrbanadsListDTOWithDefaults

`func NewGoodsFeedbackUrbanadsListDTOWithDefaults() *GoodsFeedbackUrbanadsListDTO`

NewGoodsFeedbackUrbanadsListDTOWithDefaults instantiates a new GoodsFeedbackUrbanadsListDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedbacks

`func (o *GoodsFeedbackUrbanadsListDTO) GetFeedbacks() []GoodsFeedbackUrbanadsDTO`

GetFeedbacks returns the Feedbacks field if non-nil, zero value otherwise.

### GetFeedbacksOk

`func (o *GoodsFeedbackUrbanadsListDTO) GetFeedbacksOk() (*[]GoodsFeedbackUrbanadsDTO, bool)`

GetFeedbacksOk returns a tuple with the Feedbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbacks

`func (o *GoodsFeedbackUrbanadsListDTO) SetFeedbacks(v []GoodsFeedbackUrbanadsDTO)`

SetFeedbacks sets Feedbacks field to given value.


### GetPaging

`func (o *GoodsFeedbackUrbanadsListDTO) GetPaging() PackagingForwardScrollingPagerDTO`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *GoodsFeedbackUrbanadsListDTO) GetPagingOk() (*PackagingForwardScrollingPagerDTO, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *GoodsFeedbackUrbanadsListDTO) SetPaging(v PackagingForwardScrollingPagerDTO)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *GoodsFeedbackUrbanadsListDTO) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


