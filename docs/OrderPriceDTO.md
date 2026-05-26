# OrderPriceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payment** | Pointer to [**CurrencyValueDTO**](CurrencyValueDTO.md) |  | [optional] 
**Subsidy** | Pointer to [**CurrencyValueDTO**](CurrencyValueDTO.md) |  | [optional] 
**Cashback** | Pointer to [**CurrencyValueDTO**](CurrencyValueDTO.md) |  | [optional] 
**Delivery** | Pointer to [**DeliveryPriceDTO**](DeliveryPriceDTO.md) |  | [optional] 

## Methods

### NewOrderPriceDTO

`func NewOrderPriceDTO() *OrderPriceDTO`

NewOrderPriceDTO instantiates a new OrderPriceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderPriceDTOWithDefaults

`func NewOrderPriceDTOWithDefaults() *OrderPriceDTO`

NewOrderPriceDTOWithDefaults instantiates a new OrderPriceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayment

`func (o *OrderPriceDTO) GetPayment() CurrencyValueDTO`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *OrderPriceDTO) GetPaymentOk() (*CurrencyValueDTO, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *OrderPriceDTO) SetPayment(v CurrencyValueDTO)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *OrderPriceDTO) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetSubsidy

`func (o *OrderPriceDTO) GetSubsidy() CurrencyValueDTO`

GetSubsidy returns the Subsidy field if non-nil, zero value otherwise.

### GetSubsidyOk

`func (o *OrderPriceDTO) GetSubsidyOk() (*CurrencyValueDTO, bool)`

GetSubsidyOk returns a tuple with the Subsidy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsidy

`func (o *OrderPriceDTO) SetSubsidy(v CurrencyValueDTO)`

SetSubsidy sets Subsidy field to given value.

### HasSubsidy

`func (o *OrderPriceDTO) HasSubsidy() bool`

HasSubsidy returns a boolean if a field has been set.

### GetCashback

`func (o *OrderPriceDTO) GetCashback() CurrencyValueDTO`

GetCashback returns the Cashback field if non-nil, zero value otherwise.

### GetCashbackOk

`func (o *OrderPriceDTO) GetCashbackOk() (*CurrencyValueDTO, bool)`

GetCashbackOk returns a tuple with the Cashback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashback

`func (o *OrderPriceDTO) SetCashback(v CurrencyValueDTO)`

SetCashback sets Cashback field to given value.

### HasCashback

`func (o *OrderPriceDTO) HasCashback() bool`

HasCashback returns a boolean if a field has been set.

### GetDelivery

`func (o *OrderPriceDTO) GetDelivery() DeliveryPriceDTO`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *OrderPriceDTO) GetDeliveryOk() (*DeliveryPriceDTO, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *OrderPriceDTO) SetDelivery(v DeliveryPriceDTO)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *OrderPriceDTO) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


