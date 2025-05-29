# UpdateBusinessPricesDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **float32** | Значение. | 
**CurrencyId** | [**CurrencyType**](CurrencyType.md) |  | 
**DiscountBase** | Pointer to **float32** | Зачеркнутая цена.  Число должно быть целым. Вы можете указать цену со скидкой от 5 до 99%.  Передавайте этот параметр при каждом обновлении цены, если предоставляете скидку на товар.  | [optional] 

## Methods

### NewUpdateBusinessPricesDTO

`func NewUpdateBusinessPricesDTO(value float32, currencyId CurrencyType, ) *UpdateBusinessPricesDTO`

NewUpdateBusinessPricesDTO instantiates a new UpdateBusinessPricesDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBusinessPricesDTOWithDefaults

`func NewUpdateBusinessPricesDTOWithDefaults() *UpdateBusinessPricesDTO`

NewUpdateBusinessPricesDTOWithDefaults instantiates a new UpdateBusinessPricesDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *UpdateBusinessPricesDTO) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateBusinessPricesDTO) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateBusinessPricesDTO) SetValue(v float32)`

SetValue sets Value field to given value.


### GetCurrencyId

`func (o *UpdateBusinessPricesDTO) GetCurrencyId() CurrencyType`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *UpdateBusinessPricesDTO) GetCurrencyIdOk() (*CurrencyType, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *UpdateBusinessPricesDTO) SetCurrencyId(v CurrencyType)`

SetCurrencyId sets CurrencyId field to given value.


### GetDiscountBase

`func (o *UpdateBusinessPricesDTO) GetDiscountBase() float32`

GetDiscountBase returns the DiscountBase field if non-nil, zero value otherwise.

### GetDiscountBaseOk

`func (o *UpdateBusinessPricesDTO) GetDiscountBaseOk() (*float32, bool)`

GetDiscountBaseOk returns a tuple with the DiscountBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountBase

`func (o *UpdateBusinessPricesDTO) SetDiscountBase(v float32)`

SetDiscountBase sets DiscountBase field to given value.

### HasDiscountBase

`func (o *UpdateBusinessPricesDTO) HasDiscountBase() bool`

HasDiscountBase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


