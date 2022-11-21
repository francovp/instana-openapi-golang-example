# WebpageActionConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Browser** | Pointer to **string** |  | [optional] 
**RecordVideo** | Pointer to **bool** |  | [optional] 
**Url** | **string** |  | 

## Methods

### NewWebpageActionConfiguration

`func NewWebpageActionConfiguration(url string, ) *WebpageActionConfiguration`

NewWebpageActionConfiguration instantiates a new WebpageActionConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebpageActionConfigurationWithDefaults

`func NewWebpageActionConfigurationWithDefaults() *WebpageActionConfiguration`

NewWebpageActionConfigurationWithDefaults instantiates a new WebpageActionConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrowser

`func (o *WebpageActionConfiguration) GetBrowser() string`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *WebpageActionConfiguration) GetBrowserOk() (*string, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *WebpageActionConfiguration) SetBrowser(v string)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *WebpageActionConfiguration) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### GetRecordVideo

`func (o *WebpageActionConfiguration) GetRecordVideo() bool`

GetRecordVideo returns the RecordVideo field if non-nil, zero value otherwise.

### GetRecordVideoOk

`func (o *WebpageActionConfiguration) GetRecordVideoOk() (*bool, bool)`

GetRecordVideoOk returns a tuple with the RecordVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVideo

`func (o *WebpageActionConfiguration) SetRecordVideo(v bool)`

SetRecordVideo sets RecordVideo field to given value.

### HasRecordVideo

`func (o *WebpageActionConfiguration) HasRecordVideo() bool`

HasRecordVideo returns a boolean if a field has been set.

### GetUrl

`func (o *WebpageActionConfiguration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebpageActionConfiguration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebpageActionConfiguration) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


