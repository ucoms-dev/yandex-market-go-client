# CisDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | Код маркировки в системе :no-translate[«Честный ЗНАК»]. | 
**Status** | [**CisStatusType**](CisStatusType.md) |  | 
**Substatus** | Pointer to [**CisSubstatusType**](CisSubstatusType.md) |  | [optional] 
**CrptRequestId** | Pointer to **string** | **Только для модели LaaS**  Идентификатор запроса проверки кода маркировки в [ЦРПТ](https://crpt.ru/), на основании которой принято решение о продаже товара.  | [optional] 
**CrptRequestDateTime** | Pointer to **time.Time** | **Только для модели LaaS**  Время проверки кода маркировки в [ЦРПТ](https://crpt.ru/), на основании которой принято решение о продаже товара.  | [optional] 

## Methods

### NewCisDTO

`func NewCisDTO(value string, status CisStatusType, ) *CisDTO`

NewCisDTO instantiates a new CisDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCisDTOWithDefaults

`func NewCisDTOWithDefaults() *CisDTO`

NewCisDTOWithDefaults instantiates a new CisDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CisDTO) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CisDTO) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CisDTO) SetValue(v string)`

SetValue sets Value field to given value.


### GetStatus

`func (o *CisDTO) GetStatus() CisStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CisDTO) GetStatusOk() (*CisStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CisDTO) SetStatus(v CisStatusType)`

SetStatus sets Status field to given value.


### GetSubstatus

`func (o *CisDTO) GetSubstatus() CisSubstatusType`

GetSubstatus returns the Substatus field if non-nil, zero value otherwise.

### GetSubstatusOk

`func (o *CisDTO) GetSubstatusOk() (*CisSubstatusType, bool)`

GetSubstatusOk returns a tuple with the Substatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstatus

`func (o *CisDTO) SetSubstatus(v CisSubstatusType)`

SetSubstatus sets Substatus field to given value.

### HasSubstatus

`func (o *CisDTO) HasSubstatus() bool`

HasSubstatus returns a boolean if a field has been set.

### GetCrptRequestId

`func (o *CisDTO) GetCrptRequestId() string`

GetCrptRequestId returns the CrptRequestId field if non-nil, zero value otherwise.

### GetCrptRequestIdOk

`func (o *CisDTO) GetCrptRequestIdOk() (*string, bool)`

GetCrptRequestIdOk returns a tuple with the CrptRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrptRequestId

`func (o *CisDTO) SetCrptRequestId(v string)`

SetCrptRequestId sets CrptRequestId field to given value.

### HasCrptRequestId

`func (o *CisDTO) HasCrptRequestId() bool`

HasCrptRequestId returns a boolean if a field has been set.

### GetCrptRequestDateTime

`func (o *CisDTO) GetCrptRequestDateTime() time.Time`

GetCrptRequestDateTime returns the CrptRequestDateTime field if non-nil, zero value otherwise.

### GetCrptRequestDateTimeOk

`func (o *CisDTO) GetCrptRequestDateTimeOk() (*time.Time, bool)`

GetCrptRequestDateTimeOk returns a tuple with the CrptRequestDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrptRequestDateTime

`func (o *CisDTO) SetCrptRequestDateTime(v time.Time)`

SetCrptRequestDateTime sets CrptRequestDateTime field to given value.

### HasCrptRequestDateTime

`func (o *CisDTO) HasCrptRequestDateTime() bool`

HasCrptRequestDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


