# \CustomDashboardsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCustomDashboard**](CustomDashboardsApi.md#AddCustomDashboard) | **Post** /api/custom-dashboard | Add custom dashboard
[**DeleteCustomDashboard**](CustomDashboardsApi.md#DeleteCustomDashboard) | **Delete** /api/custom-dashboard/{customDashboardId} | Remove custom dashboard
[**GetCustomDashboard**](CustomDashboardsApi.md#GetCustomDashboard) | **Get** /api/custom-dashboard/{customDashboardId} | Get custom dashboard
[**GetCustomDashboards**](CustomDashboardsApi.md#GetCustomDashboards) | **Get** /api/custom-dashboard | Get accessible custom dashboards
[**GetShareableApiTokens**](CustomDashboardsApi.md#GetShareableApiTokens) | **Get** /api/custom-dashboard/shareable-api-tokens | Get all API tokens.
[**GetShareableUsers**](CustomDashboardsApi.md#GetShareableUsers) | **Get** /api/custom-dashboard/shareable-users | Get all users (without invitations).
[**UpdateCustomDashboard**](CustomDashboardsApi.md#UpdateCustomDashboard) | **Put** /api/custom-dashboard/{customDashboardId} | Update custom dashboard



## AddCustomDashboard

> CustomDashboard AddCustomDashboard(ctx).CustomDashboard(customDashboard).Execute()

Add custom dashboard

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
    customDashboard := *openapiclient.NewCustomDashboard([]openapiclient.AccessRule{*openapiclient.NewAccessRule("AccessType_example", "RelationType_example")}, "Id_example", "Title_example", []openapiclient.Widget{*openapiclient.NewWidget(interface{}(123), "Id_example", "Type_example")}) // CustomDashboard |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDashboardsApi.AddCustomDashboard(context.Background()).CustomDashboard(customDashboard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDashboardsApi.AddCustomDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCustomDashboard`: CustomDashboard
    fmt.Fprintf(os.Stdout, "Response from `CustomDashboardsApi.AddCustomDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCustomDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customDashboard** | [**CustomDashboard**](CustomDashboard.md) |  | 

### Return type

[**CustomDashboard**](CustomDashboard.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomDashboard

> DeleteCustomDashboard(ctx, customDashboardId).Execute()

Remove custom dashboard

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
    customDashboardId := "customDashboardId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDashboardsApi.DeleteCustomDashboard(context.Background(), customDashboardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDashboardsApi.DeleteCustomDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customDashboardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomDashboard

> CustomDashboard GetCustomDashboard(ctx, customDashboardId).Execute()

Get custom dashboard

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
    customDashboardId := "customDashboardId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDashboardsApi.GetCustomDashboard(context.Background(), customDashboardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDashboardsApi.GetCustomDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomDashboard`: CustomDashboard
    fmt.Fprintf(os.Stdout, "Response from `CustomDashboardsApi.GetCustomDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customDashboardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomDashboard**](CustomDashboard.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomDashboards

> []CustomDashboardPreview GetCustomDashboards(ctx).Query(query).PageSize(pageSize).Page(page).Execute()

Get accessible custom dashboards

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
    query := "query_example" // string |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDashboardsApi.GetCustomDashboards(context.Background()).Query(query).PageSize(pageSize).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDashboardsApi.GetCustomDashboards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomDashboards`: []CustomDashboardPreview
    fmt.Fprintf(os.Stdout, "Response from `CustomDashboardsApi.GetCustomDashboards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomDashboardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **pageSize** | **int32** |  | 
 **page** | **int32** |  | 

### Return type

[**[]CustomDashboardPreview**](CustomDashboardPreview.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShareableApiTokens

> []ApiToken GetShareableApiTokens(ctx).Execute()

Get all API tokens.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDashboardsApi.GetShareableApiTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDashboardsApi.GetShareableApiTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShareableApiTokens`: []ApiToken
    fmt.Fprintf(os.Stdout, "Response from `CustomDashboardsApi.GetShareableApiTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetShareableApiTokensRequest struct via the builder pattern


### Return type

[**[]ApiToken**](ApiToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShareableUsers

> []UserResult GetShareableUsers(ctx).Execute()

Get all users (without invitations).

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDashboardsApi.GetShareableUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDashboardsApi.GetShareableUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShareableUsers`: []UserResult
    fmt.Fprintf(os.Stdout, "Response from `CustomDashboardsApi.GetShareableUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetShareableUsersRequest struct via the builder pattern


### Return type

[**[]UserResult**](UserResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomDashboard

> CustomDashboard UpdateCustomDashboard(ctx, customDashboardId).CustomDashboard(customDashboard).Execute()

Update custom dashboard

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
    customDashboardId := "customDashboardId_example" // string | 
    customDashboard := *openapiclient.NewCustomDashboard([]openapiclient.AccessRule{*openapiclient.NewAccessRule("AccessType_example", "RelationType_example")}, "Id_example", "Title_example", []openapiclient.Widget{*openapiclient.NewWidget(interface{}(123), "Id_example", "Type_example")}) // CustomDashboard |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDashboardsApi.UpdateCustomDashboard(context.Background(), customDashboardId).CustomDashboard(customDashboard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDashboardsApi.UpdateCustomDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomDashboard`: CustomDashboard
    fmt.Fprintf(os.Stdout, "Response from `CustomDashboardsApi.UpdateCustomDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customDashboardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customDashboard** | [**CustomDashboard**](CustomDashboard.md) |  | 

### Return type

[**CustomDashboard**](CustomDashboard.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

