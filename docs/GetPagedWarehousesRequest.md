# GetPagedWarehousesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]WarehouseComponentType**](WarehouseComponentType.md) | Свойства складов, которые необходимо вернуть. Если какое-то значение параметра не задано, этой информации в ответе не будет.  Передавайте параметр, только если нужна информация, которую он возвращает.  Можно передать сразу несколько значений.  | [optional] 
**CampaignIds** | Pointer to **[]int64** | Список идентификаторов кампании тех магазинов, склады которых необходимо вернуть.  Их можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не используйте вместо них идентификаторы магазинов, которые указаны в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | [optional] 

## Methods

### NewGetPagedWarehousesRequest

`func NewGetPagedWarehousesRequest() *GetPagedWarehousesRequest`

NewGetPagedWarehousesRequest instantiates a new GetPagedWarehousesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPagedWarehousesRequestWithDefaults

`func NewGetPagedWarehousesRequestWithDefaults() *GetPagedWarehousesRequest`

NewGetPagedWarehousesRequestWithDefaults instantiates a new GetPagedWarehousesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *GetPagedWarehousesRequest) GetComponents() []WarehouseComponentType`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *GetPagedWarehousesRequest) GetComponentsOk() (*[]WarehouseComponentType, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *GetPagedWarehousesRequest) SetComponents(v []WarehouseComponentType)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *GetPagedWarehousesRequest) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### SetComponentsNil

`func (o *GetPagedWarehousesRequest) SetComponentsNil(b bool)`

 SetComponentsNil sets the value for Components to be an explicit nil

### UnsetComponents
`func (o *GetPagedWarehousesRequest) UnsetComponents()`

UnsetComponents ensures that no value is present for Components, not even an explicit nil
### GetCampaignIds

`func (o *GetPagedWarehousesRequest) GetCampaignIds() []int64`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *GetPagedWarehousesRequest) GetCampaignIdsOk() (*[]int64, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignIds

`func (o *GetPagedWarehousesRequest) SetCampaignIds(v []int64)`

SetCampaignIds sets CampaignIds field to given value.

### HasCampaignIds

`func (o *GetPagedWarehousesRequest) HasCampaignIds() bool`

HasCampaignIds returns a boolean if a field has been set.

### SetCampaignIdsNil

`func (o *GetPagedWarehousesRequest) SetCampaignIdsNil(b bool)`

 SetCampaignIdsNil sets the value for CampaignIds to be an explicit nil

### UnsetCampaignIds
`func (o *GetPagedWarehousesRequest) UnsetCampaignIds()`

UnsetCampaignIds ensures that no value is present for CampaignIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


