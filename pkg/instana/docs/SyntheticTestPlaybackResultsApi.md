# \SyntheticTestPlaybackResultsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSyntheticResult**](SyntheticTestPlaybackResultsApi.md#GetSyntheticResult) | **Post** /api/synthetics/results | Experimental - Get Synthetic test playback results
[**GetSyntheticResultDetailData**](SyntheticTestPlaybackResultsApi.md#GetSyntheticResultDetailData) | **Get** /api/synthetics/results/{testid}/{testresultid}/detail | Experimental - Get Synthetic test playback result detail data
[**GetSyntheticResultList**](SyntheticTestPlaybackResultsApi.md#GetSyntheticResultList) | **Post** /api/synthetics/results/list | Experimental - Get a list of Synthetic test playback results
[**GetSyntheticResultMetadata**](SyntheticTestPlaybackResultsApi.md#GetSyntheticResultMetadata) | **Get** /api/synthetics/results/{testid}/{testresultid} | Experimental - Get Synthetic test playback detail result description(metadata)
[**GetTestSummaryList**](SyntheticTestPlaybackResultsApi.md#GetTestSummaryList) | **Post** /api/synthetics/results/testsummarylist | Experimental - Get a list of Synthetic tests with success rate and average response time data



## GetSyntheticResult

> TestResult GetSyntheticResult(ctx).GetTestResult(getTestResult).Execute()

Experimental - Get Synthetic test playback results

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
    getTestResult := *openapiclient.NewGetTestResult([]openapiclient.SyntheticMetricConfiguration{*openapiclient.NewSyntheticMetricConfiguration("Aggregation_example", "Metric_example")}) // GetTestResult |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticTestPlaybackResultsApi.GetSyntheticResult(context.Background()).GetTestResult(getTestResult).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticTestPlaybackResultsApi.GetSyntheticResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSyntheticResult`: TestResult
    fmt.Fprintf(os.Stdout, "Response from `SyntheticTestPlaybackResultsApi.GetSyntheticResult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSyntheticResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getTestResult** | [**GetTestResult**](GetTestResult.md) |  | 

### Return type

[**TestResult**](TestResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyntheticResultDetailData

> TestResultDetailData GetSyntheticResultDetailData(ctx, testid, testresultid).Type_(type_).Name(name).Execute()

Experimental - Get Synthetic test playback result detail data

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
    testid := "testid_example" // string | 
    testresultid := "testresultid_example" // string | 
    type_ := "type__example" // string | 
    name := "name_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticTestPlaybackResultsApi.GetSyntheticResultDetailData(context.Background(), testid, testresultid).Type_(type_).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticTestPlaybackResultsApi.GetSyntheticResultDetailData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSyntheticResultDetailData`: TestResultDetailData
    fmt.Fprintf(os.Stdout, "Response from `SyntheticTestPlaybackResultsApi.GetSyntheticResultDetailData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testid** | **string** |  | 
**testresultid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyntheticResultDetailDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **string** |  | 
 **name** | **string** |  | 

### Return type

[**TestResultDetailData**](TestResultDetailData.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyntheticResultList

> TestResultListResult GetSyntheticResultList(ctx).GetTestResultList(getTestResultList).Execute()

Experimental - Get a list of Synthetic test playback results

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
    getTestResultList := *openapiclient.NewGetTestResultList([]string{"SyntheticMetrics_example"}) // GetTestResultList |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticTestPlaybackResultsApi.GetSyntheticResultList(context.Background()).GetTestResultList(getTestResultList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticTestPlaybackResultsApi.GetSyntheticResultList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSyntheticResultList`: TestResultListResult
    fmt.Fprintf(os.Stdout, "Response from `SyntheticTestPlaybackResultsApi.GetSyntheticResultList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSyntheticResultListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getTestResultList** | [**GetTestResultList**](GetTestResultList.md) |  | 

### Return type

[**TestResultListResult**](TestResultListResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyntheticResultMetadata

> TestResultMetadata GetSyntheticResultMetadata(ctx, testid, testresultid).Execute()

Experimental - Get Synthetic test playback detail result description(metadata)

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
    testid := "testid_example" // string | 
    testresultid := "testresultid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticTestPlaybackResultsApi.GetSyntheticResultMetadata(context.Background(), testid, testresultid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticTestPlaybackResultsApi.GetSyntheticResultMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSyntheticResultMetadata`: TestResultMetadata
    fmt.Fprintf(os.Stdout, "Response from `SyntheticTestPlaybackResultsApi.GetSyntheticResultMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testid** | **string** |  | 
**testresultid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyntheticResultMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TestResultMetadata**](TestResultMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestSummaryList

> TestResultListResult GetTestSummaryList(ctx).GetTestResult(getTestResult).Execute()

Experimental - Get a list of Synthetic tests with success rate and average response time data

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
    getTestResult := *openapiclient.NewGetTestResult([]openapiclient.SyntheticMetricConfiguration{*openapiclient.NewSyntheticMetricConfiguration("Aggregation_example", "Metric_example")}) // GetTestResult |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticTestPlaybackResultsApi.GetTestSummaryList(context.Background()).GetTestResult(getTestResult).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticTestPlaybackResultsApi.GetTestSummaryList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTestSummaryList`: TestResultListResult
    fmt.Fprintf(os.Stdout, "Response from `SyntheticTestPlaybackResultsApi.GetTestSummaryList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTestSummaryListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getTestResult** | [**GetTestResult**](GetTestResult.md) |  | 

### Return type

[**TestResultListResult**](TestResultListResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

