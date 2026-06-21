# BusinessOrderPickupDeliveryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**BusinessOrderDeliveryAddressDTO**](BusinessOrderDeliveryAddressDTO.md) |  | [optional] 
**Region** | Pointer to [**RegionDTO**](RegionDTO.md) |  | [optional] 
**LogisticPointId** | Pointer to **int64** | Идентификатор пункта выдачи.  Его можно узнать с помощью метода [POST v1/businesses/{businessId}/logistics-points](../../reference/logistic-points/getLogisticPoints.md).  | [optional] 
**OutletCode** | Pointer to **string** | Идентификатор пункта самовывоза, присвоенный магазином. | [optional] 
**OutletStorageLimitDate** | Pointer to **string** | Дата, до которой заказ будет храниться в пункте выдачи. Возвращается, когда заказ переходит в статус &#x60;PICKUP&#x60;.  Один раз дату можно поменять с помощью метода [PUT v2/campaigns/{campaignId}/orders/{orderId}/delivery/storage-limit](../../reference/order-delivery/updateOrderStorageLimit.md).  Формат даты: &#x60;ГГГГ-ММ-ДД&#x60;.  | [optional] 

## Methods

### NewBusinessOrderPickupDeliveryDTO

`func NewBusinessOrderPickupDeliveryDTO() *BusinessOrderPickupDeliveryDTO`

NewBusinessOrderPickupDeliveryDTO instantiates a new BusinessOrderPickupDeliveryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessOrderPickupDeliveryDTOWithDefaults

`func NewBusinessOrderPickupDeliveryDTOWithDefaults() *BusinessOrderPickupDeliveryDTO`

NewBusinessOrderPickupDeliveryDTOWithDefaults instantiates a new BusinessOrderPickupDeliveryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *BusinessOrderPickupDeliveryDTO) GetAddress() BusinessOrderDeliveryAddressDTO`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BusinessOrderPickupDeliveryDTO) GetAddressOk() (*BusinessOrderDeliveryAddressDTO, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BusinessOrderPickupDeliveryDTO) SetAddress(v BusinessOrderDeliveryAddressDTO)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BusinessOrderPickupDeliveryDTO) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetRegion

`func (o *BusinessOrderPickupDeliveryDTO) GetRegion() RegionDTO`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BusinessOrderPickupDeliveryDTO) GetRegionOk() (*RegionDTO, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BusinessOrderPickupDeliveryDTO) SetRegion(v RegionDTO)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *BusinessOrderPickupDeliveryDTO) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetLogisticPointId

`func (o *BusinessOrderPickupDeliveryDTO) GetLogisticPointId() int64`

GetLogisticPointId returns the LogisticPointId field if non-nil, zero value otherwise.

### GetLogisticPointIdOk

`func (o *BusinessOrderPickupDeliveryDTO) GetLogisticPointIdOk() (*int64, bool)`

GetLogisticPointIdOk returns a tuple with the LogisticPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogisticPointId

`func (o *BusinessOrderPickupDeliveryDTO) SetLogisticPointId(v int64)`

SetLogisticPointId sets LogisticPointId field to given value.

### HasLogisticPointId

`func (o *BusinessOrderPickupDeliveryDTO) HasLogisticPointId() bool`

HasLogisticPointId returns a boolean if a field has been set.

### GetOutletCode

`func (o *BusinessOrderPickupDeliveryDTO) GetOutletCode() string`

GetOutletCode returns the OutletCode field if non-nil, zero value otherwise.

### GetOutletCodeOk

`func (o *BusinessOrderPickupDeliveryDTO) GetOutletCodeOk() (*string, bool)`

GetOutletCodeOk returns a tuple with the OutletCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutletCode

`func (o *BusinessOrderPickupDeliveryDTO) SetOutletCode(v string)`

SetOutletCode sets OutletCode field to given value.

### HasOutletCode

`func (o *BusinessOrderPickupDeliveryDTO) HasOutletCode() bool`

HasOutletCode returns a boolean if a field has been set.

### GetOutletStorageLimitDate

`func (o *BusinessOrderPickupDeliveryDTO) GetOutletStorageLimitDate() string`

GetOutletStorageLimitDate returns the OutletStorageLimitDate field if non-nil, zero value otherwise.

### GetOutletStorageLimitDateOk

`func (o *BusinessOrderPickupDeliveryDTO) GetOutletStorageLimitDateOk() (*string, bool)`

GetOutletStorageLimitDateOk returns a tuple with the OutletStorageLimitDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutletStorageLimitDate

`func (o *BusinessOrderPickupDeliveryDTO) SetOutletStorageLimitDate(v string)`

SetOutletStorageLimitDate sets OutletStorageLimitDate field to given value.

### HasOutletStorageLimitDate

`func (o *BusinessOrderPickupDeliveryDTO) HasOutletStorageLimitDate() bool`

HasOutletStorageLimitDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


