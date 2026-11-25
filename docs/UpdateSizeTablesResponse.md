# UpdateSizeTablesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ApiResponseStatusType**](ApiResponseStatusType.md) |  | 
**Result** | [**SaveSizeTablesResultDTO**](SaveSizeTablesResultDTO.md) |  | 

## Methods

### NewUpdateSizeTablesResponse

`func NewUpdateSizeTablesResponse(status ApiResponseStatusType, result SaveSizeTablesResultDTO, ) *UpdateSizeTablesResponse`

NewUpdateSizeTablesResponse instantiates a new UpdateSizeTablesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSizeTablesResponseWithDefaults

`func NewUpdateSizeTablesResponseWithDefaults() *UpdateSizeTablesResponse`

NewUpdateSizeTablesResponseWithDefaults instantiates a new UpdateSizeTablesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateSizeTablesResponse) GetStatus() ApiResponseStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateSizeTablesResponse) GetStatusOk() (*ApiResponseStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateSizeTablesResponse) SetStatus(v ApiResponseStatusType)`

SetStatus sets Status field to given value.


### GetResult

`func (o *UpdateSizeTablesResponse) GetResult() SaveSizeTablesResultDTO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UpdateSizeTablesResponse) GetResultOk() (*SaveSizeTablesResultDTO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UpdateSizeTablesResponse) SetResult(v SaveSizeTablesResultDTO)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


