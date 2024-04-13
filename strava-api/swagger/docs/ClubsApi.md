# \ClubsApi

All URIs are relative to *https://www.strava.com/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClubActivitiesById**](ClubsApi.md#GetClubActivitiesById) | **Get** /clubs/{id}/activities | List Club Activities
[**GetClubAdminsById**](ClubsApi.md#GetClubAdminsById) | **Get** /clubs/{id}/admins | List Club Administrators
[**GetClubById**](ClubsApi.md#GetClubById) | **Get** /clubs/{id} | Get Club
[**GetClubMembersById**](ClubsApi.md#GetClubMembersById) | **Get** /clubs/{id}/members | List Club Members
[**GetLoggedInAthleteClubs**](ClubsApi.md#GetLoggedInAthleteClubs) | **Get** /athlete/clubs | List Athlete Clubs


# **GetClubActivitiesById**
> []ClubActivity GetClubActivitiesById(ctx, id, optional)
List Club Activities

Retrieve recent activities from members of a specific club. The authenticated athlete must belong to the requested club in order to hit this endpoint. Pagination is supported. Athlete profile visibility is respected for all activities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The identifier of the club. | 
 **optional** | ***ClubsApiGetClubActivitiesByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClubsApiGetClubActivitiesByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. Defaults to 1. | 
 **perPage** | **optional.Int32**| Number of items per page. Defaults to 30. | [default to 30]

### Return type

[**[]ClubActivity**](ClubActivity.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClubAdminsById**
> []SummaryAthlete GetClubAdminsById(ctx, id, optional)
List Club Administrators

Returns a list of the administrators of a given club.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The identifier of the club. | 
 **optional** | ***ClubsApiGetClubAdminsByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClubsApiGetClubAdminsByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. Defaults to 1. | 
 **perPage** | **optional.Int32**| Number of items per page. Defaults to 30. | [default to 30]

### Return type

[**[]SummaryAthlete**](SummaryAthlete.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClubById**
> DetailedClub GetClubById(ctx, id)
Get Club

Returns a given club using its identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The identifier of the club. | 

### Return type

[**DetailedClub**](DetailedClub.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClubMembersById**
> []ClubAthlete GetClubMembersById(ctx, id, optional)
List Club Members

Returns a list of the athletes who are members of a given club.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The identifier of the club. | 
 **optional** | ***ClubsApiGetClubMembersByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClubsApiGetClubMembersByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page number. Defaults to 1. | 
 **perPage** | **optional.Int32**| Number of items per page. Defaults to 30. | [default to 30]

### Return type

[**[]ClubAthlete**](ClubAthlete.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoggedInAthleteClubs**
> []SummaryClub GetLoggedInAthleteClubs(ctx, optional)
List Athlete Clubs

Returns a list of the clubs whose membership includes the authenticated athlete.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClubsApiGetLoggedInAthleteClubsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClubsApiGetLoggedInAthleteClubsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number. Defaults to 1. | 
 **perPage** | **optional.Int32**| Number of items per page. Defaults to 30. | [default to 30]

### Return type

[**[]SummaryClub**](SummaryClub.md)

### Authorization

[strava_oauth](../README.md#strava_oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

