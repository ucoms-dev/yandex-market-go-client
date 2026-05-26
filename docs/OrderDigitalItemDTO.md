# OrderDigitalItemDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Идентификатор товара в заказе.  Он приходит в ответе метода [POST v1/businesses/{businessId}/orders](../../reference/orders/getBusinessOrders.md) — параметр &#x60;id&#x60; в &#x60;items&#x60;.  | 
**Codes** | Pointer to **[]string** | Ключи, относящиеся к товару.  Поле обязательно для заполнения.  | [optional] 
**Slip** | **string** | Инструкция по активации.  Для форматирования текста можно использовать теги HTML:  * \\&lt;h&gt;, \\&lt;h1&gt;, \\&lt;h2&gt; и так далее — для заголовков; * \\&lt;br&gt; и \\&lt;p&gt; — для переноса строки; * \\&lt;ol&gt; — для нумерованного списка; * \\&lt;ul&gt; — для маркированного списка; * \\&lt;li&gt; — для создания элементов списка (должен находиться внутри \\&lt;ol&gt; или \\&lt;ul&gt;); * \\&lt;div&gt; — поддерживается, но не влияет на отображение текста.  | 
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


