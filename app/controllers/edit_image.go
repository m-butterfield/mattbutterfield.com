package controllers

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
)

type editImagePage struct {
	*basePage
	*imageInfo
	AllTags    []*data.Tag
	ImageTypes []data.ImageTypeName
	TagsInput  string
	EditCamera string
	EditLens   string
	EditFilm   string
	EditDate   string
}

type updateImageRequest struct {
	ImageID     string   `json:"imageID"`
	Caption     string   `json:"caption"`
	Location    string   `json:"location"`
	CreatedDate string   `json:"createdDate"`
	Camera      string   `json:"camera"`
	Lens        string   `json:"lens"`
	Film        string   `json:"film"`
	Tags        []string `json:"tags"`
}

func editImage(c *gin.Context) {
	encodedID := c.Param("id")
	id, err := decodeImageID(encodedID)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid image id")
		return
	}

	image, err := ds.GetImage(id)
	if err == sql.ErrNoRows {
		c.String(http.StatusNotFound, "not found")
		return
	}
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	allTags, err := ds.GetAllTags()
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	var tagNames []string
	for _, t := range image.Tags {
		tagNames = append(tagNames, t.Name)
	}

	body, err := templateRender("admin/edit_image", &editImagePage{
		basePage:  makeBasePage(c),
		imageInfo: getImageInfo(image),
		AllTags:   allTags,
		ImageTypes: []data.ImageTypeName{
			data.PhotoADayImageType,
		},
		TagsInput:  strings.Join(tagNames, ", "),
		EditCamera: image.Camera,
		EditLens:   image.Lens,
		EditFilm:   image.Film,
		EditDate:   image.CreatedAt.Format("2006-01-02"),
	})
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	c.Render(200, body)
}

type deleteImageRequest struct {
	ImageID string `json:"imageID"`
}

func deleteImage(c *gin.Context) {
	req := &deleteImageRequest{}
	if err := c.Bind(req); err != nil {
		lib.InternalError(err, c)
		return
	}

	id, err := decodeImageID(req.ImageID)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid image id")
		return
	}

	if err := ds.DeleteImage(id); err != nil {
		lib.InternalError(err, c)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"redirect": "/photos"})
}

func updateImage(c *gin.Context) {
	req := &updateImageRequest{}
	if err := c.Bind(req); err != nil {
		lib.InternalError(err, c)
		return
	}

	id, err := decodeImageID(req.ImageID)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid image id")
		return
	}

	image, err := ds.GetImage(id)
	if err == sql.ErrNoRows {
		c.String(http.StatusNotFound, "not found")
		return
	}
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	image.Caption = req.Caption
	image.Location = req.Location
	image.Camera = req.Camera
	image.Lens = req.Lens
	image.Film = req.Film

	if req.CreatedDate != "" {
		t, err := time.Parse("2006-01-02", req.CreatedDate)
		if err != nil {
			c.String(http.StatusBadRequest, "invalid date")
			return
		}
		image.CreatedAt = t
	}

	var tags []data.Tag
	for _, tagName := range req.Tags {
		if tagName != "" {
			tags = append(tags, data.Tag{Name: tagName})
		}
	}
	image.Tags = tags

	if err := ds.UpdateImage(image); err != nil {
		lib.InternalError(err, c)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"redirect": makeImagePath(id)})
}
