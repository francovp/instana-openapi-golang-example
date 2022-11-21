# \EventsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](EventsApi.md#GetEvent) | **Get** /api/events/{eventId} | Get a particular event
[**GetEvents**](EventsApi.md#GetEvents) | **Get** /api/events | Get all events
[**GetEventsByIds**](EventsApi.md#GetEventsByIds) | **Post** /api/events | Get events given a set of IDs



## GetEvent

> EventResult GetEvent(ctx, eventId).Execute()

Get a particular event



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
    eventId := "eventId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEvent(context.Background(), eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvent`: EventResult
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventResult**](EventResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvents

> []EventResult GetEvents(ctx).EventTypeFilters(eventTypeFilters).WindowSize(windowSize).From(from).To(to).ExcludeTriggeredBefore(excludeTriggeredBefore).FilterEventUpdates(filterEventUpdates).Execute()

Get all events



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
    eventTypeFilters := []string{"EventTypeFilters_example"} // []string | 
    windowSize := int64(789) // int64 |  (optional)
    from := int64(789) // int64 |  (optional)
    to := int64(789) // int64 |  (optional)
    excludeTriggeredBefore := true // bool |  (optional)
    filterEventUpdates := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEvents(context.Background()).EventTypeFilters(eventTypeFilters).WindowSize(windowSize).From(from).To(to).ExcludeTriggeredBefore(excludeTriggeredBefore).FilterEventUpdates(filterEventUpdates).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvents`: []EventResult
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventTypeFilters** | **[]string** |  | 
 **windowSize** | **int64** |  | 
 **from** | **int64** |  | 
 **to** | **int64** |  | 
 **excludeTriggeredBefore** | **bool** |  | 
 **filterEventUpdates** | **bool** |  | 

### Return type

[**[]EventResult**](EventResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsByIds

> []EventResult GetEventsByIds(ctx).RequestBody(requestBody).Execute()

Get events given a set of IDs



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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEventsByIds(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEventsByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventsByIds`: []EventResult
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEventsByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

### Return type

[**[]EventResult**](EventResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

