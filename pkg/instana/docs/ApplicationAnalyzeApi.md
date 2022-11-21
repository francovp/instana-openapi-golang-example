# \ApplicationAnalyzeApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCallGroup**](ApplicationAnalyzeApi.md#GetCallGroup) | **Post** /api/application-monitoring/analyze/call-groups | Get grouped call metrics
[**GetCorrelatedTraces**](ApplicationAnalyzeApi.md#GetCorrelatedTraces) | **Get** /api/application-monitoring/analyze/backend-correlation | Resolve backend trace IDs using correlation IDs from website and mobile app monitoring beacons.
[**GetTrace**](ApplicationAnalyzeApi.md#GetTrace) | **Get** /api/application-monitoring/analyze/traces/{id} | Get trace detail
[**GetTraceGroups**](ApplicationAnalyzeApi.md#GetTraceGroups) | **Post** /api/application-monitoring/analyze/trace-groups | Get grouped trace metrics
[**GetTraces**](ApplicationAnalyzeApi.md#GetTraces) | **Post** /api/application-monitoring/analyze/traces | Get all traces



## GetCallGroup

> CallGroupsResult GetCallGroup(ctx).FillTimeSeries(fillTimeSeries).GetCallGroups(getCallGroups).Execute()

Get grouped call metrics



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
    getCallGroups := *openapiclient.NewGetCallGroups(*openapiclient.NewGroup("GroupbyTag_example", "GroupbyTagEntity_example"), []openapiclient.MetricConfig{*openapiclient.NewMetricConfig("Aggregation_example", "Metric_example")}) // GetCallGroups |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAnalyzeApi.GetCallGroup(context.Background()).FillTimeSeries(fillTimeSeries).GetCallGroups(getCallGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAnalyzeApi.GetCallGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCallGroup`: CallGroupsResult
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAnalyzeApi.GetCallGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCallGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fillTimeSeries** | **bool** |  | 
 **getCallGroups** | [**GetCallGroups**](GetCallGroups.md) |  | 

### Return type

[**CallGroupsResult**](CallGroupsResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCorrelatedTraces

> []BackendTraceReference GetCorrelatedTraces(ctx).CorrelationId(correlationId).Execute()

Resolve backend trace IDs using correlation IDs from website and mobile app monitoring beacons.

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
    correlationId := "correlationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAnalyzeApi.GetCorrelatedTraces(context.Background()).CorrelationId(correlationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAnalyzeApi.GetCorrelatedTraces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCorrelatedTraces`: []BackendTraceReference
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAnalyzeApi.GetCorrelatedTraces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCorrelatedTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **correlationId** | **string** |  | 

### Return type

[**[]BackendTraceReference**](BackendTraceReference.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrace

> FullTrace GetTrace(ctx, id).Execute()

Get trace detail

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
    resp, r, err := apiClient.ApplicationAnalyzeApi.GetTrace(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAnalyzeApi.GetTrace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrace`: FullTrace
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAnalyzeApi.GetTrace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FullTrace**](FullTrace.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceGroups

> TraceGroupsResult GetTraceGroups(ctx).FillTimeSeries(fillTimeSeries).GetTraceGroups(getTraceGroups).Execute()

Get grouped trace metrics

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
    getTraceGroups := *openapiclient.NewGetTraceGroups(*openapiclient.NewGroup("GroupbyTag_example", "GroupbyTagEntity_example"), []openapiclient.MetricConfig{*openapiclient.NewMetricConfig("Aggregation_example", "Metric_example")}) // GetTraceGroups |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAnalyzeApi.GetTraceGroups(context.Background()).FillTimeSeries(fillTimeSeries).GetTraceGroups(getTraceGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAnalyzeApi.GetTraceGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTraceGroups`: TraceGroupsResult
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAnalyzeApi.GetTraceGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fillTimeSeries** | **bool** |  | 
 **getTraceGroups** | [**GetTraceGroups**](GetTraceGroups.md) |  | 

### Return type

[**TraceGroupsResult**](TraceGroupsResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraces

> TraceResult GetTraces(ctx).GetTraces(getTraces).Execute()

Get all traces



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
    getTraces := *openapiclient.NewGetTraces() // GetTraces |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAnalyzeApi.GetTraces(context.Background()).GetTraces(getTraces).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAnalyzeApi.GetTraces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTraces`: TraceResult
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAnalyzeApi.GetTraces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getTraces** | [**GetTraces**](GetTraces.md) |  | 

### Return type

[**TraceResult**](TraceResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

