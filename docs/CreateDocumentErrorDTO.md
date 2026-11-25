# CreateDocumentErrorDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | Номер, указанный в сертификате, декларации или другом документе. | 
**Code** | [**DocumentErrorCodeType**](DocumentErrorCodeType.md) |  | 
**Message** | Pointer to **string** | Описание ошибки для человека. Для обработки ошибки используйте поле &#x60;code&#x60;. | [optional] 

## Methods

### NewCreateDocumentErrorDTO

`func NewCreateDocumentErrorDTO(number string, code DocumentErrorCodeType, ) *CreateDocumentErrorDTO`

NewCreateDocumentErrorDTO instantiates a new CreateDocumentErrorDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDocumentErrorDTOWithDefaults

`func NewCreateDocumentErrorDTOWithDefaults() *CreateDocumentErrorDTO`

NewCreateDocumentErrorDTOWithDefaults instantiates a new CreateDocumentErrorDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *CreateDocumentErrorDTO) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CreateDocumentErrorDTO) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CreateDocumentErrorDTO) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetCode

`func (o *CreateDocumentErrorDTO) GetCode() DocumentErrorCodeType`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateDocumentErrorDTO) GetCodeOk() (*DocumentErrorCodeType, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateDocumentErrorDTO) SetCode(v DocumentErrorCodeType)`

SetCode sets Code field to given value.


### GetMessage

`func (o *CreateDocumentErrorDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateDocumentErrorDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateDocumentErrorDTO) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateDocumentErrorDTO) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


