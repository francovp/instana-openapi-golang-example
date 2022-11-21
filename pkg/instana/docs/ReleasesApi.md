# \ReleasesApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRelease**](ReleasesApi.md#DeleteRelease) | **Delete** /api/releases/{releaseId} | Delete release
[**GetAllReleases**](ReleasesApi.md#GetAllReleases) | **Get** /api/releases | Get all releases
[**GetRelease**](ReleasesApi.md#GetRelease) | **Get** /api/releases/{releaseId} | Get release
[**PostRelease**](ReleasesApi.md#PostRelease) | **Post** /api/releases | Create release
[**PutRelease**](ReleasesApi.md#PutRelease) | **Put** /api/releases/{releaseId} | Update release



## DeleteRelease

> DeleteRelease(ctx, releaseId).Execute()

Delete release

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
    releaseId := "releaseId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesApi.DeleteRelease(context.Background(), releaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesApi.DeleteRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReleaseRequest struct via the builder pattern


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


## GetAllReleases

> []ReleaseWithMetadata GetAllReleases(ctx).From(from).To(to).MaxResults(maxResults).Execute()

Get all releases



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
    to := int64(789) // int64 |  (optional)
    maxResults := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesApi.GetAllReleases(context.Background()).From(from).To(to).MaxResults(maxResults).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesApi.GetAllReleases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllReleases`: []ReleaseWithMetadata
    fmt.Fprintf(os.Stdout, "Response from `ReleasesApi.GetAllReleases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** |  | 
 **to** | **int64** |  | 
 **maxResults** | **int32** |  | 

### Return type

[**[]ReleaseWithMetadata**](ReleaseWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelease

> ReleaseWithMetadata GetRelease(ctx, releaseId).Execute()

Get release

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
    releaseId := "releaseId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesApi.GetRelease(context.Background(), releaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesApi.GetRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRelease`: ReleaseWithMetadata
    fmt.Fprintf(os.Stdout, "Response from `ReleasesApi.GetRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReleaseWithMetadata**](ReleaseWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRelease

> ReleaseWithMetadata PostRelease(ctx).Release(release).Execute()

Create release

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
    release := *openapiclient.NewRelease("Name_example", int64(123)) // Release | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesApi.PostRelease(context.Background()).Release(release).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesApi.PostRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostRelease`: ReleaseWithMetadata
    fmt.Fprintf(os.Stdout, "Response from `ReleasesApi.PostRelease`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **release** | [**Release**](Release.md) |  | 

### Return type

[**ReleaseWithMetadata**](ReleaseWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRelease

> ReleaseWithMetadata PutRelease(ctx, releaseId).Release(release).Execute()

Update release

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
    releaseId := "releaseId_example" // string | 
    release := *openapiclient.NewRelease("Name_example", int64(123)) // Release | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesApi.PutRelease(context.Background(), releaseId).Release(release).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesApi.PutRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutRelease`: ReleaseWithMetadata
    fmt.Fprintf(os.Stdout, "Response from `ReleasesApi.PutRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **release** | [**Release**](Release.md) |  | 

### Return type

[**ReleaseWithMetadata**](ReleaseWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

