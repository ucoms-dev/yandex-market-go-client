# DigitalGoodsDeliveryDetailsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**DigitalGoodsDeliveryType**](DigitalGoodsDeliveryType.md) |  | 
**SteamLink** | Pointer to **string** | Ссылка на Steam-аккаунт покупателя. Передается только для типа &#x60;STEAM_GIFT&#x60;. | [optional] 

## Methods

### NewDigitalGoodsDeliveryDetailsDTO

`func NewDigitalGoodsDeliveryDetailsDTO(type_ DigitalGoodsDeliveryType, ) *DigitalGoodsDeliveryDetailsDTO`

NewDigitalGoodsDeliveryDetailsDTO instantiates a new DigitalGoodsDeliveryDetailsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalGoodsDeliveryDetailsDTOWithDefaults

`func NewDigitalGoodsDeliveryDetailsDTOWithDefaults() *DigitalGoodsDeliveryDetailsDTO`

NewDigitalGoodsDeliveryDetailsDTOWithDefaults instantiates a new DigitalGoodsDeliveryDetailsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DigitalGoodsDeliveryDetailsDTO) GetType() DigitalGoodsDeliveryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DigitalGoodsDeliveryDetailsDTO) GetTypeOk() (*DigitalGoodsDeliveryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DigitalGoodsDeliveryDetailsDTO) SetType(v DigitalGoodsDeliveryType)`

SetType sets Type field to given value.


### GetSteamLink

`func (o *DigitalGoodsDeliveryDetailsDTO) GetSteamLink() string`

GetSteamLink returns the SteamLink field if non-nil, zero value otherwise.

### GetSteamLinkOk

`func (o *DigitalGoodsDeliveryDetailsDTO) GetSteamLinkOk() (*string, bool)`

GetSteamLinkOk returns a tuple with the SteamLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamLink

`func (o *DigitalGoodsDeliveryDetailsDTO) SetSteamLink(v string)`

SetSteamLink sets SteamLink field to given value.

### HasSteamLink

`func (o *DigitalGoodsDeliveryDetailsDTO) HasSteamLink() bool`

HasSteamLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


