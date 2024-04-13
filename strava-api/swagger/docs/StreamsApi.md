# \StreamsApi

All URIs are relative to *https://www.strava.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActivityStreams**](StreamsApi.md#GetActivityStreams) | **Get** /activities/{id}/streams | Get Activity Streams
[**GetRouteStreams**](StreamsApi.md#GetRouteStreams) | **Get** /routes/{id}/streams | Get Route Streams
[**GetSegmentEffortStreams**](StreamsApi.md#GetSegmentEffortStreams) | **Get** /segment_efforts/{id}/streams | Get Segment Effort Streams
[**GetSegmentStreams**](StreamsApi.md#GetSegmentStreams) | **Get** /segments/{id}/streams | Get Segment Streams


# **GetActivityStreams**
> StreamSet GetActivityStreams(ctx, id, keys, keyByType)
Get Activity Streams

Returns the given activity's streams. Requires activity:read scope. Requires activity:read_all scope for Only Me activities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The identifier of the activity. | 
  **keys** | [**[]string**](string.md)| Desired stream types. | 
  **keyByType** | **bool**| Must be true. | [default to true]

### Return type

[**StreamSet**](StreamSet.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRouteStreams**
> StreamSet GetRouteStreams(ctx, id)
Get Route Streams

Returns the given route's streams. Requires read_all scope for private routes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The identifier of the route. | 

### Return type

[**StreamSet**](StreamSet.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegmentEffortStreams**
> StreamSet GetSegmentEffortStreams(ctx, id, keys, keyByType)
Get Segment Effort Streams

Returns a set of streams for a segment effort completed by the authenticated athlete. Requires read_all scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The identifier of the segment effort. | 
  **keys** | [**[]string**](string.md)| The types of streams to return. | 
  **keyByType** | **bool**| Must be true. | [default to true]

### Return type

[**StreamSet**](StreamSet.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegmentStreams**
> StreamSet GetSegmentStreams(ctx, id, keys, keyByType)
Get Segment Streams

Returns the given segment's streams. Requires read_all scope for private segments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The identifier of the segment. | 
  **keys** | [**[]string**](string.md)| The types of streams to return. | 
  **keyByType** | **bool**| Must be true. | [default to true]

### Return type

[**StreamSet**](StreamSet.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

