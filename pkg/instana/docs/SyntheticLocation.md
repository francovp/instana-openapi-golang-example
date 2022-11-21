# SyntheticLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** |  | [optional] 
**CustomProperties** | Pointer to **map[string]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayLabel** | Pointer to **string** |  | [optional] 
**GeoPoint** | [**SyntheticGeoPoint**](SyntheticGeoPoint.md) |  | 
**Id** | Pointer to **string** |  | [optional] 
**Label** | **string** |  | 
**LocationType** | **string** |  | 
**ModifiedAt** | Pointer to **int64** |  | [optional] 
**ObservedAt** | Pointer to **int64** |  | [optional] 
**PlaybackCapabilities** | [**SyntheticPlaybackCapabilities**](SyntheticPlaybackCapabilities.md) |  | 
**PopVersion** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TotalTests** | Pointer to **int32** |  | [optional] 

## Methods

### NewSyntheticLocation

`func NewSyntheticLocation(geoPoint SyntheticGeoPoint, label string, locationType string, playbackCapabilities SyntheticPlaybackCapabilities, ) *SyntheticLocation`

NewSyntheticLocation instantiates a new SyntheticLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticLocationWithDefaults

`func NewSyntheticLocationWithDefaults() *SyntheticLocation`

NewSyntheticLocationWithDefaults instantiates a new SyntheticLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SyntheticLocation) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SyntheticLocation) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SyntheticLocation) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SyntheticLocation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomProperties

`func (o *SyntheticLocation) GetCustomProperties() map[string]string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *SyntheticLocation) GetCustomPropertiesOk() (*map[string]string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *SyntheticLocation) SetCustomProperties(v map[string]string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *SyntheticLocation) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *SyntheticLocation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SyntheticLocation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SyntheticLocation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SyntheticLocation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *SyntheticLocation) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *SyntheticLocation) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *SyntheticLocation) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *SyntheticLocation) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetGeoPoint

`func (o *SyntheticLocation) GetGeoPoint() SyntheticGeoPoint`

GetGeoPoint returns the GeoPoint field if non-nil, zero value otherwise.

### GetGeoPointOk

`func (o *SyntheticLocation) GetGeoPointOk() (*SyntheticGeoPoint, bool)`

GetGeoPointOk returns a tuple with the GeoPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoPoint

`func (o *SyntheticLocation) SetGeoPoint(v SyntheticGeoPoint)`

SetGeoPoint sets GeoPoint field to given value.


### GetId

`func (o *SyntheticLocation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyntheticLocation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyntheticLocation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SyntheticLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *SyntheticLocation) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SyntheticLocation) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SyntheticLocation) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLocationType

`func (o *SyntheticLocation) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *SyntheticLocation) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *SyntheticLocation) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.


### GetModifiedAt

`func (o *SyntheticLocation) GetModifiedAt() int64`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SyntheticLocation) GetModifiedAtOk() (*int64, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SyntheticLocation) SetModifiedAt(v int64)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *SyntheticLocation) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetObservedAt

`func (o *SyntheticLocation) GetObservedAt() int64`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *SyntheticLocation) GetObservedAtOk() (*int64, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *SyntheticLocation) SetObservedAt(v int64)`

SetObservedAt sets ObservedAt field to given value.

### HasObservedAt

`func (o *SyntheticLocation) HasObservedAt() bool`

HasObservedAt returns a boolean if a field has been set.

### GetPlaybackCapabilities

`func (o *SyntheticLocation) GetPlaybackCapabilities() SyntheticPlaybackCapabilities`

GetPlaybackCapabilities returns the PlaybackCapabilities field if non-nil, zero value otherwise.

### GetPlaybackCapabilitiesOk

`func (o *SyntheticLocation) GetPlaybackCapabilitiesOk() (*SyntheticPlaybackCapabilities, bool)`

GetPlaybackCapabilitiesOk returns a tuple with the PlaybackCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackCapabilities

`func (o *SyntheticLocation) SetPlaybackCapabilities(v SyntheticPlaybackCapabilities)`

SetPlaybackCapabilities sets PlaybackCapabilities field to given value.


### GetPopVersion

`func (o *SyntheticLocation) GetPopVersion() string`

GetPopVersion returns the PopVersion field if non-nil, zero value otherwise.

### GetPopVersionOk

`func (o *SyntheticLocation) GetPopVersionOk() (*string, bool)`

GetPopVersionOk returns a tuple with the PopVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopVersion

`func (o *SyntheticLocation) SetPopVersion(v string)`

SetPopVersion sets PopVersion field to given value.

### HasPopVersion

`func (o *SyntheticLocation) HasPopVersion() bool`

HasPopVersion returns a boolean if a field has been set.

### GetStatus

`func (o *SyntheticLocation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyntheticLocation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyntheticLocation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyntheticLocation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalTests

`func (o *SyntheticLocation) GetTotalTests() int32`

GetTotalTests returns the TotalTests field if non-nil, zero value otherwise.

### GetTotalTestsOk

`func (o *SyntheticLocation) GetTotalTestsOk() (*int32, bool)`

GetTotalTestsOk returns a tuple with the TotalTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTests

`func (o *SyntheticLocation) SetTotalTests(v int32)`

SetTotalTests sets TotalTests field to given value.

### HasTotalTests

`func (o *SyntheticLocation) HasTotalTests() bool`

HasTotalTests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


