# ParcelRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boxes** | [**[]ParcelBoxRequestDTO**](ParcelBoxRequestDTO.md) | Список грузовых мест. По его длине Маркет определяет количество мест. | 

## Methods

### NewParcelRequestDTO

`func NewParcelRequestDTO(boxes []ParcelBoxRequestDTO, ) *ParcelRequestDTO`

NewParcelRequestDTO instantiates a new ParcelRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParcelRequestDTOWithDefaults

`func NewParcelRequestDTOWithDefaults() *ParcelRequestDTO`

NewParcelRequestDTOWithDefaults instantiates a new ParcelRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoxes

`func (o *ParcelRequestDTO) GetBoxes() []ParcelBoxRequestDTO`

GetBoxes returns the Boxes field if non-nil, zero value otherwise.

### GetBoxesOk

`func (o *ParcelRequestDTO) GetBoxesOk() (*[]ParcelBoxRequestDTO, bool)`

GetBoxesOk returns a tuple with the Boxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxes

`func (o *ParcelRequestDTO) SetBoxes(v []ParcelBoxRequestDTO)`

SetBoxes sets Boxes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


