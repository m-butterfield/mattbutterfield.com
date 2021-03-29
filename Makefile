build:
	go build -o bin/server server.go

db:
	createdb mattbutterfield && psql -d mattbutterfield -f schema.sql

run:
	DB_SOCKET="host=localhost dbname=mattbutterfield" go run server.go

fmt:
	go fmt ./...

test:
	go test -v ./app/...

vet:
	go vet ./...
