# ApiKeyDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Название токена. | 
**AuthScopes** | [**[]ApiKeyScopeType**](ApiKeyScopeType.md) | Доступы к методам по Api-Key-токену. | 

## Methods

### NewApiKeyDTO

`func NewApiKeyDTO(name string, authScopes []ApiKeyScopeType, ) *ApiKeyDTO`

NewApiKeyDTO instantiates a new ApiKeyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyDTOWithDefaults

`func NewApiKeyDTOWithDefaults() *ApiKeyDTO`

NewApiKeyDTOWithDefaults instantiates a new ApiKeyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiKeyDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKeyDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKeyDTO) SetName(v string)`

SetName sets Name field to given value.


### GetAuthScopes

`func (o *ApiKeyDTO) GetAuthScopes() []ApiKeyScopeType`

GetAuthScopes returns the AuthScopes field if non-nil, zero value otherwise.

### GetAuthScopesOk

`func (o *ApiKeyDTO) GetAuthScopesOk() (*[]ApiKeyScopeType, bool)`

GetAuthScopesOk returns a tuple with the AuthScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScopes

`func (o *ApiKeyDTO) SetAuthScopes(v []ApiKeyScopeType)`

SetAuthScopes sets AuthScopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


