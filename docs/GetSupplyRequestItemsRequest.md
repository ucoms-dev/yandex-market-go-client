# GetSupplyRequestItemsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **int64** | Идентификатор заявки.  {% note warning \&quot;Используется только в API\&quot; %}  По нему не получится найти заявки в кабинете продавца на Маркете. Для этого используйте &#x60;marketplaceRequestId&#x60; или &#x60;warehouseRequestId&#x60;.  {% endnote %}  | 

## Methods

### NewGetSupplyRequestItemsRequest

`func NewGetSupplyRequestItemsRequest(requestId int64, ) *GetSupplyRequestItemsRequest`

NewGetSupplyRequestItemsRequest instantiates a new GetSupplyRequestItemsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSupplyRequestItemsRequestWithDefaults

`func NewGetSupplyRequestItemsRequestWithDefaults() *GetSupplyRequestItemsRequest`

NewGetSupplyRequestItemsRequestWithDefaults instantiates a new GetSupplyRequestItemsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetSupplyRequestItemsRequest) GetRequestId() int64`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetSupplyRequestItemsRequest) GetRequestIdOk() (*int64, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetSupplyRequestItemsRequest) SetRequestId(v int64)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


