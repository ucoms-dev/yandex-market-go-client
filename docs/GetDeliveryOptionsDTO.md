# GetDeliveryOptionsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WarehousesDeliveryOptions** | [**[]WarehousesDeliveryOptionsDTO**](WarehousesDeliveryOptionsDTO.md) | Варианты доставки для разных складов. | 

## Methods

### NewGetDeliveryOptionsDTO

`func NewGetDeliveryOptionsDTO(warehousesDeliveryOptions []WarehousesDeliveryOptionsDTO, ) *GetDeliveryOptionsDTO`

NewGetDeliveryOptionsDTO instantiates a new GetDeliveryOptionsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeliveryOptionsDTOWithDefaults

`func NewGetDeliveryOptionsDTOWithDefaults() *GetDeliveryOptionsDTO`

NewGetDeliveryOptionsDTOWithDefaults instantiates a new GetDeliveryOptionsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarehousesDeliveryOptions

`func (o *GetDeliveryOptionsDTO) GetWarehousesDeliveryOptions() []WarehousesDeliveryOptionsDTO`

GetWarehousesDeliveryOptions returns the WarehousesDeliveryOptions field if non-nil, zero value otherwise.

### GetWarehousesDeliveryOptionsOk

`func (o *GetDeliveryOptionsDTO) GetWarehousesDeliveryOptionsOk() (*[]WarehousesDeliveryOptionsDTO, bool)`

GetWarehousesDeliveryOptionsOk returns a tuple with the WarehousesDeliveryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehousesDeliveryOptions

`func (o *GetDeliveryOptionsDTO) SetWarehousesDeliveryOptions(v []WarehousesDeliveryOptionsDTO)`

SetWarehousesDeliveryOptions sets WarehousesDeliveryOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


