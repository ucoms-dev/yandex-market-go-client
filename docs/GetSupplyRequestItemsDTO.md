# GetSupplyRequestItemsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]SupplyRequestItemDTO**](SupplyRequestItemDTO.md) | Список товаров. | 
**Paging** | Pointer to [**ForwardScrollingPagerDTO**](ForwardScrollingPagerDTO.md) |  | [optional] 

## Methods

### NewGetSupplyRequestItemsDTO

`func NewGetSupplyRequestItemsDTO(items []SupplyRequestItemDTO, ) *GetSupplyRequestItemsDTO`

NewGetSupplyRequestItemsDTO instantiates a new GetSupplyRequestItemsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSupplyRequestItemsDTOWithDefaults

`func NewGetSupplyRequestItemsDTOWithDefaults() *GetSupplyRequestItemsDTO`

NewGetSupplyRequestItemsDTOWithDefaults instantiates a new GetSupplyRequestItemsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GetSupplyRequestItemsDTO) GetItems() []SupplyRequestItemDTO`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetSupplyRequestItemsDTO) GetItemsOk() (*[]SupplyRequestItemDTO, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetSupplyRequestItemsDTO) SetItems(v []SupplyRequestItemDTO)`

SetItems sets Items field to given value.


### GetPaging

`func (o *GetSupplyRequestItemsDTO) GetPaging() ForwardScrollingPagerDTO`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *GetSupplyRequestItemsDTO) GetPagingOk() (*ForwardScrollingPagerDTO, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *GetSupplyRequestItemsDTO) SetPaging(v ForwardScrollingPagerDTO)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *GetSupplyRequestItemsDTO) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


