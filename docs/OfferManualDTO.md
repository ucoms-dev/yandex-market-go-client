# OfferManualDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Ссылка на инструкцию. | 
**Title** | Pointer to **string** | Название инструкции, которое будет отображаться на карточке товара.  | [optional] 

## Methods

### NewOfferManualDTO

`func NewOfferManualDTO(url string, ) *OfferManualDTO`

NewOfferManualDTO instantiates a new OfferManualDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferManualDTOWithDefaults

`func NewOfferManualDTOWithDefaults() *OfferManualDTO`

NewOfferManualDTOWithDefaults instantiates a new OfferManualDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *OfferManualDTO) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OfferManualDTO) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OfferManualDTO) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTitle

`func (o *OfferManualDTO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OfferManualDTO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OfferManualDTO) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OfferManualDTO) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


