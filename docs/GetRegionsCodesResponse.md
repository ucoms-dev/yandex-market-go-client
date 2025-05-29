# GetRegionsCodesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Countries** | [**[]CountryDTO**](CountryDTO.md) | Список стран с их кодами в формате ISO 3166-1 alpha-2. | 

## Methods

### NewGetRegionsCodesResponse

`func NewGetRegionsCodesResponse(countries []CountryDTO, ) *GetRegionsCodesResponse`

NewGetRegionsCodesResponse instantiates a new GetRegionsCodesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRegionsCodesResponseWithDefaults

`func NewGetRegionsCodesResponseWithDefaults() *GetRegionsCodesResponse`

NewGetRegionsCodesResponseWithDefaults instantiates a new GetRegionsCodesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountries

`func (o *GetRegionsCodesResponse) GetCountries() []CountryDTO`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *GetRegionsCodesResponse) GetCountriesOk() (*[]CountryDTO, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *GetRegionsCodesResponse) SetCountries(v []CountryDTO)`

SetCountries sets Countries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


