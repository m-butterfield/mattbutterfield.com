package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"net/http"
)

func saveSong(c *gin.Context) {
	body := &lib.SaveSongRequest{}
	err := c.Bind(body)
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	if _, err := tc.CreateTask("save_song", "save-song-uploads", body); err != nil {
		lib.InternalError(err, c)
		return
	}
	c.Status(http.StatusCreated)
}
