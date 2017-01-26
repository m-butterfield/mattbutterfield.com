package app

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/m-butterfield/mattbutterfield.com/app/data"
)

var (
	testRouter = buildRouter(true)
)

type fakeImageStore struct {
	getImage          func(string) (*data.Image, error)
	getPrevNextImages func(string) (*data.Image, *data.Image, error)
	getRandomImage    func() (*data.Image, error)
	updateImage       func(string, string, string) error
}

func (store *fakeImageStore) GetImage(id string) (*data.Image, error) {
	return store.getImage(id)
}

func (store *fakeImageStore) GetLatestImage() (*data.Image, error) {
	panic("should not call get latest image during website view tests.")
}

func (store *fakeImageStore) GetPrevNextImages(id string) (*data.Image, *data.Image, error) {
	return store.getPrevNextImages(id)
}

func (store *fakeImageStore) GetRandomImage() (*data.Image, error) {
	return store.getRandomImage()
}

func (store *fakeImageStore) SaveImage(image data.Image) error {
	panic("Should not call save during website view tests.")
}

func (store *fakeImageStore) UpdateImage(id, location, caption string) error {
	return store.updateImage(id, location, caption)
}

func TestGetImageTimeStr(t *testing.T) {
	expectedTimeStr := "September 2004"
	img := &data.Image{ID: "20040901_001.jpg"}
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

func TestIndex(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, r)
	if w.Code != http.StatusFound {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
	if value, ok := w.Header()["Location"]; ok {
		if !strings.HasSuffix(value[0], imagePathBase+encodeImageID(homeImage)) {
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
		getImage: func(id string) (*data.Image, error) {
			getImageCalled += 1
			if id != imageID {
				t.Errorf("GetImage called with unexpected image id: %s", id)
			}
			return &data.Image{ID: imageID}, nil
		},
		getRandomImage: func() (*data.Image, error) {
			randomCalled += 1
			return &data.Image{ID: randImageID}, nil
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
	if w.Code != http.StatusBadRequest {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}

func TestImageNotFound(t *testing.T) {
	getImageCalled := 0
	imageStore = &fakeImageStore{
		getImage: func(id string) (*data.Image, error) {
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

func TestAdmin(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	adminTemplateName = cwd + "/" + "templates/admin.html"

	getImageCalled := 0
	getPrevNextCalled := 0
	imageID := "20040901_001.jpg"
	imageStore = &fakeImageStore{
		getImage: func(id string) (*data.Image, error) {
			getImageCalled += 1
			if id != imageID {
				t.Errorf("GetImage called with unexpected image id: %s", id)
			}
			return &data.Image{ID: imageID}, nil
		},
		getPrevNextImages: func(id string) (*data.Image, *data.Image, error) {
			getPrevNextCalled += 1
			if id != imageID {
				t.Errorf("GetPrevNextImages called with unexpected image id: %s", id)
			}
			return &data.Image{ID: imageID}, &data.Image{ID: imageID}, nil
		},
	}

	r, err := http.NewRequest(http.MethodGet, adminPathBase+encodeImageID(imageID), nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, r)
	if getImageCalled != 1 {
		t.Errorf("Unexpected call count for GetImage(): %d", getImageCalled)
	}
	if getPrevNextCalled != 1 {
		t.Errorf("Unexpected call count for GetPrevNextImages(): %d", getPrevNextCalled)
	}
	if w.Code != http.StatusOK {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}

func TestAdminImageNotFound(t *testing.T) {
	getImageCalled := 0
	imageStore = &fakeImageStore{
		getImage: func(id string) (*data.Image, error) {
			getImageCalled += 1
			return nil, sql.ErrNoRows
		},
	}

	r, err := http.NewRequest(http.MethodGet, adminPathBase+encodeImageID("1234"), nil)
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

func TestAdminInvalidID(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	imageTemplateName = cwd + "/" + "templates/admin.html"

	r, err := http.NewRequest(http.MethodGet, adminPathBase+"MjAwO", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}
