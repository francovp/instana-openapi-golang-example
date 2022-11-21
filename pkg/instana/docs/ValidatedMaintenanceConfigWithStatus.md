# ValidatedMaintenanceConfigWithStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Invalid** | Pointer to **bool** |  | [optional] 
**LastUpdated** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**Query** | **string** |  | 
**Status** | **string** |  | 
**Windows** | Pointer to [**[]MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] 

## Methods

### NewValidatedMaintenanceConfigWithStatus

`func NewValidatedMaintenanceConfigWithStatus(id string, name string, query string, status string, ) *ValidatedMaintenanceConfigWithStatus`

NewValidatedMaintenanceConfigWithStatus instantiates a new ValidatedMaintenanceConfigWithStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatedMaintenanceConfigWithStatusWithDefaults

`func NewValidatedMaintenanceConfigWithStatusWithDefaults() *ValidatedMaintenanceConfigWithStatus`

NewValidatedMaintenanceConfigWithStatusWithDefaults instantiates a new ValidatedMaintenanceConfigWithStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ValidatedMaintenanceConfigWithStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ValidatedMaintenanceConfigWithStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ValidatedMaintenanceConfigWithStatus) SetId(v string)`

SetId sets Id field to given value.


### GetInvalid

`func (o *ValidatedMaintenanceConfigWithStatus) GetInvalid() bool`

GetInvalid returns the Invalid field if non-nil, zero value otherwise.

### GetInvalidOk

`func (o *ValidatedMaintenanceConfigWithStatus) GetInvalidOk() (*bool, bool)`

GetInvalidOk returns a tuple with the Invalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalid

`func (o *ValidatedMaintenanceConfigWithStatus) SetInvalid(v bool)`

SetInvalid sets Invalid field to given value.

### HasInvalid

`func (o *ValidatedMaintenanceConfigWithStatus) HasInvalid() bool`

HasInvalid returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ValidatedMaintenanceConfigWithStatus) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ValidatedMaintenanceConfigWithStatus) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ValidatedMaintenanceConfigWithStatus) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ValidatedMaintenanceConfigWithStatus) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *ValidatedMaintenanceConfigWithStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ValidatedMaintenanceConfigWithStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ValidatedMaintenanceConfigWithStatus) SetName(v string)`

SetName sets Name field to given value.


### GetQuery

`func (o *ValidatedMaintenanceConfigWithStatus) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ValidatedMaintenanceConfigWithStatus) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ValidatedMaintenanceConfigWithStatus) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetStatus

`func (o *ValidatedMaintenanceConfigWithStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ValidatedMaintenanceConfigWithStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ValidatedMaintenanceConfigWithStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetWindows

`func (o *ValidatedMaintenanceConfigWithStatus) GetWindows() []MaintenanceWindow`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *ValidatedMaintenanceConfigWithStatus) GetWindowsOk() (*[]MaintenanceWindow, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *ValidatedMaintenanceConfigWithStatus) SetWindows(v []MaintenanceWindow)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *ValidatedMaintenanceConfigWithStatus) HasWindows() bool`

HasWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


