package controllers

import (
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"github.com/m-butterfield/mattbutterfield.com/app/static"
	"html/template"
	"net/http"
)

var musicTemplatePath = append([]string{templatePath + "music/index.gohtml"}, baseTemplatePaths...)

func Music(w http.ResponseWriter, _ *http.Request) {
	songs, err := db.GetSongs()
	if err != nil {
		lib.InternalError(err, w)
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
		lib.InternalError(err, w)
		return
	} else if err = tmpl.Execute(w, struct {
		*basePage
		MusicBaseURL string
		Songs        []*data.Song
	}{
		basePage:     makeBasePage(),
		MusicBaseURL: lib.MusicBaseURL,
		Songs:        songs,
	}); err != nil {
		lib.InternalError(err, w)
		return
	}
}
