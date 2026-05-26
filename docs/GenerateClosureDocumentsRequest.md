# GenerateClosureDocumentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **int64** | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями.  | 
**MonthOfYear** | [**ClosureDocumentsMonthOfYearDTO**](ClosureDocumentsMonthOfYearDTO.md) |  | 
**ContractTypes** | Pointer to [**[]ClosureDocumentsContractType**](ClosureDocumentsContractType.md) | Типы договоров, по которым нужны закрывающие документы.  Если их не указать, вернется архив с документами по всем найденным договорам.  | [optional] 

## Methods

### NewGenerateClosureDocumentsRequest

`func NewGenerateClosureDocumentsRequest(campaignId int64, monthOfYear ClosureDocumentsMonthOfYearDTO, ) *GenerateClosureDocumentsRequest`

NewGenerateClosureDocumentsRequest instantiates a new GenerateClosureDocumentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateClosureDocumentsRequestWithDefaults

`func NewGenerateClosureDocumentsRequestWithDefaults() *GenerateClosureDocumentsRequest`

NewGenerateClosureDocumentsRequestWithDefaults instantiates a new GenerateClosureDocumentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *GenerateClosureDocumentsRequest) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *GenerateClosureDocumentsRequest) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *GenerateClosureDocumentsRequest) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetMonthOfYear

`func (o *GenerateClosureDocumentsRequest) GetMonthOfYear() ClosureDocumentsMonthOfYearDTO`

GetMonthOfYear returns the MonthOfYear field if non-nil, zero value otherwise.

### GetMonthOfYearOk

`func (o *GenerateClosureDocumentsRequest) GetMonthOfYearOk() (*ClosureDocumentsMonthOfYearDTO, bool)`

GetMonthOfYearOk returns a tuple with the MonthOfYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthOfYear

`func (o *GenerateClosureDocumentsRequest) SetMonthOfYear(v ClosureDocumentsMonthOfYearDTO)`

SetMonthOfYear sets MonthOfYear field to given value.


### GetContractTypes

`func (o *GenerateClosureDocumentsRequest) GetContractTypes() []ClosureDocumentsContractType`

GetContractTypes returns the ContractTypes field if non-nil, zero value otherwise.

### GetContractTypesOk

`func (o *GenerateClosureDocumentsRequest) GetContractTypesOk() (*[]ClosureDocumentsContractType, bool)`

GetContractTypesOk returns a tuple with the ContractTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTypes

`func (o *GenerateClosureDocumentsRequest) SetContractTypes(v []ClosureDocumentsContractType)`

SetContractTypes sets ContractTypes field to given value.

### HasContractTypes

`func (o *GenerateClosureDocumentsRequest) HasContractTypes() bool`

HasContractTypes returns a boolean if a field has been set.

### SetContractTypesNil

`func (o *GenerateClosureDocumentsRequest) SetContractTypesNil(b bool)`

 SetContractTypesNil sets the value for ContractTypes to be an explicit nil

### UnsetContractTypes
`func (o *GenerateClosureDocumentsRequest) UnsetContractTypes()`

UnsetContractTypes ensures that no value is present for ContractTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


