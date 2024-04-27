package data

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestGetStravaActivity(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	activity := &StravaActivity{
		ID: 1234,
	}
	if err = s.CreateStravaActivity(activity); err != nil {
		t.Fatal(err)
	}

	result, err := s.GetStravaActivity(activity.ID)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, activity.ID, result.ID)

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&StravaActivity{})
}

func TestGetStravaActivityNotFound(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}

	result, err := s.GetStravaActivity(1234)
	if err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, result)

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&StravaActivity{})
}

func TestGetLatestStravaActivity(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	activity := &StravaActivity{
		ID: 1234,
	}
	if err = s.CreateStravaActivity(activity); err != nil {
		t.Fatal(err)
	}

	result, err := s.GetLatestStravaActivity()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, activity.ID, result.ID)

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&StravaActivity{})
}

func TestGetLatestStravaActivityNotFound(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}

	result, err := s.GetLatestStravaActivity()
	if err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, result)

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&StravaActivity{})
}

func TestGetStravaActivities(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	activity := &StravaActivity{
		ID: 1234,
	}
	if err = s.CreateStravaActivity(activity); err != nil {
		t.Fatal(err)
	}

	results, err := s.GetStravaActivities()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, activity.ID, results[0].ID)

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&StravaActivity{})
}
