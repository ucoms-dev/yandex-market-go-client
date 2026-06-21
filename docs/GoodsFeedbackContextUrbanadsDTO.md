# GoodsFeedbackContextUrbanadsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferName** | Pointer to **string** | Название товара, под которым оставлен отзыв. | [optional] 
**PictureUrl** | Pointer to **string** |  | [optional] 
**BusinessId** | Pointer to **int64** | Идентификатор кабинета. {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) {% endif %}  | [optional] 
**BusinessName** | Pointer to **string** | Название бизнеса, под товаром которого оставлен отзыв. | [optional] 
**BrandId** | Pointer to **string** | Идентификатор бренда товара. | [optional] 
**BrandName** | Pointer to **string** | Название бренда товара. | [optional] 

## Methods

### NewGoodsFeedbackContextUrbanadsDTO

`func NewGoodsFeedbackContextUrbanadsDTO() *GoodsFeedbackContextUrbanadsDTO`

NewGoodsFeedbackContextUrbanadsDTO instantiates a new GoodsFeedbackContextUrbanadsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoodsFeedbackContextUrbanadsDTOWithDefaults

`func NewGoodsFeedbackContextUrbanadsDTOWithDefaults() *GoodsFeedbackContextUrbanadsDTO`

NewGoodsFeedbackContextUrbanadsDTOWithDefaults instantiates a new GoodsFeedbackContextUrbanadsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferName

`func (o *GoodsFeedbackContextUrbanadsDTO) GetOfferName() string`

GetOfferName returns the OfferName field if non-nil, zero value otherwise.

### GetOfferNameOk

`func (o *GoodsFeedbackContextUrbanadsDTO) GetOfferNameOk() (*string, bool)`

GetOfferNameOk returns a tuple with the OfferName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferName

`func (o *GoodsFeedbackContextUrbanadsDTO) SetOfferName(v string)`

SetOfferName sets OfferName field to given value.

### HasOfferName

`func (o *GoodsFeedbackContextUrbanadsDTO) HasOfferName() bool`

HasOfferName returns a boolean if a field has been set.

### GetPictureUrl

`func (o *GoodsFeedbackContextUrbanadsDTO) GetPictureUrl() string`

GetPictureUrl returns the PictureUrl field if non-nil, zero value otherwise.

### GetPictureUrlOk

`func (o *GoodsFeedbackContextUrbanadsDTO) GetPictureUrlOk() (*string, bool)`

GetPictureUrlOk returns a tuple with the PictureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictureUrl

`func (o *GoodsFeedbackContextUrbanadsDTO) SetPictureUrl(v string)`

SetPictureUrl sets PictureUrl field to given value.

### HasPictureUrl

`func (o *GoodsFeedbackContextUrbanadsDTO) HasPictureUrl() bool`

HasPictureUrl returns a boolean if a field has been set.

### GetBusinessId

`func (o *GoodsFeedbackContextUrbanadsDTO) GetBusinessId() int64`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *GoodsFeedbackContextUrbanadsDTO) GetBusinessIdOk() (*int64, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *GoodsFeedbackContextUrbanadsDTO) SetBusinessId(v int64)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *GoodsFeedbackContextUrbanadsDTO) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetBusinessName

`func (o *GoodsFeedbackContextUrbanadsDTO) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *GoodsFeedbackContextUrbanadsDTO) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *GoodsFeedbackContextUrbanadsDTO) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *GoodsFeedbackContextUrbanadsDTO) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### GetBrandId

`func (o *GoodsFeedbackContextUrbanadsDTO) GetBrandId() string`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *GoodsFeedbackContextUrbanadsDTO) GetBrandIdOk() (*string, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *GoodsFeedbackContextUrbanadsDTO) SetBrandId(v string)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *GoodsFeedbackContextUrbanadsDTO) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### GetBrandName

`func (o *GoodsFeedbackContextUrbanadsDTO) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *GoodsFeedbackContextUrbanadsDTO) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *GoodsFeedbackContextUrbanadsDTO) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.

### HasBrandName

`func (o *GoodsFeedbackContextUrbanadsDTO) HasBrandName() bool`

HasBrandName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


