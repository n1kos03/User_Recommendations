FROM golang:latest

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY cmd/user_service_main.go cmd/user_service_main.go
COPY internal/database/ internal/database/
COPY internal/kafka/producer.go internal/kafka/producer.go
COPY internal/models/user.go internal/models/user.go
COPY internal/models/kafka_events.go internal/models/kafka_events.go
COPY internal/user/ internal/user/

RUN go build -o user-service ./cmd/user_service_main.go

EXPOSE 8080