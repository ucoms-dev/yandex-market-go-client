# CourierDeliveryAddressDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullAddress** | **string** | Полный адрес с точностью до номера дома. | 
**Entrance** | Pointer to **string** | Номер подъезда. | [optional] 
**Floor** | Pointer to **int32** | Этаж. | [optional] 
**Apartment** | Pointer to **string** | Номер квартиры или офиса. | [optional] 

## Methods

### NewCourierDeliveryAddressDTO

`func NewCourierDeliveryAddressDTO(fullAddress string, ) *CourierDeliveryAddressDTO`

NewCourierDeliveryAddressDTO instantiates a new CourierDeliveryAddressDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourierDeliveryAddressDTOWithDefaults

`func NewCourierDeliveryAddressDTOWithDefaults() *CourierDeliveryAddressDTO`

NewCourierDeliveryAddressDTOWithDefaults instantiates a new CourierDeliveryAddressDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullAddress

`func (o *CourierDeliveryAddressDTO) GetFullAddress() string`

GetFullAddress returns the FullAddress field if non-nil, zero value otherwise.

### GetFullAddressOk

`func (o *CourierDeliveryAddressDTO) GetFullAddressOk() (*string, bool)`

GetFullAddressOk returns a tuple with the FullAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullAddress

`func (o *CourierDeliveryAddressDTO) SetFullAddress(v string)`

SetFullAddress sets FullAddress field to given value.


### GetEntrance

`func (o *CourierDeliveryAddressDTO) GetEntrance() string`

GetEntrance returns the Entrance field if non-nil, zero value otherwise.

### GetEntranceOk

`func (o *CourierDeliveryAddressDTO) GetEntranceOk() (*string, bool)`

GetEntranceOk returns a tuple with the Entrance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrance

`func (o *CourierDeliveryAddressDTO) SetEntrance(v string)`

SetEntrance sets Entrance field to given value.

### HasEntrance

`func (o *CourierDeliveryAddressDTO) HasEntrance() bool`

HasEntrance returns a boolean if a field has been set.

### GetFloor

`func (o *CourierDeliveryAddressDTO) GetFloor() int32`

GetFloor returns the Floor field if non-nil, zero value otherwise.

### GetFloorOk

`func (o *CourierDeliveryAddressDTO) GetFloorOk() (*int32, bool)`

GetFloorOk returns a tuple with the Floor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloor

`func (o *CourierDeliveryAddressDTO) SetFloor(v int32)`

SetFloor sets Floor field to given value.

### HasFloor

`func (o *CourierDeliveryAddressDTO) HasFloor() bool`

HasFloor returns a boolean if a field has been set.

### GetApartment

`func (o *CourierDeliveryAddressDTO) GetApartment() string`

GetApartment returns the Apartment field if non-nil, zero value otherwise.

### GetApartmentOk

`func (o *CourierDeliveryAddressDTO) GetApartmentOk() (*string, bool)`

GetApartmentOk returns a tuple with the Apartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApartment

`func (o *CourierDeliveryAddressDTO) SetApartment(v string)`

SetApartment sets Apartment field to given value.

### HasApartment

`func (o *CourierDeliveryAddressDTO) HasApartment() bool`

HasApartment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


