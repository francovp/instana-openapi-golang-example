# \ApplicationResourcesApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationEndpoints**](ApplicationResourcesApi.md#GetApplicationEndpoints) | **Get** /api/application-monitoring/applications/services/endpoints | Get endpoints
[**GetApplicationServices**](ApplicationResourcesApi.md#GetApplicationServices) | **Get** /api/application-monitoring/applications/services | Get applications/services
[**GetApplications**](ApplicationResourcesApi.md#GetApplications) | **Get** /api/application-monitoring/applications | Get applications
[**GetServices**](ApplicationResourcesApi.md#GetServices) | **Get** /api/application-monitoring/services | Get services



## GetApplicationEndpoints

> EndpointResult GetApplicationEndpoints(ctx).NameFilter(nameFilter).Types(types).Technologies(technologies).WindowSize(windowSize).To(to).Page(page).PageSize(pageSize).ApplicationBoundaryScope(applicationBoundaryScope).Execute()

Get endpoints

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
    nameFilter := "nameFilter_example" // string | Name of service (optional)
    types := []string{"Inner_example"} // []string | Type of Endpoint (optional)
    technologies := []string{"Inner_example"} // []string | List of technologies (optional)
    windowSize := int64(789) // int64 | Size of time window in milliseconds (optional)
    to := int64(789) // int64 | Timestamp since Unix Epoch in milliseconds of the end of the time window (optional)
    page := int32(56) // int32 | Page number from results (optional)
    pageSize := int32(56) // int32 | Number of items per page (optional)
    applicationBoundaryScope := "applicationBoundaryScope_example" // string | Filter for application scope, i.e: INBOUND or ALL (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationResourcesApi.GetApplicationEndpoints(context.Background()).NameFilter(nameFilter).Types(types).Technologies(technologies).WindowSize(windowSize).To(to).Page(page).PageSize(pageSize).ApplicationBoundaryScope(applicationBoundaryScope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcesApi.GetApplicationEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationEndpoints`: EndpointResult
    fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcesApi.GetApplicationEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameFilter** | **string** | Name of service | 
 **types** | **[]string** | Type of Endpoint | 
 **technologies** | **[]string** | List of technologies | 
 **windowSize** | **int64** | Size of time window in milliseconds | 
 **to** | **int64** | Timestamp since Unix Epoch in milliseconds of the end of the time window | 
 **page** | **int32** | Page number from results | 
 **pageSize** | **int32** | Number of items per page | 
 **applicationBoundaryScope** | **string** | Filter for application scope, i.e: INBOUND or ALL | 

### Return type

[**EndpointResult**](EndpointResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationServices

> ServiceResult GetApplicationServices(ctx).NameFilter(nameFilter).WindowSize(windowSize).To(to).Page(page).PageSize(pageSize).ApplicationBoundaryScope(applicationBoundaryScope).Execute()

Get applications/services

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
    nameFilter := "nameFilter_example" // string | Name of application/service (optional)
    windowSize := int64(789) // int64 | Size of time window in milliseconds (optional)
    to := int64(789) // int64 | Timestamp since Unix Epoch in milliseconds of the end of the time window (optional)
    page := int32(56) // int32 | Page number from results (optional)
    pageSize := int32(56) // int32 | Number of items per page (optional)
    applicationBoundaryScope := "applicationBoundaryScope_example" // string | Filter for application scope, i.e: INBOUND or ALL (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationResourcesApi.GetApplicationServices(context.Background()).NameFilter(nameFilter).WindowSize(windowSize).To(to).Page(page).PageSize(pageSize).ApplicationBoundaryScope(applicationBoundaryScope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcesApi.GetApplicationServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationServices`: ServiceResult
    fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcesApi.GetApplicationServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameFilter** | **string** | Name of application/service | 
 **windowSize** | **int64** | Size of time window in milliseconds | 
 **to** | **int64** | Timestamp since Unix Epoch in milliseconds of the end of the time window | 
 **page** | **int32** | Page number from results | 
 **pageSize** | **int32** | Number of items per page | 
 **applicationBoundaryScope** | **string** | Filter for application scope, i.e: INBOUND or ALL | 

### Return type

[**ServiceResult**](ServiceResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> ApplicationResult GetApplications(ctx).NameFilter(nameFilter).WindowSize(windowSize).To(to).Page(page).PageSize(pageSize).ApplicationBoundaryScope(applicationBoundaryScope).Execute()

Get applications

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
    nameFilter := "nameFilter_example" // string | Name of application (optional)
    windowSize := int64(789) // int64 | Size of time window in milliseconds (optional)
    to := int64(789) // int64 | Timestamp since Unix Epoch in milliseconds of the end of the time window (optional)
    page := int32(56) // int32 | Page number from results (optional)
    pageSize := int32(56) // int32 | Number of items per page (optional)
    applicationBoundaryScope := "applicationBoundaryScope_example" // string | Filter for application scope, i.e: INBOUND or ALL (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationResourcesApi.GetApplications(context.Background()).NameFilter(nameFilter).WindowSize(windowSize).To(to).Page(page).PageSize(pageSize).ApplicationBoundaryScope(applicationBoundaryScope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcesApi.GetApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplications`: ApplicationResult
    fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcesApi.GetApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameFilter** | **string** | Name of application | 
 **windowSize** | **int64** | Size of time window in milliseconds | 
 **to** | **int64** | Timestamp since Unix Epoch in milliseconds of the end of the time window | 
 **page** | **int32** | Page number from results | 
 **pageSize** | **int32** | Number of items per page | 
 **applicationBoundaryScope** | **string** | Filter for application scope, i.e: INBOUND or ALL | 

### Return type

[**ApplicationResult**](ApplicationResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServices

> ServiceResult GetServices(ctx).NameFilter(nameFilter).WindowSize(windowSize).To(to).Page(page).PageSize(pageSize).Execute()

Get services

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
    nameFilter := "nameFilter_example" // string | Name of service (optional)
    windowSize := int64(789) // int64 | Size of time window in milliseconds (optional)
    to := int64(789) // int64 | Timestamp since Unix Epoch in milliseconds of the end of the time window (optional)
    page := int32(56) // int32 | Page number from results (optional)
    pageSize := int32(56) // int32 | Number of items per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationResourcesApi.GetServices(context.Background()).NameFilter(nameFilter).WindowSize(windowSize).To(to).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcesApi.GetServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServices`: ServiceResult
    fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcesApi.GetServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameFilter** | **string** | Name of service | 
 **windowSize** | **int64** | Size of time window in milliseconds | 
 **to** | **int64** | Timestamp since Unix Epoch in milliseconds of the end of the time window | 
 **page** | **int32** | Page number from results | 
 **pageSize** | **int32** | Number of items per page | 

### Return type

[**ServiceResult**](ServiceResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

