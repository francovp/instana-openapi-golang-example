# MaintenanceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Query** | **string** |  | 
**Windows** | Pointer to [**[]MaintenanceWindow**](MaintenanceWindow.md) |  | [optional] 

## Methods

### NewMaintenanceConfig

`func NewMaintenanceConfig(id string, name string, query string, ) *MaintenanceConfig`

NewMaintenanceConfig instantiates a new MaintenanceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceConfigWithDefaults

`func NewMaintenanceConfigWithDefaults() *MaintenanceConfig`

NewMaintenanceConfigWithDefaults instantiates a new MaintenanceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MaintenanceConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MaintenanceConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MaintenanceConfig) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MaintenanceConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MaintenanceConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MaintenanceConfig) SetName(v string)`

SetName sets Name field to given value.


### GetQuery

`func (o *MaintenanceConfig) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MaintenanceConfig) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MaintenanceConfig) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetWindows

`func (o *MaintenanceConfig) GetWindows() []MaintenanceWindow`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *MaintenanceConfig) GetWindowsOk() (*[]MaintenanceWindow, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *MaintenanceConfig) SetWindows(v []MaintenanceWindow)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *MaintenanceConfig) HasWindows() bool`

HasWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


