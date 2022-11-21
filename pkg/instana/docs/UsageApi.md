# \UsageApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllUsage**](UsageApi.md#GetAllUsage) | **Get** /api/instana/usage/api | API usage by customer
[**GetHostsPerDay**](UsageApi.md#GetHostsPerDay) | **Get** /api/instana/usage/hosts/{day}/{month}/{year} | Host count day / month / year
[**GetHostsPerMonth**](UsageApi.md#GetHostsPerMonth) | **Get** /api/instana/usage/hosts/{month}/{year} | Host count month / year
[**GetUsagePerDay**](UsageApi.md#GetUsagePerDay) | **Get** /api/instana/usage/api/{day}/{month}/{year} | API usage day / month / year
[**GetUsagePerMonth**](UsageApi.md#GetUsagePerMonth) | **Get** /api/instana/usage/api/{month}/{year} | API usage month / year



## GetAllUsage

> []UsageResult GetAllUsage(ctx).Execute()

API usage by customer

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
    resp, r, err := apiClient.UsageApi.GetAllUsage(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetAllUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllUsage`: []UsageResult
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetAllUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllUsageRequest struct via the builder pattern


### Return type

[**[]UsageResult**](UsageResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostsPerDay

> []UsageResult GetHostsPerDay(ctx, day, month, year).Execute()

Host count day / month / year

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
    day := int32(56) // int32 | 
    month := int32(56) // int32 | 
    year := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetHostsPerDay(context.Background(), day, month, year).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetHostsPerDay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostsPerDay`: []UsageResult
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetHostsPerDay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**day** | **int32** |  | 
**month** | **int32** |  | 
**year** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostsPerDayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]UsageResult**](UsageResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostsPerMonth

> []UsageResult GetHostsPerMonth(ctx, month, year).Execute()

Host count month / year

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
    month := int32(56) // int32 | 
    year := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetHostsPerMonth(context.Background(), month, year).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetHostsPerMonth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostsPerMonth`: []UsageResult
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetHostsPerMonth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**month** | **int32** |  | 
**year** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostsPerMonthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]UsageResult**](UsageResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsagePerDay

> []UsageResult GetUsagePerDay(ctx, day, month, year).Execute()

API usage day / month / year

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
    day := int32(56) // int32 | 
    month := int32(56) // int32 | 
    year := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetUsagePerDay(context.Background(), day, month, year).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetUsagePerDay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsagePerDay`: []UsageResult
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetUsagePerDay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**day** | **int32** |  | 
**month** | **int32** |  | 
**year** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsagePerDayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]UsageResult**](UsageResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsagePerMonth

> []UsageResult GetUsagePerMonth(ctx, month, year).Execute()

API usage month / year

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
    month := int32(56) // int32 | 
    year := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetUsagePerMonth(context.Background(), month, year).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetUsagePerMonth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsagePerMonth`: []UsageResult
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetUsagePerMonth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**month** | **int32** |  | 
**year** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsagePerMonthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]UsageResult**](UsageResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

