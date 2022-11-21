# \MobileAppConfigurationApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMobileAppConfig**](MobileAppConfigurationApi.md#DeleteMobileAppConfig) | **Delete** /api/mobile-app-monitoring/config/{mobileAppId} | Remove mobile app
[**GetMobileAppConfig**](MobileAppConfigurationApi.md#GetMobileAppConfig) | **Get** /api/mobile-app-monitoring/config | Get configured mobile apps
[**GetMobileAppGeoLocationConfiguration**](MobileAppConfigurationApi.md#GetMobileAppGeoLocationConfiguration) | **Get** /api/mobile-app-monitoring/config/{mobileAppId}/geo-location | Get geo location configuration for mobile app
[**GetMobileAppIpMaskingConfiguration**](MobileAppConfigurationApi.md#GetMobileAppIpMaskingConfiguration) | **Get** /api/mobile-app-monitoring/config/{mobileAppId}/ip-masking | Get IP masking configuration for mobile app
[**PostMobileAppConfig**](MobileAppConfigurationApi.md#PostMobileAppConfig) | **Post** /api/mobile-app-monitoring/config | Configure new mobile app
[**RenameMobileAppConfig**](MobileAppConfigurationApi.md#RenameMobileAppConfig) | **Put** /api/mobile-app-monitoring/config/{mobileAppId} | Rename mobile app
[**UpdateMobileAppGeoLocationConfiguration**](MobileAppConfigurationApi.md#UpdateMobileAppGeoLocationConfiguration) | **Put** /api/mobile-app-monitoring/config/{mobileAppId}/geo-location | Update geo location configuration for mobile app
[**UpdateMobileAppIpMaskingConfiguration**](MobileAppConfigurationApi.md#UpdateMobileAppIpMaskingConfiguration) | **Put** /api/mobile-app-monitoring/config/{mobileAppId}/ip-masking | Update IP masking configuration for mobile app



## DeleteMobileAppConfig

> DeleteMobileAppConfig(ctx, mobileAppId).Execute()

Remove mobile app

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
    mobileAppId := "mobileAppId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileAppConfigurationApi.DeleteMobileAppConfig(context.Background(), mobileAppId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileAppConfigurationApi.DeleteMobileAppConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMobileAppConfigRequest struct via the builder pattern


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


## GetMobileAppConfig

> []MobileApp GetMobileAppConfig(ctx).Execute()

Get configured mobile apps

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
    resp, r, err := apiClient.MobileAppConfigurationApi.GetMobileAppConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileAppConfigurationApi.GetMobileAppConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMobileAppConfig`: []MobileApp
    fmt.Fprintf(os.Stdout, "Response from `MobileAppConfigurationApi.GetMobileAppConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMobileAppConfigRequest struct via the builder pattern


### Return type

[**[]MobileApp**](MobileApp.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMobileAppGeoLocationConfiguration

> GeoLocationConfiguration GetMobileAppGeoLocationConfiguration(ctx, mobileAppId).Execute()

Get geo location configuration for mobile app

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
    mobileAppId := "mobileAppId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileAppConfigurationApi.GetMobileAppGeoLocationConfiguration(context.Background(), mobileAppId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileAppConfigurationApi.GetMobileAppGeoLocationConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMobileAppGeoLocationConfiguration`: GeoLocationConfiguration
    fmt.Fprintf(os.Stdout, "Response from `MobileAppConfigurationApi.GetMobileAppGeoLocationConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMobileAppGeoLocationConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GeoLocationConfiguration**](GeoLocationConfiguration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMobileAppIpMaskingConfiguration

> IpMaskingConfiguration GetMobileAppIpMaskingConfiguration(ctx, mobileAppId).Execute()

Get IP masking configuration for mobile app

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
    mobileAppId := "mobileAppId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileAppConfigurationApi.GetMobileAppIpMaskingConfiguration(context.Background(), mobileAppId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileAppConfigurationApi.GetMobileAppIpMaskingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMobileAppIpMaskingConfiguration`: IpMaskingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `MobileAppConfigurationApi.GetMobileAppIpMaskingConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMobileAppIpMaskingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IpMaskingConfiguration**](IpMaskingConfiguration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMobileAppConfig

> MobileApp PostMobileAppConfig(ctx).Name(name).Execute()

Configure new mobile app

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
    name := "name_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileAppConfigurationApi.PostMobileAppConfig(context.Background()).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileAppConfigurationApi.PostMobileAppConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMobileAppConfig`: MobileApp
    fmt.Fprintf(os.Stdout, "Response from `MobileAppConfigurationApi.PostMobileAppConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMobileAppConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 

### Return type

[**MobileApp**](MobileApp.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameMobileAppConfig

> MobileApp RenameMobileAppConfig(ctx, mobileAppId).Name(name).Execute()

Rename mobile app

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
    mobileAppId := "mobileAppId_example" // string | 
    name := "name_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileAppConfigurationApi.RenameMobileAppConfig(context.Background(), mobileAppId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileAppConfigurationApi.RenameMobileAppConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenameMobileAppConfig`: MobileApp
    fmt.Fprintf(os.Stdout, "Response from `MobileAppConfigurationApi.RenameMobileAppConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameMobileAppConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 

### Return type

[**MobileApp**](MobileApp.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMobileAppGeoLocationConfiguration

> GeoLocationConfiguration UpdateMobileAppGeoLocationConfiguration(ctx, mobileAppId).GeoLocationConfiguration(geoLocationConfiguration).Execute()

Update geo location configuration for mobile app

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
    mobileAppId := "mobileAppId_example" // string | 
    geoLocationConfiguration := *openapiclient.NewGeoLocationConfiguration("GeoDetailRemoval_example") // GeoLocationConfiguration |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileAppConfigurationApi.UpdateMobileAppGeoLocationConfiguration(context.Background(), mobileAppId).GeoLocationConfiguration(geoLocationConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileAppConfigurationApi.UpdateMobileAppGeoLocationConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMobileAppGeoLocationConfiguration`: GeoLocationConfiguration
    fmt.Fprintf(os.Stdout, "Response from `MobileAppConfigurationApi.UpdateMobileAppGeoLocationConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMobileAppGeoLocationConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **geoLocationConfiguration** | [**GeoLocationConfiguration**](GeoLocationConfiguration.md) |  | 

### Return type

[**GeoLocationConfiguration**](GeoLocationConfiguration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMobileAppIpMaskingConfiguration

> IpMaskingConfiguration UpdateMobileAppIpMaskingConfiguration(ctx, mobileAppId).IpMaskingConfiguration(ipMaskingConfiguration).Execute()

Update IP masking configuration for mobile app

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
    mobileAppId := "mobileAppId_example" // string | 
    ipMaskingConfiguration := *openapiclient.NewIpMaskingConfiguration("IpMasking_example") // IpMaskingConfiguration |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileAppConfigurationApi.UpdateMobileAppIpMaskingConfiguration(context.Background(), mobileAppId).IpMaskingConfiguration(ipMaskingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileAppConfigurationApi.UpdateMobileAppIpMaskingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMobileAppIpMaskingConfiguration`: IpMaskingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `MobileAppConfigurationApi.UpdateMobileAppIpMaskingConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMobileAppIpMaskingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipMaskingConfiguration** | [**IpMaskingConfiguration**](IpMaskingConfiguration.md) |  | 

### Return type

[**IpMaskingConfiguration**](IpMaskingConfiguration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

