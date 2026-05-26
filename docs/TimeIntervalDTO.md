# TimeIntervalDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromTime** | **string** | Начало интервала.  Формат: &#x60;ЧЧ:ММ&#x60;.  | 
**ToTime** | **string** | Конец интервала.  Формат: &#x60;ЧЧ:ММ&#x60;.  | 

## Methods

### NewTimeIntervalDTO

`func NewTimeIntervalDTO(fromTime string, toTime string, ) *TimeIntervalDTO`

NewTimeIntervalDTO instantiates a new TimeIntervalDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeIntervalDTOWithDefaults

`func NewTimeIntervalDTOWithDefaults() *TimeIntervalDTO`

NewTimeIntervalDTOWithDefaults instantiates a new TimeIntervalDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromTime

`func (o *TimeIntervalDTO) GetFromTime() string`

GetFromTime returns the FromTime field if non-nil, zero value otherwise.

### GetFromTimeOk

`func (o *TimeIntervalDTO) GetFromTimeOk() (*string, bool)`

GetFromTimeOk returns a tuple with the FromTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTime

`func (o *TimeIntervalDTO) SetFromTime(v string)`

SetFromTime sets FromTime field to given value.


### GetToTime

`func (o *TimeIntervalDTO) GetToTime() string`

GetToTime returns the ToTime field if non-nil, zero value otherwise.

### GetToTimeOk

`func (o *TimeIntervalDTO) GetToTimeOk() (*string, bool)`

GetToTimeOk returns a tuple with the ToTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTime

`func (o *TimeIntervalDTO) SetToTime(v string)`

SetToTime sets ToTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


