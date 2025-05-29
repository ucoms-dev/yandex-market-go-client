# SupplyRequestDocumentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**SupplyRequestDocumentType**](SupplyRequestDocumentType.md) |  | 
**Url** | **string** | Ссылка на документ. | 
**CreatedAt** | **time.Time** | Дата и время создания документа. | 

## Methods

### NewSupplyRequestDocumentDTO

`func NewSupplyRequestDocumentDTO(type_ SupplyRequestDocumentType, url string, createdAt time.Time, ) *SupplyRequestDocumentDTO`

NewSupplyRequestDocumentDTO instantiates a new SupplyRequestDocumentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplyRequestDocumentDTOWithDefaults

`func NewSupplyRequestDocumentDTOWithDefaults() *SupplyRequestDocumentDTO`

NewSupplyRequestDocumentDTOWithDefaults instantiates a new SupplyRequestDocumentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SupplyRequestDocumentDTO) GetType() SupplyRequestDocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SupplyRequestDocumentDTO) GetTypeOk() (*SupplyRequestDocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SupplyRequestDocumentDTO) SetType(v SupplyRequestDocumentType)`

SetType sets Type field to given value.


### GetUrl

`func (o *SupplyRequestDocumentDTO) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SupplyRequestDocumentDTO) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SupplyRequestDocumentDTO) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCreatedAt

`func (o *SupplyRequestDocumentDTO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SupplyRequestDocumentDTO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SupplyRequestDocumentDTO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


