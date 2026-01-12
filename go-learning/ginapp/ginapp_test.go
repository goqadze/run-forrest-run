package ginapp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return SetupRouter()
}

func resetBooks() {
	booksMu.Lock()
	books = make(map[int]Book)
	bookID = 1
	booksMu.Unlock()
}

func TestWelcome(t *testing.T) {
	router := setupTestRouter()
	resetBooks()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if response["message"] != "Welcome to the Gin Learning API!" {
		t.Errorf("Unexpected message: %v", response["message"])
	}
	if response["version"] != "1.0.0" {
		t.Errorf("Unexpected version: %v", response["version"])
	}
}

func TestHealthCheck(t *testing.T) {
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if response["status"] != "ok" {
		t.Errorf("Expected status 'ok', got %v", response["status"])
	}
	if _, exists := response["timestamp"]; !exists {
		t.Error("Expected timestamp in response")
	}
}

func TestGetBooks_Empty(t *testing.T) {
	router := setupTestRouter()
	resetBooks()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/books", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if response["count"].(float64) != 0 {
		t.Errorf("Expected count 0, got %v", response["count"])
	}
}

func TestCreateBook(t *testing.T) {
	router := setupTestRouter()
	resetBooks()

	book := CreateBookInput{
		Title:  "Test Book",
		Author: "Test Author",
		Year:   2024,
		ISBN:   "123-456",
	}
	body, _ := json.Marshal(book)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	data := response["data"].(map[string]interface{})
	if data["title"] != "Test Book" {
		t.Errorf("Expected title 'Test Book', got %v", data["title"])
	}
	if data["author"] != "Test Author" {
		t.Errorf("Expected author 'Test Author', got %v", data["author"])
	}
}

func TestCreateBook_ValidationError(t *testing.T) {
	router := setupTestRouter()
	resetBooks()

	tests := []struct {
		name string
		body map[string]interface{}
	}{
		{
			name: "missing title",
			body: map[string]interface{}{"author": "Author", "year": 2024},
		},
		{
			name: "missing author",
			body: map[string]interface{}{"title": "Title", "year": 2024},
		},
		{
			name: "missing year",
			body: map[string]interface{}{"title": "Title", "author": "Author"},
		},
		{
			name: "year too low",
			body: map[string]interface{}{"title": "Title", "author": "Author", "year": 999},
		},
		{
			name: "year too high",
			body: map[string]interface{}{"title": "Title", "author": "Author", "year": 2101},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.body)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/api/v1/books", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			if w.Code != http.StatusBadRequest {
				t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
			}
		})
	}
}

func TestGetBook(t *testing.T) {
	router := setupTestRouter()
	resetBooks()

	// Create a book first
	book := CreateBookInput{Title: "Test Book", Author: "Test Author", Year: 2024}
	body, _ := json.Marshal(book)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Get the book
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/books/1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	data := response["data"].(map[string]interface{})
	if data["title"] != "Test Book" {
		t.Errorf("Expected title 'Test Book', got %v", data["title"])
	}
}

func TestGetBook_NotFound(t *testing.T) {
	router := setupTestRouter()
	resetBooks()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/books/999", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestGetBook_InvalidID(t *testing.T) {
	router := setupTestRouter()
	resetBooks()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/books/invalid", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestUpdateBook(t *testing.T) {
	router := setupTestRouter()
	resetBooks()

	// Create a book first
	book := CreateBookInput{Title: "Original Title", Author: "Original Author", Year: 2024}
	body, _ := json.Marshal(book)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Update the book
	newTitle := "Updated Title"
	update := UpdateBookInput{Title: &newTitle}
	body, _ = json.Marshal(update)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/api/v1/books/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	data := response["data"].(map[string]interface{})
	if data["title"] != "Updated Title" {
		t.Errorf("Expected title 'Updated Title', got %v", data["title"])
	}
	// Author should remain unchanged
	if data["author"] != "Original Author" {
		t.Errorf("Expected author 'Original Author', got %v", data["author"])
	}
}

func TestUpdateBook_NotFound(t *testing.T) {
	router := setupTestRouter()
	resetBooks()

	newTitle := "Updated Title"
	update := UpdateBookInput{Title: &newTitle}
	body, _ := json.Marshal(update)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/books/999", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestDeleteBook(t *testing.T) {
	router := setupTestRouter()
	resetBooks()

	// Create a book first
	book := CreateBookInput{Title: "Test Book", Author: "Test Author", Year: 2024}
	body, _ := json.Marshal(book)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Delete the book
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/api/v1/books/1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Verify book is deleted
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/books/1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d after delete, got %d", http.StatusNotFound, w.Code)
	}
}

func TestDeleteBook_NotFound(t *testing.T) {
	router := setupTestRouter()
	resetBooks()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/books/999", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestSearchBooks(t *testing.T) {
	router := setupTestRouter()
	resetBooks()

	// Create some books
	books := []CreateBookInput{
		{Title: "Go Book 1", Author: "Author A", Year: 2020},
		{Title: "Go Book 2", Author: "Author A", Year: 2021},
		{Title: "Python Book", Author: "Author B", Year: 2020},
	}

	for _, book := range books {
		body, _ := json.Marshal(book)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/books", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
	}

	tests := []struct {
		name          string
		query         string
		expectedCount int
	}{
		{
			name:          "search by author",
			query:         "?author=Author A",
			expectedCount: 2,
		},
		{
			name:          "search by year",
			query:         "?year=2020",
			expectedCount: 2,
		},
		{
			name:          "search by author and year",
			query:         "?author=Author A&year=2020",
			expectedCount: 1,
		},
		{
			name:          "search with limit",
			query:         "?limit=1",
			expectedCount: 1,
		},
		{
			name:          "search no match",
			query:         "?author=NonExistent",
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/v1/books/search"+tt.query, nil)
			router.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
			}

			var response map[string]interface{}
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to parse response: %v", err)
			}

			count := int(response["count"].(float64))
			if count != tt.expectedCount {
				t.Errorf("Expected count %d, got %d", tt.expectedCount, count)
			}
		})
	}
}

func TestResponseFormats(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name        string
		format      string
		contentType string
	}{
		{
			name:        "default json",
			format:      "",
			contentType: "application/json",
		},
		{
			name:        "explicit json",
			format:      "?format=json",
			contentType: "application/json",
		},
		{
			name:        "xml format",
			format:      "?format=xml",
			contentType: "application/xml",
		},
		{
			name:        "yaml format",
			format:      "?format=yaml",
			contentType: "application/yaml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/v1/formats"+tt.format, nil)
			router.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
			}

			contentType := w.Header().Get("Content-Type")
			if contentType != tt.contentType+"; charset=utf-8" {
				t.Errorf("Expected content type %s, got %s", tt.contentType, contentType)
			}
		})
	}
}

func TestAuthMiddleware_NoToken(t *testing.T) {
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/admin/stats", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status %d, got %d", http.StatusUnauthorized, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if response["error"] != "Authorization header required" {
		t.Errorf("Unexpected error message: %v", response["error"])
	}
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/admin/stats", nil)
	req.Header.Set("Authorization", "InvalidToken")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status %d, got %d", http.StatusUnauthorized, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if response["error"] != "Invalid token format" {
		t.Errorf("Unexpected error message: %v", response["error"])
	}
}

func TestAuthMiddleware_ValidToken(t *testing.T) {
	router := setupTestRouter()
	resetBooks()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/admin/stats", nil)
	req.Header.Set("Authorization", "Bearer mytoken")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if response["user"] != "demo_user" {
		t.Errorf("Expected user 'demo_user', got %v", response["user"])
	}
}

func TestRequestIDMiddleware(t *testing.T) {
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	requestID := w.Header().Get("X-Request-ID")
	if requestID == "" {
		t.Error("Expected X-Request-ID header to be set")
	}
}

func TestTimingMiddleware(t *testing.T) {
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	responseTime := w.Header().Get("X-Response-Time")
	if responseTime == "" {
		t.Error("Expected X-Response-Time header to be set")
	}
}

func TestGetBooks_WithData(t *testing.T) {
	router := setupTestRouter()
	resetBooks()

	// Create multiple books
	for i := 0; i < 3; i++ {
		book := CreateBookInput{
			Title:  "Book " + string(rune('A'+i)),
			Author: "Author",
			Year:   2020 + i,
		}
		body, _ := json.Marshal(book)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/books", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
	}

	// Get all books
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/books", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	count := int(response["count"].(float64))
	if count != 3 {
		t.Errorf("Expected count 3, got %d", count)
	}

	data := response["data"].([]interface{})
	if len(data) != 3 {
		t.Errorf("Expected 3 books, got %d", len(data))
	}
}
