# OfferCardRecommendationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**OfferCardRecommendationType**](OfferCardRecommendationType.md) |  | 
**Percent** | Pointer to **int32** | Процент выполнения рекомендации.  Указывается для рекомендаций некоторых типов:  * &#x60;PICTURE_COUNT&#x60;. * &#x60;VIDEO_COUNT&#x60;. * &#x60;MAIN&#x60;. * &#x60;ADDITIONAL&#x60;. * &#x60;DISTINCTIVE&#x60;.  | [optional] 
**RemainingRatingPoints** | Pointer to **int32** | Максимальное количество баллов рейтинга карточки, которые можно получить за выполнение рекомендаций.  | [optional] 

## Methods

### NewOfferCardRecommendationDTO

`func NewOfferCardRecommendationDTO(type_ OfferCardRecommendationType, ) *OfferCardRecommendationDTO`

NewOfferCardRecommendationDTO instantiates a new OfferCardRecommendationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferCardRecommendationDTOWithDefaults

`func NewOfferCardRecommendationDTOWithDefaults() *OfferCardRecommendationDTO`

NewOfferCardRecommendationDTOWithDefaults instantiates a new OfferCardRecommendationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OfferCardRecommendationDTO) GetType() OfferCardRecommendationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OfferCardRecommendationDTO) GetTypeOk() (*OfferCardRecommendationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OfferCardRecommendationDTO) SetType(v OfferCardRecommendationType)`

SetType sets Type field to given value.


### GetPercent

`func (o *OfferCardRecommendationDTO) GetPercent() int32`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *OfferCardRecommendationDTO) GetPercentOk() (*int32, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *OfferCardRecommendationDTO) SetPercent(v int32)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *OfferCardRecommendationDTO) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetRemainingRatingPoints

`func (o *OfferCardRecommendationDTO) GetRemainingRatingPoints() int32`

GetRemainingRatingPoints returns the RemainingRatingPoints field if non-nil, zero value otherwise.

### GetRemainingRatingPointsOk

`func (o *OfferCardRecommendationDTO) GetRemainingRatingPointsOk() (*int32, bool)`

GetRemainingRatingPointsOk returns a tuple with the RemainingRatingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingRatingPoints

`func (o *OfferCardRecommendationDTO) SetRemainingRatingPoints(v int32)`

SetRemainingRatingPoints sets RemainingRatingPoints field to given value.

### HasRemainingRatingPoints

`func (o *OfferCardRecommendationDTO) HasRemainingRatingPoints() bool`

HasRemainingRatingPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


