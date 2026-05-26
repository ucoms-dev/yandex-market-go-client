# GenerateStocksOnWarehousesReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | Pointer to **int64** | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями.  | [optional] 
**BusinessId** | Pointer to **int64** | Идентификатор кабинета. {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) {% endif %}  | [optional] 
**WarehouseIds** | Pointer to **[]int64** | Фильтр по идентификаторам складов (только модели FBY и LaaS). Чтобы узнать идентификатор, воспользуйтесь запросом [GET v2/warehouses](../../reference/warehouses/getFulfillmentWarehouses.md). | [optional] 
**ReportDate** | Pointer to **string** | Фильтр по дате (для моделей FBY и LaaS). В отчет попадут данные за **предшествующий** дате день.  Формат даты: &#x60;ГГГГ-ММ-ДД&#x60;.  | [optional] 
**CategoryIds** | Pointer to **[]int32** | Фильтр по категориям на Маркете (кроме моделей FBY и LaaS). | [optional] 
**HasStocks** | Pointer to **bool** | Фильтр по наличию остатков (кроме моделей FBY и LaaS). | [optional] 
**CampaignIds** | Pointer to **[]int64** | Фильтр по магазинам для отчета по кабинету (кроме моделей FBY и LaaS).  Передавайте вместе с &#x60;businessId&#x60;.  | [optional] 

## Methods

### NewGenerateStocksOnWarehousesReportRequest

`func NewGenerateStocksOnWarehousesReportRequest() *GenerateStocksOnWarehousesReportRequest`

NewGenerateStocksOnWarehousesReportRequest instantiates a new GenerateStocksOnWarehousesReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateStocksOnWarehousesReportRequestWithDefaults

`func NewGenerateStocksOnWarehousesReportRequestWithDefaults() *GenerateStocksOnWarehousesReportRequest`

NewGenerateStocksOnWarehousesReportRequestWithDefaults instantiates a new GenerateStocksOnWarehousesReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *GenerateStocksOnWarehousesReportRequest) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *GenerateStocksOnWarehousesReportRequest) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *GenerateStocksOnWarehousesReportRequest) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *GenerateStocksOnWarehousesReportRequest) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetBusinessId

`func (o *GenerateStocksOnWarehousesReportRequest) GetBusinessId() int64`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *GenerateStocksOnWarehousesReportRequest) GetBusinessIdOk() (*int64, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *GenerateStocksOnWarehousesReportRequest) SetBusinessId(v int64)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *GenerateStocksOnWarehousesReportRequest) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetWarehouseIds

`func (o *GenerateStocksOnWarehousesReportRequest) GetWarehouseIds() []int64`

GetWarehouseIds returns the WarehouseIds field if non-nil, zero value otherwise.

### GetWarehouseIdsOk

`func (o *GenerateStocksOnWarehousesReportRequest) GetWarehouseIdsOk() (*[]int64, bool)`

GetWarehouseIdsOk returns a tuple with the WarehouseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseIds

`func (o *GenerateStocksOnWarehousesReportRequest) SetWarehouseIds(v []int64)`

SetWarehouseIds sets WarehouseIds field to given value.

### HasWarehouseIds

`func (o *GenerateStocksOnWarehousesReportRequest) HasWarehouseIds() bool`

HasWarehouseIds returns a boolean if a field has been set.

### SetWarehouseIdsNil

`func (o *GenerateStocksOnWarehousesReportRequest) SetWarehouseIdsNil(b bool)`

 SetWarehouseIdsNil sets the value for WarehouseIds to be an explicit nil

### UnsetWarehouseIds
`func (o *GenerateStocksOnWarehousesReportRequest) UnsetWarehouseIds()`

UnsetWarehouseIds ensures that no value is present for WarehouseIds, not even an explicit nil
### GetReportDate

`func (o *GenerateStocksOnWarehousesReportRequest) GetReportDate() string`

GetReportDate returns the ReportDate field if non-nil, zero value otherwise.

### GetReportDateOk

`func (o *GenerateStocksOnWarehousesReportRequest) GetReportDateOk() (*string, bool)`

GetReportDateOk returns a tuple with the ReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDate

`func (o *GenerateStocksOnWarehousesReportRequest) SetReportDate(v string)`

SetReportDate sets ReportDate field to given value.

### HasReportDate

`func (o *GenerateStocksOnWarehousesReportRequest) HasReportDate() bool`

HasReportDate returns a boolean if a field has been set.

### GetCategoryIds

`func (o *GenerateStocksOnWarehousesReportRequest) GetCategoryIds() []int32`

GetCategoryIds returns the CategoryIds field if non-nil, zero value otherwise.

### GetCategoryIdsOk

`func (o *GenerateStocksOnWarehousesReportRequest) GetCategoryIdsOk() (*[]int32, bool)`

GetCategoryIdsOk returns a tuple with the CategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryIds

`func (o *GenerateStocksOnWarehousesReportRequest) SetCategoryIds(v []int32)`

SetCategoryIds sets CategoryIds field to given value.

### HasCategoryIds

`func (o *GenerateStocksOnWarehousesReportRequest) HasCategoryIds() bool`

HasCategoryIds returns a boolean if a field has been set.

### SetCategoryIdsNil

`func (o *GenerateStocksOnWarehousesReportRequest) SetCategoryIdsNil(b bool)`

 SetCategoryIdsNil sets the value for CategoryIds to be an explicit nil

### UnsetCategoryIds
`func (o *GenerateStocksOnWarehousesReportRequest) UnsetCategoryIds()`

UnsetCategoryIds ensures that no value is present for CategoryIds, not even an explicit nil
### GetHasStocks

`func (o *GenerateStocksOnWarehousesReportRequest) GetHasStocks() bool`

GetHasStocks returns the HasStocks field if non-nil, zero value otherwise.

### GetHasStocksOk

`func (o *GenerateStocksOnWarehousesReportRequest) GetHasStocksOk() (*bool, bool)`

GetHasStocksOk returns a tuple with the HasStocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasStocks

`func (o *GenerateStocksOnWarehousesReportRequest) SetHasStocks(v bool)`

SetHasStocks sets HasStocks field to given value.

### HasHasStocks

`func (o *GenerateStocksOnWarehousesReportRequest) HasHasStocks() bool`

HasHasStocks returns a boolean if a field has been set.

### GetCampaignIds

`func (o *GenerateStocksOnWarehousesReportRequest) GetCampaignIds() []int64`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *GenerateStocksOnWarehousesReportRequest) GetCampaignIdsOk() (*[]int64, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignIds

`func (o *GenerateStocksOnWarehousesReportRequest) SetCampaignIds(v []int64)`

SetCampaignIds sets CampaignIds field to given value.

### HasCampaignIds

`func (o *GenerateStocksOnWarehousesReportRequest) HasCampaignIds() bool`

HasCampaignIds returns a boolean if a field has been set.

### SetCampaignIdsNil

`func (o *GenerateStocksOnWarehousesReportRequest) SetCampaignIdsNil(b bool)`

 SetCampaignIdsNil sets the value for CampaignIds to be an explicit nil

### UnsetCampaignIds
`func (o *GenerateStocksOnWarehousesReportRequest) UnsetCampaignIds()`

UnsetCampaignIds ensures that no value is present for CampaignIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


