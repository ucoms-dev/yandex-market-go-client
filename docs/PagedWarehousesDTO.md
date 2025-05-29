# PagedWarehousesDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warehouses** | [**[]WarehouseDetailsDTO**](WarehouseDetailsDTO.md) | Список складов. | 
**Paging** | Pointer to [**ForwardScrollingPagerDTO**](ForwardScrollingPagerDTO.md) |  | [optional] 

## Methods

### NewPagedWarehousesDTO

`func NewPagedWarehousesDTO(warehouses []WarehouseDetailsDTO, ) *PagedWarehousesDTO`

NewPagedWarehousesDTO instantiates a new PagedWarehousesDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedWarehousesDTOWithDefaults

`func NewPagedWarehousesDTOWithDefaults() *PagedWarehousesDTO`

NewPagedWarehousesDTOWithDefaults instantiates a new PagedWarehousesDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarehouses

`func (o *PagedWarehousesDTO) GetWarehouses() []WarehouseDetailsDTO`

GetWarehouses returns the Warehouses field if non-nil, zero value otherwise.

### GetWarehousesOk

`func (o *PagedWarehousesDTO) GetWarehousesOk() (*[]WarehouseDetailsDTO, bool)`

GetWarehousesOk returns a tuple with the Warehouses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouses

`func (o *PagedWarehousesDTO) SetWarehouses(v []WarehouseDetailsDTO)`

SetWarehouses sets Warehouses field to given value.


### GetPaging

`func (o *PagedWarehousesDTO) GetPaging() ForwardScrollingPagerDTO`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *PagedWarehousesDTO) GetPagingOk() (*ForwardScrollingPagerDTO, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *PagedWarehousesDTO) SetPaging(v ForwardScrollingPagerDTO)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *PagedWarehousesDTO) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


