# RegionWithChildrenDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Идентификатор региона. | 
**Name** | **string** | Название региона. | 
**Type** | [**RegionType**](RegionType.md) |  | 
**Parent** | Pointer to [**RegionDTO**](RegionDTO.md) |  | [optional] 
**Children** | Pointer to [**[]RegionDTO**](RegionDTO.md) | Дочерние регионы. | [optional] 

## Methods

### NewRegionWithChildrenDTO

`func NewRegionWithChildrenDTO(id int64, name string, type_ RegionType, ) *RegionWithChildrenDTO`

NewRegionWithChildrenDTO instantiates a new RegionWithChildrenDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionWithChildrenDTOWithDefaults

`func NewRegionWithChildrenDTOWithDefaults() *RegionWithChildrenDTO`

NewRegionWithChildrenDTOWithDefaults instantiates a new RegionWithChildrenDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegionWithChildrenDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegionWithChildrenDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegionWithChildrenDTO) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *RegionWithChildrenDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegionWithChildrenDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegionWithChildrenDTO) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *RegionWithChildrenDTO) GetType() RegionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegionWithChildrenDTO) GetTypeOk() (*RegionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegionWithChildrenDTO) SetType(v RegionType)`

SetType sets Type field to given value.


### GetParent

`func (o *RegionWithChildrenDTO) GetParent() RegionDTO`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *RegionWithChildrenDTO) GetParentOk() (*RegionDTO, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *RegionWithChildrenDTO) SetParent(v RegionDTO)`

SetParent sets Parent field to given value.

### HasParent

`func (o *RegionWithChildrenDTO) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetChildren

`func (o *RegionWithChildrenDTO) GetChildren() []RegionDTO`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *RegionWithChildrenDTO) GetChildrenOk() (*[]RegionDTO, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *RegionWithChildrenDTO) SetChildren(v []RegionDTO)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *RegionWithChildrenDTO) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildrenNil

`func (o *RegionWithChildrenDTO) SetChildrenNil(b bool)`

 SetChildrenNil sets the value for Children to be an explicit nil

### UnsetChildren
`func (o *RegionWithChildrenDTO) UnsetChildren()`

UnsetChildren ensures that no value is present for Children, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


