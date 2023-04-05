package order

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"

	"github.com/SkaisgirisMarius/order-processor.git/helper"
)

func InitOrderRouter(db *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Get("/", getOrderListHandler(db))
	// r.Get("/{tagId}", getOrderByIDHandler)
	r.Post("/", postOrderHandler(db))
	return r
}

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

func postOrderHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order Order
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			helper.SendJsonError(w, http.StatusBadRequest, err)
			return
		}

		if err := db.Create(&order).Error; err != nil {
			helper.SendJsonError(w, http.StatusInternalServerError, err)
			return
		}

		helper.SendJsonOk(w, "Created")
	}
}
