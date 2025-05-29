# GetSupplyRequestsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestIds** | Pointer to **[]int64** | Идентификаторы заявок. | [optional] 
**RequestDateFrom** | Pointer to **NullableTime** | Дата начала периода для фильтрации заявок. | [optional] 
**RequestDateTo** | Pointer to **NullableTime** | Дата окончания периода для фильтрации заявок. | [optional] 
**RequestTypes** | Pointer to [**[]SupplyRequestType**](SupplyRequestType.md) | Типы заявок для фильтрации. | [optional] 
**RequestSubtypes** | Pointer to [**[]SupplyRequestSubType**](SupplyRequestSubType.md) | Подтипы заявок для фильтрации. | [optional] 
**RequestStatuses** | Pointer to [**[]SupplyRequestStatusType**](SupplyRequestStatusType.md) | Статусы заявок для фильтрации. | [optional] 
**Sorting** | Pointer to [**SupplyRequestSortingDTO**](SupplyRequestSortingDTO.md) |  | [optional] 

## Methods

### NewGetSupplyRequestsRequest

`func NewGetSupplyRequestsRequest() *GetSupplyRequestsRequest`

NewGetSupplyRequestsRequest instantiates a new GetSupplyRequestsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSupplyRequestsRequestWithDefaults

`func NewGetSupplyRequestsRequestWithDefaults() *GetSupplyRequestsRequest`

NewGetSupplyRequestsRequestWithDefaults instantiates a new GetSupplyRequestsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestIds

`func (o *GetSupplyRequestsRequest) GetRequestIds() []int64`

GetRequestIds returns the RequestIds field if non-nil, zero value otherwise.

### GetRequestIdsOk

`func (o *GetSupplyRequestsRequest) GetRequestIdsOk() (*[]int64, bool)`

GetRequestIdsOk returns a tuple with the RequestIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestIds

`func (o *GetSupplyRequestsRequest) SetRequestIds(v []int64)`

SetRequestIds sets RequestIds field to given value.

### HasRequestIds

`func (o *GetSupplyRequestsRequest) HasRequestIds() bool`

HasRequestIds returns a boolean if a field has been set.

### SetRequestIdsNil

`func (o *GetSupplyRequestsRequest) SetRequestIdsNil(b bool)`

 SetRequestIdsNil sets the value for RequestIds to be an explicit nil

### UnsetRequestIds
`func (o *GetSupplyRequestsRequest) UnsetRequestIds()`

UnsetRequestIds ensures that no value is present for RequestIds, not even an explicit nil
### GetRequestDateFrom

`func (o *GetSupplyRequestsRequest) GetRequestDateFrom() time.Time`

GetRequestDateFrom returns the RequestDateFrom field if non-nil, zero value otherwise.

### GetRequestDateFromOk

`func (o *GetSupplyRequestsRequest) GetRequestDateFromOk() (*time.Time, bool)`

GetRequestDateFromOk returns a tuple with the RequestDateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDateFrom

`func (o *GetSupplyRequestsRequest) SetRequestDateFrom(v time.Time)`

SetRequestDateFrom sets RequestDateFrom field to given value.

### HasRequestDateFrom

`func (o *GetSupplyRequestsRequest) HasRequestDateFrom() bool`

HasRequestDateFrom returns a boolean if a field has been set.

### SetRequestDateFromNil

`func (o *GetSupplyRequestsRequest) SetRequestDateFromNil(b bool)`

 SetRequestDateFromNil sets the value for RequestDateFrom to be an explicit nil

### UnsetRequestDateFrom
`func (o *GetSupplyRequestsRequest) UnsetRequestDateFrom()`

UnsetRequestDateFrom ensures that no value is present for RequestDateFrom, not even an explicit nil
### GetRequestDateTo

`func (o *GetSupplyRequestsRequest) GetRequestDateTo() time.Time`

GetRequestDateTo returns the RequestDateTo field if non-nil, zero value otherwise.

### GetRequestDateToOk

`func (o *GetSupplyRequestsRequest) GetRequestDateToOk() (*time.Time, bool)`

GetRequestDateToOk returns a tuple with the RequestDateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDateTo

`func (o *GetSupplyRequestsRequest) SetRequestDateTo(v time.Time)`

SetRequestDateTo sets RequestDateTo field to given value.

### HasRequestDateTo

`func (o *GetSupplyRequestsRequest) HasRequestDateTo() bool`

HasRequestDateTo returns a boolean if a field has been set.

### SetRequestDateToNil

`func (o *GetSupplyRequestsRequest) SetRequestDateToNil(b bool)`

 SetRequestDateToNil sets the value for RequestDateTo to be an explicit nil

### UnsetRequestDateTo
`func (o *GetSupplyRequestsRequest) UnsetRequestDateTo()`

UnsetRequestDateTo ensures that no value is present for RequestDateTo, not even an explicit nil
### GetRequestTypes

`func (o *GetSupplyRequestsRequest) GetRequestTypes() []SupplyRequestType`

GetRequestTypes returns the RequestTypes field if non-nil, zero value otherwise.

### GetRequestTypesOk

`func (o *GetSupplyRequestsRequest) GetRequestTypesOk() (*[]SupplyRequestType, bool)`

GetRequestTypesOk returns a tuple with the RequestTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTypes

`func (o *GetSupplyRequestsRequest) SetRequestTypes(v []SupplyRequestType)`

SetRequestTypes sets RequestTypes field to given value.

### HasRequestTypes

`func (o *GetSupplyRequestsRequest) HasRequestTypes() bool`

HasRequestTypes returns a boolean if a field has been set.

### SetRequestTypesNil

`func (o *GetSupplyRequestsRequest) SetRequestTypesNil(b bool)`

 SetRequestTypesNil sets the value for RequestTypes to be an explicit nil

### UnsetRequestTypes
`func (o *GetSupplyRequestsRequest) UnsetRequestTypes()`

UnsetRequestTypes ensures that no value is present for RequestTypes, not even an explicit nil
### GetRequestSubtypes

`func (o *GetSupplyRequestsRequest) GetRequestSubtypes() []SupplyRequestSubType`

GetRequestSubtypes returns the RequestSubtypes field if non-nil, zero value otherwise.

### GetRequestSubtypesOk

`func (o *GetSupplyRequestsRequest) GetRequestSubtypesOk() (*[]SupplyRequestSubType, bool)`

GetRequestSubtypesOk returns a tuple with the RequestSubtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSubtypes

`func (o *GetSupplyRequestsRequest) SetRequestSubtypes(v []SupplyRequestSubType)`

SetRequestSubtypes sets RequestSubtypes field to given value.

### HasRequestSubtypes

`func (o *GetSupplyRequestsRequest) HasRequestSubtypes() bool`

HasRequestSubtypes returns a boolean if a field has been set.

### SetRequestSubtypesNil

`func (o *GetSupplyRequestsRequest) SetRequestSubtypesNil(b bool)`

 SetRequestSubtypesNil sets the value for RequestSubtypes to be an explicit nil

### UnsetRequestSubtypes
`func (o *GetSupplyRequestsRequest) UnsetRequestSubtypes()`

UnsetRequestSubtypes ensures that no value is present for RequestSubtypes, not even an explicit nil
### GetRequestStatuses

`func (o *GetSupplyRequestsRequest) GetRequestStatuses() []SupplyRequestStatusType`

GetRequestStatuses returns the RequestStatuses field if non-nil, zero value otherwise.

### GetRequestStatusesOk

`func (o *GetSupplyRequestsRequest) GetRequestStatusesOk() (*[]SupplyRequestStatusType, bool)`

GetRequestStatusesOk returns a tuple with the RequestStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStatuses

`func (o *GetSupplyRequestsRequest) SetRequestStatuses(v []SupplyRequestStatusType)`

SetRequestStatuses sets RequestStatuses field to given value.

### HasRequestStatuses

`func (o *GetSupplyRequestsRequest) HasRequestStatuses() bool`

HasRequestStatuses returns a boolean if a field has been set.

### SetRequestStatusesNil

`func (o *GetSupplyRequestsRequest) SetRequestStatusesNil(b bool)`

 SetRequestStatusesNil sets the value for RequestStatuses to be an explicit nil

### UnsetRequestStatuses
`func (o *GetSupplyRequestsRequest) UnsetRequestStatuses()`

UnsetRequestStatuses ensures that no value is present for RequestStatuses, not even an explicit nil
### GetSorting

`func (o *GetSupplyRequestsRequest) GetSorting() SupplyRequestSortingDTO`

GetSorting returns the Sorting field if non-nil, zero value otherwise.

### GetSortingOk

`func (o *GetSupplyRequestsRequest) GetSortingOk() (*SupplyRequestSortingDTO, bool)`

GetSortingOk returns a tuple with the Sorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorting

`func (o *GetSupplyRequestsRequest) SetSorting(v SupplyRequestSortingDTO)`

SetSorting sets Sorting field to given value.

### HasSorting

`func (o *GetSupplyRequestsRequest) HasSorting() bool`

HasSorting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


