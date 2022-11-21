# \ApplicationAlertConfigurationApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationAlertConfig**](ApplicationAlertConfigurationApi.md#CreateApplicationAlertConfig) | **Post** /api/events/settings/application-alert-configs | Create Application Alert Config
[**DeleteApplicationAlertConfig**](ApplicationAlertConfigurationApi.md#DeleteApplicationAlertConfig) | **Delete** /api/events/settings/application-alert-configs/{id} | Delete Application Alert Config
[**DisableApplicationAlertConfig**](ApplicationAlertConfigurationApi.md#DisableApplicationAlertConfig) | **Put** /api/events/settings/application-alert-configs/{id}/disable | Disable Application Alert Config
[**EnableApplicationAlertConfig**](ApplicationAlertConfigurationApi.md#EnableApplicationAlertConfig) | **Put** /api/events/settings/application-alert-configs/{id}/enable | Enable Application Alert Config
[**FindActiveApplicationAlertConfigs**](ApplicationAlertConfigurationApi.md#FindActiveApplicationAlertConfigs) | **Get** /api/events/settings/application-alert-configs | All Application Alert Configs
[**FindApplicationAlertConfig**](ApplicationAlertConfigurationApi.md#FindApplicationAlertConfig) | **Get** /api/events/settings/application-alert-configs/{id} | Get Application Alert Config
[**FindApplicationAlertConfigVersions**](ApplicationAlertConfigurationApi.md#FindApplicationAlertConfigVersions) | **Get** /api/events/settings/application-alert-configs/{id}/versions | Get versions of Application Alert Config
[**RestoreApplicationAlertConfig**](ApplicationAlertConfigurationApi.md#RestoreApplicationAlertConfig) | **Put** /api/events/settings/application-alert-configs/{id}/restore/{created} | Restore Application Alert Config
[**UpdateApplicationAlertConfig**](ApplicationAlertConfigurationApi.md#UpdateApplicationAlertConfig) | **Post** /api/events/settings/application-alert-configs/{id} | Update Application Alert Config



## CreateApplicationAlertConfig

> ApplicationAlertConfigWithMetadata CreateApplicationAlertConfig(ctx).ApplicationAlertConfig(applicationAlertConfig).Execute()

Create Application Alert Config



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
    applicationAlertConfig := *openapiclient.NewApplicationAlertConfig([]string{"AlertChannelIds_example"}, map[string]ApplicationNode{"key": *openapiclient.NewApplicationNode("ApplicationId_example", map[string]ServiceNode{"key": *openapiclient.NewServiceNode(map[string]EndpointNode{"key": *openapiclient.NewEndpointNode("EndpointId_example")}, "ServiceId_example")})}, "BoundaryScope_example", []openapiclient.CustomPayloadField{*openapiclient.NewCustomPayloadField("Key_example", "Type_example")}, "Description_example", "EvaluationType_example", int32(123), "Name_example", *openapiclient.NewApplicationAlertRule("AlertType_example", "MetricName_example"), *openapiclient.NewTagFilterExpressionElement("Type_example"), *openapiclient.NewThreshold("Operator_example", "Type_example"), *openapiclient.NewApplicationTimeThreshold("Type_example")) // ApplicationAlertConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAlertConfigurationApi.CreateApplicationAlertConfig(context.Background()).ApplicationAlertConfig(applicationAlertConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAlertConfigurationApi.CreateApplicationAlertConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationAlertConfig`: ApplicationAlertConfigWithMetadata
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAlertConfigurationApi.CreateApplicationAlertConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationAlertConfig** | [**ApplicationAlertConfig**](ApplicationAlertConfig.md) |  | 

### Return type

[**ApplicationAlertConfigWithMetadata**](ApplicationAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationAlertConfig

> DeleteApplicationAlertConfig(ctx, id).Execute()

Delete Application Alert Config

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
    resp, r, err := apiClient.ApplicationAlertConfigurationApi.DeleteApplicationAlertConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAlertConfigurationApi.DeleteApplicationAlertConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApplicationAlertConfigRequest struct via the builder pattern


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


## DisableApplicationAlertConfig

> DisableApplicationAlertConfig(ctx, id).Body(body).Execute()

Disable Application Alert Config

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
    resp, r, err := apiClient.ApplicationAlertConfigurationApi.DisableApplicationAlertConfig(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAlertConfigurationApi.DisableApplicationAlertConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDisableApplicationAlertConfigRequest struct via the builder pattern


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


## EnableApplicationAlertConfig

> EnableApplicationAlertConfig(ctx, id).Body(body).Execute()

Enable Application Alert Config

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
    resp, r, err := apiClient.ApplicationAlertConfigurationApi.EnableApplicationAlertConfig(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAlertConfigurationApi.EnableApplicationAlertConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnableApplicationAlertConfigRequest struct via the builder pattern


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


## FindActiveApplicationAlertConfigs

> []ApplicationAlertConfigWithMetadata FindActiveApplicationAlertConfigs(ctx).ApplicationId(applicationId).AlertIds(alertIds).Execute()

All Application Alert Configs



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
    applicationId := "applicationId_example" // string |  (optional)
    alertIds := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAlertConfigurationApi.FindActiveApplicationAlertConfigs(context.Background()).ApplicationId(applicationId).AlertIds(alertIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAlertConfigurationApi.FindActiveApplicationAlertConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindActiveApplicationAlertConfigs`: []ApplicationAlertConfigWithMetadata
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAlertConfigurationApi.FindActiveApplicationAlertConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindActiveApplicationAlertConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | **string** |  | 
 **alertIds** | **[]string** |  | 

### Return type

[**[]ApplicationAlertConfigWithMetadata**](ApplicationAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindApplicationAlertConfig

> ApplicationAlertConfigWithMetadata FindApplicationAlertConfig(ctx, id).ValidOn(validOn).Execute()

Get Application Alert Config



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
    resp, r, err := apiClient.ApplicationAlertConfigurationApi.FindApplicationAlertConfig(context.Background(), id).ValidOn(validOn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAlertConfigurationApi.FindApplicationAlertConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindApplicationAlertConfig`: ApplicationAlertConfigWithMetadata
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAlertConfigurationApi.FindApplicationAlertConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindApplicationAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validOn** | **int64** |  | 

### Return type

[**ApplicationAlertConfigWithMetadata**](ApplicationAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindApplicationAlertConfigVersions

> []ConfigVersion FindApplicationAlertConfigVersions(ctx, id).Execute()

Get versions of Application Alert Config



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
    resp, r, err := apiClient.ApplicationAlertConfigurationApi.FindApplicationAlertConfigVersions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAlertConfigurationApi.FindApplicationAlertConfigVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindApplicationAlertConfigVersions`: []ConfigVersion
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAlertConfigurationApi.FindApplicationAlertConfigVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindApplicationAlertConfigVersionsRequest struct via the builder pattern


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


## RestoreApplicationAlertConfig

> RestoreApplicationAlertConfig(ctx, id, created).Body(body).Execute()

Restore Application Alert Config

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
    resp, r, err := apiClient.ApplicationAlertConfigurationApi.RestoreApplicationAlertConfig(context.Background(), id, created).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAlertConfigurationApi.RestoreApplicationAlertConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRestoreApplicationAlertConfigRequest struct via the builder pattern


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


## UpdateApplicationAlertConfig

> ApplicationAlertConfigWithMetadata UpdateApplicationAlertConfig(ctx, id).ApplicationAlertConfig(applicationAlertConfig).Execute()

Update Application Alert Config



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
    applicationAlertConfig := *openapiclient.NewApplicationAlertConfig([]string{"AlertChannelIds_example"}, map[string]ApplicationNode{"key": *openapiclient.NewApplicationNode("ApplicationId_example", map[string]ServiceNode{"key": *openapiclient.NewServiceNode(map[string]EndpointNode{"key": *openapiclient.NewEndpointNode("EndpointId_example")}, "ServiceId_example")})}, "BoundaryScope_example", []openapiclient.CustomPayloadField{*openapiclient.NewCustomPayloadField("Key_example", "Type_example")}, "Description_example", "EvaluationType_example", int32(123), "Name_example", *openapiclient.NewApplicationAlertRule("AlertType_example", "MetricName_example"), *openapiclient.NewTagFilterExpressionElement("Type_example"), *openapiclient.NewThreshold("Operator_example", "Type_example"), *openapiclient.NewApplicationTimeThreshold("Type_example")) // ApplicationAlertConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAlertConfigurationApi.UpdateApplicationAlertConfig(context.Background(), id).ApplicationAlertConfig(applicationAlertConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAlertConfigurationApi.UpdateApplicationAlertConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationAlertConfig`: ApplicationAlertConfigWithMetadata
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAlertConfigurationApi.UpdateApplicationAlertConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationAlertConfig** | [**ApplicationAlertConfig**](ApplicationAlertConfig.md) |  | 

### Return type

[**ApplicationAlertConfigWithMetadata**](ApplicationAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

