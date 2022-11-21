# WebsiteAlertRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to **string** |  | [optional] 
**AlertType** | **string** |  | 
**MetricName** | **string** |  | 

## Methods

### NewWebsiteAlertRule

`func NewWebsiteAlertRule(alertType string, metricName string, ) *WebsiteAlertRule`

NewWebsiteAlertRule instantiates a new WebsiteAlertRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteAlertRuleWithDefaults

`func NewWebsiteAlertRuleWithDefaults() *WebsiteAlertRule`

NewWebsiteAlertRuleWithDefaults instantiates a new WebsiteAlertRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *WebsiteAlertRule) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *WebsiteAlertRule) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *WebsiteAlertRule) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *WebsiteAlertRule) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetAlertType

`func (o *WebsiteAlertRule) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *WebsiteAlertRule) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *WebsiteAlertRule) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.


### GetMetricName

`func (o *WebsiteAlertRule) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *WebsiteAlertRule) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *WebsiteAlertRule) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


