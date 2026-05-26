# GenerateClosureDocumentsDetalizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | Pointer to **int64** | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями.  | [optional] 
**MonthOfYear** | [**ClosureDocumentsMonthOfYearDTO**](ClosureDocumentsMonthOfYearDTO.md) |  | 
**ContractType** | [**ClosureDocumentsContractType**](ClosureDocumentsContractType.md) |  | 

## Methods

### NewGenerateClosureDocumentsDetalizationRequest

`func NewGenerateClosureDocumentsDetalizationRequest(monthOfYear ClosureDocumentsMonthOfYearDTO, contractType ClosureDocumentsContractType, ) *GenerateClosureDocumentsDetalizationRequest`

NewGenerateClosureDocumentsDetalizationRequest instantiates a new GenerateClosureDocumentsDetalizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateClosureDocumentsDetalizationRequestWithDefaults

`func NewGenerateClosureDocumentsDetalizationRequestWithDefaults() *GenerateClosureDocumentsDetalizationRequest`

NewGenerateClosureDocumentsDetalizationRequestWithDefaults instantiates a new GenerateClosureDocumentsDetalizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *GenerateClosureDocumentsDetalizationRequest) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *GenerateClosureDocumentsDetalizationRequest) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *GenerateClosureDocumentsDetalizationRequest) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *GenerateClosureDocumentsDetalizationRequest) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetMonthOfYear

`func (o *GenerateClosureDocumentsDetalizationRequest) GetMonthOfYear() ClosureDocumentsMonthOfYearDTO`

GetMonthOfYear returns the MonthOfYear field if non-nil, zero value otherwise.

### GetMonthOfYearOk

`func (o *GenerateClosureDocumentsDetalizationRequest) GetMonthOfYearOk() (*ClosureDocumentsMonthOfYearDTO, bool)`

GetMonthOfYearOk returns a tuple with the MonthOfYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthOfYear

`func (o *GenerateClosureDocumentsDetalizationRequest) SetMonthOfYear(v ClosureDocumentsMonthOfYearDTO)`

SetMonthOfYear sets MonthOfYear field to given value.


### GetContractType

`func (o *GenerateClosureDocumentsDetalizationRequest) GetContractType() ClosureDocumentsContractType`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *GenerateClosureDocumentsDetalizationRequest) GetContractTypeOk() (*ClosureDocumentsContractType, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *GenerateClosureDocumentsDetalizationRequest) SetContractType(v ClosureDocumentsContractType)`

SetContractType sets ContractType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


