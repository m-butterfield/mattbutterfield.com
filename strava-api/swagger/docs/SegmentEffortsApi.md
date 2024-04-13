# \SegmentEffortsApi

All URIs are relative to *https://www.strava.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEffortsBySegmentId**](SegmentEffortsApi.md#GetEffortsBySegmentId) | **Get** /segment_efforts | List Segment Efforts
[**GetSegmentEffortById**](SegmentEffortsApi.md#GetSegmentEffortById) | **Get** /segment_efforts/{id} | Get Segment Effort


# **GetEffortsBySegmentId**
> []DetailedSegmentEffort GetEffortsBySegmentId(ctx, segmentId, optional)
List Segment Efforts

Returns a set of the authenticated athlete's segment efforts for a given segment.  Requires subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **segmentId** | **int32**| The identifier of the segment. | 
 **optional** | ***SegmentEffortsApiGetEffortsBySegmentIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SegmentEffortsApiGetEffortsBySegmentIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDateLocal** | **optional.Time**| ISO 8601 formatted date time. | 
 **endDateLocal** | **optional.Time**| ISO 8601 formatted date time. | 
 **perPage** | **optional.Int32**| Number of items per page. Defaults to 30. | [default to 30]

### Return type

[**[]DetailedSegmentEffort**](DetailedSegmentEffort.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegmentEffortById**
> DetailedSegmentEffort GetSegmentEffortById(ctx, id)
Get Segment Effort

Returns a segment effort from an activity that is owned by the authenticated athlete. Requires subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The identifier of the segment effort. | 

### Return type

[**DetailedSegmentEffort**](DetailedSegmentEffort.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

