# OfferMediaFileDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | Ссылка на медиафайл.  | [optional] 
**Title** | Pointer to **string** | Название медиафайла.  | [optional] 
**UploadState** | Pointer to [**MediaFileUploadStateType**](MediaFileUploadStateType.md) |  | [optional] 

## Methods

### NewOfferMediaFileDTO

`func NewOfferMediaFileDTO() *OfferMediaFileDTO`

NewOfferMediaFileDTO instantiates a new OfferMediaFileDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferMediaFileDTOWithDefaults

`func NewOfferMediaFileDTOWithDefaults() *OfferMediaFileDTO`

NewOfferMediaFileDTOWithDefaults instantiates a new OfferMediaFileDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *OfferMediaFileDTO) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OfferMediaFileDTO) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OfferMediaFileDTO) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OfferMediaFileDTO) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTitle

`func (o *OfferMediaFileDTO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OfferMediaFileDTO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OfferMediaFileDTO) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OfferMediaFileDTO) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUploadState

`func (o *OfferMediaFileDTO) GetUploadState() MediaFileUploadStateType`

GetUploadState returns the UploadState field if non-nil, zero value otherwise.

### GetUploadStateOk

`func (o *OfferMediaFileDTO) GetUploadStateOk() (*MediaFileUploadStateType, bool)`

GetUploadStateOk returns a tuple with the UploadState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadState

`func (o *OfferMediaFileDTO) SetUploadState(v MediaFileUploadStateType)`

SetUploadState sets UploadState field to given value.

### HasUploadState

`func (o *OfferMediaFileDTO) HasUploadState() bool`

HasUploadState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


