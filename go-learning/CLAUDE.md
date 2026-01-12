# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go learning project demonstrating core Go concepts through practical examples. Module name: `go-learning`.

## Common Commands

```bash
# Run all examples
go run main.go

# Run tests
go test ./...

# Run a single test
go test -v -run TestFunctionName ./package/

# Run tests with coverage
go test -cover ./...

# Format code
go fmt ./...

# Check for issues
go vet ./...

# Clean up dependencies
go mod tidy
```

## Running Web Servers

```bash
# Run net/http server (port 8080)
go run cmd/nethttp/main.go

# Run Gin server (port 8081)
go run cmd/ginapp/main.go
```

## Architecture

The project is organized as separate packages, each demonstrating specific Go concepts:

- **main.go** - Entry point that imports and runs all example packages
- **basics/** - Variables, control flow, functions, slices, maps
- **structs/** - Struct definitions, methods, pointer receivers, embedding
- **interfaces/** - Interface definitions, polymorphism, type assertions
- **concurrency/** - Goroutines, channels, WaitGroups, select statements
- **nethttp/** - HTTP server using standard library `net/http`
- **ginapp/** - HTTP server using Gin framework
- **cmd/** - Standalone executables for running web servers independently

Each package exports a `Run*` function (e.g., `basics.RunBasics()`) called from main.go. The web packages also export setup functions (`nethttp.CreateServer()`, `ginapp.SetupRouter()`) used by their respective cmd entrypoints.
