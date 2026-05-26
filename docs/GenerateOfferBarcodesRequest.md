# GenerateOfferBarcodesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferIds** | **[]string** | Список товаров, для которых нужно сгенерировать штрихкоды. | 
**SkipIfExists** | Pointer to **bool** | Для каких товаров нужно сгенерировать штрихкоды:  * &#x60;false&#x60; — для всех, которые переданы в запросе. * &#x60;true&#x60; — только для тех, у которых их нет.  | [optional] [default to true]

## Methods

### NewGenerateOfferBarcodesRequest

`func NewGenerateOfferBarcodesRequest(offerIds []string, ) *GenerateOfferBarcodesRequest`

NewGenerateOfferBarcodesRequest instantiates a new GenerateOfferBarcodesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateOfferBarcodesRequestWithDefaults

`func NewGenerateOfferBarcodesRequestWithDefaults() *GenerateOfferBarcodesRequest`

NewGenerateOfferBarcodesRequestWithDefaults instantiates a new GenerateOfferBarcodesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferIds

`func (o *GenerateOfferBarcodesRequest) GetOfferIds() []string`

GetOfferIds returns the OfferIds field if non-nil, zero value otherwise.

### GetOfferIdsOk

`func (o *GenerateOfferBarcodesRequest) GetOfferIdsOk() (*[]string, bool)`

GetOfferIdsOk returns a tuple with the OfferIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferIds

`func (o *GenerateOfferBarcodesRequest) SetOfferIds(v []string)`

SetOfferIds sets OfferIds field to given value.


### GetSkipIfExists

`func (o *GenerateOfferBarcodesRequest) GetSkipIfExists() bool`

GetSkipIfExists returns the SkipIfExists field if non-nil, zero value otherwise.

### GetSkipIfExistsOk

`func (o *GenerateOfferBarcodesRequest) GetSkipIfExistsOk() (*bool, bool)`

GetSkipIfExistsOk returns a tuple with the SkipIfExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIfExists

`func (o *GenerateOfferBarcodesRequest) SetSkipIfExists(v bool)`

SetSkipIfExists sets SkipIfExists field to given value.

### HasSkipIfExists

`func (o *GenerateOfferBarcodesRequest) HasSkipIfExists() bool`

HasSkipIfExists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


