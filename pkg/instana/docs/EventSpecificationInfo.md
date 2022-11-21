# EventSpecificationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**EntityType** | **string** |  | 
**Id** | **string** |  | 
**Invalid** | Pointer to **bool** |  | [optional] 
**Migrated** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**Severity** | Pointer to **int32** |  | [optional] 
**Triggering** | Pointer to **bool** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewEventSpecificationInfo

`func NewEventSpecificationInfo(entityType string, id string, name string, type_ string, ) *EventSpecificationInfo`

NewEventSpecificationInfo instantiates a new EventSpecificationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSpecificationInfoWithDefaults

`func NewEventSpecificationInfoWithDefaults() *EventSpecificationInfo`

NewEventSpecificationInfoWithDefaults instantiates a new EventSpecificationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EventSpecificationInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventSpecificationInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventSpecificationInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventSpecificationInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *EventSpecificationInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EventSpecificationInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EventSpecificationInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EventSpecificationInfo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEntityType

`func (o *EventSpecificationInfo) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *EventSpecificationInfo) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *EventSpecificationInfo) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetId

`func (o *EventSpecificationInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventSpecificationInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventSpecificationInfo) SetId(v string)`

SetId sets Id field to given value.


### GetInvalid

`func (o *EventSpecificationInfo) GetInvalid() bool`

GetInvalid returns the Invalid field if non-nil, zero value otherwise.

### GetInvalidOk

`func (o *EventSpecificationInfo) GetInvalidOk() (*bool, bool)`

GetInvalidOk returns a tuple with the Invalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalid

`func (o *EventSpecificationInfo) SetInvalid(v bool)`

SetInvalid sets Invalid field to given value.

### HasInvalid

`func (o *EventSpecificationInfo) HasInvalid() bool`

HasInvalid returns a boolean if a field has been set.

### GetMigrated

`func (o *EventSpecificationInfo) GetMigrated() bool`

GetMigrated returns the Migrated field if non-nil, zero value otherwise.

### GetMigratedOk

`func (o *EventSpecificationInfo) GetMigratedOk() (*bool, bool)`

GetMigratedOk returns a tuple with the Migrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrated

`func (o *EventSpecificationInfo) SetMigrated(v bool)`

SetMigrated sets Migrated field to given value.

### HasMigrated

`func (o *EventSpecificationInfo) HasMigrated() bool`

HasMigrated returns a boolean if a field has been set.

### GetName

`func (o *EventSpecificationInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventSpecificationInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventSpecificationInfo) SetName(v string)`

SetName sets Name field to given value.


### GetSeverity

`func (o *EventSpecificationInfo) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *EventSpecificationInfo) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *EventSpecificationInfo) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *EventSpecificationInfo) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTriggering

`func (o *EventSpecificationInfo) GetTriggering() bool`

GetTriggering returns the Triggering field if non-nil, zero value otherwise.

### GetTriggeringOk

`func (o *EventSpecificationInfo) GetTriggeringOk() (*bool, bool)`

GetTriggeringOk returns a tuple with the Triggering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggering

`func (o *EventSpecificationInfo) SetTriggering(v bool)`

SetTriggering sets Triggering field to given value.

### HasTriggering

`func (o *EventSpecificationInfo) HasTriggering() bool`

HasTriggering returns a boolean if a field has been set.

### GetType

`func (o *EventSpecificationInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventSpecificationInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventSpecificationInfo) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


