# GenerateStocksReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryIds** | Pointer to **[]int64** | Фильтр по категориям на Маркете. | [optional] 
**HasStocks** | Pointer to **bool** | Фильтр по наличию остатков. | [optional] 
**PartnerWarehouseIds** | Pointer to **[]int64** | Фильтр по идентификаторам складов. Чтобы узнать идентификатор, воспользуйтесь запросом [POST v3/businesses/{businessId}/warehouses](../../reference/warehouses/getPartnerWarehouses.md).  | [optional] 

## Methods

### NewGenerateStocksReportRequest

`func NewGenerateStocksReportRequest() *GenerateStocksReportRequest`

NewGenerateStocksReportRequest instantiates a new GenerateStocksReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateStocksReportRequestWithDefaults

`func NewGenerateStocksReportRequestWithDefaults() *GenerateStocksReportRequest`

NewGenerateStocksReportRequestWithDefaults instantiates a new GenerateStocksReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryIds

`func (o *GenerateStocksReportRequest) GetCategoryIds() []int64`

GetCategoryIds returns the CategoryIds field if non-nil, zero value otherwise.

### GetCategoryIdsOk

`func (o *GenerateStocksReportRequest) GetCategoryIdsOk() (*[]int64, bool)`

GetCategoryIdsOk returns a tuple with the CategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryIds

`func (o *GenerateStocksReportRequest) SetCategoryIds(v []int64)`

SetCategoryIds sets CategoryIds field to given value.

### HasCategoryIds

`func (o *GenerateStocksReportRequest) HasCategoryIds() bool`

HasCategoryIds returns a boolean if a field has been set.

### SetCategoryIdsNil

`func (o *GenerateStocksReportRequest) SetCategoryIdsNil(b bool)`

 SetCategoryIdsNil sets the value for CategoryIds to be an explicit nil

### UnsetCategoryIds
`func (o *GenerateStocksReportRequest) UnsetCategoryIds()`

UnsetCategoryIds ensures that no value is present for CategoryIds, not even an explicit nil
### GetHasStocks

`func (o *GenerateStocksReportRequest) GetHasStocks() bool`

GetHasStocks returns the HasStocks field if non-nil, zero value otherwise.

### GetHasStocksOk

`func (o *GenerateStocksReportRequest) GetHasStocksOk() (*bool, bool)`

GetHasStocksOk returns a tuple with the HasStocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasStocks

`func (o *GenerateStocksReportRequest) SetHasStocks(v bool)`

SetHasStocks sets HasStocks field to given value.

### HasHasStocks

`func (o *GenerateStocksReportRequest) HasHasStocks() bool`

HasHasStocks returns a boolean if a field has been set.

### GetPartnerWarehouseIds

`func (o *GenerateStocksReportRequest) GetPartnerWarehouseIds() []int64`

GetPartnerWarehouseIds returns the PartnerWarehouseIds field if non-nil, zero value otherwise.

### GetPartnerWarehouseIdsOk

`func (o *GenerateStocksReportRequest) GetPartnerWarehouseIdsOk() (*[]int64, bool)`

GetPartnerWarehouseIdsOk returns a tuple with the PartnerWarehouseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerWarehouseIds

`func (o *GenerateStocksReportRequest) SetPartnerWarehouseIds(v []int64)`

SetPartnerWarehouseIds sets PartnerWarehouseIds field to given value.

### HasPartnerWarehouseIds

`func (o *GenerateStocksReportRequest) HasPartnerWarehouseIds() bool`

HasPartnerWarehouseIds returns a boolean if a field has been set.

### SetPartnerWarehouseIdsNil

`func (o *GenerateStocksReportRequest) SetPartnerWarehouseIdsNil(b bool)`

 SetPartnerWarehouseIdsNil sets the value for PartnerWarehouseIds to be an explicit nil

### UnsetPartnerWarehouseIds
`func (o *GenerateStocksReportRequest) UnsetPartnerWarehouseIds()`

UnsetPartnerWarehouseIds ensures that no value is present for PartnerWarehouseIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


