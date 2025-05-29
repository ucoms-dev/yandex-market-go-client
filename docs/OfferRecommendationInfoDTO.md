# OfferRecommendationInfoDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferId** | Pointer to **string** | Ваш SKU — идентификатор товара в вашей системе.  Правила использования SKU:  * У каждого товара SKU должен быть свой.  * Уже заданный SKU нельзя освободить и использовать заново для другого товара. Каждый товар должен получать новый идентификатор, до того никогда не использовавшийся в вашем каталоге.  SKU товара можно изменить в кабинете продавца на Маркете. О том, как это сделать, читайте [в Справке Маркета для продавцов](https://yandex.ru/support2/marketplace/ru/assortment/operations/edit-sku).  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields)  | [optional] 
**CompetitivenessThresholds** | Pointer to [**PriceCompetitivenessThresholdsDTO**](PriceCompetitivenessThresholdsDTO.md) |  | [optional] 

## Methods

### NewOfferRecommendationInfoDTO

`func NewOfferRecommendationInfoDTO() *OfferRecommendationInfoDTO`

NewOfferRecommendationInfoDTO instantiates a new OfferRecommendationInfoDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferRecommendationInfoDTOWithDefaults

`func NewOfferRecommendationInfoDTOWithDefaults() *OfferRecommendationInfoDTO`

NewOfferRecommendationInfoDTOWithDefaults instantiates a new OfferRecommendationInfoDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferId

`func (o *OfferRecommendationInfoDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *OfferRecommendationInfoDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *OfferRecommendationInfoDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *OfferRecommendationInfoDTO) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetCompetitivenessThresholds

`func (o *OfferRecommendationInfoDTO) GetCompetitivenessThresholds() PriceCompetitivenessThresholdsDTO`

GetCompetitivenessThresholds returns the CompetitivenessThresholds field if non-nil, zero value otherwise.

### GetCompetitivenessThresholdsOk

`func (o *OfferRecommendationInfoDTO) GetCompetitivenessThresholdsOk() (*PriceCompetitivenessThresholdsDTO, bool)`

GetCompetitivenessThresholdsOk returns a tuple with the CompetitivenessThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitivenessThresholds

`func (o *OfferRecommendationInfoDTO) SetCompetitivenessThresholds(v PriceCompetitivenessThresholdsDTO)`

SetCompetitivenessThresholds sets CompetitivenessThresholds field to given value.

### HasCompetitivenessThresholds

`func (o *OfferRecommendationInfoDTO) HasCompetitivenessThresholds() bool`

HasCompetitivenessThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


