# \APITokenApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiToken**](APITokenApi.md#DeleteApiToken) | **Delete** /api/settings/api-tokens/{internalId} | Delete API token
[**GetApiToken**](APITokenApi.md#GetApiToken) | **Get** /api/settings/api-tokens/{internalId} | API token
[**GetApiTokens**](APITokenApi.md#GetApiTokens) | **Get** /api/settings/api-tokens | All API tokens
[**PostApiToken**](APITokenApi.md#PostApiToken) | **Post** /api/settings/api-tokens | Create an API token
[**PutApiToken**](APITokenApi.md#PutApiToken) | **Put** /api/settings/api-tokens/{internalId} | Create or update an API token



## DeleteApiToken

> DeleteApiToken(ctx, internalId).Execute()

Delete API token

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
    internalId := "internalId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APITokenApi.DeleteApiToken(context.Background(), internalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APITokenApi.DeleteApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**internalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiTokenRequest struct via the builder pattern


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


## GetApiToken

> ApiToken GetApiToken(ctx, internalId).Execute()

API token

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
    internalId := "internalId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APITokenApi.GetApiToken(context.Background(), internalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APITokenApi.GetApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiToken`: ApiToken
    fmt.Fprintf(os.Stdout, "Response from `APITokenApi.GetApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**internalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiToken**](ApiToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiTokens

> []ApiToken GetApiTokens(ctx).Execute()

All API tokens

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
    resp, r, err := apiClient.APITokenApi.GetApiTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APITokenApi.GetApiTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiTokens`: []ApiToken
    fmt.Fprintf(os.Stdout, "Response from `APITokenApi.GetApiTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiTokensRequest struct via the builder pattern


### Return type

[**[]ApiToken**](ApiToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiToken

> ApiToken PostApiToken(ctx).ApiToken(apiToken).Execute()

Create an API token

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
    apiToken := *openapiclient.NewApiToken("AccessGrantingToken_example", "InternalId_example", "Name_example") // ApiToken | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APITokenApi.PostApiToken(context.Background()).ApiToken(apiToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APITokenApi.PostApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApiToken`: ApiToken
    fmt.Fprintf(os.Stdout, "Response from `APITokenApi.PostApiToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiToken** | [**ApiToken**](ApiToken.md) |  | 

### Return type

[**ApiToken**](ApiToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutApiToken

> ApiToken PutApiToken(ctx, internalId).ApiToken(apiToken).Execute()

Create or update an API token

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
    internalId := "internalId_example" // string | 
    apiToken := *openapiclient.NewApiToken("AccessGrantingToken_example", "InternalId_example", "Name_example") // ApiToken | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APITokenApi.PutApiToken(context.Background(), internalId).ApiToken(apiToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APITokenApi.PutApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutApiToken`: ApiToken
    fmt.Fprintf(os.Stdout, "Response from `APITokenApi.PutApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**internalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiToken** | [**ApiToken**](ApiToken.md) |  | 

### Return type

[**ApiToken**](ApiToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

