# PickupOptionsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogisticPointId** | **int64** | Идентификатор пункта выдачи.  Его можно узнать с помощью метода [POST v1/businesses/{businessId}/logistics-points](../../reference/logistic-points/getLogisticPoints.md).  | 
**Options** | [**[]PickupOptionDTO**](PickupOptionDTO.md) | Варианты доставки в ПВЗ. | 

## Methods

### NewPickupOptionsDTO

`func NewPickupOptionsDTO(logisticPointId int64, options []PickupOptionDTO, ) *PickupOptionsDTO`

NewPickupOptionsDTO instantiates a new PickupOptionsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPickupOptionsDTOWithDefaults

`func NewPickupOptionsDTOWithDefaults() *PickupOptionsDTO`

NewPickupOptionsDTOWithDefaults instantiates a new PickupOptionsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogisticPointId

`func (o *PickupOptionsDTO) GetLogisticPointId() int64`

GetLogisticPointId returns the LogisticPointId field if non-nil, zero value otherwise.

### GetLogisticPointIdOk

`func (o *PickupOptionsDTO) GetLogisticPointIdOk() (*int64, bool)`

GetLogisticPointIdOk returns a tuple with the LogisticPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogisticPointId

`func (o *PickupOptionsDTO) SetLogisticPointId(v int64)`

SetLogisticPointId sets LogisticPointId field to given value.


### GetOptions

`func (o *PickupOptionsDTO) GetOptions() []PickupOptionDTO`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PickupOptionsDTO) GetOptionsOk() (*[]PickupOptionDTO, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PickupOptionsDTO) SetOptions(v []PickupOptionDTO)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


