package static

import (
	"embed"
	"errors"
	"io/fs"
	"os"
	"strings"
)

var (
	//go:embed templates
	templatesEmbedFS embed.FS
	//go:embed css
	cssEmbedFS embed.FS
	//go:embed js
	jsEmbedFS embed.FS
)

type FlexFS struct{}

func (f *FlexFS) Open(name string) (fs.File, error) {
	if os.Getenv("USE_LOCAL_FS") != "" {
		return os.Open("./app/static/" + name)
	}
	if strings.HasPrefix(name, "js/") {
		return jsEmbedFS.Open(name)
	}
	if strings.HasPrefix(name, "css/") {
		return cssEmbedFS.Open(name)
	}
	if strings.HasPrefix(name, "templates/") {
		return templatesEmbedFS.Open(name)
	}
	return nil, errors.New("could not find file")
}
