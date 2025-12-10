# Builder stage
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache \
    gcc \
    musl-dev \
    pkgconfig \
    make

RUN go install github.com/swaggo/swag/cmd/swag@latest

ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

RUN go clean -modcache && \
    go mod download -x

COPY . .

RUN swag init

RUN go build -a -ldflags="-s -w" -o server

# Runtime stage
FROM alpine:3.20

RUN apk add --no-cache \
    libwebp

WORKDIR /app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /app/server /app/server
COPY --from=builder /app/docs /app/docs
COPY --from=builder /app/.env.example /app/.env

RUN chmod +x /app/server
RUN chown -R appuser:appgroup /app

USER appuser

ENTRYPOINT ["/app/server"]
CMD []

