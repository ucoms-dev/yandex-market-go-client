# ShipmentBoxesDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boxes** | [**[]ParcelBoxDTO**](ParcelBoxDTO.md) | Список грузовых мест. По его длине Маркет определил количество мест.  | 

## Methods

### NewShipmentBoxesDTO

`func NewShipmentBoxesDTO(boxes []ParcelBoxDTO, ) *ShipmentBoxesDTO`

NewShipmentBoxesDTO instantiates a new ShipmentBoxesDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentBoxesDTOWithDefaults

`func NewShipmentBoxesDTOWithDefaults() *ShipmentBoxesDTO`

NewShipmentBoxesDTOWithDefaults instantiates a new ShipmentBoxesDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoxes

`func (o *ShipmentBoxesDTO) GetBoxes() []ParcelBoxDTO`

GetBoxes returns the Boxes field if non-nil, zero value otherwise.

### GetBoxesOk

`func (o *ShipmentBoxesDTO) GetBoxesOk() (*[]ParcelBoxDTO, bool)`

GetBoxesOk returns a tuple with the Boxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxes

`func (o *ShipmentBoxesDTO) SetBoxes(v []ParcelBoxDTO)`

SetBoxes sets Boxes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


