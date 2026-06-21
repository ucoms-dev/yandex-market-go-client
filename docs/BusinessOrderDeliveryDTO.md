# BusinessOrderDeliveryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**OrderDeliveryType**](OrderDeliveryType.md) |  | 
**ServiceName** | **string** | Название службы доставки. | 
**DeliveryServiceId** | **int64** | Идентификатор службы доставки. | 
**WarehouseId** | Pointer to **string** | Идентификатор склада в системе магазина, на который сформирован заказ. | [optional] 
**DeliveryPartnerType** | [**OrderDeliveryPartnerType**](OrderDeliveryPartnerType.md) |  | 
**DispatchType** | Pointer to [**OrderDeliveryDispatchType**](OrderDeliveryDispatchType.md) |  | [optional] 
**Dates** | [**BusinessOrderDeliveryDatesDTO**](BusinessOrderDeliveryDatesDTO.md) |  | 
**Shipment** | Pointer to [**BusinessOrderShipmentDTO**](BusinessOrderShipmentDTO.md) |  | [optional] 
**Courier** | Pointer to [**BusinessOrderCourierDeliveryDTO**](BusinessOrderCourierDeliveryDTO.md) |  | [optional] 
**Pickup** | Pointer to [**BusinessOrderPickupDeliveryDTO**](BusinessOrderPickupDeliveryDTO.md) |  | [optional] 
**Transfer** | Pointer to [**BusinessOrderTransferDTO**](BusinessOrderTransferDTO.md) |  | [optional] 
**BoxesLayout** | Pointer to [**[]BusinessOrderBoxLayoutDTO**](BusinessOrderBoxLayoutDTO.md) | Раскладка товаров по коробкам. | [optional] 
**Tracks** | Pointer to [**[]OrderTrackDTO**](OrderTrackDTO.md) | Информация для отслеживания посылки. | [optional] 
**Estimated** | Pointer to **bool** | Приблизительная ли дата доставки. | [optional] 
**ReceiveBarcode** | Pointer to **string** | **Только для модели LaaS**  Штрихкод получения заказа на ПВЗ.  | [optional] 
**ReceiveCode** | Pointer to **string** | **Только для модели LaaS**  Код получения заказа на ПВЗ.  | [optional] 
**DigitalGoods** | Pointer to [**DigitalGoodsDeliveryDetailsDTO**](DigitalGoodsDeliveryDetailsDTO.md) |  | [optional] 

## Methods

### NewBusinessOrderDeliveryDTO

`func NewBusinessOrderDeliveryDTO(type_ OrderDeliveryType, serviceName string, deliveryServiceId int64, deliveryPartnerType OrderDeliveryPartnerType, dates BusinessOrderDeliveryDatesDTO, ) *BusinessOrderDeliveryDTO`

NewBusinessOrderDeliveryDTO instantiates a new BusinessOrderDeliveryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessOrderDeliveryDTOWithDefaults

`func NewBusinessOrderDeliveryDTOWithDefaults() *BusinessOrderDeliveryDTO`

NewBusinessOrderDeliveryDTOWithDefaults instantiates a new BusinessOrderDeliveryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BusinessOrderDeliveryDTO) GetType() OrderDeliveryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BusinessOrderDeliveryDTO) GetTypeOk() (*OrderDeliveryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BusinessOrderDeliveryDTO) SetType(v OrderDeliveryType)`

SetType sets Type field to given value.


### GetServiceName

`func (o *BusinessOrderDeliveryDTO) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *BusinessOrderDeliveryDTO) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *BusinessOrderDeliveryDTO) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetDeliveryServiceId

`func (o *BusinessOrderDeliveryDTO) GetDeliveryServiceId() int64`

GetDeliveryServiceId returns the DeliveryServiceId field if non-nil, zero value otherwise.

### GetDeliveryServiceIdOk

`func (o *BusinessOrderDeliveryDTO) GetDeliveryServiceIdOk() (*int64, bool)`

GetDeliveryServiceIdOk returns a tuple with the DeliveryServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryServiceId

`func (o *BusinessOrderDeliveryDTO) SetDeliveryServiceId(v int64)`

SetDeliveryServiceId sets DeliveryServiceId field to given value.


### GetWarehouseId

`func (o *BusinessOrderDeliveryDTO) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *BusinessOrderDeliveryDTO) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *BusinessOrderDeliveryDTO) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *BusinessOrderDeliveryDTO) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetDeliveryPartnerType

`func (o *BusinessOrderDeliveryDTO) GetDeliveryPartnerType() OrderDeliveryPartnerType`

GetDeliveryPartnerType returns the DeliveryPartnerType field if non-nil, zero value otherwise.

### GetDeliveryPartnerTypeOk

`func (o *BusinessOrderDeliveryDTO) GetDeliveryPartnerTypeOk() (*OrderDeliveryPartnerType, bool)`

GetDeliveryPartnerTypeOk returns a tuple with the DeliveryPartnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryPartnerType

`func (o *BusinessOrderDeliveryDTO) SetDeliveryPartnerType(v OrderDeliveryPartnerType)`

SetDeliveryPartnerType sets DeliveryPartnerType field to given value.


### GetDispatchType

`func (o *BusinessOrderDeliveryDTO) GetDispatchType() OrderDeliveryDispatchType`

GetDispatchType returns the DispatchType field if non-nil, zero value otherwise.

### GetDispatchTypeOk

`func (o *BusinessOrderDeliveryDTO) GetDispatchTypeOk() (*OrderDeliveryDispatchType, bool)`

GetDispatchTypeOk returns a tuple with the DispatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchType

`func (o *BusinessOrderDeliveryDTO) SetDispatchType(v OrderDeliveryDispatchType)`

SetDispatchType sets DispatchType field to given value.

### HasDispatchType

`func (o *BusinessOrderDeliveryDTO) HasDispatchType() bool`

HasDispatchType returns a boolean if a field has been set.

### GetDates

`func (o *BusinessOrderDeliveryDTO) GetDates() BusinessOrderDeliveryDatesDTO`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *BusinessOrderDeliveryDTO) GetDatesOk() (*BusinessOrderDeliveryDatesDTO, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *BusinessOrderDeliveryDTO) SetDates(v BusinessOrderDeliveryDatesDTO)`

SetDates sets Dates field to given value.


### GetShipment

`func (o *BusinessOrderDeliveryDTO) GetShipment() BusinessOrderShipmentDTO`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *BusinessOrderDeliveryDTO) GetShipmentOk() (*BusinessOrderShipmentDTO, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *BusinessOrderDeliveryDTO) SetShipment(v BusinessOrderShipmentDTO)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *BusinessOrderDeliveryDTO) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetCourier

`func (o *BusinessOrderDeliveryDTO) GetCourier() BusinessOrderCourierDeliveryDTO`

GetCourier returns the Courier field if non-nil, zero value otherwise.

### GetCourierOk

`func (o *BusinessOrderDeliveryDTO) GetCourierOk() (*BusinessOrderCourierDeliveryDTO, bool)`

GetCourierOk returns a tuple with the Courier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourier

`func (o *BusinessOrderDeliveryDTO) SetCourier(v BusinessOrderCourierDeliveryDTO)`

SetCourier sets Courier field to given value.

### HasCourier

`func (o *BusinessOrderDeliveryDTO) HasCourier() bool`

HasCourier returns a boolean if a field has been set.

### GetPickup

`func (o *BusinessOrderDeliveryDTO) GetPickup() BusinessOrderPickupDeliveryDTO`

GetPickup returns the Pickup field if non-nil, zero value otherwise.

### GetPickupOk

`func (o *BusinessOrderDeliveryDTO) GetPickupOk() (*BusinessOrderPickupDeliveryDTO, bool)`

GetPickupOk returns a tuple with the Pickup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickup

`func (o *BusinessOrderDeliveryDTO) SetPickup(v BusinessOrderPickupDeliveryDTO)`

SetPickup sets Pickup field to given value.

### HasPickup

`func (o *BusinessOrderDeliveryDTO) HasPickup() bool`

HasPickup returns a boolean if a field has been set.

### GetTransfer

`func (o *BusinessOrderDeliveryDTO) GetTransfer() BusinessOrderTransferDTO`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *BusinessOrderDeliveryDTO) GetTransferOk() (*BusinessOrderTransferDTO, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *BusinessOrderDeliveryDTO) SetTransfer(v BusinessOrderTransferDTO)`

SetTransfer sets Transfer field to given value.

### HasTransfer

`func (o *BusinessOrderDeliveryDTO) HasTransfer() bool`

HasTransfer returns a boolean if a field has been set.

### GetBoxesLayout

`func (o *BusinessOrderDeliveryDTO) GetBoxesLayout() []BusinessOrderBoxLayoutDTO`

GetBoxesLayout returns the BoxesLayout field if non-nil, zero value otherwise.

### GetBoxesLayoutOk

`func (o *BusinessOrderDeliveryDTO) GetBoxesLayoutOk() (*[]BusinessOrderBoxLayoutDTO, bool)`

GetBoxesLayoutOk returns a tuple with the BoxesLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxesLayout

`func (o *BusinessOrderDeliveryDTO) SetBoxesLayout(v []BusinessOrderBoxLayoutDTO)`

SetBoxesLayout sets BoxesLayout field to given value.

### HasBoxesLayout

`func (o *BusinessOrderDeliveryDTO) HasBoxesLayout() bool`

HasBoxesLayout returns a boolean if a field has been set.

### SetBoxesLayoutNil

`func (o *BusinessOrderDeliveryDTO) SetBoxesLayoutNil(b bool)`

 SetBoxesLayoutNil sets the value for BoxesLayout to be an explicit nil

### UnsetBoxesLayout
`func (o *BusinessOrderDeliveryDTO) UnsetBoxesLayout()`

UnsetBoxesLayout ensures that no value is present for BoxesLayout, not even an explicit nil
### GetTracks

`func (o *BusinessOrderDeliveryDTO) GetTracks() []OrderTrackDTO`

GetTracks returns the Tracks field if non-nil, zero value otherwise.

### GetTracksOk

`func (o *BusinessOrderDeliveryDTO) GetTracksOk() (*[]OrderTrackDTO, bool)`

GetTracksOk returns a tuple with the Tracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracks

`func (o *BusinessOrderDeliveryDTO) SetTracks(v []OrderTrackDTO)`

SetTracks sets Tracks field to given value.

### HasTracks

`func (o *BusinessOrderDeliveryDTO) HasTracks() bool`

HasTracks returns a boolean if a field has been set.

### SetTracksNil

`func (o *BusinessOrderDeliveryDTO) SetTracksNil(b bool)`

 SetTracksNil sets the value for Tracks to be an explicit nil

### UnsetTracks
`func (o *BusinessOrderDeliveryDTO) UnsetTracks()`

UnsetTracks ensures that no value is present for Tracks, not even an explicit nil
### GetEstimated

`func (o *BusinessOrderDeliveryDTO) GetEstimated() bool`

GetEstimated returns the Estimated field if non-nil, zero value otherwise.

### GetEstimatedOk

`func (o *BusinessOrderDeliveryDTO) GetEstimatedOk() (*bool, bool)`

GetEstimatedOk returns a tuple with the Estimated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimated

`func (o *BusinessOrderDeliveryDTO) SetEstimated(v bool)`

SetEstimated sets Estimated field to given value.

### HasEstimated

`func (o *BusinessOrderDeliveryDTO) HasEstimated() bool`

HasEstimated returns a boolean if a field has been set.

### GetReceiveBarcode

`func (o *BusinessOrderDeliveryDTO) GetReceiveBarcode() string`

GetReceiveBarcode returns the ReceiveBarcode field if non-nil, zero value otherwise.

### GetReceiveBarcodeOk

`func (o *BusinessOrderDeliveryDTO) GetReceiveBarcodeOk() (*string, bool)`

GetReceiveBarcodeOk returns a tuple with the ReceiveBarcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveBarcode

`func (o *BusinessOrderDeliveryDTO) SetReceiveBarcode(v string)`

SetReceiveBarcode sets ReceiveBarcode field to given value.

### HasReceiveBarcode

`func (o *BusinessOrderDeliveryDTO) HasReceiveBarcode() bool`

HasReceiveBarcode returns a boolean if a field has been set.

### GetReceiveCode

`func (o *BusinessOrderDeliveryDTO) GetReceiveCode() string`

GetReceiveCode returns the ReceiveCode field if non-nil, zero value otherwise.

### GetReceiveCodeOk

`func (o *BusinessOrderDeliveryDTO) GetReceiveCodeOk() (*string, bool)`

GetReceiveCodeOk returns a tuple with the ReceiveCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveCode

`func (o *BusinessOrderDeliveryDTO) SetReceiveCode(v string)`

SetReceiveCode sets ReceiveCode field to given value.

### HasReceiveCode

`func (o *BusinessOrderDeliveryDTO) HasReceiveCode() bool`

HasReceiveCode returns a boolean if a field has been set.

### GetDigitalGoods

`func (o *BusinessOrderDeliveryDTO) GetDigitalGoods() DigitalGoodsDeliveryDetailsDTO`

GetDigitalGoods returns the DigitalGoods field if non-nil, zero value otherwise.

### GetDigitalGoodsOk

`func (o *BusinessOrderDeliveryDTO) GetDigitalGoodsOk() (*DigitalGoodsDeliveryDetailsDTO, bool)`

GetDigitalGoodsOk returns a tuple with the DigitalGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalGoods

`func (o *BusinessOrderDeliveryDTO) SetDigitalGoods(v DigitalGoodsDeliveryDetailsDTO)`

SetDigitalGoods sets DigitalGoods field to given value.

### HasDigitalGoods

`func (o *BusinessOrderDeliveryDTO) HasDigitalGoods() bool`

HasDigitalGoods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


