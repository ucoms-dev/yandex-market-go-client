# OfferDefaultPriceListResponseDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offers** | [**[]OfferDefaultPriceResponseDTO**](OfferDefaultPriceResponseDTO.md) | Список товаров. | 
**Paging** | Pointer to [**PackagingForwardScrollingPagerDTO**](PackagingForwardScrollingPagerDTO.md) |  | [optional] 

## Methods

### NewOfferDefaultPriceListResponseDTO

`func NewOfferDefaultPriceListResponseDTO(offers []OfferDefaultPriceResponseDTO, ) *OfferDefaultPriceListResponseDTO`

NewOfferDefaultPriceListResponseDTO instantiates a new OfferDefaultPriceListResponseDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferDefaultPriceListResponseDTOWithDefaults

`func NewOfferDefaultPriceListResponseDTOWithDefaults() *OfferDefaultPriceListResponseDTO`

NewOfferDefaultPriceListResponseDTOWithDefaults instantiates a new OfferDefaultPriceListResponseDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffers

`func (o *OfferDefaultPriceListResponseDTO) GetOffers() []OfferDefaultPriceResponseDTO`

GetOffers returns the Offers field if non-nil, zero value otherwise.

### GetOffersOk

`func (o *OfferDefaultPriceListResponseDTO) GetOffersOk() (*[]OfferDefaultPriceResponseDTO, bool)`

GetOffersOk returns a tuple with the Offers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffers

`func (o *OfferDefaultPriceListResponseDTO) SetOffers(v []OfferDefaultPriceResponseDTO)`

SetOffers sets Offers field to given value.


### GetPaging

`func (o *OfferDefaultPriceListResponseDTO) GetPaging() PackagingForwardScrollingPagerDTO`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *OfferDefaultPriceListResponseDTO) GetPagingOk() (*PackagingForwardScrollingPagerDTO, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *OfferDefaultPriceListResponseDTO) SetPaging(v PackagingForwardScrollingPagerDTO)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *OfferDefaultPriceListResponseDTO) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


