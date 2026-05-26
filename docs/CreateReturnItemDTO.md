# CreateReturnItemDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferId** | **string** | Ваш SKU — идентификатор товара в вашей системе.  Правила использования SKU:  * У каждого товара SKU должен быть свой.  * Уже заданный SKU нельзя освободить и использовать заново для другого товара. Каждый товар должен получать новый идентификатор, до того никогда не использовавшийся в вашем каталоге.  SKU товара можно изменить в кабинете продавца на Маркете. О том, как это сделать, читайте [в Справке Маркета для продавцов](https://yandex.ru/support2/marketplace/ru/assortment/operations/edit-sku).  {% note warning %}  Пробельные символы в начале и конце значения автоматически удаляются. Например, &#x60;\&quot;  SKU123  \&quot;&#x60; и &#x60;\&quot;SKU123\&quot;&#x60; будут обработаны как одинаковые значения.  {% endnote %}  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields)  | 
**Count** | **int32** | Количество единиц товара. | 
**ReasonType** | [**ExternalReturnDecisionReasonType**](ExternalReturnDecisionReasonType.md) |  | 
**SubreasonType** | [**ExternalReturnDecisionSubreasonType**](ExternalReturnDecisionSubreasonType.md) |  | 
**Comment** | Pointer to **string** | Комментарий к товару в возврате. | [optional] 
**Pictures** | Pointer to **[]string** | Ссылки (URL) на изображения товара в возврате. | [optional] 

## Methods

### NewCreateReturnItemDTO

`func NewCreateReturnItemDTO(offerId string, count int32, reasonType ExternalReturnDecisionReasonType, subreasonType ExternalReturnDecisionSubreasonType, ) *CreateReturnItemDTO`

NewCreateReturnItemDTO instantiates a new CreateReturnItemDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReturnItemDTOWithDefaults

`func NewCreateReturnItemDTOWithDefaults() *CreateReturnItemDTO`

NewCreateReturnItemDTOWithDefaults instantiates a new CreateReturnItemDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferId

`func (o *CreateReturnItemDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *CreateReturnItemDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *CreateReturnItemDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.


### GetCount

`func (o *CreateReturnItemDTO) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CreateReturnItemDTO) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CreateReturnItemDTO) SetCount(v int32)`

SetCount sets Count field to given value.


### GetReasonType

`func (o *CreateReturnItemDTO) GetReasonType() ExternalReturnDecisionReasonType`

GetReasonType returns the ReasonType field if non-nil, zero value otherwise.

### GetReasonTypeOk

`func (o *CreateReturnItemDTO) GetReasonTypeOk() (*ExternalReturnDecisionReasonType, bool)`

GetReasonTypeOk returns a tuple with the ReasonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonType

`func (o *CreateReturnItemDTO) SetReasonType(v ExternalReturnDecisionReasonType)`

SetReasonType sets ReasonType field to given value.


### GetSubreasonType

`func (o *CreateReturnItemDTO) GetSubreasonType() ExternalReturnDecisionSubreasonType`

GetSubreasonType returns the SubreasonType field if non-nil, zero value otherwise.

### GetSubreasonTypeOk

`func (o *CreateReturnItemDTO) GetSubreasonTypeOk() (*ExternalReturnDecisionSubreasonType, bool)`

GetSubreasonTypeOk returns a tuple with the SubreasonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubreasonType

`func (o *CreateReturnItemDTO) SetSubreasonType(v ExternalReturnDecisionSubreasonType)`

SetSubreasonType sets SubreasonType field to given value.


### GetComment

`func (o *CreateReturnItemDTO) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateReturnItemDTO) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateReturnItemDTO) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateReturnItemDTO) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPictures

`func (o *CreateReturnItemDTO) GetPictures() []string`

GetPictures returns the Pictures field if non-nil, zero value otherwise.

### GetPicturesOk

`func (o *CreateReturnItemDTO) GetPicturesOk() (*[]string, bool)`

GetPicturesOk returns a tuple with the Pictures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictures

`func (o *CreateReturnItemDTO) SetPictures(v []string)`

SetPictures sets Pictures field to given value.

### HasPictures

`func (o *CreateReturnItemDTO) HasPictures() bool`

HasPictures returns a boolean if a field has been set.

### SetPicturesNil

`func (o *CreateReturnItemDTO) SetPicturesNil(b bool)`

 SetPicturesNil sets the value for Pictures to be an explicit nil

### UnsetPictures
`func (o *CreateReturnItemDTO) UnsetPictures()`

UnsetPictures ensures that no value is present for Pictures, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


