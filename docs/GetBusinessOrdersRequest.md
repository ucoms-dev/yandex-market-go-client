# GetBusinessOrdersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderIds** | Pointer to **[]int64** | Идентификаторы заказов. | [optional] 
**ExternalOrderIds** | Pointer to **[]string** | Внешние идентификаторы заказов. | [optional] 
**ProgramTypes** | Pointer to [**[]SellingProgramType**](SellingProgramType.md) | Модели работы магазина на Маркете. | [optional] 
**CampaignIds** | Pointer to **[]int64** | Идентификаторы кампаний магазинов. | [optional] 
**Statuses** | Pointer to [**[]OrderStatusType**](OrderStatusType.md) | Статусы заказов. | [optional] 
**Substatuses** | Pointer to [**[]OrderSubstatusType**](OrderSubstatusType.md) | Этапы обработки или причины отмены заказов. | [optional] 
**Dates** | Pointer to [**OrderDatesFilterDTO**](OrderDatesFilterDTO.md) |  | [optional] 
**Fake** | Pointer to **bool** | Тип заказа:  * &#x60;false&#x60; — настоящий заказ покупателя.  * &#x60;true&#x60; — [тестовый заказ](../../concepts/sandbox.md) Маркета.  | [optional] 
**WaitingForCancellationApprove** | Pointer to **bool** | **Только для модели DBS**  Фильтр для получения заказов, по которым есть запросы на отмену.  При значении &#x60;true&#x60; возвращаются только те заказы, которые находятся в статусе &#x60;DELIVERY&#x60; или &#x60;PICKUP&#x60;, и пользователи решили их отменить.  | [optional] 
**SourcePlatforms** | Pointer to [**[]OrderSourcePlatformType**](OrderSourcePlatformType.md) | Площадки-источники заказов. | [optional] 

## Methods

### NewGetBusinessOrdersRequest

`func NewGetBusinessOrdersRequest() *GetBusinessOrdersRequest`

NewGetBusinessOrdersRequest instantiates a new GetBusinessOrdersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBusinessOrdersRequestWithDefaults

`func NewGetBusinessOrdersRequestWithDefaults() *GetBusinessOrdersRequest`

NewGetBusinessOrdersRequestWithDefaults instantiates a new GetBusinessOrdersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderIds

`func (o *GetBusinessOrdersRequest) GetOrderIds() []int64`

GetOrderIds returns the OrderIds field if non-nil, zero value otherwise.

### GetOrderIdsOk

`func (o *GetBusinessOrdersRequest) GetOrderIdsOk() (*[]int64, bool)`

GetOrderIdsOk returns a tuple with the OrderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderIds

`func (o *GetBusinessOrdersRequest) SetOrderIds(v []int64)`

SetOrderIds sets OrderIds field to given value.

### HasOrderIds

`func (o *GetBusinessOrdersRequest) HasOrderIds() bool`

HasOrderIds returns a boolean if a field has been set.

### SetOrderIdsNil

`func (o *GetBusinessOrdersRequest) SetOrderIdsNil(b bool)`

 SetOrderIdsNil sets the value for OrderIds to be an explicit nil

### UnsetOrderIds
`func (o *GetBusinessOrdersRequest) UnsetOrderIds()`

UnsetOrderIds ensures that no value is present for OrderIds, not even an explicit nil
### GetExternalOrderIds

`func (o *GetBusinessOrdersRequest) GetExternalOrderIds() []string`

GetExternalOrderIds returns the ExternalOrderIds field if non-nil, zero value otherwise.

### GetExternalOrderIdsOk

`func (o *GetBusinessOrdersRequest) GetExternalOrderIdsOk() (*[]string, bool)`

GetExternalOrderIdsOk returns a tuple with the ExternalOrderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOrderIds

`func (o *GetBusinessOrdersRequest) SetExternalOrderIds(v []string)`

SetExternalOrderIds sets ExternalOrderIds field to given value.

### HasExternalOrderIds

`func (o *GetBusinessOrdersRequest) HasExternalOrderIds() bool`

HasExternalOrderIds returns a boolean if a field has been set.

### SetExternalOrderIdsNil

`func (o *GetBusinessOrdersRequest) SetExternalOrderIdsNil(b bool)`

 SetExternalOrderIdsNil sets the value for ExternalOrderIds to be an explicit nil

### UnsetExternalOrderIds
`func (o *GetBusinessOrdersRequest) UnsetExternalOrderIds()`

UnsetExternalOrderIds ensures that no value is present for ExternalOrderIds, not even an explicit nil
### GetProgramTypes

`func (o *GetBusinessOrdersRequest) GetProgramTypes() []SellingProgramType`

GetProgramTypes returns the ProgramTypes field if non-nil, zero value otherwise.

### GetProgramTypesOk

`func (o *GetBusinessOrdersRequest) GetProgramTypesOk() (*[]SellingProgramType, bool)`

GetProgramTypesOk returns a tuple with the ProgramTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramTypes

`func (o *GetBusinessOrdersRequest) SetProgramTypes(v []SellingProgramType)`

SetProgramTypes sets ProgramTypes field to given value.

### HasProgramTypes

`func (o *GetBusinessOrdersRequest) HasProgramTypes() bool`

HasProgramTypes returns a boolean if a field has been set.

### SetProgramTypesNil

`func (o *GetBusinessOrdersRequest) SetProgramTypesNil(b bool)`

 SetProgramTypesNil sets the value for ProgramTypes to be an explicit nil

### UnsetProgramTypes
`func (o *GetBusinessOrdersRequest) UnsetProgramTypes()`

UnsetProgramTypes ensures that no value is present for ProgramTypes, not even an explicit nil
### GetCampaignIds

`func (o *GetBusinessOrdersRequest) GetCampaignIds() []int64`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *GetBusinessOrdersRequest) GetCampaignIdsOk() (*[]int64, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignIds

`func (o *GetBusinessOrdersRequest) SetCampaignIds(v []int64)`

SetCampaignIds sets CampaignIds field to given value.

### HasCampaignIds

`func (o *GetBusinessOrdersRequest) HasCampaignIds() bool`

HasCampaignIds returns a boolean if a field has been set.

### SetCampaignIdsNil

`func (o *GetBusinessOrdersRequest) SetCampaignIdsNil(b bool)`

 SetCampaignIdsNil sets the value for CampaignIds to be an explicit nil

### UnsetCampaignIds
`func (o *GetBusinessOrdersRequest) UnsetCampaignIds()`

UnsetCampaignIds ensures that no value is present for CampaignIds, not even an explicit nil
### GetStatuses

`func (o *GetBusinessOrdersRequest) GetStatuses() []OrderStatusType`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *GetBusinessOrdersRequest) GetStatusesOk() (*[]OrderStatusType, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *GetBusinessOrdersRequest) SetStatuses(v []OrderStatusType)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *GetBusinessOrdersRequest) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### SetStatusesNil

`func (o *GetBusinessOrdersRequest) SetStatusesNil(b bool)`

 SetStatusesNil sets the value for Statuses to be an explicit nil

### UnsetStatuses
`func (o *GetBusinessOrdersRequest) UnsetStatuses()`

UnsetStatuses ensures that no value is present for Statuses, not even an explicit nil
### GetSubstatuses

`func (o *GetBusinessOrdersRequest) GetSubstatuses() []OrderSubstatusType`

GetSubstatuses returns the Substatuses field if non-nil, zero value otherwise.

### GetSubstatusesOk

`func (o *GetBusinessOrdersRequest) GetSubstatusesOk() (*[]OrderSubstatusType, bool)`

GetSubstatusesOk returns a tuple with the Substatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstatuses

`func (o *GetBusinessOrdersRequest) SetSubstatuses(v []OrderSubstatusType)`

SetSubstatuses sets Substatuses field to given value.

### HasSubstatuses

`func (o *GetBusinessOrdersRequest) HasSubstatuses() bool`

HasSubstatuses returns a boolean if a field has been set.

### SetSubstatusesNil

`func (o *GetBusinessOrdersRequest) SetSubstatusesNil(b bool)`

 SetSubstatusesNil sets the value for Substatuses to be an explicit nil

### UnsetSubstatuses
`func (o *GetBusinessOrdersRequest) UnsetSubstatuses()`

UnsetSubstatuses ensures that no value is present for Substatuses, not even an explicit nil
### GetDates

`func (o *GetBusinessOrdersRequest) GetDates() OrderDatesFilterDTO`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *GetBusinessOrdersRequest) GetDatesOk() (*OrderDatesFilterDTO, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *GetBusinessOrdersRequest) SetDates(v OrderDatesFilterDTO)`

SetDates sets Dates field to given value.

### HasDates

`func (o *GetBusinessOrdersRequest) HasDates() bool`

HasDates returns a boolean if a field has been set.

### GetFake

`func (o *GetBusinessOrdersRequest) GetFake() bool`

GetFake returns the Fake field if non-nil, zero value otherwise.

### GetFakeOk

`func (o *GetBusinessOrdersRequest) GetFakeOk() (*bool, bool)`

GetFakeOk returns a tuple with the Fake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFake

`func (o *GetBusinessOrdersRequest) SetFake(v bool)`

SetFake sets Fake field to given value.

### HasFake

`func (o *GetBusinessOrdersRequest) HasFake() bool`

HasFake returns a boolean if a field has been set.

### GetWaitingForCancellationApprove

`func (o *GetBusinessOrdersRequest) GetWaitingForCancellationApprove() bool`

GetWaitingForCancellationApprove returns the WaitingForCancellationApprove field if non-nil, zero value otherwise.

### GetWaitingForCancellationApproveOk

`func (o *GetBusinessOrdersRequest) GetWaitingForCancellationApproveOk() (*bool, bool)`

GetWaitingForCancellationApproveOk returns a tuple with the WaitingForCancellationApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingForCancellationApprove

`func (o *GetBusinessOrdersRequest) SetWaitingForCancellationApprove(v bool)`

SetWaitingForCancellationApprove sets WaitingForCancellationApprove field to given value.

### HasWaitingForCancellationApprove

`func (o *GetBusinessOrdersRequest) HasWaitingForCancellationApprove() bool`

HasWaitingForCancellationApprove returns a boolean if a field has been set.

### GetSourcePlatforms

`func (o *GetBusinessOrdersRequest) GetSourcePlatforms() []OrderSourcePlatformType`

GetSourcePlatforms returns the SourcePlatforms field if non-nil, zero value otherwise.

### GetSourcePlatformsOk

`func (o *GetBusinessOrdersRequest) GetSourcePlatformsOk() (*[]OrderSourcePlatformType, bool)`

GetSourcePlatformsOk returns a tuple with the SourcePlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePlatforms

`func (o *GetBusinessOrdersRequest) SetSourcePlatforms(v []OrderSourcePlatformType)`

SetSourcePlatforms sets SourcePlatforms field to given value.

### HasSourcePlatforms

`func (o *GetBusinessOrdersRequest) HasSourcePlatforms() bool`

HasSourcePlatforms returns a boolean if a field has been set.

### SetSourcePlatformsNil

`func (o *GetBusinessOrdersRequest) SetSourcePlatformsNil(b bool)`

 SetSourcePlatformsNil sets the value for SourcePlatforms to be an explicit nil

### UnsetSourcePlatforms
`func (o *GetBusinessOrdersRequest) UnsetSourcePlatforms()`

UnsetSourcePlatforms ensures that no value is present for SourcePlatforms, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


