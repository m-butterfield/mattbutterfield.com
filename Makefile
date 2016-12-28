build:
	go build -i -o bin/website website.go
	go build -i -o bin/scrapes3 scrapes3.go

db:
	sqlite3 app.db ".read schema.sql"

fmt:
	@gofmt -l -s -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")

test:
	go test -v ./datastore ./website

vet:
	@go vet $(shell find . -type f -name '*.go' -path "./")
	@go vet ./datastore ./website
