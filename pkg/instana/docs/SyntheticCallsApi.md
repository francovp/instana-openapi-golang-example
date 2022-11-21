# \SyntheticCallsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSyntheticCall**](SyntheticCallsApi.md#DeleteSyntheticCall) | **Delete** /api/settings/synthetic-calls | Delete synthetic call configurations
[**GetSyntheticCalls**](SyntheticCallsApi.md#GetSyntheticCalls) | **Get** /api/settings/synthetic-calls | Synthetic call configurations
[**UpdateSyntheticCall**](SyntheticCallsApi.md#UpdateSyntheticCall) | **Put** /api/settings/synthetic-calls | Update synthetic call configurations



## DeleteSyntheticCall

> DeleteSyntheticCall(ctx).Execute()

Delete synthetic call configurations

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
    resp, r, err := apiClient.SyntheticCallsApi.DeleteSyntheticCall(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticCallsApi.DeleteSyntheticCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSyntheticCallRequest struct via the builder pattern


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


## GetSyntheticCalls

> SyntheticCallWithDefaultsConfig GetSyntheticCalls(ctx).Execute()

Synthetic call configurations

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
    resp, r, err := apiClient.SyntheticCallsApi.GetSyntheticCalls(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticCallsApi.GetSyntheticCalls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSyntheticCalls`: SyntheticCallWithDefaultsConfig
    fmt.Fprintf(os.Stdout, "Response from `SyntheticCallsApi.GetSyntheticCalls`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyntheticCallsRequest struct via the builder pattern


### Return type

[**SyntheticCallWithDefaultsConfig**](SyntheticCallWithDefaultsConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyntheticCall

> UpdateSyntheticCall(ctx).SyntheticCallConfig(syntheticCallConfig).Execute()

Update synthetic call configurations

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
    syntheticCallConfig := *openapiclient.NewSyntheticCallConfig([]openapiclient.SyntheticCallRule{*openapiclient.NewSyntheticCallRule(*openapiclient.NewMatchExpressionDTO("Type_example"), "Name_example")}) // SyntheticCallConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticCallsApi.UpdateSyntheticCall(context.Background()).SyntheticCallConfig(syntheticCallConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticCallsApi.UpdateSyntheticCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSyntheticCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syntheticCallConfig** | [**SyntheticCallConfig**](SyntheticCallConfig.md) |  | 

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

