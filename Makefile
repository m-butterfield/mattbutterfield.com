build:
	go build -i -o bin/website website.go
	go build -i -o bin/scrapes3 scrapes3.go

db:
	sqlite3 app.db ".read schema.sql"

fmt:
	@gofmt -l -s -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")

test:
	sqlite3 app_test.db ".read schema.sql"
	go test -v ./datastore ./website
	rm app_test.db
