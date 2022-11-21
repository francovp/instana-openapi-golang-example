# \MobileAppCatalogApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMobileAppMetricCatalog**](MobileAppCatalogApi.md#GetMobileAppMetricCatalog) | **Get** /api/mobile-app-monitoring/catalog/metrics | Metric catalog
[**GetMobileAppTagCatalog**](MobileAppCatalogApi.md#GetMobileAppTagCatalog) | **Get** /api/mobile-app-monitoring/catalog | Get mobile app tag catalog



## GetMobileAppMetricCatalog

> []MobileAppMonitoringMetricDescription GetMobileAppMetricCatalog(ctx).Execute()

Metric catalog

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
    resp, r, err := apiClient.MobileAppCatalogApi.GetMobileAppMetricCatalog(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileAppCatalogApi.GetMobileAppMetricCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMobileAppMetricCatalog`: []MobileAppMonitoringMetricDescription
    fmt.Fprintf(os.Stdout, "Response from `MobileAppCatalogApi.GetMobileAppMetricCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMobileAppMetricCatalogRequest struct via the builder pattern


### Return type

[**[]MobileAppMonitoringMetricDescription**](MobileAppMonitoringMetricDescription.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMobileAppTagCatalog

> TagCatalog GetMobileAppTagCatalog(ctx).BeaconType(beaconType).UseCase(useCase).Execute()

Get mobile app tag catalog

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
    beaconType := "beaconType_example" // string | 
    useCase := "useCase_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileAppCatalogApi.GetMobileAppTagCatalog(context.Background()).BeaconType(beaconType).UseCase(useCase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileAppCatalogApi.GetMobileAppTagCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMobileAppTagCatalog`: TagCatalog
    fmt.Fprintf(os.Stdout, "Response from `MobileAppCatalogApi.GetMobileAppTagCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMobileAppTagCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **beaconType** | **string** |  | 
 **useCase** | **string** |  | 

### Return type

[**TagCatalog**](TagCatalog.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

