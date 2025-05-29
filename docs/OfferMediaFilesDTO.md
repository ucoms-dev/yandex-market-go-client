# OfferMediaFilesDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstVideoAsCover** | Pointer to **bool** | Использовать первое видео в карточке как видеообложку.  Передайте &#x60;true&#x60;, чтобы первое видео использовалось как видеообложка, или &#x60;false&#x60;, чтобы видеообложка не отображалась в карточке товара.  | [optional] 
**Videos** | Pointer to [**[]OfferMediaFileDTO**](OfferMediaFileDTO.md) | Видеофайлы товара.  | [optional] 
**Pictures** | Pointer to [**[]OfferMediaFileDTO**](OfferMediaFileDTO.md) | Изображения товара.  | [optional] 
**Manuals** | Pointer to [**[]OfferMediaFileDTO**](OfferMediaFileDTO.md) | Руководства по использованию товара.  | [optional] 

## Methods

### NewOfferMediaFilesDTO

`func NewOfferMediaFilesDTO() *OfferMediaFilesDTO`

NewOfferMediaFilesDTO instantiates a new OfferMediaFilesDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferMediaFilesDTOWithDefaults

`func NewOfferMediaFilesDTOWithDefaults() *OfferMediaFilesDTO`

NewOfferMediaFilesDTOWithDefaults instantiates a new OfferMediaFilesDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstVideoAsCover

`func (o *OfferMediaFilesDTO) GetFirstVideoAsCover() bool`

GetFirstVideoAsCover returns the FirstVideoAsCover field if non-nil, zero value otherwise.

### GetFirstVideoAsCoverOk

`func (o *OfferMediaFilesDTO) GetFirstVideoAsCoverOk() (*bool, bool)`

GetFirstVideoAsCoverOk returns a tuple with the FirstVideoAsCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstVideoAsCover

`func (o *OfferMediaFilesDTO) SetFirstVideoAsCover(v bool)`

SetFirstVideoAsCover sets FirstVideoAsCover field to given value.

### HasFirstVideoAsCover

`func (o *OfferMediaFilesDTO) HasFirstVideoAsCover() bool`

HasFirstVideoAsCover returns a boolean if a field has been set.

### GetVideos

`func (o *OfferMediaFilesDTO) GetVideos() []OfferMediaFileDTO`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *OfferMediaFilesDTO) GetVideosOk() (*[]OfferMediaFileDTO, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *OfferMediaFilesDTO) SetVideos(v []OfferMediaFileDTO)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *OfferMediaFilesDTO) HasVideos() bool`

HasVideos returns a boolean if a field has been set.

### SetVideosNil

`func (o *OfferMediaFilesDTO) SetVideosNil(b bool)`

 SetVideosNil sets the value for Videos to be an explicit nil

### UnsetVideos
`func (o *OfferMediaFilesDTO) UnsetVideos()`

UnsetVideos ensures that no value is present for Videos, not even an explicit nil
### GetPictures

`func (o *OfferMediaFilesDTO) GetPictures() []OfferMediaFileDTO`

GetPictures returns the Pictures field if non-nil, zero value otherwise.

### GetPicturesOk

`func (o *OfferMediaFilesDTO) GetPicturesOk() (*[]OfferMediaFileDTO, bool)`

GetPicturesOk returns a tuple with the Pictures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictures

`func (o *OfferMediaFilesDTO) SetPictures(v []OfferMediaFileDTO)`

SetPictures sets Pictures field to given value.

### HasPictures

`func (o *OfferMediaFilesDTO) HasPictures() bool`

HasPictures returns a boolean if a field has been set.

### SetPicturesNil

`func (o *OfferMediaFilesDTO) SetPicturesNil(b bool)`

 SetPicturesNil sets the value for Pictures to be an explicit nil

### UnsetPictures
`func (o *OfferMediaFilesDTO) UnsetPictures()`

UnsetPictures ensures that no value is present for Pictures, not even an explicit nil
### GetManuals

`func (o *OfferMediaFilesDTO) GetManuals() []OfferMediaFileDTO`

GetManuals returns the Manuals field if non-nil, zero value otherwise.

### GetManualsOk

`func (o *OfferMediaFilesDTO) GetManualsOk() (*[]OfferMediaFileDTO, bool)`

GetManualsOk returns a tuple with the Manuals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuals

`func (o *OfferMediaFilesDTO) SetManuals(v []OfferMediaFileDTO)`

SetManuals sets Manuals field to given value.

### HasManuals

`func (o *OfferMediaFilesDTO) HasManuals() bool`

HasManuals returns a boolean if a field has been set.

### SetManualsNil

`func (o *OfferMediaFilesDTO) SetManualsNil(b bool)`

 SetManualsNil sets the value for Manuals to be an explicit nil

### UnsetManuals
`func (o *OfferMediaFilesDTO) UnsetManuals()`

UnsetManuals ensures that no value is present for Manuals, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


