FROM golang:latest AS builder

WORKDIR /app

COPY common/ common/
COPY services/user/ services/user/
COPY cmd/user_service_main.go cmd/user_service_main.go

RUN go work init ./common ./services/user
RUN go mod download

RUN CGO_ENABLED=1 GOOS=linux go build -o user-service ./cmd/user_service_main.go

FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
    librdkafka1 \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/user-service .
COPY common/database/migrations/user-service common/database/migrations/user-service

EXPOSE 8080