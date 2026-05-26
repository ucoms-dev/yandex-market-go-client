# BusinessOrderEacDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EacType** | [**OrderDeliveryEacType**](OrderDeliveryEacType.md) |  | 
**EacCode** | Pointer to **string** | Код подтверждения ЭАПП (для типа &#x60;MERCHANT_TO_COURIER&#x60;).  | [optional] 

## Methods

### NewBusinessOrderEacDTO

`func NewBusinessOrderEacDTO(eacType OrderDeliveryEacType, ) *BusinessOrderEacDTO`

NewBusinessOrderEacDTO instantiates a new BusinessOrderEacDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessOrderEacDTOWithDefaults

`func NewBusinessOrderEacDTOWithDefaults() *BusinessOrderEacDTO`

NewBusinessOrderEacDTOWithDefaults instantiates a new BusinessOrderEacDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEacType

`func (o *BusinessOrderEacDTO) GetEacType() OrderDeliveryEacType`

GetEacType returns the EacType field if non-nil, zero value otherwise.

### GetEacTypeOk

`func (o *BusinessOrderEacDTO) GetEacTypeOk() (*OrderDeliveryEacType, bool)`

GetEacTypeOk returns a tuple with the EacType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEacType

`func (o *BusinessOrderEacDTO) SetEacType(v OrderDeliveryEacType)`

SetEacType sets EacType field to given value.


### GetEacCode

`func (o *BusinessOrderEacDTO) GetEacCode() string`

GetEacCode returns the EacCode field if non-nil, zero value otherwise.

### GetEacCodeOk

`func (o *BusinessOrderEacDTO) GetEacCodeOk() (*string, bool)`

GetEacCodeOk returns a tuple with the EacCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEacCode

`func (o *BusinessOrderEacDTO) SetEacCode(v string)`

SetEacCode sets EacCode field to given value.

### HasEacCode

`func (o *BusinessOrderEacDTO) HasEacCode() bool`

HasEacCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


