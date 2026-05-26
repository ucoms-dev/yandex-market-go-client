# DeliveryOptionPriceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **float32** | Значение. | 
**CurrencyId** | [**CurrencyType**](CurrencyType.md) |  | 

## Methods

### NewDeliveryOptionPriceDTO

`func NewDeliveryOptionPriceDTO(value float32, currencyId CurrencyType, ) *DeliveryOptionPriceDTO`

NewDeliveryOptionPriceDTO instantiates a new DeliveryOptionPriceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryOptionPriceDTOWithDefaults

`func NewDeliveryOptionPriceDTOWithDefaults() *DeliveryOptionPriceDTO`

NewDeliveryOptionPriceDTOWithDefaults instantiates a new DeliveryOptionPriceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *DeliveryOptionPriceDTO) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DeliveryOptionPriceDTO) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DeliveryOptionPriceDTO) SetValue(v float32)`

SetValue sets Value field to given value.


### GetCurrencyId

`func (o *DeliveryOptionPriceDTO) GetCurrencyId() CurrencyType`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *DeliveryOptionPriceDTO) GetCurrencyIdOk() (*CurrencyType, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *DeliveryOptionPriceDTO) SetCurrencyId(v CurrencyType)`

SetCurrencyId sets CurrencyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


