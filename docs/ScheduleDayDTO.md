# ScheduleDayDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | [**DayOfWeekType**](DayOfWeekType.md) |  | 
**StartTime** | **string** | Время начала рабочего дня.  Формат: &#x60;ЧЧ:ММ&#x60;.  | 
**EndTime** | **string** | Время конца рабочего дня.  Формат: &#x60;ЧЧ:ММ&#x60;.  | 

## Methods

### NewScheduleDayDTO

`func NewScheduleDayDTO(day DayOfWeekType, startTime string, endTime string, ) *ScheduleDayDTO`

NewScheduleDayDTO instantiates a new ScheduleDayDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleDayDTOWithDefaults

`func NewScheduleDayDTOWithDefaults() *ScheduleDayDTO`

NewScheduleDayDTOWithDefaults instantiates a new ScheduleDayDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *ScheduleDayDTO) GetDay() DayOfWeekType`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *ScheduleDayDTO) GetDayOk() (*DayOfWeekType, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *ScheduleDayDTO) SetDay(v DayOfWeekType)`

SetDay sets Day field to given value.


### GetStartTime

`func (o *ScheduleDayDTO) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ScheduleDayDTO) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ScheduleDayDTO) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *ScheduleDayDTO) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ScheduleDayDTO) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ScheduleDayDTO) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


