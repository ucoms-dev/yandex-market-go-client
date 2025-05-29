# TokenDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | [**ApiKeyDTO**](ApiKeyDTO.md) |  | 

## Methods

### NewTokenDTO

`func NewTokenDTO(apiKey ApiKeyDTO, ) *TokenDTO`

NewTokenDTO instantiates a new TokenDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenDTOWithDefaults

`func NewTokenDTOWithDefaults() *TokenDTO`

NewTokenDTOWithDefaults instantiates a new TokenDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *TokenDTO) GetApiKey() ApiKeyDTO`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *TokenDTO) GetApiKeyOk() (*ApiKeyDTO, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *TokenDTO) SetApiKey(v ApiKeyDTO)`

SetApiKey sets ApiKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


