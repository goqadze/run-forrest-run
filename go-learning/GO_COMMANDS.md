# Go Commands Reference

A comprehensive reference guide for commonly-used Go CLI commands. This guide covers everything from basic compilation to advanced tooling.

## Table of Contents
- [Core Commands](#core-commands)
- [Module Management](#module-management)
- [Code Quality Tools](#code-quality-tools)
- [Documentation & Information](#documentation--information)
- [Advanced Commands](#advanced-commands)
- [Common Flags](#common-flags)

---

## Core Commands

### `go run`
Compile and run Go programs in one step (great for development and testing).

**Syntax:**
```bash
go run [flags] package [arguments...]
go run [flags] file.go [arguments...]
```

**Examples:**
```bash
# Run main.go
go run main.go

# Run with arguments
go run main.go arg1 arg2

# Run all files in a package
go run .

# Run specific files
go run file1.go file2.go
```

**Common Flags:**
- `-race` - Enable data race detection

---

### `go build`
Compile packages and their dependencies into an executable.

**Syntax:**
```bash
go build [flags] [packages]
```

**Examples:**
```bash
# Build current package (creates executable named after directory)
go build

# Build and name the output
go build -o myapp

# Build for different OS/architecture
GOOS=linux GOARCH=amd64 go build

# Build with optimizations disabled (for debugging)
go build -gcflags="all=-N -l"
```

**Common Flags:**
- `-o <file>` - Output file name
- `-v` - Verbose output (show packages being compiled)
- `-race` - Enable data race detector
- `-ldflags` - Pass flags to the linker
- `-tags` - Build tags

**Tips:**
- The resulting binary is platform-specific
- No output means success (unless using `-v`)
- Binary is created in current directory

---

### `go install`
Compile and install packages and dependencies.

**Syntax:**
```bash
go install [flags] [packages]
```

**Examples:**
```bash
# Install current package to $GOPATH/bin
go install

# Install a specific tool
go install github.com/user/tool@latest

# Install with version
go install golang.org/x/tools/gopls@v0.11.0
```

**What it does:**
- Compiles the package
- Places executables in `$GOPATH/bin` or `$GOBIN`
- Caches build artifacts in `$GOPATH/pkg`

---

### `go test`
Run tests for packages.

**Syntax:**
```bash
go test [flags] [packages]
```

**Examples:**
```bash
# Run tests in current package
go test

# Run all tests recursively
go test ./...

# Run tests with verbose output
go test -v

# Run specific test
go test -run TestMyFunction

# Run with coverage
go test -cover

# Generate coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run benchmarks
go test -bench=.

# Run tests with race detector
go test -race
```

**Common Flags:**
- `-v` - Verbose output
- `-run <regexp>` - Run only tests matching pattern
- `-cover` - Enable coverage analysis
- `-bench <regexp>` - Run benchmarks
- `-timeout <duration>` - Test timeout (default 10m)
- `-short` - Run shorter version of tests
- `-race` - Enable data race detector
- `-parallel <n>` - Allow N tests to run in parallel

---

### `go get`
Add dependencies to current module and install them.

**Syntax:**
```bash
go get [flags] [packages]
```

**Examples:**
```bash
# Add latest version of a package
go get github.com/gin-gonic/gin

# Add specific version
go get github.com/gin-gonic/gin@v1.9.0

# Update to latest
go get -u github.com/gin-gonic/gin

# Update all dependencies
go get -u ./...

# Remove a dependency (since Go 1.17)
go get github.com/unwanted/package@none
```

**Common Flags:**
- `-u` - Update to newer minor or patch versions
- `-t` - Consider test dependencies
- `-d` - Download only (don't install)

**Note:** As of Go 1.17+, `go install` is preferred for installing executables.

---

## Module Management

### `go mod init`
Initialize a new Go module.

**Syntax:**
```bash
go mod init [module-path]
```

**Examples:**
```bash
# Initialize with module name
go mod init myapp

# Initialize with full path (for published modules)
go mod init github.com/username/myapp
```

**Creates:** `go.mod` file in the current directory

---

### `go mod tidy`
Add missing and remove unused module dependencies.

**Syntax:**
```bash
go mod tidy [flags]
```

**Examples:**
```bash
# Clean up dependencies
go mod tidy

# Tidy with verbose output
go mod tidy -v
```

**What it does:**
- Adds missing dependencies required by code
- Removes dependencies that aren't used
- Updates `go.mod` and `go.sum`

**Best Practice:** Run after adding/removing imports

---

### `go mod download`
Download modules to local cache.

**Syntax:**
```bash
go mod download [modules]
```

**Examples:**
```bash
# Download all dependencies
go mod download

# Download specific module
go mod download github.com/gin-gonic/gin
```

**Location:** Downloads to `$GOPATH/pkg/mod`

---

### `go mod verify`
Verify dependencies have expected content.

**Syntax:**
```bash
go mod verify
```

**What it does:**
- Checks that dependencies haven't been modified
- Verifies checksums in `go.sum`

---

### `go mod vendor`
Make vendored copy of dependencies.

**Syntax:**
```bash
go mod vendor
```

**What it does:**
- Creates `vendor/` directory
- Copies all dependencies into `vendor/`
- Useful for ensuring reproducible builds

**Usage:**
```bash
# Create vendor directory
go mod vendor

# Build using vendor directory
go build -mod=vendor
```

---

### `go mod graph`
Print module requirement graph.

**Syntax:**
```bash
go mod graph
```

**Example output:**
```
myapp github.com/gin-gonic/gin@v1.9.0
github.com/gin-gonic/gin@v1.9.0 github.com/gin-contrib/sse@v0.1.0
```

---

### `go mod why`
Explain why packages or modules are needed.

**Syntax:**
```bash
go mod why [packages]
```

**Examples:**
```bash
# Why is this package needed?
go mod why github.com/some/package

# Check multiple packages
go mod why -m golang.org/x/crypto
```

---

## Code Quality Tools

### `go fmt`
Format Go source code according to standard style.

**Syntax:**
```bash
go fmt [packages]
```

**Examples:**
```bash
# Format current package
go fmt

# Format all packages recursively
go fmt ./...

# Format specific file
go fmt main.go
```

**Note:** Most editors run this automatically on save.

---

### `gofmt`
Lower-level formatting tool with more options.

**Syntax:**
```bash
gofmt [flags] [files]
```

**Examples:**
```bash
# Show diff without modifying
gofmt -d main.go

# Write changes to file
gofmt -w main.go

# Simplify code
gofmt -s -w main.go

# Format all files recursively
gofmt -w .
```

**Common Flags:**
- `-d` - Display diffs
- `-w` - Write result to file
- `-s` - Simplify code

---

### `go vet`
Examine code for suspicious constructs.

**Syntax:**
```bash
go vet [packages]
```

**Examples:**
```bash
# Check current package
go vet

# Check all packages
go vet ./...

# Check specific files
go vet main.go
```

**What it finds:**
- Unreachable code
- Useless assignments
- Printf format errors
- Struct tag mistakes
- And more...

**Best Practice:** Run before committing code

---

### `go fix`
Update Go code to use new APIs.

**Syntax:**
```bash
go fix [packages]
```

**What it does:**
- Rewrites code to use newer API versions
- Useful when upgrading Go versions

---

### `go clean`
Remove object files and cached files.

**Syntax:**
```bash
go clean [flags] [packages]
```

**Examples:**
```bash
# Clean current package
go clean

# Clean with cache
go clean -cache

# Clean module cache
go clean -modcache

# Clean test cache
go clean -testcache

# Remove all cached data
go clean -cache -modcache -testcache
```

**Common Flags:**
- `-cache` - Remove build cache
- `-modcache` - Remove module cache
- `-testcache` - Remove test cache
- `-i` - Remove installed archives

---

## Documentation & Information

### `go doc`
Show documentation for package or symbol.

**Syntax:**
```bash
go doc [package|symbol]
```

**Examples:**
```bash
# Show package documentation
go doc fmt

# Show function documentation
go doc fmt.Println

# Show method documentation
go doc http.Client.Get

# Show documentation for current package
go doc

# Show all documentation
go doc -all fmt
```

**Common Flags:**
- `-all` - Show all documentation
- `-u` - Show unexported symbols too

---

### `godoc`
Run documentation server (install separately).

**Installation:**
```bash
go install golang.org/x/tools/cmd/godoc@latest
```

**Usage:**
```bash
# Start documentation server on http://localhost:6060
godoc -http=:6060
```

---

### `go version`
Print Go version.

**Syntax:**
```bash
go version
```

**Example output:**
```
go version go1.23.0 darwin/arm64
```

---

### `go env`
Print Go environment information.

**Syntax:**
```bash
go env [var ...]
```

**Examples:**
```bash
# Show all environment variables
go env

# Show specific variable
go env GOPATH

# Show multiple variables
go env GOPATH GOROOT GOOS GOARCH

# Set environment variable (Go 1.13+)
go env -w GOPRIVATE=github.com/mycompany/*
```

**Common Variables:**
- `GOPATH` - Workspace location
- `GOROOT` - Go installation location
- `GOOS` - Target operating system
- `GOARCH` - Target architecture
- `GOBIN` - Where `go install` puts binaries
- `GOPROXY` - Module proxy URL
- `GOPRIVATE` - Private modules pattern

---

### `go list`
List packages or modules.

**Syntax:**
```bash
go list [flags] [packages]
```

**Examples:**
```bash
# List all packages in module
go list ./...

# List dependencies
go list -m all

# Show package details in JSON
go list -json ./...

# List test files
go list -f '{{.TestGoFiles}}' ./...
```

**Common Flags:**
- `-m` - List modules instead of packages
- `-json` - Print in JSON format
- `-f <format>` - Custom format using template

---

## Advanced Commands

### `go generate`
Generate Go files by processing source.

**Syntax:**
```bash
go generate [flags] [packages]
```

**How it works:**
1. Looks for `//go:generate` directives in `.go` files
2. Runs the commands specified
3. Typically used for code generation

**Example directive:**
```go
//go:generate stringer -type=Status
```

**Running:**
```bash
# Generate for current package
go generate

# Generate for all packages
go generate ./...
```

---

### `go work`
Workspace maintenance (Go 1.18+).

**Syntax:**
```bash
go work init [modules]
go work use [modules]
go work sync
```

**Examples:**
```bash
# Initialize workspace
go work init ./module1 ./module2

# Add module to workspace
go work use ./module3

# Sync workspace dependencies
go work sync
```

**What it does:**
- Allows working with multiple modules simultaneously
- Creates `go.work` file

---

### `go tool`
Run specified Go tool.

**Syntax:**
```bash
go tool <tool> [args]
```

**Examples:**
```bash
# View available tools
go tool

# View assembly
go tool compile -S main.go

# Analyze CPU profile
go tool pprof cpu.prof

# View coverage in HTML
go tool cover -html=coverage.out
```

**Common Tools:**
- `compile` - Go compiler
- `link` - Go linker
- `pprof` - Profiler
- `cover` - Coverage analysis
- `trace` - Execution tracer

---

## Common Flags

These flags work with many Go commands:

- `-n` - Print commands but don't run them
- `-x` - Print commands as they're executed
- `-v` - Verbose output
- `-race` - Enable race detector
- `-work` - Print temporary work directory
- `-a` - Force rebuilding of packages
- `-p n` - Number of programs to run in parallel

---

## Quick Tips

1. **Most used commands:**
   ```bash
   go run .          # Quick testing
   go build          # Create executable
   go test ./...     # Run all tests
   go mod tidy       # Clean dependencies
   go fmt ./...      # Format code
   go vet ./...      # Check for issues
   ```

2. **Before committing:**
   ```bash
   go fmt ./...
   go vet ./...
   go test ./...
   ```

3. **Cross-compilation:**
   ```bash
   GOOS=linux GOARCH=amd64 go build
   GOOS=windows GOARCH=amd64 go build
   GOOS=darwin GOARCH=arm64 go build
   ```

4. **Get help:**
   ```bash
   go help <command>
   go help build
   go help modules
   ```

---

## See Also

- [Go Official Documentation](https://go.dev/doc/)
- [Go Command Documentation](https://pkg.go.dev/cmd/go)
- [Effective Go](https://go.dev/doc/effective_go)
- [README.md](README.md) - Project learning guide
