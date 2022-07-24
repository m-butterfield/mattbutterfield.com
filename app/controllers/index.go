package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"net/http"
)

func index(c *gin.Context) {
	c.Redirect(http.StatusFound, makeImagePath(lib.HomeImage))
}
