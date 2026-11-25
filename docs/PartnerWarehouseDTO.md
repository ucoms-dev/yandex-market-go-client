# PartnerWarehouseDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Идентификатор склада. | 
**Name** | **string** | Название склада. | 
**Models** | [**[]PartnerWarehouseModelDTO**](PartnerWarehouseModelDTO.md) | Модели работы, доступные для склада. | 
**Address** | Pointer to [**WarehouseAddressDTO**](WarehouseAddressDTO.md) |  | [optional] 

## Methods

### NewPartnerWarehouseDTO

`func NewPartnerWarehouseDTO(id int64, name string, models []PartnerWarehouseModelDTO, ) *PartnerWarehouseDTO`

NewPartnerWarehouseDTO instantiates a new PartnerWarehouseDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerWarehouseDTOWithDefaults

`func NewPartnerWarehouseDTOWithDefaults() *PartnerWarehouseDTO`

NewPartnerWarehouseDTOWithDefaults instantiates a new PartnerWarehouseDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PartnerWarehouseDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartnerWarehouseDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartnerWarehouseDTO) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *PartnerWarehouseDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnerWarehouseDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnerWarehouseDTO) SetName(v string)`

SetName sets Name field to given value.


### GetModels

`func (o *PartnerWarehouseDTO) GetModels() []PartnerWarehouseModelDTO`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *PartnerWarehouseDTO) GetModelsOk() (*[]PartnerWarehouseModelDTO, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *PartnerWarehouseDTO) SetModels(v []PartnerWarehouseModelDTO)`

SetModels sets Models field to given value.


### GetAddress

`func (o *PartnerWarehouseDTO) GetAddress() WarehouseAddressDTO`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PartnerWarehouseDTO) GetAddressOk() (*WarehouseAddressDTO, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PartnerWarehouseDTO) SetAddress(v WarehouseAddressDTO)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PartnerWarehouseDTO) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


