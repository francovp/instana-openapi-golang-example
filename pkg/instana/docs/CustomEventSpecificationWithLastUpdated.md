# CustomEventSpecificationWithLastUpdated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]Action**](Action.md) |  | [optional] 
**ApplicationAlertConfigId** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**EntityType** | **string** |  | 
**ExpirationTime** | Pointer to **int64** |  | [optional] 
**Id** | **string** |  | 
**LastUpdated** | Pointer to **int64** |  | [optional] 
**Migrated** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**Query** | Pointer to **string** |  | [optional] 
**Rules** | [**[]AbstractRule**](AbstractRule.md) |  | 
**Triggering** | Pointer to **bool** |  | [optional] 
**ValidVersion** | Pointer to **int32** |  | [optional] 

## Methods

### NewCustomEventSpecificationWithLastUpdated

`func NewCustomEventSpecificationWithLastUpdated(entityType string, id string, name string, rules []AbstractRule, ) *CustomEventSpecificationWithLastUpdated`

NewCustomEventSpecificationWithLastUpdated instantiates a new CustomEventSpecificationWithLastUpdated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEventSpecificationWithLastUpdatedWithDefaults

`func NewCustomEventSpecificationWithLastUpdatedWithDefaults() *CustomEventSpecificationWithLastUpdated`

NewCustomEventSpecificationWithLastUpdatedWithDefaults instantiates a new CustomEventSpecificationWithLastUpdated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *CustomEventSpecificationWithLastUpdated) GetActions() []Action`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CustomEventSpecificationWithLastUpdated) GetActionsOk() (*[]Action, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CustomEventSpecificationWithLastUpdated) SetActions(v []Action)`

SetActions sets Actions field to given value.

### HasActions

`func (o *CustomEventSpecificationWithLastUpdated) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetApplicationAlertConfigId

`func (o *CustomEventSpecificationWithLastUpdated) GetApplicationAlertConfigId() string`

GetApplicationAlertConfigId returns the ApplicationAlertConfigId field if non-nil, zero value otherwise.

### GetApplicationAlertConfigIdOk

`func (o *CustomEventSpecificationWithLastUpdated) GetApplicationAlertConfigIdOk() (*string, bool)`

GetApplicationAlertConfigIdOk returns a tuple with the ApplicationAlertConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAlertConfigId

`func (o *CustomEventSpecificationWithLastUpdated) SetApplicationAlertConfigId(v string)`

SetApplicationAlertConfigId sets ApplicationAlertConfigId field to given value.

### HasApplicationAlertConfigId

`func (o *CustomEventSpecificationWithLastUpdated) HasApplicationAlertConfigId() bool`

HasApplicationAlertConfigId returns a boolean if a field has been set.

### GetDeleted

`func (o *CustomEventSpecificationWithLastUpdated) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *CustomEventSpecificationWithLastUpdated) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *CustomEventSpecificationWithLastUpdated) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *CustomEventSpecificationWithLastUpdated) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDescription

`func (o *CustomEventSpecificationWithLastUpdated) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomEventSpecificationWithLastUpdated) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomEventSpecificationWithLastUpdated) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomEventSpecificationWithLastUpdated) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CustomEventSpecificationWithLastUpdated) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustomEventSpecificationWithLastUpdated) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustomEventSpecificationWithLastUpdated) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustomEventSpecificationWithLastUpdated) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEntityType

`func (o *CustomEventSpecificationWithLastUpdated) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CustomEventSpecificationWithLastUpdated) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CustomEventSpecificationWithLastUpdated) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetExpirationTime

`func (o *CustomEventSpecificationWithLastUpdated) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *CustomEventSpecificationWithLastUpdated) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *CustomEventSpecificationWithLastUpdated) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *CustomEventSpecificationWithLastUpdated) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetId

`func (o *CustomEventSpecificationWithLastUpdated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomEventSpecificationWithLastUpdated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomEventSpecificationWithLastUpdated) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdated

`func (o *CustomEventSpecificationWithLastUpdated) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CustomEventSpecificationWithLastUpdated) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CustomEventSpecificationWithLastUpdated) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CustomEventSpecificationWithLastUpdated) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetMigrated

`func (o *CustomEventSpecificationWithLastUpdated) GetMigrated() bool`

GetMigrated returns the Migrated field if non-nil, zero value otherwise.

### GetMigratedOk

`func (o *CustomEventSpecificationWithLastUpdated) GetMigratedOk() (*bool, bool)`

GetMigratedOk returns a tuple with the Migrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrated

`func (o *CustomEventSpecificationWithLastUpdated) SetMigrated(v bool)`

SetMigrated sets Migrated field to given value.

### HasMigrated

`func (o *CustomEventSpecificationWithLastUpdated) HasMigrated() bool`

HasMigrated returns a boolean if a field has been set.

### GetName

`func (o *CustomEventSpecificationWithLastUpdated) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEventSpecificationWithLastUpdated) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomEventSpecificationWithLastUpdated) SetName(v string)`

SetName sets Name field to given value.


### GetQuery

`func (o *CustomEventSpecificationWithLastUpdated) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CustomEventSpecificationWithLastUpdated) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CustomEventSpecificationWithLastUpdated) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CustomEventSpecificationWithLastUpdated) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetRules

`func (o *CustomEventSpecificationWithLastUpdated) GetRules() []AbstractRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CustomEventSpecificationWithLastUpdated) GetRulesOk() (*[]AbstractRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CustomEventSpecificationWithLastUpdated) SetRules(v []AbstractRule)`

SetRules sets Rules field to given value.


### GetTriggering

`func (o *CustomEventSpecificationWithLastUpdated) GetTriggering() bool`

GetTriggering returns the Triggering field if non-nil, zero value otherwise.

### GetTriggeringOk

`func (o *CustomEventSpecificationWithLastUpdated) GetTriggeringOk() (*bool, bool)`

GetTriggeringOk returns a tuple with the Triggering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggering

`func (o *CustomEventSpecificationWithLastUpdated) SetTriggering(v bool)`

SetTriggering sets Triggering field to given value.

### HasTriggering

`func (o *CustomEventSpecificationWithLastUpdated) HasTriggering() bool`

HasTriggering returns a boolean if a field has been set.

### GetValidVersion

`func (o *CustomEventSpecificationWithLastUpdated) GetValidVersion() int32`

GetValidVersion returns the ValidVersion field if non-nil, zero value otherwise.

### GetValidVersionOk

`func (o *CustomEventSpecificationWithLastUpdated) GetValidVersionOk() (*int32, bool)`

GetValidVersionOk returns a tuple with the ValidVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidVersion

`func (o *CustomEventSpecificationWithLastUpdated) SetValidVersion(v int32)`

SetValidVersion sets ValidVersion field to given value.

### HasValidVersion

`func (o *CustomEventSpecificationWithLastUpdated) HasValidVersion() bool`

HasValidVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


