# DetailedSegment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The unique identifier of this segment | [optional] [default to null]
**Name** | **string** | The name of this segment | [optional] [default to null]
**ActivityType** | **string** |  | [optional] [default to null]
**Distance** | **float32** | The segment&#39;s distance, in meters | [optional] [default to null]
**AverageGrade** | **float32** | The segment&#39;s average grade, in percents | [optional] [default to null]
**MaximumGrade** | **float32** | The segments&#39;s maximum grade, in percents | [optional] [default to null]
**ElevationHigh** | **float32** | The segments&#39;s highest elevation, in meters | [optional] [default to null]
**ElevationLow** | **float32** | The segments&#39;s lowest elevation, in meters | [optional] [default to null]
**StartLatlng** | [***LatLng**](LatLng.md) |  | [optional] [default to null]
**EndLatlng** | [***LatLng**](LatLng.md) |  | [optional] [default to null]
**ClimbCategory** | **int32** | The category of the climb [0, 5]. Higher is harder ie. 5 is Hors catégorie, 0 is uncategorized in climb_category. | [optional] [default to null]
**City** | **string** | The segments&#39;s city. | [optional] [default to null]
**State** | **string** | The segments&#39;s state or geographical region. | [optional] [default to null]
**Country** | **string** | The segment&#39;s country. | [optional] [default to null]
**Private** | **bool** | Whether this segment is private. | [optional] [default to null]
**AthletePrEffort** | [***SummaryPrSegmentEffort**](SummaryPRSegmentEffort.md) |  | [optional] [default to null]
**AthleteSegmentStats** | [***SummarySegmentEffort**](SummarySegmentEffort.md) |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The time at which the segment was created. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The time at which the segment was last updated. | [optional] [default to null]
**TotalElevationGain** | **float32** | The segment&#39;s total elevation gain. | [optional] [default to null]
**Map_** | [***PolylineMap**](PolylineMap.md) |  | [optional] [default to null]
**EffortCount** | **int32** | The total number of efforts for this segment | [optional] [default to null]
**AthleteCount** | **int32** | The number of unique athletes who have an effort for this segment | [optional] [default to null]
**Hazardous** | **bool** | Whether this segment is considered hazardous | [optional] [default to null]
**StarCount** | **int32** | The number of stars for this segment | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


