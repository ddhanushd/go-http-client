# Go HTTP Client with Tests & CI

![Go CI](https://github.com/ddhanushd/go-http-client/actions/workflows/ci.yml/badge.svg)

This project is a small Go module that demonstrates how to build a testable HTTP client with mocked external calls using `httpmatter` and continuous integration using GitHub Actions.

## Features

- Simple HTTP client to fetch city data
- Dependency injection for testability
- Mocked HTTP calls using fixture-based approach (httpmatter)
- Unit tests
- GitHub Actions CI that runs tests on every push

## Running Tests Locally

To run the tests locally, use:

```bash
go test ./...
