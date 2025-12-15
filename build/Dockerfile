FROM golang:1.25.5-alpine AS builder

RUN apk add --no-cache build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o application ./cmd

FROM alpine:3.23.0 AS release

RUN apk add --no-cache

WORKDIR /app
COPY --from=builder /app/application /app/application
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
ENV ZONEINFO=/zoneinfo.zip

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

CMD ["/app/application"]
