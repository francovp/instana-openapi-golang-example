# BrowserScriptConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Browser** | Pointer to **string** |  | [optional] 
**RecordVideo** | Pointer to **bool** |  | [optional] 
**Script** | Pointer to **string** |  | [optional] 
**ScriptType** | Pointer to **string** |  | [optional] 
**Scripts** | Pointer to [**MultipleScripts**](MultipleScripts.md) |  | [optional] 

## Methods

### NewBrowserScriptConfiguration

`func NewBrowserScriptConfiguration() *BrowserScriptConfiguration`

NewBrowserScriptConfiguration instantiates a new BrowserScriptConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserScriptConfigurationWithDefaults

`func NewBrowserScriptConfigurationWithDefaults() *BrowserScriptConfiguration`

NewBrowserScriptConfigurationWithDefaults instantiates a new BrowserScriptConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrowser

`func (o *BrowserScriptConfiguration) GetBrowser() string`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *BrowserScriptConfiguration) GetBrowserOk() (*string, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *BrowserScriptConfiguration) SetBrowser(v string)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *BrowserScriptConfiguration) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### GetRecordVideo

`func (o *BrowserScriptConfiguration) GetRecordVideo() bool`

GetRecordVideo returns the RecordVideo field if non-nil, zero value otherwise.

### GetRecordVideoOk

`func (o *BrowserScriptConfiguration) GetRecordVideoOk() (*bool, bool)`

GetRecordVideoOk returns a tuple with the RecordVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVideo

`func (o *BrowserScriptConfiguration) SetRecordVideo(v bool)`

SetRecordVideo sets RecordVideo field to given value.

### HasRecordVideo

`func (o *BrowserScriptConfiguration) HasRecordVideo() bool`

HasRecordVideo returns a boolean if a field has been set.

### GetScript

`func (o *BrowserScriptConfiguration) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *BrowserScriptConfiguration) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *BrowserScriptConfiguration) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *BrowserScriptConfiguration) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetScriptType

`func (o *BrowserScriptConfiguration) GetScriptType() string`

GetScriptType returns the ScriptType field if non-nil, zero value otherwise.

### GetScriptTypeOk

`func (o *BrowserScriptConfiguration) GetScriptTypeOk() (*string, bool)`

GetScriptTypeOk returns a tuple with the ScriptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptType

`func (o *BrowserScriptConfiguration) SetScriptType(v string)`

SetScriptType sets ScriptType field to given value.

### HasScriptType

`func (o *BrowserScriptConfiguration) HasScriptType() bool`

HasScriptType returns a boolean if a field has been set.

### GetScripts

`func (o *BrowserScriptConfiguration) GetScripts() MultipleScripts`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *BrowserScriptConfiguration) GetScriptsOk() (*MultipleScripts, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *BrowserScriptConfiguration) SetScripts(v MultipleScripts)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *BrowserScriptConfiguration) HasScripts() bool`

HasScripts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


