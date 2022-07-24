package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"github.com/m-butterfield/mattbutterfield.com/app/static"
	"io/fs"
	"net/http"
)

func blogEntry(c *gin.Context) {
	image, err := ds.GetRandomImage()
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	entryName := c.Param("entryName")
	entryPath := fmt.Sprintf("blog/%s", entryName)
	ffs := &static.FS{}
	if list, err := fs.Glob(ffs, entryPath); err != nil {
		lib.InternalError(err, c)
		return
	} else if len(list) == 0 {
		c.String(http.StatusNotFound, "not found")
		return
	}

	body, err := templateRender(entryPath, makeSingleImagePage(image))
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	c.Render(200, body)
}
