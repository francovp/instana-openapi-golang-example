# HealthState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | Pointer to **string** |  | [optional] 
**Messages** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHealthState

`func NewHealthState() *HealthState`

NewHealthState instantiates a new HealthState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthStateWithDefaults

`func NewHealthStateWithDefaults() *HealthState`

NewHealthStateWithDefaults instantiates a new HealthState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *HealthState) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *HealthState) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *HealthState) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *HealthState) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetMessages

`func (o *HealthState) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *HealthState) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *HealthState) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *HealthState) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


