# SupplyRequestLocationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedDate** | Pointer to **time.Time** | Дата и время поставки на склад или в ПВЗ. | [optional] 
**ServiceId** | **int64** | Идентификатор склада или логистического партнера ПВЗ. | 
**Name** | **string** | Название склада или ПВЗ. | 
**Address** | [**SupplyRequestLocationAddressDTO**](SupplyRequestLocationAddressDTO.md) |  | 
**Type** | [**SupplyRequestLocationType**](SupplyRequestLocationType.md) |  | 

## Methods

### NewSupplyRequestLocationDTO

`func NewSupplyRequestLocationDTO(serviceId int64, name string, address SupplyRequestLocationAddressDTO, type_ SupplyRequestLocationType, ) *SupplyRequestLocationDTO`

NewSupplyRequestLocationDTO instantiates a new SupplyRequestLocationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplyRequestLocationDTOWithDefaults

`func NewSupplyRequestLocationDTOWithDefaults() *SupplyRequestLocationDTO`

NewSupplyRequestLocationDTOWithDefaults instantiates a new SupplyRequestLocationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedDate

`func (o *SupplyRequestLocationDTO) GetRequestedDate() time.Time`

GetRequestedDate returns the RequestedDate field if non-nil, zero value otherwise.

### GetRequestedDateOk

`func (o *SupplyRequestLocationDTO) GetRequestedDateOk() (*time.Time, bool)`

GetRequestedDateOk returns a tuple with the RequestedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDate

`func (o *SupplyRequestLocationDTO) SetRequestedDate(v time.Time)`

SetRequestedDate sets RequestedDate field to given value.

### HasRequestedDate

`func (o *SupplyRequestLocationDTO) HasRequestedDate() bool`

HasRequestedDate returns a boolean if a field has been set.

### GetServiceId

`func (o *SupplyRequestLocationDTO) GetServiceId() int64`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SupplyRequestLocationDTO) GetServiceIdOk() (*int64, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SupplyRequestLocationDTO) SetServiceId(v int64)`

SetServiceId sets ServiceId field to given value.


### GetName

`func (o *SupplyRequestLocationDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SupplyRequestLocationDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SupplyRequestLocationDTO) SetName(v string)`

SetName sets Name field to given value.


### GetAddress

`func (o *SupplyRequestLocationDTO) GetAddress() SupplyRequestLocationAddressDTO`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SupplyRequestLocationDTO) GetAddressOk() (*SupplyRequestLocationAddressDTO, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SupplyRequestLocationDTO) SetAddress(v SupplyRequestLocationAddressDTO)`

SetAddress sets Address field to given value.


### GetType

`func (o *SupplyRequestLocationDTO) GetType() SupplyRequestLocationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SupplyRequestLocationDTO) GetTypeOk() (*SupplyRequestLocationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SupplyRequestLocationDTO) SetType(v SupplyRequestLocationType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


