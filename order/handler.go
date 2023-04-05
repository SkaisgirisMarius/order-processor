package order

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"

	"github.com/SkaisgirisMarius/order-processor.git/helper"
)

func InitOrderRouter(db *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Get("/", getOrderListHandler)
	// r.Get("/{tagId}", getOrderByIDHandler)
	r.Post("/", postOrderHandler)
	return r
}

func getOrderListHandler(w http.ResponseWriter, r *http.Request) {
	tags, err := getOrderList()
	if err != nil {
		log.Println(err)
		helper.SendJsonError(w, http.StatusInternalServerError, err)
		return
	}
	helper.SendJsonOk(w, tags)
}

func getOrderList() (*[]Order, error) {
	var tagsList = make([]Order, 0)
	for _, value := range tagsMap.ByID {
		if !value.Disabled {
			tagsList = append(tagsList, value)
		}
	}
	return &tagsList, nil
}

func postOrderHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order Order
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := db.Create(&order).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
