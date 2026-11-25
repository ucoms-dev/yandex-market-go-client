# PartnerWarehouseModelDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlacementType** | [**WarehouseProgramType**](WarehouseProgramType.md) |  | 
**CampaignId** | Pointer to **int64** | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями.  | [optional] 
**ApiAvailability** | [**ApiAvailabilityStatusType**](ApiAvailabilityStatusType.md) |  | 

## Methods

### NewPartnerWarehouseModelDTO

`func NewPartnerWarehouseModelDTO(placementType WarehouseProgramType, apiAvailability ApiAvailabilityStatusType, ) *PartnerWarehouseModelDTO`

NewPartnerWarehouseModelDTO instantiates a new PartnerWarehouseModelDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerWarehouseModelDTOWithDefaults

`func NewPartnerWarehouseModelDTOWithDefaults() *PartnerWarehouseModelDTO`

NewPartnerWarehouseModelDTOWithDefaults instantiates a new PartnerWarehouseModelDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlacementType

`func (o *PartnerWarehouseModelDTO) GetPlacementType() WarehouseProgramType`

GetPlacementType returns the PlacementType field if non-nil, zero value otherwise.

### GetPlacementTypeOk

`func (o *PartnerWarehouseModelDTO) GetPlacementTypeOk() (*WarehouseProgramType, bool)`

GetPlacementTypeOk returns a tuple with the PlacementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementType

`func (o *PartnerWarehouseModelDTO) SetPlacementType(v WarehouseProgramType)`

SetPlacementType sets PlacementType field to given value.


### GetCampaignId

`func (o *PartnerWarehouseModelDTO) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *PartnerWarehouseModelDTO) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *PartnerWarehouseModelDTO) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *PartnerWarehouseModelDTO) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetApiAvailability

`func (o *PartnerWarehouseModelDTO) GetApiAvailability() ApiAvailabilityStatusType`

GetApiAvailability returns the ApiAvailability field if non-nil, zero value otherwise.

### GetApiAvailabilityOk

`func (o *PartnerWarehouseModelDTO) GetApiAvailabilityOk() (*ApiAvailabilityStatusType, bool)`

GetApiAvailabilityOk returns a tuple with the ApiAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAvailability

`func (o *PartnerWarehouseModelDTO) SetApiAvailability(v ApiAvailabilityStatusType)`

SetApiAvailability sets ApiAvailability field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


