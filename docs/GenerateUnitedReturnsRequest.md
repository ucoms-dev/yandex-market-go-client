# GenerateUnitedReturnsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | **int64** | Идентификатор кабинета. | 
**DateFrom** | **string** | Начало периода, включительно. | 
**DateTo** | **string** | Конец периода, включительно. | 
**CampaignIds** | Pointer to **[]int64** | Список идентификаторов кампании тех магазинов, которые нужны в отчете.  Их можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не используйте вместо них идентификаторы магазинов, которые указаны в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | [optional] 
**ReturnType** | Pointer to [**ReturnType**](ReturnType.md) |  | [optional] 
**ReturnStatusTypes** | Pointer to [**[]ReturnShipmentStatusType**](ReturnShipmentStatusType.md) | Статусы передачи возвратов, которые нужны в отчете.  Если их не указать, вернется информация по всем возвратам.  | [optional] 

## Methods

### NewGenerateUnitedReturnsRequest

`func NewGenerateUnitedReturnsRequest(businessId int64, dateFrom string, dateTo string, ) *GenerateUnitedReturnsRequest`

NewGenerateUnitedReturnsRequest instantiates a new GenerateUnitedReturnsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateUnitedReturnsRequestWithDefaults

`func NewGenerateUnitedReturnsRequestWithDefaults() *GenerateUnitedReturnsRequest`

NewGenerateUnitedReturnsRequestWithDefaults instantiates a new GenerateUnitedReturnsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessId

`func (o *GenerateUnitedReturnsRequest) GetBusinessId() int64`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *GenerateUnitedReturnsRequest) GetBusinessIdOk() (*int64, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *GenerateUnitedReturnsRequest) SetBusinessId(v int64)`

SetBusinessId sets BusinessId field to given value.


### GetDateFrom

`func (o *GenerateUnitedReturnsRequest) GetDateFrom() string`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *GenerateUnitedReturnsRequest) GetDateFromOk() (*string, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *GenerateUnitedReturnsRequest) SetDateFrom(v string)`

SetDateFrom sets DateFrom field to given value.


### GetDateTo

`func (o *GenerateUnitedReturnsRequest) GetDateTo() string`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *GenerateUnitedReturnsRequest) GetDateToOk() (*string, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *GenerateUnitedReturnsRequest) SetDateTo(v string)`

SetDateTo sets DateTo field to given value.


### GetCampaignIds

`func (o *GenerateUnitedReturnsRequest) GetCampaignIds() []int64`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *GenerateUnitedReturnsRequest) GetCampaignIdsOk() (*[]int64, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignIds

`func (o *GenerateUnitedReturnsRequest) SetCampaignIds(v []int64)`

SetCampaignIds sets CampaignIds field to given value.

### HasCampaignIds

`func (o *GenerateUnitedReturnsRequest) HasCampaignIds() bool`

HasCampaignIds returns a boolean if a field has been set.

### SetCampaignIdsNil

`func (o *GenerateUnitedReturnsRequest) SetCampaignIdsNil(b bool)`

 SetCampaignIdsNil sets the value for CampaignIds to be an explicit nil

### UnsetCampaignIds
`func (o *GenerateUnitedReturnsRequest) UnsetCampaignIds()`

UnsetCampaignIds ensures that no value is present for CampaignIds, not even an explicit nil
### GetReturnType

`func (o *GenerateUnitedReturnsRequest) GetReturnType() ReturnType`

GetReturnType returns the ReturnType field if non-nil, zero value otherwise.

### GetReturnTypeOk

`func (o *GenerateUnitedReturnsRequest) GetReturnTypeOk() (*ReturnType, bool)`

GetReturnTypeOk returns a tuple with the ReturnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnType

`func (o *GenerateUnitedReturnsRequest) SetReturnType(v ReturnType)`

SetReturnType sets ReturnType field to given value.

### HasReturnType

`func (o *GenerateUnitedReturnsRequest) HasReturnType() bool`

HasReturnType returns a boolean if a field has been set.

### GetReturnStatusTypes

`func (o *GenerateUnitedReturnsRequest) GetReturnStatusTypes() []ReturnShipmentStatusType`

GetReturnStatusTypes returns the ReturnStatusTypes field if non-nil, zero value otherwise.

### GetReturnStatusTypesOk

`func (o *GenerateUnitedReturnsRequest) GetReturnStatusTypesOk() (*[]ReturnShipmentStatusType, bool)`

GetReturnStatusTypesOk returns a tuple with the ReturnStatusTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnStatusTypes

`func (o *GenerateUnitedReturnsRequest) SetReturnStatusTypes(v []ReturnShipmentStatusType)`

SetReturnStatusTypes sets ReturnStatusTypes field to given value.

### HasReturnStatusTypes

`func (o *GenerateUnitedReturnsRequest) HasReturnStatusTypes() bool`

HasReturnStatusTypes returns a boolean if a field has been set.

### SetReturnStatusTypesNil

`func (o *GenerateUnitedReturnsRequest) SetReturnStatusTypesNil(b bool)`

 SetReturnStatusTypesNil sets the value for ReturnStatusTypes to be an explicit nil

### UnsetReturnStatusTypes
`func (o *GenerateUnitedReturnsRequest) UnsetReturnStatusTypes()`

UnsetReturnStatusTypes ensures that no value is present for ReturnStatusTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


