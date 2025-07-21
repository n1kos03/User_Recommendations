FROM golang:latest

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY internal/database/ internal/database/
COPY internal/kafka/ internal/kafka/
COPY internal/models/ internal/models/ 
COPY internal/user/ internal/user/

RUN go build -o user-service ./cmd/main.go

# FROM debian:bullseye-slim

# WORKDIR /root/
# COPY --from=builder /app/user-service .

EXPOSE 8080

# CMD ["./user-service"]