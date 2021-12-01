package controllers

import (
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"google.golang.org/genproto/googleapis/cloud/tasks/v2"
)

var testRouter = Router()

type testStore struct {
	getImage       func(string) (*data.Image, error)
	getRandomImage func() (*data.Image, error)
	getSongs       func() ([]*data.Song, error)
	saveSong       func(string, string) error
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

func (s *testStore) SaveSong(string, string) error {
	panic("do not call this from controllers")
}

type testTaskCreator struct {
	createTask func(string, string, interface{}) (*tasks.Task, error)
}

func (t *testTaskCreator) CreateTask(taskName, queueID string, body interface{}) (*tasks.Task, error) {
	return t.createTask(taskName, queueID, body)
}
