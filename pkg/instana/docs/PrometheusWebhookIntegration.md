# PrometheusWebhookIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Receiver** | Pointer to **string** |  | [optional] 
**WebhookUrl** | **string** |  | 

## Methods

### NewPrometheusWebhookIntegration

`func NewPrometheusWebhookIntegration(webhookUrl string, ) *PrometheusWebhookIntegration`

NewPrometheusWebhookIntegration instantiates a new PrometheusWebhookIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusWebhookIntegrationWithDefaults

`func NewPrometheusWebhookIntegrationWithDefaults() *PrometheusWebhookIntegration`

NewPrometheusWebhookIntegrationWithDefaults instantiates a new PrometheusWebhookIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceiver

`func (o *PrometheusWebhookIntegration) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *PrometheusWebhookIntegration) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *PrometheusWebhookIntegration) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *PrometheusWebhookIntegration) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *PrometheusWebhookIntegration) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *PrometheusWebhookIntegration) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *PrometheusWebhookIntegration) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


