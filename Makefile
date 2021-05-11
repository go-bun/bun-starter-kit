db_migrate:
	go run cmd/bun/main.go -env=test db init
	go run cmd/bun/main.go -env=test db migrate

test:
	go test ./...
