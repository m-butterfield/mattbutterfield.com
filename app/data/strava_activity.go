package data

import (
	"gorm.io/gorm"
	"time"
)

type StravaActivity struct {
	ID                 int64
	Name               string
	Distance           float32 // in meters
	MovingTime         int32   // in seconds
	ElapsedTime        int32   // in seconds
	TotalElevationGain float32 // in meters
	ElevHigh           float32 // in meters
	ElevLow            float32 // in meters
	SportType          string
	StartDate          time.Time
	StartDateLocal     time.Time
	Timezone           string
	AverageSpeed       float32 // in meters per second
	MaxSpeed           float32 // in meters per second
	GearId             string
	Kilojoules         float32 // total work done
	AverageWatts       float32
}

func (s *ds) GetStravaActivity(id int64) (*StravaActivity, error) {
	activity := &StravaActivity{}
	tx := s.db.First(activity, "id = ?", id)
	if tx.Error != nil {
		if tx.Error != gorm.ErrRecordNotFound {
			return nil, tx.Error
		}
		return nil, nil
	}
	return activity, nil
}

func (s *ds) GetLatestStravaActivity() (*StravaActivity, error) {
	activity := &StravaActivity{}
	tx := s.db.Order("start_date desc").First(activity)
	if tx.Error != nil {
		if tx.Error != gorm.ErrRecordNotFound {
			return nil, tx.Error
		}
		return nil, nil
	}
	return activity, nil
}

func (s *ds) GetStravaActivities() ([]*StravaActivity, error) {
	var activities []*StravaActivity
	tx := s.db.Find(&activities)
	if tx.Error != nil {
		return nil, nil
	}
	return activities, nil
}

func (s *ds) CreateStravaActivity(activity *StravaActivity) error {
	if tx := s.db.Create(activity); tx.Error != nil {
		return tx.Error
	}
	return nil
}
