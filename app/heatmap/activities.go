package heatmap

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"github.com/m-butterfield/mattbutterfield.com/strava-api/swagger"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func getLatestActivities(ds data.Store) error {
	client, auth, err := getStravaClient(ds)
	if err != nil {
		return err
	}

	latestActivity, err := ds.GetLatestStravaActivity()
	if err != nil {
		return err
	}

	after := int32(0)
	if latestActivity == nil {
		log.Print("Fetching initial activities")
	} else {
		log.Printf("Fetching activities after last saved activity on %s", latestActivity.StartDate.Format(time.RFC3339))
		after = int32(latestActivity.StartDate.Unix())
	}
	activities, r, err := client.ActivitiesApi.GetLoggedInAthleteActivities(auth, &swagger.ActivitiesApiGetLoggedInAthleteActivitiesOpts{
		After: optional.NewInt32(after),
	})
	if r.StatusCode == 429 {
		if err = checkRateLimitUsage(r.Header); err != nil {
			return err
		}
		return getLatestActivities(ds)
	}
	if err != nil {
		return err
	}
	if r.StatusCode != 200 {
		return fmt.Errorf("unexpected status code when fetching activities: %d", r.StatusCode)
	}

	for len(activities) > 0 {
		for _, activity := range activities {
			if *activity.SportType == "VirtualRide" {
				log.Print("Skipping virtual ride")
				continue
			}
			if activity.Private {
				log.Print("Skipping private activity")
				continue
			}
			if activity.Manual {
				log.Print("Skipping manually created activity")
				continue
			}
			if isNew, err := saveActivity(ds, client, auth, activity); err != nil {
				return err
			} else if !isNew {
				return nil
			}
		}
		activities, r, err = client.ActivitiesApi.GetLoggedInAthleteActivities(auth, &swagger.ActivitiesApiGetLoggedInAthleteActivitiesOpts{
			After: optional.NewInt32(int32(activities[len(activities)-1].StartDate.Unix())),
		})
		if r.StatusCode == 429 {
			if err = checkRateLimitUsage(r.Header); err != nil {
				return err
			}
			return getLatestActivities(ds)
		}
		if err != nil {
			return err
		}
		if r.StatusCode != 200 {
			return fmt.Errorf("unexpected status code when fetching activities: %d", r.StatusCode)
		}
	}

	return nil
}

func checkRateLimitUsage(header http.Header) error {
	limitInfo := header["X-Readratelimit-Limit"][0]
	limits := strings.Split(limitInfo, ",")
	minute15Limit, _ := strconv.Atoi(limits[0])
	dayLimit, _ := strconv.Atoi(limits[1])

	usageInfo := header["X-Readratelimit-Usage"][0]
	usages := strings.Split(usageInfo, ",")
	minute15Usage, _ := strconv.Atoi(usages[0])
	dayUsage, _ := strconv.Atoi(usages[1])

	if dayUsage >= dayLimit {
		return fmt.Errorf("exceeded daily rate limit")
	}
	if minute15Usage >= minute15Limit {
		log.Print("Exceeded 15-minute rate limit, sleeping for 15 minutes")
		time.Sleep(15 * time.Minute)
	}
	return nil
}

func saveActivity(ds data.Store, client *swagger.APIClient, auth context.Context, activity swagger.SummaryActivity) (bool, error) {
	log.Printf("Checking if activity %d is new", activity.Id)
	if existingActivity, err := ds.GetStravaActivity(activity.Id); err != nil {
		return false, err
	} else if existingActivity != nil {
		log.Print("Already saved this activity")
		return false, nil
	}

	log.Print("Fetching activity stream")
	stream, r, err := client.StreamsApi.GetActivityStreams(auth, activity.Id, []string{"latlng"}, false)
	if r.StatusCode == 429 {
		if err = checkRateLimitUsage(r.Header); err != nil {
			return false, err
		}
		return saveActivity(ds, client, auth, activity)
	}
	if err != nil {
		return false, err
	}
	if r.StatusCode != 200 {
		return false, fmt.Errorf("unexpected status code when fetching activity stream: %d", r.StatusCode)
	}

	log.Print("Saving GeoJSON result")
	coordinates := [][][]float64{{}}
	for _, latlng := range stream.Latlng.Data {
		coordinates[0] = append(coordinates[0], []float64{latlng[1], latlng[0]})
	}

	err = saveGeoJSONResult(coordinates, strconv.FormatInt(activity.Id, 10))
	if err != nil {
		return false, err
	}

	log.Print("Saving activity")
	newActivity := &data.StravaActivity{
		ID:                 activity.Id,
		Name:               activity.Name,
		Distance:           activity.Distance,
		MovingTime:         activity.MovingTime,
		ElapsedTime:        activity.ElapsedTime,
		TotalElevationGain: activity.TotalElevationGain,
		ElevHigh:           activity.ElevHigh,
		ElevLow:            activity.ElevLow,
		SportType:          string(*activity.SportType),
		StartDate:          activity.StartDate,
		StartDateLocal:     activity.StartDateLocal,
		Timezone:           activity.Timezone,
		AverageSpeed:       activity.AverageSpeed,
		MaxSpeed:           activity.MaxSpeed,
		GearId:             activity.GearId,
		Kilojoules:         activity.Kilojoules,
		AverageWatts:       activity.AverageWatts,
	}
	if err = ds.CreateStravaActivity(newActivity); err != nil {
		return false, err
	}

	return true, nil
}

type geoJSONResult struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}
