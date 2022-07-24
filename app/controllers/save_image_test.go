package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"google.golang.org/genproto/googleapis/cloud/tasks/v2"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSaveImage(t *testing.T) {
	expectedBody := &lib.SaveImageRequest{
		ImageFileName: "test.jpg?123456",
		CreatedDate:   lib.CreatedDateJSON{Time: time.Date(2021, time.December, 1, 0, 0, 0, 0, time.UTC)},
		Location:      "NYC",
		Caption:       "Central Park",
	}
	body, err := json.Marshal(expectedBody)
	r, err := http.NewRequest(http.MethodPost, "/admin/save_image", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.AddCookie(&http.Cookie{Name: "auth", Value: "1234"})
	authArray = []byte("1234")
	taskCalled := false
	tc = &testTaskCreator{
		createTask: func(taskName, queueID string, body interface{}) (*tasks.Task, error) {
			taskCalled = true
			if taskName != "save_image" {
				t.Error("Unexpected task name: ", taskName)
			}
			if queueID != "save-image-uploads" {
				t.Error("Unexpected queueID")
			}
			if *body.(*lib.SaveImageRequest) != *expectedBody {
				t.Error("Unexpected task body")
			}
			return &tasks.Task{}, nil
		},
	}

	w := httptest.NewRecorder()
	testRouter().ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
	if !taskCalled {
		t.Errorf("create task never called")
	}
}
