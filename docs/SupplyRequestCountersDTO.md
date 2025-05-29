# SupplyRequestCountersDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanCount** | Pointer to **int32** | Количество товаров в заявке на поставку. | [optional] 
**FactCount** | Pointer to **int32** | Количество товаров, которые приняты на складе. | [optional] 
**UndefinedCount** | Pointer to **int32** | Количество непринятых товаров. | [optional] 
**SurplusCount** | Pointer to **int32** | Количество лишних товаров. | [optional] 
**ShortageCount** | Pointer to **int32** | Количество товаров с недостатками. | [optional] 
**DefectCount** | Pointer to **int32** | Количество товаров с браком. | [optional] 
**AcceptableCount** | Pointer to **int32** | Количество товаров, которые можно привезти дополнительно. | [optional] 
**UnacceptableCount** | Pointer to **int32** | Количество товаров, которые нельзя привезти дополнительно. | [optional] 
**ActualPalletsCount** | Pointer to **int32** | Количество палет, которые приняты на складе. | [optional] 
**ActualBoxCount** | Pointer to **int32** | Количество коробок, которые приняты на складе. | [optional] 

## Methods

### NewSupplyRequestCountersDTO

`func NewSupplyRequestCountersDTO() *SupplyRequestCountersDTO`

NewSupplyRequestCountersDTO instantiates a new SupplyRequestCountersDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplyRequestCountersDTOWithDefaults

`func NewSupplyRequestCountersDTOWithDefaults() *SupplyRequestCountersDTO`

NewSupplyRequestCountersDTOWithDefaults instantiates a new SupplyRequestCountersDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanCount

`func (o *SupplyRequestCountersDTO) GetPlanCount() int32`

GetPlanCount returns the PlanCount field if non-nil, zero value otherwise.

### GetPlanCountOk

`func (o *SupplyRequestCountersDTO) GetPlanCountOk() (*int32, bool)`

GetPlanCountOk returns a tuple with the PlanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCount

`func (o *SupplyRequestCountersDTO) SetPlanCount(v int32)`

SetPlanCount sets PlanCount field to given value.

### HasPlanCount

`func (o *SupplyRequestCountersDTO) HasPlanCount() bool`

HasPlanCount returns a boolean if a field has been set.

### GetFactCount

`func (o *SupplyRequestCountersDTO) GetFactCount() int32`

GetFactCount returns the FactCount field if non-nil, zero value otherwise.

### GetFactCountOk

`func (o *SupplyRequestCountersDTO) GetFactCountOk() (*int32, bool)`

GetFactCountOk returns a tuple with the FactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactCount

`func (o *SupplyRequestCountersDTO) SetFactCount(v int32)`

SetFactCount sets FactCount field to given value.

### HasFactCount

`func (o *SupplyRequestCountersDTO) HasFactCount() bool`

HasFactCount returns a boolean if a field has been set.

### GetUndefinedCount

`func (o *SupplyRequestCountersDTO) GetUndefinedCount() int32`

GetUndefinedCount returns the UndefinedCount field if non-nil, zero value otherwise.

### GetUndefinedCountOk

`func (o *SupplyRequestCountersDTO) GetUndefinedCountOk() (*int32, bool)`

GetUndefinedCountOk returns a tuple with the UndefinedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndefinedCount

`func (o *SupplyRequestCountersDTO) SetUndefinedCount(v int32)`

SetUndefinedCount sets UndefinedCount field to given value.

### HasUndefinedCount

`func (o *SupplyRequestCountersDTO) HasUndefinedCount() bool`

HasUndefinedCount returns a boolean if a field has been set.

### GetSurplusCount

`func (o *SupplyRequestCountersDTO) GetSurplusCount() int32`

GetSurplusCount returns the SurplusCount field if non-nil, zero value otherwise.

### GetSurplusCountOk

`func (o *SupplyRequestCountersDTO) GetSurplusCountOk() (*int32, bool)`

GetSurplusCountOk returns a tuple with the SurplusCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurplusCount

`func (o *SupplyRequestCountersDTO) SetSurplusCount(v int32)`

SetSurplusCount sets SurplusCount field to given value.

### HasSurplusCount

`func (o *SupplyRequestCountersDTO) HasSurplusCount() bool`

HasSurplusCount returns a boolean if a field has been set.

### GetShortageCount

`func (o *SupplyRequestCountersDTO) GetShortageCount() int32`

GetShortageCount returns the ShortageCount field if non-nil, zero value otherwise.

### GetShortageCountOk

`func (o *SupplyRequestCountersDTO) GetShortageCountOk() (*int32, bool)`

GetShortageCountOk returns a tuple with the ShortageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortageCount

`func (o *SupplyRequestCountersDTO) SetShortageCount(v int32)`

SetShortageCount sets ShortageCount field to given value.

### HasShortageCount

`func (o *SupplyRequestCountersDTO) HasShortageCount() bool`

HasShortageCount returns a boolean if a field has been set.

### GetDefectCount

`func (o *SupplyRequestCountersDTO) GetDefectCount() int32`

GetDefectCount returns the DefectCount field if non-nil, zero value otherwise.

### GetDefectCountOk

`func (o *SupplyRequestCountersDTO) GetDefectCountOk() (*int32, bool)`

GetDefectCountOk returns a tuple with the DefectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectCount

`func (o *SupplyRequestCountersDTO) SetDefectCount(v int32)`

SetDefectCount sets DefectCount field to given value.

### HasDefectCount

`func (o *SupplyRequestCountersDTO) HasDefectCount() bool`

HasDefectCount returns a boolean if a field has been set.

### GetAcceptableCount

`func (o *SupplyRequestCountersDTO) GetAcceptableCount() int32`

GetAcceptableCount returns the AcceptableCount field if non-nil, zero value otherwise.

### GetAcceptableCountOk

`func (o *SupplyRequestCountersDTO) GetAcceptableCountOk() (*int32, bool)`

GetAcceptableCountOk returns a tuple with the AcceptableCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptableCount

`func (o *SupplyRequestCountersDTO) SetAcceptableCount(v int32)`

SetAcceptableCount sets AcceptableCount field to given value.

### HasAcceptableCount

`func (o *SupplyRequestCountersDTO) HasAcceptableCount() bool`

HasAcceptableCount returns a boolean if a field has been set.

### GetUnacceptableCount

`func (o *SupplyRequestCountersDTO) GetUnacceptableCount() int32`

GetUnacceptableCount returns the UnacceptableCount field if non-nil, zero value otherwise.

### GetUnacceptableCountOk

`func (o *SupplyRequestCountersDTO) GetUnacceptableCountOk() (*int32, bool)`

GetUnacceptableCountOk returns a tuple with the UnacceptableCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnacceptableCount

`func (o *SupplyRequestCountersDTO) SetUnacceptableCount(v int32)`

SetUnacceptableCount sets UnacceptableCount field to given value.

### HasUnacceptableCount

`func (o *SupplyRequestCountersDTO) HasUnacceptableCount() bool`

HasUnacceptableCount returns a boolean if a field has been set.

### GetActualPalletsCount

`func (o *SupplyRequestCountersDTO) GetActualPalletsCount() int32`

GetActualPalletsCount returns the ActualPalletsCount field if non-nil, zero value otherwise.

### GetActualPalletsCountOk

`func (o *SupplyRequestCountersDTO) GetActualPalletsCountOk() (*int32, bool)`

GetActualPalletsCountOk returns a tuple with the ActualPalletsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualPalletsCount

`func (o *SupplyRequestCountersDTO) SetActualPalletsCount(v int32)`

SetActualPalletsCount sets ActualPalletsCount field to given value.

### HasActualPalletsCount

`func (o *SupplyRequestCountersDTO) HasActualPalletsCount() bool`

HasActualPalletsCount returns a boolean if a field has been set.

### GetActualBoxCount

`func (o *SupplyRequestCountersDTO) GetActualBoxCount() int32`

GetActualBoxCount returns the ActualBoxCount field if non-nil, zero value otherwise.

### GetActualBoxCountOk

`func (o *SupplyRequestCountersDTO) GetActualBoxCountOk() (*int32, bool)`

GetActualBoxCountOk returns a tuple with the ActualBoxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualBoxCount

`func (o *SupplyRequestCountersDTO) SetActualBoxCount(v int32)`

SetActualBoxCount sets ActualBoxCount field to given value.

### HasActualBoxCount

`func (o *SupplyRequestCountersDTO) HasActualBoxCount() bool`

HasActualBoxCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


