build:
	go build -o bin/server server.go

db:
	sqlite3 app.db ".read schema.sql"

fmt:
	gofmt -l -s -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")

test:
	go test -v ./app/...

vet:
	go vet -n $(shell find . -type f -name '*.go' -maxdepth 1)
	go vet -n ./app/...
