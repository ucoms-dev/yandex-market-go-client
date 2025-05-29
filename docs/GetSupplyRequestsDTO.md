# GetSupplyRequestsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | [**[]SupplyRequestDTO**](SupplyRequestDTO.md) | Список заявок. | 
**Paging** | Pointer to [**ForwardScrollingPagerDTO**](ForwardScrollingPagerDTO.md) |  | [optional] 

## Methods

### NewGetSupplyRequestsDTO

`func NewGetSupplyRequestsDTO(requests []SupplyRequestDTO, ) *GetSupplyRequestsDTO`

NewGetSupplyRequestsDTO instantiates a new GetSupplyRequestsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSupplyRequestsDTOWithDefaults

`func NewGetSupplyRequestsDTOWithDefaults() *GetSupplyRequestsDTO`

NewGetSupplyRequestsDTOWithDefaults instantiates a new GetSupplyRequestsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *GetSupplyRequestsDTO) GetRequests() []SupplyRequestDTO`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *GetSupplyRequestsDTO) GetRequestsOk() (*[]SupplyRequestDTO, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *GetSupplyRequestsDTO) SetRequests(v []SupplyRequestDTO)`

SetRequests sets Requests field to given value.


### GetPaging

`func (o *GetSupplyRequestsDTO) GetPaging() ForwardScrollingPagerDTO`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *GetSupplyRequestsDTO) GetPagingOk() (*ForwardScrollingPagerDTO, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *GetSupplyRequestsDTO) SetPaging(v ForwardScrollingPagerDTO)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *GetSupplyRequestsDTO) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


