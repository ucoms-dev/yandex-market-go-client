# BusinessOrderShipmentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Идентификатор отгрузки. | [optional] 
**ShipmentDate** | **string** | Дата отгрузки.  Формат даты: &#x60;ГГГГ-ММ-ДД&#x60;.  | 
**ShipmentTime** | Pointer to **NullableString** | Время отгрузки. | [optional] 

## Methods

### NewBusinessOrderShipmentDTO

`func NewBusinessOrderShipmentDTO(shipmentDate string, ) *BusinessOrderShipmentDTO`

NewBusinessOrderShipmentDTO instantiates a new BusinessOrderShipmentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessOrderShipmentDTOWithDefaults

`func NewBusinessOrderShipmentDTOWithDefaults() *BusinessOrderShipmentDTO`

NewBusinessOrderShipmentDTOWithDefaults instantiates a new BusinessOrderShipmentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BusinessOrderShipmentDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BusinessOrderShipmentDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BusinessOrderShipmentDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BusinessOrderShipmentDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BusinessOrderShipmentDTO) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BusinessOrderShipmentDTO) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetShipmentDate

`func (o *BusinessOrderShipmentDTO) GetShipmentDate() string`

GetShipmentDate returns the ShipmentDate field if non-nil, zero value otherwise.

### GetShipmentDateOk

`func (o *BusinessOrderShipmentDTO) GetShipmentDateOk() (*string, bool)`

GetShipmentDateOk returns a tuple with the ShipmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentDate

`func (o *BusinessOrderShipmentDTO) SetShipmentDate(v string)`

SetShipmentDate sets ShipmentDate field to given value.


### GetShipmentTime

`func (o *BusinessOrderShipmentDTO) GetShipmentTime() string`

GetShipmentTime returns the ShipmentTime field if non-nil, zero value otherwise.

### GetShipmentTimeOk

`func (o *BusinessOrderShipmentDTO) GetShipmentTimeOk() (*string, bool)`

GetShipmentTimeOk returns a tuple with the ShipmentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentTime

`func (o *BusinessOrderShipmentDTO) SetShipmentTime(v string)`

SetShipmentTime sets ShipmentTime field to given value.

### HasShipmentTime

`func (o *BusinessOrderShipmentDTO) HasShipmentTime() bool`

HasShipmentTime returns a boolean if a field has been set.

### SetShipmentTimeNil

`func (o *BusinessOrderShipmentDTO) SetShipmentTimeNil(b bool)`

 SetShipmentTimeNil sets the value for ShipmentTime to be an explicit nil

### UnsetShipmentTime
`func (o *BusinessOrderShipmentDTO) UnsetShipmentTime()`

UnsetShipmentTime ensures that no value is present for ShipmentTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


