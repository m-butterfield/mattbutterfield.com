package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"log"
	"net/http"
)

func saveSong(c *gin.Context) {
	body := &lib.SaveSongRequest{}
	err := c.Bind(body)
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	if task, err := tc.CreateTask("save_song", "save-song-uploads", body); err != nil {
		lib.InternalError(err, c)
		return
	} else {
		log.Println("Created task: " + task.Name)
	}
	c.Status(http.StatusCreated)
}
