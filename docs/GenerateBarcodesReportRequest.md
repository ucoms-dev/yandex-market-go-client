# GenerateBarcodesReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **int64** | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями.  | 
**BarcodeFormat** | [**BarcodeFormatType**](BarcodeFormatType.md) |  | 
**BarcodeOfferInfos** | Pointer to [**[]BarcodeOfferInfoDTO**](BarcodeOfferInfoDTO.md) | Список товаров.  Передайте этот параметр и уникальные &#x60;offerId&#x60;, чтобы вернуть файл со штрихкодами конкретных товаров.  В запросе обязательно должен быть либо &#x60;barcodeOfferInfos&#x60;, либо &#x60;supplyRequestId&#x60;, но не оба сразу.  Если передается информация о товаре, у которого несколько штрихкодов, все штрихкоды будут добавлены в файл, их количество будет задано параметром &#x60;barcodeCount&#x60;.  | [optional] 
**SupplyRequestId** | Pointer to **int64** | Идентификатор заявки.  {% note warning \&quot;Используется только в API\&quot; %}  По нему не получится найти заявки в кабинете продавца на Маркете. Для этого используйте &#x60;marketplaceRequestId&#x60; или &#x60;warehouseRequestId&#x60;.  {% endnote %}  | [optional] 

## Methods

### NewGenerateBarcodesReportRequest

`func NewGenerateBarcodesReportRequest(campaignId int64, barcodeFormat BarcodeFormatType, ) *GenerateBarcodesReportRequest`

NewGenerateBarcodesReportRequest instantiates a new GenerateBarcodesReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateBarcodesReportRequestWithDefaults

`func NewGenerateBarcodesReportRequestWithDefaults() *GenerateBarcodesReportRequest`

NewGenerateBarcodesReportRequestWithDefaults instantiates a new GenerateBarcodesReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *GenerateBarcodesReportRequest) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *GenerateBarcodesReportRequest) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *GenerateBarcodesReportRequest) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetBarcodeFormat

`func (o *GenerateBarcodesReportRequest) GetBarcodeFormat() BarcodeFormatType`

GetBarcodeFormat returns the BarcodeFormat field if non-nil, zero value otherwise.

### GetBarcodeFormatOk

`func (o *GenerateBarcodesReportRequest) GetBarcodeFormatOk() (*BarcodeFormatType, bool)`

GetBarcodeFormatOk returns a tuple with the BarcodeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeFormat

`func (o *GenerateBarcodesReportRequest) SetBarcodeFormat(v BarcodeFormatType)`

SetBarcodeFormat sets BarcodeFormat field to given value.


### GetBarcodeOfferInfos

`func (o *GenerateBarcodesReportRequest) GetBarcodeOfferInfos() []BarcodeOfferInfoDTO`

GetBarcodeOfferInfos returns the BarcodeOfferInfos field if non-nil, zero value otherwise.

### GetBarcodeOfferInfosOk

`func (o *GenerateBarcodesReportRequest) GetBarcodeOfferInfosOk() (*[]BarcodeOfferInfoDTO, bool)`

GetBarcodeOfferInfosOk returns a tuple with the BarcodeOfferInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeOfferInfos

`func (o *GenerateBarcodesReportRequest) SetBarcodeOfferInfos(v []BarcodeOfferInfoDTO)`

SetBarcodeOfferInfos sets BarcodeOfferInfos field to given value.

### HasBarcodeOfferInfos

`func (o *GenerateBarcodesReportRequest) HasBarcodeOfferInfos() bool`

HasBarcodeOfferInfos returns a boolean if a field has been set.

### SetBarcodeOfferInfosNil

`func (o *GenerateBarcodesReportRequest) SetBarcodeOfferInfosNil(b bool)`

 SetBarcodeOfferInfosNil sets the value for BarcodeOfferInfos to be an explicit nil

### UnsetBarcodeOfferInfos
`func (o *GenerateBarcodesReportRequest) UnsetBarcodeOfferInfos()`

UnsetBarcodeOfferInfos ensures that no value is present for BarcodeOfferInfos, not even an explicit nil
### GetSupplyRequestId

`func (o *GenerateBarcodesReportRequest) GetSupplyRequestId() int64`

GetSupplyRequestId returns the SupplyRequestId field if non-nil, zero value otherwise.

### GetSupplyRequestIdOk

`func (o *GenerateBarcodesReportRequest) GetSupplyRequestIdOk() (*int64, bool)`

GetSupplyRequestIdOk returns a tuple with the SupplyRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplyRequestId

`func (o *GenerateBarcodesReportRequest) SetSupplyRequestId(v int64)`

SetSupplyRequestId sets SupplyRequestId field to given value.

### HasSupplyRequestId

`func (o *GenerateBarcodesReportRequest) HasSupplyRequestId() bool`

HasSupplyRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


