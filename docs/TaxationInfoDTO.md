# TaxationInfoDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vat** | Pointer to [**VatRateType**](VatRateType.md) |  | [optional] 

## Methods

### NewTaxationInfoDTO

`func NewTaxationInfoDTO() *TaxationInfoDTO`

NewTaxationInfoDTO instantiates a new TaxationInfoDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxationInfoDTOWithDefaults

`func NewTaxationInfoDTOWithDefaults() *TaxationInfoDTO`

NewTaxationInfoDTOWithDefaults instantiates a new TaxationInfoDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVat

`func (o *TaxationInfoDTO) GetVat() VatRateType`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *TaxationInfoDTO) GetVatOk() (*VatRateType, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *TaxationInfoDTO) SetVat(v VatRateType)`

SetVat sets Vat field to given value.

### HasVat

`func (o *TaxationInfoDTO) HasVat() bool`

HasVat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


