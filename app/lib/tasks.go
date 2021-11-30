package lib

import (
	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2"
	"os"
)

const (
	projectID  = "mattbutterfield"
	locationID = "us-central1"
)

type SaveSongRequest struct {
	FileName    string `json:"fileName"`
	SongName    string `json:"songName"`
	Description string `json:"description"`
}

type TaskCreator interface {
	CreateTask(string, string, interface{}) (*taskspb.Task, error)
}

func NewTaskCreator() (TaskCreator, error) {
	workerBaseURL := os.Getenv("WORKER_BASE_URL")
	if workerBaseURL == "" {
		return nil, errors.New("WORKER_BASE_URL not set")
	}
	serviceAccountEmail := os.Getenv("TASK_SERVICE_ACCOUNT_EMAIL")
	if serviceAccountEmail == "" {
		return nil, errors.New("TASK_SERVICE_ACCOUNT_EMAIL not set")
	}
	return &taskCreator{
		workerBaseURL:       workerBaseURL,
		serviceAccountEmail: serviceAccountEmail,
	}, nil
}

type taskCreator struct {
	workerBaseURL       string
	serviceAccountEmail string
}

func (t *taskCreator) CreateTask(taskName, queueID string, body interface{}) (*taskspb.Task, error) {
	url := t.workerBaseURL + taskName
	ctx := context.Background()
	client, err := cloudtasks.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("NewClient: %v", err)
	}
	defer func(client *cloudtasks.Client) {
		err := client.Close()
		if err != nil {
			log.Print(err.Error())
		}
	}(client)

	req := &taskspb.CreateTaskRequest{
		Parent: fmt.Sprintf("projects/%s/locations/%s/queues/%s", projectID, locationID, queueID),
		Task: &taskspb.Task{
			MessageType: &taskspb.Task_HttpRequest{
				HttpRequest: &taskspb.HttpRequest{
					HttpMethod: taskspb.HttpMethod_POST,
					Url:        url,
					AuthorizationHeader: &taskspb.HttpRequest_OidcToken{
						OidcToken: &taskspb.OidcToken{
							ServiceAccountEmail: t.serviceAccountEmail,
						},
					},
				},
			},
		},
	}

	if message, err := json.Marshal(body); err != nil {
		return nil, err
	} else {
		req.Task.GetHttpRequest().Body = message
	}

	createdTask, err := client.CreateTask(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cloudtasks.CreateTask: %v", err)
	}

	return createdTask, nil
}
