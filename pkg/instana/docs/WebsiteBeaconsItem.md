# WebsiteBeaconsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Beacon** | [**WebsiteMonitoringBeacon**](WebsiteMonitoringBeacon.md) |  | 
**Cursor** | [**IngestionOffsetCursor**](IngestionOffsetCursor.md) |  | 

## Methods

### NewWebsiteBeaconsItem

`func NewWebsiteBeaconsItem(beacon WebsiteMonitoringBeacon, cursor IngestionOffsetCursor, ) *WebsiteBeaconsItem`

NewWebsiteBeaconsItem instantiates a new WebsiteBeaconsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteBeaconsItemWithDefaults

`func NewWebsiteBeaconsItemWithDefaults() *WebsiteBeaconsItem`

NewWebsiteBeaconsItemWithDefaults instantiates a new WebsiteBeaconsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeacon

`func (o *WebsiteBeaconsItem) GetBeacon() WebsiteMonitoringBeacon`

GetBeacon returns the Beacon field if non-nil, zero value otherwise.

### GetBeaconOk

`func (o *WebsiteBeaconsItem) GetBeaconOk() (*WebsiteMonitoringBeacon, bool)`

GetBeaconOk returns a tuple with the Beacon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeacon

`func (o *WebsiteBeaconsItem) SetBeacon(v WebsiteMonitoringBeacon)`

SetBeacon sets Beacon field to given value.


### GetCursor

`func (o *WebsiteBeaconsItem) GetCursor() IngestionOffsetCursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *WebsiteBeaconsItem) GetCursorOk() (*IngestionOffsetCursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *WebsiteBeaconsItem) SetCursor(v IngestionOffsetCursor)`

SetCursor sets Cursor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


