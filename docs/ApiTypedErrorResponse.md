# ApiTypedErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ApiResponseStatusType**](ApiResponseStatusType.md) |  | 
**Errors** | Pointer to [**[]ApiTypedErrorDTO**](ApiTypedErrorDTO.md) | Список ошибок. | [optional] 

## Methods

### NewApiTypedErrorResponse

`func NewApiTypedErrorResponse(status ApiResponseStatusType, ) *ApiTypedErrorResponse`

NewApiTypedErrorResponse instantiates a new ApiTypedErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTypedErrorResponseWithDefaults

`func NewApiTypedErrorResponseWithDefaults() *ApiTypedErrorResponse`

NewApiTypedErrorResponseWithDefaults instantiates a new ApiTypedErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ApiTypedErrorResponse) GetStatus() ApiResponseStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiTypedErrorResponse) GetStatusOk() (*ApiResponseStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiTypedErrorResponse) SetStatus(v ApiResponseStatusType)`

SetStatus sets Status field to given value.


### GetErrors

`func (o *ApiTypedErrorResponse) GetErrors() []ApiTypedErrorDTO`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ApiTypedErrorResponse) GetErrorsOk() (*[]ApiTypedErrorDTO, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ApiTypedErrorResponse) SetErrors(v []ApiTypedErrorDTO)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ApiTypedErrorResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *ApiTypedErrorResponse) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *ApiTypedErrorResponse) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


