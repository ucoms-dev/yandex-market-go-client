# SupplyRequestLocationAddressDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullAddress** | **string** | Полный адрес склада или ПВЗ. | 
**Gps** | [**GpsDTO**](GpsDTO.md) |  | 

## Methods

### NewSupplyRequestLocationAddressDTO

`func NewSupplyRequestLocationAddressDTO(fullAddress string, gps GpsDTO, ) *SupplyRequestLocationAddressDTO`

NewSupplyRequestLocationAddressDTO instantiates a new SupplyRequestLocationAddressDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplyRequestLocationAddressDTOWithDefaults

`func NewSupplyRequestLocationAddressDTOWithDefaults() *SupplyRequestLocationAddressDTO`

NewSupplyRequestLocationAddressDTOWithDefaults instantiates a new SupplyRequestLocationAddressDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullAddress

`func (o *SupplyRequestLocationAddressDTO) GetFullAddress() string`

GetFullAddress returns the FullAddress field if non-nil, zero value otherwise.

### GetFullAddressOk

`func (o *SupplyRequestLocationAddressDTO) GetFullAddressOk() (*string, bool)`

GetFullAddressOk returns a tuple with the FullAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullAddress

`func (o *SupplyRequestLocationAddressDTO) SetFullAddress(v string)`

SetFullAddress sets FullAddress field to given value.


### GetGps

`func (o *SupplyRequestLocationAddressDTO) GetGps() GpsDTO`

GetGps returns the Gps field if non-nil, zero value otherwise.

### GetGpsOk

`func (o *SupplyRequestLocationAddressDTO) GetGpsOk() (*GpsDTO, bool)`

GetGpsOk returns a tuple with the Gps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGps

`func (o *SupplyRequestLocationAddressDTO) SetGps(v GpsDTO)`

SetGps sets Gps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


