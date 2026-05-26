# OrderDatesFilterDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDateFrom** | Pointer to **string** | Начальная дата для фильтрации заказов по дате оформления.  Формат даты: &#x60;ГГГГ-ММ-ДД&#x60;.  Между начальной и конечной датой (параметр &#x60;creationDateTo&#x60;) должно быть не больше 30 дней.  Значение по умолчанию: 30 дней назад от текущей даты.  Начальная дата включается в интервал для фильтрации.  | [optional] 
**CreationDateTo** | Pointer to **string** | Конечная дата для фильтрации заказов по дате оформления.  Формат даты: &#x60;ГГГГ-ММ-ДД&#x60;.  Между начальной (параметр &#x60;creationDateFrom&#x60;) и конечной датой должно быть не больше 30 дней.  Значение по умолчанию: текущая дата.  Если промежуток времени между &#x60;creationDateTo&#x60; и &#x60;creationDateFrom&#x60; меньше суток, то &#x60;creationDateTo&#x60; равен &#x60;creationDateFrom&#x60; + сутки.  Конечная дата не включается в интервал для фильтрации.  | [optional] 
**ShipmentDateFrom** | Pointer to **string** | Начальная дата для фильтрации заказов по дате отгрузки в службу доставки (параметр &#x60;shipmentDate&#x60;).  Формат даты: &#x60;ГГГГ-ММ-ДД&#x60;.  Между начальной и конечной датой (параметр &#x60;shipmentDateTo&#x60;) должно быть не больше 30 дней.  Начальная дата включается в интервал для фильтрации.  | [optional] 
**ShipmentDateTo** | Pointer to **string** | Конечная дата для фильтрации заказов по дате отгрузки в службу доставки (параметр &#x60;shipmentDate&#x60;).  Формат даты: &#x60;ГГГГ-ММ-ДД&#x60;.  Между начальной (параметр &#x60;shipmentDateFrom&#x60;) и конечной датой должно быть не больше 30 дней.  Если промежуток времени между &#x60;shipmentDateTo&#x60; и &#x60;shipmentDateFrom&#x60; меньше суток, то &#x60;shipmentDateTo&#x60; равен &#x60;shipmentDateFrom&#x60; + сутки.  Конечная дата не включается в интервал для фильтрации.  | [optional] 
**UpdateDateFrom** | Pointer to **time.Time** | Начальная дата обновления заказа (ISO 8601). | [optional] 
**UpdateDateTo** | Pointer to **time.Time** | Конечная дата обновления заказа (ISO 8601). | [optional] 

## Methods

### NewOrderDatesFilterDTO

`func NewOrderDatesFilterDTO() *OrderDatesFilterDTO`

NewOrderDatesFilterDTO instantiates a new OrderDatesFilterDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDatesFilterDTOWithDefaults

`func NewOrderDatesFilterDTOWithDefaults() *OrderDatesFilterDTO`

NewOrderDatesFilterDTOWithDefaults instantiates a new OrderDatesFilterDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDateFrom

`func (o *OrderDatesFilterDTO) GetCreationDateFrom() string`

GetCreationDateFrom returns the CreationDateFrom field if non-nil, zero value otherwise.

### GetCreationDateFromOk

`func (o *OrderDatesFilterDTO) GetCreationDateFromOk() (*string, bool)`

GetCreationDateFromOk returns a tuple with the CreationDateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateFrom

`func (o *OrderDatesFilterDTO) SetCreationDateFrom(v string)`

SetCreationDateFrom sets CreationDateFrom field to given value.

### HasCreationDateFrom

`func (o *OrderDatesFilterDTO) HasCreationDateFrom() bool`

HasCreationDateFrom returns a boolean if a field has been set.

### GetCreationDateTo

`func (o *OrderDatesFilterDTO) GetCreationDateTo() string`

GetCreationDateTo returns the CreationDateTo field if non-nil, zero value otherwise.

### GetCreationDateToOk

`func (o *OrderDatesFilterDTO) GetCreationDateToOk() (*string, bool)`

GetCreationDateToOk returns a tuple with the CreationDateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTo

`func (o *OrderDatesFilterDTO) SetCreationDateTo(v string)`

SetCreationDateTo sets CreationDateTo field to given value.

### HasCreationDateTo

`func (o *OrderDatesFilterDTO) HasCreationDateTo() bool`

HasCreationDateTo returns a boolean if a field has been set.

### GetShipmentDateFrom

`func (o *OrderDatesFilterDTO) GetShipmentDateFrom() string`

GetShipmentDateFrom returns the ShipmentDateFrom field if non-nil, zero value otherwise.

### GetShipmentDateFromOk

`func (o *OrderDatesFilterDTO) GetShipmentDateFromOk() (*string, bool)`

GetShipmentDateFromOk returns a tuple with the ShipmentDateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentDateFrom

`func (o *OrderDatesFilterDTO) SetShipmentDateFrom(v string)`

SetShipmentDateFrom sets ShipmentDateFrom field to given value.

### HasShipmentDateFrom

`func (o *OrderDatesFilterDTO) HasShipmentDateFrom() bool`

HasShipmentDateFrom returns a boolean if a field has been set.

### GetShipmentDateTo

`func (o *OrderDatesFilterDTO) GetShipmentDateTo() string`

GetShipmentDateTo returns the ShipmentDateTo field if non-nil, zero value otherwise.

### GetShipmentDateToOk

`func (o *OrderDatesFilterDTO) GetShipmentDateToOk() (*string, bool)`

GetShipmentDateToOk returns a tuple with the ShipmentDateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentDateTo

`func (o *OrderDatesFilterDTO) SetShipmentDateTo(v string)`

SetShipmentDateTo sets ShipmentDateTo field to given value.

### HasShipmentDateTo

`func (o *OrderDatesFilterDTO) HasShipmentDateTo() bool`

HasShipmentDateTo returns a boolean if a field has been set.

### GetUpdateDateFrom

`func (o *OrderDatesFilterDTO) GetUpdateDateFrom() time.Time`

GetUpdateDateFrom returns the UpdateDateFrom field if non-nil, zero value otherwise.

### GetUpdateDateFromOk

`func (o *OrderDatesFilterDTO) GetUpdateDateFromOk() (*time.Time, bool)`

GetUpdateDateFromOk returns a tuple with the UpdateDateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDateFrom

`func (o *OrderDatesFilterDTO) SetUpdateDateFrom(v time.Time)`

SetUpdateDateFrom sets UpdateDateFrom field to given value.

### HasUpdateDateFrom

`func (o *OrderDatesFilterDTO) HasUpdateDateFrom() bool`

HasUpdateDateFrom returns a boolean if a field has been set.

### GetUpdateDateTo

`func (o *OrderDatesFilterDTO) GetUpdateDateTo() time.Time`

GetUpdateDateTo returns the UpdateDateTo field if non-nil, zero value otherwise.

### GetUpdateDateToOk

`func (o *OrderDatesFilterDTO) GetUpdateDateToOk() (*time.Time, bool)`

GetUpdateDateToOk returns a tuple with the UpdateDateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDateTo

`func (o *OrderDatesFilterDTO) SetUpdateDateTo(v time.Time)`

SetUpdateDateTo sets UpdateDateTo field to given value.

### HasUpdateDateTo

`func (o *OrderDatesFilterDTO) HasUpdateDateTo() bool`

HasUpdateDateTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


