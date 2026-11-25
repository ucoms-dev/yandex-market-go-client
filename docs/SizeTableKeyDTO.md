# SizeTableKeyDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **int64** | Идентификатор категории из запроса. | 
**VendorId** | Pointer to **int64** | Идентификатор бренда на Маркете. Возвращается только для таблиц бренда. | [optional] 
**GenderId** | Pointer to **int64** | Идентификатор пола на Маркете. | [optional] 

## Methods

### NewSizeTableKeyDTO

`func NewSizeTableKeyDTO(categoryId int64, ) *SizeTableKeyDTO`

NewSizeTableKeyDTO instantiates a new SizeTableKeyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSizeTableKeyDTOWithDefaults

`func NewSizeTableKeyDTOWithDefaults() *SizeTableKeyDTO`

NewSizeTableKeyDTOWithDefaults instantiates a new SizeTableKeyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *SizeTableKeyDTO) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *SizeTableKeyDTO) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *SizeTableKeyDTO) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.


### GetVendorId

`func (o *SizeTableKeyDTO) GetVendorId() int64`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *SizeTableKeyDTO) GetVendorIdOk() (*int64, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *SizeTableKeyDTO) SetVendorId(v int64)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *SizeTableKeyDTO) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetGenderId

`func (o *SizeTableKeyDTO) GetGenderId() int64`

GetGenderId returns the GenderId field if non-nil, zero value otherwise.

### GetGenderIdOk

`func (o *SizeTableKeyDTO) GetGenderIdOk() (*int64, bool)`

GetGenderIdOk returns a tuple with the GenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenderId

`func (o *SizeTableKeyDTO) SetGenderId(v int64)`

SetGenderId sets GenderId field to given value.

### HasGenderId

`func (o *SizeTableKeyDTO) HasGenderId() bool`

HasGenderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


