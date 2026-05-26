# GetReturnAvailableDecisionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **int64** | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями.  | 
**ReturnId** | **int64** | Идентификатор невыкупа или возврата. | 

## Methods

### NewGetReturnAvailableDecisionsRequest

`func NewGetReturnAvailableDecisionsRequest(campaignId int64, returnId int64, ) *GetReturnAvailableDecisionsRequest`

NewGetReturnAvailableDecisionsRequest instantiates a new GetReturnAvailableDecisionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReturnAvailableDecisionsRequestWithDefaults

`func NewGetReturnAvailableDecisionsRequestWithDefaults() *GetReturnAvailableDecisionsRequest`

NewGetReturnAvailableDecisionsRequestWithDefaults instantiates a new GetReturnAvailableDecisionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *GetReturnAvailableDecisionsRequest) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *GetReturnAvailableDecisionsRequest) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *GetReturnAvailableDecisionsRequest) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetReturnId

`func (o *GetReturnAvailableDecisionsRequest) GetReturnId() int64`

GetReturnId returns the ReturnId field if non-nil, zero value otherwise.

### GetReturnIdOk

`func (o *GetReturnAvailableDecisionsRequest) GetReturnIdOk() (*int64, bool)`

GetReturnIdOk returns a tuple with the ReturnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnId

`func (o *GetReturnAvailableDecisionsRequest) SetReturnId(v int64)`

SetReturnId sets ReturnId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


