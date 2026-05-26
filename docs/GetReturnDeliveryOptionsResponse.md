# GetReturnDeliveryOptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ApiResponseStatusType**](ApiResponseStatusType.md) |  | 
**Result** | Pointer to [**GetReturnDeliveryOptionsDTO**](GetReturnDeliveryOptionsDTO.md) |  | [optional] 

## Methods

### NewGetReturnDeliveryOptionsResponse

`func NewGetReturnDeliveryOptionsResponse(status ApiResponseStatusType, ) *GetReturnDeliveryOptionsResponse`

NewGetReturnDeliveryOptionsResponse instantiates a new GetReturnDeliveryOptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReturnDeliveryOptionsResponseWithDefaults

`func NewGetReturnDeliveryOptionsResponseWithDefaults() *GetReturnDeliveryOptionsResponse`

NewGetReturnDeliveryOptionsResponseWithDefaults instantiates a new GetReturnDeliveryOptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetReturnDeliveryOptionsResponse) GetStatus() ApiResponseStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetReturnDeliveryOptionsResponse) GetStatusOk() (*ApiResponseStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetReturnDeliveryOptionsResponse) SetStatus(v ApiResponseStatusType)`

SetStatus sets Status field to given value.


### GetResult

`func (o *GetReturnDeliveryOptionsResponse) GetResult() GetReturnDeliveryOptionsDTO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetReturnDeliveryOptionsResponse) GetResultOk() (*GetReturnDeliveryOptionsDTO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetReturnDeliveryOptionsResponse) SetResult(v GetReturnDeliveryOptionsDTO)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetReturnDeliveryOptionsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


