# ApiTypedUnauthorizedErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ApiResponseStatusType**](ApiResponseStatusType.md) |  | 
**Errors** | Pointer to [**[]ApiTypedErrorDTO**](ApiTypedErrorDTO.md) | Список ошибок. | [optional] 

## Methods

### NewApiTypedUnauthorizedErrorResponse

`func NewApiTypedUnauthorizedErrorResponse(status ApiResponseStatusType, ) *ApiTypedUnauthorizedErrorResponse`

NewApiTypedUnauthorizedErrorResponse instantiates a new ApiTypedUnauthorizedErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTypedUnauthorizedErrorResponseWithDefaults

`func NewApiTypedUnauthorizedErrorResponseWithDefaults() *ApiTypedUnauthorizedErrorResponse`

NewApiTypedUnauthorizedErrorResponseWithDefaults instantiates a new ApiTypedUnauthorizedErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ApiTypedUnauthorizedErrorResponse) GetStatus() ApiResponseStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiTypedUnauthorizedErrorResponse) GetStatusOk() (*ApiResponseStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiTypedUnauthorizedErrorResponse) SetStatus(v ApiResponseStatusType)`

SetStatus sets Status field to given value.


### GetErrors

`func (o *ApiTypedUnauthorizedErrorResponse) GetErrors() []ApiTypedErrorDTO`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ApiTypedUnauthorizedErrorResponse) GetErrorsOk() (*[]ApiTypedErrorDTO, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ApiTypedUnauthorizedErrorResponse) SetErrors(v []ApiTypedErrorDTO)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ApiTypedUnauthorizedErrorResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *ApiTypedUnauthorizedErrorResponse) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *ApiTypedUnauthorizedErrorResponse) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


