package order

import (
	"log"
	"net/http"
)

func InitOrderRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", getOrderListHandler)
	// r.Get("/{tagId}", getOrderByIDHandler)
	// r.Post("/", postOrderHandler)
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

func getOrderList() (*[]schema.Order, error) {
	var tagsList = make([]schema.Order, 0)
	for _, value := range tagsMap.ByID {
		if !value.Disabled {
			tagsList = append(tagsList, value)
		}
	}
	return &tagsList, nil
}
