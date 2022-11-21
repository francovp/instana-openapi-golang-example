# \SLIReportApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSli**](SLIReportApi.md#GetSli) | **Get** /api/sli/report/{sliId} | Generate SLI report



## GetSli

> []SliReport GetSli(ctx, sliId).Slo(slo).From(from).To(to).Execute()

Generate SLI report

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
    sliId := "sliId_example" // string | 
    slo := float64(1.2) // float64 |  (optional)
    from := int64(789) // int64 |  (optional)
    to := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SLIReportApi.GetSli(context.Background(), sliId).Slo(slo).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLIReportApi.GetSli``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSli`: []SliReport
    fmt.Fprintf(os.Stdout, "Response from `SLIReportApi.GetSli`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sliId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSliRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **slo** | **float64** |  | 
 **from** | **int64** |  | 
 **to** | **int64** |  | 

### Return type

[**[]SliReport**](SliReport.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

