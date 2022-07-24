package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"github.com/m-butterfield/mattbutterfield.com/app/static"
	"net/http"
	"path"
)

func router() (*gin.Engine, error) {
	r, err := lib.BaseRouter()
	if err != nil {
		return nil, err
	}

	fs := http.FileServer(http.FS(static.FS{}))
	addStaticHandler(r, "/css", fs)
	addStaticHandler(r, "/js", fs)

	r.GET("/favicon.ico", favicon)
	r.GET("/", index)
	r.GET("/login", login)
	r.POST("/login", loginUser)
	r.GET("/img/:id", home)
	r.GET("/blog", blog)
	r.GET("/blog/:entryName", blogEntry)
	r.GET("/music", music)
	r.GET("/photos", photos)
	r.GET("/video", video)
	r.GET("/video/connections", videoConnections)

	adminGroup := r.Group("/adminGroup")
	adminGroup.Use(authRequired)
	adminGroup.GET("/adminGroup", admin)
	adminGroup.GET("/adminGroup/upload_music", uploadMusic)
	adminGroup.GET("/adminGroup/upload_image", uploadImage)
	adminGroup.POST("/adminGroup/signed_upload_url", signedUploadURL)
	adminGroup.POST("/adminGroup/save_song", saveSong)
	adminGroup.POST("/adminGroup/song_image", saveImage)

	return r, nil
}

func addStaticHandler(r *gin.Engine, prefix string, fileServer http.Handler) {
	handler := func(c *gin.Context) { fileServer.ServeHTTP(c.Writer, c.Request) }
	pattern := path.Join(prefix, "/*filepath")
	r.GET(pattern, handler)
	r.HEAD(pattern, handler)
}
