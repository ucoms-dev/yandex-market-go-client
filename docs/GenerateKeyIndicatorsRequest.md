# GenerateKeyIndicatorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | Pointer to **int64** | Идентификатор кабинета. {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) {% endif %}  | [optional] 
**CampaignId** | Pointer to **int64** | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями.  | [optional] 
**DetalizationLevel** | [**KeyIndicatorsReportDetalizationLevelType**](KeyIndicatorsReportDetalizationLevelType.md) |  | 

## Methods

### NewGenerateKeyIndicatorsRequest

`func NewGenerateKeyIndicatorsRequest(detalizationLevel KeyIndicatorsReportDetalizationLevelType, ) *GenerateKeyIndicatorsRequest`

NewGenerateKeyIndicatorsRequest instantiates a new GenerateKeyIndicatorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateKeyIndicatorsRequestWithDefaults

`func NewGenerateKeyIndicatorsRequestWithDefaults() *GenerateKeyIndicatorsRequest`

NewGenerateKeyIndicatorsRequestWithDefaults instantiates a new GenerateKeyIndicatorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessId

`func (o *GenerateKeyIndicatorsRequest) GetBusinessId() int64`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *GenerateKeyIndicatorsRequest) GetBusinessIdOk() (*int64, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *GenerateKeyIndicatorsRequest) SetBusinessId(v int64)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *GenerateKeyIndicatorsRequest) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCampaignId

`func (o *GenerateKeyIndicatorsRequest) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *GenerateKeyIndicatorsRequest) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *GenerateKeyIndicatorsRequest) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *GenerateKeyIndicatorsRequest) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetDetalizationLevel

`func (o *GenerateKeyIndicatorsRequest) GetDetalizationLevel() KeyIndicatorsReportDetalizationLevelType`

GetDetalizationLevel returns the DetalizationLevel field if non-nil, zero value otherwise.

### GetDetalizationLevelOk

`func (o *GenerateKeyIndicatorsRequest) GetDetalizationLevelOk() (*KeyIndicatorsReportDetalizationLevelType, bool)`

GetDetalizationLevelOk returns a tuple with the DetalizationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetalizationLevel

`func (o *GenerateKeyIndicatorsRequest) SetDetalizationLevel(v KeyIndicatorsReportDetalizationLevelType)`

SetDetalizationLevel sets DetalizationLevel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


