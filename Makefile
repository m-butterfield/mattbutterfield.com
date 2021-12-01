build:
	go build -o bin/server cmd/server.go
	go build -o bin/worker cmd/worker.go

deploy: deploy-web deploy-worker

deploy-web: docker-build docker-push
	gcloud run deploy mattbutterfield \
	  --project=mattbutterfield \
	  --region=us-central1 \
	  --platform=managed \
	  --image=gcr.io/mattbutterfield/mattbutterfield.com

deploy-worker: docker-build docker-push
	gcloud run deploy mattbutterfield-worker \
	  --project=mattbutterfield \
	  --region=us-central1 \
	  --platform=managed \
	  --image=gcr.io/mattbutterfield/mattbutterfield.com-worker

docker-build:
	docker-compose build

docker-push:
	docker-compose push

db:
	createdb mattbutterfield

fmt:
	go fmt ./...
	npx eslint app/static/js/ --fix

run:
	DB_SOCKET="host=localhost dbname=mattbutterfield" USE_LOCAL_FS=true go run cmd/server.go

test:
	dropdb --if-exists mattbutterfield_test && createdb mattbutterfield_test && psql -d mattbutterfield_test -f schema.sql
	DB_SOCKET="host=localhost dbname=mattbutterfield_test" PUBSUB_EMULATOR_HOST=localhost go test -v ./app/...
