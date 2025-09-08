ALL_SOURCES=services/user/handlers/*.go common/database/*.go common/kafka/*.go cmd/*.go common/models/*.go services/product/handlers/*.go services/recommendations/handlers/*.go

all: run-services

run-services: #lint
	docker-compose up --build -d

stop-services:
	docker-compose down -v

lint: fmt
	golangci-lint run --disable=depguard ./...

fmt:
	gofmt -w $(ALL_SOURCES)