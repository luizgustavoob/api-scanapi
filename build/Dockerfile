FROM golang:1.16.0-buster AS base
WORKDIR $GOPATH/src/gitlab.neoway.com.br/companies

# Dependencies
FROM base as dependencies
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download

# Test
FROM dependencies AS test
COPY . .
ARG POSTGRES_URL
RUN POSTGRES_URL=$POSTGRES_URL go test -v -cpu 1 -failfast -coverprofile=coverage.out -covermode=set ./...
RUN go get -u github.com/axw/gocov/gocov && GO111MODULE=off go get -u gopkg.in/matm/v1/gocov-html
RUN gocov convert coverage.out | gocov-html > /index.html
RUN grep -v "_mock" coverage.out >> filtered_coverage.out
RUN go tool cover -func filtered_coverage.out

# Build
FROM dependencies AS build
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/companies ./cmd/server

# Release
FROM registry.neoway.com.br/platform/docker-base-image:stable AS image
COPY --from=build /go/bin/companies /companies
ENTRYPOINT ["/companies"]
EXPOSE 9998