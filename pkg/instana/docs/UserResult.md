# UserResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**FullName** | **string** |  | 
**Id** | **string** |  | 
**LastLoggedIn** | Pointer to **int64** |  | [optional] 

## Methods

### NewUserResult

`func NewUserResult(email string, fullName string, id string, ) *UserResult`

NewUserResult instantiates a new UserResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResultWithDefaults

`func NewUserResultWithDefaults() *UserResult`

NewUserResultWithDefaults instantiates a new UserResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserResult) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFullName

`func (o *UserResult) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *UserResult) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *UserResult) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetId

`func (o *UserResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserResult) SetId(v string)`

SetId sets Id field to given value.


### GetLastLoggedIn

`func (o *UserResult) GetLastLoggedIn() int64`

GetLastLoggedIn returns the LastLoggedIn field if non-nil, zero value otherwise.

### GetLastLoggedInOk

`func (o *UserResult) GetLastLoggedInOk() (*int64, bool)`

GetLastLoggedInOk returns a tuple with the LastLoggedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoggedIn

`func (o *UserResult) SetLastLoggedIn(v int64)`

SetLastLoggedIn sets LastLoggedIn field to given value.

### HasLastLoggedIn

`func (o *UserResult) HasLastLoggedIn() bool`

HasLastLoggedIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


