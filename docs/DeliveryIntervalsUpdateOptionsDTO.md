# DeliveryIntervalsUpdateOptionsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableDeliveryIntervals** | [**[]DeliveryIntervalsUpdateOptionDTO**](DeliveryIntervalsUpdateOptionDTO.md) | Интервалы дат и времени.  Если доступных интервалов нет, возвращается пустой массив.  | 

## Methods

### NewDeliveryIntervalsUpdateOptionsDTO

`func NewDeliveryIntervalsUpdateOptionsDTO(availableDeliveryIntervals []DeliveryIntervalsUpdateOptionDTO, ) *DeliveryIntervalsUpdateOptionsDTO`

NewDeliveryIntervalsUpdateOptionsDTO instantiates a new DeliveryIntervalsUpdateOptionsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryIntervalsUpdateOptionsDTOWithDefaults

`func NewDeliveryIntervalsUpdateOptionsDTOWithDefaults() *DeliveryIntervalsUpdateOptionsDTO`

NewDeliveryIntervalsUpdateOptionsDTOWithDefaults instantiates a new DeliveryIntervalsUpdateOptionsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableDeliveryIntervals

`func (o *DeliveryIntervalsUpdateOptionsDTO) GetAvailableDeliveryIntervals() []DeliveryIntervalsUpdateOptionDTO`

GetAvailableDeliveryIntervals returns the AvailableDeliveryIntervals field if non-nil, zero value otherwise.

### GetAvailableDeliveryIntervalsOk

`func (o *DeliveryIntervalsUpdateOptionsDTO) GetAvailableDeliveryIntervalsOk() (*[]DeliveryIntervalsUpdateOptionDTO, bool)`

GetAvailableDeliveryIntervalsOk returns a tuple with the AvailableDeliveryIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableDeliveryIntervals

`func (o *DeliveryIntervalsUpdateOptionsDTO) SetAvailableDeliveryIntervals(v []DeliveryIntervalsUpdateOptionDTO)`

SetAvailableDeliveryIntervals sets AvailableDeliveryIntervals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


