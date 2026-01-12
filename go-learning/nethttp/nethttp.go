package nethttp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// ============================================================
// NET/HTTP MODULE - Learning Guide
// ============================================================
// This module covers:
// 1. Basic HTTP Server
// 2. HTTP Handlers and HandlerFunc
// 3. Routing with ServeMux
// 4. Middleware pattern
// 5. JSON APIs
// 6. HTTP Client
// 7. Request/Response handling
// ============================================================

// User represents a simple user model for JSON examples
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// In-memory store for demo purposes
var (
	users   = make(map[int]User)
	usersMu sync.RWMutex
	nextID  = 1
)

// ============================================================
// 1. BASIC HANDLERS
// ============================================================

// HelloHandler - simplest possible handler
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// TimeHandler - returns current server time
func TimeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC1123)
	fmt.Fprintf(w, "Current server time: %s\n", currentTime)
}

// ============================================================
// 2. REQUEST INFORMATION
// ============================================================

// RequestInfoHandler - demonstrates how to read request details
func RequestInfoHandler(w http.ResponseWriter, r *http.Request) {
	info := fmt.Sprintf(`Request Information:
  Method:     %s
  URL:        %s
  Path:       %s
  Host:       %s
  RemoteAddr: %s
  User-Agent: %s
  Headers:    %v
`, r.Method, r.URL.String(), r.URL.Path, r.Host, r.RemoteAddr, r.UserAgent(), r.Header)

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, info)
}

// ============================================================
// 3. QUERY PARAMETERS
// ============================================================

// QueryParamsHandler - demonstrates reading query parameters
// Example: /query?name=John&age=30
func QueryParamsHandler(w http.ResponseWriter, r *http.Request) {
	// Get single value
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	// Get all values for a key (useful for ?color=red&color=blue)
	allValues := r.URL.Query()

	response := fmt.Sprintf(`Query Parameters:
  name: %s
  age:  %s
  All params: %v

Try: /query?name=John&age=30
`, name, age, allValues)

	fmt.Fprint(w, response)
}

// ============================================================
// 4. JSON RESPONSES
// ============================================================

// JSONResponse helper - sends JSON with proper headers
func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// ErrorResponse helper - sends error as JSON
func ErrorResponse(w http.ResponseWriter, status int, message string) {
	JSONResponse(w, status, map[string]string{"error": message})
}

// GetUsersHandler - GET /api/users - returns all users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	usersMu.RLock()
	userList := make([]User, 0, len(users))
	for _, u := range users {
		userList = append(userList, u)
	}
	usersMu.RUnlock()

	JSONResponse(w, http.StatusOK, userList)
}

// CreateUserHandler - POST /api/users - creates a new user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid JSON: "+err.Error())
		return
	}

	if input.Name == "" || input.Email == "" {
		ErrorResponse(w, http.StatusBadRequest, "Name and email are required")
		return
	}

	usersMu.Lock()
	user := User{
		ID:        nextID,
		Name:      input.Name,
		Email:     input.Email,
		CreatedAt: time.Now(),
	}
	users[nextID] = user
	nextID++
	usersMu.Unlock()

	JSONResponse(w, http.StatusCreated, user)
}

// ============================================================
// 5. MIDDLEWARE PATTERN
// ============================================================

// Middleware type for chaining
type Middleware func(http.Handler) http.Handler

// LoggingMiddleware - logs all requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("[%s] %s %s\n", time.Now().Format("15:04:05"), r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		fmt.Printf("[%s] Completed in %v\n", time.Now().Format("15:04:05"), time.Since(start))
	})
}

// RecoveryMiddleware - recovers from panics
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Panic recovered: %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// CORSMiddleware - adds CORS headers
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ChainMiddleware - chains multiple middlewares
func ChainMiddleware(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// ============================================================
// 6. HTTP CLIENT EXAMPLES
// ============================================================

// HTTPClientExample - demonstrates making HTTP requests
func HTTPClientExample() {
	fmt.Println("\n--- HTTP Client Examples ---")

	// Create a client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Simple GET request
	fmt.Println("\n1. Simple GET request:")
	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("   Status: %s\n", resp.Status)
	fmt.Printf("   Headers: %v\n", resp.Header.Get("Content-Type"))

	// Read response body
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("   Body length: %d bytes\n", len(body))
}

// ============================================================
// 7. CUSTOM SERVER CONFIGURATION
// ============================================================

// CreateServer - creates a configured HTTP server
func CreateServer(addr string) *http.Server {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/", HelloHandler)
	mux.HandleFunc("/time", TimeHandler)
	mux.HandleFunc("/info", RequestInfoHandler)
	mux.HandleFunc("/query", QueryParamsHandler)
	mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetUsersHandler(w, r)
		case http.MethodPost:
			CreateUserHandler(w, r)
		default:
			ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})

	// Wrap with middleware
	handler := ChainMiddleware(mux,
		RecoveryMiddleware,
		LoggingMiddleware,
		CORSMiddleware,
	)

	return &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

// ============================================================
// MAIN DEMO FUNCTION
// ============================================================

// RunNetHTTP demonstrates net/http concepts (without starting server)
func RunNetHTTP() {
	fmt.Println("NET/HTTP Module - Go Standard Library HTTP Package")
	fmt.Println("==================================================")

	fmt.Println("\n1. HTTP Handlers:")
	fmt.Println("   - HelloHandler:       Basic 'Hello, World!' response")
	fmt.Println("   - TimeHandler:        Returns server time")
	fmt.Println("   - RequestInfoHandler: Shows request details")
	fmt.Println("   - QueryParamsHandler: Demonstrates query parsing")

	fmt.Println("\n2. JSON API Handlers:")
	fmt.Println("   - GET  /api/users:    List all users")
	fmt.Println("   - POST /api/users:    Create a user")

	fmt.Println("\n3. Middleware:")
	fmt.Println("   - LoggingMiddleware:  Logs request duration")
	fmt.Println("   - RecoveryMiddleware: Catches panics")
	fmt.Println("   - CORSMiddleware:     Adds CORS headers")

	fmt.Println("\n4. To start the server, run:")
	fmt.Println("   go run cmd/nethttp/main.go")

	// Demo: Add sample users
	usersMu.Lock()
	users[nextID] = User{ID: nextID, Name: "Alice", Email: "alice@example.com", CreatedAt: time.Now()}
	nextID++
	users[nextID] = User{ID: nextID, Name: "Bob", Email: "bob@example.com", CreatedAt: time.Now()}
	nextID++
	usersMu.Unlock()

	fmt.Printf("\n5. Sample users loaded: %d users\n", len(users))
}
