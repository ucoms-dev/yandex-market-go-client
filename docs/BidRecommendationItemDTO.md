# BidRecommendationItemDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bid** | **int32** | Значение ставки. | 
**ShowPercent** | **int64** | Доля показов.  | 
**Benefits** | Pointer to [**[]BenefitType**](BenefitType.md) | Список доступных субсидий.  Чтобы получить необходимый инструмент продвижения, установите ставку, которая будет рекомендована для этого инструмента или выше.  | [optional] 

## Methods

### NewBidRecommendationItemDTO

`func NewBidRecommendationItemDTO(bid int32, showPercent int64, ) *BidRecommendationItemDTO`

NewBidRecommendationItemDTO instantiates a new BidRecommendationItemDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBidRecommendationItemDTOWithDefaults

`func NewBidRecommendationItemDTOWithDefaults() *BidRecommendationItemDTO`

NewBidRecommendationItemDTOWithDefaults instantiates a new BidRecommendationItemDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBid

`func (o *BidRecommendationItemDTO) GetBid() int32`

GetBid returns the Bid field if non-nil, zero value otherwise.

### GetBidOk

`func (o *BidRecommendationItemDTO) GetBidOk() (*int32, bool)`

GetBidOk returns a tuple with the Bid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBid

`func (o *BidRecommendationItemDTO) SetBid(v int32)`

SetBid sets Bid field to given value.


### GetShowPercent

`func (o *BidRecommendationItemDTO) GetShowPercent() int64`

GetShowPercent returns the ShowPercent field if non-nil, zero value otherwise.

### GetShowPercentOk

`func (o *BidRecommendationItemDTO) GetShowPercentOk() (*int64, bool)`

GetShowPercentOk returns a tuple with the ShowPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPercent

`func (o *BidRecommendationItemDTO) SetShowPercent(v int64)`

SetShowPercent sets ShowPercent field to given value.


### GetBenefits

`func (o *BidRecommendationItemDTO) GetBenefits() []BenefitType`

GetBenefits returns the Benefits field if non-nil, zero value otherwise.

### GetBenefitsOk

`func (o *BidRecommendationItemDTO) GetBenefitsOk() (*[]BenefitType, bool)`

GetBenefitsOk returns a tuple with the Benefits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefits

`func (o *BidRecommendationItemDTO) SetBenefits(v []BenefitType)`

SetBenefits sets Benefits field to given value.

### HasBenefits

`func (o *BidRecommendationItemDTO) HasBenefits() bool`

HasBenefits returns a boolean if a field has been set.

### SetBenefitsNil

`func (o *BidRecommendationItemDTO) SetBenefitsNil(b bool)`

 SetBenefitsNil sets the value for Benefits to be an explicit nil

### UnsetBenefits
`func (o *BidRecommendationItemDTO) UnsetBenefits()`

UnsetBenefits ensures that no value is present for Benefits, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


