# BusinessOrderBoxLayoutItemDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Идентификатор товара в заказе.  Параметр &#x60;id&#x60; в &#x60;items&#x60;.  | 
**FullCount** | Pointer to **int32** | Количество единиц товара в коробке.  | [optional] 
**PartialCount** | Pointer to [**BusinessOrderBoxLayoutPartialCountDTO**](BusinessOrderBoxLayoutPartialCountDTO.md) |  | [optional] 
**Instances** | Pointer to [**[]BriefOrderItemInstanceDTO**](BriefOrderItemInstanceDTO.md) | Переданные коды маркировки. | [optional] 

## Methods

### NewBusinessOrderBoxLayoutItemDTO

`func NewBusinessOrderBoxLayoutItemDTO(id int64, ) *BusinessOrderBoxLayoutItemDTO`

NewBusinessOrderBoxLayoutItemDTO instantiates a new BusinessOrderBoxLayoutItemDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessOrderBoxLayoutItemDTOWithDefaults

`func NewBusinessOrderBoxLayoutItemDTOWithDefaults() *BusinessOrderBoxLayoutItemDTO`

NewBusinessOrderBoxLayoutItemDTOWithDefaults instantiates a new BusinessOrderBoxLayoutItemDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BusinessOrderBoxLayoutItemDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BusinessOrderBoxLayoutItemDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BusinessOrderBoxLayoutItemDTO) SetId(v int64)`

SetId sets Id field to given value.


### GetFullCount

`func (o *BusinessOrderBoxLayoutItemDTO) GetFullCount() int32`

GetFullCount returns the FullCount field if non-nil, zero value otherwise.

### GetFullCountOk

`func (o *BusinessOrderBoxLayoutItemDTO) GetFullCountOk() (*int32, bool)`

GetFullCountOk returns a tuple with the FullCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullCount

`func (o *BusinessOrderBoxLayoutItemDTO) SetFullCount(v int32)`

SetFullCount sets FullCount field to given value.

### HasFullCount

`func (o *BusinessOrderBoxLayoutItemDTO) HasFullCount() bool`

HasFullCount returns a boolean if a field has been set.

### GetPartialCount

`func (o *BusinessOrderBoxLayoutItemDTO) GetPartialCount() BusinessOrderBoxLayoutPartialCountDTO`

GetPartialCount returns the PartialCount field if non-nil, zero value otherwise.

### GetPartialCountOk

`func (o *BusinessOrderBoxLayoutItemDTO) GetPartialCountOk() (*BusinessOrderBoxLayoutPartialCountDTO, bool)`

GetPartialCountOk returns a tuple with the PartialCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialCount

`func (o *BusinessOrderBoxLayoutItemDTO) SetPartialCount(v BusinessOrderBoxLayoutPartialCountDTO)`

SetPartialCount sets PartialCount field to given value.

### HasPartialCount

`func (o *BusinessOrderBoxLayoutItemDTO) HasPartialCount() bool`

HasPartialCount returns a boolean if a field has been set.

### GetInstances

`func (o *BusinessOrderBoxLayoutItemDTO) GetInstances() []BriefOrderItemInstanceDTO`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *BusinessOrderBoxLayoutItemDTO) GetInstancesOk() (*[]BriefOrderItemInstanceDTO, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *BusinessOrderBoxLayoutItemDTO) SetInstances(v []BriefOrderItemInstanceDTO)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *BusinessOrderBoxLayoutItemDTO) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### SetInstancesNil

`func (o *BusinessOrderBoxLayoutItemDTO) SetInstancesNil(b bool)`

 SetInstancesNil sets the value for Instances to be an explicit nil

### UnsetInstances
`func (o *BusinessOrderBoxLayoutItemDTO) UnsetInstances()`

UnsetInstances ensures that no value is present for Instances, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


