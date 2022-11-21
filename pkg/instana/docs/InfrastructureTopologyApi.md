# \InfrastructureTopologyApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRelatedHosts**](InfrastructureTopologyApi.md#GetRelatedHosts) | **Get** /api/infrastructure-monitoring/graph/related-hosts/{snapshotId} | Related hosts
[**GetTopology**](InfrastructureTopologyApi.md#GetTopology) | **Get** /api/infrastructure-monitoring/topology | Gets the infrastructure topology



## GetRelatedHosts

> []string GetRelatedHosts(ctx, snapshotId).Execute()

Related hosts

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
    snapshotId := "snapshotId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfrastructureTopologyApi.GetRelatedHosts(context.Background(), snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureTopologyApi.GetRelatedHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRelatedHosts`: []string
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureTopologyApi.GetRelatedHosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRelatedHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopology

> Topology GetTopology(ctx).IncludeData(includeData).Execute()

Gets the infrastructure topology

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
    includeData := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfrastructureTopologyApi.GetTopology(context.Background()).IncludeData(includeData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfrastructureTopologyApi.GetTopology``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopology`: Topology
    fmt.Fprintf(os.Stdout, "Response from `InfrastructureTopologyApi.GetTopology`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTopologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeData** | **bool** |  | 

### Return type

[**Topology**](Topology.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

