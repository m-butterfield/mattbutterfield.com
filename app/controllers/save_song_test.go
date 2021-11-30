package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaveSong(t *testing.T) {
	expectedBody := &lib.SaveSongRequest{
		FileName:    "test.wav?123456",
		SongName:    "test song!",
		Description: "test description",
	}
	body, err := json.Marshal(expectedBody)
	r, err := http.NewRequest(http.MethodPost, "/admin/save_song", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	r.AddCookie(&http.Cookie{Name: "auth", Value: "1234"})
	authArray = []byte("1234")
	taskCalled := false
	taskCreator = &testTaskCreator{
		createTask: func(taskName, queueID string, body interface{}) (*taskspb.Task, error) {
			taskCalled = true
			if taskName != "save_song" {
				t.Error("Unexpected task name")
			}
			if queueID != "save-song-uploads" {
				t.Error("Unexpected queueID")
			}
			if *body.(*lib.SaveSongRequest) != *expectedBody {
				t.Error("Unexpected task body")
			}
			return &taskspb.Task{}, nil
		},
	}

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
	if !taskCalled {
		t.Errorf("create task never called")
	}
}
