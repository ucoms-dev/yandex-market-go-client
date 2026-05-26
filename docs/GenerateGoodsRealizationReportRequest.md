# GenerateGoodsRealizationReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **int64** | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями.  | 
**Year** | **int32** | Год. | 
**Month** | **int32** | Номер месяца. | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


