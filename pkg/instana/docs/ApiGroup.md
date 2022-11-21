# ApiGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Members** | [**[]ApiMember**](ApiMember.md) |  | 
**Name** | **string** |  | 
**PermissionSet** | [**ApiPermissionSetWithRoles**](ApiPermissionSetWithRoles.md) |  | 

## Methods

### NewApiGroup

`func NewApiGroup(members []ApiMember, name string, permissionSet ApiPermissionSetWithRoles, ) *ApiGroup`

NewApiGroup instantiates a new ApiGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiGroupWithDefaults

`func NewApiGroupWithDefaults() *ApiGroup`

NewApiGroupWithDefaults instantiates a new ApiGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMembers

`func (o *ApiGroup) GetMembers() []ApiMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ApiGroup) GetMembersOk() (*[]ApiMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ApiGroup) SetMembers(v []ApiMember)`

SetMembers sets Members field to given value.


### GetName

`func (o *ApiGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiGroup) SetName(v string)`

SetName sets Name field to given value.


### GetPermissionSet

`func (o *ApiGroup) GetPermissionSet() ApiPermissionSetWithRoles`

GetPermissionSet returns the PermissionSet field if non-nil, zero value otherwise.

### GetPermissionSetOk

`func (o *ApiGroup) GetPermissionSetOk() (*ApiPermissionSetWithRoles, bool)`

GetPermissionSetOk returns a tuple with the PermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSet

`func (o *ApiGroup) SetPermissionSet(v ApiPermissionSetWithRoles)`

SetPermissionSet sets PermissionSet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


