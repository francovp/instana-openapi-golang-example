# MobileAppMonitoringBeacon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccuracyRadius** | Pointer to **int64** |  | [optional] 
**AgentVersion** | Pointer to **string** |  | [optional] 
**AppBuild** | Pointer to **string** |  | [optional] 
**AppVersion** | Pointer to **string** |  | [optional] 
**BackendTraceId** | Pointer to **string** |  | [optional] 
**BatchSize** | Pointer to **int64** |  | [optional] 
**BeaconId** | **string** |  | 
**BundleIdentifier** | Pointer to **string** |  | [optional] 
**Carrier** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**ClockSkew** | Pointer to **int64** |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**Continent** | Pointer to **string** |  | [optional] 
**ContinentCode** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**CustomEventName** | Pointer to **string** |  | [optional] 
**DecodedBodySize** | Pointer to **int64** |  | [optional] 
**DeviceHardware** | Pointer to **string** |  | [optional] 
**DeviceManufacturer** | Pointer to **string** |  | [optional] 
**DeviceModel** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**EffectiveConnectionType** | Pointer to **string** |  | [optional] 
**EncodedBodySize** | Pointer to **int64** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**ErrorCount** | Pointer to **int64** |  | [optional] 
**ErrorId** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**ErrorType** | Pointer to **string** |  | [optional] 
**GooglePlayServicesMissing** | Pointer to **bool** |  | [optional] 
**HttpCallHeaders** | Pointer to **map[string]string** |  | [optional] 
**HttpCallMethod** | Pointer to **string** |  | [optional] 
**HttpCallOrigin** | Pointer to **string** |  | [optional] 
**HttpCallPath** | Pointer to **string** |  | [optional] 
**HttpCallStatus** | Pointer to **int32** |  | [optional] 
**HttpCallUrl** | Pointer to **string** |  | [optional] 
**IngestionTime** | Pointer to **int64** |  | [optional] 
**Latitude** | Pointer to **float64** |  | [optional] 
**Longitude** | Pointer to **float64** |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**MobileAppId** | **string** |  | 
**MobileAppLabel** | Pointer to **string** |  | [optional] 
**OsName** | Pointer to **string** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Rooted** | Pointer to **bool** |  | [optional] 
**SessionId** | **string** |  | 
**StackTrace** | Pointer to **string** |  | [optional] 
**Subdivision** | Pointer to **string** |  | [optional] 
**SubdivisionCode** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**TransferSize** | Pointer to **int64** |  | [optional] 
**Type** | **string** |  | 
**Unit** | Pointer to **string** |  | [optional] 
**UserEmail** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**UserIp** | Pointer to **string** |  | [optional] 
**UserLanguages** | Pointer to **[]string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**View** | Pointer to **string** |  | [optional] 
**ViewportHeight** | Pointer to **int32** |  | [optional] 
**ViewportWidth** | Pointer to **int32** |  | [optional] 

## Methods

### NewMobileAppMonitoringBeacon

`func NewMobileAppMonitoringBeacon(beaconId string, mobileAppId string, sessionId string, type_ string, ) *MobileAppMonitoringBeacon`

NewMobileAppMonitoringBeacon instantiates a new MobileAppMonitoringBeacon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileAppMonitoringBeaconWithDefaults

`func NewMobileAppMonitoringBeaconWithDefaults() *MobileAppMonitoringBeacon`

NewMobileAppMonitoringBeaconWithDefaults instantiates a new MobileAppMonitoringBeacon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccuracyRadius

`func (o *MobileAppMonitoringBeacon) GetAccuracyRadius() int64`

GetAccuracyRadius returns the AccuracyRadius field if non-nil, zero value otherwise.

### GetAccuracyRadiusOk

`func (o *MobileAppMonitoringBeacon) GetAccuracyRadiusOk() (*int64, bool)`

GetAccuracyRadiusOk returns a tuple with the AccuracyRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracyRadius

`func (o *MobileAppMonitoringBeacon) SetAccuracyRadius(v int64)`

SetAccuracyRadius sets AccuracyRadius field to given value.

### HasAccuracyRadius

`func (o *MobileAppMonitoringBeacon) HasAccuracyRadius() bool`

HasAccuracyRadius returns a boolean if a field has been set.

### GetAgentVersion

`func (o *MobileAppMonitoringBeacon) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *MobileAppMonitoringBeacon) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *MobileAppMonitoringBeacon) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *MobileAppMonitoringBeacon) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetAppBuild

`func (o *MobileAppMonitoringBeacon) GetAppBuild() string`

GetAppBuild returns the AppBuild field if non-nil, zero value otherwise.

### GetAppBuildOk

`func (o *MobileAppMonitoringBeacon) GetAppBuildOk() (*string, bool)`

GetAppBuildOk returns a tuple with the AppBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppBuild

`func (o *MobileAppMonitoringBeacon) SetAppBuild(v string)`

SetAppBuild sets AppBuild field to given value.

### HasAppBuild

`func (o *MobileAppMonitoringBeacon) HasAppBuild() bool`

HasAppBuild returns a boolean if a field has been set.

### GetAppVersion

`func (o *MobileAppMonitoringBeacon) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *MobileAppMonitoringBeacon) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *MobileAppMonitoringBeacon) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *MobileAppMonitoringBeacon) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetBackendTraceId

`func (o *MobileAppMonitoringBeacon) GetBackendTraceId() string`

GetBackendTraceId returns the BackendTraceId field if non-nil, zero value otherwise.

### GetBackendTraceIdOk

`func (o *MobileAppMonitoringBeacon) GetBackendTraceIdOk() (*string, bool)`

GetBackendTraceIdOk returns a tuple with the BackendTraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendTraceId

`func (o *MobileAppMonitoringBeacon) SetBackendTraceId(v string)`

SetBackendTraceId sets BackendTraceId field to given value.

### HasBackendTraceId

`func (o *MobileAppMonitoringBeacon) HasBackendTraceId() bool`

HasBackendTraceId returns a boolean if a field has been set.

### GetBatchSize

`func (o *MobileAppMonitoringBeacon) GetBatchSize() int64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *MobileAppMonitoringBeacon) GetBatchSizeOk() (*int64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *MobileAppMonitoringBeacon) SetBatchSize(v int64)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *MobileAppMonitoringBeacon) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### GetBeaconId

`func (o *MobileAppMonitoringBeacon) GetBeaconId() string`

GetBeaconId returns the BeaconId field if non-nil, zero value otherwise.

### GetBeaconIdOk

`func (o *MobileAppMonitoringBeacon) GetBeaconIdOk() (*string, bool)`

GetBeaconIdOk returns a tuple with the BeaconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconId

`func (o *MobileAppMonitoringBeacon) SetBeaconId(v string)`

SetBeaconId sets BeaconId field to given value.


### GetBundleIdentifier

`func (o *MobileAppMonitoringBeacon) GetBundleIdentifier() string`

GetBundleIdentifier returns the BundleIdentifier field if non-nil, zero value otherwise.

### GetBundleIdentifierOk

`func (o *MobileAppMonitoringBeacon) GetBundleIdentifierOk() (*string, bool)`

GetBundleIdentifierOk returns a tuple with the BundleIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleIdentifier

`func (o *MobileAppMonitoringBeacon) SetBundleIdentifier(v string)`

SetBundleIdentifier sets BundleIdentifier field to given value.

### HasBundleIdentifier

`func (o *MobileAppMonitoringBeacon) HasBundleIdentifier() bool`

HasBundleIdentifier returns a boolean if a field has been set.

### GetCarrier

`func (o *MobileAppMonitoringBeacon) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *MobileAppMonitoringBeacon) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *MobileAppMonitoringBeacon) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *MobileAppMonitoringBeacon) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCity

`func (o *MobileAppMonitoringBeacon) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *MobileAppMonitoringBeacon) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *MobileAppMonitoringBeacon) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *MobileAppMonitoringBeacon) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetClockSkew

`func (o *MobileAppMonitoringBeacon) GetClockSkew() int64`

GetClockSkew returns the ClockSkew field if non-nil, zero value otherwise.

### GetClockSkewOk

`func (o *MobileAppMonitoringBeacon) GetClockSkewOk() (*int64, bool)`

GetClockSkewOk returns a tuple with the ClockSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockSkew

`func (o *MobileAppMonitoringBeacon) SetClockSkew(v int64)`

SetClockSkew sets ClockSkew field to given value.

### HasClockSkew

`func (o *MobileAppMonitoringBeacon) HasClockSkew() bool`

HasClockSkew returns a boolean if a field has been set.

### GetConnectionType

`func (o *MobileAppMonitoringBeacon) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *MobileAppMonitoringBeacon) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *MobileAppMonitoringBeacon) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *MobileAppMonitoringBeacon) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetContinent

`func (o *MobileAppMonitoringBeacon) GetContinent() string`

GetContinent returns the Continent field if non-nil, zero value otherwise.

### GetContinentOk

`func (o *MobileAppMonitoringBeacon) GetContinentOk() (*string, bool)`

GetContinentOk returns a tuple with the Continent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinent

`func (o *MobileAppMonitoringBeacon) SetContinent(v string)`

SetContinent sets Continent field to given value.

### HasContinent

`func (o *MobileAppMonitoringBeacon) HasContinent() bool`

HasContinent returns a boolean if a field has been set.

### GetContinentCode

`func (o *MobileAppMonitoringBeacon) GetContinentCode() string`

GetContinentCode returns the ContinentCode field if non-nil, zero value otherwise.

### GetContinentCodeOk

`func (o *MobileAppMonitoringBeacon) GetContinentCodeOk() (*string, bool)`

GetContinentCodeOk returns a tuple with the ContinentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinentCode

`func (o *MobileAppMonitoringBeacon) SetContinentCode(v string)`

SetContinentCode sets ContinentCode field to given value.

### HasContinentCode

`func (o *MobileAppMonitoringBeacon) HasContinentCode() bool`

HasContinentCode returns a boolean if a field has been set.

### GetCountry

`func (o *MobileAppMonitoringBeacon) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MobileAppMonitoringBeacon) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *MobileAppMonitoringBeacon) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *MobileAppMonitoringBeacon) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *MobileAppMonitoringBeacon) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *MobileAppMonitoringBeacon) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *MobileAppMonitoringBeacon) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *MobileAppMonitoringBeacon) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCustomEventName

`func (o *MobileAppMonitoringBeacon) GetCustomEventName() string`

GetCustomEventName returns the CustomEventName field if non-nil, zero value otherwise.

### GetCustomEventNameOk

`func (o *MobileAppMonitoringBeacon) GetCustomEventNameOk() (*string, bool)`

GetCustomEventNameOk returns a tuple with the CustomEventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEventName

`func (o *MobileAppMonitoringBeacon) SetCustomEventName(v string)`

SetCustomEventName sets CustomEventName field to given value.

### HasCustomEventName

`func (o *MobileAppMonitoringBeacon) HasCustomEventName() bool`

HasCustomEventName returns a boolean if a field has been set.

### GetDecodedBodySize

`func (o *MobileAppMonitoringBeacon) GetDecodedBodySize() int64`

GetDecodedBodySize returns the DecodedBodySize field if non-nil, zero value otherwise.

### GetDecodedBodySizeOk

`func (o *MobileAppMonitoringBeacon) GetDecodedBodySizeOk() (*int64, bool)`

GetDecodedBodySizeOk returns a tuple with the DecodedBodySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecodedBodySize

`func (o *MobileAppMonitoringBeacon) SetDecodedBodySize(v int64)`

SetDecodedBodySize sets DecodedBodySize field to given value.

### HasDecodedBodySize

`func (o *MobileAppMonitoringBeacon) HasDecodedBodySize() bool`

HasDecodedBodySize returns a boolean if a field has been set.

### GetDeviceHardware

`func (o *MobileAppMonitoringBeacon) GetDeviceHardware() string`

GetDeviceHardware returns the DeviceHardware field if non-nil, zero value otherwise.

### GetDeviceHardwareOk

`func (o *MobileAppMonitoringBeacon) GetDeviceHardwareOk() (*string, bool)`

GetDeviceHardwareOk returns a tuple with the DeviceHardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceHardware

`func (o *MobileAppMonitoringBeacon) SetDeviceHardware(v string)`

SetDeviceHardware sets DeviceHardware field to given value.

### HasDeviceHardware

`func (o *MobileAppMonitoringBeacon) HasDeviceHardware() bool`

HasDeviceHardware returns a boolean if a field has been set.

### GetDeviceManufacturer

`func (o *MobileAppMonitoringBeacon) GetDeviceManufacturer() string`

GetDeviceManufacturer returns the DeviceManufacturer field if non-nil, zero value otherwise.

### GetDeviceManufacturerOk

`func (o *MobileAppMonitoringBeacon) GetDeviceManufacturerOk() (*string, bool)`

GetDeviceManufacturerOk returns a tuple with the DeviceManufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceManufacturer

`func (o *MobileAppMonitoringBeacon) SetDeviceManufacturer(v string)`

SetDeviceManufacturer sets DeviceManufacturer field to given value.

### HasDeviceManufacturer

`func (o *MobileAppMonitoringBeacon) HasDeviceManufacturer() bool`

HasDeviceManufacturer returns a boolean if a field has been set.

### GetDeviceModel

`func (o *MobileAppMonitoringBeacon) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *MobileAppMonitoringBeacon) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *MobileAppMonitoringBeacon) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *MobileAppMonitoringBeacon) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDuration

`func (o *MobileAppMonitoringBeacon) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MobileAppMonitoringBeacon) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MobileAppMonitoringBeacon) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *MobileAppMonitoringBeacon) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEffectiveConnectionType

`func (o *MobileAppMonitoringBeacon) GetEffectiveConnectionType() string`

GetEffectiveConnectionType returns the EffectiveConnectionType field if non-nil, zero value otherwise.

### GetEffectiveConnectionTypeOk

`func (o *MobileAppMonitoringBeacon) GetEffectiveConnectionTypeOk() (*string, bool)`

GetEffectiveConnectionTypeOk returns a tuple with the EffectiveConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveConnectionType

`func (o *MobileAppMonitoringBeacon) SetEffectiveConnectionType(v string)`

SetEffectiveConnectionType sets EffectiveConnectionType field to given value.

### HasEffectiveConnectionType

`func (o *MobileAppMonitoringBeacon) HasEffectiveConnectionType() bool`

HasEffectiveConnectionType returns a boolean if a field has been set.

### GetEncodedBodySize

`func (o *MobileAppMonitoringBeacon) GetEncodedBodySize() int64`

GetEncodedBodySize returns the EncodedBodySize field if non-nil, zero value otherwise.

### GetEncodedBodySizeOk

`func (o *MobileAppMonitoringBeacon) GetEncodedBodySizeOk() (*int64, bool)`

GetEncodedBodySizeOk returns a tuple with the EncodedBodySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedBodySize

`func (o *MobileAppMonitoringBeacon) SetEncodedBodySize(v int64)`

SetEncodedBodySize sets EncodedBodySize field to given value.

### HasEncodedBodySize

`func (o *MobileAppMonitoringBeacon) HasEncodedBodySize() bool`

HasEncodedBodySize returns a boolean if a field has been set.

### GetEnvironment

`func (o *MobileAppMonitoringBeacon) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *MobileAppMonitoringBeacon) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *MobileAppMonitoringBeacon) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *MobileAppMonitoringBeacon) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetErrorCount

`func (o *MobileAppMonitoringBeacon) GetErrorCount() int64`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *MobileAppMonitoringBeacon) GetErrorCountOk() (*int64, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *MobileAppMonitoringBeacon) SetErrorCount(v int64)`

SetErrorCount sets ErrorCount field to given value.

### HasErrorCount

`func (o *MobileAppMonitoringBeacon) HasErrorCount() bool`

HasErrorCount returns a boolean if a field has been set.

### GetErrorId

`func (o *MobileAppMonitoringBeacon) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *MobileAppMonitoringBeacon) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *MobileAppMonitoringBeacon) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *MobileAppMonitoringBeacon) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.

### GetErrorMessage

`func (o *MobileAppMonitoringBeacon) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *MobileAppMonitoringBeacon) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *MobileAppMonitoringBeacon) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *MobileAppMonitoringBeacon) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetErrorType

`func (o *MobileAppMonitoringBeacon) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *MobileAppMonitoringBeacon) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *MobileAppMonitoringBeacon) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *MobileAppMonitoringBeacon) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetGooglePlayServicesMissing

`func (o *MobileAppMonitoringBeacon) GetGooglePlayServicesMissing() bool`

GetGooglePlayServicesMissing returns the GooglePlayServicesMissing field if non-nil, zero value otherwise.

### GetGooglePlayServicesMissingOk

`func (o *MobileAppMonitoringBeacon) GetGooglePlayServicesMissingOk() (*bool, bool)`

GetGooglePlayServicesMissingOk returns a tuple with the GooglePlayServicesMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePlayServicesMissing

`func (o *MobileAppMonitoringBeacon) SetGooglePlayServicesMissing(v bool)`

SetGooglePlayServicesMissing sets GooglePlayServicesMissing field to given value.

### HasGooglePlayServicesMissing

`func (o *MobileAppMonitoringBeacon) HasGooglePlayServicesMissing() bool`

HasGooglePlayServicesMissing returns a boolean if a field has been set.

### GetHttpCallHeaders

`func (o *MobileAppMonitoringBeacon) GetHttpCallHeaders() map[string]string`

GetHttpCallHeaders returns the HttpCallHeaders field if non-nil, zero value otherwise.

### GetHttpCallHeadersOk

`func (o *MobileAppMonitoringBeacon) GetHttpCallHeadersOk() (*map[string]string, bool)`

GetHttpCallHeadersOk returns a tuple with the HttpCallHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCallHeaders

`func (o *MobileAppMonitoringBeacon) SetHttpCallHeaders(v map[string]string)`

SetHttpCallHeaders sets HttpCallHeaders field to given value.

### HasHttpCallHeaders

`func (o *MobileAppMonitoringBeacon) HasHttpCallHeaders() bool`

HasHttpCallHeaders returns a boolean if a field has been set.

### GetHttpCallMethod

`func (o *MobileAppMonitoringBeacon) GetHttpCallMethod() string`

GetHttpCallMethod returns the HttpCallMethod field if non-nil, zero value otherwise.

### GetHttpCallMethodOk

`func (o *MobileAppMonitoringBeacon) GetHttpCallMethodOk() (*string, bool)`

GetHttpCallMethodOk returns a tuple with the HttpCallMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCallMethod

`func (o *MobileAppMonitoringBeacon) SetHttpCallMethod(v string)`

SetHttpCallMethod sets HttpCallMethod field to given value.

### HasHttpCallMethod

`func (o *MobileAppMonitoringBeacon) HasHttpCallMethod() bool`

HasHttpCallMethod returns a boolean if a field has been set.

### GetHttpCallOrigin

`func (o *MobileAppMonitoringBeacon) GetHttpCallOrigin() string`

GetHttpCallOrigin returns the HttpCallOrigin field if non-nil, zero value otherwise.

### GetHttpCallOriginOk

`func (o *MobileAppMonitoringBeacon) GetHttpCallOriginOk() (*string, bool)`

GetHttpCallOriginOk returns a tuple with the HttpCallOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCallOrigin

`func (o *MobileAppMonitoringBeacon) SetHttpCallOrigin(v string)`

SetHttpCallOrigin sets HttpCallOrigin field to given value.

### HasHttpCallOrigin

`func (o *MobileAppMonitoringBeacon) HasHttpCallOrigin() bool`

HasHttpCallOrigin returns a boolean if a field has been set.

### GetHttpCallPath

`func (o *MobileAppMonitoringBeacon) GetHttpCallPath() string`

GetHttpCallPath returns the HttpCallPath field if non-nil, zero value otherwise.

### GetHttpCallPathOk

`func (o *MobileAppMonitoringBeacon) GetHttpCallPathOk() (*string, bool)`

GetHttpCallPathOk returns a tuple with the HttpCallPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCallPath

`func (o *MobileAppMonitoringBeacon) SetHttpCallPath(v string)`

SetHttpCallPath sets HttpCallPath field to given value.

### HasHttpCallPath

`func (o *MobileAppMonitoringBeacon) HasHttpCallPath() bool`

HasHttpCallPath returns a boolean if a field has been set.

### GetHttpCallStatus

`func (o *MobileAppMonitoringBeacon) GetHttpCallStatus() int32`

GetHttpCallStatus returns the HttpCallStatus field if non-nil, zero value otherwise.

### GetHttpCallStatusOk

`func (o *MobileAppMonitoringBeacon) GetHttpCallStatusOk() (*int32, bool)`

GetHttpCallStatusOk returns a tuple with the HttpCallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCallStatus

`func (o *MobileAppMonitoringBeacon) SetHttpCallStatus(v int32)`

SetHttpCallStatus sets HttpCallStatus field to given value.

### HasHttpCallStatus

`func (o *MobileAppMonitoringBeacon) HasHttpCallStatus() bool`

HasHttpCallStatus returns a boolean if a field has been set.

### GetHttpCallUrl

`func (o *MobileAppMonitoringBeacon) GetHttpCallUrl() string`

GetHttpCallUrl returns the HttpCallUrl field if non-nil, zero value otherwise.

### GetHttpCallUrlOk

`func (o *MobileAppMonitoringBeacon) GetHttpCallUrlOk() (*string, bool)`

GetHttpCallUrlOk returns a tuple with the HttpCallUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCallUrl

`func (o *MobileAppMonitoringBeacon) SetHttpCallUrl(v string)`

SetHttpCallUrl sets HttpCallUrl field to given value.

### HasHttpCallUrl

`func (o *MobileAppMonitoringBeacon) HasHttpCallUrl() bool`

HasHttpCallUrl returns a boolean if a field has been set.

### GetIngestionTime

`func (o *MobileAppMonitoringBeacon) GetIngestionTime() int64`

GetIngestionTime returns the IngestionTime field if non-nil, zero value otherwise.

### GetIngestionTimeOk

`func (o *MobileAppMonitoringBeacon) GetIngestionTimeOk() (*int64, bool)`

GetIngestionTimeOk returns a tuple with the IngestionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionTime

`func (o *MobileAppMonitoringBeacon) SetIngestionTime(v int64)`

SetIngestionTime sets IngestionTime field to given value.

### HasIngestionTime

`func (o *MobileAppMonitoringBeacon) HasIngestionTime() bool`

HasIngestionTime returns a boolean if a field has been set.

### GetLatitude

`func (o *MobileAppMonitoringBeacon) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *MobileAppMonitoringBeacon) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *MobileAppMonitoringBeacon) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *MobileAppMonitoringBeacon) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *MobileAppMonitoringBeacon) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *MobileAppMonitoringBeacon) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *MobileAppMonitoringBeacon) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *MobileAppMonitoringBeacon) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetMeta

`func (o *MobileAppMonitoringBeacon) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MobileAppMonitoringBeacon) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MobileAppMonitoringBeacon) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MobileAppMonitoringBeacon) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetMobileAppId

`func (o *MobileAppMonitoringBeacon) GetMobileAppId() string`

GetMobileAppId returns the MobileAppId field if non-nil, zero value otherwise.

### GetMobileAppIdOk

`func (o *MobileAppMonitoringBeacon) GetMobileAppIdOk() (*string, bool)`

GetMobileAppIdOk returns a tuple with the MobileAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileAppId

`func (o *MobileAppMonitoringBeacon) SetMobileAppId(v string)`

SetMobileAppId sets MobileAppId field to given value.


### GetMobileAppLabel

`func (o *MobileAppMonitoringBeacon) GetMobileAppLabel() string`

GetMobileAppLabel returns the MobileAppLabel field if non-nil, zero value otherwise.

### GetMobileAppLabelOk

`func (o *MobileAppMonitoringBeacon) GetMobileAppLabelOk() (*string, bool)`

GetMobileAppLabelOk returns a tuple with the MobileAppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileAppLabel

`func (o *MobileAppMonitoringBeacon) SetMobileAppLabel(v string)`

SetMobileAppLabel sets MobileAppLabel field to given value.

### HasMobileAppLabel

`func (o *MobileAppMonitoringBeacon) HasMobileAppLabel() bool`

HasMobileAppLabel returns a boolean if a field has been set.

### GetOsName

`func (o *MobileAppMonitoringBeacon) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *MobileAppMonitoringBeacon) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *MobileAppMonitoringBeacon) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *MobileAppMonitoringBeacon) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetOsVersion

`func (o *MobileAppMonitoringBeacon) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *MobileAppMonitoringBeacon) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *MobileAppMonitoringBeacon) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *MobileAppMonitoringBeacon) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetPlatform

`func (o *MobileAppMonitoringBeacon) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *MobileAppMonitoringBeacon) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *MobileAppMonitoringBeacon) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *MobileAppMonitoringBeacon) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetRegion

`func (o *MobileAppMonitoringBeacon) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MobileAppMonitoringBeacon) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MobileAppMonitoringBeacon) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MobileAppMonitoringBeacon) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRooted

`func (o *MobileAppMonitoringBeacon) GetRooted() bool`

GetRooted returns the Rooted field if non-nil, zero value otherwise.

### GetRootedOk

`func (o *MobileAppMonitoringBeacon) GetRootedOk() (*bool, bool)`

GetRootedOk returns a tuple with the Rooted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooted

`func (o *MobileAppMonitoringBeacon) SetRooted(v bool)`

SetRooted sets Rooted field to given value.

### HasRooted

`func (o *MobileAppMonitoringBeacon) HasRooted() bool`

HasRooted returns a boolean if a field has been set.

### GetSessionId

`func (o *MobileAppMonitoringBeacon) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *MobileAppMonitoringBeacon) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *MobileAppMonitoringBeacon) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetStackTrace

`func (o *MobileAppMonitoringBeacon) GetStackTrace() string`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *MobileAppMonitoringBeacon) GetStackTraceOk() (*string, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *MobileAppMonitoringBeacon) SetStackTrace(v string)`

SetStackTrace sets StackTrace field to given value.

### HasStackTrace

`func (o *MobileAppMonitoringBeacon) HasStackTrace() bool`

HasStackTrace returns a boolean if a field has been set.

### GetSubdivision

`func (o *MobileAppMonitoringBeacon) GetSubdivision() string`

GetSubdivision returns the Subdivision field if non-nil, zero value otherwise.

### GetSubdivisionOk

`func (o *MobileAppMonitoringBeacon) GetSubdivisionOk() (*string, bool)`

GetSubdivisionOk returns a tuple with the Subdivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdivision

`func (o *MobileAppMonitoringBeacon) SetSubdivision(v string)`

SetSubdivision sets Subdivision field to given value.

### HasSubdivision

`func (o *MobileAppMonitoringBeacon) HasSubdivision() bool`

HasSubdivision returns a boolean if a field has been set.

### GetSubdivisionCode

`func (o *MobileAppMonitoringBeacon) GetSubdivisionCode() string`

GetSubdivisionCode returns the SubdivisionCode field if non-nil, zero value otherwise.

### GetSubdivisionCodeOk

`func (o *MobileAppMonitoringBeacon) GetSubdivisionCodeOk() (*string, bool)`

GetSubdivisionCodeOk returns a tuple with the SubdivisionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdivisionCode

`func (o *MobileAppMonitoringBeacon) SetSubdivisionCode(v string)`

SetSubdivisionCode sets SubdivisionCode field to given value.

### HasSubdivisionCode

`func (o *MobileAppMonitoringBeacon) HasSubdivisionCode() bool`

HasSubdivisionCode returns a boolean if a field has been set.

### GetTenant

`func (o *MobileAppMonitoringBeacon) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *MobileAppMonitoringBeacon) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *MobileAppMonitoringBeacon) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *MobileAppMonitoringBeacon) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetTimestamp

`func (o *MobileAppMonitoringBeacon) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MobileAppMonitoringBeacon) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MobileAppMonitoringBeacon) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MobileAppMonitoringBeacon) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTransferSize

`func (o *MobileAppMonitoringBeacon) GetTransferSize() int64`

GetTransferSize returns the TransferSize field if non-nil, zero value otherwise.

### GetTransferSizeOk

`func (o *MobileAppMonitoringBeacon) GetTransferSizeOk() (*int64, bool)`

GetTransferSizeOk returns a tuple with the TransferSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferSize

`func (o *MobileAppMonitoringBeacon) SetTransferSize(v int64)`

SetTransferSize sets TransferSize field to given value.

### HasTransferSize

`func (o *MobileAppMonitoringBeacon) HasTransferSize() bool`

HasTransferSize returns a boolean if a field has been set.

### GetType

`func (o *MobileAppMonitoringBeacon) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MobileAppMonitoringBeacon) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MobileAppMonitoringBeacon) SetType(v string)`

SetType sets Type field to given value.


### GetUnit

`func (o *MobileAppMonitoringBeacon) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MobileAppMonitoringBeacon) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MobileAppMonitoringBeacon) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MobileAppMonitoringBeacon) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUserEmail

`func (o *MobileAppMonitoringBeacon) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *MobileAppMonitoringBeacon) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *MobileAppMonitoringBeacon) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *MobileAppMonitoringBeacon) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUserId

`func (o *MobileAppMonitoringBeacon) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MobileAppMonitoringBeacon) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MobileAppMonitoringBeacon) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MobileAppMonitoringBeacon) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserIp

`func (o *MobileAppMonitoringBeacon) GetUserIp() string`

GetUserIp returns the UserIp field if non-nil, zero value otherwise.

### GetUserIpOk

`func (o *MobileAppMonitoringBeacon) GetUserIpOk() (*string, bool)`

GetUserIpOk returns a tuple with the UserIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIp

`func (o *MobileAppMonitoringBeacon) SetUserIp(v string)`

SetUserIp sets UserIp field to given value.

### HasUserIp

`func (o *MobileAppMonitoringBeacon) HasUserIp() bool`

HasUserIp returns a boolean if a field has been set.

### GetUserLanguages

`func (o *MobileAppMonitoringBeacon) GetUserLanguages() []string`

GetUserLanguages returns the UserLanguages field if non-nil, zero value otherwise.

### GetUserLanguagesOk

`func (o *MobileAppMonitoringBeacon) GetUserLanguagesOk() (*[]string, bool)`

GetUserLanguagesOk returns a tuple with the UserLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLanguages

`func (o *MobileAppMonitoringBeacon) SetUserLanguages(v []string)`

SetUserLanguages sets UserLanguages field to given value.

### HasUserLanguages

`func (o *MobileAppMonitoringBeacon) HasUserLanguages() bool`

HasUserLanguages returns a boolean if a field has been set.

### GetUserName

`func (o *MobileAppMonitoringBeacon) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *MobileAppMonitoringBeacon) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *MobileAppMonitoringBeacon) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *MobileAppMonitoringBeacon) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetView

`func (o *MobileAppMonitoringBeacon) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *MobileAppMonitoringBeacon) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *MobileAppMonitoringBeacon) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *MobileAppMonitoringBeacon) HasView() bool`

HasView returns a boolean if a field has been set.

### GetViewportHeight

`func (o *MobileAppMonitoringBeacon) GetViewportHeight() int32`

GetViewportHeight returns the ViewportHeight field if non-nil, zero value otherwise.

### GetViewportHeightOk

`func (o *MobileAppMonitoringBeacon) GetViewportHeightOk() (*int32, bool)`

GetViewportHeightOk returns a tuple with the ViewportHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewportHeight

`func (o *MobileAppMonitoringBeacon) SetViewportHeight(v int32)`

SetViewportHeight sets ViewportHeight field to given value.

### HasViewportHeight

`func (o *MobileAppMonitoringBeacon) HasViewportHeight() bool`

HasViewportHeight returns a boolean if a field has been set.

### GetViewportWidth

`func (o *MobileAppMonitoringBeacon) GetViewportWidth() int32`

GetViewportWidth returns the ViewportWidth field if non-nil, zero value otherwise.

### GetViewportWidthOk

`func (o *MobileAppMonitoringBeacon) GetViewportWidthOk() (*int32, bool)`

GetViewportWidthOk returns a tuple with the ViewportWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewportWidth

`func (o *MobileAppMonitoringBeacon) SetViewportWidth(v int32)`

SetViewportWidth sets ViewportWidth field to given value.

### HasViewportWidth

`func (o *MobileAppMonitoringBeacon) HasViewportWidth() bool`

HasViewportWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


