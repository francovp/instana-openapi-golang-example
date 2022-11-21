# HttpScriptConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | Pointer to **string** |  | [optional] 
**ScriptType** | Pointer to **string** |  | [optional] 
**Scripts** | Pointer to [**MultipleScripts**](MultipleScripts.md) |  | [optional] 

## Methods

### NewHttpScriptConfiguration

`func NewHttpScriptConfiguration() *HttpScriptConfiguration`

NewHttpScriptConfiguration instantiates a new HttpScriptConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpScriptConfigurationWithDefaults

`func NewHttpScriptConfigurationWithDefaults() *HttpScriptConfiguration`

NewHttpScriptConfigurationWithDefaults instantiates a new HttpScriptConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *HttpScriptConfiguration) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *HttpScriptConfiguration) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *HttpScriptConfiguration) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *HttpScriptConfiguration) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetScriptType

`func (o *HttpScriptConfiguration) GetScriptType() string`

GetScriptType returns the ScriptType field if non-nil, zero value otherwise.

### GetScriptTypeOk

`func (o *HttpScriptConfiguration) GetScriptTypeOk() (*string, bool)`

GetScriptTypeOk returns a tuple with the ScriptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptType

`func (o *HttpScriptConfiguration) SetScriptType(v string)`

SetScriptType sets ScriptType field to given value.

### HasScriptType

`func (o *HttpScriptConfiguration) HasScriptType() bool`

HasScriptType returns a boolean if a field has been set.

### GetScripts

`func (o *HttpScriptConfiguration) GetScripts() MultipleScripts`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *HttpScriptConfiguration) GetScriptsOk() (*MultipleScripts, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *HttpScriptConfiguration) SetScripts(v MultipleScripts)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *HttpScriptConfiguration) HasScripts() bool`

HasScripts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


