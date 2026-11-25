# UpdateStockOnWarehouseItemDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | **string** | Ваш SKU — идентификатор товара в вашей системе.  Правила использования SKU:  * У каждого товара SKU должен быть свой.  * Уже заданный SKU нельзя освободить и использовать заново для другого товара. Каждый товар должен получать новый идентификатор, до того никогда не использовавшийся в вашем каталоге.  SKU товара можно изменить в кабинете продавца на Маркете. О том, как это сделать, читайте [в Справке Маркета для продавцов](https://yandex.ru/support2/marketplace/ru/assortment/operations/edit-sku).  {% note warning %}  Пробельные символы в начале и конце значения автоматически удаляются. Например, &#x60;\&quot;  SKU123  \&quot;&#x60; и &#x60;\&quot;SKU123\&quot;&#x60; будут обработаны как одинаковые значения.  {% endnote %}  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields)  | 
**PartnerWarehouseId** | **int64** | Идентификатор склада, на котором требуется обновить остатки.  Чтобы узнать идентификатор склада, воспользуйтесь запросом [POST v3/businesses/{businessId}/warehouses](../../reference/warehouses/getPartnerWarehouses.md).  | 
**Count** | **int64** | Количество доступного товара.  | 
**UpdatedAt** | Pointer to **time.Time** | Дата и время последнего обновления информации об остатках. &lt;br&gt;&lt;br&gt; Если вы не передали параметр &#x60;updatedAt&#x60;, используется текущее время. &lt;br&gt;&lt;br&gt; Формат даты и времени: ISO 8601 со смещением относительно UTC. Например, &#x60;2017-11-21T00:42:42+03:00&#x60;.  | [optional] 

## Methods

### NewUpdateStockOnWarehouseItemDTO

`func NewUpdateStockOnWarehouseItemDTO(sku string, partnerWarehouseId int64, count int64, ) *UpdateStockOnWarehouseItemDTO`

NewUpdateStockOnWarehouseItemDTO instantiates a new UpdateStockOnWarehouseItemDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStockOnWarehouseItemDTOWithDefaults

`func NewUpdateStockOnWarehouseItemDTOWithDefaults() *UpdateStockOnWarehouseItemDTO`

NewUpdateStockOnWarehouseItemDTOWithDefaults instantiates a new UpdateStockOnWarehouseItemDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *UpdateStockOnWarehouseItemDTO) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *UpdateStockOnWarehouseItemDTO) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *UpdateStockOnWarehouseItemDTO) SetSku(v string)`

SetSku sets Sku field to given value.


### GetPartnerWarehouseId

`func (o *UpdateStockOnWarehouseItemDTO) GetPartnerWarehouseId() int64`

GetPartnerWarehouseId returns the PartnerWarehouseId field if non-nil, zero value otherwise.

### GetPartnerWarehouseIdOk

`func (o *UpdateStockOnWarehouseItemDTO) GetPartnerWarehouseIdOk() (*int64, bool)`

GetPartnerWarehouseIdOk returns a tuple with the PartnerWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerWarehouseId

`func (o *UpdateStockOnWarehouseItemDTO) SetPartnerWarehouseId(v int64)`

SetPartnerWarehouseId sets PartnerWarehouseId field to given value.


### GetCount

`func (o *UpdateStockOnWarehouseItemDTO) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UpdateStockOnWarehouseItemDTO) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UpdateStockOnWarehouseItemDTO) SetCount(v int64)`

SetCount sets Count field to given value.


### GetUpdatedAt

`func (o *UpdateStockOnWarehouseItemDTO) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdateStockOnWarehouseItemDTO) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdateStockOnWarehouseItemDTO) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UpdateStockOnWarehouseItemDTO) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


