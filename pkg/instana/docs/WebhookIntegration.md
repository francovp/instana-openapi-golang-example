# WebhookIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to **[]string** |  | [optional] 
**WebhookUrls** | **[]string** |  | 

## Methods

### NewWebhookIntegration

`func NewWebhookIntegration(webhookUrls []string, ) *WebhookIntegration`

NewWebhookIntegration instantiates a new WebhookIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookIntegrationWithDefaults

`func NewWebhookIntegrationWithDefaults() *WebhookIntegration`

NewWebhookIntegrationWithDefaults instantiates a new WebhookIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *WebhookIntegration) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebhookIntegration) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WebhookIntegration) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *WebhookIntegration) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetWebhookUrls

`func (o *WebhookIntegration) GetWebhookUrls() []string`

GetWebhookUrls returns the WebhookUrls field if non-nil, zero value otherwise.

### GetWebhookUrlsOk

`func (o *WebhookIntegration) GetWebhookUrlsOk() (*[]string, bool)`

GetWebhookUrlsOk returns a tuple with the WebhookUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrls

`func (o *WebhookIntegration) SetWebhookUrls(v []string)`

SetWebhookUrls sets WebhookUrls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


