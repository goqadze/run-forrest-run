package ginapp

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// ============================================================
// GIN FRAMEWORK MODULE - Learning Guide
// ============================================================
// This module covers:
// 1. Basic Gin setup and routing
// 2. Route groups and versioning
// 3. Middleware (built-in and custom)
// 4. Request binding (JSON, Query, URI)
// 5. Response types (JSON, XML, HTML)
// 6. Error handling
// 7. Validation
// ============================================================

// ============================================================
// MODELS
// ============================================================

// Book represents a book in our API
type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" binding:"required,min=1,max=200"`
	Author    string    `json:"author" binding:"required"`
	Year      int       `json:"year" binding:"required,gte=1000,lte=2100"`
	ISBN      string    `json:"isbn,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateBookInput - input for creating a book (without ID and CreatedAt)
type CreateBookInput struct {
	Title  string `json:"title" binding:"required,min=1,max=200"`
	Author string `json:"author" binding:"required"`
	Year   int    `json:"year" binding:"required,gte=1000,lte=2100"`
	ISBN   string `json:"isbn,omitempty"`
}

// UpdateBookInput - input for updating a book (all fields optional)
type UpdateBookInput struct {
	Title  *string `json:"title,omitempty" binding:"omitempty,min=1,max=200"`
	Author *string `json:"author,omitempty"`
	Year   *int    `json:"year,omitempty" binding:"omitempty,gte=1000,lte=2100"`
	ISBN   *string `json:"isbn,omitempty"`
}

// In-memory store
var (
	books   = make(map[int]Book)
	booksMu sync.RWMutex
	bookID  = 1
)

// ============================================================
// 1. BASIC HANDLERS
// ============================================================

// HealthCheck - GET /health
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

// Welcome - GET /
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Gin Learning API!",
		"version": "1.0.0",
		"docs":    "/api/v1",
	})
}

// ============================================================
// 2. CRUD HANDLERS FOR BOOKS
// ============================================================

// GetBooks - GET /api/v1/books
func GetBooks(c *gin.Context) {
	booksMu.RLock()
	defer booksMu.RUnlock()

	bookList := make([]Book, 0, len(books))
	for _, b := range books {
		bookList = append(bookList, b)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  bookList,
		"count": len(bookList),
	})
}

// GetBook - GET /api/v1/books/:id
func GetBook(c *gin.Context) {
	// URI binding - get ID from path parameter
	var uri struct {
		ID int `uri:"id" binding:"required,min=1"`
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	booksMu.RLock()
	book, exists := books[uri.ID]
	booksMu.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// CreateBook - POST /api/v1/books
func CreateBook(c *gin.Context) {
	var input CreateBookInput

	// JSON binding with validation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	booksMu.Lock()
	book := Book{
		ID:        bookID,
		Title:     input.Title,
		Author:    input.Author,
		Year:      input.Year,
		ISBN:      input.ISBN,
		CreatedAt: time.Now(),
	}
	books[bookID] = book
	bookID++
	booksMu.Unlock()

	c.JSON(http.StatusCreated, gin.H{"data": book})
}

// UpdateBook - PUT /api/v1/books/:id
func UpdateBook(c *gin.Context) {
	var uri struct {
		ID int `uri:"id" binding:"required,min=1"`
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	booksMu.Lock()
	defer booksMu.Unlock()

	book, exists := books[uri.ID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Update only provided fields
	if input.Title != nil {
		book.Title = *input.Title
	}
	if input.Author != nil {
		book.Author = *input.Author
	}
	if input.Year != nil {
		book.Year = *input.Year
	}
	if input.ISBN != nil {
		book.ISBN = *input.ISBN
	}

	books[uri.ID] = book
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook - DELETE /api/v1/books/:id
func DeleteBook(c *gin.Context) {
	var uri struct {
		ID int `uri:"id" binding:"required,min=1"`
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	booksMu.Lock()
	defer booksMu.Unlock()

	if _, exists := books[uri.ID]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	delete(books, uri.ID)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

// ============================================================
// 3. QUERY PARAMETERS EXAMPLE
// ============================================================

// SearchBooks - GET /api/v1/books/search?author=...&year=...
func SearchBooks(c *gin.Context) {
	var query struct {
		Author string `form:"author"`
		Year   int    `form:"year"`
		Limit  int    `form:"limit,default=10"`
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	booksMu.RLock()
	defer booksMu.RUnlock()

	results := make([]Book, 0)
	count := 0

	for _, b := range books {
		if count >= query.Limit {
			break
		}

		matchAuthor := query.Author == "" || b.Author == query.Author
		matchYear := query.Year == 0 || b.Year == query.Year

		if matchAuthor && matchYear {
			results = append(results, b)
			count++
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   results,
		"count":  len(results),
		"filter": query,
	})
}

// ============================================================
// 4. CUSTOM MIDDLEWARE
// ============================================================

// RequestIDMiddleware - adds unique request ID to each request
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := fmt.Sprintf("%d", time.Now().UnixNano())
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}

// TimingMiddleware - logs request duration
func TimingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		c.Header("X-Response-Time", duration.String())
	}
}

// AuthMiddleware - simple auth middleware example
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		// Simple demo - in real app, validate JWT or session
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// For demo: accept any token starting with "Bearer "
		if len(token) < 7 || token[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// Set user info in context (in real app, decode from token)
		c.Set("user", "demo_user")
		c.Next()
	}
}

// ============================================================
// 5. DIFFERENT RESPONSE FORMATS
// ============================================================

// ResponseFormats - GET /api/v1/formats
func ResponseFormats(c *gin.Context) {
	format := c.DefaultQuery("format", "json")
	data := gin.H{
		"message": "Hello from Gin!",
		"formats": []string{"json", "xml", "yaml"},
	}

	switch format {
	case "xml":
		c.XML(http.StatusOK, data)
	case "yaml":
		c.YAML(http.StatusOK, data)
	default:
		c.JSON(http.StatusOK, data)
	}
}

// ============================================================
// 6. ERROR HANDLING
// ============================================================

// Custom error type
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrorHandlerMiddleware - centralized error handling
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check for errors after handler execution
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	}
}

// ============================================================
// ROUTER SETUP
// ============================================================

// SetupRouter creates and configures the Gin router
func SetupRouter() *gin.Engine {
	// Set release mode for production (less logging)
	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default() // Includes Logger and Recovery middleware

	// Add custom middleware
	router.Use(RequestIDMiddleware())
	router.Use(TimingMiddleware())
	router.Use(ErrorHandlerMiddleware())

	// Root routes
	router.GET("/", Welcome)
	router.GET("/health", HealthCheck)

	// API v1 group
	v1 := router.Group("/api/v1")
	{
		// Public routes
		v1.GET("/formats", ResponseFormats)

		// Books CRUD
		booksGroup := v1.Group("/books")
		{
			booksGroup.GET("", GetBooks)
			booksGroup.GET("/:id", GetBook)
			booksGroup.GET("/search", SearchBooks)
			booksGroup.POST("", CreateBook)
			booksGroup.PUT("/:id", UpdateBook)
			booksGroup.DELETE("/:id", DeleteBook)
		}

		// Protected routes (require auth)
		protected := v1.Group("/admin")
		protected.Use(AuthMiddleware())
		{
			protected.GET("/stats", func(c *gin.Context) {
				user := c.GetString("user")
				booksMu.RLock()
				count := len(books)
				booksMu.RUnlock()

				c.JSON(http.StatusOK, gin.H{
					"user":        user,
					"total_books": count,
					"timestamp":   time.Now().Format(time.RFC3339),
				})
			})
		}
	}

	return router
}

// ============================================================
// MAIN DEMO FUNCTION
// ============================================================

// RunGinApp demonstrates Gin concepts (without starting server)
func RunGinApp() {
	fmt.Println("GIN Framework Module - High-Performance HTTP Web Framework")
	fmt.Println("============================================================")

	fmt.Println("\n1. Key Features:")
	fmt.Println("   - Fast routing with radix tree")
	fmt.Println("   - Middleware support")
	fmt.Println("   - JSON validation")
	fmt.Println("   - Route grouping")
	fmt.Println("   - Error handling")

	fmt.Println("\n2. Available Endpoints:")
	fmt.Println("   GET  /                     - Welcome message")
	fmt.Println("   GET  /health               - Health check")
	fmt.Println("   GET  /api/v1/formats       - Response format demo")
	fmt.Println("   GET  /api/v1/books         - List all books")
	fmt.Println("   GET  /api/v1/books/:id     - Get book by ID")
	fmt.Println("   GET  /api/v1/books/search  - Search books (?author=...&year=...)")
	fmt.Println("   POST /api/v1/books         - Create book")
	fmt.Println("   PUT  /api/v1/books/:id     - Update book")
	fmt.Println("   DELETE /api/v1/books/:id   - Delete book")
	fmt.Println("   GET  /api/v1/admin/stats   - Admin stats (requires auth)")

	fmt.Println("\n3. To start the server, run:")
	fmt.Println("   go run cmd/ginapp/main.go")

	// Load sample data
	booksMu.Lock()
	books[bookID] = Book{ID: bookID, Title: "The Go Programming Language", Author: "Donovan & Kernighan", Year: 2015, CreatedAt: time.Now()}
	bookID++
	books[bookID] = Book{ID: bookID, Title: "Learning Go", Author: "Jon Bodner", Year: 2021, CreatedAt: time.Now()}
	bookID++
	books[bookID] = Book{ID: bookID, Title: "Concurrency in Go", Author: "Katherine Cox-Buday", Year: 2017, CreatedAt: time.Now()}
	bookID++
	booksMu.Unlock()

	fmt.Printf("\n4. Sample books loaded: %d books\n", len(books))
}
