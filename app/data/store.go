package data

import (
	"time"
)

type Store interface {
	GetImage(string) (*Image, error)
	GetImages(time.Time, int) ([]*Image, error)
	GetRandomImage() (*Image, error)
	GetSongs() ([]*Song, error)
	SaveSong(*Song) error
	SaveImage(*Image) error
}

func Connect() (Store, error) {
	s, err := getDS()
	if err != nil {
		return nil, err
	}
	return s, nil
}
