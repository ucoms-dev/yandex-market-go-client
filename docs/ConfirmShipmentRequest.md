# ConfirmShipmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalShipmentId** | Pointer to **string** | Идентификатор отгрузки в системе поставщика. | [optional] 
**Signatory** | Pointer to **string** | Логин пользователя в Яндекс ID, от имени которого будет подписываться электронный акт приема-передачи.  Указывается без &#x60;@yandex.ru&#x60;.  Где его найти:  * на странице [Яндекс ID](https://id.yandex.ru); * в [кабинете продавца на Маркете](https://partner.market.yandex.ru/):    * слева снизу под иконкой пользователя;   * на странице **Настройки** → **Сотрудники и доступы**.  | [optional] 

## Methods

### NewConfirmShipmentRequest

`func NewConfirmShipmentRequest() *ConfirmShipmentRequest`

NewConfirmShipmentRequest instantiates a new ConfirmShipmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmShipmentRequestWithDefaults

`func NewConfirmShipmentRequestWithDefaults() *ConfirmShipmentRequest`

NewConfirmShipmentRequestWithDefaults instantiates a new ConfirmShipmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalShipmentId

`func (o *ConfirmShipmentRequest) GetExternalShipmentId() string`

GetExternalShipmentId returns the ExternalShipmentId field if non-nil, zero value otherwise.

### GetExternalShipmentIdOk

`func (o *ConfirmShipmentRequest) GetExternalShipmentIdOk() (*string, bool)`

GetExternalShipmentIdOk returns a tuple with the ExternalShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalShipmentId

`func (o *ConfirmShipmentRequest) SetExternalShipmentId(v string)`

SetExternalShipmentId sets ExternalShipmentId field to given value.

### HasExternalShipmentId

`func (o *ConfirmShipmentRequest) HasExternalShipmentId() bool`

HasExternalShipmentId returns a boolean if a field has been set.

### GetSignatory

`func (o *ConfirmShipmentRequest) GetSignatory() string`

GetSignatory returns the Signatory field if non-nil, zero value otherwise.

### GetSignatoryOk

`func (o *ConfirmShipmentRequest) GetSignatoryOk() (*string, bool)`

GetSignatoryOk returns a tuple with the Signatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatory

`func (o *ConfirmShipmentRequest) SetSignatory(v string)`

SetSignatory sets Signatory field to given value.

### HasSignatory

`func (o *ConfirmShipmentRequest) HasSignatory() bool`

HasSignatory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


