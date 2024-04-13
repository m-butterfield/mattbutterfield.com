package data

import (
	"time"
)

type Store interface {
	GetImage(string) (*Image, error)
	GetImages(time.Time, int) ([]*Image, error)
	GetYearImages(int, time.Time, int) ([]*Image, error)
	GetRandomImage() (*Image, error)
	GetSongs() ([]*Song, error)
	GetStravaAccessToken(string) (*StravaAccessToken, error)
	SaveSong(*Song) error
	SaveImage(*Image) error
	CreateStravaAccessToken(*StravaAccessToken) error
	UpdateStravaAccessToken(*StravaAccessToken) error
	GetStravaActivity(int64) (*StravaActivity, error)
	GetStravaActivities() ([]*StravaActivity, error)
	GetLatestStravaActivity() (*StravaActivity, error)
	CreateStravaActivity(*StravaActivity) error
	GetMapBoxConfig() (*MapBoxConfig, error)
	UpdateMapBoxConfig(config *MapBoxConfig) error
	CreateMapBoxConfig(config *MapBoxConfig) error
}

func Connect() (Store, error) {
	s, err := getDS()
	if err != nil {
		return nil, err
	}
	return s, nil
}
