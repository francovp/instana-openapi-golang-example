# \MaintenanceConfigurationApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMaintenanceConfig**](MaintenanceConfigurationApi.md#DeleteMaintenanceConfig) | **Delete** /api/settings/maintenance/{id} | Delete maintenance configuration
[**GetMaintenanceConfig**](MaintenanceConfigurationApi.md#GetMaintenanceConfig) | **Get** /api/settings/maintenance/{id} | Maintenance configuration
[**GetMaintenanceConfigs**](MaintenanceConfigurationApi.md#GetMaintenanceConfigs) | **Get** /api/settings/maintenance | All maintenance configurations
[**PutMaintenanceConfig**](MaintenanceConfigurationApi.md#PutMaintenanceConfig) | **Put** /api/settings/maintenance/{id} | Create or update maintenance configuration



## DeleteMaintenanceConfig

> DeleteMaintenanceConfig(ctx, id).Execute()

Delete maintenance configuration

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
    resp, r, err := apiClient.MaintenanceConfigurationApi.DeleteMaintenanceConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceConfigurationApi.DeleteMaintenanceConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMaintenanceConfigRequest struct via the builder pattern


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


## GetMaintenanceConfig

> MaintenanceConfigWithLastUpdated GetMaintenanceConfig(ctx, id).Execute()

Maintenance configuration

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
    resp, r, err := apiClient.MaintenanceConfigurationApi.GetMaintenanceConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceConfigurationApi.GetMaintenanceConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMaintenanceConfig`: MaintenanceConfigWithLastUpdated
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceConfigurationApi.GetMaintenanceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaintenanceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MaintenanceConfigWithLastUpdated**](MaintenanceConfigWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaintenanceConfigs

> []ValidatedMaintenanceConfigWithStatus GetMaintenanceConfigs(ctx).Execute()

All maintenance configurations

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
    resp, r, err := apiClient.MaintenanceConfigurationApi.GetMaintenanceConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceConfigurationApi.GetMaintenanceConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMaintenanceConfigs`: []ValidatedMaintenanceConfigWithStatus
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceConfigurationApi.GetMaintenanceConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaintenanceConfigsRequest struct via the builder pattern


### Return type

[**[]ValidatedMaintenanceConfigWithStatus**](ValidatedMaintenanceConfigWithStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMaintenanceConfig

> PutMaintenanceConfig(ctx, id).MaintenanceConfig(maintenanceConfig).Execute()

Create or update maintenance configuration

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
    maintenanceConfig := *openapiclient.NewMaintenanceConfig("Id_example", "Name_example", "Query_example") // MaintenanceConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MaintenanceConfigurationApi.PutMaintenanceConfig(context.Background(), id).MaintenanceConfig(maintenanceConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceConfigurationApi.PutMaintenanceConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutMaintenanceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maintenanceConfig** | [**MaintenanceConfig**](MaintenanceConfig.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

