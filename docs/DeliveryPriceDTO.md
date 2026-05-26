# DeliveryPriceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payment** | Pointer to [**CurrencyValueDTO**](CurrencyValueDTO.md) |  | [optional] 
**Subsidy** | Pointer to [**CurrencyValueDTO**](CurrencyValueDTO.md) |  | [optional] 
**Vat** | Pointer to [**OrderVatType**](OrderVatType.md) |  | [optional] 

## Methods

### NewDeliveryPriceDTO

`func NewDeliveryPriceDTO() *DeliveryPriceDTO`

NewDeliveryPriceDTO instantiates a new DeliveryPriceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryPriceDTOWithDefaults

`func NewDeliveryPriceDTOWithDefaults() *DeliveryPriceDTO`

NewDeliveryPriceDTOWithDefaults instantiates a new DeliveryPriceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayment

`func (o *DeliveryPriceDTO) GetPayment() CurrencyValueDTO`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *DeliveryPriceDTO) GetPaymentOk() (*CurrencyValueDTO, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *DeliveryPriceDTO) SetPayment(v CurrencyValueDTO)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *DeliveryPriceDTO) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetSubsidy

`func (o *DeliveryPriceDTO) GetSubsidy() CurrencyValueDTO`

GetSubsidy returns the Subsidy field if non-nil, zero value otherwise.

### GetSubsidyOk

`func (o *DeliveryPriceDTO) GetSubsidyOk() (*CurrencyValueDTO, bool)`

GetSubsidyOk returns a tuple with the Subsidy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsidy

`func (o *DeliveryPriceDTO) SetSubsidy(v CurrencyValueDTO)`

SetSubsidy sets Subsidy field to given value.

### HasSubsidy

`func (o *DeliveryPriceDTO) HasSubsidy() bool`

HasSubsidy returns a boolean if a field has been set.

### GetVat

`func (o *DeliveryPriceDTO) GetVat() OrderVatType`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *DeliveryPriceDTO) GetVatOk() (*OrderVatType, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *DeliveryPriceDTO) SetVat(v OrderVatType)`

SetVat sets Vat field to given value.

### HasVat

`func (o *DeliveryPriceDTO) HasVat() bool`

HasVat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


