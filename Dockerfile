# syntax=docker/dockerfile:1

FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum .env ./
RUN go mod download

COPY *.go ./
COPY public/ ./public/

RUN CGO_ENABLED=0 GOOS=linux go build -o vault

EXPOSE 8082

CMD [ "./vault" ]