# ItemPriceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payment** | Pointer to [**CurrencyValueDTO**](CurrencyValueDTO.md) |  | [optional] 
**Subsidy** | Pointer to [**CurrencyValueDTO**](CurrencyValueDTO.md) |  | [optional] 
**Cashback** | Pointer to [**CurrencyValueDTO**](CurrencyValueDTO.md) |  | [optional] 
**Vat** | Pointer to [**OrderVatType**](OrderVatType.md) |  | [optional] 

## Methods

### NewItemPriceDTO

`func NewItemPriceDTO() *ItemPriceDTO`

NewItemPriceDTO instantiates a new ItemPriceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemPriceDTOWithDefaults

`func NewItemPriceDTOWithDefaults() *ItemPriceDTO`

NewItemPriceDTOWithDefaults instantiates a new ItemPriceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayment

`func (o *ItemPriceDTO) GetPayment() CurrencyValueDTO`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *ItemPriceDTO) GetPaymentOk() (*CurrencyValueDTO, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *ItemPriceDTO) SetPayment(v CurrencyValueDTO)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *ItemPriceDTO) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetSubsidy

`func (o *ItemPriceDTO) GetSubsidy() CurrencyValueDTO`

GetSubsidy returns the Subsidy field if non-nil, zero value otherwise.

### GetSubsidyOk

`func (o *ItemPriceDTO) GetSubsidyOk() (*CurrencyValueDTO, bool)`

GetSubsidyOk returns a tuple with the Subsidy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsidy

`func (o *ItemPriceDTO) SetSubsidy(v CurrencyValueDTO)`

SetSubsidy sets Subsidy field to given value.

### HasSubsidy

`func (o *ItemPriceDTO) HasSubsidy() bool`

HasSubsidy returns a boolean if a field has been set.

### GetCashback

`func (o *ItemPriceDTO) GetCashback() CurrencyValueDTO`

GetCashback returns the Cashback field if non-nil, zero value otherwise.

### GetCashbackOk

`func (o *ItemPriceDTO) GetCashbackOk() (*CurrencyValueDTO, bool)`

GetCashbackOk returns a tuple with the Cashback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashback

`func (o *ItemPriceDTO) SetCashback(v CurrencyValueDTO)`

SetCashback sets Cashback field to given value.

### HasCashback

`func (o *ItemPriceDTO) HasCashback() bool`

HasCashback returns a boolean if a field has been set.

### GetVat

`func (o *ItemPriceDTO) GetVat() OrderVatType`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *ItemPriceDTO) GetVatOk() (*OrderVatType, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *ItemPriceDTO) SetVat(v OrderVatType)`

SetVat sets Vat field to given value.

### HasVat

`func (o *ItemPriceDTO) HasVat() bool`

HasVat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


