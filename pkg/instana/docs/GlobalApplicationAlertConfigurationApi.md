# \GlobalApplicationAlertConfigurationApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGlobalApplicationAlertConfig**](GlobalApplicationAlertConfigurationApi.md#CreateGlobalApplicationAlertConfig) | **Post** /api/events/settings/global-alert-configs/applications | Create Global Application Alert Config
[**DeleteGlobalApplicationAlertConfig**](GlobalApplicationAlertConfigurationApi.md#DeleteGlobalApplicationAlertConfig) | **Delete** /api/events/settings/global-alert-configs/applications/{id} | Delete Global Application Alert Config
[**DisableGlobalApplicationAlertConfig**](GlobalApplicationAlertConfigurationApi.md#DisableGlobalApplicationAlertConfig) | **Put** /api/events/settings/global-alert-configs/applications/{id}/disable | Disable Global Application Alert Config
[**EnableGlobalApplicationAlertConfig**](GlobalApplicationAlertConfigurationApi.md#EnableGlobalApplicationAlertConfig) | **Put** /api/events/settings/global-alert-configs/applications/{id}/enable | Enable Global Application Alert Config
[**FindActiveGlobalApplicationAlertConfigs**](GlobalApplicationAlertConfigurationApi.md#FindActiveGlobalApplicationAlertConfigs) | **Get** /api/events/settings/global-alert-configs/applications | All Global Application Alert Configs
[**FindGlobalApplicationAlertConfig**](GlobalApplicationAlertConfigurationApi.md#FindGlobalApplicationAlertConfig) | **Get** /api/events/settings/global-alert-configs/applications/{id} | Get Global Application Alert Config
[**FindGlobalApplicationAlertConfigVersions**](GlobalApplicationAlertConfigurationApi.md#FindGlobalApplicationAlertConfigVersions) | **Get** /api/events/settings/global-alert-configs/applications/{id}/versions | Get versions of Global Application Alert Config
[**RestoreGlobalApplicationAlertConfig**](GlobalApplicationAlertConfigurationApi.md#RestoreGlobalApplicationAlertConfig) | **Put** /api/events/settings/global-alert-configs/applications/{id}/restore/{created} | Restore Global Application Alert Config
[**UpdateGlobalApplicationAlertConfig**](GlobalApplicationAlertConfigurationApi.md#UpdateGlobalApplicationAlertConfig) | **Post** /api/events/settings/global-alert-configs/applications/{id} | Update Global Application Alert Config



## CreateGlobalApplicationAlertConfig

> GlobalApplicationAlertConfigWithMetadata CreateGlobalApplicationAlertConfig(ctx).GlobalApplicationsAlertConfig(globalApplicationsAlertConfig).Execute()

Create Global Application Alert Config

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
    globalApplicationsAlertConfig := *openapiclient.NewGlobalApplicationsAlertConfig([]string{"AlertChannelIds_example"}, map[string]ApplicationNode{"key": *openapiclient.NewApplicationNode("ApplicationId_example", map[string]ServiceNode{"key": *openapiclient.NewServiceNode(map[string]EndpointNode{"key": *openapiclient.NewEndpointNode("EndpointId_example")}, "ServiceId_example")})}, "BoundaryScope_example", []openapiclient.CustomPayloadField{*openapiclient.NewCustomPayloadField("Key_example", "Type_example")}, "Description_example", "EvaluationType_example", int32(123), "Name_example", *openapiclient.NewApplicationAlertRule("AlertType_example", "MetricName_example"), *openapiclient.NewTagFilterExpressionElement("Type_example"), *openapiclient.NewThreshold("Operator_example", "Type_example"), *openapiclient.NewApplicationTimeThreshold("Type_example")) // GlobalApplicationsAlertConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalApplicationAlertConfigurationApi.CreateGlobalApplicationAlertConfig(context.Background()).GlobalApplicationsAlertConfig(globalApplicationsAlertConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalApplicationAlertConfigurationApi.CreateGlobalApplicationAlertConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGlobalApplicationAlertConfig`: GlobalApplicationAlertConfigWithMetadata
    fmt.Fprintf(os.Stdout, "Response from `GlobalApplicationAlertConfigurationApi.CreateGlobalApplicationAlertConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGlobalApplicationAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalApplicationsAlertConfig** | [**GlobalApplicationsAlertConfig**](GlobalApplicationsAlertConfig.md) |  | 

### Return type

[**GlobalApplicationAlertConfigWithMetadata**](GlobalApplicationAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGlobalApplicationAlertConfig

> DeleteGlobalApplicationAlertConfig(ctx, id).Execute()

Delete Global Application Alert Config

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
    resp, r, err := apiClient.GlobalApplicationAlertConfigurationApi.DeleteGlobalApplicationAlertConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalApplicationAlertConfigurationApi.DeleteGlobalApplicationAlertConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteGlobalApplicationAlertConfigRequest struct via the builder pattern


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


## DisableGlobalApplicationAlertConfig

> DisableGlobalApplicationAlertConfig(ctx, id).Body(body).Execute()

Disable Global Application Alert Config

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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalApplicationAlertConfigurationApi.DisableGlobalApplicationAlertConfig(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalApplicationAlertConfigurationApi.DisableGlobalApplicationAlertConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDisableGlobalApplicationAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

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


## EnableGlobalApplicationAlertConfig

> EnableGlobalApplicationAlertConfig(ctx, id).Body(body).Execute()

Enable Global Application Alert Config

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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalApplicationAlertConfigurationApi.EnableGlobalApplicationAlertConfig(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalApplicationAlertConfigurationApi.EnableGlobalApplicationAlertConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnableGlobalApplicationAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

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


## FindActiveGlobalApplicationAlertConfigs

> []GlobalApplicationAlertConfigWithMetadata FindActiveGlobalApplicationAlertConfigs(ctx).ApplicationId(applicationId).AlertIds(alertIds).Execute()

All Global Application Alert Configs



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
    applicationId := "applicationId_example" // string |  (optional)
    alertIds := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalApplicationAlertConfigurationApi.FindActiveGlobalApplicationAlertConfigs(context.Background()).ApplicationId(applicationId).AlertIds(alertIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalApplicationAlertConfigurationApi.FindActiveGlobalApplicationAlertConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindActiveGlobalApplicationAlertConfigs`: []GlobalApplicationAlertConfigWithMetadata
    fmt.Fprintf(os.Stdout, "Response from `GlobalApplicationAlertConfigurationApi.FindActiveGlobalApplicationAlertConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindActiveGlobalApplicationAlertConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | **string** |  | 
 **alertIds** | **[]string** |  | 

### Return type

[**[]GlobalApplicationAlertConfigWithMetadata**](GlobalApplicationAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindGlobalApplicationAlertConfig

> GlobalApplicationAlertConfigWithMetadata FindGlobalApplicationAlertConfig(ctx, id).ValidOn(validOn).Execute()

Get Global Application Alert Config



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
    validOn := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalApplicationAlertConfigurationApi.FindGlobalApplicationAlertConfig(context.Background(), id).ValidOn(validOn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalApplicationAlertConfigurationApi.FindGlobalApplicationAlertConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindGlobalApplicationAlertConfig`: GlobalApplicationAlertConfigWithMetadata
    fmt.Fprintf(os.Stdout, "Response from `GlobalApplicationAlertConfigurationApi.FindGlobalApplicationAlertConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindGlobalApplicationAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validOn** | **int64** |  | 

### Return type

[**GlobalApplicationAlertConfigWithMetadata**](GlobalApplicationAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindGlobalApplicationAlertConfigVersions

> []ConfigVersion FindGlobalApplicationAlertConfigVersions(ctx, id).Execute()

Get versions of Global Application Alert Config



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
    resp, r, err := apiClient.GlobalApplicationAlertConfigurationApi.FindGlobalApplicationAlertConfigVersions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalApplicationAlertConfigurationApi.FindGlobalApplicationAlertConfigVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindGlobalApplicationAlertConfigVersions`: []ConfigVersion
    fmt.Fprintf(os.Stdout, "Response from `GlobalApplicationAlertConfigurationApi.FindGlobalApplicationAlertConfigVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindGlobalApplicationAlertConfigVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ConfigVersion**](ConfigVersion.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreGlobalApplicationAlertConfig

> RestoreGlobalApplicationAlertConfig(ctx, id, created).Body(body).Execute()

Restore Global Application Alert Config

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
    created := int64(789) // int64 | 
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalApplicationAlertConfigurationApi.RestoreGlobalApplicationAlertConfig(context.Background(), id, created).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalApplicationAlertConfigurationApi.RestoreGlobalApplicationAlertConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**created** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreGlobalApplicationAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

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


## UpdateGlobalApplicationAlertConfig

> GlobalApplicationAlertConfigWithMetadata UpdateGlobalApplicationAlertConfig(ctx, id).GlobalApplicationsAlertConfig(globalApplicationsAlertConfig).Execute()

Update Global Application Alert Config

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
    globalApplicationsAlertConfig := *openapiclient.NewGlobalApplicationsAlertConfig([]string{"AlertChannelIds_example"}, map[string]ApplicationNode{"key": *openapiclient.NewApplicationNode("ApplicationId_example", map[string]ServiceNode{"key": *openapiclient.NewServiceNode(map[string]EndpointNode{"key": *openapiclient.NewEndpointNode("EndpointId_example")}, "ServiceId_example")})}, "BoundaryScope_example", []openapiclient.CustomPayloadField{*openapiclient.NewCustomPayloadField("Key_example", "Type_example")}, "Description_example", "EvaluationType_example", int32(123), "Name_example", *openapiclient.NewApplicationAlertRule("AlertType_example", "MetricName_example"), *openapiclient.NewTagFilterExpressionElement("Type_example"), *openapiclient.NewThreshold("Operator_example", "Type_example"), *openapiclient.NewApplicationTimeThreshold("Type_example")) // GlobalApplicationsAlertConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalApplicationAlertConfigurationApi.UpdateGlobalApplicationAlertConfig(context.Background(), id).GlobalApplicationsAlertConfig(globalApplicationsAlertConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalApplicationAlertConfigurationApi.UpdateGlobalApplicationAlertConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGlobalApplicationAlertConfig`: GlobalApplicationAlertConfigWithMetadata
    fmt.Fprintf(os.Stdout, "Response from `GlobalApplicationAlertConfigurationApi.UpdateGlobalApplicationAlertConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlobalApplicationAlertConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **globalApplicationsAlertConfig** | [**GlobalApplicationsAlertConfig**](GlobalApplicationsAlertConfig.md) |  | 

### Return type

[**GlobalApplicationAlertConfigWithMetadata**](GlobalApplicationAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

