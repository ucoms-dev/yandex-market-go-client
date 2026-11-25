# SuggestSizeTablesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **int32** | Идентификатор категории на Маркете.  При изменении категории убедитесь, что характеристики товара и их значения в параметре &#x60;parameterValues&#x60; вы передаете для новой категории.  Список категорий Маркета можно получить с помощью запроса  [POST v2/categories/tree](../../reference/categories/getCategoriesTree.md).  | 
**VendorId** | Pointer to **int64** | Идентификатор бренда: значение характеристики **Бренд** (&#x60;id: 7893318&#x60;).  Чтобы его узнать, воспользуйтесь запросом [POST v2/category/{categoryId}/parameters](../../reference/content/getCategoryContentParameters.md) — идентификатор придет в &#x60;parameters[].values[].id&#x60;.  | [optional] 
**GenderId** | Pointer to **int64** | Идентификатор пола: значение характеристики **Пол** (&#x60;id: 14805991&#x60;).  Чтобы его узнать, воспользуйтесь запросом [POST v2/category/{categoryId}/parameters](../../reference/content/getCategoryContentParameters.md) — идентификатор придет в &#x60;parameters[].values[].id&#x60;.  | [optional] 

## Methods

### NewSuggestSizeTablesRequest

`func NewSuggestSizeTablesRequest(categoryId int32, ) *SuggestSizeTablesRequest`

NewSuggestSizeTablesRequest instantiates a new SuggestSizeTablesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestSizeTablesRequestWithDefaults

`func NewSuggestSizeTablesRequestWithDefaults() *SuggestSizeTablesRequest`

NewSuggestSizeTablesRequestWithDefaults instantiates a new SuggestSizeTablesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *SuggestSizeTablesRequest) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *SuggestSizeTablesRequest) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *SuggestSizeTablesRequest) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.


### GetVendorId

`func (o *SuggestSizeTablesRequest) GetVendorId() int64`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *SuggestSizeTablesRequest) GetVendorIdOk() (*int64, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *SuggestSizeTablesRequest) SetVendorId(v int64)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *SuggestSizeTablesRequest) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetGenderId

`func (o *SuggestSizeTablesRequest) GetGenderId() int64`

GetGenderId returns the GenderId field if non-nil, zero value otherwise.

### GetGenderIdOk

`func (o *SuggestSizeTablesRequest) GetGenderIdOk() (*int64, bool)`

GetGenderIdOk returns a tuple with the GenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenderId

`func (o *SuggestSizeTablesRequest) SetGenderId(v int64)`

SetGenderId sets GenderId field to given value.

### HasGenderId

`func (o *SuggestSizeTablesRequest) HasGenderId() bool`

HasGenderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


