# HttpEndpointConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointNameByCollectedPathTemplateRuleEnabled** | Pointer to **bool** |  | [optional] 
**EndpointNameByFirstPathSegmentRuleEnabled** | Pointer to **bool** |  | [optional] 
**Rules** | [**[]HttpEndpointRule**](HttpEndpointRule.md) |  | 
**ServiceId** | **string** |  | 

## Methods

### NewHttpEndpointConfig

`func NewHttpEndpointConfig(rules []HttpEndpointRule, serviceId string, ) *HttpEndpointConfig`

NewHttpEndpointConfig instantiates a new HttpEndpointConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpEndpointConfigWithDefaults

`func NewHttpEndpointConfigWithDefaults() *HttpEndpointConfig`

NewHttpEndpointConfigWithDefaults instantiates a new HttpEndpointConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointNameByCollectedPathTemplateRuleEnabled

`func (o *HttpEndpointConfig) GetEndpointNameByCollectedPathTemplateRuleEnabled() bool`

GetEndpointNameByCollectedPathTemplateRuleEnabled returns the EndpointNameByCollectedPathTemplateRuleEnabled field if non-nil, zero value otherwise.

### GetEndpointNameByCollectedPathTemplateRuleEnabledOk

`func (o *HttpEndpointConfig) GetEndpointNameByCollectedPathTemplateRuleEnabledOk() (*bool, bool)`

GetEndpointNameByCollectedPathTemplateRuleEnabledOk returns a tuple with the EndpointNameByCollectedPathTemplateRuleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointNameByCollectedPathTemplateRuleEnabled

`func (o *HttpEndpointConfig) SetEndpointNameByCollectedPathTemplateRuleEnabled(v bool)`

SetEndpointNameByCollectedPathTemplateRuleEnabled sets EndpointNameByCollectedPathTemplateRuleEnabled field to given value.

### HasEndpointNameByCollectedPathTemplateRuleEnabled

`func (o *HttpEndpointConfig) HasEndpointNameByCollectedPathTemplateRuleEnabled() bool`

HasEndpointNameByCollectedPathTemplateRuleEnabled returns a boolean if a field has been set.

### GetEndpointNameByFirstPathSegmentRuleEnabled

`func (o *HttpEndpointConfig) GetEndpointNameByFirstPathSegmentRuleEnabled() bool`

GetEndpointNameByFirstPathSegmentRuleEnabled returns the EndpointNameByFirstPathSegmentRuleEnabled field if non-nil, zero value otherwise.

### GetEndpointNameByFirstPathSegmentRuleEnabledOk

`func (o *HttpEndpointConfig) GetEndpointNameByFirstPathSegmentRuleEnabledOk() (*bool, bool)`

GetEndpointNameByFirstPathSegmentRuleEnabledOk returns a tuple with the EndpointNameByFirstPathSegmentRuleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointNameByFirstPathSegmentRuleEnabled

`func (o *HttpEndpointConfig) SetEndpointNameByFirstPathSegmentRuleEnabled(v bool)`

SetEndpointNameByFirstPathSegmentRuleEnabled sets EndpointNameByFirstPathSegmentRuleEnabled field to given value.

### HasEndpointNameByFirstPathSegmentRuleEnabled

`func (o *HttpEndpointConfig) HasEndpointNameByFirstPathSegmentRuleEnabled() bool`

HasEndpointNameByFirstPathSegmentRuleEnabled returns a boolean if a field has been set.

### GetRules

`func (o *HttpEndpointConfig) GetRules() []HttpEndpointRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *HttpEndpointConfig) GetRulesOk() (*[]HttpEndpointRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *HttpEndpointConfig) SetRules(v []HttpEndpointRule)`

SetRules sets Rules field to given value.


### GetServiceId

`func (o *HttpEndpointConfig) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *HttpEndpointConfig) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *HttpEndpointConfig) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


