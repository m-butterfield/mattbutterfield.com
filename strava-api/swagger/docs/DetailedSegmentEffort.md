# DetailedSegmentEffort

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The unique identifier of this effort | [optional] [default to null]
**ActivityId** | **int64** | The unique identifier of the activity related to this effort | [optional] [default to null]
**ElapsedTime** | **int32** | The effort&#39;s elapsed time | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | The time at which the effort was started. | [optional] [default to null]
**StartDateLocal** | [**time.Time**](time.Time.md) | The time at which the effort was started in the local timezone. | [optional] [default to null]
**Distance** | **float32** | The effort&#39;s distance in meters | [optional] [default to null]
**IsKom** | **bool** | Whether this effort is the current best on the leaderboard | [optional] [default to null]
**Name** | **string** | The name of the segment on which this effort was performed | [optional] [default to null]
**Activity** | [***MetaActivity**](MetaActivity.md) |  | [optional] [default to null]
**Athlete** | [***MetaAthlete**](MetaAthlete.md) |  | [optional] [default to null]
**MovingTime** | **int32** | The effort&#39;s moving time | [optional] [default to null]
**StartIndex** | **int32** | The start index of this effort in its activity&#39;s stream | [optional] [default to null]
**EndIndex** | **int32** | The end index of this effort in its activity&#39;s stream | [optional] [default to null]
**AverageCadence** | **float32** | The effort&#39;s average cadence | [optional] [default to null]
**AverageWatts** | **float32** | The average wattage of this effort | [optional] [default to null]
**DeviceWatts** | **bool** | For riding efforts, whether the wattage was reported by a dedicated recording device | [optional] [default to null]
**AverageHeartrate** | **float32** | The heart heart rate of the athlete during this effort | [optional] [default to null]
**MaxHeartrate** | **float32** | The maximum heart rate of the athlete during this effort | [optional] [default to null]
**Segment** | [***SummarySegment**](SummarySegment.md) |  | [optional] [default to null]
**KomRank** | **int32** | The rank of the effort on the global leaderboard if it belongs in the top 10 at the time of upload | [optional] [default to null]
**PrRank** | **int32** | The rank of the effort on the athlete&#39;s leaderboard if it belongs in the top 3 at the time of upload | [optional] [default to null]
**Hidden** | **bool** | Whether this effort should be hidden when viewed within an activity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


