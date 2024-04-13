# \RoutesApi

All URIs are relative to *https://www.strava.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRouteAsGPX**](RoutesApi.md#GetRouteAsGPX) | **Get** /routes/{id}/export_gpx | Export Route GPX
[**GetRouteAsTCX**](RoutesApi.md#GetRouteAsTCX) | **Get** /routes/{id}/export_tcx | Export Route TCX
[**GetRouteById**](RoutesApi.md#GetRouteById) | **Get** /routes/{id} | Get Route
[**GetRoutesByAthleteId**](RoutesApi.md#GetRoutesByAthleteId) | **Get** /athletes/{id}/routes | List Athlete Routes


# **GetRouteAsGPX**
> GetRouteAsGPX(ctx, id)
Export Route GPX

Returns a GPX file of the route. Requires read_all scope for private routes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The identifier of the route. | 

### Return type

 (empty response body)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRouteAsTCX**
> GetRouteAsTCX(ctx, id)
Export Route TCX

Returns a TCX file of the route. Requires read_all scope for private routes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The identifier of the route. | 

### Return type

 (empty response body)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRouteById**
> Route GetRouteById(ctx, id)
Get Route

Returns a route using its identifier. Requires read_all scope for private routes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The identifier of the route. | 

### Return type

[**Route**](Route.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoutesByAthleteId**
> []Route GetRoutesByAthleteId(ctx, optional)
List Athlete Routes

Returns a list of the routes created by the authenticated athlete. Private routes are filtered out unless requested by a token with read_all scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RoutesApiGetRoutesByAthleteIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoutesApiGetRoutesByAthleteIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number. Defaults to 1. | 
 **perPage** | **optional.Int32**| Number of items per page. Defaults to 30. | [default to 30]

### Return type

[**[]Route**](Route.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

