# \AuditLogApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditLogs**](AuditLogApi.md#GetAuditLogs) | **Get** /api/settings/auditlog | Audit log



## GetAuditLogs

> AuditLogUiResponse GetAuditLogs(ctx).Offset(offset).Query(query).PageSize(pageSize).Execute()

Audit log

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
    resp, r, err := apiClient.AuditLogApi.GetAuditLogs(context.Background()).Offset(offset).Query(query).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogApi.GetAuditLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuditLogs`: AuditLogUiResponse
    fmt.Fprintf(os.Stdout, "Response from `AuditLogApi.GetAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **query** | **string** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**AuditLogUiResponse**](AuditLogUiResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

