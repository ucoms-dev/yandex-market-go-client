# BusinessOrderDeliveryAddressDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | Страна. | [optional] 
**Postcode** | Pointer to **string** | Почтовый индекс. | [optional] 
**City** | Pointer to **string** | Город или населенный пункт. | [optional] 
**District** | Pointer to **string** | Район. | [optional] 
**Subway** | Pointer to **string** | Станция метро. | [optional] 
**Street** | Pointer to **string** | Улица. | [optional] 
**House** | Pointer to **string** | Номер дома. | [optional] 
**Block** | Pointer to **string** | Корпус. | [optional] 
**Entrance** | Pointer to **string** | Номер подъезда. | [optional] 
**Entryphone** | Pointer to **string** | Код домофона. | [optional] 
**Floor** | Pointer to **string** | Этаж. | [optional] 
**Apartment** | Pointer to **string** | Номер квартиры или офиса. | [optional] 
**Gps** | Pointer to [**GpsDTO**](GpsDTO.md) |  | [optional] 

## Methods

### NewBusinessOrderDeliveryAddressDTO

`func NewBusinessOrderDeliveryAddressDTO() *BusinessOrderDeliveryAddressDTO`

NewBusinessOrderDeliveryAddressDTO instantiates a new BusinessOrderDeliveryAddressDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessOrderDeliveryAddressDTOWithDefaults

`func NewBusinessOrderDeliveryAddressDTOWithDefaults() *BusinessOrderDeliveryAddressDTO`

NewBusinessOrderDeliveryAddressDTOWithDefaults instantiates a new BusinessOrderDeliveryAddressDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *BusinessOrderDeliveryAddressDTO) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BusinessOrderDeliveryAddressDTO) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BusinessOrderDeliveryAddressDTO) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *BusinessOrderDeliveryAddressDTO) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPostcode

`func (o *BusinessOrderDeliveryAddressDTO) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *BusinessOrderDeliveryAddressDTO) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *BusinessOrderDeliveryAddressDTO) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.

### HasPostcode

`func (o *BusinessOrderDeliveryAddressDTO) HasPostcode() bool`

HasPostcode returns a boolean if a field has been set.

### GetCity

`func (o *BusinessOrderDeliveryAddressDTO) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BusinessOrderDeliveryAddressDTO) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BusinessOrderDeliveryAddressDTO) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *BusinessOrderDeliveryAddressDTO) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetDistrict

`func (o *BusinessOrderDeliveryAddressDTO) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *BusinessOrderDeliveryAddressDTO) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *BusinessOrderDeliveryAddressDTO) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *BusinessOrderDeliveryAddressDTO) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### GetSubway

`func (o *BusinessOrderDeliveryAddressDTO) GetSubway() string`

GetSubway returns the Subway field if non-nil, zero value otherwise.

### GetSubwayOk

`func (o *BusinessOrderDeliveryAddressDTO) GetSubwayOk() (*string, bool)`

GetSubwayOk returns a tuple with the Subway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubway

`func (o *BusinessOrderDeliveryAddressDTO) SetSubway(v string)`

SetSubway sets Subway field to given value.

### HasSubway

`func (o *BusinessOrderDeliveryAddressDTO) HasSubway() bool`

HasSubway returns a boolean if a field has been set.

### GetStreet

`func (o *BusinessOrderDeliveryAddressDTO) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *BusinessOrderDeliveryAddressDTO) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *BusinessOrderDeliveryAddressDTO) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *BusinessOrderDeliveryAddressDTO) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetHouse

`func (o *BusinessOrderDeliveryAddressDTO) GetHouse() string`

GetHouse returns the House field if non-nil, zero value otherwise.

### GetHouseOk

`func (o *BusinessOrderDeliveryAddressDTO) GetHouseOk() (*string, bool)`

GetHouseOk returns a tuple with the House field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouse

`func (o *BusinessOrderDeliveryAddressDTO) SetHouse(v string)`

SetHouse sets House field to given value.

### HasHouse

`func (o *BusinessOrderDeliveryAddressDTO) HasHouse() bool`

HasHouse returns a boolean if a field has been set.

### GetBlock

`func (o *BusinessOrderDeliveryAddressDTO) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *BusinessOrderDeliveryAddressDTO) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *BusinessOrderDeliveryAddressDTO) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *BusinessOrderDeliveryAddressDTO) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetEntrance

`func (o *BusinessOrderDeliveryAddressDTO) GetEntrance() string`

GetEntrance returns the Entrance field if non-nil, zero value otherwise.

### GetEntranceOk

`func (o *BusinessOrderDeliveryAddressDTO) GetEntranceOk() (*string, bool)`

GetEntranceOk returns a tuple with the Entrance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrance

`func (o *BusinessOrderDeliveryAddressDTO) SetEntrance(v string)`

SetEntrance sets Entrance field to given value.

### HasEntrance

`func (o *BusinessOrderDeliveryAddressDTO) HasEntrance() bool`

HasEntrance returns a boolean if a field has been set.

### GetEntryphone

`func (o *BusinessOrderDeliveryAddressDTO) GetEntryphone() string`

GetEntryphone returns the Entryphone field if non-nil, zero value otherwise.

### GetEntryphoneOk

`func (o *BusinessOrderDeliveryAddressDTO) GetEntryphoneOk() (*string, bool)`

GetEntryphoneOk returns a tuple with the Entryphone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryphone

`func (o *BusinessOrderDeliveryAddressDTO) SetEntryphone(v string)`

SetEntryphone sets Entryphone field to given value.

### HasEntryphone

`func (o *BusinessOrderDeliveryAddressDTO) HasEntryphone() bool`

HasEntryphone returns a boolean if a field has been set.

### GetFloor

`func (o *BusinessOrderDeliveryAddressDTO) GetFloor() string`

GetFloor returns the Floor field if non-nil, zero value otherwise.

### GetFloorOk

`func (o *BusinessOrderDeliveryAddressDTO) GetFloorOk() (*string, bool)`

GetFloorOk returns a tuple with the Floor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloor

`func (o *BusinessOrderDeliveryAddressDTO) SetFloor(v string)`

SetFloor sets Floor field to given value.

### HasFloor

`func (o *BusinessOrderDeliveryAddressDTO) HasFloor() bool`

HasFloor returns a boolean if a field has been set.

### GetApartment

`func (o *BusinessOrderDeliveryAddressDTO) GetApartment() string`

GetApartment returns the Apartment field if non-nil, zero value otherwise.

### GetApartmentOk

`func (o *BusinessOrderDeliveryAddressDTO) GetApartmentOk() (*string, bool)`

GetApartmentOk returns a tuple with the Apartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApartment

`func (o *BusinessOrderDeliveryAddressDTO) SetApartment(v string)`

SetApartment sets Apartment field to given value.

### HasApartment

`func (o *BusinessOrderDeliveryAddressDTO) HasApartment() bool`

HasApartment returns a boolean if a field has been set.

### GetGps

`func (o *BusinessOrderDeliveryAddressDTO) GetGps() GpsDTO`

GetGps returns the Gps field if non-nil, zero value otherwise.

### GetGpsOk

`func (o *BusinessOrderDeliveryAddressDTO) GetGpsOk() (*GpsDTO, bool)`

GetGpsOk returns a tuple with the Gps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGps

`func (o *BusinessOrderDeliveryAddressDTO) SetGps(v GpsDTO)`

SetGps sets Gps field to given value.

### HasGps

`func (o *BusinessOrderDeliveryAddressDTO) HasGps() bool`

HasGps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


