# BusinessOrderTransferDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Courier** | Pointer to [**OrderCourierDTO**](OrderCourierDTO.md) |  | [optional] 
**Eac** | Pointer to [**BusinessOrderEacDTO**](BusinessOrderEacDTO.md) |  | [optional] 

## Methods

### NewBusinessOrderTransferDTO

`func NewBusinessOrderTransferDTO() *BusinessOrderTransferDTO`

NewBusinessOrderTransferDTO instantiates a new BusinessOrderTransferDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessOrderTransferDTOWithDefaults

`func NewBusinessOrderTransferDTOWithDefaults() *BusinessOrderTransferDTO`

NewBusinessOrderTransferDTOWithDefaults instantiates a new BusinessOrderTransferDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCourier

`func (o *BusinessOrderTransferDTO) GetCourier() OrderCourierDTO`

GetCourier returns the Courier field if non-nil, zero value otherwise.

### GetCourierOk

`func (o *BusinessOrderTransferDTO) GetCourierOk() (*OrderCourierDTO, bool)`

GetCourierOk returns a tuple with the Courier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourier

`func (o *BusinessOrderTransferDTO) SetCourier(v OrderCourierDTO)`

SetCourier sets Courier field to given value.

### HasCourier

`func (o *BusinessOrderTransferDTO) HasCourier() bool`

HasCourier returns a boolean if a field has been set.

### GetEac

`func (o *BusinessOrderTransferDTO) GetEac() BusinessOrderEacDTO`

GetEac returns the Eac field if non-nil, zero value otherwise.

### GetEacOk

`func (o *BusinessOrderTransferDTO) GetEacOk() (*BusinessOrderEacDTO, bool)`

GetEacOk returns a tuple with the Eac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEac

`func (o *BusinessOrderTransferDTO) SetEac(v BusinessOrderEacDTO)`

SetEac sets Eac field to given value.

### HasEac

`func (o *BusinessOrderTransferDTO) HasEac() bool`

HasEac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


