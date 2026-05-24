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
			{Name: "Travel"},
			{Name: "Urban"},
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
	assert.Len(t, result.Tags, 2)

	tags, err := s.GetImageTags(image.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, tags, 2)

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Image{})
	s.db.Exec("DELETE FROM image_tags")
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Tag{})
}

func TestSaveImageReusesExistingTag(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	image1 := &Image{
		ID:     "first.jpg",
		Width:  100,
		Height: 100,
		Tags: []Tag{
			{Name: "Travel"},
		},
	}
	if err = s.SaveImage(image1); err != nil {
		t.Fatal(err)
	}

	image2 := &Image{
		ID:     "second.jpg",
		Width:  200,
		Height: 200,
		Tags: []Tag{
			{Name: "Travel"},
		},
	}
	if err = s.SaveImage(image2); err != nil {
		t.Fatal(err)
	}

	tags, err := s.GetImageTags(image2.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, tags, 1)
	assert.Equal(t, "Travel", tags[0].Name)

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
			{Name: "Travel"},
			{Name: "Food"},
		},
	}
	if err = s.SaveImage(image); err != nil {
		t.Fatal(err)
	}

	tags, err := s.GetImageTags(image.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, tags, 2)

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
			{Name: "Travel"},
		},
	}
	image2 := &Image{
		ID:     "test_tag2.jpg",
		Width:  200,
		Height: 200,
		Tags: []Tag{
			{Name: "Travel"},
		},
	}
	if err = s.SaveImage(image1); err != nil {
		t.Fatal(err)
	}
	if err = s.SaveImage(image2); err != nil {
		t.Fatal(err)
	}

	images, err := s.GetImagesByTag([]string{"Travel"}, time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC), 10)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, images, 2)

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Image{})
	s.db.Exec("DELETE FROM image_tags")
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Tag{})
}

func TestGetImagesByMultipleTags(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	image1 := &Image{
		ID:     "test_multi1.jpg",
		Width:  100,
		Height: 100,
		Tags: []Tag{
			{Name: "Travel"},
			{Name: "Food"},
		},
	}
	image2 := &Image{
		ID:     "test_multi2.jpg",
		Width:  200,
		Height: 200,
		Tags: []Tag{
			{Name: "Travel"},
		},
	}
	if err = s.SaveImage(image1); err != nil {
		t.Fatal(err)
	}
	if err = s.SaveImage(image2); err != nil {
		t.Fatal(err)
	}

	images, err := s.GetImagesByTag([]string{"Travel", "Food"}, time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC), 10)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, images, 1)
	assert.Equal(t, "test_multi1.jpg", images[0].ID)

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
	assert.Equal(t, "Night Photography", tags[0].Name)

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Image{})
	s.db.Exec("DELETE FROM image_tags")
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Tag{})
}

func TestUpdateImageTags(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	image := &Image{
		ID:     "test_update_tags.jpg",
		Width:  100,
		Height: 100,
	}

	if err = s.SaveImage(image); err != nil {
		t.Fatal(err)
	}

	image.Caption = "updated"
	image.Tags = []Tag{
		{Name: "Travel"},
		{Name: "Food"},
	}

	if err = s.UpdateImage(image); err != nil {
		t.Fatal(err)
	}

	result, err := s.GetImage(image.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "updated", result.Caption)
	assert.Len(t, result.Tags, 2)

	image.Tags = []Tag{
		{Name: "Travel"},
	}

	if err = s.UpdateImage(image); err != nil {
		t.Fatal(err)
	}

	result, err = s.GetImage(image.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, result.Tags, 1)
	assert.Equal(t, "Travel", result.Tags[0].Name)

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
			{Name: "Travel"},
			{Name: "Food"},
		},
	}
	if err = s.SaveImage(image); err != nil {
		t.Fatal(err)
	}

	if err = s.RemoveImageTag(image.ID, "Food"); err != nil {
		t.Fatal(err)
	}

	tags, err := s.GetImageTags(image.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, tags, 1)
	assert.Equal(t, "Travel", tags[0].Name)

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Image{})
	s.db.Exec("DELETE FROM image_tags")
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Tag{})
}
