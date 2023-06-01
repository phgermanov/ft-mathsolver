# FT-Mathsolver Service

FT-Mathsolver is a service designed to parse and evaluate simple arithmetic expressions provided in natural language format.

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

To run this service, you will need to have Golang installed on your machine. You can download it [here](https://golang.org/dl/).

### Installing

Clone the repository:

```bash
git clone https://github.com/phgermanov/ft-mathsolver.git
```

Navigate into the directory:

```bash
cd ft-mathsolver
```
Run the application:

```bash
go run cmd/mathsolver/mathsolver.go
```

or 

```bash
docker-compose up
```
### Running the tests

To run the unit tests, execute the following command:
```bash
go test ./internal/...
```
To run integration tests, a service needs to be running locally on port *8080* before running the following command:
```bash
go test ./integration/...
```
### API Endpoints

* `POST /evaluate`: Evaluates a mathematical expression provided in natural language format.
* `POST /validate`: Validates a mathematical expression provided in natural language format.
* `GET /errors`: Retrieves recorded errors.