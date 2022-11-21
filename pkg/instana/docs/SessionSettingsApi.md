# \SessionSettingsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSessionSettings**](SessionSettingsApi.md#DeleteSessionSettings) | **Delete** /api/settings/session | Delete session settings
[**GetSessionSettings**](SessionSettingsApi.md#GetSessionSettings) | **Get** /api/settings/session | Get session settings
[**SetSessionSettings**](SessionSettingsApi.md#SetSessionSettings) | **Put** /api/settings/session | Configure session settings



## DeleteSessionSettings

> DeleteSessionSettings(ctx).Execute()

Delete session settings

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
    resp, r, err := apiClient.SessionSettingsApi.DeleteSessionSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionSettingsApi.DeleteSessionSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionSettingsRequest struct via the builder pattern


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


## GetSessionSettings

> SessionSettings GetSessionSettings(ctx).Execute()

Get session settings

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
    resp, r, err := apiClient.SessionSettingsApi.GetSessionSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionSettingsApi.GetSessionSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionSettings`: SessionSettings
    fmt.Fprintf(os.Stdout, "Response from `SessionSettingsApi.GetSessionSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionSettingsRequest struct via the builder pattern


### Return type

[**SessionSettings**](SessionSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSessionSettings

> SessionSettings SetSessionSettings(ctx).SessionSettings(sessionSettings).Execute()

Configure session settings

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
    sessionSettings := *openapiclient.NewSessionSettings() // SessionSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionSettingsApi.SetSessionSettings(context.Background()).SessionSettings(sessionSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionSettingsApi.SetSessionSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetSessionSettings`: SessionSettings
    fmt.Fprintf(os.Stdout, "Response from `SessionSettingsApi.SetSessionSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetSessionSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionSettings** | [**SessionSettings**](SessionSettings.md) |  | 

### Return type

[**SessionSettings**](SessionSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

