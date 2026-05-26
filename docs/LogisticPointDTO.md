# LogisticPointDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogisticPointId** | **int64** | Идентификатор пункта выдачи.  Его можно узнать с помощью метода [POST v1/businesses/{businessId}/logistics-points](../../reference/logistic-points/getLogisticPoints.md).  | 
**Brand** | [**LogisticPointBrandType**](LogisticPointBrandType.md) |  | 
**Address** | [**LogisticPointAddressDTO**](LogisticPointAddressDTO.md) |  | 
**WorkingSchedule** | [**LogisticPointScheduleDTO**](LogisticPointScheduleDTO.md) |  | 
**DeliveryRestrictions** | [**LogisticPointDeliveryRestrictionDTO**](LogisticPointDeliveryRestrictionDTO.md) |  | 
**Features** | Pointer to [**[]LogisticPointFeatureType**](LogisticPointFeatureType.md) | Свойства пункта выдачи. | [optional] 
**StoragePeriod** | **int64** | Срок хранения заказа в пункте выдачи.  Указывается в днях.  | 

## Methods

### NewLogisticPointDTO

`func NewLogisticPointDTO(logisticPointId int64, brand LogisticPointBrandType, address LogisticPointAddressDTO, workingSchedule LogisticPointScheduleDTO, deliveryRestrictions LogisticPointDeliveryRestrictionDTO, storagePeriod int64, ) *LogisticPointDTO`

NewLogisticPointDTO instantiates a new LogisticPointDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogisticPointDTOWithDefaults

`func NewLogisticPointDTOWithDefaults() *LogisticPointDTO`

NewLogisticPointDTOWithDefaults instantiates a new LogisticPointDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogisticPointId

`func (o *LogisticPointDTO) GetLogisticPointId() int64`

GetLogisticPointId returns the LogisticPointId field if non-nil, zero value otherwise.

### GetLogisticPointIdOk

`func (o *LogisticPointDTO) GetLogisticPointIdOk() (*int64, bool)`

GetLogisticPointIdOk returns a tuple with the LogisticPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogisticPointId

`func (o *LogisticPointDTO) SetLogisticPointId(v int64)`

SetLogisticPointId sets LogisticPointId field to given value.


### GetBrand

`func (o *LogisticPointDTO) GetBrand() LogisticPointBrandType`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *LogisticPointDTO) GetBrandOk() (*LogisticPointBrandType, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *LogisticPointDTO) SetBrand(v LogisticPointBrandType)`

SetBrand sets Brand field to given value.


### GetAddress

`func (o *LogisticPointDTO) GetAddress() LogisticPointAddressDTO`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LogisticPointDTO) GetAddressOk() (*LogisticPointAddressDTO, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LogisticPointDTO) SetAddress(v LogisticPointAddressDTO)`

SetAddress sets Address field to given value.


### GetWorkingSchedule

`func (o *LogisticPointDTO) GetWorkingSchedule() LogisticPointScheduleDTO`

GetWorkingSchedule returns the WorkingSchedule field if non-nil, zero value otherwise.

### GetWorkingScheduleOk

`func (o *LogisticPointDTO) GetWorkingScheduleOk() (*LogisticPointScheduleDTO, bool)`

GetWorkingScheduleOk returns a tuple with the WorkingSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingSchedule

`func (o *LogisticPointDTO) SetWorkingSchedule(v LogisticPointScheduleDTO)`

SetWorkingSchedule sets WorkingSchedule field to given value.


### GetDeliveryRestrictions

`func (o *LogisticPointDTO) GetDeliveryRestrictions() LogisticPointDeliveryRestrictionDTO`

GetDeliveryRestrictions returns the DeliveryRestrictions field if non-nil, zero value otherwise.

### GetDeliveryRestrictionsOk

`func (o *LogisticPointDTO) GetDeliveryRestrictionsOk() (*LogisticPointDeliveryRestrictionDTO, bool)`

GetDeliveryRestrictionsOk returns a tuple with the DeliveryRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryRestrictions

`func (o *LogisticPointDTO) SetDeliveryRestrictions(v LogisticPointDeliveryRestrictionDTO)`

SetDeliveryRestrictions sets DeliveryRestrictions field to given value.


### GetFeatures

`func (o *LogisticPointDTO) GetFeatures() []LogisticPointFeatureType`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *LogisticPointDTO) GetFeaturesOk() (*[]LogisticPointFeatureType, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *LogisticPointDTO) SetFeatures(v []LogisticPointFeatureType)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *LogisticPointDTO) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeaturesNil

`func (o *LogisticPointDTO) SetFeaturesNil(b bool)`

 SetFeaturesNil sets the value for Features to be an explicit nil

### UnsetFeatures
`func (o *LogisticPointDTO) UnsetFeatures()`

UnsetFeatures ensures that no value is present for Features, not even an explicit nil
### GetStoragePeriod

`func (o *LogisticPointDTO) GetStoragePeriod() int64`

GetStoragePeriod returns the StoragePeriod field if non-nil, zero value otherwise.

### GetStoragePeriodOk

`func (o *LogisticPointDTO) GetStoragePeriodOk() (*int64, bool)`

GetStoragePeriodOk returns a tuple with the StoragePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePeriod

`func (o *LogisticPointDTO) SetStoragePeriod(v int64)`

SetStoragePeriod sets StoragePeriod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


