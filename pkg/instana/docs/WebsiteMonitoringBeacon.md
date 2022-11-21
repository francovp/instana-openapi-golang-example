# WebsiteMonitoringBeacon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccuracyRadius** | Pointer to **int64** |  | [optional] 
**AccurateTimingsAvailable** | Pointer to **bool** |  | [optional] 
**AppCacheTime** | Pointer to **int64** |  | [optional] 
**BackendTime** | Pointer to **int64** |  | [optional] 
**BackendTraceId** | Pointer to **string** |  | [optional] 
**BatchSize** | Pointer to **int64** |  | [optional] 
**BeaconId** | **string** |  | 
**BrowserName** | Pointer to **string** |  | [optional] 
**BrowserVersion** | Pointer to **string** |  | [optional] 
**CacheInteraction** | Pointer to **string** |  | [optional] 
**ChildrenTime** | Pointer to **int64** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**ClockSkew** | Pointer to **int64** |  | [optional] 
**ComponentStack** | Pointer to **string** |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**Continent** | Pointer to **string** |  | [optional] 
**ContinentCode** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**CspBlockedUri** | Pointer to **string** |  | [optional] 
**CspColumnNumber** | Pointer to **int64** |  | [optional] 
**CspDisposition** | Pointer to **string** |  | [optional] 
**CspEffectiveDirective** | Pointer to **string** |  | [optional] 
**CspLineNumber** | Pointer to **int64** |  | [optional] 
**CspOriginalPolicy** | Pointer to **string** |  | [optional] 
**CspSample** | Pointer to **string** |  | [optional] 
**CspSourceFile** | Pointer to **string** |  | [optional] 
**CumulativeLayoutShift** | Pointer to **float64** |  | [optional] 
**CustomEventName** | Pointer to **string** |  | [optional] 
**DecodedBodySize** | Pointer to **int64** |  | [optional] 
**Deprecations** | Pointer to **[]string** |  | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 
**DnsTime** | Pointer to **int64** |  | [optional] 
**DomTime** | Pointer to **int64** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**EncodedBodySize** | Pointer to **int64** |  | [optional] 
**ErrorCount** | Pointer to **int64** |  | [optional] 
**ErrorId** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**ErrorType** | Pointer to **string** |  | [optional] 
**FirstContentfulPaintTime** | Pointer to **int64** |  | [optional] 
**FirstInputDelayTime** | Pointer to **int64** |  | [optional] 
**FirstPaintTime** | Pointer to **int64** |  | [optional] 
**FrontendTime** | Pointer to **int64** |  | [optional] 
**GraphqlOperationName** | Pointer to **string** |  | [optional] 
**GraphqlOperationType** | Pointer to **string** |  | [optional] 
**HttpCallAsynchronous** | Pointer to **bool** |  | [optional] 
**HttpCallCorrelationAttempted** | Pointer to **bool** |  | [optional] 
**HttpCallHeaders** | Pointer to **map[string]string** |  | [optional] 
**HttpCallMethod** | Pointer to **string** |  | [optional] 
**HttpCallOrigin** | Pointer to **string** |  | [optional] 
**HttpCallPath** | Pointer to **string** |  | [optional] 
**HttpCallStatus** | Pointer to **int32** |  | [optional] 
**HttpCallUrl** | Pointer to **string** |  | [optional] 
**Initiator** | Pointer to **string** |  | [optional] 
**LargestContentfulPaintTime** | Pointer to **int64** |  | [optional] 
**Latitude** | Pointer to **float64** |  | [optional] 
**LocationOrigin** | **string** |  | 
**LocationPath** | Pointer to **string** |  | [optional] 
**LocationUrl** | **string** |  | 
**Longitude** | Pointer to **float64** |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**OnLoadTime** | Pointer to **int64** |  | [optional] 
**OsName** | Pointer to **string** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **string** |  | [optional] 
**PageLoadId** | **string** |  | 
**ParsedStackTrace** | Pointer to [**[]StackTraceLine**](StackTraceLine.md) |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**ProcessingTime** | Pointer to **int64** |  | [optional] 
**RedirectTime** | Pointer to **int64** |  | [optional] 
**RequestTime** | Pointer to **int64** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**ResponseTime** | Pointer to **int64** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**SnippetVersion** | Pointer to **string** |  | [optional] 
**SslTime** | Pointer to **int64** |  | [optional] 
**StackTrace** | Pointer to **string** |  | [optional] 
**StackTraceParsingStatus** | Pointer to **int32** |  | [optional] 
**StackTraceReadability** | Pointer to **int32** |  | [optional] 
**Subdivision** | Pointer to **string** |  | [optional] 
**SubdivisionCode** | Pointer to **string** |  | [optional] 
**TcpTime** | Pointer to **int64** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**TransferSize** | Pointer to **int64** |  | [optional] 
**Type** | **string** |  | 
**UnloadTime** | Pointer to **int64** |  | [optional] 
**UserEmail** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**UserIp** | Pointer to **string** |  | [optional] 
**UserLanguages** | Pointer to **[]string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**WebsiteId** | **string** |  | 
**WebsiteLabel** | **string** |  | 
**WindowHeight** | Pointer to **int32** |  | [optional] 
**WindowHidden** | Pointer to **bool** |  | [optional] 
**WindowWidth** | Pointer to **int32** |  | [optional] 

## Methods

### NewWebsiteMonitoringBeacon

`func NewWebsiteMonitoringBeacon(beaconId string, locationOrigin string, locationUrl string, pageLoadId string, type_ string, websiteId string, websiteLabel string, ) *WebsiteMonitoringBeacon`

NewWebsiteMonitoringBeacon instantiates a new WebsiteMonitoringBeacon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteMonitoringBeaconWithDefaults

`func NewWebsiteMonitoringBeaconWithDefaults() *WebsiteMonitoringBeacon`

NewWebsiteMonitoringBeaconWithDefaults instantiates a new WebsiteMonitoringBeacon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccuracyRadius

`func (o *WebsiteMonitoringBeacon) GetAccuracyRadius() int64`

GetAccuracyRadius returns the AccuracyRadius field if non-nil, zero value otherwise.

### GetAccuracyRadiusOk

`func (o *WebsiteMonitoringBeacon) GetAccuracyRadiusOk() (*int64, bool)`

GetAccuracyRadiusOk returns a tuple with the AccuracyRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracyRadius

`func (o *WebsiteMonitoringBeacon) SetAccuracyRadius(v int64)`

SetAccuracyRadius sets AccuracyRadius field to given value.

### HasAccuracyRadius

`func (o *WebsiteMonitoringBeacon) HasAccuracyRadius() bool`

HasAccuracyRadius returns a boolean if a field has been set.

### GetAccurateTimingsAvailable

`func (o *WebsiteMonitoringBeacon) GetAccurateTimingsAvailable() bool`

GetAccurateTimingsAvailable returns the AccurateTimingsAvailable field if non-nil, zero value otherwise.

### GetAccurateTimingsAvailableOk

`func (o *WebsiteMonitoringBeacon) GetAccurateTimingsAvailableOk() (*bool, bool)`

GetAccurateTimingsAvailableOk returns a tuple with the AccurateTimingsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccurateTimingsAvailable

`func (o *WebsiteMonitoringBeacon) SetAccurateTimingsAvailable(v bool)`

SetAccurateTimingsAvailable sets AccurateTimingsAvailable field to given value.

### HasAccurateTimingsAvailable

`func (o *WebsiteMonitoringBeacon) HasAccurateTimingsAvailable() bool`

HasAccurateTimingsAvailable returns a boolean if a field has been set.

### GetAppCacheTime

`func (o *WebsiteMonitoringBeacon) GetAppCacheTime() int64`

GetAppCacheTime returns the AppCacheTime field if non-nil, zero value otherwise.

### GetAppCacheTimeOk

`func (o *WebsiteMonitoringBeacon) GetAppCacheTimeOk() (*int64, bool)`

GetAppCacheTimeOk returns a tuple with the AppCacheTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCacheTime

`func (o *WebsiteMonitoringBeacon) SetAppCacheTime(v int64)`

SetAppCacheTime sets AppCacheTime field to given value.

### HasAppCacheTime

`func (o *WebsiteMonitoringBeacon) HasAppCacheTime() bool`

HasAppCacheTime returns a boolean if a field has been set.

### GetBackendTime

`func (o *WebsiteMonitoringBeacon) GetBackendTime() int64`

GetBackendTime returns the BackendTime field if non-nil, zero value otherwise.

### GetBackendTimeOk

`func (o *WebsiteMonitoringBeacon) GetBackendTimeOk() (*int64, bool)`

GetBackendTimeOk returns a tuple with the BackendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendTime

`func (o *WebsiteMonitoringBeacon) SetBackendTime(v int64)`

SetBackendTime sets BackendTime field to given value.

### HasBackendTime

`func (o *WebsiteMonitoringBeacon) HasBackendTime() bool`

HasBackendTime returns a boolean if a field has been set.

### GetBackendTraceId

`func (o *WebsiteMonitoringBeacon) GetBackendTraceId() string`

GetBackendTraceId returns the BackendTraceId field if non-nil, zero value otherwise.

### GetBackendTraceIdOk

`func (o *WebsiteMonitoringBeacon) GetBackendTraceIdOk() (*string, bool)`

GetBackendTraceIdOk returns a tuple with the BackendTraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendTraceId

`func (o *WebsiteMonitoringBeacon) SetBackendTraceId(v string)`

SetBackendTraceId sets BackendTraceId field to given value.

### HasBackendTraceId

`func (o *WebsiteMonitoringBeacon) HasBackendTraceId() bool`

HasBackendTraceId returns a boolean if a field has been set.

### GetBatchSize

`func (o *WebsiteMonitoringBeacon) GetBatchSize() int64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *WebsiteMonitoringBeacon) GetBatchSizeOk() (*int64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *WebsiteMonitoringBeacon) SetBatchSize(v int64)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *WebsiteMonitoringBeacon) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### GetBeaconId

`func (o *WebsiteMonitoringBeacon) GetBeaconId() string`

GetBeaconId returns the BeaconId field if non-nil, zero value otherwise.

### GetBeaconIdOk

`func (o *WebsiteMonitoringBeacon) GetBeaconIdOk() (*string, bool)`

GetBeaconIdOk returns a tuple with the BeaconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconId

`func (o *WebsiteMonitoringBeacon) SetBeaconId(v string)`

SetBeaconId sets BeaconId field to given value.


### GetBrowserName

`func (o *WebsiteMonitoringBeacon) GetBrowserName() string`

GetBrowserName returns the BrowserName field if non-nil, zero value otherwise.

### GetBrowserNameOk

`func (o *WebsiteMonitoringBeacon) GetBrowserNameOk() (*string, bool)`

GetBrowserNameOk returns a tuple with the BrowserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserName

`func (o *WebsiteMonitoringBeacon) SetBrowserName(v string)`

SetBrowserName sets BrowserName field to given value.

### HasBrowserName

`func (o *WebsiteMonitoringBeacon) HasBrowserName() bool`

HasBrowserName returns a boolean if a field has been set.

### GetBrowserVersion

`func (o *WebsiteMonitoringBeacon) GetBrowserVersion() string`

GetBrowserVersion returns the BrowserVersion field if non-nil, zero value otherwise.

### GetBrowserVersionOk

`func (o *WebsiteMonitoringBeacon) GetBrowserVersionOk() (*string, bool)`

GetBrowserVersionOk returns a tuple with the BrowserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserVersion

`func (o *WebsiteMonitoringBeacon) SetBrowserVersion(v string)`

SetBrowserVersion sets BrowserVersion field to given value.

### HasBrowserVersion

`func (o *WebsiteMonitoringBeacon) HasBrowserVersion() bool`

HasBrowserVersion returns a boolean if a field has been set.

### GetCacheInteraction

`func (o *WebsiteMonitoringBeacon) GetCacheInteraction() string`

GetCacheInteraction returns the CacheInteraction field if non-nil, zero value otherwise.

### GetCacheInteractionOk

`func (o *WebsiteMonitoringBeacon) GetCacheInteractionOk() (*string, bool)`

GetCacheInteractionOk returns a tuple with the CacheInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheInteraction

`func (o *WebsiteMonitoringBeacon) SetCacheInteraction(v string)`

SetCacheInteraction sets CacheInteraction field to given value.

### HasCacheInteraction

`func (o *WebsiteMonitoringBeacon) HasCacheInteraction() bool`

HasCacheInteraction returns a boolean if a field has been set.

### GetChildrenTime

`func (o *WebsiteMonitoringBeacon) GetChildrenTime() int64`

GetChildrenTime returns the ChildrenTime field if non-nil, zero value otherwise.

### GetChildrenTimeOk

`func (o *WebsiteMonitoringBeacon) GetChildrenTimeOk() (*int64, bool)`

GetChildrenTimeOk returns a tuple with the ChildrenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenTime

`func (o *WebsiteMonitoringBeacon) SetChildrenTime(v int64)`

SetChildrenTime sets ChildrenTime field to given value.

### HasChildrenTime

`func (o *WebsiteMonitoringBeacon) HasChildrenTime() bool`

HasChildrenTime returns a boolean if a field has been set.

### GetCity

`func (o *WebsiteMonitoringBeacon) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *WebsiteMonitoringBeacon) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *WebsiteMonitoringBeacon) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *WebsiteMonitoringBeacon) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetClockSkew

`func (o *WebsiteMonitoringBeacon) GetClockSkew() int64`

GetClockSkew returns the ClockSkew field if non-nil, zero value otherwise.

### GetClockSkewOk

`func (o *WebsiteMonitoringBeacon) GetClockSkewOk() (*int64, bool)`

GetClockSkewOk returns a tuple with the ClockSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockSkew

`func (o *WebsiteMonitoringBeacon) SetClockSkew(v int64)`

SetClockSkew sets ClockSkew field to given value.

### HasClockSkew

`func (o *WebsiteMonitoringBeacon) HasClockSkew() bool`

HasClockSkew returns a boolean if a field has been set.

### GetComponentStack

`func (o *WebsiteMonitoringBeacon) GetComponentStack() string`

GetComponentStack returns the ComponentStack field if non-nil, zero value otherwise.

### GetComponentStackOk

`func (o *WebsiteMonitoringBeacon) GetComponentStackOk() (*string, bool)`

GetComponentStackOk returns a tuple with the ComponentStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentStack

`func (o *WebsiteMonitoringBeacon) SetComponentStack(v string)`

SetComponentStack sets ComponentStack field to given value.

### HasComponentStack

`func (o *WebsiteMonitoringBeacon) HasComponentStack() bool`

HasComponentStack returns a boolean if a field has been set.

### GetConnectionType

`func (o *WebsiteMonitoringBeacon) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *WebsiteMonitoringBeacon) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *WebsiteMonitoringBeacon) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *WebsiteMonitoringBeacon) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetContinent

`func (o *WebsiteMonitoringBeacon) GetContinent() string`

GetContinent returns the Continent field if non-nil, zero value otherwise.

### GetContinentOk

`func (o *WebsiteMonitoringBeacon) GetContinentOk() (*string, bool)`

GetContinentOk returns a tuple with the Continent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinent

`func (o *WebsiteMonitoringBeacon) SetContinent(v string)`

SetContinent sets Continent field to given value.

### HasContinent

`func (o *WebsiteMonitoringBeacon) HasContinent() bool`

HasContinent returns a boolean if a field has been set.

### GetContinentCode

`func (o *WebsiteMonitoringBeacon) GetContinentCode() string`

GetContinentCode returns the ContinentCode field if non-nil, zero value otherwise.

### GetContinentCodeOk

`func (o *WebsiteMonitoringBeacon) GetContinentCodeOk() (*string, bool)`

GetContinentCodeOk returns a tuple with the ContinentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinentCode

`func (o *WebsiteMonitoringBeacon) SetContinentCode(v string)`

SetContinentCode sets ContinentCode field to given value.

### HasContinentCode

`func (o *WebsiteMonitoringBeacon) HasContinentCode() bool`

HasContinentCode returns a boolean if a field has been set.

### GetCountry

`func (o *WebsiteMonitoringBeacon) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *WebsiteMonitoringBeacon) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *WebsiteMonitoringBeacon) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *WebsiteMonitoringBeacon) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *WebsiteMonitoringBeacon) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *WebsiteMonitoringBeacon) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *WebsiteMonitoringBeacon) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *WebsiteMonitoringBeacon) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCspBlockedUri

`func (o *WebsiteMonitoringBeacon) GetCspBlockedUri() string`

GetCspBlockedUri returns the CspBlockedUri field if non-nil, zero value otherwise.

### GetCspBlockedUriOk

`func (o *WebsiteMonitoringBeacon) GetCspBlockedUriOk() (*string, bool)`

GetCspBlockedUriOk returns a tuple with the CspBlockedUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspBlockedUri

`func (o *WebsiteMonitoringBeacon) SetCspBlockedUri(v string)`

SetCspBlockedUri sets CspBlockedUri field to given value.

### HasCspBlockedUri

`func (o *WebsiteMonitoringBeacon) HasCspBlockedUri() bool`

HasCspBlockedUri returns a boolean if a field has been set.

### GetCspColumnNumber

`func (o *WebsiteMonitoringBeacon) GetCspColumnNumber() int64`

GetCspColumnNumber returns the CspColumnNumber field if non-nil, zero value otherwise.

### GetCspColumnNumberOk

`func (o *WebsiteMonitoringBeacon) GetCspColumnNumberOk() (*int64, bool)`

GetCspColumnNumberOk returns a tuple with the CspColumnNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspColumnNumber

`func (o *WebsiteMonitoringBeacon) SetCspColumnNumber(v int64)`

SetCspColumnNumber sets CspColumnNumber field to given value.

### HasCspColumnNumber

`func (o *WebsiteMonitoringBeacon) HasCspColumnNumber() bool`

HasCspColumnNumber returns a boolean if a field has been set.

### GetCspDisposition

`func (o *WebsiteMonitoringBeacon) GetCspDisposition() string`

GetCspDisposition returns the CspDisposition field if non-nil, zero value otherwise.

### GetCspDispositionOk

`func (o *WebsiteMonitoringBeacon) GetCspDispositionOk() (*string, bool)`

GetCspDispositionOk returns a tuple with the CspDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspDisposition

`func (o *WebsiteMonitoringBeacon) SetCspDisposition(v string)`

SetCspDisposition sets CspDisposition field to given value.

### HasCspDisposition

`func (o *WebsiteMonitoringBeacon) HasCspDisposition() bool`

HasCspDisposition returns a boolean if a field has been set.

### GetCspEffectiveDirective

`func (o *WebsiteMonitoringBeacon) GetCspEffectiveDirective() string`

GetCspEffectiveDirective returns the CspEffectiveDirective field if non-nil, zero value otherwise.

### GetCspEffectiveDirectiveOk

`func (o *WebsiteMonitoringBeacon) GetCspEffectiveDirectiveOk() (*string, bool)`

GetCspEffectiveDirectiveOk returns a tuple with the CspEffectiveDirective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspEffectiveDirective

`func (o *WebsiteMonitoringBeacon) SetCspEffectiveDirective(v string)`

SetCspEffectiveDirective sets CspEffectiveDirective field to given value.

### HasCspEffectiveDirective

`func (o *WebsiteMonitoringBeacon) HasCspEffectiveDirective() bool`

HasCspEffectiveDirective returns a boolean if a field has been set.

### GetCspLineNumber

`func (o *WebsiteMonitoringBeacon) GetCspLineNumber() int64`

GetCspLineNumber returns the CspLineNumber field if non-nil, zero value otherwise.

### GetCspLineNumberOk

`func (o *WebsiteMonitoringBeacon) GetCspLineNumberOk() (*int64, bool)`

GetCspLineNumberOk returns a tuple with the CspLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspLineNumber

`func (o *WebsiteMonitoringBeacon) SetCspLineNumber(v int64)`

SetCspLineNumber sets CspLineNumber field to given value.

### HasCspLineNumber

`func (o *WebsiteMonitoringBeacon) HasCspLineNumber() bool`

HasCspLineNumber returns a boolean if a field has been set.

### GetCspOriginalPolicy

`func (o *WebsiteMonitoringBeacon) GetCspOriginalPolicy() string`

GetCspOriginalPolicy returns the CspOriginalPolicy field if non-nil, zero value otherwise.

### GetCspOriginalPolicyOk

`func (o *WebsiteMonitoringBeacon) GetCspOriginalPolicyOk() (*string, bool)`

GetCspOriginalPolicyOk returns a tuple with the CspOriginalPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspOriginalPolicy

`func (o *WebsiteMonitoringBeacon) SetCspOriginalPolicy(v string)`

SetCspOriginalPolicy sets CspOriginalPolicy field to given value.

### HasCspOriginalPolicy

`func (o *WebsiteMonitoringBeacon) HasCspOriginalPolicy() bool`

HasCspOriginalPolicy returns a boolean if a field has been set.

### GetCspSample

`func (o *WebsiteMonitoringBeacon) GetCspSample() string`

GetCspSample returns the CspSample field if non-nil, zero value otherwise.

### GetCspSampleOk

`func (o *WebsiteMonitoringBeacon) GetCspSampleOk() (*string, bool)`

GetCspSampleOk returns a tuple with the CspSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspSample

`func (o *WebsiteMonitoringBeacon) SetCspSample(v string)`

SetCspSample sets CspSample field to given value.

### HasCspSample

`func (o *WebsiteMonitoringBeacon) HasCspSample() bool`

HasCspSample returns a boolean if a field has been set.

### GetCspSourceFile

`func (o *WebsiteMonitoringBeacon) GetCspSourceFile() string`

GetCspSourceFile returns the CspSourceFile field if non-nil, zero value otherwise.

### GetCspSourceFileOk

`func (o *WebsiteMonitoringBeacon) GetCspSourceFileOk() (*string, bool)`

GetCspSourceFileOk returns a tuple with the CspSourceFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspSourceFile

`func (o *WebsiteMonitoringBeacon) SetCspSourceFile(v string)`

SetCspSourceFile sets CspSourceFile field to given value.

### HasCspSourceFile

`func (o *WebsiteMonitoringBeacon) HasCspSourceFile() bool`

HasCspSourceFile returns a boolean if a field has been set.

### GetCumulativeLayoutShift

`func (o *WebsiteMonitoringBeacon) GetCumulativeLayoutShift() float64`

GetCumulativeLayoutShift returns the CumulativeLayoutShift field if non-nil, zero value otherwise.

### GetCumulativeLayoutShiftOk

`func (o *WebsiteMonitoringBeacon) GetCumulativeLayoutShiftOk() (*float64, bool)`

GetCumulativeLayoutShiftOk returns a tuple with the CumulativeLayoutShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeLayoutShift

`func (o *WebsiteMonitoringBeacon) SetCumulativeLayoutShift(v float64)`

SetCumulativeLayoutShift sets CumulativeLayoutShift field to given value.

### HasCumulativeLayoutShift

`func (o *WebsiteMonitoringBeacon) HasCumulativeLayoutShift() bool`

HasCumulativeLayoutShift returns a boolean if a field has been set.

### GetCustomEventName

`func (o *WebsiteMonitoringBeacon) GetCustomEventName() string`

GetCustomEventName returns the CustomEventName field if non-nil, zero value otherwise.

### GetCustomEventNameOk

`func (o *WebsiteMonitoringBeacon) GetCustomEventNameOk() (*string, bool)`

GetCustomEventNameOk returns a tuple with the CustomEventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEventName

`func (o *WebsiteMonitoringBeacon) SetCustomEventName(v string)`

SetCustomEventName sets CustomEventName field to given value.

### HasCustomEventName

`func (o *WebsiteMonitoringBeacon) HasCustomEventName() bool`

HasCustomEventName returns a boolean if a field has been set.

### GetDecodedBodySize

`func (o *WebsiteMonitoringBeacon) GetDecodedBodySize() int64`

GetDecodedBodySize returns the DecodedBodySize field if non-nil, zero value otherwise.

### GetDecodedBodySizeOk

`func (o *WebsiteMonitoringBeacon) GetDecodedBodySizeOk() (*int64, bool)`

GetDecodedBodySizeOk returns a tuple with the DecodedBodySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecodedBodySize

`func (o *WebsiteMonitoringBeacon) SetDecodedBodySize(v int64)`

SetDecodedBodySize sets DecodedBodySize field to given value.

### HasDecodedBodySize

`func (o *WebsiteMonitoringBeacon) HasDecodedBodySize() bool`

HasDecodedBodySize returns a boolean if a field has been set.

### GetDeprecations

`func (o *WebsiteMonitoringBeacon) GetDeprecations() []string`

GetDeprecations returns the Deprecations field if non-nil, zero value otherwise.

### GetDeprecationsOk

`func (o *WebsiteMonitoringBeacon) GetDeprecationsOk() (*[]string, bool)`

GetDeprecationsOk returns a tuple with the Deprecations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecations

`func (o *WebsiteMonitoringBeacon) SetDeprecations(v []string)`

SetDeprecations sets Deprecations field to given value.

### HasDeprecations

`func (o *WebsiteMonitoringBeacon) HasDeprecations() bool`

HasDeprecations returns a boolean if a field has been set.

### GetDeviceType

`func (o *WebsiteMonitoringBeacon) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *WebsiteMonitoringBeacon) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *WebsiteMonitoringBeacon) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *WebsiteMonitoringBeacon) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDnsTime

`func (o *WebsiteMonitoringBeacon) GetDnsTime() int64`

GetDnsTime returns the DnsTime field if non-nil, zero value otherwise.

### GetDnsTimeOk

`func (o *WebsiteMonitoringBeacon) GetDnsTimeOk() (*int64, bool)`

GetDnsTimeOk returns a tuple with the DnsTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTime

`func (o *WebsiteMonitoringBeacon) SetDnsTime(v int64)`

SetDnsTime sets DnsTime field to given value.

### HasDnsTime

`func (o *WebsiteMonitoringBeacon) HasDnsTime() bool`

HasDnsTime returns a boolean if a field has been set.

### GetDomTime

`func (o *WebsiteMonitoringBeacon) GetDomTime() int64`

GetDomTime returns the DomTime field if non-nil, zero value otherwise.

### GetDomTimeOk

`func (o *WebsiteMonitoringBeacon) GetDomTimeOk() (*int64, bool)`

GetDomTimeOk returns a tuple with the DomTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomTime

`func (o *WebsiteMonitoringBeacon) SetDomTime(v int64)`

SetDomTime sets DomTime field to given value.

### HasDomTime

`func (o *WebsiteMonitoringBeacon) HasDomTime() bool`

HasDomTime returns a boolean if a field has been set.

### GetDuration

`func (o *WebsiteMonitoringBeacon) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WebsiteMonitoringBeacon) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WebsiteMonitoringBeacon) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *WebsiteMonitoringBeacon) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEncodedBodySize

`func (o *WebsiteMonitoringBeacon) GetEncodedBodySize() int64`

GetEncodedBodySize returns the EncodedBodySize field if non-nil, zero value otherwise.

### GetEncodedBodySizeOk

`func (o *WebsiteMonitoringBeacon) GetEncodedBodySizeOk() (*int64, bool)`

GetEncodedBodySizeOk returns a tuple with the EncodedBodySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedBodySize

`func (o *WebsiteMonitoringBeacon) SetEncodedBodySize(v int64)`

SetEncodedBodySize sets EncodedBodySize field to given value.

### HasEncodedBodySize

`func (o *WebsiteMonitoringBeacon) HasEncodedBodySize() bool`

HasEncodedBodySize returns a boolean if a field has been set.

### GetErrorCount

`func (o *WebsiteMonitoringBeacon) GetErrorCount() int64`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *WebsiteMonitoringBeacon) GetErrorCountOk() (*int64, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *WebsiteMonitoringBeacon) SetErrorCount(v int64)`

SetErrorCount sets ErrorCount field to given value.

### HasErrorCount

`func (o *WebsiteMonitoringBeacon) HasErrorCount() bool`

HasErrorCount returns a boolean if a field has been set.

### GetErrorId

`func (o *WebsiteMonitoringBeacon) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *WebsiteMonitoringBeacon) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *WebsiteMonitoringBeacon) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *WebsiteMonitoringBeacon) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.

### GetErrorMessage

`func (o *WebsiteMonitoringBeacon) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *WebsiteMonitoringBeacon) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *WebsiteMonitoringBeacon) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *WebsiteMonitoringBeacon) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetErrorType

`func (o *WebsiteMonitoringBeacon) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *WebsiteMonitoringBeacon) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *WebsiteMonitoringBeacon) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *WebsiteMonitoringBeacon) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetFirstContentfulPaintTime

`func (o *WebsiteMonitoringBeacon) GetFirstContentfulPaintTime() int64`

GetFirstContentfulPaintTime returns the FirstContentfulPaintTime field if non-nil, zero value otherwise.

### GetFirstContentfulPaintTimeOk

`func (o *WebsiteMonitoringBeacon) GetFirstContentfulPaintTimeOk() (*int64, bool)`

GetFirstContentfulPaintTimeOk returns a tuple with the FirstContentfulPaintTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstContentfulPaintTime

`func (o *WebsiteMonitoringBeacon) SetFirstContentfulPaintTime(v int64)`

SetFirstContentfulPaintTime sets FirstContentfulPaintTime field to given value.

### HasFirstContentfulPaintTime

`func (o *WebsiteMonitoringBeacon) HasFirstContentfulPaintTime() bool`

HasFirstContentfulPaintTime returns a boolean if a field has been set.

### GetFirstInputDelayTime

`func (o *WebsiteMonitoringBeacon) GetFirstInputDelayTime() int64`

GetFirstInputDelayTime returns the FirstInputDelayTime field if non-nil, zero value otherwise.

### GetFirstInputDelayTimeOk

`func (o *WebsiteMonitoringBeacon) GetFirstInputDelayTimeOk() (*int64, bool)`

GetFirstInputDelayTimeOk returns a tuple with the FirstInputDelayTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstInputDelayTime

`func (o *WebsiteMonitoringBeacon) SetFirstInputDelayTime(v int64)`

SetFirstInputDelayTime sets FirstInputDelayTime field to given value.

### HasFirstInputDelayTime

`func (o *WebsiteMonitoringBeacon) HasFirstInputDelayTime() bool`

HasFirstInputDelayTime returns a boolean if a field has been set.

### GetFirstPaintTime

`func (o *WebsiteMonitoringBeacon) GetFirstPaintTime() int64`

GetFirstPaintTime returns the FirstPaintTime field if non-nil, zero value otherwise.

### GetFirstPaintTimeOk

`func (o *WebsiteMonitoringBeacon) GetFirstPaintTimeOk() (*int64, bool)`

GetFirstPaintTimeOk returns a tuple with the FirstPaintTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPaintTime

`func (o *WebsiteMonitoringBeacon) SetFirstPaintTime(v int64)`

SetFirstPaintTime sets FirstPaintTime field to given value.

### HasFirstPaintTime

`func (o *WebsiteMonitoringBeacon) HasFirstPaintTime() bool`

HasFirstPaintTime returns a boolean if a field has been set.

### GetFrontendTime

`func (o *WebsiteMonitoringBeacon) GetFrontendTime() int64`

GetFrontendTime returns the FrontendTime field if non-nil, zero value otherwise.

### GetFrontendTimeOk

`func (o *WebsiteMonitoringBeacon) GetFrontendTimeOk() (*int64, bool)`

GetFrontendTimeOk returns a tuple with the FrontendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontendTime

`func (o *WebsiteMonitoringBeacon) SetFrontendTime(v int64)`

SetFrontendTime sets FrontendTime field to given value.

### HasFrontendTime

`func (o *WebsiteMonitoringBeacon) HasFrontendTime() bool`

HasFrontendTime returns a boolean if a field has been set.

### GetGraphqlOperationName

`func (o *WebsiteMonitoringBeacon) GetGraphqlOperationName() string`

GetGraphqlOperationName returns the GraphqlOperationName field if non-nil, zero value otherwise.

### GetGraphqlOperationNameOk

`func (o *WebsiteMonitoringBeacon) GetGraphqlOperationNameOk() (*string, bool)`

GetGraphqlOperationNameOk returns a tuple with the GraphqlOperationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphqlOperationName

`func (o *WebsiteMonitoringBeacon) SetGraphqlOperationName(v string)`

SetGraphqlOperationName sets GraphqlOperationName field to given value.

### HasGraphqlOperationName

`func (o *WebsiteMonitoringBeacon) HasGraphqlOperationName() bool`

HasGraphqlOperationName returns a boolean if a field has been set.

### GetGraphqlOperationType

`func (o *WebsiteMonitoringBeacon) GetGraphqlOperationType() string`

GetGraphqlOperationType returns the GraphqlOperationType field if non-nil, zero value otherwise.

### GetGraphqlOperationTypeOk

`func (o *WebsiteMonitoringBeacon) GetGraphqlOperationTypeOk() (*string, bool)`

GetGraphqlOperationTypeOk returns a tuple with the GraphqlOperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphqlOperationType

`func (o *WebsiteMonitoringBeacon) SetGraphqlOperationType(v string)`

SetGraphqlOperationType sets GraphqlOperationType field to given value.

### HasGraphqlOperationType

`func (o *WebsiteMonitoringBeacon) HasGraphqlOperationType() bool`

HasGraphqlOperationType returns a boolean if a field has been set.

### GetHttpCallAsynchronous

`func (o *WebsiteMonitoringBeacon) GetHttpCallAsynchronous() bool`

GetHttpCallAsynchronous returns the HttpCallAsynchronous field if non-nil, zero value otherwise.

### GetHttpCallAsynchronousOk

`func (o *WebsiteMonitoringBeacon) GetHttpCallAsynchronousOk() (*bool, bool)`

GetHttpCallAsynchronousOk returns a tuple with the HttpCallAsynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCallAsynchronous

`func (o *WebsiteMonitoringBeacon) SetHttpCallAsynchronous(v bool)`

SetHttpCallAsynchronous sets HttpCallAsynchronous field to given value.

### HasHttpCallAsynchronous

`func (o *WebsiteMonitoringBeacon) HasHttpCallAsynchronous() bool`

HasHttpCallAsynchronous returns a boolean if a field has been set.

### GetHttpCallCorrelationAttempted

`func (o *WebsiteMonitoringBeacon) GetHttpCallCorrelationAttempted() bool`

GetHttpCallCorrelationAttempted returns the HttpCallCorrelationAttempted field if non-nil, zero value otherwise.

### GetHttpCallCorrelationAttemptedOk

`func (o *WebsiteMonitoringBeacon) GetHttpCallCorrelationAttemptedOk() (*bool, bool)`

GetHttpCallCorrelationAttemptedOk returns a tuple with the HttpCallCorrelationAttempted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCallCorrelationAttempted

`func (o *WebsiteMonitoringBeacon) SetHttpCallCorrelationAttempted(v bool)`

SetHttpCallCorrelationAttempted sets HttpCallCorrelationAttempted field to given value.

### HasHttpCallCorrelationAttempted

`func (o *WebsiteMonitoringBeacon) HasHttpCallCorrelationAttempted() bool`

HasHttpCallCorrelationAttempted returns a boolean if a field has been set.

### GetHttpCallHeaders

`func (o *WebsiteMonitoringBeacon) GetHttpCallHeaders() map[string]string`

GetHttpCallHeaders returns the HttpCallHeaders field if non-nil, zero value otherwise.

### GetHttpCallHeadersOk

`func (o *WebsiteMonitoringBeacon) GetHttpCallHeadersOk() (*map[string]string, bool)`

GetHttpCallHeadersOk returns a tuple with the HttpCallHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCallHeaders

`func (o *WebsiteMonitoringBeacon) SetHttpCallHeaders(v map[string]string)`

SetHttpCallHeaders sets HttpCallHeaders field to given value.

### HasHttpCallHeaders

`func (o *WebsiteMonitoringBeacon) HasHttpCallHeaders() bool`

HasHttpCallHeaders returns a boolean if a field has been set.

### GetHttpCallMethod

`func (o *WebsiteMonitoringBeacon) GetHttpCallMethod() string`

GetHttpCallMethod returns the HttpCallMethod field if non-nil, zero value otherwise.

### GetHttpCallMethodOk

`func (o *WebsiteMonitoringBeacon) GetHttpCallMethodOk() (*string, bool)`

GetHttpCallMethodOk returns a tuple with the HttpCallMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCallMethod

`func (o *WebsiteMonitoringBeacon) SetHttpCallMethod(v string)`

SetHttpCallMethod sets HttpCallMethod field to given value.

### HasHttpCallMethod

`func (o *WebsiteMonitoringBeacon) HasHttpCallMethod() bool`

HasHttpCallMethod returns a boolean if a field has been set.

### GetHttpCallOrigin

`func (o *WebsiteMonitoringBeacon) GetHttpCallOrigin() string`

GetHttpCallOrigin returns the HttpCallOrigin field if non-nil, zero value otherwise.

### GetHttpCallOriginOk

`func (o *WebsiteMonitoringBeacon) GetHttpCallOriginOk() (*string, bool)`

GetHttpCallOriginOk returns a tuple with the HttpCallOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCallOrigin

`func (o *WebsiteMonitoringBeacon) SetHttpCallOrigin(v string)`

SetHttpCallOrigin sets HttpCallOrigin field to given value.

### HasHttpCallOrigin

`func (o *WebsiteMonitoringBeacon) HasHttpCallOrigin() bool`

HasHttpCallOrigin returns a boolean if a field has been set.

### GetHttpCallPath

`func (o *WebsiteMonitoringBeacon) GetHttpCallPath() string`

GetHttpCallPath returns the HttpCallPath field if non-nil, zero value otherwise.

### GetHttpCallPathOk

`func (o *WebsiteMonitoringBeacon) GetHttpCallPathOk() (*string, bool)`

GetHttpCallPathOk returns a tuple with the HttpCallPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCallPath

`func (o *WebsiteMonitoringBeacon) SetHttpCallPath(v string)`

SetHttpCallPath sets HttpCallPath field to given value.

### HasHttpCallPath

`func (o *WebsiteMonitoringBeacon) HasHttpCallPath() bool`

HasHttpCallPath returns a boolean if a field has been set.

### GetHttpCallStatus

`func (o *WebsiteMonitoringBeacon) GetHttpCallStatus() int32`

GetHttpCallStatus returns the HttpCallStatus field if non-nil, zero value otherwise.

### GetHttpCallStatusOk

`func (o *WebsiteMonitoringBeacon) GetHttpCallStatusOk() (*int32, bool)`

GetHttpCallStatusOk returns a tuple with the HttpCallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCallStatus

`func (o *WebsiteMonitoringBeacon) SetHttpCallStatus(v int32)`

SetHttpCallStatus sets HttpCallStatus field to given value.

### HasHttpCallStatus

`func (o *WebsiteMonitoringBeacon) HasHttpCallStatus() bool`

HasHttpCallStatus returns a boolean if a field has been set.

### GetHttpCallUrl

`func (o *WebsiteMonitoringBeacon) GetHttpCallUrl() string`

GetHttpCallUrl returns the HttpCallUrl field if non-nil, zero value otherwise.

### GetHttpCallUrlOk

`func (o *WebsiteMonitoringBeacon) GetHttpCallUrlOk() (*string, bool)`

GetHttpCallUrlOk returns a tuple with the HttpCallUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCallUrl

`func (o *WebsiteMonitoringBeacon) SetHttpCallUrl(v string)`

SetHttpCallUrl sets HttpCallUrl field to given value.

### HasHttpCallUrl

`func (o *WebsiteMonitoringBeacon) HasHttpCallUrl() bool`

HasHttpCallUrl returns a boolean if a field has been set.

### GetInitiator

`func (o *WebsiteMonitoringBeacon) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *WebsiteMonitoringBeacon) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *WebsiteMonitoringBeacon) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *WebsiteMonitoringBeacon) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetLargestContentfulPaintTime

`func (o *WebsiteMonitoringBeacon) GetLargestContentfulPaintTime() int64`

GetLargestContentfulPaintTime returns the LargestContentfulPaintTime field if non-nil, zero value otherwise.

### GetLargestContentfulPaintTimeOk

`func (o *WebsiteMonitoringBeacon) GetLargestContentfulPaintTimeOk() (*int64, bool)`

GetLargestContentfulPaintTimeOk returns a tuple with the LargestContentfulPaintTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargestContentfulPaintTime

`func (o *WebsiteMonitoringBeacon) SetLargestContentfulPaintTime(v int64)`

SetLargestContentfulPaintTime sets LargestContentfulPaintTime field to given value.

### HasLargestContentfulPaintTime

`func (o *WebsiteMonitoringBeacon) HasLargestContentfulPaintTime() bool`

HasLargestContentfulPaintTime returns a boolean if a field has been set.

### GetLatitude

`func (o *WebsiteMonitoringBeacon) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *WebsiteMonitoringBeacon) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *WebsiteMonitoringBeacon) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *WebsiteMonitoringBeacon) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLocationOrigin

`func (o *WebsiteMonitoringBeacon) GetLocationOrigin() string`

GetLocationOrigin returns the LocationOrigin field if non-nil, zero value otherwise.

### GetLocationOriginOk

`func (o *WebsiteMonitoringBeacon) GetLocationOriginOk() (*string, bool)`

GetLocationOriginOk returns a tuple with the LocationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationOrigin

`func (o *WebsiteMonitoringBeacon) SetLocationOrigin(v string)`

SetLocationOrigin sets LocationOrigin field to given value.


### GetLocationPath

`func (o *WebsiteMonitoringBeacon) GetLocationPath() string`

GetLocationPath returns the LocationPath field if non-nil, zero value otherwise.

### GetLocationPathOk

`func (o *WebsiteMonitoringBeacon) GetLocationPathOk() (*string, bool)`

GetLocationPathOk returns a tuple with the LocationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationPath

`func (o *WebsiteMonitoringBeacon) SetLocationPath(v string)`

SetLocationPath sets LocationPath field to given value.

### HasLocationPath

`func (o *WebsiteMonitoringBeacon) HasLocationPath() bool`

HasLocationPath returns a boolean if a field has been set.

### GetLocationUrl

`func (o *WebsiteMonitoringBeacon) GetLocationUrl() string`

GetLocationUrl returns the LocationUrl field if non-nil, zero value otherwise.

### GetLocationUrlOk

`func (o *WebsiteMonitoringBeacon) GetLocationUrlOk() (*string, bool)`

GetLocationUrlOk returns a tuple with the LocationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationUrl

`func (o *WebsiteMonitoringBeacon) SetLocationUrl(v string)`

SetLocationUrl sets LocationUrl field to given value.


### GetLongitude

`func (o *WebsiteMonitoringBeacon) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *WebsiteMonitoringBeacon) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *WebsiteMonitoringBeacon) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *WebsiteMonitoringBeacon) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetMeta

`func (o *WebsiteMonitoringBeacon) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WebsiteMonitoringBeacon) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WebsiteMonitoringBeacon) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *WebsiteMonitoringBeacon) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetOnLoadTime

`func (o *WebsiteMonitoringBeacon) GetOnLoadTime() int64`

GetOnLoadTime returns the OnLoadTime field if non-nil, zero value otherwise.

### GetOnLoadTimeOk

`func (o *WebsiteMonitoringBeacon) GetOnLoadTimeOk() (*int64, bool)`

GetOnLoadTimeOk returns a tuple with the OnLoadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnLoadTime

`func (o *WebsiteMonitoringBeacon) SetOnLoadTime(v int64)`

SetOnLoadTime sets OnLoadTime field to given value.

### HasOnLoadTime

`func (o *WebsiteMonitoringBeacon) HasOnLoadTime() bool`

HasOnLoadTime returns a boolean if a field has been set.

### GetOsName

`func (o *WebsiteMonitoringBeacon) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *WebsiteMonitoringBeacon) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *WebsiteMonitoringBeacon) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *WebsiteMonitoringBeacon) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetOsVersion

`func (o *WebsiteMonitoringBeacon) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *WebsiteMonitoringBeacon) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *WebsiteMonitoringBeacon) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *WebsiteMonitoringBeacon) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetPage

`func (o *WebsiteMonitoringBeacon) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *WebsiteMonitoringBeacon) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *WebsiteMonitoringBeacon) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *WebsiteMonitoringBeacon) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageLoadId

`func (o *WebsiteMonitoringBeacon) GetPageLoadId() string`

GetPageLoadId returns the PageLoadId field if non-nil, zero value otherwise.

### GetPageLoadIdOk

`func (o *WebsiteMonitoringBeacon) GetPageLoadIdOk() (*string, bool)`

GetPageLoadIdOk returns a tuple with the PageLoadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageLoadId

`func (o *WebsiteMonitoringBeacon) SetPageLoadId(v string)`

SetPageLoadId sets PageLoadId field to given value.


### GetParsedStackTrace

`func (o *WebsiteMonitoringBeacon) GetParsedStackTrace() []StackTraceLine`

GetParsedStackTrace returns the ParsedStackTrace field if non-nil, zero value otherwise.

### GetParsedStackTraceOk

`func (o *WebsiteMonitoringBeacon) GetParsedStackTraceOk() (*[]StackTraceLine, bool)`

GetParsedStackTraceOk returns a tuple with the ParsedStackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsedStackTrace

`func (o *WebsiteMonitoringBeacon) SetParsedStackTrace(v []StackTraceLine)`

SetParsedStackTrace sets ParsedStackTrace field to given value.

### HasParsedStackTrace

`func (o *WebsiteMonitoringBeacon) HasParsedStackTrace() bool`

HasParsedStackTrace returns a boolean if a field has been set.

### GetPhase

`func (o *WebsiteMonitoringBeacon) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *WebsiteMonitoringBeacon) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *WebsiteMonitoringBeacon) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *WebsiteMonitoringBeacon) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetProcessingTime

`func (o *WebsiteMonitoringBeacon) GetProcessingTime() int64`

GetProcessingTime returns the ProcessingTime field if non-nil, zero value otherwise.

### GetProcessingTimeOk

`func (o *WebsiteMonitoringBeacon) GetProcessingTimeOk() (*int64, bool)`

GetProcessingTimeOk returns a tuple with the ProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTime

`func (o *WebsiteMonitoringBeacon) SetProcessingTime(v int64)`

SetProcessingTime sets ProcessingTime field to given value.

### HasProcessingTime

`func (o *WebsiteMonitoringBeacon) HasProcessingTime() bool`

HasProcessingTime returns a boolean if a field has been set.

### GetRedirectTime

`func (o *WebsiteMonitoringBeacon) GetRedirectTime() int64`

GetRedirectTime returns the RedirectTime field if non-nil, zero value otherwise.

### GetRedirectTimeOk

`func (o *WebsiteMonitoringBeacon) GetRedirectTimeOk() (*int64, bool)`

GetRedirectTimeOk returns a tuple with the RedirectTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTime

`func (o *WebsiteMonitoringBeacon) SetRedirectTime(v int64)`

SetRedirectTime sets RedirectTime field to given value.

### HasRedirectTime

`func (o *WebsiteMonitoringBeacon) HasRedirectTime() bool`

HasRedirectTime returns a boolean if a field has been set.

### GetRequestTime

`func (o *WebsiteMonitoringBeacon) GetRequestTime() int64`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *WebsiteMonitoringBeacon) GetRequestTimeOk() (*int64, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *WebsiteMonitoringBeacon) SetRequestTime(v int64)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *WebsiteMonitoringBeacon) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.

### GetResourceType

`func (o *WebsiteMonitoringBeacon) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *WebsiteMonitoringBeacon) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *WebsiteMonitoringBeacon) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *WebsiteMonitoringBeacon) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResponseTime

`func (o *WebsiteMonitoringBeacon) GetResponseTime() int64`

GetResponseTime returns the ResponseTime field if non-nil, zero value otherwise.

### GetResponseTimeOk

`func (o *WebsiteMonitoringBeacon) GetResponseTimeOk() (*int64, bool)`

GetResponseTimeOk returns a tuple with the ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTime

`func (o *WebsiteMonitoringBeacon) SetResponseTime(v int64)`

SetResponseTime sets ResponseTime field to given value.

### HasResponseTime

`func (o *WebsiteMonitoringBeacon) HasResponseTime() bool`

HasResponseTime returns a boolean if a field has been set.

### GetSessionId

`func (o *WebsiteMonitoringBeacon) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *WebsiteMonitoringBeacon) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *WebsiteMonitoringBeacon) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *WebsiteMonitoringBeacon) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSnippetVersion

`func (o *WebsiteMonitoringBeacon) GetSnippetVersion() string`

GetSnippetVersion returns the SnippetVersion field if non-nil, zero value otherwise.

### GetSnippetVersionOk

`func (o *WebsiteMonitoringBeacon) GetSnippetVersionOk() (*string, bool)`

GetSnippetVersionOk returns a tuple with the SnippetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippetVersion

`func (o *WebsiteMonitoringBeacon) SetSnippetVersion(v string)`

SetSnippetVersion sets SnippetVersion field to given value.

### HasSnippetVersion

`func (o *WebsiteMonitoringBeacon) HasSnippetVersion() bool`

HasSnippetVersion returns a boolean if a field has been set.

### GetSslTime

`func (o *WebsiteMonitoringBeacon) GetSslTime() int64`

GetSslTime returns the SslTime field if non-nil, zero value otherwise.

### GetSslTimeOk

`func (o *WebsiteMonitoringBeacon) GetSslTimeOk() (*int64, bool)`

GetSslTimeOk returns a tuple with the SslTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslTime

`func (o *WebsiteMonitoringBeacon) SetSslTime(v int64)`

SetSslTime sets SslTime field to given value.

### HasSslTime

`func (o *WebsiteMonitoringBeacon) HasSslTime() bool`

HasSslTime returns a boolean if a field has been set.

### GetStackTrace

`func (o *WebsiteMonitoringBeacon) GetStackTrace() string`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *WebsiteMonitoringBeacon) GetStackTraceOk() (*string, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *WebsiteMonitoringBeacon) SetStackTrace(v string)`

SetStackTrace sets StackTrace field to given value.

### HasStackTrace

`func (o *WebsiteMonitoringBeacon) HasStackTrace() bool`

HasStackTrace returns a boolean if a field has been set.

### GetStackTraceParsingStatus

`func (o *WebsiteMonitoringBeacon) GetStackTraceParsingStatus() int32`

GetStackTraceParsingStatus returns the StackTraceParsingStatus field if non-nil, zero value otherwise.

### GetStackTraceParsingStatusOk

`func (o *WebsiteMonitoringBeacon) GetStackTraceParsingStatusOk() (*int32, bool)`

GetStackTraceParsingStatusOk returns a tuple with the StackTraceParsingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTraceParsingStatus

`func (o *WebsiteMonitoringBeacon) SetStackTraceParsingStatus(v int32)`

SetStackTraceParsingStatus sets StackTraceParsingStatus field to given value.

### HasStackTraceParsingStatus

`func (o *WebsiteMonitoringBeacon) HasStackTraceParsingStatus() bool`

HasStackTraceParsingStatus returns a boolean if a field has been set.

### GetStackTraceReadability

`func (o *WebsiteMonitoringBeacon) GetStackTraceReadability() int32`

GetStackTraceReadability returns the StackTraceReadability field if non-nil, zero value otherwise.

### GetStackTraceReadabilityOk

`func (o *WebsiteMonitoringBeacon) GetStackTraceReadabilityOk() (*int32, bool)`

GetStackTraceReadabilityOk returns a tuple with the StackTraceReadability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTraceReadability

`func (o *WebsiteMonitoringBeacon) SetStackTraceReadability(v int32)`

SetStackTraceReadability sets StackTraceReadability field to given value.

### HasStackTraceReadability

`func (o *WebsiteMonitoringBeacon) HasStackTraceReadability() bool`

HasStackTraceReadability returns a boolean if a field has been set.

### GetSubdivision

`func (o *WebsiteMonitoringBeacon) GetSubdivision() string`

GetSubdivision returns the Subdivision field if non-nil, zero value otherwise.

### GetSubdivisionOk

`func (o *WebsiteMonitoringBeacon) GetSubdivisionOk() (*string, bool)`

GetSubdivisionOk returns a tuple with the Subdivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdivision

`func (o *WebsiteMonitoringBeacon) SetSubdivision(v string)`

SetSubdivision sets Subdivision field to given value.

### HasSubdivision

`func (o *WebsiteMonitoringBeacon) HasSubdivision() bool`

HasSubdivision returns a boolean if a field has been set.

### GetSubdivisionCode

`func (o *WebsiteMonitoringBeacon) GetSubdivisionCode() string`

GetSubdivisionCode returns the SubdivisionCode field if non-nil, zero value otherwise.

### GetSubdivisionCodeOk

`func (o *WebsiteMonitoringBeacon) GetSubdivisionCodeOk() (*string, bool)`

GetSubdivisionCodeOk returns a tuple with the SubdivisionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdivisionCode

`func (o *WebsiteMonitoringBeacon) SetSubdivisionCode(v string)`

SetSubdivisionCode sets SubdivisionCode field to given value.

### HasSubdivisionCode

`func (o *WebsiteMonitoringBeacon) HasSubdivisionCode() bool`

HasSubdivisionCode returns a boolean if a field has been set.

### GetTcpTime

`func (o *WebsiteMonitoringBeacon) GetTcpTime() int64`

GetTcpTime returns the TcpTime field if non-nil, zero value otherwise.

### GetTcpTimeOk

`func (o *WebsiteMonitoringBeacon) GetTcpTimeOk() (*int64, bool)`

GetTcpTimeOk returns a tuple with the TcpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpTime

`func (o *WebsiteMonitoringBeacon) SetTcpTime(v int64)`

SetTcpTime sets TcpTime field to given value.

### HasTcpTime

`func (o *WebsiteMonitoringBeacon) HasTcpTime() bool`

HasTcpTime returns a boolean if a field has been set.

### GetTimestamp

`func (o *WebsiteMonitoringBeacon) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebsiteMonitoringBeacon) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebsiteMonitoringBeacon) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WebsiteMonitoringBeacon) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTransferSize

`func (o *WebsiteMonitoringBeacon) GetTransferSize() int64`

GetTransferSize returns the TransferSize field if non-nil, zero value otherwise.

### GetTransferSizeOk

`func (o *WebsiteMonitoringBeacon) GetTransferSizeOk() (*int64, bool)`

GetTransferSizeOk returns a tuple with the TransferSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferSize

`func (o *WebsiteMonitoringBeacon) SetTransferSize(v int64)`

SetTransferSize sets TransferSize field to given value.

### HasTransferSize

`func (o *WebsiteMonitoringBeacon) HasTransferSize() bool`

HasTransferSize returns a boolean if a field has been set.

### GetType

`func (o *WebsiteMonitoringBeacon) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebsiteMonitoringBeacon) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebsiteMonitoringBeacon) SetType(v string)`

SetType sets Type field to given value.


### GetUnloadTime

`func (o *WebsiteMonitoringBeacon) GetUnloadTime() int64`

GetUnloadTime returns the UnloadTime field if non-nil, zero value otherwise.

### GetUnloadTimeOk

`func (o *WebsiteMonitoringBeacon) GetUnloadTimeOk() (*int64, bool)`

GetUnloadTimeOk returns a tuple with the UnloadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnloadTime

`func (o *WebsiteMonitoringBeacon) SetUnloadTime(v int64)`

SetUnloadTime sets UnloadTime field to given value.

### HasUnloadTime

`func (o *WebsiteMonitoringBeacon) HasUnloadTime() bool`

HasUnloadTime returns a boolean if a field has been set.

### GetUserEmail

`func (o *WebsiteMonitoringBeacon) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *WebsiteMonitoringBeacon) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *WebsiteMonitoringBeacon) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *WebsiteMonitoringBeacon) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUserId

`func (o *WebsiteMonitoringBeacon) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WebsiteMonitoringBeacon) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WebsiteMonitoringBeacon) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WebsiteMonitoringBeacon) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserIp

`func (o *WebsiteMonitoringBeacon) GetUserIp() string`

GetUserIp returns the UserIp field if non-nil, zero value otherwise.

### GetUserIpOk

`func (o *WebsiteMonitoringBeacon) GetUserIpOk() (*string, bool)`

GetUserIpOk returns a tuple with the UserIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIp

`func (o *WebsiteMonitoringBeacon) SetUserIp(v string)`

SetUserIp sets UserIp field to given value.

### HasUserIp

`func (o *WebsiteMonitoringBeacon) HasUserIp() bool`

HasUserIp returns a boolean if a field has been set.

### GetUserLanguages

`func (o *WebsiteMonitoringBeacon) GetUserLanguages() []string`

GetUserLanguages returns the UserLanguages field if non-nil, zero value otherwise.

### GetUserLanguagesOk

`func (o *WebsiteMonitoringBeacon) GetUserLanguagesOk() (*[]string, bool)`

GetUserLanguagesOk returns a tuple with the UserLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLanguages

`func (o *WebsiteMonitoringBeacon) SetUserLanguages(v []string)`

SetUserLanguages sets UserLanguages field to given value.

### HasUserLanguages

`func (o *WebsiteMonitoringBeacon) HasUserLanguages() bool`

HasUserLanguages returns a boolean if a field has been set.

### GetUserName

`func (o *WebsiteMonitoringBeacon) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *WebsiteMonitoringBeacon) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *WebsiteMonitoringBeacon) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *WebsiteMonitoringBeacon) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetWebsiteId

`func (o *WebsiteMonitoringBeacon) GetWebsiteId() string`

GetWebsiteId returns the WebsiteId field if non-nil, zero value otherwise.

### GetWebsiteIdOk

`func (o *WebsiteMonitoringBeacon) GetWebsiteIdOk() (*string, bool)`

GetWebsiteIdOk returns a tuple with the WebsiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteId

`func (o *WebsiteMonitoringBeacon) SetWebsiteId(v string)`

SetWebsiteId sets WebsiteId field to given value.


### GetWebsiteLabel

`func (o *WebsiteMonitoringBeacon) GetWebsiteLabel() string`

GetWebsiteLabel returns the WebsiteLabel field if non-nil, zero value otherwise.

### GetWebsiteLabelOk

`func (o *WebsiteMonitoringBeacon) GetWebsiteLabelOk() (*string, bool)`

GetWebsiteLabelOk returns a tuple with the WebsiteLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteLabel

`func (o *WebsiteMonitoringBeacon) SetWebsiteLabel(v string)`

SetWebsiteLabel sets WebsiteLabel field to given value.


### GetWindowHeight

`func (o *WebsiteMonitoringBeacon) GetWindowHeight() int32`

GetWindowHeight returns the WindowHeight field if non-nil, zero value otherwise.

### GetWindowHeightOk

`func (o *WebsiteMonitoringBeacon) GetWindowHeightOk() (*int32, bool)`

GetWindowHeightOk returns a tuple with the WindowHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowHeight

`func (o *WebsiteMonitoringBeacon) SetWindowHeight(v int32)`

SetWindowHeight sets WindowHeight field to given value.

### HasWindowHeight

`func (o *WebsiteMonitoringBeacon) HasWindowHeight() bool`

HasWindowHeight returns a boolean if a field has been set.

### GetWindowHidden

`func (o *WebsiteMonitoringBeacon) GetWindowHidden() bool`

GetWindowHidden returns the WindowHidden field if non-nil, zero value otherwise.

### GetWindowHiddenOk

`func (o *WebsiteMonitoringBeacon) GetWindowHiddenOk() (*bool, bool)`

GetWindowHiddenOk returns a tuple with the WindowHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowHidden

`func (o *WebsiteMonitoringBeacon) SetWindowHidden(v bool)`

SetWindowHidden sets WindowHidden field to given value.

### HasWindowHidden

`func (o *WebsiteMonitoringBeacon) HasWindowHidden() bool`

HasWindowHidden returns a boolean if a field has been set.

### GetWindowWidth

`func (o *WebsiteMonitoringBeacon) GetWindowWidth() int32`

GetWindowWidth returns the WindowWidth field if non-nil, zero value otherwise.

### GetWindowWidthOk

`func (o *WebsiteMonitoringBeacon) GetWindowWidthOk() (*int32, bool)`

GetWindowWidthOk returns a tuple with the WindowWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowWidth

`func (o *WebsiteMonitoringBeacon) SetWindowWidth(v int32)`

SetWindowWidth sets WindowWidth field to given value.

### HasWindowWidth

`func (o *WebsiteMonitoringBeacon) HasWindowWidth() bool`

HasWindowWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


