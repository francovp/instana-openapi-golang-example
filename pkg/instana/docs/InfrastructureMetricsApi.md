# \InfrastructureMetricsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInfrastructureMetrics**](InfrastructureMetricsApi.md#GetInfrastructureMetrics) | **Post** /api/infrastructure-monitoring/metrics | Get infrastructure metrics



## GetInfrastructureMetrics

> InfrastructureMetricResult GetInfrastructureMetrics(ctx).Offline(offline).GetCombinedMetrics(getCombinedMetrics).Execute()

Get infrastructure metrics



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
    offline := true // bool |  (optional)
    getCombinedMetrics := *openapiclient.NewGetCombinedMetrics([]string{"Metrics_example"}, "Plugin_example") // GetCombinedMetrics |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfrastructureMetricsApi.GetInfrastructureMetrics(context.Background()).Offline(offline).GetCombinedMetrics(getCombinedMetrics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureMetricsApi.GetInfrastructureMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfrastructureMetrics`: InfrastructureMetricResult
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureMetricsApi.GetInfrastructureMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInfrastructureMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offline** | **bool** |  | 
 **getCombinedMetrics** | [**GetCombinedMetrics**](GetCombinedMetrics.md) |  | 

### Return type

[**InfrastructureMetricResult**](InfrastructureMetricResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

