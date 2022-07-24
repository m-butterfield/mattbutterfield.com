package tasks

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
)

func router() (*gin.Engine, error) {
	r, err := lib.BaseRouter()
	if err != nil {
		return nil, err
	}

	r.POST("/save_song", saveSong)
	r.POST("/save_image", saveImage)

	return r, nil
}
