# UpdateWarehouseStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Статус склада:  * &#x60;true&#x60; — включен. * &#x60;false&#x60; — отключен.  | 

## Methods

### NewUpdateWarehouseStatusRequest

`func NewUpdateWarehouseStatusRequest(enabled bool, ) *UpdateWarehouseStatusRequest`

NewUpdateWarehouseStatusRequest instantiates a new UpdateWarehouseStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWarehouseStatusRequestWithDefaults

`func NewUpdateWarehouseStatusRequestWithDefaults() *UpdateWarehouseStatusRequest`

NewUpdateWarehouseStatusRequestWithDefaults instantiates a new UpdateWarehouseStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateWarehouseStatusRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateWarehouseStatusRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateWarehouseStatusRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


