build:
	go build -o bin/server server.go

db:
	createdb mattbutterfield && psql -d mattbutterfield -f schema.sql
	createdb mattbutterfield_test && psql -d mattbutterfield_test -f schema.sql

run:
	DB_SOCKET="host=localhost dbname=mattbutterfield" go run server.go

fmt:
	go fmt ./...

test:
	DB_SOCKET="host=localhost dbname=mattbutterfield_test" go test -v ./app/...

vet:
	go vet ./...
