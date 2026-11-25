# SaveSizeTablesResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]SaveSizeTableResultDTO**](SaveSizeTableResultDTO.md) | Результаты по таблицам из запроса в том же порядке: &#x60;results[i]&#x60; относится к &#x60;sizeTables[i]&#x60;.  | 

## Methods

### NewSaveSizeTablesResultDTO

`func NewSaveSizeTablesResultDTO(results []SaveSizeTableResultDTO, ) *SaveSizeTablesResultDTO`

NewSaveSizeTablesResultDTO instantiates a new SaveSizeTablesResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveSizeTablesResultDTOWithDefaults

`func NewSaveSizeTablesResultDTOWithDefaults() *SaveSizeTablesResultDTO`

NewSaveSizeTablesResultDTOWithDefaults instantiates a new SaveSizeTablesResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *SaveSizeTablesResultDTO) GetResults() []SaveSizeTableResultDTO`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SaveSizeTablesResultDTO) GetResultsOk() (*[]SaveSizeTableResultDTO, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SaveSizeTablesResultDTO) SetResults(v []SaveSizeTableResultDTO)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


