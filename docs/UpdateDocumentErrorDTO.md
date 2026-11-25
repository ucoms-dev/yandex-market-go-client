# UpdateDocumentErrorDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Идентификатор документа. | 
**Code** | [**DocumentErrorCodeType**](DocumentErrorCodeType.md) |  | 
**Message** | Pointer to **string** | Описание ошибки для человека. Для обработки ошибки используйте поле &#x60;code&#x60;. | [optional] 

## Methods

### NewUpdateDocumentErrorDTO

`func NewUpdateDocumentErrorDTO(id int64, code DocumentErrorCodeType, ) *UpdateDocumentErrorDTO`

NewUpdateDocumentErrorDTO instantiates a new UpdateDocumentErrorDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDocumentErrorDTOWithDefaults

`func NewUpdateDocumentErrorDTOWithDefaults() *UpdateDocumentErrorDTO`

NewUpdateDocumentErrorDTOWithDefaults instantiates a new UpdateDocumentErrorDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateDocumentErrorDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateDocumentErrorDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateDocumentErrorDTO) SetId(v int64)`

SetId sets Id field to given value.


### GetCode

`func (o *UpdateDocumentErrorDTO) GetCode() DocumentErrorCodeType`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateDocumentErrorDTO) GetCodeOk() (*DocumentErrorCodeType, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateDocumentErrorDTO) SetCode(v DocumentErrorCodeType)`

SetCode sets Code field to given value.


### GetMessage

`func (o *UpdateDocumentErrorDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateDocumentErrorDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateDocumentErrorDTO) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdateDocumentErrorDTO) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


