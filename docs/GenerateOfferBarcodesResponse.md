# GenerateOfferBarcodesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ApiResponseStatusType**](ApiResponseStatusType.md) |  | 
**Result** | Pointer to [**GenerateOfferBarcodesResultDTO**](GenerateOfferBarcodesResultDTO.md) |  | [optional] 

## Methods

### NewGenerateOfferBarcodesResponse

`func NewGenerateOfferBarcodesResponse(status ApiResponseStatusType, ) *GenerateOfferBarcodesResponse`

NewGenerateOfferBarcodesResponse instantiates a new GenerateOfferBarcodesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateOfferBarcodesResponseWithDefaults

`func NewGenerateOfferBarcodesResponseWithDefaults() *GenerateOfferBarcodesResponse`

NewGenerateOfferBarcodesResponseWithDefaults instantiates a new GenerateOfferBarcodesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GenerateOfferBarcodesResponse) GetStatus() ApiResponseStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GenerateOfferBarcodesResponse) GetStatusOk() (*ApiResponseStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GenerateOfferBarcodesResponse) SetStatus(v ApiResponseStatusType)`

SetStatus sets Status field to given value.


### GetResult

`func (o *GenerateOfferBarcodesResponse) GetResult() GenerateOfferBarcodesResultDTO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GenerateOfferBarcodesResponse) GetResultOk() (*GenerateOfferBarcodesResultDTO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GenerateOfferBarcodesResponse) SetResult(v GenerateOfferBarcodesResultDTO)`

SetResult sets Result field to given value.

### HasResult

`func (o *GenerateOfferBarcodesResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


