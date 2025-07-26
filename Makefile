SOURCES=internal/user/handlers/*.go internal/database/*.go internal/kafka/*.go cmd/*.go

all: run-services

run-services: lint
	docker-compose up --build -d

stop-services:
	docker-compose down -v

lint: fmt
	golangci-lint run --disable=depguard ./...

fmt:
	gofmt -w $(SOURCES)