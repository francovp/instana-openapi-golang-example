# SyntheticTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** |  | 
**ApplicationId** | Pointer to **string** |  | [optional] 
**ApplicationLabel** | Pointer to **string** |  | [optional] 
**Configuration** | [**SyntheticTypeConfiguration**](SyntheticTypeConfiguration.md) |  | 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CustomProperties** | Pointer to **map[string]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Label** | **string** |  | 
**LocationDisplayLabels** | Pointer to **[]string** |  | [optional] 
**LocationLabels** | Pointer to **[]string** |  | [optional] 
**Locations** | **[]string** |  | 
**ModifiedAt** | Pointer to **int64** |  | [optional] 
**ModifiedBy** | Pointer to **string** |  | [optional] 
**PlaybackMode** | **string** |  | 
**TenantId** | Pointer to **string** |  | [optional] 
**TestFrequency** | **int32** |  | 

## Methods

### NewSyntheticTest

`func NewSyntheticTest(active bool, configuration SyntheticTypeConfiguration, label string, locations []string, playbackMode string, testFrequency int32, ) *SyntheticTest`

NewSyntheticTest instantiates a new SyntheticTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticTestWithDefaults

`func NewSyntheticTestWithDefaults() *SyntheticTest`

NewSyntheticTestWithDefaults instantiates a new SyntheticTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *SyntheticTest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SyntheticTest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SyntheticTest) SetActive(v bool)`

SetActive sets Active field to given value.


### GetApplicationId

`func (o *SyntheticTest) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *SyntheticTest) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *SyntheticTest) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *SyntheticTest) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationLabel

`func (o *SyntheticTest) GetApplicationLabel() string`

GetApplicationLabel returns the ApplicationLabel field if non-nil, zero value otherwise.

### GetApplicationLabelOk

`func (o *SyntheticTest) GetApplicationLabelOk() (*string, bool)`

GetApplicationLabelOk returns a tuple with the ApplicationLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationLabel

`func (o *SyntheticTest) SetApplicationLabel(v string)`

SetApplicationLabel sets ApplicationLabel field to given value.

### HasApplicationLabel

`func (o *SyntheticTest) HasApplicationLabel() bool`

HasApplicationLabel returns a boolean if a field has been set.

### GetConfiguration

`func (o *SyntheticTest) GetConfiguration() SyntheticTypeConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *SyntheticTest) GetConfigurationOk() (*SyntheticTypeConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *SyntheticTest) SetConfiguration(v SyntheticTypeConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetCreatedAt

`func (o *SyntheticTest) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SyntheticTest) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SyntheticTest) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SyntheticTest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SyntheticTest) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SyntheticTest) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SyntheticTest) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SyntheticTest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCustomProperties

`func (o *SyntheticTest) GetCustomProperties() map[string]string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *SyntheticTest) GetCustomPropertiesOk() (*map[string]string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *SyntheticTest) SetCustomProperties(v map[string]string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *SyntheticTest) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *SyntheticTest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SyntheticTest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SyntheticTest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SyntheticTest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *SyntheticTest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyntheticTest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyntheticTest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SyntheticTest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *SyntheticTest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SyntheticTest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SyntheticTest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLocationDisplayLabels

`func (o *SyntheticTest) GetLocationDisplayLabels() []string`

GetLocationDisplayLabels returns the LocationDisplayLabels field if non-nil, zero value otherwise.

### GetLocationDisplayLabelsOk

`func (o *SyntheticTest) GetLocationDisplayLabelsOk() (*[]string, bool)`

GetLocationDisplayLabelsOk returns a tuple with the LocationDisplayLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationDisplayLabels

`func (o *SyntheticTest) SetLocationDisplayLabels(v []string)`

SetLocationDisplayLabels sets LocationDisplayLabels field to given value.

### HasLocationDisplayLabels

`func (o *SyntheticTest) HasLocationDisplayLabels() bool`

HasLocationDisplayLabels returns a boolean if a field has been set.

### GetLocationLabels

`func (o *SyntheticTest) GetLocationLabels() []string`

GetLocationLabels returns the LocationLabels field if non-nil, zero value otherwise.

### GetLocationLabelsOk

`func (o *SyntheticTest) GetLocationLabelsOk() (*[]string, bool)`

GetLocationLabelsOk returns a tuple with the LocationLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationLabels

`func (o *SyntheticTest) SetLocationLabels(v []string)`

SetLocationLabels sets LocationLabels field to given value.

### HasLocationLabels

`func (o *SyntheticTest) HasLocationLabels() bool`

HasLocationLabels returns a boolean if a field has been set.

### GetLocations

`func (o *SyntheticTest) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *SyntheticTest) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *SyntheticTest) SetLocations(v []string)`

SetLocations sets Locations field to given value.


### GetModifiedAt

`func (o *SyntheticTest) GetModifiedAt() int64`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SyntheticTest) GetModifiedAtOk() (*int64, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SyntheticTest) SetModifiedAt(v int64)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *SyntheticTest) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *SyntheticTest) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *SyntheticTest) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *SyntheticTest) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *SyntheticTest) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetPlaybackMode

`func (o *SyntheticTest) GetPlaybackMode() string`

GetPlaybackMode returns the PlaybackMode field if non-nil, zero value otherwise.

### GetPlaybackModeOk

`func (o *SyntheticTest) GetPlaybackModeOk() (*string, bool)`

GetPlaybackModeOk returns a tuple with the PlaybackMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackMode

`func (o *SyntheticTest) SetPlaybackMode(v string)`

SetPlaybackMode sets PlaybackMode field to given value.


### GetTenantId

`func (o *SyntheticTest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SyntheticTest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SyntheticTest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SyntheticTest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTestFrequency

`func (o *SyntheticTest) GetTestFrequency() int32`

GetTestFrequency returns the TestFrequency field if non-nil, zero value otherwise.

### GetTestFrequencyOk

`func (o *SyntheticTest) GetTestFrequencyOk() (*int32, bool)`

GetTestFrequencyOk returns a tuple with the TestFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestFrequency

`func (o *SyntheticTest) SetTestFrequency(v int32)`

SetTestFrequency sets TestFrequency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


