# ChatCustomerDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Публичное имя покупателя в Яндекс Паспорте, которое отображается в сервисах Яндекса.  | [optional] 
**PublicId** | Pointer to **string** | Публичный идентификатор пользователя в Яндекс Паспорте.  {% cut \&quot;Примеры, где используется\&quot; %}  * Маркет: &#x60;https://market.yandex.ru/user/{public-id}/reviews&#x60; * Дзен: &#x60;https://zen.yandex.ru/user/{public-id}&#x60; * Отзывы: &#x60;https://yandex.ru/user/{public-id}&#x60;  {% endcut %}  Подробнее о публичных данных читайте в [документации Яндекс ID](https://yandex.ru/support/id/ru/data/public-data).  | [optional] 

## Methods

### NewChatCustomerDTO

`func NewChatCustomerDTO() *ChatCustomerDTO`

NewChatCustomerDTO instantiates a new ChatCustomerDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCustomerDTOWithDefaults

`func NewChatCustomerDTOWithDefaults() *ChatCustomerDTO`

NewChatCustomerDTOWithDefaults instantiates a new ChatCustomerDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ChatCustomerDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChatCustomerDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChatCustomerDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChatCustomerDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicId

`func (o *ChatCustomerDTO) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *ChatCustomerDTO) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *ChatCustomerDTO) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *ChatCustomerDTO) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


