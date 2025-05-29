# GetSuggestedOfferMappingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ApiResponseStatusType**](ApiResponseStatusType.md) |  | [optional] 
**Result** | Pointer to [**GetSuggestedOfferMappingsResultDTO**](GetSuggestedOfferMappingsResultDTO.md) |  | [optional] 

## Methods

### NewGetSuggestedOfferMappingsResponse

`func NewGetSuggestedOfferMappingsResponse() *GetSuggestedOfferMappingsResponse`

NewGetSuggestedOfferMappingsResponse instantiates a new GetSuggestedOfferMappingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSuggestedOfferMappingsResponseWithDefaults

`func NewGetSuggestedOfferMappingsResponseWithDefaults() *GetSuggestedOfferMappingsResponse`

NewGetSuggestedOfferMappingsResponseWithDefaults instantiates a new GetSuggestedOfferMappingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetSuggestedOfferMappingsResponse) GetStatus() ApiResponseStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSuggestedOfferMappingsResponse) GetStatusOk() (*ApiResponseStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSuggestedOfferMappingsResponse) SetStatus(v ApiResponseStatusType)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSuggestedOfferMappingsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *GetSuggestedOfferMappingsResponse) GetResult() GetSuggestedOfferMappingsResultDTO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetSuggestedOfferMappingsResponse) GetResultOk() (*GetSuggestedOfferMappingsResultDTO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetSuggestedOfferMappingsResponse) SetResult(v GetSuggestedOfferMappingsResultDTO)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetSuggestedOfferMappingsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


