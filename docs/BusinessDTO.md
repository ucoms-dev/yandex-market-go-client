# BusinessDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Идентификатор кабинета. | [optional] 
**Name** | Pointer to **string** | Название бизнеса. | [optional] 

## Methods

### NewBusinessDTO

`func NewBusinessDTO() *BusinessDTO`

NewBusinessDTO instantiates a new BusinessDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessDTOWithDefaults

`func NewBusinessDTOWithDefaults() *BusinessDTO`

NewBusinessDTOWithDefaults instantiates a new BusinessDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BusinessDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BusinessDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BusinessDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BusinessDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BusinessDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BusinessDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BusinessDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BusinessDTO) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


