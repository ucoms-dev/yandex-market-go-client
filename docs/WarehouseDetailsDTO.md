# WarehouseDetailsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Идентификатор склада. | 
**Name** | **string** | Название склада. | 
**CampaignId** | **int64** | Идентификатор кампании того магазина, который связан со складом.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**Express** | **bool** | Возможна ли доставка для модели Экспресс. | 
**GroupInfo** | Pointer to [**WarehouseGroupInfoDTO**](WarehouseGroupInfoDTO.md) |  | [optional] 
**Address** | Pointer to [**WarehouseAddressDTO**](WarehouseAddressDTO.md) |  | [optional] 
**Status** | Pointer to [**WarehouseStatusDTO**](WarehouseStatusDTO.md) |  | [optional] 

## Methods

### NewWarehouseDetailsDTO

`func NewWarehouseDetailsDTO(id int64, name string, campaignId int64, express bool, ) *WarehouseDetailsDTO`

NewWarehouseDetailsDTO instantiates a new WarehouseDetailsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseDetailsDTOWithDefaults

`func NewWarehouseDetailsDTOWithDefaults() *WarehouseDetailsDTO`

NewWarehouseDetailsDTOWithDefaults instantiates a new WarehouseDetailsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WarehouseDetailsDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WarehouseDetailsDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WarehouseDetailsDTO) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *WarehouseDetailsDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WarehouseDetailsDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WarehouseDetailsDTO) SetName(v string)`

SetName sets Name field to given value.


### GetCampaignId

`func (o *WarehouseDetailsDTO) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *WarehouseDetailsDTO) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *WarehouseDetailsDTO) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetExpress

`func (o *WarehouseDetailsDTO) GetExpress() bool`

GetExpress returns the Express field if non-nil, zero value otherwise.

### GetExpressOk

`func (o *WarehouseDetailsDTO) GetExpressOk() (*bool, bool)`

GetExpressOk returns a tuple with the Express field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpress

`func (o *WarehouseDetailsDTO) SetExpress(v bool)`

SetExpress sets Express field to given value.


### GetGroupInfo

`func (o *WarehouseDetailsDTO) GetGroupInfo() WarehouseGroupInfoDTO`

GetGroupInfo returns the GroupInfo field if non-nil, zero value otherwise.

### GetGroupInfoOk

`func (o *WarehouseDetailsDTO) GetGroupInfoOk() (*WarehouseGroupInfoDTO, bool)`

GetGroupInfoOk returns a tuple with the GroupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupInfo

`func (o *WarehouseDetailsDTO) SetGroupInfo(v WarehouseGroupInfoDTO)`

SetGroupInfo sets GroupInfo field to given value.

### HasGroupInfo

`func (o *WarehouseDetailsDTO) HasGroupInfo() bool`

HasGroupInfo returns a boolean if a field has been set.

### GetAddress

`func (o *WarehouseDetailsDTO) GetAddress() WarehouseAddressDTO`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WarehouseDetailsDTO) GetAddressOk() (*WarehouseAddressDTO, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WarehouseDetailsDTO) SetAddress(v WarehouseAddressDTO)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *WarehouseDetailsDTO) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetStatus

`func (o *WarehouseDetailsDTO) GetStatus() WarehouseStatusDTO`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WarehouseDetailsDTO) GetStatusOk() (*WarehouseStatusDTO, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WarehouseDetailsDTO) SetStatus(v WarehouseStatusDTO)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WarehouseDetailsDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


