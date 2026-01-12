package main

import (
	"fmt"
	"go-learning/ginapp"
)

func main() {
	fmt.Println("=== GIN Web Application ===")
	fmt.Println()

	router := ginapp.SetupRouter()

	fmt.Println("Available endpoints:")
	fmt.Println("  GET  /                     - Welcome message")
	fmt.Println("  GET  /health               - Health check")
	fmt.Println("  GET  /api/v1/formats       - Response formats (?format=json|xml|yaml)")
	fmt.Println("  GET  /api/v1/books         - List all books")
	fmt.Println("  GET  /api/v1/books/:id     - Get book by ID")
	fmt.Println("  GET  /api/v1/books/search  - Search (?author=...&year=...&limit=10)")
	fmt.Println("  POST /api/v1/books         - Create book (JSON body)")
	fmt.Println("  PUT  /api/v1/books/:id     - Update book (JSON body)")
	fmt.Println("  DELETE /api/v1/books/:id   - Delete book")
	fmt.Println("  GET  /api/v1/admin/stats   - Stats (requires: Authorization: Bearer <token>)")
	fmt.Println()
	fmt.Println("Example curl commands:")
	fmt.Println("  curl http://localhost:8081/api/v1/books")
	fmt.Println("  curl -X POST http://localhost:8081/api/v1/books -H 'Content-Type: application/json' -d '{\"title\":\"My Book\",\"author\":\"Me\",\"year\":2024}'")
	fmt.Println("  curl http://localhost:8081/api/v1/admin/stats -H 'Authorization: Bearer mytoken'")
	fmt.Println()
	fmt.Println("Starting server on http://localhost:8081")
	fmt.Println("Press Ctrl+C to stop")
	fmt.Println()

	if err := router.Run(":8081"); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
