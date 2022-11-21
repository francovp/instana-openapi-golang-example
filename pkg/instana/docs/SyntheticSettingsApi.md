# \SyntheticSettingsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSyntheticTest**](SyntheticSettingsApi.md#CreateSyntheticTest) | **Post** /api/synthetics/settings/tests | Experimental - Create a Synthetic test
[**DeleteSyntheticLocation**](SyntheticSettingsApi.md#DeleteSyntheticLocation) | **Delete** /api/synthetics/settings/locations/{id} | Experimental - Delete Synthetic location
[**DeleteSyntheticTest**](SyntheticSettingsApi.md#DeleteSyntheticTest) | **Delete** /api/synthetics/settings/tests/{id} | Experimental - Delete a Synthetic test
[**GetSyntheticLocation**](SyntheticSettingsApi.md#GetSyntheticLocation) | **Get** /api/synthetics/settings/locations/{id} | Experimental - Synthetic location
[**GetSyntheticLocations**](SyntheticSettingsApi.md#GetSyntheticLocations) | **Get** /api/synthetics/settings/locations | Experimental - All Synthetic locations
[**GetSyntheticTest**](SyntheticSettingsApi.md#GetSyntheticTest) | **Get** /api/synthetics/settings/tests/{id} | Experimental - A Synthetic test
[**GetSyntheticTests**](SyntheticSettingsApi.md#GetSyntheticTests) | **Get** /api/synthetics/settings/tests | Experimental - All Synthetic tests
[**UpdateSyntheticTest**](SyntheticSettingsApi.md#UpdateSyntheticTest) | **Put** /api/synthetics/settings/tests/{id} | Experimental - Update a Synthetic test



## CreateSyntheticTest

> SyntheticTest CreateSyntheticTest(ctx).SyntheticTest(syntheticTest).Execute()

Experimental - Create a Synthetic test



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
    syntheticTest := *openapiclient.NewSyntheticTest(false, *openapiclient.NewSyntheticTypeConfiguration(false, "SyntheticType_example"), "Label_example", []string{"Locations_example"}, "PlaybackMode_example", int32(123)) // SyntheticTest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticSettingsApi.CreateSyntheticTest(context.Background()).SyntheticTest(syntheticTest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticSettingsApi.CreateSyntheticTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSyntheticTest`: SyntheticTest
    fmt.Fprintf(os.Stdout, "Response from `SyntheticSettingsApi.CreateSyntheticTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSyntheticTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syntheticTest** | [**SyntheticTest**](SyntheticTest.md) |  | 

### Return type

[**SyntheticTest**](SyntheticTest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSyntheticLocation

> DeleteSyntheticLocation(ctx, id).Execute()

Experimental - Delete Synthetic location

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
    resp, r, err := apiClient.SyntheticSettingsApi.DeleteSyntheticLocation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticSettingsApi.DeleteSyntheticLocation``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSyntheticLocationRequest struct via the builder pattern


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


## DeleteSyntheticTest

> DeleteSyntheticTest(ctx, id).Execute()

Experimental - Delete a Synthetic test

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
    resp, r, err := apiClient.SyntheticSettingsApi.DeleteSyntheticTest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticSettingsApi.DeleteSyntheticTest``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSyntheticTestRequest struct via the builder pattern


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


## GetSyntheticLocation

> SyntheticLocation GetSyntheticLocation(ctx, id).Execute()

Experimental - Synthetic location

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
    resp, r, err := apiClient.SyntheticSettingsApi.GetSyntheticLocation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticSettingsApi.GetSyntheticLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSyntheticLocation`: SyntheticLocation
    fmt.Fprintf(os.Stdout, "Response from `SyntheticSettingsApi.GetSyntheticLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyntheticLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyntheticLocation**](SyntheticLocation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyntheticLocations

> []SyntheticLocation GetSyntheticLocations(ctx).Execute()

Experimental - All Synthetic locations

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
    resp, r, err := apiClient.SyntheticSettingsApi.GetSyntheticLocations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticSettingsApi.GetSyntheticLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSyntheticLocations`: []SyntheticLocation
    fmt.Fprintf(os.Stdout, "Response from `SyntheticSettingsApi.GetSyntheticLocations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyntheticLocationsRequest struct via the builder pattern


### Return type

[**[]SyntheticLocation**](SyntheticLocation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyntheticTest

> SyntheticTest GetSyntheticTest(ctx, id).Execute()

Experimental - A Synthetic test

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
    resp, r, err := apiClient.SyntheticSettingsApi.GetSyntheticTest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticSettingsApi.GetSyntheticTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSyntheticTest`: SyntheticTest
    fmt.Fprintf(os.Stdout, "Response from `SyntheticSettingsApi.GetSyntheticTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyntheticTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyntheticTest**](SyntheticTest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyntheticTests

> []SyntheticTest GetSyntheticTests(ctx).ApplicationId(applicationId).Execute()

Experimental - All Synthetic tests

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticSettingsApi.GetSyntheticTests(context.Background()).ApplicationId(applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticSettingsApi.GetSyntheticTests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSyntheticTests`: []SyntheticTest
    fmt.Fprintf(os.Stdout, "Response from `SyntheticSettingsApi.GetSyntheticTests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSyntheticTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | **string** |  | 

### Return type

[**[]SyntheticTest**](SyntheticTest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyntheticTest

> SyntheticTest UpdateSyntheticTest(ctx, id).SyntheticTest(syntheticTest).Execute()

Experimental - Update a Synthetic test

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
    syntheticTest := *openapiclient.NewSyntheticTest(false, *openapiclient.NewSyntheticTypeConfiguration(false, "SyntheticType_example"), "Label_example", []string{"Locations_example"}, "PlaybackMode_example", int32(123)) // SyntheticTest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticSettingsApi.UpdateSyntheticTest(context.Background(), id).SyntheticTest(syntheticTest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticSettingsApi.UpdateSyntheticTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSyntheticTest`: SyntheticTest
    fmt.Fprintf(os.Stdout, "Response from `SyntheticSettingsApi.UpdateSyntheticTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSyntheticTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syntheticTest** | [**SyntheticTest**](SyntheticTest.md) |  | 

### Return type

[**SyntheticTest**](SyntheticTest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

