# ── Stage 1: Build ────────────────────────────────────────────────────────────
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /bin/selection-service ./cmd/main.go

# --- migrate CLI ---
FROM migrate/migrate:v4.17.1 AS migrate

# ── Stage 2: Runtime ──────────────────────────────────────────────────────────
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata

COPY --from=migrate /migrate /usr/local/bin/migrate
COPY --from=builder /bin/selection-service /selection-service

COPY internal/migrations /migrations
COPY deploy/entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh /selection-service \
    && adduser -D -u 65532 svc

USER svc

EXPOSE 8080 8081

ENTRYPOINT ["/entrypoint.sh"]
