# UpdatableActivity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commute** | **bool** | Whether this activity is a commute | [optional] [default to null]
**Trainer** | **bool** | Whether this activity was recorded on a training machine | [optional] [default to null]
**HideFromHome** | **bool** | Whether this activity is muted | [optional] [default to null]
**Description** | **string** | The description of the activity | [optional] [default to null]
**Name** | **string** | The name of the activity | [optional] [default to null]
**Type_** | [***ActivityType**](ActivityType.md) | Deprecated. Prefer to use sport_type. In a request where both type and sport_type are present, this field will be ignored | [optional] [default to null]
**SportType** | [***SportType**](SportType.md) |  | [optional] [default to null]
**GearId** | **string** | Identifier for the gear associated with the activity. ‘none’ clears gear from activity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


