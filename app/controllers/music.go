package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"github.com/m-butterfield/mattbutterfield.com/app/static"
	"html/template"
)

var musicTemplatePath = append([]string{templatePath + "music/index.gohtml"}, baseTemplatePaths...)

func music(c *gin.Context) {
	songs, err := ds.GetSongs()
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	var tmpl *template.Template
	if tmpl, err = template.New("index.gohtml").Funcs(map[string]interface{}{
		"getDataNext": func(songs []*data.Song, i int) string {
			if len(songs)-1 == i {
				return ""
			}
			return songs[i+1].ID
		},
	}).ParseFS(&static.FS{}, musicTemplatePath...); err != nil {
		lib.InternalError(err, c)
		return
	}
	c.Render(200, render.HTML{
		Template: tmpl,
		Data: struct {
			*basePage
			MusicBaseURL string
			Songs        []*data.Song
		}{
			basePage:     makeBasePage(),
			MusicBaseURL: lib.MusicBaseURL,
			Songs:        songs,
		},
	})
}
