# UsersResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invitations** | Pointer to [**[]InvitationResult**](InvitationResult.md) |  | [optional] 
**Users** | Pointer to [**[]UserResult**](UserResult.md) |  | [optional] 

## Methods

### NewUsersResult

`func NewUsersResult() *UsersResult`

NewUsersResult instantiates a new UsersResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersResultWithDefaults

`func NewUsersResultWithDefaults() *UsersResult`

NewUsersResultWithDefaults instantiates a new UsersResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvitations

`func (o *UsersResult) GetInvitations() []InvitationResult`

GetInvitations returns the Invitations field if non-nil, zero value otherwise.

### GetInvitationsOk

`func (o *UsersResult) GetInvitationsOk() (*[]InvitationResult, bool)`

GetInvitationsOk returns a tuple with the Invitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitations

`func (o *UsersResult) SetInvitations(v []InvitationResult)`

SetInvitations sets Invitations field to given value.

### HasInvitations

`func (o *UsersResult) HasInvitations() bool`

HasInvitations returns a boolean if a field has been set.

### GetUsers

`func (o *UsersResult) GetUsers() []UserResult`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UsersResult) GetUsersOk() (*[]UserResult, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UsersResult) SetUsers(v []UserResult)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UsersResult) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


