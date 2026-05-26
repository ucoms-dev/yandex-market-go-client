# LogisticPointScheduleDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | [**[]ScheduleDayDTO**](ScheduleDayDTO.md) | График работы. | 
**Holidays** | Pointer to **[]string** | Расписание праздничных дней. | [optional] 

## Methods

### NewLogisticPointScheduleDTO

`func NewLogisticPointScheduleDTO(schedule []ScheduleDayDTO, ) *LogisticPointScheduleDTO`

NewLogisticPointScheduleDTO instantiates a new LogisticPointScheduleDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogisticPointScheduleDTOWithDefaults

`func NewLogisticPointScheduleDTOWithDefaults() *LogisticPointScheduleDTO`

NewLogisticPointScheduleDTOWithDefaults instantiates a new LogisticPointScheduleDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *LogisticPointScheduleDTO) GetSchedule() []ScheduleDayDTO`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *LogisticPointScheduleDTO) GetScheduleOk() (*[]ScheduleDayDTO, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *LogisticPointScheduleDTO) SetSchedule(v []ScheduleDayDTO)`

SetSchedule sets Schedule field to given value.


### GetHolidays

`func (o *LogisticPointScheduleDTO) GetHolidays() []string`

GetHolidays returns the Holidays field if non-nil, zero value otherwise.

### GetHolidaysOk

`func (o *LogisticPointScheduleDTO) GetHolidaysOk() (*[]string, bool)`

GetHolidaysOk returns a tuple with the Holidays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidays

`func (o *LogisticPointScheduleDTO) SetHolidays(v []string)`

SetHolidays sets Holidays field to given value.

### HasHolidays

`func (o *LogisticPointScheduleDTO) HasHolidays() bool`

HasHolidays returns a boolean if a field has been set.

### SetHolidaysNil

`func (o *LogisticPointScheduleDTO) SetHolidaysNil(b bool)`

 SetHolidaysNil sets the value for Holidays to be an explicit nil

### UnsetHolidays
`func (o *LogisticPointScheduleDTO) UnsetHolidays()`

UnsetHolidays ensures that no value is present for Holidays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


