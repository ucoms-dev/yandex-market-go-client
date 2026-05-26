# VotesDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Likes** | **int64** | Количество лайков. | 
**Dislikes** | **int64** | Количество дизлайков. | 

## Methods

### NewVotesDTO

`func NewVotesDTO(likes int64, dislikes int64, ) *VotesDTO`

NewVotesDTO instantiates a new VotesDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVotesDTOWithDefaults

`func NewVotesDTOWithDefaults() *VotesDTO`

NewVotesDTOWithDefaults instantiates a new VotesDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLikes

`func (o *VotesDTO) GetLikes() int64`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *VotesDTO) GetLikesOk() (*int64, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *VotesDTO) SetLikes(v int64)`

SetLikes sets Likes field to given value.


### GetDislikes

`func (o *VotesDTO) GetDislikes() int64`

GetDislikes returns the Dislikes field if non-nil, zero value otherwise.

### GetDislikesOk

`func (o *VotesDTO) GetDislikesOk() (*int64, bool)`

GetDislikesOk returns a tuple with the Dislikes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDislikes

`func (o *VotesDTO) SetDislikes(v int64)`

SetDislikes sets Dislikes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


