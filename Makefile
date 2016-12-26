GO_FILES = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

build:
	go build -o bin/website website.go
	go build -o bin/scrapes3 scrapes3.go

db:
	sqlite3 app.db ".read schema.sql"

fmt:
	gofmt -l -s -w ${GO_FILES}
