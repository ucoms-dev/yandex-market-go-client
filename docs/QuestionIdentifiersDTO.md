# QuestionIdentifiersDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Идентификатор вопроса.  | 
**CategoryId** | Pointer to **int64** | Идентификатор категории. | [optional] 
**OfferId** | **string** | Ваш SKU — идентификатор товара в вашей системе.  Правила использования SKU:  * У каждого товара SKU должен быть свой.  * Уже заданный SKU нельзя освободить и использовать заново для другого товара. Каждый товар должен получать новый идентификатор, до того никогда не использовавшийся в вашем каталоге.  SKU товара можно изменить в кабинете продавца на Маркете. О том, как это сделать, читайте [в Справке Маркета для продавцов](https://yandex.ru/support2/marketplace/ru/assortment/operations/edit-sku).  {% note warning %}  Пробельные символы в начале и конце значения автоматически удаляются. Например, &#x60;\&quot;  SKU123  \&quot;&#x60; и &#x60;\&quot;SKU123\&quot;&#x60; будут обработаны как одинаковые значения.  {% endnote %}  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields)  | 

## Methods

### NewQuestionIdentifiersDTO

`func NewQuestionIdentifiersDTO(id int64, offerId string, ) *QuestionIdentifiersDTO`

NewQuestionIdentifiersDTO instantiates a new QuestionIdentifiersDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionIdentifiersDTOWithDefaults

`func NewQuestionIdentifiersDTOWithDefaults() *QuestionIdentifiersDTO`

NewQuestionIdentifiersDTOWithDefaults instantiates a new QuestionIdentifiersDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionIdentifiersDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionIdentifiersDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionIdentifiersDTO) SetId(v int64)`

SetId sets Id field to given value.


### GetCategoryId

`func (o *QuestionIdentifiersDTO) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *QuestionIdentifiersDTO) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *QuestionIdentifiersDTO) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *QuestionIdentifiersDTO) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetOfferId

`func (o *QuestionIdentifiersDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *QuestionIdentifiersDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *QuestionIdentifiersDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


