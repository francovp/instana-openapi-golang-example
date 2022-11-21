# GeoMappingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccuracyRadius** | Pointer to **int64** |  | [optional] 
**Cidr** | **string** |  | 
**City** | Pointer to **string** |  | [optional] 
**Continent** | Pointer to **string** |  | [optional] 
**ContinentCode** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**Latitude** | Pointer to **float64** |  | [optional] 
**LeastSpecificSubdivision** | Pointer to [**GeoSubdivision**](GeoSubdivision.md) |  | [optional] 
**Longitude** | Pointer to **float64** |  | [optional] 
**Subdivisions** | [**[]GeoSubdivision**](GeoSubdivision.md) |  | 

## Methods

### NewGeoMappingRule

`func NewGeoMappingRule(cidr string, subdivisions []GeoSubdivision, ) *GeoMappingRule`

NewGeoMappingRule instantiates a new GeoMappingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoMappingRuleWithDefaults

`func NewGeoMappingRuleWithDefaults() *GeoMappingRule`

NewGeoMappingRuleWithDefaults instantiates a new GeoMappingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccuracyRadius

`func (o *GeoMappingRule) GetAccuracyRadius() int64`

GetAccuracyRadius returns the AccuracyRadius field if non-nil, zero value otherwise.

### GetAccuracyRadiusOk

`func (o *GeoMappingRule) GetAccuracyRadiusOk() (*int64, bool)`

GetAccuracyRadiusOk returns a tuple with the AccuracyRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracyRadius

`func (o *GeoMappingRule) SetAccuracyRadius(v int64)`

SetAccuracyRadius sets AccuracyRadius field to given value.

### HasAccuracyRadius

`func (o *GeoMappingRule) HasAccuracyRadius() bool`

HasAccuracyRadius returns a boolean if a field has been set.

### GetCidr

`func (o *GeoMappingRule) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *GeoMappingRule) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *GeoMappingRule) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetCity

`func (o *GeoMappingRule) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GeoMappingRule) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GeoMappingRule) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *GeoMappingRule) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetContinent

`func (o *GeoMappingRule) GetContinent() string`

GetContinent returns the Continent field if non-nil, zero value otherwise.

### GetContinentOk

`func (o *GeoMappingRule) GetContinentOk() (*string, bool)`

GetContinentOk returns a tuple with the Continent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinent

`func (o *GeoMappingRule) SetContinent(v string)`

SetContinent sets Continent field to given value.

### HasContinent

`func (o *GeoMappingRule) HasContinent() bool`

HasContinent returns a boolean if a field has been set.

### GetContinentCode

`func (o *GeoMappingRule) GetContinentCode() string`

GetContinentCode returns the ContinentCode field if non-nil, zero value otherwise.

### GetContinentCodeOk

`func (o *GeoMappingRule) GetContinentCodeOk() (*string, bool)`

GetContinentCodeOk returns a tuple with the ContinentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinentCode

`func (o *GeoMappingRule) SetContinentCode(v string)`

SetContinentCode sets ContinentCode field to given value.

### HasContinentCode

`func (o *GeoMappingRule) HasContinentCode() bool`

HasContinentCode returns a boolean if a field has been set.

### GetCountry

`func (o *GeoMappingRule) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GeoMappingRule) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GeoMappingRule) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GeoMappingRule) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *GeoMappingRule) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *GeoMappingRule) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *GeoMappingRule) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *GeoMappingRule) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetLatitude

`func (o *GeoMappingRule) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *GeoMappingRule) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *GeoMappingRule) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *GeoMappingRule) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLeastSpecificSubdivision

`func (o *GeoMappingRule) GetLeastSpecificSubdivision() GeoSubdivision`

GetLeastSpecificSubdivision returns the LeastSpecificSubdivision field if non-nil, zero value otherwise.

### GetLeastSpecificSubdivisionOk

`func (o *GeoMappingRule) GetLeastSpecificSubdivisionOk() (*GeoSubdivision, bool)`

GetLeastSpecificSubdivisionOk returns a tuple with the LeastSpecificSubdivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeastSpecificSubdivision

`func (o *GeoMappingRule) SetLeastSpecificSubdivision(v GeoSubdivision)`

SetLeastSpecificSubdivision sets LeastSpecificSubdivision field to given value.

### HasLeastSpecificSubdivision

`func (o *GeoMappingRule) HasLeastSpecificSubdivision() bool`

HasLeastSpecificSubdivision returns a boolean if a field has been set.

### GetLongitude

`func (o *GeoMappingRule) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *GeoMappingRule) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *GeoMappingRule) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *GeoMappingRule) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetSubdivisions

`func (o *GeoMappingRule) GetSubdivisions() []GeoSubdivision`

GetSubdivisions returns the Subdivisions field if non-nil, zero value otherwise.

### GetSubdivisionsOk

`func (o *GeoMappingRule) GetSubdivisionsOk() (*[]GeoSubdivision, bool)`

GetSubdivisionsOk returns a tuple with the Subdivisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdivisions

`func (o *GeoMappingRule) SetSubdivisions(v []GeoSubdivision)`

SetSubdivisions sets Subdivisions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


