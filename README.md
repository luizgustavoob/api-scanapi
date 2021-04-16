# Companies API

Companies API is an example project, created with the objective of demonstrating the use of [ScanAPI](https://scanapi.dev/), a library for integration tests.

## Getting Started

### Prerequisites

- [Golang](http://golang.org/)
- [GNU Make](https://www.gnu.org/software/make/)
- [Docker](http://docker.com)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Installing and running locally

```shell script
# Install dependencies
make install

# Run postgres locally as a container
make env

# Run server locally
make run
```

## Running the tests and coverage report

To view report of tests locally use the following command:

```bash
make env # prepares environment for testing
make test
```

## Running the integration tests with ScanAPI

To scan the external APIs and view report of integration tests use the following command:

```bash
make scan-external
```

To scan your application and view report of integration tests use the following command:

```bash
make env        # prepares environment for testing
make build      # prepares your application
make image      # prepares your application
make run-docker # run your application as container
make scan-internal
```

## Deployment

### Build

```bash
make build
```

### Create image

```bash
make image
```

### Run registry image locally

```bash
make run-docker
make remove-docker
```
  
## Inspiration

### Package organization

The package structure used in this project was inspired by the [golang-standards](https://github.com/golang-standards/project-layout) project.
