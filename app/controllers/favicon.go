package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"net/http"
)

func favicon(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, lib.ImagesBaseURL+"/favicon.ico")
}
