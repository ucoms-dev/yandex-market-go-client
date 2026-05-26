# DeliveryDateIntervalDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromDate** | **string** | Начало интервала.  Формат даты: &#x60;ГГГГ-ММ-ДД&#x60;.  | 
**ToDate** | **string** | Конец интервала.  Формат даты: &#x60;ГГГГ-ММ-ДД&#x60;.  | 

## Methods

### NewDeliveryDateIntervalDTO

`func NewDeliveryDateIntervalDTO(fromDate string, toDate string, ) *DeliveryDateIntervalDTO`

NewDeliveryDateIntervalDTO instantiates a new DeliveryDateIntervalDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryDateIntervalDTOWithDefaults

`func NewDeliveryDateIntervalDTOWithDefaults() *DeliveryDateIntervalDTO`

NewDeliveryDateIntervalDTOWithDefaults instantiates a new DeliveryDateIntervalDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromDate

`func (o *DeliveryDateIntervalDTO) GetFromDate() string`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *DeliveryDateIntervalDTO) GetFromDateOk() (*string, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *DeliveryDateIntervalDTO) SetFromDate(v string)`

SetFromDate sets FromDate field to given value.


### GetToDate

`func (o *DeliveryDateIntervalDTO) GetToDate() string`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *DeliveryDateIntervalDTO) GetToDateOk() (*string, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *DeliveryDateIntervalDTO) SetToDate(v string)`

SetToDate sets ToDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


