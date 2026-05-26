# GenerateGoodsTurnoverRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **int64** | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями.  | 
**Date** | Pointer to **string** | Дата, за которую нужно рассчитать оборачиваемость. Если параметр не указан, используется текущая дата.  Формат даты: &#x60;ГГГГ-ММ-ДД&#x60;.  | [optional] 

## Methods

### NewGenerateGoodsTurnoverRequest

`func NewGenerateGoodsTurnoverRequest(campaignId int64, ) *GenerateGoodsTurnoverRequest`

NewGenerateGoodsTurnoverRequest instantiates a new GenerateGoodsTurnoverRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateGoodsTurnoverRequestWithDefaults

`func NewGenerateGoodsTurnoverRequestWithDefaults() *GenerateGoodsTurnoverRequest`

NewGenerateGoodsTurnoverRequestWithDefaults instantiates a new GenerateGoodsTurnoverRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *GenerateGoodsTurnoverRequest) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *GenerateGoodsTurnoverRequest) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *GenerateGoodsTurnoverRequest) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetDate

`func (o *GenerateGoodsTurnoverRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GenerateGoodsTurnoverRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GenerateGoodsTurnoverRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *GenerateGoodsTurnoverRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


