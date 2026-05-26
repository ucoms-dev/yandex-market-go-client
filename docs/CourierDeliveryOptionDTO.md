# CourierDeliveryOptionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryDateInterval** | [**DeliveryDateIntervalDTO**](DeliveryDateIntervalDTO.md) |  | 
**DeliveryTimeInterval** | [**TimeIntervalDTO**](TimeIntervalDTO.md) |  | 
**Price** | [**DeliveryOptionPriceDTO**](DeliveryOptionPriceDTO.md) |  | 

## Methods

### NewCourierDeliveryOptionDTO

`func NewCourierDeliveryOptionDTO(deliveryDateInterval DeliveryDateIntervalDTO, deliveryTimeInterval TimeIntervalDTO, price DeliveryOptionPriceDTO, ) *CourierDeliveryOptionDTO`

NewCourierDeliveryOptionDTO instantiates a new CourierDeliveryOptionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourierDeliveryOptionDTOWithDefaults

`func NewCourierDeliveryOptionDTOWithDefaults() *CourierDeliveryOptionDTO`

NewCourierDeliveryOptionDTOWithDefaults instantiates a new CourierDeliveryOptionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryDateInterval

`func (o *CourierDeliveryOptionDTO) GetDeliveryDateInterval() DeliveryDateIntervalDTO`

GetDeliveryDateInterval returns the DeliveryDateInterval field if non-nil, zero value otherwise.

### GetDeliveryDateIntervalOk

`func (o *CourierDeliveryOptionDTO) GetDeliveryDateIntervalOk() (*DeliveryDateIntervalDTO, bool)`

GetDeliveryDateIntervalOk returns a tuple with the DeliveryDateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDateInterval

`func (o *CourierDeliveryOptionDTO) SetDeliveryDateInterval(v DeliveryDateIntervalDTO)`

SetDeliveryDateInterval sets DeliveryDateInterval field to given value.


### GetDeliveryTimeInterval

`func (o *CourierDeliveryOptionDTO) GetDeliveryTimeInterval() TimeIntervalDTO`

GetDeliveryTimeInterval returns the DeliveryTimeInterval field if non-nil, zero value otherwise.

### GetDeliveryTimeIntervalOk

`func (o *CourierDeliveryOptionDTO) GetDeliveryTimeIntervalOk() (*TimeIntervalDTO, bool)`

GetDeliveryTimeIntervalOk returns a tuple with the DeliveryTimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTimeInterval

`func (o *CourierDeliveryOptionDTO) SetDeliveryTimeInterval(v TimeIntervalDTO)`

SetDeliveryTimeInterval sets DeliveryTimeInterval field to given value.


### GetPrice

`func (o *CourierDeliveryOptionDTO) GetPrice() DeliveryOptionPriceDTO`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CourierDeliveryOptionDTO) GetPriceOk() (*DeliveryOptionPriceDTO, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CourierDeliveryOptionDTO) SetPrice(v DeliveryOptionPriceDTO)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


