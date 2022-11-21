# \WebsiteMetricsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBeaconMetrics**](WebsiteMetricsApi.md#GetBeaconMetrics) | **Post** /api/website-monitoring/metrics | Get beacon metrics
[**GetBeaconMetricsV2**](WebsiteMetricsApi.md#GetBeaconMetricsV2) | **Post** /api/website-monitoring/v2/metrics | Get beacon metrics
[**GetPageLoad**](WebsiteMetricsApi.md#GetPageLoad) | **Get** /api/website-monitoring/page-load | Get page load



## GetBeaconMetrics

> WebsiteMetricResult GetBeaconMetrics(ctx).GetWebsiteMetrics(getWebsiteMetrics).Execute()

Get beacon metrics

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
    getWebsiteMetrics := *openapiclient.NewGetWebsiteMetrics([]openapiclient.WebsiteMonitoringMetricsConfiguration{*openapiclient.NewWebsiteMonitoringMetricsConfiguration("Aggregation_example", "Metric_example")}, "Type_example") // GetWebsiteMetrics |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebsiteMetricsApi.GetBeaconMetrics(context.Background()).GetWebsiteMetrics(getWebsiteMetrics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteMetricsApi.GetBeaconMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBeaconMetrics`: WebsiteMetricResult
    fmt.Fprintf(os.Stdout, "Response from `WebsiteMetricsApi.GetBeaconMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBeaconMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getWebsiteMetrics** | [**GetWebsiteMetrics**](GetWebsiteMetrics.md) |  | 

### Return type

[**WebsiteMetricResult**](WebsiteMetricResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBeaconMetricsV2

> MetricAPIResult GetBeaconMetricsV2(ctx).GetWebsiteMetricsV2(getWebsiteMetricsV2).Execute()

Get beacon metrics

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
    getWebsiteMetricsV2 := *openapiclient.NewGetWebsiteMetricsV2([]openapiclient.WebsiteMonitoringMetricsConfiguration{*openapiclient.NewWebsiteMonitoringMetricsConfiguration("Aggregation_example", "Metric_example")}, "Type_example") // GetWebsiteMetricsV2 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebsiteMetricsApi.GetBeaconMetricsV2(context.Background()).GetWebsiteMetricsV2(getWebsiteMetricsV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteMetricsApi.GetBeaconMetricsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBeaconMetricsV2`: MetricAPIResult
    fmt.Fprintf(os.Stdout, "Response from `WebsiteMetricsApi.GetBeaconMetricsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBeaconMetricsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getWebsiteMetricsV2** | [**GetWebsiteMetricsV2**](GetWebsiteMetricsV2.md) |  | 

### Return type

[**MetricAPIResult**](MetricAPIResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageLoad

> []WebsiteMonitoringBeacon GetPageLoad(ctx, id, timestamp).Execute()

Get page load

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
    timestamp := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebsiteMetricsApi.GetPageLoad(context.Background(), id, timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteMetricsApi.GetPageLoad``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageLoad`: []WebsiteMonitoringBeacon
    fmt.Fprintf(os.Stdout, "Response from `WebsiteMetricsApi.GetPageLoad`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**timestamp** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageLoadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]WebsiteMonitoringBeacon**](WebsiteMonitoringBeacon.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

