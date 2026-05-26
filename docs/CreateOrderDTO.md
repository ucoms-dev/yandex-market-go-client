# CreateOrderDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalOrderId** | **string** | Внешний идентификатор заказа в системе магазина. | 
**ItemsDelivery** | [**[]CreateOrderWarehouseItemsDTO**](CreateOrderWarehouseItemsDTO.md) | Список товаров в заказе. | 
**Destination** | [**CreateOrderDeliveryOptionDTO**](CreateOrderDeliveryOptionDTO.md) |  | 
**Customer** | [**CustomerDTO**](CustomerDTO.md) |  | 
**Packaging** | [**CreateOrderPackagingDTO**](CreateOrderPackagingDTO.md) |  | 
**PaymentType** | [**DeliveryPaymentType**](DeliveryPaymentType.md) |  | 
**Draft** | Pointer to **bool** | Признак создания черновика заказа.  * &#x60;true&#x60; — Маркет создаст заказ в статусе &#x60;RESERVED&#x60; и будет ждать подтверждения от магазина. * &#x60;false&#x60; — Маркет создаст заказ в статусе &#x60;PROCESSING&#x60; с подстатусом &#x60;STARTED&#x60; и начнёт его обработку, дополнительных подтверждений не требуется.  | [optional] [default to false]

## Methods

### NewCreateOrderDTO

`func NewCreateOrderDTO(externalOrderId string, itemsDelivery []CreateOrderWarehouseItemsDTO, destination CreateOrderDeliveryOptionDTO, customer CustomerDTO, packaging CreateOrderPackagingDTO, paymentType DeliveryPaymentType, ) *CreateOrderDTO`

NewCreateOrderDTO instantiates a new CreateOrderDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderDTOWithDefaults

`func NewCreateOrderDTOWithDefaults() *CreateOrderDTO`

NewCreateOrderDTOWithDefaults instantiates a new CreateOrderDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalOrderId

`func (o *CreateOrderDTO) GetExternalOrderId() string`

GetExternalOrderId returns the ExternalOrderId field if non-nil, zero value otherwise.

### GetExternalOrderIdOk

`func (o *CreateOrderDTO) GetExternalOrderIdOk() (*string, bool)`

GetExternalOrderIdOk returns a tuple with the ExternalOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOrderId

`func (o *CreateOrderDTO) SetExternalOrderId(v string)`

SetExternalOrderId sets ExternalOrderId field to given value.


### GetItemsDelivery

`func (o *CreateOrderDTO) GetItemsDelivery() []CreateOrderWarehouseItemsDTO`

GetItemsDelivery returns the ItemsDelivery field if non-nil, zero value otherwise.

### GetItemsDeliveryOk

`func (o *CreateOrderDTO) GetItemsDeliveryOk() (*[]CreateOrderWarehouseItemsDTO, bool)`

GetItemsDeliveryOk returns a tuple with the ItemsDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsDelivery

`func (o *CreateOrderDTO) SetItemsDelivery(v []CreateOrderWarehouseItemsDTO)`

SetItemsDelivery sets ItemsDelivery field to given value.


### GetDestination

`func (o *CreateOrderDTO) GetDestination() CreateOrderDeliveryOptionDTO`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *CreateOrderDTO) GetDestinationOk() (*CreateOrderDeliveryOptionDTO, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *CreateOrderDTO) SetDestination(v CreateOrderDeliveryOptionDTO)`

SetDestination sets Destination field to given value.


### GetCustomer

`func (o *CreateOrderDTO) GetCustomer() CustomerDTO`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CreateOrderDTO) GetCustomerOk() (*CustomerDTO, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CreateOrderDTO) SetCustomer(v CustomerDTO)`

SetCustomer sets Customer field to given value.


### GetPackaging

`func (o *CreateOrderDTO) GetPackaging() CreateOrderPackagingDTO`

GetPackaging returns the Packaging field if non-nil, zero value otherwise.

### GetPackagingOk

`func (o *CreateOrderDTO) GetPackagingOk() (*CreateOrderPackagingDTO, bool)`

GetPackagingOk returns a tuple with the Packaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackaging

`func (o *CreateOrderDTO) SetPackaging(v CreateOrderPackagingDTO)`

SetPackaging sets Packaging field to given value.


### GetPaymentType

`func (o *CreateOrderDTO) GetPaymentType() DeliveryPaymentType`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *CreateOrderDTO) GetPaymentTypeOk() (*DeliveryPaymentType, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *CreateOrderDTO) SetPaymentType(v DeliveryPaymentType)`

SetPaymentType sets PaymentType field to given value.


### GetDraft

`func (o *CreateOrderDTO) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *CreateOrderDTO) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *CreateOrderDTO) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *CreateOrderDTO) HasDraft() bool`

HasDraft returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


