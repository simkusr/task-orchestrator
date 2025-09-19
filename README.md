# task-orchestrator
Task Orchestrator is a lightweight Go microservice for managing and executing background tasks.

## Setup

### Environment variables
Default environment variables file name is held `.env` but can be
changed via environment variable `ENV_FILE_NAME`.

#### used environment variables
1. `PORT` - defines what kind of port we want to use, passed as in
`PORT=:8080` format


## Development

### Makefile commands

1. `make run` -> starts `task-orchestrator`
2. `make lint` -> runs `golangci-lint`
3. `make test` -> runs `go test ./... -v`
