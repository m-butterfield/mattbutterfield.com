package controllers

import (
	"cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTagImages(t *testing.T) {
	ds = &testStore{
		getTags: func() ([]*data.Tag, error) {
			return []*data.Tag{{Name: "Travel", Slug: "travel"}}, nil
		},
		getImagesByTag: func(slug string, before time.Time, limit int) ([]*data.Image, error) {
			return []*data.Image{
				{ID: "test.jpg", Width: 100, Height: 100},
			}, nil
		},
	}
	tc = &testTaskCreator{
		createTask: func(string, string, interface{}) (*cloudtaskspb.Task, error) {
			return &cloudtaskspb.Task{}, nil
		},
	}

	r, err := http.NewRequest(http.MethodGet, "/tag/travel", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	testRouter().ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}

func TestTagImagesNotFound(t *testing.T) {
	ds = &testStore{
		getTags: func() ([]*data.Tag, error) {
			return []*data.Tag{}, nil
		},
	}
	tc = &testTaskCreator{
		createTask: func(string, string, interface{}) (*cloudtaskspb.Task, error) {
			return &cloudtaskspb.Task{}, nil
		},
	}

	r, err := http.NewRequest(http.MethodGet, "/tag/nonexistent", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	testRouter().ServeHTTP(w, r)

	if w.Code != http.StatusNotFound {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}
