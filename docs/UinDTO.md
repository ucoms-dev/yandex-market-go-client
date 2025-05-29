# UinDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | УИН товара. | 
**Status** | [**UinStatusType**](UinStatusType.md) |  | 

## Methods

### NewUinDTO

`func NewUinDTO(value string, status UinStatusType, ) *UinDTO`

NewUinDTO instantiates a new UinDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUinDTOWithDefaults

`func NewUinDTOWithDefaults() *UinDTO`

NewUinDTOWithDefaults instantiates a new UinDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *UinDTO) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UinDTO) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UinDTO) SetValue(v string)`

SetValue sets Value field to given value.


### GetStatus

`func (o *UinDTO) GetStatus() UinStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UinDTO) GetStatusOk() (*UinStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UinDTO) SetStatus(v UinStatusType)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


