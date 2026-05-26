# CreateReturnResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ApiResponseStatusType**](ApiResponseStatusType.md) |  | 
**Result** | Pointer to [**CreatedReturnDTO**](CreatedReturnDTO.md) |  | [optional] 

## Methods

### NewCreateReturnResponse

`func NewCreateReturnResponse(status ApiResponseStatusType, ) *CreateReturnResponse`

NewCreateReturnResponse instantiates a new CreateReturnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReturnResponseWithDefaults

`func NewCreateReturnResponseWithDefaults() *CreateReturnResponse`

NewCreateReturnResponseWithDefaults instantiates a new CreateReturnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateReturnResponse) GetStatus() ApiResponseStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateReturnResponse) GetStatusOk() (*ApiResponseStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateReturnResponse) SetStatus(v ApiResponseStatusType)`

SetStatus sets Status field to given value.


### GetResult

`func (o *CreateReturnResponse) GetResult() CreatedReturnDTO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CreateReturnResponse) GetResultOk() (*CreatedReturnDTO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CreateReturnResponse) SetResult(v CreatedReturnDTO)`

SetResult sets Result field to given value.

### HasResult

`func (o *CreateReturnResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


