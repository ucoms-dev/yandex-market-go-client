# ReturnItemDecisionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnItemId** | **int64** | Идентификатор товара в возврате. | 
**DecisionType** | [**ReturnRequestDecisionType**](ReturnRequestDecisionType.md) |  | 
**DecisionReasonType** | Pointer to [**ReturnRequestDecisionReasonType**](ReturnRequestDecisionReasonType.md) |  | [optional] 
**Comment** | Pointer to **string** | Комментарий к решению. Укажите:  * для &#x60;REFUND_MONEY_INCLUDING_SHIPMENT&#x60;— стоимость обратной пересылки.  * для &#x60;REPAIR&#x60; — когда вы устраните недостатки товара.  * для &#x60;DECLINE_REFUND&#x60; — причину отказа.  * для &#x60;OTHER_DECISION&#x60; — какое решение вы предлагаете.  | [optional] 
**Compensation** | Pointer to [**BasePriceDTO**](BasePriceDTO.md) |  | [optional] 

## Methods

### NewReturnItemDecisionDTO

`func NewReturnItemDecisionDTO(returnItemId int64, decisionType ReturnRequestDecisionType, ) *ReturnItemDecisionDTO`

NewReturnItemDecisionDTO instantiates a new ReturnItemDecisionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnItemDecisionDTOWithDefaults

`func NewReturnItemDecisionDTOWithDefaults() *ReturnItemDecisionDTO`

NewReturnItemDecisionDTOWithDefaults instantiates a new ReturnItemDecisionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnItemId

`func (o *ReturnItemDecisionDTO) GetReturnItemId() int64`

GetReturnItemId returns the ReturnItemId field if non-nil, zero value otherwise.

### GetReturnItemIdOk

`func (o *ReturnItemDecisionDTO) GetReturnItemIdOk() (*int64, bool)`

GetReturnItemIdOk returns a tuple with the ReturnItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnItemId

`func (o *ReturnItemDecisionDTO) SetReturnItemId(v int64)`

SetReturnItemId sets ReturnItemId field to given value.


### GetDecisionType

`func (o *ReturnItemDecisionDTO) GetDecisionType() ReturnRequestDecisionType`

GetDecisionType returns the DecisionType field if non-nil, zero value otherwise.

### GetDecisionTypeOk

`func (o *ReturnItemDecisionDTO) GetDecisionTypeOk() (*ReturnRequestDecisionType, bool)`

GetDecisionTypeOk returns a tuple with the DecisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionType

`func (o *ReturnItemDecisionDTO) SetDecisionType(v ReturnRequestDecisionType)`

SetDecisionType sets DecisionType field to given value.


### GetDecisionReasonType

`func (o *ReturnItemDecisionDTO) GetDecisionReasonType() ReturnRequestDecisionReasonType`

GetDecisionReasonType returns the DecisionReasonType field if non-nil, zero value otherwise.

### GetDecisionReasonTypeOk

`func (o *ReturnItemDecisionDTO) GetDecisionReasonTypeOk() (*ReturnRequestDecisionReasonType, bool)`

GetDecisionReasonTypeOk returns a tuple with the DecisionReasonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionReasonType

`func (o *ReturnItemDecisionDTO) SetDecisionReasonType(v ReturnRequestDecisionReasonType)`

SetDecisionReasonType sets DecisionReasonType field to given value.

### HasDecisionReasonType

`func (o *ReturnItemDecisionDTO) HasDecisionReasonType() bool`

HasDecisionReasonType returns a boolean if a field has been set.

### GetComment

`func (o *ReturnItemDecisionDTO) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ReturnItemDecisionDTO) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ReturnItemDecisionDTO) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ReturnItemDecisionDTO) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCompensation

`func (o *ReturnItemDecisionDTO) GetCompensation() BasePriceDTO`

GetCompensation returns the Compensation field if non-nil, zero value otherwise.

### GetCompensationOk

`func (o *ReturnItemDecisionDTO) GetCompensationOk() (*BasePriceDTO, bool)`

GetCompensationOk returns a tuple with the Compensation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompensation

`func (o *ReturnItemDecisionDTO) SetCompensation(v BasePriceDTO)`

SetCompensation sets Compensation field to given value.

### HasCompensation

`func (o *ReturnItemDecisionDTO) HasCompensation() bool`

HasCompensation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


