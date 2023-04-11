package order

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TestGetOrderListHandler is a quick test that runs in sqlite to not create a new mySQL server just for tests
func TestGetOrderListHandler(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	if err = db.AutoMigrate(&Order{}); err != nil {
		t.Fatalf("Failed to auto-migrate database schema: %v", err)
	}
	db.Create(&Order{ProxyCount: 10, Name: "Test Order"})
	db.Create(&Order{ProxyCount: 20, Name: "Test Order2"})

	req := httptest.NewRequest(http.MethodGet, "/api/order", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/api/order", getOrderListHandler(db))

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var orders []Order
	err = json.Unmarshal(w.Body.Bytes(), &orders)
	assert.NoError(t, err)
	assert.Len(t, orders, 2)
	assert.Equal(t, "Test Order", orders[0].Name)
	assert.Equal(t, int64(10), orders[0].ProxyCount)
	assert.Equal(t, "Test Order2", orders[1].Name)
	assert.Equal(t, int64(20), orders[1].ProxyCount)
}

// This is a quick test that runs in sqlite to not create a new mySQL server just for tests
func TestPostOrderHandler(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to create database connection: %v", err)
	}

	if err = db.AutoMigrate(&Order{}); err != nil {
		t.Fatalf("Failed to auto-migrate database schema: %v", err)
	}

	newOrder := Order{ProxyCount: 10, Name: "Test Order"}
	payload, err := json.Marshal(newOrder)
	if err != nil {
		t.Fatalf("Failed to marshal JSON payload: %v", err)
	}
	req := httptest.NewRequest("POST", "/orders", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	postOrderHandler(db)(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected status code: got %v, want %v", w.Code, http.StatusOK)
	}

	var order Order
	if err = json.NewDecoder(w.Body).Decode(&order); err != nil {
		t.Fatalf("Failed to decode response JSON: %v", err)
	}

	if order.ID == 0 {
		t.Error("Expected order ID to be non-zero")
	}
	if order.ProxyCount != newOrder.ProxyCount {
		t.Errorf("Unexpected order proxy count: got %v, want %v", order.ProxyCount, newOrder.ProxyCount)
	}
	if order.Name != newOrder.Name {
		t.Errorf("Unexpected order name: got %v, want %v", order.Name, newOrder.Name)
	}
}
