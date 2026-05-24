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
		getTagsByNames: func(names []string) ([]*data.Tag, error) {
			return []*data.Tag{{Name: "Travel"}}, nil
		},
		getImagesByTag: func(names []string, before time.Time, limit int) ([]*data.Image, error) {
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

	r, err := http.NewRequest(http.MethodGet, "/tag/Travel", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	testRouter().ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}

func TestTagImagesMulti(t *testing.T) {
	ds = &testStore{
		getTagsByNames: func(names []string) ([]*data.Tag, error) {
			return []*data.Tag{
				{Name: "Travel"},
				{Name: "Food"},
			}, nil
		},
		getImagesByTag: func(names []string, before time.Time, limit int) ([]*data.Image, error) {
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

	r, err := http.NewRequest(http.MethodGet, "/tag/Travel,Food", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	testRouter().ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}

func TestTagImagesEmpty(t *testing.T) {
	ds = &testStore{
		getTagsByNames: func(names []string) ([]*data.Tag, error) {
			return []*data.Tag{}, nil
		},
		getImagesByTag: func(names []string, before time.Time, limit int) ([]*data.Image, error) {
			return []*data.Image{}, nil
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

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}
