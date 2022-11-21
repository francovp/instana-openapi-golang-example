# \SLISettingsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSliConfig**](SLISettingsApi.md#CreateSliConfig) | **Post** /api/settings/sli | Create SLI Config
[**CreateSliConfigV2**](SLISettingsApi.md#CreateSliConfigV2) | **Post** /api/settings/v2/sli | Create SLI Config
[**DeleteSliConfig**](SLISettingsApi.md#DeleteSliConfig) | **Delete** /api/settings/sli/{id} | Delete SLI Config
[**DeleteSliConfigV2**](SLISettingsApi.md#DeleteSliConfigV2) | **Delete** /api/settings/v2/sli/{id} | Delete SLI Config
[**GetAllSliConfigs**](SLISettingsApi.md#GetAllSliConfigs) | **Get** /api/settings/sli | Get All SLI Configs
[**GetAllSliConfigsV2**](SLISettingsApi.md#GetAllSliConfigsV2) | **Get** /api/settings/v2/sli | Get All SLI Configs
[**GetSliConfig**](SLISettingsApi.md#GetSliConfig) | **Get** /api/settings/sli/{id} | Get SLI Config
[**GetSliConfigV2**](SLISettingsApi.md#GetSliConfigV2) | **Get** /api/settings/v2/sli/{id} | Get SLI Config
[**UpdateSliConfig**](SLISettingsApi.md#UpdateSliConfig) | **Put** /api/settings/sli/{id} | Update SLI Config



## CreateSliConfig

> SliConfigurationWithLastUpdated CreateSliConfig(ctx).SliConfiguration(sliConfiguration).Execute()

Create SLI Config



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
    sliConfiguration := *openapiclient.NewSliConfiguration("Id_example", *openapiclient.NewSliEntity("SliType_example"), "SliName_example") // SliConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SLISettingsApi.CreateSliConfig(context.Background()).SliConfiguration(sliConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLISettingsApi.CreateSliConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSliConfig`: SliConfigurationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `SLISettingsApi.CreateSliConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSliConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sliConfiguration** | [**SliConfiguration**](SliConfiguration.md) |  | 

### Return type

[**SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSliConfigV2

> SliConfigurationWithLastUpdated CreateSliConfigV2(ctx).SliConfiguration(sliConfiguration).Execute()

Create SLI Config



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
    sliConfiguration := *openapiclient.NewSliConfiguration("Id_example", *openapiclient.NewSliEntity("SliType_example"), "SliName_example") // SliConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SLISettingsApi.CreateSliConfigV2(context.Background()).SliConfiguration(sliConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLISettingsApi.CreateSliConfigV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSliConfigV2`: SliConfigurationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `SLISettingsApi.CreateSliConfigV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSliConfigV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sliConfiguration** | [**SliConfiguration**](SliConfiguration.md) |  | 

### Return type

[**SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSliConfig

> DeleteSliConfig(ctx, id).Execute()

Delete SLI Config

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
    resp, r, err := apiClient.SLISettingsApi.DeleteSliConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLISettingsApi.DeleteSliConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSliConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSliConfigV2

> DeleteSliConfigV2(ctx, id).Execute()

Delete SLI Config

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
    resp, r, err := apiClient.SLISettingsApi.DeleteSliConfigV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLISettingsApi.DeleteSliConfigV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSliConfigV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSliConfigs

> []SliConfigurationWithLastUpdated GetAllSliConfigs(ctx).Execute()

Get All SLI Configs

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
    resp, r, err := apiClient.SLISettingsApi.GetAllSliConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLISettingsApi.GetAllSliConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllSliConfigs`: []SliConfigurationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `SLISettingsApi.GetAllSliConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSliConfigsRequest struct via the builder pattern


### Return type

[**[]SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSliConfigsV2

> []SliConfigurationWithLastUpdated GetAllSliConfigsV2(ctx).Execute()

Get All SLI Configs

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
    resp, r, err := apiClient.SLISettingsApi.GetAllSliConfigsV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLISettingsApi.GetAllSliConfigsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllSliConfigsV2`: []SliConfigurationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `SLISettingsApi.GetAllSliConfigsV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSliConfigsV2Request struct via the builder pattern


### Return type

[**[]SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSliConfig

> SliConfigurationWithLastUpdated GetSliConfig(ctx, id).Execute()

Get SLI Config

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
    resp, r, err := apiClient.SLISettingsApi.GetSliConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLISettingsApi.GetSliConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSliConfig`: SliConfigurationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `SLISettingsApi.GetSliConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSliConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSliConfigV2

> SliConfigurationWithLastUpdated GetSliConfigV2(ctx, id).Execute()

Get SLI Config

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
    resp, r, err := apiClient.SLISettingsApi.GetSliConfigV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLISettingsApi.GetSliConfigV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSliConfigV2`: SliConfigurationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `SLISettingsApi.GetSliConfigV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSliConfigV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSliConfig

> SliConfigurationWithLastUpdated UpdateSliConfig(ctx, id).SliConfiguration(sliConfiguration).Execute()

Update SLI Config

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
    sliConfiguration := *openapiclient.NewSliConfiguration("Id_example", *openapiclient.NewSliEntity("SliType_example"), "SliName_example") // SliConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SLISettingsApi.UpdateSliConfig(context.Background(), id).SliConfiguration(sliConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SLISettingsApi.UpdateSliConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSliConfig`: SliConfigurationWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `SLISettingsApi.UpdateSliConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSliConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sliConfiguration** | [**SliConfiguration**](SliConfiguration.md) |  | 

### Return type

[**SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

