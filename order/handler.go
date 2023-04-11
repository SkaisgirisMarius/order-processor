package order

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"

	"github.com/SkaisgirisMarius/order-processor.git/helper"
)

// InitOrderRouter initializes the order router and returns a http.Handler
func InitOrderRouter(db *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Get("/", getOrderListHandler(db))
	r.Get("/{orderID}", getOrderByIDHandler(db))
	r.Post("/", postOrderHandler(db))
	return r
}

// getOrderListHandler returns a http.HandlerFunc that handles GET requests for order list
func getOrderListHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var orders []Order
		if err := db.Find(&orders).Error; err != nil {
			helper.SendJsonError(w, http.StatusInternalServerError, err)
			return
		}
		helper.SendJsonOk(w, orders)
	}
}

// getOrderByIDHandler returns a http.HandlerFunc that handles GET requests for a specific order
func getOrderByIDHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderID := chi.URLParam(r, "orderID")
		var order Order
		if err := db.Where("id = ?", orderID).First(&order).Error; err != nil {
			if strings.Contains(err.Error(), "record not found") {
				helper.SendJsonError(w, http.StatusNotFound, fmt.Errorf("order with ID %s not found", orderID))
			} else {
				helper.SendJsonError(w, http.StatusInternalServerError, err)
			}
			return
		}
		helper.SendJsonOk(w, order)
	}
}

// postOrderHandler returns a http.HandlerFunc that handles POST requests to create a new order
func postOrderHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order Order
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			helper.SendJsonError(w, http.StatusBadRequest, err)
			return
		}

		// Create the new order
		if err := db.Create(&order).Error; err != nil {
			helper.SendJsonError(w, http.StatusInternalServerError, err)
			return
		}

		// Fetch the newly created order
		if err := db.First(&order, order.ID).Error; err != nil {
			helper.SendJsonError(w, http.StatusInternalServerError, err)
			return
		}

		helper.SendJsonOk(w, order)
	}
}
