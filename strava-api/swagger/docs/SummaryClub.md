# SummaryClub

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The club&#39;s unique identifier. | [optional] [default to null]
**ResourceState** | **int32** | Resource state, indicates level of detail. Possible values: 1 -&gt; \&quot;meta\&quot;, 2 -&gt; \&quot;summary\&quot;, 3 -&gt; \&quot;detail\&quot; | [optional] [default to null]
**Name** | **string** | The club&#39;s name. | [optional] [default to null]
**ProfileMedium** | **string** | URL to a 60x60 pixel profile picture. | [optional] [default to null]
**CoverPhoto** | **string** | URL to a ~1185x580 pixel cover photo. | [optional] [default to null]
**CoverPhotoSmall** | **string** | URL to a ~360x176  pixel cover photo. | [optional] [default to null]
**SportType** | **string** | Deprecated. Prefer to use activity_types. | [optional] [default to null]
**ActivityTypes** | [**[]ActivityType**](ActivityType.md) | The activity types that count for a club. This takes precedence over sport_type. | [optional] [default to null]
**City** | **string** | The club&#39;s city. | [optional] [default to null]
**State** | **string** | The club&#39;s state or geographical region. | [optional] [default to null]
**Country** | **string** | The club&#39;s country. | [optional] [default to null]
**Private** | **bool** | Whether the club is private. | [optional] [default to null]
**MemberCount** | **int32** | The club&#39;s member count. | [optional] [default to null]
**Featured** | **bool** | Whether the club is featured or not. | [optional] [default to null]
**Verified** | **bool** | Whether the club is verified or not. | [optional] [default to null]
**Url** | **string** | The club&#39;s vanity URL. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


