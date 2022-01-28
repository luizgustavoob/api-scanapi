FROM golang:alpine AS base
WORKDIR /go/src/github.com/api-scanapi/

# Dependencies
FROM base as dependencies
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download

# Build
FROM dependencies AS build
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/companies ./cmd/api

# Release
FROM alpine:latest AS image
COPY --from=build /go/bin/companies /companies
ENTRYPOINT ["/companies"]
EXPOSE 9998