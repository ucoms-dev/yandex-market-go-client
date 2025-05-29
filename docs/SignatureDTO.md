# SignatureDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signed** | **bool** | Подписан ли акт приема-передачи. | 

## Methods

### NewSignatureDTO

`func NewSignatureDTO(signed bool, ) *SignatureDTO`

NewSignatureDTO instantiates a new SignatureDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignatureDTOWithDefaults

`func NewSignatureDTOWithDefaults() *SignatureDTO`

NewSignatureDTOWithDefaults instantiates a new SignatureDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigned

`func (o *SignatureDTO) GetSigned() bool`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *SignatureDTO) GetSignedOk() (*bool, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *SignatureDTO) SetSigned(v bool)`

SetSigned sets Signed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


