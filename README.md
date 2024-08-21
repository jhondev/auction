# Auction

## Dev Setup

### Prerequisites

- Go

  https://go.dev/doc/install

- Task tool

  ```
  go install github.com/go-task/task/v3/cmd/task@latest
  ```

### Run Setup

```
task setup
```

## Dev Workflow

### Tasks

- Lint and Build

  ```
  task build
  ```

- Run unit tests

  ```
  task test:unit
  ```

- You can list all the available tasks executing `task`

  ```
  task
  ```
