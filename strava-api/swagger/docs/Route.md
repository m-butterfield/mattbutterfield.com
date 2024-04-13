# Route

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Athlete** | [***SummaryAthlete**](SummaryAthlete.md) |  | [optional] [default to null]
**Description** | **string** | The description of the route | [optional] [default to null]
**Distance** | **float32** | The route&#39;s distance, in meters | [optional] [default to null]
**ElevationGain** | **float32** | The route&#39;s elevation gain. | [optional] [default to null]
**Id** | **int64** | The unique identifier of this route | [optional] [default to null]
**IdStr** | **string** | The unique identifier of the route in string format | [optional] [default to null]
**Map_** | [***PolylineMap**](PolylineMap.md) |  | [optional] [default to null]
**Name** | **string** | The name of this route | [optional] [default to null]
**Private** | **bool** | Whether this route is private | [optional] [default to null]
**Starred** | **bool** | Whether this route is starred by the logged-in athlete | [optional] [default to null]
**Timestamp** | **int32** | An epoch timestamp of when the route was created | [optional] [default to null]
**Type_** | **int32** | This route&#39;s type (1 for ride, 2 for runs) | [optional] [default to null]
**SubType** | **int32** | This route&#39;s sub-type (1 for road, 2 for mountain bike, 3 for cross, 4 for trail, 5 for mixed) | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The time at which the route was created | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The time at which the route was last updated | [optional] [default to null]
**EstimatedMovingTime** | **int32** | Estimated time in seconds for the authenticated athlete to complete route | [optional] [default to null]
**Segments** | [**[]SummarySegment**](SummarySegment.md) | The segments traversed by this route | [optional] [default to null]
**Waypoints** | [**[]Waypoint**](Waypoint.md) | The custom waypoints along this route | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


