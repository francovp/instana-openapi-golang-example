# \ApplicationCatalogApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationCatalogMetrics**](ApplicationCatalogApi.md#GetApplicationCatalogMetrics) | **Get** /api/application-monitoring/catalog/metrics | Get Metric catalog
[**GetApplicationTagCatalog**](ApplicationCatalogApi.md#GetApplicationTagCatalog) | **Get** /api/application-monitoring/catalog | Get application tag catalog
[**GetApplicationTags**](ApplicationCatalogApi.md#GetApplicationTags) | **Get** /api/application-monitoring/catalog/tags | Get application tags



## GetApplicationCatalogMetrics

> []MetricDescription GetApplicationCatalogMetrics(ctx).Execute()

Get Metric catalog



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
    resp, r, err := apiClient.ApplicationCatalogApi.GetApplicationCatalogMetrics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCatalogApi.GetApplicationCatalogMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationCatalogMetrics`: []MetricDescription
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCatalogApi.GetApplicationCatalogMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCatalogMetricsRequest struct via the builder pattern


### Return type

[**[]MetricDescription**](MetricDescription.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationTagCatalog

> TagCatalog GetApplicationTagCatalog(ctx).From(from).DataSource(dataSource).UseCase(useCase).Execute()

Get application tag catalog

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
    from := int64(789) // int64 |  (optional)
    dataSource := "dataSource_example" // string |  (optional)
    useCase := "useCase_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCatalogApi.GetApplicationTagCatalog(context.Background()).From(from).DataSource(dataSource).UseCase(useCase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCatalogApi.GetApplicationTagCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationTagCatalog`: TagCatalog
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCatalogApi.GetApplicationTagCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationTagCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** |  | 
 **dataSource** | **string** |  | 
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


## GetApplicationTags

> []Tag GetApplicationTags(ctx).From(from).DataSource(dataSource).UseCase(useCase).Execute()

Get application tags



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
    from := int64(789) // int64 |  (optional)
    dataSource := "dataSource_example" // string |  (optional)
    useCase := "useCase_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCatalogApi.GetApplicationTags(context.Background()).From(from).DataSource(dataSource).UseCase(useCase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCatalogApi.GetApplicationTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationTags`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCatalogApi.GetApplicationTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** |  | 
 **dataSource** | **string** |  | 
 **useCase** | **string** |  | 

### Return type

[**[]Tag**](Tag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

