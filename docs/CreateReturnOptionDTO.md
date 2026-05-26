# CreateReturnOptionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PickupReturn** | [**OrderPickupReturnDTO**](OrderPickupReturnDTO.md) |  | 

## Methods

### NewCreateReturnOptionDTO

`func NewCreateReturnOptionDTO(pickupReturn OrderPickupReturnDTO, ) *CreateReturnOptionDTO`

NewCreateReturnOptionDTO instantiates a new CreateReturnOptionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReturnOptionDTOWithDefaults

`func NewCreateReturnOptionDTOWithDefaults() *CreateReturnOptionDTO`

NewCreateReturnOptionDTOWithDefaults instantiates a new CreateReturnOptionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPickupReturn

`func (o *CreateReturnOptionDTO) GetPickupReturn() OrderPickupReturnDTO`

GetPickupReturn returns the PickupReturn field if non-nil, zero value otherwise.

### GetPickupReturnOk

`func (o *CreateReturnOptionDTO) GetPickupReturnOk() (*OrderPickupReturnDTO, bool)`

GetPickupReturnOk returns a tuple with the PickupReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupReturn

`func (o *CreateReturnOptionDTO) SetPickupReturn(v OrderPickupReturnDTO)`

SetPickupReturn sets PickupReturn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


