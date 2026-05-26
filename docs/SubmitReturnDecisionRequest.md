# SubmitReturnDecisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnItemDecisions** | [**[]ReturnItemDecisionDTO**](ReturnItemDecisionDTO.md) | Решения по товарам в возврате. | 

## Methods

### NewSubmitReturnDecisionRequest

`func NewSubmitReturnDecisionRequest(returnItemDecisions []ReturnItemDecisionDTO, ) *SubmitReturnDecisionRequest`

NewSubmitReturnDecisionRequest instantiates a new SubmitReturnDecisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitReturnDecisionRequestWithDefaults

`func NewSubmitReturnDecisionRequestWithDefaults() *SubmitReturnDecisionRequest`

NewSubmitReturnDecisionRequestWithDefaults instantiates a new SubmitReturnDecisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnItemDecisions

`func (o *SubmitReturnDecisionRequest) GetReturnItemDecisions() []ReturnItemDecisionDTO`

GetReturnItemDecisions returns the ReturnItemDecisions field if non-nil, zero value otherwise.

### GetReturnItemDecisionsOk

`func (o *SubmitReturnDecisionRequest) GetReturnItemDecisionsOk() (*[]ReturnItemDecisionDTO, bool)`

GetReturnItemDecisionsOk returns a tuple with the ReturnItemDecisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnItemDecisions

`func (o *SubmitReturnDecisionRequest) SetReturnItemDecisions(v []ReturnItemDecisionDTO)`

SetReturnItemDecisions sets ReturnItemDecisions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


