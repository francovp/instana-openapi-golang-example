# GeoLocationConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeoDetailRemoval** | **string** |  | 
**GeoMappingRules** | Pointer to [**[]GeoMappingRule**](GeoMappingRule.md) |  | [optional] 

## Methods

### NewGeoLocationConfiguration

`func NewGeoLocationConfiguration(geoDetailRemoval string, ) *GeoLocationConfiguration`

NewGeoLocationConfiguration instantiates a new GeoLocationConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoLocationConfigurationWithDefaults

`func NewGeoLocationConfigurationWithDefaults() *GeoLocationConfiguration`

NewGeoLocationConfigurationWithDefaults instantiates a new GeoLocationConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeoDetailRemoval

`func (o *GeoLocationConfiguration) GetGeoDetailRemoval() string`

GetGeoDetailRemoval returns the GeoDetailRemoval field if non-nil, zero value otherwise.

### GetGeoDetailRemovalOk

`func (o *GeoLocationConfiguration) GetGeoDetailRemovalOk() (*string, bool)`

GetGeoDetailRemovalOk returns a tuple with the GeoDetailRemoval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoDetailRemoval

`func (o *GeoLocationConfiguration) SetGeoDetailRemoval(v string)`

SetGeoDetailRemoval sets GeoDetailRemoval field to given value.


### GetGeoMappingRules

`func (o *GeoLocationConfiguration) GetGeoMappingRules() []GeoMappingRule`

GetGeoMappingRules returns the GeoMappingRules field if non-nil, zero value otherwise.

### GetGeoMappingRulesOk

`func (o *GeoLocationConfiguration) GetGeoMappingRulesOk() (*[]GeoMappingRule, bool)`

GetGeoMappingRulesOk returns a tuple with the GeoMappingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoMappingRules

`func (o *GeoLocationConfiguration) SetGeoMappingRules(v []GeoMappingRule)`

SetGeoMappingRules sets GeoMappingRules field to given value.

### HasGeoMappingRules

`func (o *GeoLocationConfiguration) HasGeoMappingRules() bool`

HasGeoMappingRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


