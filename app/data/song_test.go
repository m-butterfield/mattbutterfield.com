package data

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestGetSongs(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	song := &Song{
		ID: "test.jpg",
	}
	if err = s.SaveSong(song); err != nil {
		t.Fatal(err)
	}

	result, err := s.GetSongs()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, *song, *result[0])

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Song{})
}
