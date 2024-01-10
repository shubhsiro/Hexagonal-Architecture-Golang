# Go Hexagonal Architecture

This project demonstrates a simple Go application following the Hexagonal Architecture pattern. The application is structured with separate layers for adapters, core logic, and delivery.

## Project Structure

| Directory                    | Description                                      |
| ---------------------------- | ------------------------------------------------ |
| `cmd/`                       | Application entry point                          |
| `internal/`                  | Internal application components                  |
| `internal/adapter/`          | Adapters for external systems                    |
| `internal/core/`             | Core business logic components                   |
| `internal/delivery/`         | Delivery mechanisms (e.g., HTTP)                 |
| `tests/`                     | Unit tests for various components                |

### Components

#### 1. Main

The `cmd/main.go` file serves as the entry point for the application. It initializes dependencies and starts the application.

#### 2. Adapters

| Directory                          | Description                                         |
| ---------------------------------- | --------------------------------------------------- |
| `internal/adapter/consensus/`      | RAFT Consensus Adapter (`raft.go`)                 |
| `internal/adapter/database/`       | Redis Database Adapter (`redis.go`)                |
| `internal/adapter/logger/`         | Logger Adapter (`hclog.go`)                        |

#### 3. Core Logic

| Directory                              | Description                                               |
| -------------------------------------- | --------------------------------------------------------- |
| `internal/core/fileprocessing/`        | File Processing Logic (`fileprocessor.go`)                |
| `internal/core/consensus/`             | Consensus Logic (`consensus.go`)                          |
| `internal/core/database/`              | Redis Database Logic (`redisdb.go`)                       |
| `internal/core/logger/`                | Logger Logic (`logger.go`)                                |

#### 4. HTTP Delivery

| Directory                              | Description                                   |
| -------------------------------------- | --------------------------------------------- |
| `internal/delivery/handler/`              | HTTP Handler (`handler.go`)                    |


## Getting Started

To run the application, execute the following command:

```bash
go run cmd/main.go
```

## Dependencies

- **go-hclog:** Used for structured logging.

- **github.com/stretchr/testify/assert:** Used for assertions in unit tests.

