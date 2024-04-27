# \AthletesApi

All URIs are relative to *https://www.strava.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLoggedInAthlete**](AthletesApi.md#GetLoggedInAthlete) | **Get** /athlete | Get Authenticated Athlete
[**GetLoggedInAthleteZones**](AthletesApi.md#GetLoggedInAthleteZones) | **Get** /athlete/zones | Get Zones
[**GetStats**](AthletesApi.md#GetStats) | **Get** /athletes/{id}/stats | Get Athlete Stats
[**UpdateLoggedInAthlete**](AthletesApi.md#UpdateLoggedInAthlete) | **Put** /athlete | Update Athlete


# **GetLoggedInAthlete**
> DetailedAthlete GetLoggedInAthlete(ctx, )
Get Authenticated Athlete

Returns the currently authenticated athlete. Tokens with profile:read_all scope will receive a detailed athlete representation; all others will receive a summary representation.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DetailedAthlete**](DetailedAthlete.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoggedInAthleteZones**
> Zones GetLoggedInAthleteZones(ctx, )
Get Zones

Returns the the authenticated athlete's heart rate and power zones. Requires profile:read_all.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Zones**](Zones.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStats**
> ActivityStats GetStats(ctx, id)
Get Athlete Stats

Returns the activity stats of an athlete. Only includes data from activities set to Everyone visibilty.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The identifier of the athlete. Must match the authenticated athlete. | 

### Return type

[**ActivityStats**](ActivityStats.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoggedInAthlete**
> DetailedAthlete UpdateLoggedInAthlete(ctx, weight)
Update Athlete

Update the currently authenticated athlete. Requires profile:write scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **weight** | **float32**| The weight of the athlete in kilograms. | 

### Return type

[**DetailedAthlete**](DetailedAthlete.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

