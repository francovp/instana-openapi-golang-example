# HttpEndpointRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**PathSegments** | [**[]HttpPathSegmentMatchingRule**](HttpPathSegmentMatchingRule.md) |  | 
**TestCases** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHttpEndpointRule

`func NewHttpEndpointRule(pathSegments []HttpPathSegmentMatchingRule, ) *HttpEndpointRule`

NewHttpEndpointRule instantiates a new HttpEndpointRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpEndpointRuleWithDefaults

`func NewHttpEndpointRuleWithDefaults() *HttpEndpointRule`

NewHttpEndpointRuleWithDefaults instantiates a new HttpEndpointRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *HttpEndpointRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HttpEndpointRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HttpEndpointRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *HttpEndpointRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPathSegments

`func (o *HttpEndpointRule) GetPathSegments() []HttpPathSegmentMatchingRule`

GetPathSegments returns the PathSegments field if non-nil, zero value otherwise.

### GetPathSegmentsOk

`func (o *HttpEndpointRule) GetPathSegmentsOk() (*[]HttpPathSegmentMatchingRule, bool)`

GetPathSegmentsOk returns a tuple with the PathSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathSegments

`func (o *HttpEndpointRule) SetPathSegments(v []HttpPathSegmentMatchingRule)`

SetPathSegments sets PathSegments field to given value.


### GetTestCases

`func (o *HttpEndpointRule) GetTestCases() []string`

GetTestCases returns the TestCases field if non-nil, zero value otherwise.

### GetTestCasesOk

`func (o *HttpEndpointRule) GetTestCasesOk() (*[]string, bool)`

GetTestCasesOk returns a tuple with the TestCases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCases

`func (o *HttpEndpointRule) SetTestCases(v []string)`

SetTestCases sets TestCases field to given value.

### HasTestCases

`func (o *HttpEndpointRule) HasTestCases() bool`

HasTestCases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


