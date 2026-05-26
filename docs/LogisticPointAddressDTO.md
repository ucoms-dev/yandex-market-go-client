# LogisticPointAddressDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullAddress** | **string** | Полный адрес. | 
**Gps** | [**GpsDTO**](GpsDTO.md) |  | 
**RegionId** | **int64** | Идентификатор региона.  Информацию о регионе можно получить c помощью метода [GET v2/regions](../../reference/regions/searchRegionsById.md).  | 
**City** | Pointer to **string** | Город. | [optional] 
**Street** | Pointer to **string** | Улица. | [optional] 
**House** | Pointer to **string** | Номер дома. | [optional] 
**Building** | Pointer to **string** | Номер строения. | [optional] 
**Block** | Pointer to **string** | Номер корпуса. | [optional] 
**Km** | Pointer to **int32** | Порядковый номер километра, на котором располагается пункт выдачи.  Указывается, если в адресе нет улицы.  | [optional] 
**Additional** | Pointer to **string** | Дополнительная информация. | [optional] 

## Methods

### NewLogisticPointAddressDTO

`func NewLogisticPointAddressDTO(fullAddress string, gps GpsDTO, regionId int64, ) *LogisticPointAddressDTO`

NewLogisticPointAddressDTO instantiates a new LogisticPointAddressDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogisticPointAddressDTOWithDefaults

`func NewLogisticPointAddressDTOWithDefaults() *LogisticPointAddressDTO`

NewLogisticPointAddressDTOWithDefaults instantiates a new LogisticPointAddressDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullAddress

`func (o *LogisticPointAddressDTO) GetFullAddress() string`

GetFullAddress returns the FullAddress field if non-nil, zero value otherwise.

### GetFullAddressOk

`func (o *LogisticPointAddressDTO) GetFullAddressOk() (*string, bool)`

GetFullAddressOk returns a tuple with the FullAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullAddress

`func (o *LogisticPointAddressDTO) SetFullAddress(v string)`

SetFullAddress sets FullAddress field to given value.


### GetGps

`func (o *LogisticPointAddressDTO) GetGps() GpsDTO`

GetGps returns the Gps field if non-nil, zero value otherwise.

### GetGpsOk

`func (o *LogisticPointAddressDTO) GetGpsOk() (*GpsDTO, bool)`

GetGpsOk returns a tuple with the Gps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGps

`func (o *LogisticPointAddressDTO) SetGps(v GpsDTO)`

SetGps sets Gps field to given value.


### GetRegionId

`func (o *LogisticPointAddressDTO) GetRegionId() int64`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *LogisticPointAddressDTO) GetRegionIdOk() (*int64, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *LogisticPointAddressDTO) SetRegionId(v int64)`

SetRegionId sets RegionId field to given value.


### GetCity

`func (o *LogisticPointAddressDTO) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *LogisticPointAddressDTO) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *LogisticPointAddressDTO) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *LogisticPointAddressDTO) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetStreet

`func (o *LogisticPointAddressDTO) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *LogisticPointAddressDTO) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *LogisticPointAddressDTO) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *LogisticPointAddressDTO) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetHouse

`func (o *LogisticPointAddressDTO) GetHouse() string`

GetHouse returns the House field if non-nil, zero value otherwise.

### GetHouseOk

`func (o *LogisticPointAddressDTO) GetHouseOk() (*string, bool)`

GetHouseOk returns a tuple with the House field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouse

`func (o *LogisticPointAddressDTO) SetHouse(v string)`

SetHouse sets House field to given value.

### HasHouse

`func (o *LogisticPointAddressDTO) HasHouse() bool`

HasHouse returns a boolean if a field has been set.

### GetBuilding

`func (o *LogisticPointAddressDTO) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *LogisticPointAddressDTO) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *LogisticPointAddressDTO) SetBuilding(v string)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *LogisticPointAddressDTO) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### GetBlock

`func (o *LogisticPointAddressDTO) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *LogisticPointAddressDTO) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *LogisticPointAddressDTO) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *LogisticPointAddressDTO) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetKm

`func (o *LogisticPointAddressDTO) GetKm() int32`

GetKm returns the Km field if non-nil, zero value otherwise.

### GetKmOk

`func (o *LogisticPointAddressDTO) GetKmOk() (*int32, bool)`

GetKmOk returns a tuple with the Km field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKm

`func (o *LogisticPointAddressDTO) SetKm(v int32)`

SetKm sets Km field to given value.

### HasKm

`func (o *LogisticPointAddressDTO) HasKm() bool`

HasKm returns a boolean if a field has been set.

### GetAdditional

`func (o *LogisticPointAddressDTO) GetAdditional() string`

GetAdditional returns the Additional field if non-nil, zero value otherwise.

### GetAdditionalOk

`func (o *LogisticPointAddressDTO) GetAdditionalOk() (*string, bool)`

GetAdditionalOk returns a tuple with the Additional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditional

`func (o *LogisticPointAddressDTO) SetAdditional(v string)`

SetAdditional sets Additional field to given value.

### HasAdditional

`func (o *LogisticPointAddressDTO) HasAdditional() bool`

HasAdditional returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


