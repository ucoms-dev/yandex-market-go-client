# OrderBuyerInfoDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Идентификатор покупателя. | [optional] 
**LastName** | Pointer to **string** | Фамилия покупателя. | [optional] 
**FirstName** | Pointer to **string** | Имя покупателя. | [optional]
**Email** | Pointer to **string** | Электронная почта покупателя. | [optional]
**MiddleName** | Pointer to **string** | Отчество покупателя. | [optional]
**Type** | [**OrderBuyerType**](OrderBuyerType.md) |  |
**Phone** | Pointer to **string** | Подменный номер телефона покупателя. Подробнее о таких номерах читайте [в Справке Маркета для продавцов](https://yandex.ru/support2/marketplace/ru/orders/dbs/call#fake-number).  Формат номера: &#x60;+&lt;код_страны&gt;&lt;код_региона&gt;&lt;номер_телефона&gt;&#x60;.  | [optional] 
**Trusted** | Pointer to **bool** | Проверенный покупатель.  Если параметр &#x60;trusted&#x60; вернулся со значением &#x60;true&#x60;, Маркет уже проверил покупателя — не звоните ему. Обработайте заказ как обычно и передайте его курьеру или отвезите в ПВЗ.  При необходимости свяжитесь с покупателем в чате. [Как это сделать](../../step-by-step/chats.md)  Подробнее о звонках покупателю читайте [в Справке Маркета для продавцов](https://yandex.ru/support/marketplace/ru/orders/dbs/call).  | [optional] 

## Methods

### NewOrderBuyerInfoDTO

`func NewOrderBuyerInfoDTO(type_ OrderBuyerType, ) *OrderBuyerInfoDTO`

NewOrderBuyerInfoDTO instantiates a new OrderBuyerInfoDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderBuyerInfoDTOWithDefaults

`func NewOrderBuyerInfoDTOWithDefaults() *OrderBuyerInfoDTO`

NewOrderBuyerInfoDTOWithDefaults instantiates a new OrderBuyerInfoDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderBuyerInfoDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderBuyerInfoDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderBuyerInfoDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderBuyerInfoDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastName

`func (o *OrderBuyerInfoDTO) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *OrderBuyerInfoDTO) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *OrderBuyerInfoDTO) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *OrderBuyerInfoDTO) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetFirstName

`func (o *OrderBuyerInfoDTO) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *OrderBuyerInfoDTO) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *OrderBuyerInfoDTO) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *OrderBuyerInfoDTO) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetEmail

`func (o *OrderBuyerInfoDTO) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrderBuyerInfoDTO) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrderBuyerInfoDTO) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrderBuyerInfoDTO) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMiddleName

`func (o *OrderBuyerInfoDTO) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *OrderBuyerInfoDTO) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *OrderBuyerInfoDTO) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *OrderBuyerInfoDTO) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetType

`func (o *OrderBuyerInfoDTO) GetType() OrderBuyerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderBuyerInfoDTO) GetTypeOk() (*OrderBuyerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderBuyerInfoDTO) SetType(v OrderBuyerType)`

SetType sets Type field to given value.


### GetPhone

`func (o *OrderBuyerInfoDTO) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *OrderBuyerInfoDTO) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *OrderBuyerInfoDTO) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *OrderBuyerInfoDTO) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetTrusted

`func (o *OrderBuyerInfoDTO) GetTrusted() bool`

GetTrusted returns the Trusted field if non-nil, zero value otherwise.

### GetTrustedOk

`func (o *OrderBuyerInfoDTO) GetTrustedOk() (*bool, bool)`

GetTrustedOk returns a tuple with the Trusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrusted

`func (o *OrderBuyerInfoDTO) SetTrusted(v bool)`

SetTrusted sets Trusted field to given value.

### HasTrusted

`func (o *OrderBuyerInfoDTO) HasTrusted() bool`

HasTrusted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


