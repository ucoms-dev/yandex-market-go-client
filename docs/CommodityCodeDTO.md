# CommodityCodeDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Товарный код. | 
**Type** | [**CommodityCodeType**](CommodityCodeType.md) |  | 

## Methods

### NewCommodityCodeDTO

`func NewCommodityCodeDTO(code string, type_ CommodityCodeType, ) *CommodityCodeDTO`

NewCommodityCodeDTO instantiates a new CommodityCodeDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommodityCodeDTOWithDefaults

`func NewCommodityCodeDTOWithDefaults() *CommodityCodeDTO`

NewCommodityCodeDTOWithDefaults instantiates a new CommodityCodeDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CommodityCodeDTO) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CommodityCodeDTO) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CommodityCodeDTO) SetCode(v string)`

SetCode sets Code field to given value.


### GetType

`func (o *CommodityCodeDTO) GetType() CommodityCodeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommodityCodeDTO) GetTypeOk() (*CommodityCodeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommodityCodeDTO) SetType(v CommodityCodeType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


