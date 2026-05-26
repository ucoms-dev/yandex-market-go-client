# ReturnAvailableDecisionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecisionType** | [**ReturnDecisionType**](ReturnDecisionType.md) |  | 
**DecisionReasonTypes** | Pointer to [**[]ReturnRequestDecisionReasonType**](ReturnRequestDecisionReasonType.md) | Возможные причины отказа (только для решения DECLINE_REFUND). | [optional] 
**PartialCompensationBounds** | Pointer to [**PartialCompensationBoundsDTO**](PartialCompensationBoundsDTO.md) |  | [optional] 

## Methods

### NewReturnAvailableDecisionDTO

`func NewReturnAvailableDecisionDTO(decisionType ReturnDecisionType, ) *ReturnAvailableDecisionDTO`

NewReturnAvailableDecisionDTO instantiates a new ReturnAvailableDecisionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnAvailableDecisionDTOWithDefaults

`func NewReturnAvailableDecisionDTOWithDefaults() *ReturnAvailableDecisionDTO`

NewReturnAvailableDecisionDTOWithDefaults instantiates a new ReturnAvailableDecisionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecisionType

`func (o *ReturnAvailableDecisionDTO) GetDecisionType() ReturnDecisionType`

GetDecisionType returns the DecisionType field if non-nil, zero value otherwise.

### GetDecisionTypeOk

`func (o *ReturnAvailableDecisionDTO) GetDecisionTypeOk() (*ReturnDecisionType, bool)`

GetDecisionTypeOk returns a tuple with the DecisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionType

`func (o *ReturnAvailableDecisionDTO) SetDecisionType(v ReturnDecisionType)`

SetDecisionType sets DecisionType field to given value.


### GetDecisionReasonTypes

`func (o *ReturnAvailableDecisionDTO) GetDecisionReasonTypes() []ReturnRequestDecisionReasonType`

GetDecisionReasonTypes returns the DecisionReasonTypes field if non-nil, zero value otherwise.

### GetDecisionReasonTypesOk

`func (o *ReturnAvailableDecisionDTO) GetDecisionReasonTypesOk() (*[]ReturnRequestDecisionReasonType, bool)`

GetDecisionReasonTypesOk returns a tuple with the DecisionReasonTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionReasonTypes

`func (o *ReturnAvailableDecisionDTO) SetDecisionReasonTypes(v []ReturnRequestDecisionReasonType)`

SetDecisionReasonTypes sets DecisionReasonTypes field to given value.

### HasDecisionReasonTypes

`func (o *ReturnAvailableDecisionDTO) HasDecisionReasonTypes() bool`

HasDecisionReasonTypes returns a boolean if a field has been set.

### SetDecisionReasonTypesNil

`func (o *ReturnAvailableDecisionDTO) SetDecisionReasonTypesNil(b bool)`

 SetDecisionReasonTypesNil sets the value for DecisionReasonTypes to be an explicit nil

### UnsetDecisionReasonTypes
`func (o *ReturnAvailableDecisionDTO) UnsetDecisionReasonTypes()`

UnsetDecisionReasonTypes ensures that no value is present for DecisionReasonTypes, not even an explicit nil
### GetPartialCompensationBounds

`func (o *ReturnAvailableDecisionDTO) GetPartialCompensationBounds() PartialCompensationBoundsDTO`

GetPartialCompensationBounds returns the PartialCompensationBounds field if non-nil, zero value otherwise.

### GetPartialCompensationBoundsOk

`func (o *ReturnAvailableDecisionDTO) GetPartialCompensationBoundsOk() (*PartialCompensationBoundsDTO, bool)`

GetPartialCompensationBoundsOk returns a tuple with the PartialCompensationBounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialCompensationBounds

`func (o *ReturnAvailableDecisionDTO) SetPartialCompensationBounds(v PartialCompensationBoundsDTO)`

SetPartialCompensationBounds sets PartialCompensationBounds field to given value.

### HasPartialCompensationBounds

`func (o *ReturnAvailableDecisionDTO) HasPartialCompensationBounds() bool`

HasPartialCompensationBounds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


