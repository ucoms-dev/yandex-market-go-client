# GetStocksOnPartnerWarehousesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ApiResponseStatusType**](ApiResponseStatusType.md) |  | 
**Result** | Pointer to [**GetStocksOnPartnerWarehousesDTO**](GetStocksOnPartnerWarehousesDTO.md) |  | [optional] 

## Methods

### NewGetStocksOnPartnerWarehousesResponse

`func NewGetStocksOnPartnerWarehousesResponse(status ApiResponseStatusType, ) *GetStocksOnPartnerWarehousesResponse`

NewGetStocksOnPartnerWarehousesResponse instantiates a new GetStocksOnPartnerWarehousesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStocksOnPartnerWarehousesResponseWithDefaults

`func NewGetStocksOnPartnerWarehousesResponseWithDefaults() *GetStocksOnPartnerWarehousesResponse`

NewGetStocksOnPartnerWarehousesResponseWithDefaults instantiates a new GetStocksOnPartnerWarehousesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetStocksOnPartnerWarehousesResponse) GetStatus() ApiResponseStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetStocksOnPartnerWarehousesResponse) GetStatusOk() (*ApiResponseStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetStocksOnPartnerWarehousesResponse) SetStatus(v ApiResponseStatusType)`

SetStatus sets Status field to given value.


### GetResult

`func (o *GetStocksOnPartnerWarehousesResponse) GetResult() GetStocksOnPartnerWarehousesDTO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetStocksOnPartnerWarehousesResponse) GetResultOk() (*GetStocksOnPartnerWarehousesDTO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetStocksOnPartnerWarehousesResponse) SetResult(v GetStocksOnPartnerWarehousesDTO)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetStocksOnPartnerWarehousesResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


