package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"net/http"
)

func saveImage(c *gin.Context) {
	body := &lib.SaveImageRequest{}
	err := c.Bind(body)
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	if _, err := tc.CreateTask("save_image", "save-image-uploads", body); err != nil {
		lib.InternalError(err, c)
		return
	}
	c.Status(http.StatusCreated)
}
