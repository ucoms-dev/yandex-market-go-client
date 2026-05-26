# GetRegionByIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | Pointer to [**[]RegionDTO**](RegionDTO.md) | Регион доставки.  {% note warning %}  В массиве всегда возвращается один регион, используйте поле &#x60;region&#x60; вместо него.  {% endnote %}  | [optional] 
**Region** | [**RegionDTO**](RegionDTO.md) |  | 

## Methods

### NewGetRegionByIdResponse

`func NewGetRegionByIdResponse(region RegionDTO, ) *GetRegionByIdResponse`

NewGetRegionByIdResponse instantiates a new GetRegionByIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRegionByIdResponseWithDefaults

`func NewGetRegionByIdResponseWithDefaults() *GetRegionByIdResponse`

NewGetRegionByIdResponseWithDefaults instantiates a new GetRegionByIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegions

`func (o *GetRegionByIdResponse) GetRegions() []RegionDTO`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *GetRegionByIdResponse) GetRegionsOk() (*[]RegionDTO, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *GetRegionByIdResponse) SetRegions(v []RegionDTO)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *GetRegionByIdResponse) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### SetRegionsNil

`func (o *GetRegionByIdResponse) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *GetRegionByIdResponse) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil
### GetRegion

`func (o *GetRegionByIdResponse) GetRegion() RegionDTO`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetRegionByIdResponse) GetRegionOk() (*RegionDTO, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetRegionByIdResponse) SetRegion(v RegionDTO)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


