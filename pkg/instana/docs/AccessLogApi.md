# \AccessLogApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessLogs**](AccessLogApi.md#GetAccessLogs) | **Get** /api/settings/accesslog | Access log



## GetAccessLogs

> AccessLogResponse GetAccessLogs(ctx).Offset(offset).Query(query).PageSize(pageSize).Execute()

Access log

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
    offset := int32(56) // int32 |  (optional)
    query := "query_example" // string |  (optional)
    pageSize := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessLogApi.GetAccessLogs(context.Background()).Offset(offset).Query(query).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessLogApi.GetAccessLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessLogs`: AccessLogResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessLogApi.GetAccessLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **query** | **string** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**AccessLogResponse**](AccessLogResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

