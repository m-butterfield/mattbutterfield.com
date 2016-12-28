package website

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/m-butterfield/mattbutterfield.com/datastore"
)

var (
	testRouter = buildRouter()
)

type fakeImageStore struct {
	getImage       func(id string) (*datastore.Image, error)
	getRandomImage func() (*datastore.Image, error)
}

func (store *fakeImageStore) SaveImage(image datastore.Image) error {
	panic("Should not call save during website view tests.")
}

func (store *fakeImageStore) GetImage(id string) (*datastore.Image, error) {
	return store.getImage(id)
}

func (store *fakeImageStore) GetLatestImage() (*datastore.Image, error) {
	panic("should not call get latest image suring website view tests.")
}

func (store *fakeImageStore) GetRandomImage() (*datastore.Image, error) {
	return store.getRandomImage()
}

func TestIndex(t *testing.T) {
	imageID := "20040901_001.jpg"
	randomCalled := 0
	imageStore = &fakeImageStore{
		getRandomImage: func() (*datastore.Image, error) {
			randomCalled += 1
			return &datastore.Image{ID: imageID}, nil
		},
	}

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, r)
	if randomCalled != 1 {
		t.Errorf("Unexpected call count for GetRandomImage(): %d", randomCalled)
	}
	if w.Code != http.StatusFound {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
	if value, ok := w.Header()["Location"]; ok {
		if !strings.HasSuffix(value[0], imagePathBase+encodeImageID(imageID)) {
			t.Errorf("Unexpected redirect location: %s", value)
		}
	} else {
		t.Error("Location header not found in response.")
	}
}

func TestImg(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	imageTemplateName = cwd + "/" + "templates/image.html"

	imageID := "20040901_001.jpg"
	randImageID := "blerp"
	getImageCalled, randomCalled := 0, 0
	imageStore = &fakeImageStore{
		getImage: func(id string) (*datastore.Image, error) {
			getImageCalled += 1
			if id != imageID {
				t.Errorf("GetImage called with unexpected image id: %s", id)
			}
			return &datastore.Image{ID: imageID}, nil
		},
		getRandomImage: func() (*datastore.Image, error) {
			randomCalled += 1
			return &datastore.Image{ID: randImageID}, nil
		},
	}

	r, err := http.NewRequest(http.MethodGet, imagePathBase+encodeImageID(imageID), nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, r)
	if getImageCalled != 1 {
		t.Errorf("Unexpected call count for GetImage(): %d", getImageCalled)
	}
	if randomCalled != 1 {
		t.Errorf("Unexpected call count for GetRandomImage(): %d", randomCalled)
	}
	if w.Code != http.StatusOK {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}

func TestInvalidID(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	imageTemplateName = cwd + "/" + "templates/image.html"

	r, err := http.NewRequest(http.MethodGet, imagePathBase+"MjAwO", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}

func TestImageNotFound(t *testing.T) {
	getImageCalled := 0
	imageStore = &fakeImageStore{
		getImage: func(id string) (*datastore.Image, error) {
			getImageCalled += 1
			return nil, sql.ErrNoRows
		},
	}

	r, err := http.NewRequest(http.MethodGet, imagePathBase+encodeImageID("1234"), nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, r)
	if getImageCalled != 1 {
		t.Errorf("Unexpected call count for GetImage(): %d", getImageCalled)
	}
	if w.Code != http.StatusNotFound {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}

func TestGetImageTimeStr(t *testing.T) {
	expectedTimeStr := "September 2004"
	img := &datastore.Image{ID: "20040901_001.jpg"}
	timeStr := getImageTimeStr(img)
	if timeStr != expectedTimeStr {
		t.Errorf("Unexpected time string: %s != %s", expectedTimeStr, timeStr)
	}
	img.ID = "blerp"
	timeStr = getImageTimeStr(img)
	if timeStr != "" {
		t.Errorf("Expected empty string from id: %s, instead got: %s", img.ID, timeStr)
	}
}
