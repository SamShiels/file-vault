# syntax=docker/dockerfile:1

FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum .env ./
RUN go mod download

COPY . .
COPY public/ ./public/

RUN go build -o vault

# Final image to run the app with debugging support
FROM golang:1.21

WORKDIR /app

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/vault .

# Install Delve debugger
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Expose the port used by Delve
EXPOSE 40000
EXPOSE 8082

CMD ["dlv", "exec", "./vault", "--headless", "--listen", "0.0.0.0:40000", "--api-version", "2", "--accept-multiclient"]