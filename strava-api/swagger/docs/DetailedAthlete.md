# DetailedAthlete

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The unique identifier of the athlete | [optional] [default to null]
**ResourceState** | **int32** | Resource state, indicates level of detail. Possible values: 1 -&gt; \&quot;meta\&quot;, 2 -&gt; \&quot;summary\&quot;, 3 -&gt; \&quot;detail\&quot; | [optional] [default to null]
**Firstname** | **string** | The athlete&#39;s first name. | [optional] [default to null]
**Lastname** | **string** | The athlete&#39;s last name. | [optional] [default to null]
**ProfileMedium** | **string** | URL to a 62x62 pixel profile picture. | [optional] [default to null]
**Profile** | **string** | URL to a 124x124 pixel profile picture. | [optional] [default to null]
**City** | **string** | The athlete&#39;s city. | [optional] [default to null]
**State** | **string** | The athlete&#39;s state or geographical region. | [optional] [default to null]
**Country** | **string** | The athlete&#39;s country. | [optional] [default to null]
**Sex** | **string** | The athlete&#39;s sex. | [optional] [default to null]
**Premium** | **bool** | Deprecated.  Use summit field instead. Whether the athlete has any Summit subscription. | [optional] [default to null]
**Summit** | **bool** | Whether the athlete has any Summit subscription. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The time at which the athlete was created. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The time at which the athlete was last updated. | [optional] [default to null]
**FollowerCount** | **int32** | The athlete&#39;s follower count. | [optional] [default to null]
**FriendCount** | **int32** | The athlete&#39;s friend count. | [optional] [default to null]
**MeasurementPreference** | **string** | The athlete&#39;s preferred unit system. | [optional] [default to null]
**Ftp** | **int32** | The athlete&#39;s FTP (Functional Threshold Power). | [optional] [default to null]
**Weight** | **float32** | The athlete&#39;s weight. | [optional] [default to null]
**Clubs** | [**[]SummaryClub**](SummaryClub.md) | The athlete&#39;s clubs. | [optional] [default to null]
**Bikes** | [**[]SummaryGear**](SummaryGear.md) | The athlete&#39;s bikes. | [optional] [default to null]
**Shoes** | [**[]SummaryGear**](SummaryGear.md) | The athlete&#39;s shoes. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


