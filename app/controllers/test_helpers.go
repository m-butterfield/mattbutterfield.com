package controllers

import "github.com/m-butterfield/mattbutterfield.com/app/data"

var testRouter = Router()

type testStore struct {
	getImage       func(string) (*data.Image, error)
	getRandomImage func() (*data.Image, error)
}

func (s *testStore) GetImage(id string) (*data.Image, error) {
	return s.getImage(id)
}

func (s *testStore) GetRandomImage() (*data.Image, error) {
	return s.getRandomImage()
}
