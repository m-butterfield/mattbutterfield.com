package data

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestGetImage(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	image := &Image{
		ID:     "test.jpg",
		Width:  100,
		Height: 100,
	}
	if err = s.SaveImage(image); err != nil {
		t.Fatal(err)
	}

	result, err := s.GetImage(image.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, *image, *result)

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Image{})
}

func TestGetRandomImage(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	image := &Image{
		ID:     "test.jpg",
		Width:  100,
		Height: 100,
	}
	if err = s.SaveImage(image); err != nil {
		t.Fatal(err)
	}

	result, err := s.GetRandomImage()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, *image, *result)

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Image{})
}
