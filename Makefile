db_migrate:
	go run cmd/bun/main.go -env=test db init
	go run cmd/bun/main.go -env=test db migrate

test:
	go test ./...

fmt:
	gofmt -w -s ./
	goimports -w  -local github.com/go-bun/bun-starter-kit ./
