# GetBusinessSettingsInfoDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to [**BusinessDTO**](BusinessDTO.md) |  | [optional] 
**Settings** | Pointer to [**BusinessSettingsDTO**](BusinessSettingsDTO.md) |  | [optional] 
**SubscriptionLevel** | Pointer to [**BusinessSubscriptionLevelType**](BusinessSubscriptionLevelType.md) |  | [optional] 
**Traits** | Pointer to [**[]BusinessTraitType**](BusinessTraitType.md) | Свойства кабинета. | [optional] 

## Methods

### NewGetBusinessSettingsInfoDTO

`func NewGetBusinessSettingsInfoDTO() *GetBusinessSettingsInfoDTO`

NewGetBusinessSettingsInfoDTO instantiates a new GetBusinessSettingsInfoDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBusinessSettingsInfoDTOWithDefaults

`func NewGetBusinessSettingsInfoDTOWithDefaults() *GetBusinessSettingsInfoDTO`

NewGetBusinessSettingsInfoDTOWithDefaults instantiates a new GetBusinessSettingsInfoDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *GetBusinessSettingsInfoDTO) GetInfo() BusinessDTO`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *GetBusinessSettingsInfoDTO) GetInfoOk() (*BusinessDTO, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *GetBusinessSettingsInfoDTO) SetInfo(v BusinessDTO)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *GetBusinessSettingsInfoDTO) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetSettings

`func (o *GetBusinessSettingsInfoDTO) GetSettings() BusinessSettingsDTO`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *GetBusinessSettingsInfoDTO) GetSettingsOk() (*BusinessSettingsDTO, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *GetBusinessSettingsInfoDTO) SetSettings(v BusinessSettingsDTO)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *GetBusinessSettingsInfoDTO) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSubscriptionLevel

`func (o *GetBusinessSettingsInfoDTO) GetSubscriptionLevel() BusinessSubscriptionLevelType`

GetSubscriptionLevel returns the SubscriptionLevel field if non-nil, zero value otherwise.

### GetSubscriptionLevelOk

`func (o *GetBusinessSettingsInfoDTO) GetSubscriptionLevelOk() (*BusinessSubscriptionLevelType, bool)`

GetSubscriptionLevelOk returns a tuple with the SubscriptionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionLevel

`func (o *GetBusinessSettingsInfoDTO) SetSubscriptionLevel(v BusinessSubscriptionLevelType)`

SetSubscriptionLevel sets SubscriptionLevel field to given value.

### HasSubscriptionLevel

`func (o *GetBusinessSettingsInfoDTO) HasSubscriptionLevel() bool`

HasSubscriptionLevel returns a boolean if a field has been set.

### GetTraits

`func (o *GetBusinessSettingsInfoDTO) GetTraits() []BusinessTraitType`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *GetBusinessSettingsInfoDTO) GetTraitsOk() (*[]BusinessTraitType, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *GetBusinessSettingsInfoDTO) SetTraits(v []BusinessTraitType)`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *GetBusinessSettingsInfoDTO) HasTraits() bool`

HasTraits returns a boolean if a field has been set.

### SetTraitsNil

`func (o *GetBusinessSettingsInfoDTO) SetTraitsNil(b bool)`

 SetTraitsNil sets the value for Traits to be an explicit nil

### UnsetTraits
`func (o *GetBusinessSettingsInfoDTO) UnsetTraits()`

UnsetTraits ensures that no value is present for Traits, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


