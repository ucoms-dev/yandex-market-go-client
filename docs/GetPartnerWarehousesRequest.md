# GetPartnerWarehousesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WarehouseIds** | Pointer to **[]int64** | Список идентификаторов складов, которые необходимо вернуть.  Если параметр не указан, возвращаются все склады кабинета.  | [optional] 
**Components** | Pointer to [**[]PartnerWarehouseComponentType**](PartnerWarehouseComponentType.md) | Свойства складов, которые необходимо вернуть. Если какое-то значение параметра не задано, этой информации в ответе не будет.  Передавайте параметр, только если нужна информация, которую он возвращает.  | [optional] 

## Methods

### NewGetPartnerWarehousesRequest

`func NewGetPartnerWarehousesRequest() *GetPartnerWarehousesRequest`

NewGetPartnerWarehousesRequest instantiates a new GetPartnerWarehousesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPartnerWarehousesRequestWithDefaults

`func NewGetPartnerWarehousesRequestWithDefaults() *GetPartnerWarehousesRequest`

NewGetPartnerWarehousesRequestWithDefaults instantiates a new GetPartnerWarehousesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarehouseIds

`func (o *GetPartnerWarehousesRequest) GetWarehouseIds() []int64`

GetWarehouseIds returns the WarehouseIds field if non-nil, zero value otherwise.

### GetWarehouseIdsOk

`func (o *GetPartnerWarehousesRequest) GetWarehouseIdsOk() (*[]int64, bool)`

GetWarehouseIdsOk returns a tuple with the WarehouseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseIds

`func (o *GetPartnerWarehousesRequest) SetWarehouseIds(v []int64)`

SetWarehouseIds sets WarehouseIds field to given value.

### HasWarehouseIds

`func (o *GetPartnerWarehousesRequest) HasWarehouseIds() bool`

HasWarehouseIds returns a boolean if a field has been set.

### SetWarehouseIdsNil

`func (o *GetPartnerWarehousesRequest) SetWarehouseIdsNil(b bool)`

 SetWarehouseIdsNil sets the value for WarehouseIds to be an explicit nil

### UnsetWarehouseIds
`func (o *GetPartnerWarehousesRequest) UnsetWarehouseIds()`

UnsetWarehouseIds ensures that no value is present for WarehouseIds, not even an explicit nil
### GetComponents

`func (o *GetPartnerWarehousesRequest) GetComponents() []PartnerWarehouseComponentType`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *GetPartnerWarehousesRequest) GetComponentsOk() (*[]PartnerWarehouseComponentType, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *GetPartnerWarehousesRequest) SetComponents(v []PartnerWarehouseComponentType)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *GetPartnerWarehousesRequest) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### SetComponentsNil

`func (o *GetPartnerWarehousesRequest) SetComponentsNil(b bool)`

 SetComponentsNil sets the value for Components to be an explicit nil

### UnsetComponents
`func (o *GetPartnerWarehousesRequest) UnsetComponents()`

UnsetComponents ensures that no value is present for Components, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


