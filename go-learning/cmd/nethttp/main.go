package main

import (
	"fmt"
	"go-learning/nethttp"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("=== NET/HTTP Server ===")
	fmt.Println()

	addr := ":8080"
	server := nethttp.CreateServer(addr)

	// Graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		fmt.Println("\nShutting down server...")
		os.Exit(0)
	}()

	fmt.Printf("Server starting on http://localhost%s\n", addr)
	fmt.Println()
	fmt.Println("Available endpoints:")
	fmt.Println("  GET  /           - Hello World")
	fmt.Println("  GET  /time       - Current server time")
	fmt.Println("  GET  /info       - Request information")
	fmt.Println("  GET  /query      - Query params demo (?name=John&age=30)")
	fmt.Println("  GET  /api/users  - List all users (JSON)")
	fmt.Println("  POST /api/users  - Create user (JSON body: {\"name\":\"...\",\"email\":\"...\"})")
	fmt.Println()
	fmt.Println("Press Ctrl+C to stop")
	fmt.Println()

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
