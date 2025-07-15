# syntax=docker/dockerfile:1
# BULD STAGE
FROM golang:1.24 AS build-stage

WORKDIR /app

RUN apt-get update && apt-get install -y git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o assistant-server ./app/main.go

# TEST STAGE
FROM build-stage AS run-test-stage
RUN go test -v ./...

# RUN STAGE
FROM alpine AS run-stage

WORKDIR /app

COPY --from=build-stage /app/assistant-server .
EXPOSE 8080

ENTRYPOINT [ "./assistant-server" ]