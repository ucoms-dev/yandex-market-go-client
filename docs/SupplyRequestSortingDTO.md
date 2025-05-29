# SupplyRequestSortingDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | [**SortOrderType**](SortOrderType.md) |  | 
**Attribute** | [**SupplyRequestSortAttributeType**](SupplyRequestSortAttributeType.md) |  | 

## Methods

### NewSupplyRequestSortingDTO

`func NewSupplyRequestSortingDTO(direction SortOrderType, attribute SupplyRequestSortAttributeType, ) *SupplyRequestSortingDTO`

NewSupplyRequestSortingDTO instantiates a new SupplyRequestSortingDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplyRequestSortingDTOWithDefaults

`func NewSupplyRequestSortingDTOWithDefaults() *SupplyRequestSortingDTO`

NewSupplyRequestSortingDTOWithDefaults instantiates a new SupplyRequestSortingDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *SupplyRequestSortingDTO) GetDirection() SortOrderType`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SupplyRequestSortingDTO) GetDirectionOk() (*SortOrderType, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SupplyRequestSortingDTO) SetDirection(v SortOrderType)`

SetDirection sets Direction field to given value.


### GetAttribute

`func (o *SupplyRequestSortingDTO) GetAttribute() SupplyRequestSortAttributeType`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *SupplyRequestSortingDTO) GetAttributeOk() (*SupplyRequestSortAttributeType, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *SupplyRequestSortingDTO) SetAttribute(v SupplyRequestSortAttributeType)`

SetAttribute sets Attribute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


