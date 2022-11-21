# \InfrastructureResourcesApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMonitoringState**](InfrastructureResourcesApi.md#GetMonitoringState) | **Get** /api/infrastructure-monitoring/monitoring-state | Monitored host count
[**GetSnapshot**](InfrastructureResourcesApi.md#GetSnapshot) | **Get** /api/infrastructure-monitoring/snapshots/{id} | Get snapshot details
[**GetSnapshots**](InfrastructureResourcesApi.md#GetSnapshots) | **Get** /api/infrastructure-monitoring/snapshots | Search snapshots
[**SoftwareVersions**](InfrastructureResourcesApi.md#SoftwareVersions) | **Get** /api/infrastructure-monitoring/software/versions | Get installed software



## GetMonitoringState

> MonitoringState GetMonitoringState(ctx).Execute()

Monitored host count

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
    resp, r, err := apiClient.InfrastructureResourcesApi.GetMonitoringState(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureResourcesApi.GetMonitoringState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitoringState`: MonitoringState
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureResourcesApi.GetMonitoringState`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitoringStateRequest struct via the builder pattern


### Return type

[**MonitoringState**](MonitoringState.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshot

> SnapshotItem GetSnapshot(ctx, id).To(to).WindowSize(windowSize).Execute()

Get snapshot details

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
    to := int64(789) // int64 |  (optional)
    windowSize := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfrastructureResourcesApi.GetSnapshot(context.Background(), id).To(to).WindowSize(windowSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureResourcesApi.GetSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSnapshot`: SnapshotItem
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureResourcesApi.GetSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **to** | **int64** |  | 
 **windowSize** | **int64** |  | 

### Return type

[**SnapshotItem**](SnapshotItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshots

> SnapshotResult GetSnapshots(ctx).Query(query).To(to).WindowSize(windowSize).Size(size).Plugin(plugin).Offline(offline).Execute()

Search snapshots



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
    query := "query_example" // string |  (optional)
    to := int64(789) // int64 |  (optional)
    windowSize := int64(789) // int64 |  (optional)
    size := int32(56) // int32 |  (optional)
    plugin := "plugin_example" // string |  (optional)
    offline := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfrastructureResourcesApi.GetSnapshots(context.Background()).Query(query).To(to).WindowSize(windowSize).Size(size).Plugin(plugin).Offline(offline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureResourcesApi.GetSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSnapshots`: SnapshotResult
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureResourcesApi.GetSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **to** | **int64** |  | 
 **windowSize** | **int64** |  | 
 **size** | **int32** |  | 
 **plugin** | **string** |  | 
 **offline** | **bool** |  | 

### Return type

[**SnapshotResult**](SnapshotResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoftwareVersions

> []SoftwareVersion SoftwareVersions(ctx).Time(time).Origin(origin).Type_(type_).Name(name).Version(version).Execute()

Get installed software



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
    time := int64(789) // int64 |  (optional)
    origin := "origin_example" // string |  (optional)
    type_ := "type__example" // string |  (optional)
    name := "name_example" // string |  (optional)
    version := "version_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfrastructureResourcesApi.SoftwareVersions(context.Background()).Time(time).Origin(origin).Type_(type_).Name(name).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureResourcesApi.SoftwareVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SoftwareVersions`: []SoftwareVersion
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureResourcesApi.SoftwareVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSoftwareVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **time** | **int64** |  | 
 **origin** | **string** |  | 
 **type_** | **string** |  | 
 **name** | **string** |  | 
 **version** | **string** |  | 

### Return type

[**[]SoftwareVersion**](SoftwareVersion.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

