.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: migrate_up
migrate_up:
	migrate -path migrations -database "postgres://localhost/restapi_dev?sslmode=disable" up

.PHONY: migrate_down
migrate_down:
	migrate -path migrations -database "postgres://localhost/restapi_dev?sslmode=disable" down


.DEFAULT_GOAL := build