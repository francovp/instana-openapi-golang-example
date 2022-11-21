# SessionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdleTimeInMillis** | Pointer to **int64** |  | [optional] 
**TokenLifeTimeInMillis** | Pointer to **int64** |  | [optional] 

## Methods

### NewSessionSettings

`func NewSessionSettings() *SessionSettings`

NewSessionSettings instantiates a new SessionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionSettingsWithDefaults

`func NewSessionSettingsWithDefaults() *SessionSettings`

NewSessionSettingsWithDefaults instantiates a new SessionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdleTimeInMillis

`func (o *SessionSettings) GetIdleTimeInMillis() int64`

GetIdleTimeInMillis returns the IdleTimeInMillis field if non-nil, zero value otherwise.

### GetIdleTimeInMillisOk

`func (o *SessionSettings) GetIdleTimeInMillisOk() (*int64, bool)`

GetIdleTimeInMillisOk returns a tuple with the IdleTimeInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeInMillis

`func (o *SessionSettings) SetIdleTimeInMillis(v int64)`

SetIdleTimeInMillis sets IdleTimeInMillis field to given value.

### HasIdleTimeInMillis

`func (o *SessionSettings) HasIdleTimeInMillis() bool`

HasIdleTimeInMillis returns a boolean if a field has been set.

### GetTokenLifeTimeInMillis

`func (o *SessionSettings) GetTokenLifeTimeInMillis() int64`

GetTokenLifeTimeInMillis returns the TokenLifeTimeInMillis field if non-nil, zero value otherwise.

### GetTokenLifeTimeInMillisOk

`func (o *SessionSettings) GetTokenLifeTimeInMillisOk() (*int64, bool)`

GetTokenLifeTimeInMillisOk returns a tuple with the TokenLifeTimeInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLifeTimeInMillis

`func (o *SessionSettings) SetTokenLifeTimeInMillis(v int64)`

SetTokenLifeTimeInMillis sets TokenLifeTimeInMillis field to given value.

### HasTokenLifeTimeInMillis

`func (o *SessionSettings) HasTokenLifeTimeInMillis() bool`

HasTokenLifeTimeInMillis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


