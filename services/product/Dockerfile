FROM golang:latest AS builder

WORKDIR /app

COPY common/ common/
COPY services/product/ services/product/
COPY cmd/product_service_main.go cmd/product_service_main.go

RUN go work init ./common ./services/product
RUN go mod download

RUN CGO_ENABLED=1 GOOS=linux go build -o product-service ./cmd/product_service_main.go

FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
    librdkafka1 \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/product-service .
COPY common/database/migrations/product-service/ common/database/migrations/product-service/

EXPOSE 8080