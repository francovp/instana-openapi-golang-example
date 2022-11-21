# MaintenanceConfigWithLastUpdated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**LastUpdated** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**Query** | **string** |  | 
**Windows** | Pointer to [**[]MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] 

## Methods

### NewMaintenanceConfigWithLastUpdated

`func NewMaintenanceConfigWithLastUpdated(id string, name string, query string, ) *MaintenanceConfigWithLastUpdated`

NewMaintenanceConfigWithLastUpdated instantiates a new MaintenanceConfigWithLastUpdated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceConfigWithLastUpdatedWithDefaults

`func NewMaintenanceConfigWithLastUpdatedWithDefaults() *MaintenanceConfigWithLastUpdated`

NewMaintenanceConfigWithLastUpdatedWithDefaults instantiates a new MaintenanceConfigWithLastUpdated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MaintenanceConfigWithLastUpdated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MaintenanceConfigWithLastUpdated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MaintenanceConfigWithLastUpdated) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdated

`func (o *MaintenanceConfigWithLastUpdated) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *MaintenanceConfigWithLastUpdated) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *MaintenanceConfigWithLastUpdated) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *MaintenanceConfigWithLastUpdated) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *MaintenanceConfigWithLastUpdated) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MaintenanceConfigWithLastUpdated) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MaintenanceConfigWithLastUpdated) SetName(v string)`

SetName sets Name field to given value.


### GetQuery

`func (o *MaintenanceConfigWithLastUpdated) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MaintenanceConfigWithLastUpdated) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MaintenanceConfigWithLastUpdated) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetWindows

`func (o *MaintenanceConfigWithLastUpdated) GetWindows() []MaintenanceWindow`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *MaintenanceConfigWithLastUpdated) GetWindowsOk() (*[]MaintenanceWindow, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *MaintenanceConfigWithLastUpdated) SetWindows(v []MaintenanceWindow)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *MaintenanceConfigWithLastUpdated) HasWindows() bool`

HasWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


