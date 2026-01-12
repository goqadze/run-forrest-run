package main

import (
	"fmt"
	"go-learning/basics"
	"go-learning/concurrency"
	"go-learning/ginapp"
	"go-learning/interfaces"
	"go-learning/nethttp"
	"go-learning/structs"
)

func main() {
	fmt.Println("=== Welcome to Go Learning! ===\n")

	// Run basics examples
	fmt.Println("--- 1. Basics ---")
	basics.RunBasics()

	// Run struct examples
	fmt.Println("\n--- 2. Structs ---")
	structs.RunStructs()

	// Run interface examples
	fmt.Println("\n--- 3. Interfaces ---")
	interfaces.RunInterfaces()

	// Run concurrency examples
	fmt.Println("\n--- 4. Concurrency ---")
	concurrency.RunConcurrency()

	// Run net/http examples
	fmt.Println("\n--- 5. Net/HTTP ---")
	nethttp.RunNetHTTP()

	// Run Gin examples
	fmt.Println("\n--- 6. Gin Framework ---")
	ginapp.RunGinApp()

	fmt.Println("\n=== All examples completed! ===")
}
