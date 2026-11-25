# GetStocksOnPartnerWarehousesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerWarehouseId** | **int64** | Идентификатор склада.  Чтобы узнать идентификатор склада, воспользуйтесь запросом [POST v3/businesses/{businessId}/warehouses](../../reference/warehouses/getPartnerWarehouses.md).  | 
**Archived** | Pointer to **bool** | Фильтр по нахождению в архиве.  Передайте &#x60;true&#x60;, чтобы получить информацию об остатках товаров, которые находятся в архиве. Если фильтр не заполнен или передано &#x60;false&#x60;, в ответе возвращается информация о товарах, которые не находятся в архиве.  | [optional] 
**OfferIds** | Pointer to **[]string** | Фильтр по вашим SKU товаров.  Возвращается информация об остатках всех переданных SKU, включая товары в архиве.  {% note warning \&quot;Такой список возвращается только целиком\&quot; %}  Если вы запрашиваете информацию по конкретным SKU, не заполняйте:  * &#x60;pageToken&#x60; * &#x60;limit&#x60; * &#x60;archived&#x60;  {% endnote %}     | [optional] 

## Methods

### NewGetStocksOnPartnerWarehousesRequest

`func NewGetStocksOnPartnerWarehousesRequest(partnerWarehouseId int64, ) *GetStocksOnPartnerWarehousesRequest`

NewGetStocksOnPartnerWarehousesRequest instantiates a new GetStocksOnPartnerWarehousesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStocksOnPartnerWarehousesRequestWithDefaults

`func NewGetStocksOnPartnerWarehousesRequestWithDefaults() *GetStocksOnPartnerWarehousesRequest`

NewGetStocksOnPartnerWarehousesRequestWithDefaults instantiates a new GetStocksOnPartnerWarehousesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerWarehouseId

`func (o *GetStocksOnPartnerWarehousesRequest) GetPartnerWarehouseId() int64`

GetPartnerWarehouseId returns the PartnerWarehouseId field if non-nil, zero value otherwise.

### GetPartnerWarehouseIdOk

`func (o *GetStocksOnPartnerWarehousesRequest) GetPartnerWarehouseIdOk() (*int64, bool)`

GetPartnerWarehouseIdOk returns a tuple with the PartnerWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerWarehouseId

`func (o *GetStocksOnPartnerWarehousesRequest) SetPartnerWarehouseId(v int64)`

SetPartnerWarehouseId sets PartnerWarehouseId field to given value.


### GetArchived

`func (o *GetStocksOnPartnerWarehousesRequest) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *GetStocksOnPartnerWarehousesRequest) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *GetStocksOnPartnerWarehousesRequest) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *GetStocksOnPartnerWarehousesRequest) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetOfferIds

`func (o *GetStocksOnPartnerWarehousesRequest) GetOfferIds() []string`

GetOfferIds returns the OfferIds field if non-nil, zero value otherwise.

### GetOfferIdsOk

`func (o *GetStocksOnPartnerWarehousesRequest) GetOfferIdsOk() (*[]string, bool)`

GetOfferIdsOk returns a tuple with the OfferIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferIds

`func (o *GetStocksOnPartnerWarehousesRequest) SetOfferIds(v []string)`

SetOfferIds sets OfferIds field to given value.

### HasOfferIds

`func (o *GetStocksOnPartnerWarehousesRequest) HasOfferIds() bool`

HasOfferIds returns a boolean if a field has been set.

### SetOfferIdsNil

`func (o *GetStocksOnPartnerWarehousesRequest) SetOfferIdsNil(b bool)`

 SetOfferIdsNil sets the value for OfferIds to be an explicit nil

### UnsetOfferIds
`func (o *GetStocksOnPartnerWarehousesRequest) UnsetOfferIds()`

UnsetOfferIds ensures that no value is present for OfferIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


