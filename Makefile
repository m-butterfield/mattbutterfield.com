cloudrunbasecommand := gcloud run deploy --project=mattbutterfield --region=us-central1 --platform=managed
deployservercommand := $(cloudrunbasecommand) --image=gcr.io/mattbutterfield/mattbutterfield.com mattbutterfield
deployworkercommand := $(cloudrunbasecommand) --image=gcr.io/mattbutterfield/mattbutterfield.com-worker mattbutterfield-worker
gobuild := go build

build:
	$(gobuild) -o bin/server cmd/server.go
	$(gobuild) -o bin/worker cmd/worker.go

deploy: docker-build docker-push
	$(deployservercommand)
	$(deployworkercommand)

deploy-server: docker-build-server docker-push-server
	$(deployservercommand)

deploy-worker: docker-build-worker docker-push-worker
	$(deployworkercommand)

docker-build:
	docker-compose build

docker-build-server:
	docker-compose build server

docker-build-worker:
	docker-compose build worker

docker-push:
	docker-compose push

docker-push-server:
	docker-compose push server

docker-push-worker:
	docker-compose push worker

db:
	createdb mattbutterfield

fmt:
	go fmt ./...
	npx eslint app/static/js/ --fix

run-server:
	DB_SOCKET="host=localhost dbname=mattbutterfield" USE_LOCAL_FS=true go run cmd/server.go

test:
	dropdb --if-exists mattbutterfield_test && createdb mattbutterfield_test && psql -d mattbutterfield_test -f schema.sql
	DB_SOCKET="host=localhost dbname=mattbutterfield_test" go test -v ./app/...

update-deps:
	go get -u ./app/...
	go mod tidy
	npm upgrade
