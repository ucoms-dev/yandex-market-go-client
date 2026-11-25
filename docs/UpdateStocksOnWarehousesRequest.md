# UpdateStocksOnWarehousesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkuItems** | [**[]UpdateStockOnWarehouseItemDTO**](UpdateStockOnWarehouseItemDTO.md) | Данные об остатках товаров.  В рамках одного запроса все пары &#x60;sku&#x60; и &#x60;partnerWarehouseId&#x60; должны быть уникальными. Не допускается передача двух объектов с одинаковыми &#x60;sku&#x60; и &#x60;partnerWarehouseId&#x60;.  | 

## Methods

### NewUpdateStocksOnWarehousesRequest

`func NewUpdateStocksOnWarehousesRequest(skuItems []UpdateStockOnWarehouseItemDTO, ) *UpdateStocksOnWarehousesRequest`

NewUpdateStocksOnWarehousesRequest instantiates a new UpdateStocksOnWarehousesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStocksOnWarehousesRequestWithDefaults

`func NewUpdateStocksOnWarehousesRequestWithDefaults() *UpdateStocksOnWarehousesRequest`

NewUpdateStocksOnWarehousesRequestWithDefaults instantiates a new UpdateStocksOnWarehousesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkuItems

`func (o *UpdateStocksOnWarehousesRequest) GetSkuItems() []UpdateStockOnWarehouseItemDTO`

GetSkuItems returns the SkuItems field if non-nil, zero value otherwise.

### GetSkuItemsOk

`func (o *UpdateStocksOnWarehousesRequest) GetSkuItemsOk() (*[]UpdateStockOnWarehouseItemDTO, bool)`

GetSkuItemsOk returns a tuple with the SkuItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuItems

`func (o *UpdateStocksOnWarehousesRequest) SetSkuItems(v []UpdateStockOnWarehouseItemDTO)`

SetSkuItems sets SkuItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


