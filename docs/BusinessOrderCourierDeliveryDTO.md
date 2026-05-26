# BusinessOrderCourierDeliveryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**BusinessOrderDeliveryAddressDTO**](BusinessOrderDeliveryAddressDTO.md) |  | [optional] 
**Region** | Pointer to [**RegionDTO**](RegionDTO.md) |  | [optional] 

## Methods

### NewBusinessOrderCourierDeliveryDTO

`func NewBusinessOrderCourierDeliveryDTO() *BusinessOrderCourierDeliveryDTO`

NewBusinessOrderCourierDeliveryDTO instantiates a new BusinessOrderCourierDeliveryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessOrderCourierDeliveryDTOWithDefaults

`func NewBusinessOrderCourierDeliveryDTOWithDefaults() *BusinessOrderCourierDeliveryDTO`

NewBusinessOrderCourierDeliveryDTOWithDefaults instantiates a new BusinessOrderCourierDeliveryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *BusinessOrderCourierDeliveryDTO) GetAddress() BusinessOrderDeliveryAddressDTO`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BusinessOrderCourierDeliveryDTO) GetAddressOk() (*BusinessOrderDeliveryAddressDTO, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BusinessOrderCourierDeliveryDTO) SetAddress(v BusinessOrderDeliveryAddressDTO)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BusinessOrderCourierDeliveryDTO) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetRegion

`func (o *BusinessOrderCourierDeliveryDTO) GetRegion() RegionDTO`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BusinessOrderCourierDeliveryDTO) GetRegionOk() (*RegionDTO, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BusinessOrderCourierDeliveryDTO) SetRegion(v RegionDTO)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *BusinessOrderCourierDeliveryDTO) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


