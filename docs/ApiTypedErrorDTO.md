# ApiTypedErrorDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | [**ApiErrorCodeType**](ApiErrorCodeType.md) |  | 
**Message** | Pointer to **string** | Описание ошибки. | [optional] 

## Methods

### NewApiTypedErrorDTO

`func NewApiTypedErrorDTO(code ApiErrorCodeType, ) *ApiTypedErrorDTO`

NewApiTypedErrorDTO instantiates a new ApiTypedErrorDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTypedErrorDTOWithDefaults

`func NewApiTypedErrorDTOWithDefaults() *ApiTypedErrorDTO`

NewApiTypedErrorDTOWithDefaults instantiates a new ApiTypedErrorDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ApiTypedErrorDTO) GetCode() ApiErrorCodeType`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiTypedErrorDTO) GetCodeOk() (*ApiErrorCodeType, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiTypedErrorDTO) SetCode(v ApiErrorCodeType)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ApiTypedErrorDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiTypedErrorDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiTypedErrorDTO) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiTypedErrorDTO) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


