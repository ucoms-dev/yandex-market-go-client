# CourierDeliveryParametersDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullAddress** | **string** | Полный адрес с точностью до номера дома. | 

## Methods

### NewCourierDeliveryParametersDTO

`func NewCourierDeliveryParametersDTO(fullAddress string, ) *CourierDeliveryParametersDTO`

NewCourierDeliveryParametersDTO instantiates a new CourierDeliveryParametersDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCourierDeliveryParametersDTOWithDefaults

`func NewCourierDeliveryParametersDTOWithDefaults() *CourierDeliveryParametersDTO`

NewCourierDeliveryParametersDTOWithDefaults instantiates a new CourierDeliveryParametersDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullAddress

`func (o *CourierDeliveryParametersDTO) GetFullAddress() string`

GetFullAddress returns the FullAddress field if non-nil, zero value otherwise.

### GetFullAddressOk

`func (o *CourierDeliveryParametersDTO) GetFullAddressOk() (*string, bool)`

GetFullAddressOk returns a tuple with the FullAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullAddress

`func (o *CourierDeliveryParametersDTO) SetFullAddress(v string)`

SetFullAddress sets FullAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


