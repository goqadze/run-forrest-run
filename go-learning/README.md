# Go Learning Project

A hands-on project to learn Go programming through practical examples.

## Project Structure

```
go-learning/
├── main.go              # Entry point that runs all examples
├── basics/              # Fundamental Go concepts
│   └── basics.go
├── structs/             # Structs, methods, and pointers
│   └── structs.go
├── interfaces/          # Interfaces and polymorphism
│   └── interfaces.go
└── concurrency/         # Goroutines, channels, and sync
    └── concurrency.go
```

## Getting Started

### Run All Examples

```bash
go run main.go
```

### Run Specific Examples

You can also run specific packages to focus on particular concepts:

```bash
# Just basics
go run main.go  # Then comment out other sections

# Or create test files to experiment
go run basics/basics.go
```

## Topics Covered

### 1. Basics (`basics/basics.go`)
- **Variables and Types**: Different ways to declare variables
- **Control Flow**: if-else, for loops, switch statements
- **Functions**: Multiple return values, named returns
- **Collections**: Slices (dynamic arrays) and maps

Key concepts:
- `:=` short declaration
- Go only has `for` loops (no while!)
- Functions can return multiple values

### 2. Structs (`structs/structs.go`)
- **Basic Structs**: Defining and creating structs
- **Methods**: Value receivers vs pointer receivers
- **Embedding**: Struct composition (Go's "inheritance")
- **Pointers**: Working with memory addresses

Key concepts:
- Methods with value receivers can't modify the struct
- Methods with pointer receivers can modify the struct
- Struct embedding allows composition over inheritance

### 3. Interfaces (`interfaces/interfaces.go`)
- **Interface Definition**: Defining contracts
- **Polymorphism**: Different types implementing the same interface
- **Empty Interface**: `interface{}` can hold any type
- **Type Assertions**: Checking and converting interface types

Key concepts:
- Interfaces are implemented implicitly (no "implements" keyword)
- Empty interface `interface{}` is like `any` in other languages
- Type assertions let you extract concrete types from interfaces

### 4. Concurrency (`concurrency/concurrency.go`)
- **Goroutines**: Lightweight threads
- **Channels**: Communication between goroutines
- **Buffered Channels**: Non-blocking sends with capacity
- **WaitGroups**: Synchronizing multiple goroutines
- **Select**: Multiplexing channel operations

Key concepts:
- Use `go` keyword to start a goroutine
- Channels provide safe communication between goroutines
- `select` allows waiting on multiple channel operations
- WaitGroups help coordinate goroutine completion

## Learning Path

1. **Start with Basics**: Understand variables, control flow, and functions
2. **Move to Structs**: Learn Go's approach to data structures
3. **Master Interfaces**: Understand Go's polymorphism
4. **Explore Concurrency**: Discover Go's superpower

## Exercises to Try

1. Add a new shape (Triangle) to the interfaces example
2. Create a new struct with methods for a BankAccount
3. Modify the concurrency example to use a channel to collect results from workers
4. Add error handling examples to the basics
5. Create a simple concurrent web scraper

## Common Go Patterns

**Error Handling:**
```go
result, err := someFunction()
if err != nil {
    // Handle error
    return err
}
// Use result
```

**Defer (cleanup):**
```go
file, err := os.Open("file.txt")
if err != nil {
    return err
}
defer file.Close() // Runs when function exits
```

**Range (iteration):**
```go
for index, value := range slice {
    fmt.Println(index, value)
}
```

## Resources

- **[Go Commands Reference](GO_COMMANDS.md)** - Quick reference for all Go CLI commands
- Official Tour: https://go.dev/tour/
- Go by Example: https://gobyexample.com/
- Effective Go: https://go.dev/doc/effective_go
- Go Documentation: https://go.dev/doc/

## Tips

- Go is opinionated: `gofmt` formats your code automatically
- Use `go fmt ./...` to format all files
- Use `go vet ./...` to find potential bugs
- Use `go mod tidy` to clean up dependencies
- Run `go test ./...` to run tests (add `_test.go` files!)

## Next Steps

After completing these examples, try building:
- A CLI tool using the `flag` package
- A REST API using `net/http` or a framework like Gin
- A concurrent file processor
- A simple database application using `database/sql`

Happy coding with Go!
