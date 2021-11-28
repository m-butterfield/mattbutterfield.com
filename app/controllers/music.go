package controllers

import (
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"github.com/m-butterfield/mattbutterfield.com/app/static"
	"html/template"
	"net/http"
	"time"
)

var musicTemplatePath = append([]string{templatePath + "music/index.gohtml"}, baseTemplatePaths...)

func Music(w http.ResponseWriter, _ *http.Request) {
	songs, err := db.GetSongs()
	if err != nil {
		internalError(err, w)
		return
	}
	if tmpl, err := template.New("index.gohtml").Funcs(map[string]interface{}{
		"getDataNext": func(songs []*data.Song, i int) string {
			if len(songs)-1 == i {
				return ""
			}
			return songs[i+1].ID
		},
	}).ParseFS(&static.FlexFS{}, musicTemplatePath...); err != nil {
		internalError(err, w)
		return
	} else if err = tmpl.Execute(w, struct {
		Songs []*data.Song
		Year  string
	}{
		Songs: songs,
		Year:  time.Now().Format("2006"),
	}); err != nil {
		internalError(err, w)
		return
	}
}
