# ConfigVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeSummary** | Pointer to [**ChangeSummary**](ChangeSummary.md) |  | [optional] 
**Created** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | **string** |  | 

## Methods

### NewConfigVersion

`func NewConfigVersion(id string, ) *ConfigVersion`

NewConfigVersion instantiates a new ConfigVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigVersionWithDefaults

`func NewConfigVersionWithDefaults() *ConfigVersion`

NewConfigVersionWithDefaults instantiates a new ConfigVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeSummary

`func (o *ConfigVersion) GetChangeSummary() ChangeSummary`

GetChangeSummary returns the ChangeSummary field if non-nil, zero value otherwise.

### GetChangeSummaryOk

`func (o *ConfigVersion) GetChangeSummaryOk() (*ChangeSummary, bool)`

GetChangeSummaryOk returns a tuple with the ChangeSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeSummary

`func (o *ConfigVersion) SetChangeSummary(v ChangeSummary)`

SetChangeSummary sets ChangeSummary field to given value.

### HasChangeSummary

`func (o *ConfigVersion) HasChangeSummary() bool`

HasChangeSummary returns a boolean if a field has been set.

### GetCreated

`func (o *ConfigVersion) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ConfigVersion) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ConfigVersion) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ConfigVersion) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDeleted

`func (o *ConfigVersion) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ConfigVersion) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ConfigVersion) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ConfigVersion) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetEnabled

`func (o *ConfigVersion) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConfigVersion) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConfigVersion) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ConfigVersion) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *ConfigVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigVersion) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


