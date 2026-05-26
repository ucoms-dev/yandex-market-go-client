# ChatFullContextDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ChatContextType**](ChatContextType.md) |  | 
**Customer** | Pointer to [**ChatCustomerDTO**](ChatCustomerDTO.md) |  | [optional] 
**CampaignId** | Pointer to **int64** | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями.  | [optional] 
**OrderId** | Pointer to **int64** | Идентификатор заказа.  Возвращается для заказов и возвратов.  | [optional] 
**ReturnId** | Pointer to **int64** | Идентификатор возврата.  Возвращается только для возвратов.  | [optional] 

## Methods

### NewChatFullContextDTO

`func NewChatFullContextDTO(type_ ChatContextType, ) *ChatFullContextDTO`

NewChatFullContextDTO instantiates a new ChatFullContextDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatFullContextDTOWithDefaults

`func NewChatFullContextDTOWithDefaults() *ChatFullContextDTO`

NewChatFullContextDTOWithDefaults instantiates a new ChatFullContextDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChatFullContextDTO) GetType() ChatContextType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChatFullContextDTO) GetTypeOk() (*ChatContextType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChatFullContextDTO) SetType(v ChatContextType)`

SetType sets Type field to given value.


### GetCustomer

`func (o *ChatFullContextDTO) GetCustomer() ChatCustomerDTO`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *ChatFullContextDTO) GetCustomerOk() (*ChatCustomerDTO, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *ChatFullContextDTO) SetCustomer(v ChatCustomerDTO)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *ChatFullContextDTO) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCampaignId

`func (o *ChatFullContextDTO) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *ChatFullContextDTO) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *ChatFullContextDTO) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *ChatFullContextDTO) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetOrderId

`func (o *ChatFullContextDTO) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ChatFullContextDTO) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ChatFullContextDTO) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ChatFullContextDTO) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetReturnId

`func (o *ChatFullContextDTO) GetReturnId() int64`

GetReturnId returns the ReturnId field if non-nil, zero value otherwise.

### GetReturnIdOk

`func (o *ChatFullContextDTO) GetReturnIdOk() (*int64, bool)`

GetReturnIdOk returns a tuple with the ReturnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnId

`func (o *ChatFullContextDTO) SetReturnId(v int64)`

SetReturnId sets ReturnId field to given value.

### HasReturnId

`func (o *ChatFullContextDTO) HasReturnId() bool`

HasReturnId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


