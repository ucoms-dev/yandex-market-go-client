# UpdateWarehouseModelStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerWarehouseId** | **int64** | Идентификатор склада.  Чтобы узнать идентификатор склада, воспользуйтесь запросом [POST v3/businesses/{businessId}/warehouses](../../reference/warehouses/getPartnerWarehouses.md).  | 
**Model** | [**WarehouseProgramType**](WarehouseProgramType.md) |  | 
**Enabled** | **bool** | Статус модели работы склада:  * &#x60;true&#x60; — включена. * &#x60;false&#x60; — отключена.  | 

## Methods

### NewUpdateWarehouseModelStatusRequest

`func NewUpdateWarehouseModelStatusRequest(partnerWarehouseId int64, model WarehouseProgramType, enabled bool, ) *UpdateWarehouseModelStatusRequest`

NewUpdateWarehouseModelStatusRequest instantiates a new UpdateWarehouseModelStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWarehouseModelStatusRequestWithDefaults

`func NewUpdateWarehouseModelStatusRequestWithDefaults() *UpdateWarehouseModelStatusRequest`

NewUpdateWarehouseModelStatusRequestWithDefaults instantiates a new UpdateWarehouseModelStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerWarehouseId

`func (o *UpdateWarehouseModelStatusRequest) GetPartnerWarehouseId() int64`

GetPartnerWarehouseId returns the PartnerWarehouseId field if non-nil, zero value otherwise.

### GetPartnerWarehouseIdOk

`func (o *UpdateWarehouseModelStatusRequest) GetPartnerWarehouseIdOk() (*int64, bool)`

GetPartnerWarehouseIdOk returns a tuple with the PartnerWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerWarehouseId

`func (o *UpdateWarehouseModelStatusRequest) SetPartnerWarehouseId(v int64)`

SetPartnerWarehouseId sets PartnerWarehouseId field to given value.


### GetModel

`func (o *UpdateWarehouseModelStatusRequest) GetModel() WarehouseProgramType`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *UpdateWarehouseModelStatusRequest) GetModelOk() (*WarehouseProgramType, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *UpdateWarehouseModelStatusRequest) SetModel(v WarehouseProgramType)`

SetModel sets Model field to given value.


### GetEnabled

`func (o *UpdateWarehouseModelStatusRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateWarehouseModelStatusRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateWarehouseModelStatusRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


