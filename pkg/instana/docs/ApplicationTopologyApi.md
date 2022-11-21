# \ApplicationTopologyApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServicesMap**](ApplicationTopologyApi.md#GetServicesMap) | **Get** /api/application-monitoring/topology/services | Gets the service topology



## GetServicesMap

> ServiceMap GetServicesMap(ctx).WindowSize(windowSize).To(to).ApplicationId(applicationId).ApplicationBoundaryScope(applicationBoundaryScope).Execute()

Gets the service topology

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
    windowSize := int64(789) // int64 | Size of time window in milliseconds (optional)
    to := int64(789) // int64 | Timestamp since Unix Epoch in milliseconds of the end of the time window (optional)
    applicationId := "applicationId_example" // string | Filter by application ID (optional)
    applicationBoundaryScope := "applicationBoundaryScope_example" // string | Filter by application scope, i.e., INBOUND or ALL. The default value is INBOUND. Applies only if application ID filter is specified. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationTopologyApi.GetServicesMap(context.Background()).WindowSize(windowSize).To(to).ApplicationId(applicationId).ApplicationBoundaryScope(applicationBoundaryScope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationTopologyApi.GetServicesMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServicesMap`: ServiceMap
    fmt.Fprintf(os.Stdout, "Response from `ApplicationTopologyApi.GetServicesMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **windowSize** | **int64** | Size of time window in milliseconds | 
 **to** | **int64** | Timestamp since Unix Epoch in milliseconds of the end of the time window | 
 **applicationId** | **string** | Filter by application ID | 
 **applicationBoundaryScope** | **string** | Filter by application scope, i.e., INBOUND or ALL. The default value is INBOUND. Applies only if application ID filter is specified. | 

### Return type

[**ServiceMap**](ServiceMap.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

