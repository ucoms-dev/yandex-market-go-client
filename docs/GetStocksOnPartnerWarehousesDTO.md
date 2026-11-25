# GetStocksOnPartnerWarehousesDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerWarehouseId** | **int64** | Идентификатор склада.  | 
**Offers** | [**[]PartnerWarehouseOfferDTO**](PartnerWarehouseOfferDTO.md) | Страница списка товаров.  | 
**Paging** | Pointer to [**PackagingForwardScrollingPagerDTO**](PackagingForwardScrollingPagerDTO.md) |  | [optional] 

## Methods

### NewGetStocksOnPartnerWarehousesDTO

`func NewGetStocksOnPartnerWarehousesDTO(partnerWarehouseId int64, offers []PartnerWarehouseOfferDTO, ) *GetStocksOnPartnerWarehousesDTO`

NewGetStocksOnPartnerWarehousesDTO instantiates a new GetStocksOnPartnerWarehousesDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStocksOnPartnerWarehousesDTOWithDefaults

`func NewGetStocksOnPartnerWarehousesDTOWithDefaults() *GetStocksOnPartnerWarehousesDTO`

NewGetStocksOnPartnerWarehousesDTOWithDefaults instantiates a new GetStocksOnPartnerWarehousesDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerWarehouseId

`func (o *GetStocksOnPartnerWarehousesDTO) GetPartnerWarehouseId() int64`

GetPartnerWarehouseId returns the PartnerWarehouseId field if non-nil, zero value otherwise.

### GetPartnerWarehouseIdOk

`func (o *GetStocksOnPartnerWarehousesDTO) GetPartnerWarehouseIdOk() (*int64, bool)`

GetPartnerWarehouseIdOk returns a tuple with the PartnerWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerWarehouseId

`func (o *GetStocksOnPartnerWarehousesDTO) SetPartnerWarehouseId(v int64)`

SetPartnerWarehouseId sets PartnerWarehouseId field to given value.


### GetOffers

`func (o *GetStocksOnPartnerWarehousesDTO) GetOffers() []PartnerWarehouseOfferDTO`

GetOffers returns the Offers field if non-nil, zero value otherwise.

### GetOffersOk

`func (o *GetStocksOnPartnerWarehousesDTO) GetOffersOk() (*[]PartnerWarehouseOfferDTO, bool)`

GetOffersOk returns a tuple with the Offers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffers

`func (o *GetStocksOnPartnerWarehousesDTO) SetOffers(v []PartnerWarehouseOfferDTO)`

SetOffers sets Offers field to given value.


### GetPaging

`func (o *GetStocksOnPartnerWarehousesDTO) GetPaging() PackagingForwardScrollingPagerDTO`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *GetStocksOnPartnerWarehousesDTO) GetPagingOk() (*PackagingForwardScrollingPagerDTO, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *GetStocksOnPartnerWarehousesDTO) SetPaging(v PackagingForwardScrollingPagerDTO)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *GetStocksOnPartnerWarehousesDTO) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


