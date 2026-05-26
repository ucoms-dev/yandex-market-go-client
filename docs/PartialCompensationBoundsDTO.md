# PartialCompensationBoundsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinAmount** | [**BasePriceDTO**](BasePriceDTO.md) |  | 
**MaxAmount** | [**BasePriceDTO**](BasePriceDTO.md) |  | 
**MaxPercent** | **int64** | Верхний предел доли суммы позиции, которую можно компенсировать (в процентах). | 

## Methods

### NewPartialCompensationBoundsDTO

`func NewPartialCompensationBoundsDTO(minAmount BasePriceDTO, maxAmount BasePriceDTO, maxPercent int64, ) *PartialCompensationBoundsDTO`

NewPartialCompensationBoundsDTO instantiates a new PartialCompensationBoundsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialCompensationBoundsDTOWithDefaults

`func NewPartialCompensationBoundsDTOWithDefaults() *PartialCompensationBoundsDTO`

NewPartialCompensationBoundsDTOWithDefaults instantiates a new PartialCompensationBoundsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinAmount

`func (o *PartialCompensationBoundsDTO) GetMinAmount() BasePriceDTO`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *PartialCompensationBoundsDTO) GetMinAmountOk() (*BasePriceDTO, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *PartialCompensationBoundsDTO) SetMinAmount(v BasePriceDTO)`

SetMinAmount sets MinAmount field to given value.


### GetMaxAmount

`func (o *PartialCompensationBoundsDTO) GetMaxAmount() BasePriceDTO`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *PartialCompensationBoundsDTO) GetMaxAmountOk() (*BasePriceDTO, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *PartialCompensationBoundsDTO) SetMaxAmount(v BasePriceDTO)`

SetMaxAmount sets MaxAmount field to given value.


### GetMaxPercent

`func (o *PartialCompensationBoundsDTO) GetMaxPercent() int64`

GetMaxPercent returns the MaxPercent field if non-nil, zero value otherwise.

### GetMaxPercentOk

`func (o *PartialCompensationBoundsDTO) GetMaxPercentOk() (*int64, bool)`

GetMaxPercentOk returns a tuple with the MaxPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPercent

`func (o *PartialCompensationBoundsDTO) SetMaxPercent(v int64)`

SetMaxPercent sets MaxPercent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


