# Lap

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The unique identifier of this lap | [optional] [default to null]
**Activity** | [***MetaActivity**](MetaActivity.md) |  | [optional] [default to null]
**Athlete** | [***MetaAthlete**](MetaAthlete.md) |  | [optional] [default to null]
**AverageCadence** | **float32** | The lap&#39;s average cadence | [optional] [default to null]
**AverageSpeed** | **float32** | The lap&#39;s average speed | [optional] [default to null]
**Distance** | **float32** | The lap&#39;s distance, in meters | [optional] [default to null]
**ElapsedTime** | **int32** | The lap&#39;s elapsed time, in seconds | [optional] [default to null]
**StartIndex** | **int32** | The start index of this effort in its activity&#39;s stream | [optional] [default to null]
**EndIndex** | **int32** | The end index of this effort in its activity&#39;s stream | [optional] [default to null]
**LapIndex** | **int32** | The index of this lap in the activity it belongs to | [optional] [default to null]
**MaxSpeed** | **float32** | The maximum speed of this lat, in meters per second | [optional] [default to null]
**MovingTime** | **int32** | The lap&#39;s moving time, in seconds | [optional] [default to null]
**Name** | **string** | The name of the lap | [optional] [default to null]
**PaceZone** | **int32** | The athlete&#39;s pace zone during this lap | [optional] [default to null]
**Split** | **int32** |  | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | The time at which the lap was started. | [optional] [default to null]
**StartDateLocal** | [**time.Time**](time.Time.md) | The time at which the lap was started in the local timezone. | [optional] [default to null]
**TotalElevationGain** | **float32** | The elevation gain of this lap, in meters | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


