package data

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestSaveImageWithTags(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	image := &Image{
		ID:     "test_tags.jpg",
		Width:  100,
		Height: 100,
		Tags: []Tag{
			{Name: "Travel", Slug: "travel"},
			{Name: "Urban", Slug: "urban"},
		},
	}
	if err = s.SaveImage(image); err != nil {
		t.Fatal(err)
	}

	result, err := s.GetImage(image.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, image.ID, result.ID)

	tags, err := s.GetImageTags(image.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, tags, 2)

	// cleanup
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Image{})
	s.db.Exec("DELETE FROM image_tags")
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Tag{})
}

func TestGetTags(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	image := &Image{
		ID:     "test_tags_list.jpg",
		Width:  100,
		Height: 100,
		Tags: []Tag{
			{Name: "Travel", Slug: "travel"},
			{Name: "Food", Slug: "food"},
		},
	}
	if err = s.SaveImage(image); err != nil {
		t.Fatal(err)
	}

	tags, err := s.GetTags()
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, tags, 2)

	// cleanup
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Image{})
	s.db.Exec("DELETE FROM image_tags")
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Tag{})
}

func TestGetImagesByTag(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	image1 := &Image{
		ID:     "test_tag1.jpg",
		Width:  100,
		Height: 100,
		Tags: []Tag{
			{Name: "Travel", Slug: "travel"},
		},
	}
	image2 := &Image{
		ID:     "test_tag2.jpg",
		Width:  200,
		Height: 200,
		Tags: []Tag{
			{Name: "Travel", Slug: "travel"},
		},
	}
	if err = s.SaveImage(image1); err != nil {
		t.Fatal(err)
	}
	if err = s.SaveImage(image2); err != nil {
		t.Fatal(err)
	}

	images, err := s.GetImagesByTag("travel", time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC), 10)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, images, 2)

	// cleanup
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Image{})
	s.db.Exec("DELETE FROM image_tags")
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Tag{})
}

func TestAddImageTag(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	image := &Image{
		ID:     "test_add_tag.jpg",
		Width:  100,
		Height: 100,
	}
	if err = s.SaveImage(image); err != nil {
		t.Fatal(err)
	}

	if err = s.AddImageTag(image.ID, "Night Photography"); err != nil {
		t.Fatal(err)
	}

	tags, err := s.GetImageTags(image.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, tags, 1)
	assert.Equal(t, "night-photography", tags[0].Slug)
	assert.Equal(t, "Night Photography", tags[0].Name)

	// cleanup
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Image{})
	s.db.Exec("DELETE FROM image_tags")
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Tag{})
}

func TestRemoveImageTag(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	image := &Image{
		ID:     "test_remove_tag.jpg",
		Width:  100,
		Height: 100,
		Tags: []Tag{
			{Name: "Travel", Slug: "travel"},
			{Name: "Food", Slug: "food"},
		},
	}
	if err = s.SaveImage(image); err != nil {
		t.Fatal(err)
	}

	if err = s.RemoveImageTag(image.ID, "food"); err != nil {
		t.Fatal(err)
	}

	tags, err := s.GetImageTags(image.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, tags, 1)
	assert.Equal(t, "travel", tags[0].Slug)

	// cleanup
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Image{})
	s.db.Exec("DELETE FROM image_tags")
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Tag{})
}
