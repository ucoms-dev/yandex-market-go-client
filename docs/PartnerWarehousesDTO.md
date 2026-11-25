# PartnerWarehousesDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warehouses** | [**[]PartnerWarehouseDTO**](PartnerWarehouseDTO.md) | Список складов. | 
**Paging** | Pointer to [**PackagingForwardScrollingPagerDTO**](PackagingForwardScrollingPagerDTO.md) |  | [optional] 

## Methods

### NewPartnerWarehousesDTO

`func NewPartnerWarehousesDTO(warehouses []PartnerWarehouseDTO, ) *PartnerWarehousesDTO`

NewPartnerWarehousesDTO instantiates a new PartnerWarehousesDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerWarehousesDTOWithDefaults

`func NewPartnerWarehousesDTOWithDefaults() *PartnerWarehousesDTO`

NewPartnerWarehousesDTOWithDefaults instantiates a new PartnerWarehousesDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarehouses

`func (o *PartnerWarehousesDTO) GetWarehouses() []PartnerWarehouseDTO`

GetWarehouses returns the Warehouses field if non-nil, zero value otherwise.

### GetWarehousesOk

`func (o *PartnerWarehousesDTO) GetWarehousesOk() (*[]PartnerWarehouseDTO, bool)`

GetWarehousesOk returns a tuple with the Warehouses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouses

`func (o *PartnerWarehousesDTO) SetWarehouses(v []PartnerWarehouseDTO)`

SetWarehouses sets Warehouses field to given value.


### GetPaging

`func (o *PartnerWarehousesDTO) GetPaging() PackagingForwardScrollingPagerDTO`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *PartnerWarehousesDTO) GetPagingOk() (*PackagingForwardScrollingPagerDTO, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *PartnerWarehousesDTO) SetPaging(v PackagingForwardScrollingPagerDTO)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *PartnerWarehousesDTO) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


