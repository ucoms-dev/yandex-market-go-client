# SupplyRequestItemCountersDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanCount** | Pointer to **int32** | Количество товаров в заявке на поставку. | [optional] 
**FactCount** | Pointer to **int32** | Количество товаров, которые приняты на складе. | [optional] 
**SurplusCount** | Pointer to **int32** | Количество лишних товаров. | [optional] 
**ShortageCount** | Pointer to **int32** | Количество товаров с недостатками. | [optional] 
**DefectCount** | Pointer to **int32** | Количество товаров с браком. | [optional] 

## Methods

### NewSupplyRequestItemCountersDTO

`func NewSupplyRequestItemCountersDTO() *SupplyRequestItemCountersDTO`

NewSupplyRequestItemCountersDTO instantiates a new SupplyRequestItemCountersDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplyRequestItemCountersDTOWithDefaults

`func NewSupplyRequestItemCountersDTOWithDefaults() *SupplyRequestItemCountersDTO`

NewSupplyRequestItemCountersDTOWithDefaults instantiates a new SupplyRequestItemCountersDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanCount

`func (o *SupplyRequestItemCountersDTO) GetPlanCount() int32`

GetPlanCount returns the PlanCount field if non-nil, zero value otherwise.

### GetPlanCountOk

`func (o *SupplyRequestItemCountersDTO) GetPlanCountOk() (*int32, bool)`

GetPlanCountOk returns a tuple with the PlanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCount

`func (o *SupplyRequestItemCountersDTO) SetPlanCount(v int32)`

SetPlanCount sets PlanCount field to given value.

### HasPlanCount

`func (o *SupplyRequestItemCountersDTO) HasPlanCount() bool`

HasPlanCount returns a boolean if a field has been set.

### GetFactCount

`func (o *SupplyRequestItemCountersDTO) GetFactCount() int32`

GetFactCount returns the FactCount field if non-nil, zero value otherwise.

### GetFactCountOk

`func (o *SupplyRequestItemCountersDTO) GetFactCountOk() (*int32, bool)`

GetFactCountOk returns a tuple with the FactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactCount

`func (o *SupplyRequestItemCountersDTO) SetFactCount(v int32)`

SetFactCount sets FactCount field to given value.

### HasFactCount

`func (o *SupplyRequestItemCountersDTO) HasFactCount() bool`

HasFactCount returns a boolean if a field has been set.

### GetSurplusCount

`func (o *SupplyRequestItemCountersDTO) GetSurplusCount() int32`

GetSurplusCount returns the SurplusCount field if non-nil, zero value otherwise.

### GetSurplusCountOk

`func (o *SupplyRequestItemCountersDTO) GetSurplusCountOk() (*int32, bool)`

GetSurplusCountOk returns a tuple with the SurplusCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurplusCount

`func (o *SupplyRequestItemCountersDTO) SetSurplusCount(v int32)`

SetSurplusCount sets SurplusCount field to given value.

### HasSurplusCount

`func (o *SupplyRequestItemCountersDTO) HasSurplusCount() bool`

HasSurplusCount returns a boolean if a field has been set.

### GetShortageCount

`func (o *SupplyRequestItemCountersDTO) GetShortageCount() int32`

GetShortageCount returns the ShortageCount field if non-nil, zero value otherwise.

### GetShortageCountOk

`func (o *SupplyRequestItemCountersDTO) GetShortageCountOk() (*int32, bool)`

GetShortageCountOk returns a tuple with the ShortageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortageCount

`func (o *SupplyRequestItemCountersDTO) SetShortageCount(v int32)`

SetShortageCount sets ShortageCount field to given value.

### HasShortageCount

`func (o *SupplyRequestItemCountersDTO) HasShortageCount() bool`

HasShortageCount returns a boolean if a field has been set.

### GetDefectCount

`func (o *SupplyRequestItemCountersDTO) GetDefectCount() int32`

GetDefectCount returns the DefectCount field if non-nil, zero value otherwise.

### GetDefectCountOk

`func (o *SupplyRequestItemCountersDTO) GetDefectCountOk() (*int32, bool)`

GetDefectCountOk returns a tuple with the DefectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectCount

`func (o *SupplyRequestItemCountersDTO) SetDefectCount(v int32)`

SetDefectCount sets DefectCount field to given value.

### HasDefectCount

`func (o *SupplyRequestItemCountersDTO) HasDefectCount() bool`

HasDefectCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


