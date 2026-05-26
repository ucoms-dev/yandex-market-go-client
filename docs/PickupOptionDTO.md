# PickupOptionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryDateInterval** | [**DeliveryDateIntervalDTO**](DeliveryDateIntervalDTO.md) |  | 
**Price** | [**DeliveryOptionPriceDTO**](DeliveryOptionPriceDTO.md) |  | 

## Methods

### NewPickupOptionDTO

`func NewPickupOptionDTO(deliveryDateInterval DeliveryDateIntervalDTO, price DeliveryOptionPriceDTO, ) *PickupOptionDTO`

NewPickupOptionDTO instantiates a new PickupOptionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPickupOptionDTOWithDefaults

`func NewPickupOptionDTOWithDefaults() *PickupOptionDTO`

NewPickupOptionDTOWithDefaults instantiates a new PickupOptionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryDateInterval

`func (o *PickupOptionDTO) GetDeliveryDateInterval() DeliveryDateIntervalDTO`

GetDeliveryDateInterval returns the DeliveryDateInterval field if non-nil, zero value otherwise.

### GetDeliveryDateIntervalOk

`func (o *PickupOptionDTO) GetDeliveryDateIntervalOk() (*DeliveryDateIntervalDTO, bool)`

GetDeliveryDateIntervalOk returns a tuple with the DeliveryDateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDateInterval

`func (o *PickupOptionDTO) SetDeliveryDateInterval(v DeliveryDateIntervalDTO)`

SetDeliveryDateInterval sets DeliveryDateInterval field to given value.


### GetPrice

`func (o *PickupOptionDTO) GetPrice() DeliveryOptionPriceDTO`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PickupOptionDTO) GetPriceOk() (*DeliveryOptionPriceDTO, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PickupOptionDTO) SetPrice(v DeliveryOptionPriceDTO)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


