# \MobileAppAnalyzeApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMobileAppBeaconGroups**](MobileAppAnalyzeApi.md#GetMobileAppBeaconGroups) | **Post** /api/mobile-app-monitoring/analyze/beacon-groups | Get grouped beacon metrics
[**GetMobileAppBeacons**](MobileAppAnalyzeApi.md#GetMobileAppBeacons) | **Post** /api/mobile-app-monitoring/analyze/beacons | Get all beacons



## GetMobileAppBeaconGroups

> MobileAppBeaconGroupsResult GetMobileAppBeaconGroups(ctx).FillTimeSeries(fillTimeSeries).GetMobileAppBeaconGroups(getMobileAppBeaconGroups).Execute()

Get grouped beacon metrics

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
    fillTimeSeries := true // bool |  (optional)
    getMobileAppBeaconGroups := *openapiclient.NewGetMobileAppBeaconGroups(*openapiclient.NewMobileAppBeaconTagGroup("GroupbyTag_example", "GroupbyTagEntity_example"), []openapiclient.MobileAppMonitoringMetricsConfiguration{*openapiclient.NewMobileAppMonitoringMetricsConfiguration("Aggregation_example", "Metric_example")}, "Type_example") // GetMobileAppBeaconGroups |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileAppAnalyzeApi.GetMobileAppBeaconGroups(context.Background()).FillTimeSeries(fillTimeSeries).GetMobileAppBeaconGroups(getMobileAppBeaconGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileAppAnalyzeApi.GetMobileAppBeaconGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMobileAppBeaconGroups`: MobileAppBeaconGroupsResult
    fmt.Fprintf(os.Stdout, "Response from `MobileAppAnalyzeApi.GetMobileAppBeaconGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMobileAppBeaconGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fillTimeSeries** | **bool** |  | 
 **getMobileAppBeaconGroups** | [**GetMobileAppBeaconGroups**](GetMobileAppBeaconGroups.md) |  | 

### Return type

[**MobileAppBeaconGroupsResult**](MobileAppBeaconGroupsResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMobileAppBeacons

> MobileAppBeaconResult GetMobileAppBeacons(ctx).GetMobileAppBeacons(getMobileAppBeacons).Execute()

Get all beacons

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
    getMobileAppBeacons := *openapiclient.NewGetMobileAppBeacons("Type_example") // GetMobileAppBeacons |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileAppAnalyzeApi.GetMobileAppBeacons(context.Background()).GetMobileAppBeacons(getMobileAppBeacons).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileAppAnalyzeApi.GetMobileAppBeacons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMobileAppBeacons`: MobileAppBeaconResult
    fmt.Fprintf(os.Stdout, "Response from `MobileAppAnalyzeApi.GetMobileAppBeacons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMobileAppBeaconsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getMobileAppBeacons** | [**GetMobileAppBeacons**](GetMobileAppBeacons.md) |  | 

### Return type

[**MobileAppBeaconResult**](MobileAppBeaconResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

