all:
	@make db-up || echo "Failure start DB"
	@sleep 1
	@make migrate-up
	@go run cmd/main.go


migrate-up:
	migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path migration/ up

migrate-down:
	migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path migration/ down

db-up:
	docker run --rm \
    --detach \
    --publish 5432:5432 \
    --env POSTGRES_DB=postgres \
    --env POSTGRES_USER=postgres \
    --env POSTGRES_PASSWORD=postgres \
    postgres

test:
	@go test -v -cover ./...

db-down:

