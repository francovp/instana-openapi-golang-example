# \InfrastructureCatalogApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInfrastructureCatalogMetrics**](InfrastructureCatalogApi.md#GetInfrastructureCatalogMetrics) | **Get** /api/infrastructure-monitoring/catalog/metrics/{plugin} | Get metric catalog
[**GetInfrastructureCatalogPlugins**](InfrastructureCatalogApi.md#GetInfrastructureCatalogPlugins) | **Get** /api/infrastructure-monitoring/catalog/plugins | Get plugin catalog
[**GetInfrastructureCatalogPluginsWithCustomMetrics**](InfrastructureCatalogApi.md#GetInfrastructureCatalogPluginsWithCustomMetrics) | **Get** /api/infrastructure-monitoring/catalog/plugins-with-custom-metrics | Get all plugins with custom metrics catalog
[**GetInfrastructureCatalogSearchFields**](InfrastructureCatalogApi.md#GetInfrastructureCatalogSearchFields) | **Get** /api/infrastructure-monitoring/catalog/search | get search field catalog



## GetInfrastructureCatalogMetrics

> []MetricInstance GetInfrastructureCatalogMetrics(ctx, plugin).Filter(filter).Execute()

Get metric catalog



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
    plugin := "plugin_example" // string | 
    filter := "filter_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfrastructureCatalogApi.GetInfrastructureCatalogMetrics(context.Background(), plugin).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureCatalogApi.GetInfrastructureCatalogMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfrastructureCatalogMetrics`: []MetricInstance
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureCatalogApi.GetInfrastructureCatalogMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plugin** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureCatalogMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** |  | 

### Return type

[**[]MetricInstance**](MetricInstance.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureCatalogPlugins

> []PluginResult GetInfrastructureCatalogPlugins(ctx).Execute()

Get plugin catalog



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
    resp, r, err := apiClient.InfrastructureCatalogApi.GetInfrastructureCatalogPlugins(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureCatalogApi.GetInfrastructureCatalogPlugins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfrastructureCatalogPlugins`: []PluginResult
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureCatalogApi.GetInfrastructureCatalogPlugins`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureCatalogPluginsRequest struct via the builder pattern


### Return type

[**[]PluginResult**](PluginResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureCatalogPluginsWithCustomMetrics

> []PluginResult GetInfrastructureCatalogPluginsWithCustomMetrics(ctx).Execute()

Get all plugins with custom metrics catalog

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
    resp, r, err := apiClient.InfrastructureCatalogApi.GetInfrastructureCatalogPluginsWithCustomMetrics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureCatalogApi.GetInfrastructureCatalogPluginsWithCustomMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfrastructureCatalogPluginsWithCustomMetrics`: []PluginResult
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureCatalogApi.GetInfrastructureCatalogPluginsWithCustomMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureCatalogPluginsWithCustomMetricsRequest struct via the builder pattern


### Return type

[**[]PluginResult**](PluginResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfrastructureCatalogSearchFields

> []SearchFieldResult GetInfrastructureCatalogSearchFields(ctx).Execute()

get search field catalog



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
    resp, r, err := apiClient.InfrastructureCatalogApi.GetInfrastructureCatalogSearchFields(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureCatalogApi.GetInfrastructureCatalogSearchFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfrastructureCatalogSearchFields`: []SearchFieldResult
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureCatalogApi.GetInfrastructureCatalogSearchFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureCatalogSearchFieldsRequest struct via the builder pattern


### Return type

[**[]SearchFieldResult**](SearchFieldResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

