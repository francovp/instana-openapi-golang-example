# \MobileAppMetricsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMobileAppBeaconMetrics**](MobileAppMetricsApi.md#GetMobileAppBeaconMetrics) | **Post** /api/mobile-app-monitoring/metrics | Get beacon metrics
[**GetMobileAppBeaconMetricsV2**](MobileAppMetricsApi.md#GetMobileAppBeaconMetricsV2) | **Post** /api/mobile-app-monitoring/v2/metrics | Get beacon metrics
[**GetSession**](MobileAppMetricsApi.md#GetSession) | **Get** /api/mobile-app-monitoring/session | Get session



## GetMobileAppBeaconMetrics

> MobileAppMetricResult GetMobileAppBeaconMetrics(ctx).GetMobileAppMetrics(getMobileAppMetrics).Execute()

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
    getMobileAppMetrics := *openapiclient.NewGetMobileAppMetrics([]openapiclient.MobileAppMonitoringMetricsConfiguration{*openapiclient.NewMobileAppMonitoringMetricsConfiguration("Aggregation_example", "Metric_example")}, "Type_example") // GetMobileAppMetrics |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileAppMetricsApi.GetMobileAppBeaconMetrics(context.Background()).GetMobileAppMetrics(getMobileAppMetrics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileAppMetricsApi.GetMobileAppBeaconMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMobileAppBeaconMetrics`: MobileAppMetricResult
    fmt.Fprintf(os.Stdout, "Response from `MobileAppMetricsApi.GetMobileAppBeaconMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMobileAppBeaconMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getMobileAppMetrics** | [**GetMobileAppMetrics**](GetMobileAppMetrics.md) |  | 

### Return type

[**MobileAppMetricResult**](MobileAppMetricResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMobileAppBeaconMetricsV2

> MetricAPIResult GetMobileAppBeaconMetricsV2(ctx).GetMobileAppMetricsV2(getMobileAppMetricsV2).Execute()

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
    getMobileAppMetricsV2 := *openapiclient.NewGetMobileAppMetricsV2([]openapiclient.MobileAppMonitoringMetricsConfiguration{*openapiclient.NewMobileAppMonitoringMetricsConfiguration("Aggregation_example", "Metric_example")}, "Type_example") // GetMobileAppMetricsV2 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileAppMetricsApi.GetMobileAppBeaconMetricsV2(context.Background()).GetMobileAppMetricsV2(getMobileAppMetricsV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileAppMetricsApi.GetMobileAppBeaconMetricsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMobileAppBeaconMetricsV2`: MetricAPIResult
    fmt.Fprintf(os.Stdout, "Response from `MobileAppMetricsApi.GetMobileAppBeaconMetricsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMobileAppBeaconMetricsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getMobileAppMetricsV2** | [**GetMobileAppMetricsV2**](GetMobileAppMetricsV2.md) |  | 

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


## GetSession

> []MobileAppMonitoringBeacon GetSession(ctx, id, timestamp).Execute()

Get session

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
    resp, r, err := apiClient.MobileAppMetricsApi.GetSession(context.Background(), id, timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileAppMetricsApi.GetSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSession`: []MobileAppMonitoringBeacon
    fmt.Fprintf(os.Stdout, "Response from `MobileAppMetricsApi.GetSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**timestamp** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]MobileAppMonitoringBeacon**](MobileAppMonitoringBeacon.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

