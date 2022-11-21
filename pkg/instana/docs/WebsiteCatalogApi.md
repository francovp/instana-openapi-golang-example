# \WebsiteCatalogApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebsiteCatalogMetrics**](WebsiteCatalogApi.md#GetWebsiteCatalogMetrics) | **Get** /api/website-monitoring/catalog/metrics | Metric catalog
[**GetWebsiteCatalogTags**](WebsiteCatalogApi.md#GetWebsiteCatalogTags) | **Get** /api/website-monitoring/catalog/tags | Get all existing website tags
[**GetWebsiteTagCatalog**](WebsiteCatalogApi.md#GetWebsiteTagCatalog) | **Get** /api/website-monitoring/catalog | Get website tag catalog



## GetWebsiteCatalogMetrics

> []WebsiteMonitoringMetricDescription GetWebsiteCatalogMetrics(ctx).Execute()

Metric catalog



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
    resp, r, err := apiClient.WebsiteCatalogApi.GetWebsiteCatalogMetrics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteCatalogApi.GetWebsiteCatalogMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebsiteCatalogMetrics`: []WebsiteMonitoringMetricDescription
    fmt.Fprintf(os.Stdout, "Response from `WebsiteCatalogApi.GetWebsiteCatalogMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebsiteCatalogMetricsRequest struct via the builder pattern


### Return type

[**[]WebsiteMonitoringMetricDescription**](WebsiteMonitoringMetricDescription.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebsiteCatalogTags

> []Tag GetWebsiteCatalogTags(ctx).Execute()

Get all existing website tags



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
    resp, r, err := apiClient.WebsiteCatalogApi.GetWebsiteCatalogTags(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteCatalogApi.GetWebsiteCatalogTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebsiteCatalogTags`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `WebsiteCatalogApi.GetWebsiteCatalogTags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebsiteCatalogTagsRequest struct via the builder pattern


### Return type

[**[]Tag**](Tag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebsiteTagCatalog

> TagCatalog GetWebsiteTagCatalog(ctx).BeaconType(beaconType).UseCase(useCase).Execute()

Get website tag catalog

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
    beaconType := "beaconType_example" // string | 
    useCase := "useCase_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebsiteCatalogApi.GetWebsiteTagCatalog(context.Background()).BeaconType(beaconType).UseCase(useCase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteCatalogApi.GetWebsiteTagCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebsiteTagCatalog`: TagCatalog
    fmt.Fprintf(os.Stdout, "Response from `WebsiteCatalogApi.GetWebsiteTagCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebsiteTagCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **beaconType** | **string** |  | 
 **useCase** | **string** |  | 

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

