# ValidatedAlertingChannelInputInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**EventTypes** | Pointer to **[]string** |  | [optional] 
**Id** | **string** |  | 
**Invalid** | Pointer to **bool** |  | [optional] 
**Label** | **string** |  | 
**Query** | Pointer to **string** |  | [optional] 
**SelectedEvents** | Pointer to **int32** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewValidatedAlertingChannelInputInfo

`func NewValidatedAlertingChannelInputInfo(id string, label string, type_ string, ) *ValidatedAlertingChannelInputInfo`

NewValidatedAlertingChannelInputInfo instantiates a new ValidatedAlertingChannelInputInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatedAlertingChannelInputInfoWithDefaults

`func NewValidatedAlertingChannelInputInfoWithDefaults() *ValidatedAlertingChannelInputInfo`

NewValidatedAlertingChannelInputInfoWithDefaults instantiates a new ValidatedAlertingChannelInputInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ValidatedAlertingChannelInputInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ValidatedAlertingChannelInputInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ValidatedAlertingChannelInputInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ValidatedAlertingChannelInputInfo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEntityId

`func (o *ValidatedAlertingChannelInputInfo) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ValidatedAlertingChannelInputInfo) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ValidatedAlertingChannelInputInfo) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *ValidatedAlertingChannelInputInfo) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEventTypes

`func (o *ValidatedAlertingChannelInputInfo) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *ValidatedAlertingChannelInputInfo) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *ValidatedAlertingChannelInputInfo) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *ValidatedAlertingChannelInputInfo) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetId

`func (o *ValidatedAlertingChannelInputInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ValidatedAlertingChannelInputInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ValidatedAlertingChannelInputInfo) SetId(v string)`

SetId sets Id field to given value.


### GetInvalid

`func (o *ValidatedAlertingChannelInputInfo) GetInvalid() bool`

GetInvalid returns the Invalid field if non-nil, zero value otherwise.

### GetInvalidOk

`func (o *ValidatedAlertingChannelInputInfo) GetInvalidOk() (*bool, bool)`

GetInvalidOk returns a tuple with the Invalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalid

`func (o *ValidatedAlertingChannelInputInfo) SetInvalid(v bool)`

SetInvalid sets Invalid field to given value.

### HasInvalid

`func (o *ValidatedAlertingChannelInputInfo) HasInvalid() bool`

HasInvalid returns a boolean if a field has been set.

### GetLabel

`func (o *ValidatedAlertingChannelInputInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ValidatedAlertingChannelInputInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ValidatedAlertingChannelInputInfo) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetQuery

`func (o *ValidatedAlertingChannelInputInfo) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ValidatedAlertingChannelInputInfo) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ValidatedAlertingChannelInputInfo) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ValidatedAlertingChannelInputInfo) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSelectedEvents

`func (o *ValidatedAlertingChannelInputInfo) GetSelectedEvents() int32`

GetSelectedEvents returns the SelectedEvents field if non-nil, zero value otherwise.

### GetSelectedEventsOk

`func (o *ValidatedAlertingChannelInputInfo) GetSelectedEventsOk() (*int32, bool)`

GetSelectedEventsOk returns a tuple with the SelectedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedEvents

`func (o *ValidatedAlertingChannelInputInfo) SetSelectedEvents(v int32)`

SetSelectedEvents sets SelectedEvents field to given value.

### HasSelectedEvents

`func (o *ValidatedAlertingChannelInputInfo) HasSelectedEvents() bool`

HasSelectedEvents returns a boolean if a field has been set.

### GetType

`func (o *ValidatedAlertingChannelInputInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ValidatedAlertingChannelInputInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ValidatedAlertingChannelInputInfo) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


