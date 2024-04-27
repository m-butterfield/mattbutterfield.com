package controllers

import (
	"bytes"
	"cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStravaWebhook(t *testing.T) {
	body := &StravaWebhookRequest{
		AspectType:     "create",
		EventTime:      1234,
		ObjectId:       12345,
		ObjectType:     "activity",
		OwnerId:        123456,
		SubscriptionId: 1234567,
		Updates:        map[string]any{},
	}
	data, err := json.Marshal(body)
	r, err := http.NewRequest(http.MethodPost, "/strava_webhook", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")
	taskCalled := false
	tc = &testTaskCreator{
		createTask: func(taskName, queueID string, body interface{}) (*cloudtaskspb.Task, error) {
			taskCalled = true
			if taskName != "update_heatmap" {
				t.Error("Unexpected task name")
			}
			if queueID != "update-heatmap" {
				t.Error("Unexpected queueID")
			}
			if body != nil {
				t.Error("Unexpected task body")
			}
			return &cloudtaskspb.Task{}, nil
		},
	}

	w := httptest.NewRecorder()
	testRouter().ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
	if !taskCalled {
		t.Errorf("create task never called")
	}
}
