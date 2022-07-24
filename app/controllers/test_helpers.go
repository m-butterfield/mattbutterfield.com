package controllers

import (
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"google.golang.org/genproto/googleapis/cloud/tasks/v2"
	"time"
)

var testRouter = Router()

type testStore struct {
	getImage       func(string) (*data.Image, error)
	getImages      func(time.Time, int) ([]*data.Image, error)
	getRandomImage func() (*data.Image, error)
	getSongs       func() ([]*data.Song, error)
	saveSong       func(*data.Song) error
	saveImage      func(*data.Image) error
}

func (s *testStore) GetImage(id string) (*data.Image, error) {
	return s.getImage(id)
}

func (s *testStore) GetImages(before time.Time, limit int) ([]*data.Image, error) {
	return s.getImages(before, limit)
}

func (s *testStore) GetRandomImage() (*data.Image, error) {
	return s.getRandomImage()
}

func (s *testStore) GetSongs() ([]*data.Song, error) {
	return s.getSongs()
}

func (s *testStore) SaveSong(*data.Song) error {
	panic("do not call this from controllers")
}

func (s *testStore) SaveImage(*data.Image) error {
	panic("do not call this from controllers")
}

type testTaskCreator struct {
	createTask func(string, string, interface{}) (*tasks.Task, error)
}

func (t *testTaskCreator) CreateTask(taskName, queueID string, body interface{}) (*tasks.Task, error) {
	return t.createTask(taskName, queueID, body)
}
