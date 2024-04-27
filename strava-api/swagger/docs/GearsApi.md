# \GearsApi

All URIs are relative to *https://www.strava.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGearById**](GearsApi.md#GetGearById) | **Get** /gear/{id} | Get Equipment


# **GetGearById**
> DetailedGear GetGearById(ctx, id)
Get Equipment

Returns an equipment using its identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The identifier of the gear. | 

### Return type

[**DetailedGear**](DetailedGear.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

