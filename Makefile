all:
	@go run cmd/main.go

test:
	@go test -v -cover ./...

.PHONY: swag
swag:
	@swag init -g internal/app/blog.go

db-down:

