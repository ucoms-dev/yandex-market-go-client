# BusinessOrderDeliveryDatesDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromDate** | **string** | Ближайшая дата доставки.  Формат даты: &#x60;ГГГГ-ММ-ДД&#x60;.  | 
**ToDate** | Pointer to **string** | Самая поздняя дата доставки.  Если &#x60;toDate&#x60; не указан, считается дата в параметре &#x60;fromDate&#x60;.  Формат даты: &#x60;ГГГГ-ММ-ДД&#x60;.  | [optional] 
**FromTime** | Pointer to **string** | Начало интервала времени доставки.  Передается только вместе с параметром &#x60;type&#x3D;DELIVERY&#x60;.  Формат времени: 24-часовой, &#x60;ЧЧ:ММ&#x60;. Вместо &#x60;ММ&#x60; всегда указывайте &#x60;00&#x60; (исключение — &#x60;23:59&#x60;).  Минимальное значение: &#x60;00:00&#x60;.  | [optional] 
**ToTime** | Pointer to **string** | Конец интервала времени доставки.  Передается только вместе с параметром &#x60;type&#x3D;DELIVERY&#x60;.  Формат времени: 24-часовой, &#x60;ЧЧ:ММ&#x60;. Вместо &#x60;ММ&#x60; всегда указывайте &#x60;00&#x60; (исключение — &#x60;23:59&#x60;).  Максимальное значение: &#x60;23:59&#x60;.  | [optional] 
**RealDeliveryDate** | Pointer to **string** | Дата, когда товар доставлен до пункта выдачи заказа (в случае самовывоза) или до покупателя (если заказ доставляет курьер).  Формат даты: &#x60;ГГГГ-ММ-ДД&#x60;.  | [optional] 

## Methods

### NewBusinessOrderDeliveryDatesDTO

`func NewBusinessOrderDeliveryDatesDTO(fromDate string, ) *BusinessOrderDeliveryDatesDTO`

NewBusinessOrderDeliveryDatesDTO instantiates a new BusinessOrderDeliveryDatesDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessOrderDeliveryDatesDTOWithDefaults

`func NewBusinessOrderDeliveryDatesDTOWithDefaults() *BusinessOrderDeliveryDatesDTO`

NewBusinessOrderDeliveryDatesDTOWithDefaults instantiates a new BusinessOrderDeliveryDatesDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromDate

`func (o *BusinessOrderDeliveryDatesDTO) GetFromDate() string`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *BusinessOrderDeliveryDatesDTO) GetFromDateOk() (*string, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *BusinessOrderDeliveryDatesDTO) SetFromDate(v string)`

SetFromDate sets FromDate field to given value.


### GetToDate

`func (o *BusinessOrderDeliveryDatesDTO) GetToDate() string`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *BusinessOrderDeliveryDatesDTO) GetToDateOk() (*string, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *BusinessOrderDeliveryDatesDTO) SetToDate(v string)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *BusinessOrderDeliveryDatesDTO) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### GetFromTime

`func (o *BusinessOrderDeliveryDatesDTO) GetFromTime() string`

GetFromTime returns the FromTime field if non-nil, zero value otherwise.

### GetFromTimeOk

`func (o *BusinessOrderDeliveryDatesDTO) GetFromTimeOk() (*string, bool)`

GetFromTimeOk returns a tuple with the FromTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTime

`func (o *BusinessOrderDeliveryDatesDTO) SetFromTime(v string)`

SetFromTime sets FromTime field to given value.

### HasFromTime

`func (o *BusinessOrderDeliveryDatesDTO) HasFromTime() bool`

HasFromTime returns a boolean if a field has been set.

### GetToTime

`func (o *BusinessOrderDeliveryDatesDTO) GetToTime() string`

GetToTime returns the ToTime field if non-nil, zero value otherwise.

### GetToTimeOk

`func (o *BusinessOrderDeliveryDatesDTO) GetToTimeOk() (*string, bool)`

GetToTimeOk returns a tuple with the ToTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTime

`func (o *BusinessOrderDeliveryDatesDTO) SetToTime(v string)`

SetToTime sets ToTime field to given value.

### HasToTime

`func (o *BusinessOrderDeliveryDatesDTO) HasToTime() bool`

HasToTime returns a boolean if a field has been set.

### GetRealDeliveryDate

`func (o *BusinessOrderDeliveryDatesDTO) GetRealDeliveryDate() string`

GetRealDeliveryDate returns the RealDeliveryDate field if non-nil, zero value otherwise.

### GetRealDeliveryDateOk

`func (o *BusinessOrderDeliveryDatesDTO) GetRealDeliveryDateOk() (*string, bool)`

GetRealDeliveryDateOk returns a tuple with the RealDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealDeliveryDate

`func (o *BusinessOrderDeliveryDatesDTO) SetRealDeliveryDate(v string)`

SetRealDeliveryDate sets RealDeliveryDate field to given value.

### HasRealDeliveryDate

`func (o *BusinessOrderDeliveryDatesDTO) HasRealDeliveryDate() bool`

HasRealDeliveryDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


