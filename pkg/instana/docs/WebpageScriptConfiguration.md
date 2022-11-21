# WebpageScriptConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Browser** | Pointer to **string** |  | [optional] 
**RecordVideo** | Pointer to **bool** |  | [optional] 
**Script** | **string** |  | 

## Methods

### NewWebpageScriptConfiguration

`func NewWebpageScriptConfiguration(script string, ) *WebpageScriptConfiguration`

NewWebpageScriptConfiguration instantiates a new WebpageScriptConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebpageScriptConfigurationWithDefaults

`func NewWebpageScriptConfigurationWithDefaults() *WebpageScriptConfiguration`

NewWebpageScriptConfigurationWithDefaults instantiates a new WebpageScriptConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrowser

`func (o *WebpageScriptConfiguration) GetBrowser() string`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *WebpageScriptConfiguration) GetBrowserOk() (*string, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *WebpageScriptConfiguration) SetBrowser(v string)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *WebpageScriptConfiguration) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### GetRecordVideo

`func (o *WebpageScriptConfiguration) GetRecordVideo() bool`

GetRecordVideo returns the RecordVideo field if non-nil, zero value otherwise.

### GetRecordVideoOk

`func (o *WebpageScriptConfiguration) GetRecordVideoOk() (*bool, bool)`

GetRecordVideoOk returns a tuple with the RecordVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVideo

`func (o *WebpageScriptConfiguration) SetRecordVideo(v bool)`

SetRecordVideo sets RecordVideo field to given value.

### HasRecordVideo

`func (o *WebpageScriptConfiguration) HasRecordVideo() bool`

HasRecordVideo returns a boolean if a field has been set.

### GetScript

`func (o *WebpageScriptConfiguration) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *WebpageScriptConfiguration) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *WebpageScriptConfiguration) SetScript(v string)`

SetScript sets Script field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


