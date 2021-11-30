package controllers

import (
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2"
)

var testRouter = Router()

type testStore struct {
	getImage       func(string) (*data.Image, error)
	getRandomImage func() (*data.Image, error)
	getSongs       func() ([]*data.Song, error)
}

func (s *testStore) GetImage(id string) (*data.Image, error) {
	return s.getImage(id)
}

func (s *testStore) GetRandomImage() (*data.Image, error) {
	return s.getRandomImage()
}

func (s *testStore) GetSongs() ([]*data.Song, error) {
	return s.getSongs()
}

type testTaskCreator struct {
	createTask func(string, string, interface{}) (*taskspb.Task, error)
}

func (t *testTaskCreator) CreateTask(taskName, queueID string, body interface{}) (*taskspb.Task, error) {
	return t.createTask(taskName, queueID, body)
}
