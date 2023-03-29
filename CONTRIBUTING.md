# 💼 Contributing

- [💽 Development](#-development)
- [🖥 Requirements](#-requirements)
- [✅ Testing](#-testing)
- [🚀 Releasing](#-releasing)


## 💽 Development

Run the following command to run the application:

    go run github.com/techygrrrl/timerrr

## 🖥 Requirements

- `go 1.20`
- [cobra-cli](https://github.com/spf13/cobra#usage)


## ✅ Testing

Run the tests with the following command:

    go test -v ./...

To run tests with coverage run:

    go test -v ./... -cover

Write any operating-specific tests in a file for that operating system, e.g. `os_utils_linux_test.go` will only run on Linux.


## 🚀 Releasing

Uses [goreleaser](https://goreleaser.com/quick-start/).
