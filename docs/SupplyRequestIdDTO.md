# SupplyRequestIdDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Идентификатор заявки.  {% note warning \&quot;Используется только в API\&quot; %}  По нему не получится найти заявки в кабинете продавца на Маркете. Для этого используйте &#x60;marketplaceRequestId&#x60; или &#x60;warehouseRequestId&#x60;.  {% endnote %}  | 
**MarketplaceRequestId** | Pointer to **string** | Номер заявки на маркетплейсе.  Также указывается в кабинете продавца на Маркете.  | [optional] 
**WarehouseRequestId** | Pointer to **string** | Номер заявки на складе.  Также указывается в кабинете продавца на Маркете.  | [optional] 

## Methods

### NewSupplyRequestIdDTO

`func NewSupplyRequestIdDTO(id int64, ) *SupplyRequestIdDTO`

NewSupplyRequestIdDTO instantiates a new SupplyRequestIdDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplyRequestIdDTOWithDefaults

`func NewSupplyRequestIdDTOWithDefaults() *SupplyRequestIdDTO`

NewSupplyRequestIdDTOWithDefaults instantiates a new SupplyRequestIdDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupplyRequestIdDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupplyRequestIdDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupplyRequestIdDTO) SetId(v int64)`

SetId sets Id field to given value.


### GetMarketplaceRequestId

`func (o *SupplyRequestIdDTO) GetMarketplaceRequestId() string`

GetMarketplaceRequestId returns the MarketplaceRequestId field if non-nil, zero value otherwise.

### GetMarketplaceRequestIdOk

`func (o *SupplyRequestIdDTO) GetMarketplaceRequestIdOk() (*string, bool)`

GetMarketplaceRequestIdOk returns a tuple with the MarketplaceRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceRequestId

`func (o *SupplyRequestIdDTO) SetMarketplaceRequestId(v string)`

SetMarketplaceRequestId sets MarketplaceRequestId field to given value.

### HasMarketplaceRequestId

`func (o *SupplyRequestIdDTO) HasMarketplaceRequestId() bool`

HasMarketplaceRequestId returns a boolean if a field has been set.

### GetWarehouseRequestId

`func (o *SupplyRequestIdDTO) GetWarehouseRequestId() string`

GetWarehouseRequestId returns the WarehouseRequestId field if non-nil, zero value otherwise.

### GetWarehouseRequestIdOk

`func (o *SupplyRequestIdDTO) GetWarehouseRequestIdOk() (*string, bool)`

GetWarehouseRequestIdOk returns a tuple with the WarehouseRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseRequestId

`func (o *SupplyRequestIdDTO) SetWarehouseRequestId(v string)`

SetWarehouseRequestId sets WarehouseRequestId field to given value.

### HasWarehouseRequestId

`func (o *SupplyRequestIdDTO) HasWarehouseRequestId() bool`

HasWarehouseRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


