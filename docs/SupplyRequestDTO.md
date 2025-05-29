# SupplyRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**SupplyRequestIdDTO**](SupplyRequestIdDTO.md) |  | 
**Type** | [**SupplyRequestType**](SupplyRequestType.md) |  | 
**Subtype** | [**SupplyRequestSubType**](SupplyRequestSubType.md) |  | 
**Status** | [**SupplyRequestStatusType**](SupplyRequestStatusType.md) |  | 
**UpdatedAt** | **time.Time** | Дата и время последнего обновления заявки. | 
**Counters** | [**SupplyRequestCountersDTO**](SupplyRequestCountersDTO.md) |  | 
**ParentLink** | Pointer to [**SupplyRequestReferenceDTO**](SupplyRequestReferenceDTO.md) |  | [optional] 
**ChildrenLinks** | Pointer to [**[]SupplyRequestReferenceDTO**](SupplyRequestReferenceDTO.md) | Ссылки на дочерние заявки. | [optional] 
**TargetLocation** | [**SupplyRequestLocationDTO**](SupplyRequestLocationDTO.md) |  | 
**TransitLocation** | Pointer to [**SupplyRequestLocationDTO**](SupplyRequestLocationDTO.md) |  | [optional] 

## Methods

### NewSupplyRequestDTO

`func NewSupplyRequestDTO(id SupplyRequestIdDTO, type_ SupplyRequestType, subtype SupplyRequestSubType, status SupplyRequestStatusType, updatedAt time.Time, counters SupplyRequestCountersDTO, targetLocation SupplyRequestLocationDTO, ) *SupplyRequestDTO`

NewSupplyRequestDTO instantiates a new SupplyRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplyRequestDTOWithDefaults

`func NewSupplyRequestDTOWithDefaults() *SupplyRequestDTO`

NewSupplyRequestDTOWithDefaults instantiates a new SupplyRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupplyRequestDTO) GetId() SupplyRequestIdDTO`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupplyRequestDTO) GetIdOk() (*SupplyRequestIdDTO, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupplyRequestDTO) SetId(v SupplyRequestIdDTO)`

SetId sets Id field to given value.


### GetType

`func (o *SupplyRequestDTO) GetType() SupplyRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SupplyRequestDTO) GetTypeOk() (*SupplyRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SupplyRequestDTO) SetType(v SupplyRequestType)`

SetType sets Type field to given value.


### GetSubtype

`func (o *SupplyRequestDTO) GetSubtype() SupplyRequestSubType`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *SupplyRequestDTO) GetSubtypeOk() (*SupplyRequestSubType, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *SupplyRequestDTO) SetSubtype(v SupplyRequestSubType)`

SetSubtype sets Subtype field to given value.


### GetStatus

`func (o *SupplyRequestDTO) GetStatus() SupplyRequestStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SupplyRequestDTO) GetStatusOk() (*SupplyRequestStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SupplyRequestDTO) SetStatus(v SupplyRequestStatusType)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *SupplyRequestDTO) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SupplyRequestDTO) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SupplyRequestDTO) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCounters

`func (o *SupplyRequestDTO) GetCounters() SupplyRequestCountersDTO`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *SupplyRequestDTO) GetCountersOk() (*SupplyRequestCountersDTO, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *SupplyRequestDTO) SetCounters(v SupplyRequestCountersDTO)`

SetCounters sets Counters field to given value.


### GetParentLink

`func (o *SupplyRequestDTO) GetParentLink() SupplyRequestReferenceDTO`

GetParentLink returns the ParentLink field if non-nil, zero value otherwise.

### GetParentLinkOk

`func (o *SupplyRequestDTO) GetParentLinkOk() (*SupplyRequestReferenceDTO, bool)`

GetParentLinkOk returns a tuple with the ParentLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentLink

`func (o *SupplyRequestDTO) SetParentLink(v SupplyRequestReferenceDTO)`

SetParentLink sets ParentLink field to given value.

### HasParentLink

`func (o *SupplyRequestDTO) HasParentLink() bool`

HasParentLink returns a boolean if a field has been set.

### GetChildrenLinks

`func (o *SupplyRequestDTO) GetChildrenLinks() []SupplyRequestReferenceDTO`

GetChildrenLinks returns the ChildrenLinks field if non-nil, zero value otherwise.

### GetChildrenLinksOk

`func (o *SupplyRequestDTO) GetChildrenLinksOk() (*[]SupplyRequestReferenceDTO, bool)`

GetChildrenLinksOk returns a tuple with the ChildrenLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenLinks

`func (o *SupplyRequestDTO) SetChildrenLinks(v []SupplyRequestReferenceDTO)`

SetChildrenLinks sets ChildrenLinks field to given value.

### HasChildrenLinks

`func (o *SupplyRequestDTO) HasChildrenLinks() bool`

HasChildrenLinks returns a boolean if a field has been set.

### SetChildrenLinksNil

`func (o *SupplyRequestDTO) SetChildrenLinksNil(b bool)`

 SetChildrenLinksNil sets the value for ChildrenLinks to be an explicit nil

### UnsetChildrenLinks
`func (o *SupplyRequestDTO) UnsetChildrenLinks()`

UnsetChildrenLinks ensures that no value is present for ChildrenLinks, not even an explicit nil
### GetTargetLocation

`func (o *SupplyRequestDTO) GetTargetLocation() SupplyRequestLocationDTO`

GetTargetLocation returns the TargetLocation field if non-nil, zero value otherwise.

### GetTargetLocationOk

`func (o *SupplyRequestDTO) GetTargetLocationOk() (*SupplyRequestLocationDTO, bool)`

GetTargetLocationOk returns a tuple with the TargetLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLocation

`func (o *SupplyRequestDTO) SetTargetLocation(v SupplyRequestLocationDTO)`

SetTargetLocation sets TargetLocation field to given value.


### GetTransitLocation

`func (o *SupplyRequestDTO) GetTransitLocation() SupplyRequestLocationDTO`

GetTransitLocation returns the TransitLocation field if non-nil, zero value otherwise.

### GetTransitLocationOk

`func (o *SupplyRequestDTO) GetTransitLocationOk() (*SupplyRequestLocationDTO, bool)`

GetTransitLocationOk returns a tuple with the TransitLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitLocation

`func (o *SupplyRequestDTO) SetTransitLocation(v SupplyRequestLocationDTO)`

SetTransitLocation sets TransitLocation field to given value.

### HasTransitLocation

`func (o *SupplyRequestDTO) HasTransitLocation() bool`

HasTransitLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


