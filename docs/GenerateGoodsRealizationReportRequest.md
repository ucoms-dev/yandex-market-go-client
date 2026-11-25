# GenerateGoodsRealizationReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **int64** | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями.  | 
**Year** | **int32** | Год. | 
**Month** | **int32** | Номер месяца. | 
**Inns** | Pointer to **[]string** | Список ИНН (доступно продавцам-резидентам РФ). | [optional] 
**PlacementPrograms** | Pointer to [**[]PlacementType**](PlacementType.md) | Список моделей размещения. | [optional] 
**PlacementContracts** | Pointer to **[]string** | Список номеров договоров на размещение. | [optional] 

## Methods

### NewGenerateGoodsRealizationReportRequest

`func NewGenerateGoodsRealizationReportRequest(campaignId int64, year int32, month int32, ) *GenerateGoodsRealizationReportRequest`

NewGenerateGoodsRealizationReportRequest instantiates a new GenerateGoodsRealizationReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateGoodsRealizationReportRequestWithDefaults

`func NewGenerateGoodsRealizationReportRequestWithDefaults() *GenerateGoodsRealizationReportRequest`

NewGenerateGoodsRealizationReportRequestWithDefaults instantiates a new GenerateGoodsRealizationReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *GenerateGoodsRealizationReportRequest) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *GenerateGoodsRealizationReportRequest) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *GenerateGoodsRealizationReportRequest) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetYear

`func (o *GenerateGoodsRealizationReportRequest) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *GenerateGoodsRealizationReportRequest) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *GenerateGoodsRealizationReportRequest) SetYear(v int32)`

SetYear sets Year field to given value.


### GetMonth

`func (o *GenerateGoodsRealizationReportRequest) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *GenerateGoodsRealizationReportRequest) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *GenerateGoodsRealizationReportRequest) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetInns

`func (o *GenerateGoodsRealizationReportRequest) GetInns() []string`

GetInns returns the Inns field if non-nil, zero value otherwise.

### GetInnsOk

`func (o *GenerateGoodsRealizationReportRequest) GetInnsOk() (*[]string, bool)`

GetInnsOk returns a tuple with the Inns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInns

`func (o *GenerateGoodsRealizationReportRequest) SetInns(v []string)`

SetInns sets Inns field to given value.

### HasInns

`func (o *GenerateGoodsRealizationReportRequest) HasInns() bool`

HasInns returns a boolean if a field has been set.

### SetInnsNil

`func (o *GenerateGoodsRealizationReportRequest) SetInnsNil(b bool)`

 SetInnsNil sets the value for Inns to be an explicit nil

### UnsetInns
`func (o *GenerateGoodsRealizationReportRequest) UnsetInns()`

UnsetInns ensures that no value is present for Inns, not even an explicit nil
### GetPlacementPrograms

`func (o *GenerateGoodsRealizationReportRequest) GetPlacementPrograms() []PlacementType`

GetPlacementPrograms returns the PlacementPrograms field if non-nil, zero value otherwise.

### GetPlacementProgramsOk

`func (o *GenerateGoodsRealizationReportRequest) GetPlacementProgramsOk() (*[]PlacementType, bool)`

GetPlacementProgramsOk returns a tuple with the PlacementPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementPrograms

`func (o *GenerateGoodsRealizationReportRequest) SetPlacementPrograms(v []PlacementType)`

SetPlacementPrograms sets PlacementPrograms field to given value.

### HasPlacementPrograms

`func (o *GenerateGoodsRealizationReportRequest) HasPlacementPrograms() bool`

HasPlacementPrograms returns a boolean if a field has been set.

### SetPlacementProgramsNil

`func (o *GenerateGoodsRealizationReportRequest) SetPlacementProgramsNil(b bool)`

 SetPlacementProgramsNil sets the value for PlacementPrograms to be an explicit nil

### UnsetPlacementPrograms
`func (o *GenerateGoodsRealizationReportRequest) UnsetPlacementPrograms()`

UnsetPlacementPrograms ensures that no value is present for PlacementPrograms, not even an explicit nil
### GetPlacementContracts

`func (o *GenerateGoodsRealizationReportRequest) GetPlacementContracts() []string`

GetPlacementContracts returns the PlacementContracts field if non-nil, zero value otherwise.

### GetPlacementContractsOk

`func (o *GenerateGoodsRealizationReportRequest) GetPlacementContractsOk() (*[]string, bool)`

GetPlacementContractsOk returns a tuple with the PlacementContracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementContracts

`func (o *GenerateGoodsRealizationReportRequest) SetPlacementContracts(v []string)`

SetPlacementContracts sets PlacementContracts field to given value.

### HasPlacementContracts

`func (o *GenerateGoodsRealizationReportRequest) HasPlacementContracts() bool`

HasPlacementContracts returns a boolean if a field has been set.

### SetPlacementContractsNil

`func (o *GenerateGoodsRealizationReportRequest) SetPlacementContractsNil(b bool)`

 SetPlacementContractsNil sets the value for PlacementContracts to be an explicit nil

### UnsetPlacementContracts
`func (o *GenerateGoodsRealizationReportRequest) UnsetPlacementContracts()`

UnsetPlacementContracts ensures that no value is present for PlacementContracts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


