package controllers

import (
	"cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"log"
	"time"
)

func testRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r, err := router()
	if err != nil {
		log.Fatal(err)
	}
	return r
}

type testStore struct {
	getImage       func(string) (*data.Image, error)
	getImages      func(time.Time, int) ([]*data.Image, error)
	getYearImages  func(int, time.Time, int) ([]*data.Image, error)
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

func (s *testStore) GetYearImages(year int, before time.Time, limit int) ([]*data.Image, error) {
	return s.getYearImages(year, before, limit)
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

func (s *testStore) GetStravaAccessToken(string) (*data.StravaAccessToken, error) {
	panic("do not call this from controllers")
}

func (s *testStore) CreateStravaAccessToken(*data.StravaAccessToken) error {
	panic("do not call this from controllers")
}

func (s *testStore) UpdateStravaAccessToken(*data.StravaAccessToken) error {
	panic("do not call this from controllers")
}

func (s *testStore) GetStravaActivity(int64) (*data.StravaActivity, error) {
	panic("do not call this from controllers")
}

func (s *testStore) GetStravaActivities() ([]*data.StravaActivity, error) {
	panic("do not call this from controllers")
}

func (s *testStore) GetLatestStravaActivity() (*data.StravaActivity, error) {
	panic("do not call this from controllers")
}

func (s *testStore) CreateStravaActivity(*data.StravaActivity) error {
	panic("do not call this from controllers")
}

func (s *testStore) GetMapBoxConfig() (*data.MapBoxConfig, error) {
	panic("do not call this from controllers")
}

func (s *testStore) UpdateMapBoxConfig(config *data.MapBoxConfig) error {
	panic("do not call this from controllers")
}

func (s *testStore) CreateMapBoxConfig(config *data.MapBoxConfig) error {
	panic("do not call this from controllers")
}

type testTaskCreator struct {
	createTask func(string, string, interface{}) (*cloudtaskspb.Task, error)
}

func (t *testTaskCreator) CreateTask(taskName, queueID string, body interface{}) (*cloudtaskspb.Task, error) {
	return t.createTask(taskName, queueID, body)
}
