# GetPartnerWarehousesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ApiResponseStatusType**](ApiResponseStatusType.md) |  | 
**Result** | Pointer to [**PartnerWarehousesDTO**](PartnerWarehousesDTO.md) |  | [optional] 

## Methods

### NewGetPartnerWarehousesResponse

`func NewGetPartnerWarehousesResponse(status ApiResponseStatusType, ) *GetPartnerWarehousesResponse`

NewGetPartnerWarehousesResponse instantiates a new GetPartnerWarehousesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPartnerWarehousesResponseWithDefaults

`func NewGetPartnerWarehousesResponseWithDefaults() *GetPartnerWarehousesResponse`

NewGetPartnerWarehousesResponseWithDefaults instantiates a new GetPartnerWarehousesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetPartnerWarehousesResponse) GetStatus() ApiResponseStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPartnerWarehousesResponse) GetStatusOk() (*ApiResponseStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPartnerWarehousesResponse) SetStatus(v ApiResponseStatusType)`

SetStatus sets Status field to given value.


### GetResult

`func (o *GetPartnerWarehousesResponse) GetResult() PartnerWarehousesDTO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetPartnerWarehousesResponse) GetResultOk() (*PartnerWarehousesDTO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetPartnerWarehousesResponse) SetResult(v PartnerWarehousesDTO)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetPartnerWarehousesResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


