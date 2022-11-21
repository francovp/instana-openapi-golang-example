# \EventSettingsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebsiteAlertConfig**](EventSettingsApi.md#CreateWebsiteAlertConfig) | **Post** /api/events/settings/website-alert-configs | Create Website Alert Config
[**DeleteAlert**](EventSettingsApi.md#DeleteAlert) | **Delete** /api/events/settings/alerts/{id} | Delete Alert Configuration
[**DeleteAlertingChannel**](EventSettingsApi.md#DeleteAlertingChannel) | **Delete** /api/events/settings/alertingChannels/{id} | Delete alerting channel
[**DeleteBuiltInEventSpecification**](EventSettingsApi.md#DeleteBuiltInEventSpecification) | **Delete** /api/events/settings/event-specifications/built-in/{eventSpecificationId} | Delete built-in event specification
[**DeleteCustomEventSpecification**](EventSettingsApi.md#DeleteCustomEventSpecification) | **Delete** /api/events/settings/event-specifications/custom/{eventSpecificationId} | Delete custom event specification
[**DeleteCustomPayloadConfiguration**](EventSettingsApi.md#DeleteCustomPayloadConfiguration) | **Delete** /api/events/settings/custom-payload-configurations | Delete Custom Payload Configuration
[**DeleteWebsiteAlertConfig**](EventSettingsApi.md#DeleteWebsiteAlertConfig) | **Delete** /api/events/settings/website-alert-configs/{id} | Delete Website Alert Config
[**DisableBuiltInEventSpecification**](EventSettingsApi.md#DisableBuiltInEventSpecification) | **Post** /api/events/settings/event-specifications/built-in/{eventSpecificationId}/disable | Disable built-in event specification
[**DisableCustomEventSpecification**](EventSettingsApi.md#DisableCustomEventSpecification) | **Post** /api/events/settings/event-specifications/custom/{eventSpecificationId}/disable | Disable custom event specification
[**DisableWebsiteAlertConfig**](EventSettingsApi.md#DisableWebsiteAlertConfig) | **Put** /api/events/settings/website-alert-configs/{id}/disable | Disable Website Alert Config
[**EnableBuiltInEventSpecification**](EventSettingsApi.md#EnableBuiltInEventSpecification) | **Post** /api/events/settings/event-specifications/built-in/{eventSpecificationId}/enable | Enable built-in event specification
[**EnableCustomEventSpecification**](EventSettingsApi.md#EnableCustomEventSpecification) | **Post** /api/events/settings/event-specifications/custom/{eventSpecificationId}/enable | Enable custom event specification
[**EnableWebsiteAlertConfig**](EventSettingsApi.md#EnableWebsiteAlertConfig) | **Put** /api/events/settings/website-alert-configs/{id}/enable | Enable Website Alert Config
[**FindActiveWebsiteAlertConfigs**](EventSettingsApi.md#FindActiveWebsiteAlertConfigs) | **Get** /api/events/settings/website-alert-configs | All Website Alert Configs
[**FindWebsiteAlertConfig**](EventSettingsApi.md#FindWebsiteAlertConfig) | **Get** /api/events/settings/website-alert-configs/{id} | Get Website Alert Config
[**FindWebsiteAlertConfigVersions**](EventSettingsApi.md#FindWebsiteAlertConfigVersions) | **Get** /api/events/settings/website-alert-configs/{id}/versions | Get versions of Website Alert Config
[**GetAlert**](EventSettingsApi.md#GetAlert) | **Get** /api/events/settings/alerts/{id} | Find an Alert Configuration by ID
[**GetAlertingChannel**](EventSettingsApi.md#GetAlertingChannel) | **Get** /api/events/settings/alertingChannels/{id} | Alerting channel
[**GetAlertingChannels**](EventSettingsApi.md#GetAlertingChannels) | **Get** /api/events/settings/alertingChannels | All alerting channels
[**GetAlertingChannelsOverview**](EventSettingsApi.md#GetAlertingChannelsOverview) | **Get** /api/events/settings/alertingChannels/infos | Overview over all alerting channels
[**GetAlertingConfigurationInfos**](EventSettingsApi.md#GetAlertingConfigurationInfos) | **Get** /api/events/settings/alerts/infos | All alerting configuration info
[**GetAlerts**](EventSettingsApi.md#GetAlerts) | **Get** /api/events/settings/alerts | Get all Alert Configurations
[**GetBuiltInEventSpecification**](EventSettingsApi.md#GetBuiltInEventSpecification) | **Get** /api/events/settings/event-specifications/built-in/{eventSpecificationId} | Built-in event specifications
[**GetBuiltInEventSpecifications**](EventSettingsApi.md#GetBuiltInEventSpecifications) | **Get** /api/events/settings/event-specifications/built-in | All built-in event specification
[**GetCustomEventSpecification**](EventSettingsApi.md#GetCustomEventSpecification) | **Get** /api/events/settings/event-specifications/custom/{eventSpecificationId} | Custom event specification
[**GetCustomEventSpecifications**](EventSettingsApi.md#GetCustomEventSpecifications) | **Get** /api/events/settings/event-specifications/custom | All custom event specifications
[**GetCustomPayloadConfigurations**](EventSettingsApi.md#GetCustomPayloadConfigurations) | **Get** /api/events/settings/custom-payload-configurations | Get Custom Payload Configuration
[**GetCustomPayloadTagCatalog**](EventSettingsApi.md#GetCustomPayloadTagCatalog) | **Get** /api/events/settings/custom-payload-configurations/catalog | Get tag catalog for custom payload in alerting
[**GetEventSpecificationInfos**](EventSettingsApi.md#GetEventSpecificationInfos) | **Get** /api/events/settings/event-specifications/infos | Summary of all built-in and custom event specifications
[**GetEventSpecificationInfosByIds**](EventSettingsApi.md#GetEventSpecificationInfosByIds) | **Post** /api/events/settings/event-specifications/infos | All built-in and custom event specifications
[**GetSystemRules**](EventSettingsApi.md#GetSystemRules) | **Get** /api/events/settings/event-specifications/custom/systemRules | All system rules for custom event specifications
[**PostCustomEventSpecification**](EventSettingsApi.md#PostCustomEventSpecification) | **Post** /api/events/settings/event-specifications/custom | Create new custom event specification
[**PutAlert**](EventSettingsApi.md#PutAlert) | **Put** /api/events/settings/alerts/{id} | Update Alert Configuration
[**PutAlertingChannel**](EventSettingsApi.md#PutAlertingChannel) | **Put** /api/events/settings/alertingChannels/{id} | Update alerting channel
[**PutCustomEventSpecification**](EventSettingsApi.md#PutCustomEventSpecification) | **Put** /api/events/settings/event-specifications/custom/{eventSpecificationId} | Create or update custom event specification
[**RestoreWebsiteAlertConfig**](EventSettingsApi.md#RestoreWebsiteAlertConfig) | **Put** /api/events/settings/website-alert-configs/{id}/restore/{created} | Restore Website Alert Config
[**SendTestAlerting**](EventSettingsApi.md#SendTestAlerting) | **Put** /api/events/settings/alertingChannels/test | Test alerting channel
[**UpdateWebsiteAlertConfig**](EventSettingsApi.md#UpdateWebsiteAlertConfig) | **Post** /api/events/settings/website-alert-configs/{id} | Update Website Alert Config
[**UpsertCustomPayloadConfiguration**](EventSettingsApi.md#UpsertCustomPayloadConfiguration) | **Put** /api/events/settings/custom-payload-configurations | Create / Update Custom Payload Configuration



## CreateWebsiteAlertConfig

> WebsiteAlertConfigWithMetadata CreateWebsiteAlertConfig(ctx).WebsiteAlertConfig(websiteAlertConfig).Execute()

Create Website Alert Config



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    websiteAlertConfig := *openapiclient.NewWebsiteAlertConfig([]string{"AlertChannelIds_example"}, []openapiclient.CustomPayloadField{*openapiclient.NewCustomPayloadField("Key_example", "Type_example")}, "Description_example", int32(123), "Name_example", *openapiclient.NewWebsiteAlertRule("AlertType_example", "MetricName_example"), *openapiclient.NewTagFilterExpressionElement("Type_example"), *openapiclient.NewThreshold("Operator_example", "Type_example"), *openapiclient.NewWebsiteTimeThreshold("Type_example"), "WebsiteId_example") // WebsiteAlertConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.CreateWebsiteAlertConfig(context.Background()).WebsiteAlertConfig(websiteAlertConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.CreateWebsiteAlertConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebsiteAlertConfig`: WebsiteAlertConfigWithMetadata
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.CreateWebsiteAlertConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebsiteAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **websiteAlertConfig** | [**WebsiteAlertConfig**](WebsiteAlertConfig.md) |  | 

### Return type

[**WebsiteAlertConfigWithMetadata**](WebsiteAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlert

> DeleteAlert(ctx, id).Execute()

Delete Alert Configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.DeleteAlert(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.DeleteAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertingChannel

> DeleteAlertingChannel(ctx, id).Execute()

Delete alerting channel

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.DeleteAlertingChannel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.DeleteAlertingChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertingChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBuiltInEventSpecification

> DeleteBuiltInEventSpecification(ctx, eventSpecificationId).Execute()

Delete built-in event specification

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    eventSpecificationId := "eventSpecificationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.DeleteBuiltInEventSpecification(context.Background(), eventSpecificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.DeleteBuiltInEventSpecification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuiltInEventSpecificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomEventSpecification

> DeleteCustomEventSpecification(ctx, eventSpecificationId).Execute()

Delete custom event specification



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    eventSpecificationId := "eventSpecificationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.DeleteCustomEventSpecification(context.Background(), eventSpecificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.DeleteCustomEventSpecification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomEventSpecificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomPayloadConfiguration

> DeleteCustomPayloadConfiguration(ctx).Execute()

Delete Custom Payload Configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.DeleteCustomPayloadConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.DeleteCustomPayloadConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomPayloadConfigurationRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebsiteAlertConfig

> DeleteWebsiteAlertConfig(ctx, id).Execute()

Delete Website Alert Config

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.DeleteWebsiteAlertConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.DeleteWebsiteAlertConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebsiteAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableBuiltInEventSpecification

> BuiltInEventSpecificationWithLastUpdated DisableBuiltInEventSpecification(ctx, eventSpecificationId).Body(body).Execute()

Disable built-in event specification

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    eventSpecificationId := "eventSpecificationId_example" // string | 
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.DisableBuiltInEventSpecification(context.Background(), eventSpecificationId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.DisableBuiltInEventSpecification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableBuiltInEventSpecification`: BuiltInEventSpecificationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.DisableBuiltInEventSpecification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableBuiltInEventSpecificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

[**BuiltInEventSpecificationWithLastUpdated**](BuiltInEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableCustomEventSpecification

> CustomEventSpecificationWithLastUpdated DisableCustomEventSpecification(ctx, eventSpecificationId).Body(body).Execute()

Disable custom event specification

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    eventSpecificationId := "eventSpecificationId_example" // string | 
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.DisableCustomEventSpecification(context.Background(), eventSpecificationId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.DisableCustomEventSpecification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableCustomEventSpecification`: CustomEventSpecificationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.DisableCustomEventSpecification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableCustomEventSpecificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

[**CustomEventSpecificationWithLastUpdated**](CustomEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableWebsiteAlertConfig

> DisableWebsiteAlertConfig(ctx, id).Body(body).Execute()

Disable Website Alert Config

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.DisableWebsiteAlertConfig(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.DisableWebsiteAlertConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableWebsiteAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableBuiltInEventSpecification

> BuiltInEventSpecificationWithLastUpdated EnableBuiltInEventSpecification(ctx, eventSpecificationId).Body(body).Execute()

Enable built-in event specification

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    eventSpecificationId := "eventSpecificationId_example" // string | 
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.EnableBuiltInEventSpecification(context.Background(), eventSpecificationId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.EnableBuiltInEventSpecification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableBuiltInEventSpecification`: BuiltInEventSpecificationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.EnableBuiltInEventSpecification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableBuiltInEventSpecificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

[**BuiltInEventSpecificationWithLastUpdated**](BuiltInEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableCustomEventSpecification

> CustomEventSpecificationWithLastUpdated EnableCustomEventSpecification(ctx, eventSpecificationId).Body(body).Execute()

Enable custom event specification

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    eventSpecificationId := "eventSpecificationId_example" // string | 
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.EnableCustomEventSpecification(context.Background(), eventSpecificationId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.EnableCustomEventSpecification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableCustomEventSpecification`: CustomEventSpecificationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.EnableCustomEventSpecification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableCustomEventSpecificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

[**CustomEventSpecificationWithLastUpdated**](CustomEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableWebsiteAlertConfig

> EnableWebsiteAlertConfig(ctx, id).Body(body).Execute()

Enable Website Alert Config

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.EnableWebsiteAlertConfig(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.EnableWebsiteAlertConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableWebsiteAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindActiveWebsiteAlertConfigs

> []WebsiteAlertConfigWithMetadata FindActiveWebsiteAlertConfigs(ctx).WebsiteId(websiteId).AlertIds(alertIds).Execute()

All Website Alert Configs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    websiteId := "websiteId_example" // string |  (optional)
    alertIds := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.FindActiveWebsiteAlertConfigs(context.Background()).WebsiteId(websiteId).AlertIds(alertIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.FindActiveWebsiteAlertConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindActiveWebsiteAlertConfigs`: []WebsiteAlertConfigWithMetadata
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.FindActiveWebsiteAlertConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindActiveWebsiteAlertConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **websiteId** | **string** |  | 
 **alertIds** | **[]string** |  | 

### Return type

[**[]WebsiteAlertConfigWithMetadata**](WebsiteAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindWebsiteAlertConfig

> WebsiteAlertConfigWithMetadata FindWebsiteAlertConfig(ctx, id).ValidOn(validOn).Execute()

Get Website Alert Config



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    validOn := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.FindWebsiteAlertConfig(context.Background(), id).ValidOn(validOn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.FindWebsiteAlertConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindWebsiteAlertConfig`: WebsiteAlertConfigWithMetadata
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.FindWebsiteAlertConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindWebsiteAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validOn** | **int64** |  | 

### Return type

[**WebsiteAlertConfigWithMetadata**](WebsiteAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindWebsiteAlertConfigVersions

> []ConfigVersion FindWebsiteAlertConfigVersions(ctx, id).Execute()

Get versions of Website Alert Config



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.FindWebsiteAlertConfigVersions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.FindWebsiteAlertConfigVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindWebsiteAlertConfigVersions`: []ConfigVersion
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.FindWebsiteAlertConfigVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindWebsiteAlertConfigVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ConfigVersion**](ConfigVersion.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlert

> AlertingConfigurationWithLastUpdated GetAlert(ctx, id).Execute()

Find an Alert Configuration by ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.GetAlert(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.GetAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlert`: AlertingConfigurationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.GetAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertingConfigurationWithLastUpdated**](AlertingConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertingChannel

> AbstractIntegration GetAlertingChannel(ctx, id).Execute()

Alerting channel

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.GetAlertingChannel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.GetAlertingChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertingChannel`: AbstractIntegration
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.GetAlertingChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertingChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AbstractIntegration**](AbstractIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertingChannels

> []AbstractIntegration GetAlertingChannels(ctx).Ids(ids).Execute()

All alerting channels

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ids := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.GetAlertingChannels(context.Background()).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.GetAlertingChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertingChannels`: []AbstractIntegration
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.GetAlertingChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertingChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** |  | 

### Return type

[**[]AbstractIntegration**](AbstractIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertingChannelsOverview

> []IntegrationOverview GetAlertingChannelsOverview(ctx).Ids(ids).Execute()

Overview over all alerting channels

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ids := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.GetAlertingChannelsOverview(context.Background()).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.GetAlertingChannelsOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertingChannelsOverview`: []IntegrationOverview
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.GetAlertingChannelsOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertingChannelsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** |  | 

### Return type

[**[]IntegrationOverview**](IntegrationOverview.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertingConfigurationInfos

> []ValidatedAlertingChannelInputInfo GetAlertingConfigurationInfos(ctx).IntegrationId(integrationId).Execute()

All alerting configuration info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    integrationId := "integrationId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.GetAlertingConfigurationInfos(context.Background()).IntegrationId(integrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.GetAlertingConfigurationInfos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertingConfigurationInfos`: []ValidatedAlertingChannelInputInfo
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.GetAlertingConfigurationInfos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertingConfigurationInfosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationId** | **string** |  | 

### Return type

[**[]ValidatedAlertingChannelInputInfo**](ValidatedAlertingChannelInputInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlerts

> []ValidatedAlertingConfiguration GetAlerts(ctx).Execute()

Get all Alert Configurations

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.GetAlerts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.GetAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlerts`: []ValidatedAlertingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.GetAlerts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsRequest struct via the builder pattern


### Return type

[**[]ValidatedAlertingConfiguration**](ValidatedAlertingConfiguration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuiltInEventSpecification

> BuiltInEventSpecification GetBuiltInEventSpecification(ctx, eventSpecificationId).Execute()

Built-in event specifications

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    eventSpecificationId := "eventSpecificationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.GetBuiltInEventSpecification(context.Background(), eventSpecificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.GetBuiltInEventSpecification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuiltInEventSpecification`: BuiltInEventSpecification
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.GetBuiltInEventSpecification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuiltInEventSpecificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BuiltInEventSpecification**](BuiltInEventSpecification.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuiltInEventSpecifications

> []BuiltInEventSpecificationWithLastUpdated GetBuiltInEventSpecifications(ctx).Ids(ids).Execute()

All built-in event specification

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ids := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.GetBuiltInEventSpecifications(context.Background()).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.GetBuiltInEventSpecifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuiltInEventSpecifications`: []BuiltInEventSpecificationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.GetBuiltInEventSpecifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBuiltInEventSpecificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** |  | 

### Return type

[**[]BuiltInEventSpecificationWithLastUpdated**](BuiltInEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomEventSpecification

> CustomEventSpecificationWithLastUpdated GetCustomEventSpecification(ctx, eventSpecificationId).Execute()

Custom event specification

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    eventSpecificationId := "eventSpecificationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.GetCustomEventSpecification(context.Background(), eventSpecificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.GetCustomEventSpecification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomEventSpecification`: CustomEventSpecificationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.GetCustomEventSpecification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomEventSpecificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomEventSpecificationWithLastUpdated**](CustomEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomEventSpecifications

> []CustomEventSpecificationWithLastUpdated GetCustomEventSpecifications(ctx).Execute()

All custom event specifications

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.GetCustomEventSpecifications(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.GetCustomEventSpecifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomEventSpecifications`: []CustomEventSpecificationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.GetCustomEventSpecifications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomEventSpecificationsRequest struct via the builder pattern


### Return type

[**[]CustomEventSpecificationWithLastUpdated**](CustomEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomPayloadConfigurations

> CustomPayloadWithLastUpdated GetCustomPayloadConfigurations(ctx).Execute()

Get Custom Payload Configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.GetCustomPayloadConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.GetCustomPayloadConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomPayloadConfigurations`: CustomPayloadWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.GetCustomPayloadConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomPayloadConfigurationsRequest struct via the builder pattern


### Return type

[**CustomPayloadWithLastUpdated**](CustomPayloadWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomPayloadTagCatalog

> TagCatalog GetCustomPayloadTagCatalog(ctx).Execute()

Get tag catalog for custom payload in alerting

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.GetCustomPayloadTagCatalog(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.GetCustomPayloadTagCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomPayloadTagCatalog`: TagCatalog
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.GetCustomPayloadTagCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomPayloadTagCatalogRequest struct via the builder pattern


### Return type

[**TagCatalog**](TagCatalog.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventSpecificationInfos

> []EventSpecificationInfo GetEventSpecificationInfos(ctx).Execute()

Summary of all built-in and custom event specifications

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.GetEventSpecificationInfos(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.GetEventSpecificationInfos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventSpecificationInfos`: []EventSpecificationInfo
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.GetEventSpecificationInfos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventSpecificationInfosRequest struct via the builder pattern


### Return type

[**[]EventSpecificationInfo**](EventSpecificationInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventSpecificationInfosByIds

> []EventSpecificationInfo GetEventSpecificationInfosByIds(ctx).RequestBody(requestBody).Execute()

All built-in and custom event specifications



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.GetEventSpecificationInfosByIds(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.GetEventSpecificationInfosByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventSpecificationInfosByIds`: []EventSpecificationInfo
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.GetEventSpecificationInfosByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventSpecificationInfosByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

### Return type

[**[]EventSpecificationInfo**](EventSpecificationInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemRules

> []SystemRuleLabel GetSystemRules(ctx).Execute()

All system rules for custom event specifications

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.GetSystemRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.GetSystemRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemRules`: []SystemRuleLabel
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.GetSystemRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemRulesRequest struct via the builder pattern


### Return type

[**[]SystemRuleLabel**](SystemRuleLabel.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCustomEventSpecification

> CustomEventSpecificationWithLastUpdated PostCustomEventSpecification(ctx).CustomEventSpecification(customEventSpecification).Execute()

Create new custom event specification



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    customEventSpecification := *openapiclient.NewCustomEventSpecification("EntityType_example", "Name_example", []openapiclient.AbstractRule{*openapiclient.NewAbstractRule("RuleType_example")}) // CustomEventSpecification | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.PostCustomEventSpecification(context.Background()).CustomEventSpecification(customEventSpecification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.PostCustomEventSpecification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCustomEventSpecification`: CustomEventSpecificationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.PostCustomEventSpecification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCustomEventSpecificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customEventSpecification** | [**CustomEventSpecification**](CustomEventSpecification.md) |  | 

### Return type

[**CustomEventSpecificationWithLastUpdated**](CustomEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAlert

> AlertingConfigurationWithLastUpdated PutAlert(ctx, id).AlertingConfiguration(alertingConfiguration).Execute()

Update Alert Configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    alertingConfiguration := *openapiclient.NewAlertingConfiguration("AlertName_example", []openapiclient.StaticStringField{*openapiclient.NewStaticStringField("Value_example", "Key_example", "Type_example")}, *openapiclient.NewEventFilteringConfiguration(), "Id_example", []string{"IntegrationIds_example"}) // AlertingConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.PutAlert(context.Background(), id).AlertingConfiguration(alertingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.PutAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutAlert`: AlertingConfigurationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.PutAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertingConfiguration** | [**AlertingConfiguration**](AlertingConfiguration.md) |  | 

### Return type

[**AlertingConfigurationWithLastUpdated**](AlertingConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAlertingChannel

> PutAlertingChannel(ctx, id).AbstractIntegration(abstractIntegration).Execute()

Update alerting channel

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    abstractIntegration := *openapiclient.NewAbstractIntegration("Id_example", "Kind_example", "Name_example") // AbstractIntegration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.PutAlertingChannel(context.Background(), id).AbstractIntegration(abstractIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.PutAlertingChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAlertingChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **abstractIntegration** | [**AbstractIntegration**](AbstractIntegration.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCustomEventSpecification

> CustomEventSpecificationWithLastUpdated PutCustomEventSpecification(ctx, eventSpecificationId).CustomEventSpecification(customEventSpecification).AllowRestore(allowRestore).Execute()

Create or update custom event specification



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    eventSpecificationId := "eventSpecificationId_example" // string | 
    customEventSpecification := *openapiclient.NewCustomEventSpecification("EntityType_example", "Name_example", []openapiclient.AbstractRule{*openapiclient.NewAbstractRule("RuleType_example")}) // CustomEventSpecification | 
    allowRestore := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.PutCustomEventSpecification(context.Background(), eventSpecificationId).CustomEventSpecification(customEventSpecification).AllowRestore(allowRestore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.PutCustomEventSpecification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutCustomEventSpecification`: CustomEventSpecificationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.PutCustomEventSpecification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCustomEventSpecificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customEventSpecification** | [**CustomEventSpecification**](CustomEventSpecification.md) |  | 
 **allowRestore** | **bool** |  | 

### Return type

[**CustomEventSpecificationWithLastUpdated**](CustomEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreWebsiteAlertConfig

> RestoreWebsiteAlertConfig(ctx, id, created).Body(body).Execute()

Restore Website Alert Config

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    created := int64(789) // int64 | 
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.RestoreWebsiteAlertConfig(context.Background(), id, created).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.RestoreWebsiteAlertConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**created** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreWebsiteAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTestAlerting

> SendTestAlerting(ctx).AbstractIntegration(abstractIntegration).Execute()

Test alerting channel

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    abstractIntegration := *openapiclient.NewAbstractIntegration("Id_example", "Kind_example", "Name_example") // AbstractIntegration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.SendTestAlerting(context.Background()).AbstractIntegration(abstractIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.SendTestAlerting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendTestAlertingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **abstractIntegration** | [**AbstractIntegration**](AbstractIntegration.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebsiteAlertConfig

> WebsiteAlertConfigWithMetadata UpdateWebsiteAlertConfig(ctx, id).WebsiteAlertConfig(websiteAlertConfig).Execute()

Update Website Alert Config



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    websiteAlertConfig := *openapiclient.NewWebsiteAlertConfig([]string{"AlertChannelIds_example"}, []openapiclient.CustomPayloadField{*openapiclient.NewCustomPayloadField("Key_example", "Type_example")}, "Description_example", int32(123), "Name_example", *openapiclient.NewWebsiteAlertRule("AlertType_example", "MetricName_example"), *openapiclient.NewTagFilterExpressionElement("Type_example"), *openapiclient.NewThreshold("Operator_example", "Type_example"), *openapiclient.NewWebsiteTimeThreshold("Type_example"), "WebsiteId_example") // WebsiteAlertConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.UpdateWebsiteAlertConfig(context.Background(), id).WebsiteAlertConfig(websiteAlertConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.UpdateWebsiteAlertConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebsiteAlertConfig`: WebsiteAlertConfigWithMetadata
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.UpdateWebsiteAlertConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebsiteAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **websiteAlertConfig** | [**WebsiteAlertConfig**](WebsiteAlertConfig.md) |  | 

### Return type

[**WebsiteAlertConfigWithMetadata**](WebsiteAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertCustomPayloadConfiguration

> []CustomPayloadWithLastUpdated UpsertCustomPayloadConfiguration(ctx).CustomPayloadConfiguration(customPayloadConfiguration).Execute()

Create / Update Custom Payload Configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    customPayloadConfiguration := *openapiclient.NewCustomPayloadConfiguration([]openapiclient.CustomPayloadField{*openapiclient.NewCustomPayloadField("Key_example", "Type_example")}) // CustomPayloadConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventSettingsApi.UpsertCustomPayloadConfiguration(context.Background()).CustomPayloadConfiguration(customPayloadConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventSettingsApi.UpsertCustomPayloadConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpsertCustomPayloadConfiguration`: []CustomPayloadWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `EventSettingsApi.UpsertCustomPayloadConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertCustomPayloadConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customPayloadConfiguration** | [**CustomPayloadConfiguration**](CustomPayloadConfiguration.md) |  | 

### Return type

[**[]CustomPayloadWithLastUpdated**](CustomPayloadWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

