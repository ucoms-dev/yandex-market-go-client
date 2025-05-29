# OrderDigitalItemDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Идентификатор товара в заказе.  Он приходит в ответе на запрос [GET campaigns/{campaignId}/orders/{orderId}](../../reference/orders/getOrder.md) — параметр &#x60;id&#x60; в &#x60;items&#x60;.  | 
**Code** | Pointer to **string** | {% note warning \&quot;Вместо него используйте &#x60;codes&#x60;\&quot; %}  Совместное использование обоих параметров приведет к ошибке.  {% endnote %}  Сам ключ.  | [optional] 
**Codes** | Pointer to **[]string** | Ключи, относящиеся к товару. | [optional] 
**Slip** | **string** | Инструкция по активации. | 
**ActivateTill** | **string** | Дата, до которой нужно активировать ключи. Если ключи действуют бессрочно, укажите любую дату в отдаленном будущем.  Формат даты: &#x60;ГГГГ-ММ-ДД&#x60;.  | 

## Methods

### NewOrderDigitalItemDTO

`func NewOrderDigitalItemDTO(id int64, slip string, activateTill string, ) *OrderDigitalItemDTO`

NewOrderDigitalItemDTO instantiates a new OrderDigitalItemDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDigitalItemDTOWithDefaults

`func NewOrderDigitalItemDTOWithDefaults() *OrderDigitalItemDTO`

NewOrderDigitalItemDTOWithDefaults instantiates a new OrderDigitalItemDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderDigitalItemDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderDigitalItemDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderDigitalItemDTO) SetId(v int64)`

SetId sets Id field to given value.


### GetCode

`func (o *OrderDigitalItemDTO) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OrderDigitalItemDTO) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OrderDigitalItemDTO) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *OrderDigitalItemDTO) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCodes

`func (o *OrderDigitalItemDTO) GetCodes() []string`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *OrderDigitalItemDTO) GetCodesOk() (*[]string, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *OrderDigitalItemDTO) SetCodes(v []string)`

SetCodes sets Codes field to given value.

### HasCodes

`func (o *OrderDigitalItemDTO) HasCodes() bool`

HasCodes returns a boolean if a field has been set.

### SetCodesNil

`func (o *OrderDigitalItemDTO) SetCodesNil(b bool)`

 SetCodesNil sets the value for Codes to be an explicit nil

### UnsetCodes
`func (o *OrderDigitalItemDTO) UnsetCodes()`

UnsetCodes ensures that no value is present for Codes, not even an explicit nil
### GetSlip

`func (o *OrderDigitalItemDTO) GetSlip() string`

GetSlip returns the Slip field if non-nil, zero value otherwise.

### GetSlipOk

`func (o *OrderDigitalItemDTO) GetSlipOk() (*string, bool)`

GetSlipOk returns a tuple with the Slip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlip

`func (o *OrderDigitalItemDTO) SetSlip(v string)`

SetSlip sets Slip field to given value.


### GetActivateTill

`func (o *OrderDigitalItemDTO) GetActivateTill() string`

GetActivateTill returns the ActivateTill field if non-nil, zero value otherwise.

### GetActivateTillOk

`func (o *OrderDigitalItemDTO) GetActivateTillOk() (*string, bool)`

GetActivateTillOk returns a tuple with the ActivateTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateTill

`func (o *OrderDigitalItemDTO) SetActivateTill(v string)`

SetActivateTill sets ActivateTill field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


