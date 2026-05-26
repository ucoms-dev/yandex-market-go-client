# CreateOrderPackagingDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageType** | Pointer to [**CreateOrderPackageType**](CreateOrderPackageType.md) |  | [optional] 

## Methods

### NewCreateOrderPackagingDTO

`func NewCreateOrderPackagingDTO() *CreateOrderPackagingDTO`

NewCreateOrderPackagingDTO instantiates a new CreateOrderPackagingDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderPackagingDTOWithDefaults

`func NewCreateOrderPackagingDTOWithDefaults() *CreateOrderPackagingDTO`

NewCreateOrderPackagingDTOWithDefaults instantiates a new CreateOrderPackagingDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageType

`func (o *CreateOrderPackagingDTO) GetPackageType() CreateOrderPackageType`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *CreateOrderPackagingDTO) GetPackageTypeOk() (*CreateOrderPackageType, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *CreateOrderPackagingDTO) SetPackageType(v CreateOrderPackageType)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *CreateOrderPackagingDTO) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


