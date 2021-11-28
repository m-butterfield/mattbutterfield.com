build:
	go build -o bin/server server.go

db:
	createdb mattbutterfield

run:
	DB_SOCKET="host=localhost dbname=mattbutterfield" USE_LOCAL_FS=true go run server.go

fmt:
	go fmt ./...
	npx eslint app/static/js/ --fix

test:
	dropdb --if-exists mattbutterfield_test && createdb mattbutterfield_test && psql -d mattbutterfield_test -f schema.sql
	DB_SOCKET="host=localhost dbname=mattbutterfield_test" PUBSUB_EMULATOR_HOST=localhost go test -v ./app/...
