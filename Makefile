build:
	go build -o bin/website website.go
	go build -o bin/scrapes3 scrapes3.go

db:
	sqlite3 app.db ".read schema.sql"
