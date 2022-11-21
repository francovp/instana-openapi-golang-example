# \WebsiteConfigurationApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebsite**](WebsiteConfigurationApi.md#CreateWebsite) | **Post** /api/website-monitoring/config | Configure new website
[**DeleteWebsite**](WebsiteConfigurationApi.md#DeleteWebsite) | **Delete** /api/website-monitoring/config/{websiteId} | Remove website
[**GetMobileAppGeoMappingRules**](WebsiteConfigurationApi.md#GetMobileAppGeoMappingRules) | **Get** /api/mobile-app-monitoring/config/{mobileAppId}/geo-mapping-rules | Get custom geo mapping rules for mobile app
[**GetWebsite**](WebsiteConfigurationApi.md#GetWebsite) | **Get** /api/website-monitoring/config/{websiteId} | Get configured website
[**GetWebsiteGeoLocationConfiguration**](WebsiteConfigurationApi.md#GetWebsiteGeoLocationConfiguration) | **Get** /api/website-monitoring/config/{websiteId}/geo-location | Get geo location configuration for website
[**GetWebsiteGeoMappingRules**](WebsiteConfigurationApi.md#GetWebsiteGeoMappingRules) | **Get** /api/website-monitoring/config/{websiteId}/geo-mapping-rules | Get custom geo mapping rules for website
[**GetWebsiteIpMaskingConfiguration**](WebsiteConfigurationApi.md#GetWebsiteIpMaskingConfiguration) | **Get** /api/website-monitoring/config/{websiteId}/ip-masking | Get IP masking configuration for website
[**GetWebsites**](WebsiteConfigurationApi.md#GetWebsites) | **Get** /api/website-monitoring/config | Get configured websites
[**RenameWebsite**](WebsiteConfigurationApi.md#RenameWebsite) | **Put** /api/website-monitoring/config/{websiteId} | Rename website
[**SetMobileAppGeoMappingRules**](WebsiteConfigurationApi.md#SetMobileAppGeoMappingRules) | **Put** /api/mobile-app-monitoring/config/{mobileAppId}/geo-mapping-rules | Set custom geo mapping rules for mobile app
[**SetWebsiteGeoMappingRules**](WebsiteConfigurationApi.md#SetWebsiteGeoMappingRules) | **Put** /api/website-monitoring/config/{websiteId}/geo-mapping-rules | Set custom geo mapping rules for website
[**UpdateWebsiteGeoLocationConfiguration**](WebsiteConfigurationApi.md#UpdateWebsiteGeoLocationConfiguration) | **Put** /api/website-monitoring/config/{websiteId}/geo-location | Update geo location configuration for website
[**UpdateWebsiteIpMaskingConfiguration**](WebsiteConfigurationApi.md#UpdateWebsiteIpMaskingConfiguration) | **Put** /api/website-monitoring/config/{websiteId}/ip-masking | Update IP masking configuration for website



## CreateWebsite

> Website CreateWebsite(ctx).Name(name).Execute()

Configure new website

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
    resp, r, err := apiClient.WebsiteConfigurationApi.CreateWebsite(context.Background()).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteConfigurationApi.CreateWebsite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebsite`: Website
    fmt.Fprintf(os.Stdout, "Response from `WebsiteConfigurationApi.CreateWebsite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebsiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 

### Return type

[**Website**](Website.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebsite

> DeleteWebsite(ctx, websiteId).Execute()

Remove website

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
    websiteId := "websiteId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebsiteConfigurationApi.DeleteWebsite(context.Background(), websiteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteConfigurationApi.DeleteWebsite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebsiteRequest struct via the builder pattern


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


## GetMobileAppGeoMappingRules

> GetMobileAppGeoMappingRules(ctx, mobileAppId).Execute()

Get custom geo mapping rules for mobile app

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
    resp, r, err := apiClient.WebsiteConfigurationApi.GetMobileAppGeoMappingRules(context.Background(), mobileAppId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteConfigurationApi.GetMobileAppGeoMappingRules``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetMobileAppGeoMappingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebsite

> Website GetWebsite(ctx, websiteId).Execute()

Get configured website

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
    websiteId := "websiteId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebsiteConfigurationApi.GetWebsite(context.Background(), websiteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteConfigurationApi.GetWebsite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebsite`: Website
    fmt.Fprintf(os.Stdout, "Response from `WebsiteConfigurationApi.GetWebsite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebsiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Website**](Website.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebsiteGeoLocationConfiguration

> GeoLocationConfiguration GetWebsiteGeoLocationConfiguration(ctx, websiteId).Execute()

Get geo location configuration for website

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
    websiteId := "websiteId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebsiteConfigurationApi.GetWebsiteGeoLocationConfiguration(context.Background(), websiteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteConfigurationApi.GetWebsiteGeoLocationConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebsiteGeoLocationConfiguration`: GeoLocationConfiguration
    fmt.Fprintf(os.Stdout, "Response from `WebsiteConfigurationApi.GetWebsiteGeoLocationConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebsiteGeoLocationConfigurationRequest struct via the builder pattern


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


## GetWebsiteGeoMappingRules

> GetWebsiteGeoMappingRules(ctx, websiteId).Execute()

Get custom geo mapping rules for website

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
    websiteId := "websiteId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebsiteConfigurationApi.GetWebsiteGeoMappingRules(context.Background(), websiteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteConfigurationApi.GetWebsiteGeoMappingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebsiteGeoMappingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebsiteIpMaskingConfiguration

> IpMaskingConfiguration GetWebsiteIpMaskingConfiguration(ctx, websiteId).Execute()

Get IP masking configuration for website

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
    websiteId := "websiteId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebsiteConfigurationApi.GetWebsiteIpMaskingConfiguration(context.Background(), websiteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteConfigurationApi.GetWebsiteIpMaskingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebsiteIpMaskingConfiguration`: IpMaskingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `WebsiteConfigurationApi.GetWebsiteIpMaskingConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebsiteIpMaskingConfigurationRequest struct via the builder pattern


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


## GetWebsites

> []Website GetWebsites(ctx).Execute()

Get configured websites

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
    resp, r, err := apiClient.WebsiteConfigurationApi.GetWebsites(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteConfigurationApi.GetWebsites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebsites`: []Website
    fmt.Fprintf(os.Stdout, "Response from `WebsiteConfigurationApi.GetWebsites`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebsitesRequest struct via the builder pattern


### Return type

[**[]Website**](Website.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameWebsite

> Website RenameWebsite(ctx, websiteId).Name(name).Execute()

Rename website

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
    websiteId := "websiteId_example" // string | 
    name := "name_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebsiteConfigurationApi.RenameWebsite(context.Background(), websiteId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteConfigurationApi.RenameWebsite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenameWebsite`: Website
    fmt.Fprintf(os.Stdout, "Response from `WebsiteConfigurationApi.RenameWebsite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameWebsiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 

### Return type

[**Website**](Website.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMobileAppGeoMappingRules

> SetMobileAppGeoMappingRules(ctx, mobileAppId).Body(body).Execute()

Set custom geo mapping rules for mobile app

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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebsiteConfigurationApi.SetMobileAppGeoMappingRules(context.Background(), mobileAppId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteConfigurationApi.SetMobileAppGeoMappingRules``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetMobileAppGeoMappingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: text/csv
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetWebsiteGeoMappingRules

> SetWebsiteGeoMappingRules(ctx, websiteId).Body(body).Execute()

Set custom geo mapping rules for website

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
    websiteId := "websiteId_example" // string | 
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebsiteConfigurationApi.SetWebsiteGeoMappingRules(context.Background(), websiteId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteConfigurationApi.SetWebsiteGeoMappingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetWebsiteGeoMappingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: text/csv
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebsiteGeoLocationConfiguration

> GeoLocationConfiguration UpdateWebsiteGeoLocationConfiguration(ctx, websiteId).GeoLocationConfiguration(geoLocationConfiguration).Execute()

Update geo location configuration for website

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
    websiteId := "websiteId_example" // string | 
    geoLocationConfiguration := *openapiclient.NewGeoLocationConfiguration("GeoDetailRemoval_example") // GeoLocationConfiguration |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebsiteConfigurationApi.UpdateWebsiteGeoLocationConfiguration(context.Background(), websiteId).GeoLocationConfiguration(geoLocationConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteConfigurationApi.UpdateWebsiteGeoLocationConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebsiteGeoLocationConfiguration`: GeoLocationConfiguration
    fmt.Fprintf(os.Stdout, "Response from `WebsiteConfigurationApi.UpdateWebsiteGeoLocationConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebsiteGeoLocationConfigurationRequest struct via the builder pattern


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


## UpdateWebsiteIpMaskingConfiguration

> IpMaskingConfiguration UpdateWebsiteIpMaskingConfiguration(ctx, websiteId).IpMaskingConfiguration(ipMaskingConfiguration).Execute()

Update IP masking configuration for website

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
    websiteId := "websiteId_example" // string | 
    ipMaskingConfiguration := *openapiclient.NewIpMaskingConfiguration("IpMasking_example") // IpMaskingConfiguration |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebsiteConfigurationApi.UpdateWebsiteIpMaskingConfiguration(context.Background(), websiteId).IpMaskingConfiguration(ipMaskingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebsiteConfigurationApi.UpdateWebsiteIpMaskingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebsiteIpMaskingConfiguration`: IpMaskingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `WebsiteConfigurationApi.UpdateWebsiteIpMaskingConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebsiteIpMaskingConfigurationRequest struct via the builder pattern


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

