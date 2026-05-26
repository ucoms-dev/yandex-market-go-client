# GetBusinessOrdersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | [**[]BusinessOrderDTO**](BusinessOrderDTO.md) | Список заказов в кабинете. | 
**Paging** | Pointer to [**PackagingForwardScrollingPagerDTO**](PackagingForwardScrollingPagerDTO.md) |  | [optional] 

## Methods

### NewGetBusinessOrdersResponse

`func NewGetBusinessOrdersResponse(orders []BusinessOrderDTO, ) *GetBusinessOrdersResponse`

NewGetBusinessOrdersResponse instantiates a new GetBusinessOrdersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBusinessOrdersResponseWithDefaults

`func NewGetBusinessOrdersResponseWithDefaults() *GetBusinessOrdersResponse`

NewGetBusinessOrdersResponseWithDefaults instantiates a new GetBusinessOrdersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *GetBusinessOrdersResponse) GetOrders() []BusinessOrderDTO`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *GetBusinessOrdersResponse) GetOrdersOk() (*[]BusinessOrderDTO, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *GetBusinessOrdersResponse) SetOrders(v []BusinessOrderDTO)`

SetOrders sets Orders field to given value.


### GetPaging

`func (o *GetBusinessOrdersResponse) GetPaging() PackagingForwardScrollingPagerDTO`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *GetBusinessOrdersResponse) GetPagingOk() (*PackagingForwardScrollingPagerDTO, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *GetBusinessOrdersResponse) SetPaging(v PackagingForwardScrollingPagerDTO)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *GetBusinessOrdersResponse) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


