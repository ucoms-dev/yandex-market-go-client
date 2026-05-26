# GenerateOfferBarcodesResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnprocessedOfferIds** | **[]string** | Список товаров, для которых не удалось сгенерировать штрихкоды. | 

## Methods

### NewGenerateOfferBarcodesResultDTO

`func NewGenerateOfferBarcodesResultDTO(unprocessedOfferIds []string, ) *GenerateOfferBarcodesResultDTO`

NewGenerateOfferBarcodesResultDTO instantiates a new GenerateOfferBarcodesResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateOfferBarcodesResultDTOWithDefaults

`func NewGenerateOfferBarcodesResultDTOWithDefaults() *GenerateOfferBarcodesResultDTO`

NewGenerateOfferBarcodesResultDTOWithDefaults instantiates a new GenerateOfferBarcodesResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnprocessedOfferIds

`func (o *GenerateOfferBarcodesResultDTO) GetUnprocessedOfferIds() []string`

GetUnprocessedOfferIds returns the UnprocessedOfferIds field if non-nil, zero value otherwise.

### GetUnprocessedOfferIdsOk

`func (o *GenerateOfferBarcodesResultDTO) GetUnprocessedOfferIdsOk() (*[]string, bool)`

GetUnprocessedOfferIdsOk returns a tuple with the UnprocessedOfferIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprocessedOfferIds

`func (o *GenerateOfferBarcodesResultDTO) SetUnprocessedOfferIds(v []string)`

SetUnprocessedOfferIds sets UnprocessedOfferIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


