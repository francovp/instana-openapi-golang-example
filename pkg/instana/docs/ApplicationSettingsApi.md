# \ApplicationSettingsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddApplicationConfig**](ApplicationSettingsApi.md#AddApplicationConfig) | **Post** /api/application-monitoring/settings/application | Add application configuration
[**AddManualServiceConfig**](ApplicationSettingsApi.md#AddManualServiceConfig) | **Post** /api/application-monitoring/settings/manual-service | Add manual service configuration
[**AddServiceConfig**](ApplicationSettingsApi.md#AddServiceConfig) | **Post** /api/application-monitoring/settings/service | Add service configuration
[**CreateEndpointConfig**](ApplicationSettingsApi.md#CreateEndpointConfig) | **Post** /api/application-monitoring/settings/http-endpoint | Create endpoint configuration
[**DeleteApplicationConfig**](ApplicationSettingsApi.md#DeleteApplicationConfig) | **Delete** /api/application-monitoring/settings/application/{id} | Delete application configuration
[**DeleteEndpointConfig**](ApplicationSettingsApi.md#DeleteEndpointConfig) | **Delete** /api/application-monitoring/settings/http-endpoint/{id} | Delete endpoint configuration
[**DeleteManualServiceConfig**](ApplicationSettingsApi.md#DeleteManualServiceConfig) | **Delete** /api/application-monitoring/settings/manual-service/{id} | Delete manual service configuration
[**DeleteServiceConfig**](ApplicationSettingsApi.md#DeleteServiceConfig) | **Delete** /api/application-monitoring/settings/service/{id} | Delete service configuration
[**GetAllManualServiceConfigs**](ApplicationSettingsApi.md#GetAllManualServiceConfigs) | **Get** /api/application-monitoring/settings/manual-service | All manual service configurations
[**GetApplicationConfig**](ApplicationSettingsApi.md#GetApplicationConfig) | **Get** /api/application-monitoring/settings/application/{id} | Application configuration
[**GetApplicationConfigs**](ApplicationSettingsApi.md#GetApplicationConfigs) | **Get** /api/application-monitoring/settings/application | All Application configurations
[**GetEndpointConfig**](ApplicationSettingsApi.md#GetEndpointConfig) | **Get** /api/application-monitoring/settings/http-endpoint/{id} | Endpoint configuration
[**GetEndpointConfigs**](ApplicationSettingsApi.md#GetEndpointConfigs) | **Get** /api/application-monitoring/settings/http-endpoint | All Endpoint configurations
[**GetServiceConfig**](ApplicationSettingsApi.md#GetServiceConfig) | **Get** /api/application-monitoring/settings/service/{id} | Service configuration
[**GetServiceConfigs**](ApplicationSettingsApi.md#GetServiceConfigs) | **Get** /api/application-monitoring/settings/service | All service configurations
[**OrderServiceConfig**](ApplicationSettingsApi.md#OrderServiceConfig) | **Put** /api/application-monitoring/settings/service/order | Order of service configuration
[**PutApplicationConfig**](ApplicationSettingsApi.md#PutApplicationConfig) | **Put** /api/application-monitoring/settings/application/{id} | Update application configuration
[**PutServiceConfig**](ApplicationSettingsApi.md#PutServiceConfig) | **Put** /api/application-monitoring/settings/service/{id} | Update service configuration
[**ReplaceAll**](ApplicationSettingsApi.md#ReplaceAll) | **Put** /api/application-monitoring/settings/service | Replace all service configurations
[**ReplaceAllManualServiceConfigs**](ApplicationSettingsApi.md#ReplaceAllManualServiceConfigs) | **Put** /api/application-monitoring/settings/manual-service | Replace all manual service configurations
[**UpdateEndpointConfig**](ApplicationSettingsApi.md#UpdateEndpointConfig) | **Put** /api/application-monitoring/settings/http-endpoint/{id} | Update endpoint configuration
[**UpdateManualServiceConfig**](ApplicationSettingsApi.md#UpdateManualServiceConfig) | **Put** /api/application-monitoring/settings/manual-service/{id} | Update manual service configuration



## AddApplicationConfig

> ApplicationConfig AddApplicationConfig(ctx).NewApplicationConfig(newApplicationConfig).Execute()

Add application configuration



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
    newApplicationConfig := *openapiclient.NewNewApplicationConfig([]openapiclient.AccessRule{*openapiclient.NewAccessRule("AccessType_example", "RelationType_example")}, "BoundaryScope_example", "Label_example", "Scope_example") // NewApplicationConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationSettingsApi.AddApplicationConfig(context.Background()).NewApplicationConfig(newApplicationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.AddApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddApplicationConfig`: ApplicationConfig
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSettingsApi.AddApplicationConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newApplicationConfig** | [**NewApplicationConfig**](NewApplicationConfig.md) |  | 

### Return type

[**ApplicationConfig**](ApplicationConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddManualServiceConfig

> ManualServiceConfig AddManualServiceConfig(ctx).NewManualServiceConfig(newManualServiceConfig).Execute()

Add manual service configuration



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
    newManualServiceConfig := *openapiclient.NewNewManualServiceConfig(*openapiclient.NewTagFilterExpressionElement("Type_example")) // NewManualServiceConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationSettingsApi.AddManualServiceConfig(context.Background()).NewManualServiceConfig(newManualServiceConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.AddManualServiceConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddManualServiceConfig`: ManualServiceConfig
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSettingsApi.AddManualServiceConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddManualServiceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newManualServiceConfig** | [**NewManualServiceConfig**](NewManualServiceConfig.md) |  | 

### Return type

[**ManualServiceConfig**](ManualServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddServiceConfig

> ServiceConfig AddServiceConfig(ctx).ServiceConfig(serviceConfig).Execute()

Add service configuration

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
    serviceConfig := *openapiclient.NewServiceConfig(false, "Id_example", "Label_example", []openapiclient.ServiceMatchingRule{*openapiclient.NewServiceMatchingRule("Key_example", "Value_example")}, "Name_example") // ServiceConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationSettingsApi.AddServiceConfig(context.Background()).ServiceConfig(serviceConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.AddServiceConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddServiceConfig`: ServiceConfig
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSettingsApi.AddServiceConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddServiceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceConfig** | [**ServiceConfig**](ServiceConfig.md) |  | 

### Return type

[**ServiceConfig**](ServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEndpointConfig

> HttpEndpointConfig CreateEndpointConfig(ctx).HttpEndpointConfig(httpEndpointConfig).Execute()

Create endpoint configuration

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
    httpEndpointConfig := *openapiclient.NewHttpEndpointConfig([]openapiclient.HttpEndpointRule{*openapiclient.NewHttpEndpointRule([]openapiclient.HttpPathSegmentMatchingRule{*openapiclient.NewHttpPathSegmentMatchingRule("Type_example")})}, "ServiceId_example") // HttpEndpointConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationSettingsApi.CreateEndpointConfig(context.Background()).HttpEndpointConfig(httpEndpointConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.CreateEndpointConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEndpointConfig`: HttpEndpointConfig
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSettingsApi.CreateEndpointConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **httpEndpointConfig** | [**HttpEndpointConfig**](HttpEndpointConfig.md) |  | 

### Return type

[**HttpEndpointConfig**](HttpEndpointConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationConfig

> DeleteApplicationConfig(ctx, id).Execute()

Delete application configuration

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
    resp, r, err := apiClient.ApplicationSettingsApi.DeleteApplicationConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.DeleteApplicationConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApplicationConfigRequest struct via the builder pattern


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


## DeleteEndpointConfig

> DeleteEndpointConfig(ctx, id).Execute()

Delete endpoint configuration

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
    resp, r, err := apiClient.ApplicationSettingsApi.DeleteEndpointConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.DeleteEndpointConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteEndpointConfigRequest struct via the builder pattern


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


## DeleteManualServiceConfig

> DeleteManualServiceConfig(ctx, id).Execute()

Delete manual service configuration



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
    resp, r, err := apiClient.ApplicationSettingsApi.DeleteManualServiceConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.DeleteManualServiceConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteManualServiceConfigRequest struct via the builder pattern


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


## DeleteServiceConfig

> DeleteServiceConfig(ctx, id).Execute()

Delete service configuration

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
    resp, r, err := apiClient.ApplicationSettingsApi.DeleteServiceConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.DeleteServiceConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteServiceConfigRequest struct via the builder pattern


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


## GetAllManualServiceConfigs

> []ManualServiceConfig GetAllManualServiceConfigs(ctx).Execute()

All manual service configurations



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
    resp, r, err := apiClient.ApplicationSettingsApi.GetAllManualServiceConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.GetAllManualServiceConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllManualServiceConfigs`: []ManualServiceConfig
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSettingsApi.GetAllManualServiceConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllManualServiceConfigsRequest struct via the builder pattern


### Return type

[**[]ManualServiceConfig**](ManualServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationConfig

> ApplicationConfig GetApplicationConfig(ctx, id).Execute()

Application configuration

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
    resp, r, err := apiClient.ApplicationSettingsApi.GetApplicationConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.GetApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationConfig`: ApplicationConfig
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSettingsApi.GetApplicationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationConfig**](ApplicationConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationConfigs

> []ApplicationConfig GetApplicationConfigs(ctx).Execute()

All Application configurations

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
    resp, r, err := apiClient.ApplicationSettingsApi.GetApplicationConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.GetApplicationConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationConfigs`: []ApplicationConfig
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSettingsApi.GetApplicationConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationConfigsRequest struct via the builder pattern


### Return type

[**[]ApplicationConfig**](ApplicationConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointConfig

> HttpEndpointConfig GetEndpointConfig(ctx, id).Execute()

Endpoint configuration

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
    resp, r, err := apiClient.ApplicationSettingsApi.GetEndpointConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.GetEndpointConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEndpointConfig`: HttpEndpointConfig
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSettingsApi.GetEndpointConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HttpEndpointConfig**](HttpEndpointConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointConfigs

> []HttpEndpointConfig GetEndpointConfigs(ctx).Execute()

All Endpoint configurations

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
    resp, r, err := apiClient.ApplicationSettingsApi.GetEndpointConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.GetEndpointConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEndpointConfigs`: []HttpEndpointConfig
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSettingsApi.GetEndpointConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointConfigsRequest struct via the builder pattern


### Return type

[**[]HttpEndpointConfig**](HttpEndpointConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceConfig

> ServiceConfig GetServiceConfig(ctx, id).Execute()

Service configuration

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
    resp, r, err := apiClient.ApplicationSettingsApi.GetServiceConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.GetServiceConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceConfig`: ServiceConfig
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSettingsApi.GetServiceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceConfig**](ServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceConfigs

> []ServiceConfig GetServiceConfigs(ctx).Execute()

All service configurations

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
    resp, r, err := apiClient.ApplicationSettingsApi.GetServiceConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.GetServiceConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceConfigs`: []ServiceConfig
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSettingsApi.GetServiceConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceConfigsRequest struct via the builder pattern


### Return type

[**[]ServiceConfig**](ServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderServiceConfig

> OrderServiceConfig(ctx).RequestBody(requestBody).Execute()

Order of service configuration

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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationSettingsApi.OrderServiceConfig(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.OrderServiceConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderServiceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

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


## PutApplicationConfig

> ApplicationConfig PutApplicationConfig(ctx, id).ApplicationConfig(applicationConfig).Execute()

Update application configuration



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
    applicationConfig := *openapiclient.NewApplicationConfig([]openapiclient.AccessRule{*openapiclient.NewAccessRule("AccessType_example", "RelationType_example")}, "BoundaryScope_example", "Id_example", "Label_example", "Scope_example") // ApplicationConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationSettingsApi.PutApplicationConfig(context.Background(), id).ApplicationConfig(applicationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.PutApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutApplicationConfig`: ApplicationConfig
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSettingsApi.PutApplicationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationConfig** | [**ApplicationConfig**](ApplicationConfig.md) |  | 

### Return type

[**ApplicationConfig**](ApplicationConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceConfig

> ServiceConfig PutServiceConfig(ctx, id).ServiceConfig(serviceConfig).Execute()

Update service configuration

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
    serviceConfig := *openapiclient.NewServiceConfig(false, "Id_example", "Label_example", []openapiclient.ServiceMatchingRule{*openapiclient.NewServiceMatchingRule("Key_example", "Value_example")}, "Name_example") // ServiceConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationSettingsApi.PutServiceConfig(context.Background(), id).ServiceConfig(serviceConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.PutServiceConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutServiceConfig`: ServiceConfig
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSettingsApi.PutServiceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceConfig** | [**ServiceConfig**](ServiceConfig.md) |  | 

### Return type

[**ServiceConfig**](ServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceAll

> []ServiceConfig ReplaceAll(ctx).ServiceConfig(serviceConfig).Execute()

Replace all service configurations

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
    serviceConfig := []openapiclient.ServiceConfig{*openapiclient.NewServiceConfig(false, "Id_example", "Label_example", []openapiclient.ServiceMatchingRule{*openapiclient.NewServiceMatchingRule("Key_example", "Value_example")}, "Name_example")} // []ServiceConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationSettingsApi.ReplaceAll(context.Background()).ServiceConfig(serviceConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.ReplaceAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceAll`: []ServiceConfig
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSettingsApi.ReplaceAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceConfig** | [**[]ServiceConfig**](ServiceConfig.md) |  | 

### Return type

[**[]ServiceConfig**](ServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceAllManualServiceConfigs

> []ManualServiceConfig ReplaceAllManualServiceConfigs(ctx).NewManualServiceConfig(newManualServiceConfig).Execute()

Replace all manual service configurations



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
    newManualServiceConfig := []openapiclient.NewManualServiceConfig{*openapiclient.NewNewManualServiceConfig(*openapiclient.NewTagFilterExpressionElement("Type_example"))} // []NewManualServiceConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationSettingsApi.ReplaceAllManualServiceConfigs(context.Background()).NewManualServiceConfig(newManualServiceConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.ReplaceAllManualServiceConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceAllManualServiceConfigs`: []ManualServiceConfig
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSettingsApi.ReplaceAllManualServiceConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceAllManualServiceConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newManualServiceConfig** | [**[]NewManualServiceConfig**](NewManualServiceConfig.md) |  | 

### Return type

[**[]ManualServiceConfig**](ManualServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpointConfig

> HttpEndpointConfig UpdateEndpointConfig(ctx, id).HttpEndpointConfig(httpEndpointConfig).Execute()

Update endpoint configuration

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
    httpEndpointConfig := *openapiclient.NewHttpEndpointConfig([]openapiclient.HttpEndpointRule{*openapiclient.NewHttpEndpointRule([]openapiclient.HttpPathSegmentMatchingRule{*openapiclient.NewHttpPathSegmentMatchingRule("Type_example")})}, "ServiceId_example") // HttpEndpointConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationSettingsApi.UpdateEndpointConfig(context.Background(), id).HttpEndpointConfig(httpEndpointConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.UpdateEndpointConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEndpointConfig`: HttpEndpointConfig
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSettingsApi.UpdateEndpointConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **httpEndpointConfig** | [**HttpEndpointConfig**](HttpEndpointConfig.md) |  | 

### Return type

[**HttpEndpointConfig**](HttpEndpointConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManualServiceConfig

> ManualServiceConfig UpdateManualServiceConfig(ctx, id).ManualServiceConfig(manualServiceConfig).Execute()

Update manual service configuration



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
    manualServiceConfig := *openapiclient.NewManualServiceConfig("Id_example", *openapiclient.NewTagFilterExpressionElement("Type_example")) // ManualServiceConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationSettingsApi.UpdateManualServiceConfig(context.Background(), id).ManualServiceConfig(manualServiceConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSettingsApi.UpdateManualServiceConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateManualServiceConfig`: ManualServiceConfig
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSettingsApi.UpdateManualServiceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateManualServiceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **manualServiceConfig** | [**ManualServiceConfig**](ManualServiceConfig.md) |  | 

### Return type

[**ManualServiceConfig**](ManualServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

