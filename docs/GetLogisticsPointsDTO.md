# GetLogisticsPointsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogisticPoints** | [**[]LogisticPointDTO**](LogisticPointDTO.md) | Пункты выдачи заказов. | 
**Paging** | Pointer to [**PackagingForwardScrollingPagerDTO**](PackagingForwardScrollingPagerDTO.md) |  | [optional] 

## Methods

### NewGetLogisticsPointsDTO

`func NewGetLogisticsPointsDTO(logisticPoints []LogisticPointDTO, ) *GetLogisticsPointsDTO`

NewGetLogisticsPointsDTO instantiates a new GetLogisticsPointsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLogisticsPointsDTOWithDefaults

`func NewGetLogisticsPointsDTOWithDefaults() *GetLogisticsPointsDTO`

NewGetLogisticsPointsDTOWithDefaults instantiates a new GetLogisticsPointsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogisticPoints

`func (o *GetLogisticsPointsDTO) GetLogisticPoints() []LogisticPointDTO`

GetLogisticPoints returns the LogisticPoints field if non-nil, zero value otherwise.

### GetLogisticPointsOk

`func (o *GetLogisticsPointsDTO) GetLogisticPointsOk() (*[]LogisticPointDTO, bool)`

GetLogisticPointsOk returns a tuple with the LogisticPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogisticPoints

`func (o *GetLogisticsPointsDTO) SetLogisticPoints(v []LogisticPointDTO)`

SetLogisticPoints sets LogisticPoints field to given value.


### GetPaging

`func (o *GetLogisticsPointsDTO) GetPaging() PackagingForwardScrollingPagerDTO`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *GetLogisticsPointsDTO) GetPagingOk() (*PackagingForwardScrollingPagerDTO, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *GetLogisticsPointsDTO) SetPaging(v PackagingForwardScrollingPagerDTO)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *GetLogisticsPointsDTO) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


