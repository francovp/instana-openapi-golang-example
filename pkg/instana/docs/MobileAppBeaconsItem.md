# MobileAppBeaconsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Beacon** | [**MobileAppMonitoringBeacon**](MobileAppMonitoringBeacon.md) |  | 
**Cursor** | [**IngestionOffsetCursor**](IngestionOffsetCursor.md) |  | 

## Methods

### NewMobileAppBeaconsItem

`func NewMobileAppBeaconsItem(beacon MobileAppMonitoringBeacon, cursor IngestionOffsetCursor, ) *MobileAppBeaconsItem`

NewMobileAppBeaconsItem instantiates a new MobileAppBeaconsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileAppBeaconsItemWithDefaults

`func NewMobileAppBeaconsItemWithDefaults() *MobileAppBeaconsItem`

NewMobileAppBeaconsItemWithDefaults instantiates a new MobileAppBeaconsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeacon

`func (o *MobileAppBeaconsItem) GetBeacon() MobileAppMonitoringBeacon`

GetBeacon returns the Beacon field if non-nil, zero value otherwise.

### GetBeaconOk

`func (o *MobileAppBeaconsItem) GetBeaconOk() (*MobileAppMonitoringBeacon, bool)`

GetBeaconOk returns a tuple with the Beacon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeacon

`func (o *MobileAppBeaconsItem) SetBeacon(v MobileAppMonitoringBeacon)`

SetBeacon sets Beacon field to given value.


### GetCursor

`func (o *MobileAppBeaconsItem) GetCursor() IngestionOffsetCursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *MobileAppBeaconsItem) GetCursorOk() (*IngestionOffsetCursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *MobileAppBeaconsItem) SetCursor(v IngestionOffsetCursor)`

SetCursor sets Cursor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


