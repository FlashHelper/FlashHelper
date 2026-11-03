# Stage 1: Build Go application
FROM golang:1.26.10-alpine AS build

WORKDIR /src

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Install Swagger generator
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.6

# Copy source code
COPY . .

# Generate Swagger documentation
RUN swag init -g main.go --dir ./cmd/server,./internal --output ./docs

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -o /out/server \
    ./cmd/server

# Stage 2: Run application
FROM alpine:3.23

RUN apk add --no-cache ca-certificates tzdata \
    && adduser -D -u 10001 app

WORKDIR /app

COPY --from=build --chown=app:app /out/server ./server

USER app

EXPOSE 8080

CMD ["./server"]