# \ApplicationMetricsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationDataMetricsV2**](ApplicationMetricsApi.md#GetApplicationDataMetricsV2) | **Post** /api/application-monitoring/v2/metrics | Get Application Data Metrics
[**GetApplicationMetrics**](ApplicationMetricsApi.md#GetApplicationMetrics) | **Post** /api/application-monitoring/metrics/applications | Get Application Metrics
[**GetEndpointsMetrics**](ApplicationMetricsApi.md#GetEndpointsMetrics) | **Post** /api/application-monitoring/metrics/endpoints | Get Endpoint metrics
[**GetServicesMetrics**](ApplicationMetricsApi.md#GetServicesMetrics) | **Post** /api/application-monitoring/metrics/services | Get Service metrics



## GetApplicationDataMetricsV2

> MetricAPIResult GetApplicationDataMetricsV2(ctx).GetApplicationMetrics(getApplicationMetrics).Execute()

Get Application Data Metrics

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
    getApplicationMetrics := *openapiclient.NewGetApplicationMetrics([]openapiclient.AppDataMetricConfiguration{*openapiclient.NewAppDataMetricConfiguration("Aggregation_example", "Metric_example")}, *openapiclient.NewTimeFrame()) // GetApplicationMetrics |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationMetricsApi.GetApplicationDataMetricsV2(context.Background()).GetApplicationMetrics(getApplicationMetrics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationMetricsApi.GetApplicationDataMetricsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationDataMetricsV2`: MetricAPIResult
    fmt.Fprintf(os.Stdout, "Response from `ApplicationMetricsApi.GetApplicationDataMetricsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationDataMetricsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getApplicationMetrics** | [**GetApplicationMetrics**](GetApplicationMetrics.md) |  | 

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


## GetApplicationMetrics

> ApplicationMetricResult GetApplicationMetrics(ctx).FillTimeSeries(fillTimeSeries).GetApplications(getApplications).Execute()

Get Application Metrics

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
    fillTimeSeries := true // bool |  (optional)
    getApplications := *openapiclient.NewGetApplications([]openapiclient.AppDataMetricConfiguration{*openapiclient.NewAppDataMetricConfiguration("Aggregation_example", "Metric_example")}) // GetApplications |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationMetricsApi.GetApplicationMetrics(context.Background()).FillTimeSeries(fillTimeSeries).GetApplications(getApplications).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationMetricsApi.GetApplicationMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationMetrics`: ApplicationMetricResult
    fmt.Fprintf(os.Stdout, "Response from `ApplicationMetricsApi.GetApplicationMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fillTimeSeries** | **bool** |  | 
 **getApplications** | [**GetApplications**](GetApplications.md) |  | 

### Return type

[**ApplicationMetricResult**](ApplicationMetricResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointsMetrics

> EndpointMetricResult GetEndpointsMetrics(ctx).FillTimeSeries(fillTimeSeries).GetEndpoints(getEndpoints).Execute()

Get Endpoint metrics

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
    fillTimeSeries := true // bool |  (optional)
    getEndpoints := *openapiclient.NewGetEndpoints([]openapiclient.AppDataMetricConfiguration{*openapiclient.NewAppDataMetricConfiguration("Aggregation_example", "Metric_example")}) // GetEndpoints |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationMetricsApi.GetEndpointsMetrics(context.Background()).FillTimeSeries(fillTimeSeries).GetEndpoints(getEndpoints).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationMetricsApi.GetEndpointsMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEndpointsMetrics`: EndpointMetricResult
    fmt.Fprintf(os.Stdout, "Response from `ApplicationMetricsApi.GetEndpointsMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointsMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fillTimeSeries** | **bool** |  | 
 **getEndpoints** | [**GetEndpoints**](GetEndpoints.md) |  | 

### Return type

[**EndpointMetricResult**](EndpointMetricResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicesMetrics

> ServiceMetricResult GetServicesMetrics(ctx).FillTimeSeries(fillTimeSeries).GetServices(getServices).Execute()

Get Service metrics

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
    fillTimeSeries := true // bool |  (optional)
    getServices := *openapiclient.NewGetServices([]openapiclient.AppDataMetricConfiguration{*openapiclient.NewAppDataMetricConfiguration("Aggregation_example", "Metric_example")}) // GetServices |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationMetricsApi.GetServicesMetrics(context.Background()).FillTimeSeries(fillTimeSeries).GetServices(getServices).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationMetricsApi.GetServicesMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServicesMetrics`: ServiceMetricResult
    fmt.Fprintf(os.Stdout, "Response from `ApplicationMetricsApi.GetServicesMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fillTimeSeries** | **bool** |  | 
 **getServices** | [**GetServices**](GetServices.md) |  | 

### Return type

[**ServiceMetricResult**](ServiceMetricResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

